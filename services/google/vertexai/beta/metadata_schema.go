// Copyright 2022 Google LLC. All Rights Reserved.
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
package beta

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type MetadataSchema struct {
	Name          *string                       `json:"name"`
	SchemaVersion *string                       `json:"schemaVersion"`
	Schema        *string                       `json:"schema"`
	SchemaType    *MetadataSchemaSchemaTypeEnum `json:"schemaType"`
	CreateTime    *string                       `json:"createTime"`
	Project       *string                       `json:"project"`
	Location      *string                       `json:"location"`
	MetadataStore *string                       `json:"metadataStore"`
}

func (r *MetadataSchema) String() string {
	return dcl.SprintResource(r)
}

// The enum MetadataSchemaSchemaTypeEnum.
type MetadataSchemaSchemaTypeEnum string

// MetadataSchemaSchemaTypeEnumRef returns a *MetadataSchemaSchemaTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func MetadataSchemaSchemaTypeEnumRef(s string) *MetadataSchemaSchemaTypeEnum {
	v := MetadataSchemaSchemaTypeEnum(s)
	return &v
}

func (v MetadataSchemaSchemaTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"METADATA_SCHEMA_TYPE_UNSPECIFIED", "ARTIFACT_TYPE", "EXECUTION_TYPE", "CONTEXT_TYPE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "MetadataSchemaSchemaTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *MetadataSchema) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "vertex_ai",
		Type:    "MetadataSchema",
		Version: "beta",
	}
}

func (r *MetadataSchema) ID() (string, error) {
	if err := extractMetadataSchemaFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":           dcl.ValueOrEmptyString(nr.Name),
		"schema_version": dcl.ValueOrEmptyString(nr.SchemaVersion),
		"schema":         dcl.ValueOrEmptyString(nr.Schema),
		"schema_type":    dcl.ValueOrEmptyString(nr.SchemaType),
		"create_time":    dcl.ValueOrEmptyString(nr.CreateTime),
		"project":        dcl.ValueOrEmptyString(nr.Project),
		"location":       dcl.ValueOrEmptyString(nr.Location),
		"metadata_store": dcl.ValueOrEmptyString(nr.MetadataStore),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/metadataStores/{{metadata_store}}/metadataSchemas/{{name}}", params), nil
}

const MetadataSchemaMaxPage = -1

type MetadataSchemaList struct {
	Items []*MetadataSchema

	nextToken string

	pageSize int32

	resource *MetadataSchema
}

func (l *MetadataSchemaList) HasNext() bool {
	return l.nextToken != ""
}

func (l *MetadataSchemaList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listMetadataSchema(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListMetadataSchema(ctx context.Context, project, location, metadataStore string) (*MetadataSchemaList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListMetadataSchemaWithMaxResults(ctx, project, location, metadataStore, MetadataSchemaMaxPage)

}

func (c *Client) ListMetadataSchemaWithMaxResults(ctx context.Context, project, location, metadataStore string, pageSize int32) (*MetadataSchemaList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &MetadataSchema{
		Project:       &project,
		Location:      &location,
		MetadataStore: &metadataStore,
	}
	items, token, err := c.listMetadataSchema(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &MetadataSchemaList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetMetadataSchema(ctx context.Context, r *MetadataSchema) (*MetadataSchema, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractMetadataSchemaFields(r)

	b, err := c.getMetadataSchemaRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalMetadataSchema(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.MetadataStore = r.MetadataStore
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeMetadataSchemaNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractMetadataSchemaFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) ApplyMetadataSchema(ctx context.Context, rawDesired *MetadataSchema, opts ...dcl.ApplyOption) (*MetadataSchema, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *MetadataSchema
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyMetadataSchemaHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyMetadataSchemaHelper(c *Client, ctx context.Context, rawDesired *MetadataSchema, opts ...dcl.ApplyOption) (*MetadataSchema, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyMetadataSchema...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractMetadataSchemaFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.metadataSchemaDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToMetadataSchemaDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []metadataSchemaApiOperation
	if create {
		ops = append(ops, &createMetadataSchemaOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyMetadataSchemaDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyMetadataSchemaDiff(c *Client, ctx context.Context, desired *MetadataSchema, rawDesired *MetadataSchema, ops []metadataSchemaApiOperation, opts ...dcl.ApplyOption) (*MetadataSchema, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetMetadataSchema(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createMetadataSchemaOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapMetadataSchema(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeMetadataSchemaNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeMetadataSchemaNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeMetadataSchemaDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractMetadataSchemaFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractMetadataSchemaFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffMetadataSchema(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
