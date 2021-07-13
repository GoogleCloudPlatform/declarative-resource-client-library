// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package storage contains handwritten support code for the storage service.
package storage

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/textproto"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// expandStorageBucketLifecycleWithState expands the with_state enum into the
// ternary boolean is_live. It can be true, false, or unset each corresponding
// to a different value.
func expandStorageBucketLifecycleWithState(_ *BucketLifecycleRuleCondition, v *BucketLifecycleRuleConditionWithStateEnum) (interface{}, error) {
	switch *v {
	case "LIVE":
		b := true
		return &b, nil
	case "ARCHIVED":
		b := false
		return &b, nil
	case "ANY":
		return nil, nil
	}

	return nil, fmt.Errorf("unrecognized BucketLifecycleRuleConditionWithStateEnum value: %v", v)
}

// flattenStorageBucketLifecycleWithState flattens the ternary boolean is_live
// into the with_state enum. It can be true, false, or unset each corresponding
// to a different value.
func flattenStorageBucketLifecycleWithState(v interface{}) *BucketLifecycleRuleConditionWithStateEnum {
	b, ok := v.(bool)
	if !ok { // b is unset
		return BucketLifecycleRuleConditionWithStateEnumRef("ANY")
	}

	if b {
		return BucketLifecycleRuleConditionWithStateEnumRef("LIVE")
	}

	// b is false
	return BucketLifecycleRuleConditionWithStateEnumRef("ARCHIVED")
}

// EncodeStorageHmacKeyCreateRequest encapsulates fields into a metadata {}
// block, as expected by https://cloud.google.com/storage/docs/json_api/v1/projects/hmacKeys
// Note that HmacKeys is an unusual API- the JSON body shown as the "main" body
// is only sent/returned on Create, and the "storage#hmacKeyMetadata" type is
// the one recorded in serviceprotos.
func EncodeStorageHmacKeyCreateRequest(m map[string]interface{}) map[string]interface{} {
	req := make(map[string]interface{})
	dcl.PutMapEntry(req, []string{"metadata"}, m)
	return req
}

// do creates an HmacKey using the storage API. A custom method is required because the initial
// response nests some fields under metadata object.
func (op *createHmacKeyOperation) do(ctx context.Context, r *HmacKey, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project, serviceAccountEmail := r.createFields()
	u, err := hmacKeyCreateURL(c.Config.BasePath, project, serviceAccountEmail)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	// Resource fields are nested under metadata.
	metadata, ok := o["metadata"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("expected metadata to be an object, was %T", metadata)
	}
	// Secret is not nested under metadata.
	secret, ok := o["secret"].(string)
	if !ok {
		return fmt.Errorf("expected secret to be a string, was %T", secret)
	}
	metadata["secret"] = secret
	op.response = metadata
	// Include Name in URL substitution for initial GET request.
	accessID, ok := op.response["accessId"].(string)
	if !ok {
		return fmt.Errorf("expected accessID to be a string, was %T", accessID)
	}
	r.Name = &accessID

	// Poll for the HmacKey resource to be created. HmacKey resources are eventually consistent but do not support operations
	// so we must repeatedly poll to check for their creation.
	return dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		u, err := hmacKeyGetURL(c.Config.BasePath, r.URLNormalized())
		if err != nil {
			return nil, err
		}
		getResp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, nil)
		if err != nil {
			// If the error is a transient server error (e.g., 500) or not found (i.e., the resource has not yet been created),
			// continue retrying until the transient error is resolved, the resource is created, or we time out.
			if dcl.IsRetryableRequestError(c.Config, err, true) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		getResp.Response.Body.Close()
		return getResp, nil
	}, c.Config.RetryProvider)
}

// The HmacKey is already effectively deleted if it's in DELETED state.
func hmacKeyDeletePrecondition(r *HmacKey) bool {
	return *r.State == *HmacKeyStateEnumRef("DELETED")
}

// HmacKeyEnsureStateInactive ensures that an HmacKey resource is in the
// INACTIVE state prior to deletion.
func (r *HmacKey) HmacKeyEnsureStateInactive(ctx context.Context, c *Client) error {
	inactiveState := HmacKeyStateEnumRef("INACTIVE")
	if r.State == inactiveState {
		return nil
	}

	// otherwise, set to INACTIVE
	r.State = inactiveState
	c.Config.Logger.Infof("Performing sub-apply to reach desired state 'INACTIVE' for HmacKey: %v", r)
	if _, err := c.ApplyHmacKey(ctx, r); err != nil {
		return err
	}

	return nil
}

// retrieves metadata.accessID value
func retrieveAccessIDFromJSON(m map[string]interface{}) (*string, error) {
	metadata, ok := m["metadata"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("metadata key not found")
	}

	accessID, ok := metadata["accessId"].(string)
	if !ok {
		return nil, fmt.Errorf("metadata.AccessId not found")
	}

	return &accessID, nil
}

// returns url for object create request
func objectCreateURL(userBasePath, name, bucket string) (string, error) {
	params := map[string]interface{}{
		"name":   name,
		"bucket": bucket,
	}
	return dcl.URL("b/{{bucket}}/o?uploadType=multipart&name={{name}}", "https://storage.googleapis.com/upload/storage/v1/", userBasePath, params), nil
}

func (op *createObjectOperation) do(ctx context.Context, r *Object, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	bucket, name := r.createFields()
	u, err := objectCreateURL(c.Config.BasePath, name, bucket)

	if err != nil {
		return err
	}

	metadata, err := r.marshal(c)
	if err != nil {
		return err
	}
	buffer := bytes.Buffer{}
	writer := multipart.NewWriter(&buffer)
	header := textproto.MIMEHeader{}
	header.Add("Content-Type", mime.FormatMediaType("application/json", map[string]string{"charset": "UTF-8"}))
	writer.CreatePart(header)
	buffer.Write(metadata)
	header.Set("Content-Type", dcl.ValueOrEmptyString(r.ContentType))
	writer.CreatePart(header)
	buffer.Write([]byte(dcl.ValueOrEmptyString(r.Content)))
	writer.Close()
	req := buffer.Bytes()

	var config *dcl.Config
	if r.CustomerEncryption == nil {
		config = c.Config.Clone(dcl.WithContentType(fmt.Sprintf("multipart/related; boundary=%s", writer.Boundary())))
	} else {
		config = c.Config.Clone(
			dcl.WithHeader("X-Goog-Encryption-Algorithm", dcl.ValueOrEmptyString(r.CustomerEncryption.EncryptionAlgorithm)),
			dcl.WithHeader("X-Goog-Encryption-Key", dcl.ValueOrEmptyString(r.CustomerEncryption.Key)),
			dcl.WithHeader("X-Goog-Encryption-Key-Sha256", dcl.ValueOrEmptyString(r.CustomerEncryption.KeySha256)),
			dcl.WithContentType(fmt.Sprintf("multipart/related; boundary=%s", writer.Boundary())),
		)
	}

	resp, err := dcl.SendRequest(ctx, config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	c.Config.Logger.Infof("Successfully waited for operation")

	o, err := dcl.ResponseBodyAsJSON(resp)

	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetObject(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getObjectRaw(ctx context.Context, r *Object) ([]byte, error) {

	u, err := objectGetURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) getObjectContentRaw(ctx context.Context, r *Object) ([]byte, error) {

	u, err := objectGetURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return nil, err
	}
	var config *dcl.Config
	if r.CustomerEncryption == nil {
		config = c.Config.Clone(dcl.WithQueryParams(map[string]string{"alt": "media"}))
	} else {
		config = c.Config.Clone(
			dcl.WithHeader("X-Goog-Encryption-Algorithm", dcl.ValueOrEmptyString(r.CustomerEncryption.EncryptionAlgorithm)),
			dcl.WithHeader("X-Goog-Encryption-Key", dcl.ValueOrEmptyString(r.CustomerEncryption.Key)),
			dcl.WithHeader("X-Goog-Encryption-Key-Sha256", dcl.ValueOrEmptyString(r.CustomerEncryption.KeySha256)),
			dcl.WithQueryParams(map[string]string{"alt": "media"}),
		)
	}
	resp, err := dcl.SendRequest(ctx, config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) GetObject(ctx context.Context, r *Object) (*Object, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0))
	defer cancel()

	b, err := c.getObjectRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalObject(b, c)
	if err != nil {
		return nil, err
	}
	b, err = c.getObjectContentRaw(ctx, r)
	if err != nil {
		return nil, err
	}
	content := string(b)
	result.Content = &content
	result.Bucket = r.Bucket
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeObjectNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}
