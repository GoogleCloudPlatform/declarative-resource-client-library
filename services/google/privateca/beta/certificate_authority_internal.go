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
package beta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *CertificateAuthority) validate() error {

	if err := dcl.Required(r, "type"); err != nil {
		return err
	}
	if err := dcl.Required(r, "config"); err != nil {
		return err
	}
	if err := dcl.Required(r, "lifetime"); err != nil {
		return err
	}
	if err := dcl.Required(r, "keySpec"); err != nil {
		return err
	}
	if err := dcl.Required(r, "tier"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Config) {
		if err := r.Config.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.KeySpec) {
		if err := r.KeySpec.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SubordinateConfig) {
		if err := r.SubordinateConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AccessUrls) {
		if err := r.AccessUrls.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CertificatePolicy) {
		if err := r.CertificatePolicy.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.IssuingOptions) {
		if err := r.IssuingOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityConfig) validate() error {
	if err := dcl.Required(r, "subjectConfig"); err != nil {
		return err
	}
	if err := dcl.Required(r, "reusableConfig"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.SubjectConfig) {
		if err := r.SubjectConfig.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PublicKey) {
		if err := r.PublicKey.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ReusableConfig) {
		if err := r.ReusableConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityConfigSubjectConfig) validate() error {
	if err := dcl.Required(r, "subject"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Subject) {
		if err := r.Subject.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SubjectAltName) {
		if err := r.SubjectAltName.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityConfigSubjectConfigSubject) validate() error {
	return nil
}
func (r *CertificateAuthorityConfigSubjectConfigSubjectAltName) validate() error {
	return nil
}
func (r *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) validate() error {
	if err := dcl.Required(r, "objectId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "critical"); err != nil {
		return err
	}
	if err := dcl.Required(r, "value"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ObjectId) {
		if err := r.ObjectId.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) validate() error {
	if err := dcl.Required(r, "objectIdPath"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityConfigPublicKey) validate() error {
	if err := dcl.Required(r, "key"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityConfigReusableConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ReusableConfigValues) {
		if err := r.ReusableConfigValues.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityConfigReusableConfigReusableConfigValues) validate() error {
	if !dcl.IsEmptyValueIndirect(r.KeyUsage) {
		if err := r.KeyUsage.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CaOptions) {
		if err := r.CaOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage) validate() error {
	if !dcl.IsEmptyValueIndirect(r.BaseKeyUsage) {
		if err := r.BaseKeyUsage.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExtendedKeyUsage) {
		if err := r.ExtendedKeyUsage.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage) validate() error {
	return nil
}
func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage) validate() error {
	return nil
}
func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) validate() error {
	if err := dcl.Required(r, "objectIdPath"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions) validate() error {
	return nil
}
func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds) validate() error {
	if err := dcl.Required(r, "objectIdPath"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions) validate() error {
	if err := dcl.Required(r, "objectId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "critical"); err != nil {
		return err
	}
	if err := dcl.Required(r, "value"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ObjectId) {
		if err := r.ObjectId.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId) validate() error {
	if err := dcl.Required(r, "objectIdPath"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityKeySpec) validate() error {
	return nil
}
func (r *CertificateAuthoritySubordinateConfig) validate() error {
	if err := dcl.Required(r, "certificateAuthority"); err != nil {
		return err
	}
	if err := dcl.Required(r, "pemIssuerChain"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.PemIssuerChain) {
		if err := r.PemIssuerChain.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthoritySubordinateConfigPemIssuerChain) validate() error {
	if err := dcl.Required(r, "pemCertificates"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptions) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SubjectDescription) {
		if err := r.SubjectDescription.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PublicKey) {
		if err := r.PublicKey.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SubjectKeyId) {
		if err := r.SubjectKeyId.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AuthorityKeyId) {
		if err := r.AuthorityKeyId.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CertFingerprint) {
		if err := r.CertFingerprint.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ConfigValues) {
		if err := r.ConfigValues.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescription) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Subject) {
		if err := r.Subject.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SubjectAltName) {
		if err := r.SubjectAltName.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) validate() error {
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) validate() error {
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) validate() error {
	if err := dcl.Required(r, "objectId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "critical"); err != nil {
		return err
	}
	if err := dcl.Required(r, "value"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ObjectId) {
		if err := r.ObjectId.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) validate() error {
	if err := dcl.Required(r, "objectIdPath"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsPublicKey) validate() error {
	if err := dcl.Required(r, "key"); err != nil {
		return err
	}
	if err := dcl.Required(r, "format"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) validate() error {
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) validate() error {
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsCertFingerprint) validate() error {
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsConfigValues) validate() error {
	if !dcl.IsEmptyValueIndirect(r.KeyUsage) {
		if err := r.KeyUsage.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CaOptions) {
		if err := r.CaOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage) validate() error {
	if !dcl.IsEmptyValueIndirect(r.BaseKeyUsage) {
		if err := r.BaseKeyUsage.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExtendedKeyUsage) {
		if err := r.ExtendedKeyUsage.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage) validate() error {
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage) validate() error {
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages) validate() error {
	if err := dcl.Required(r, "objectIdPath"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions) validate() error {
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds) validate() error {
	if err := dcl.Required(r, "objectIdPath"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions) validate() error {
	if err := dcl.Required(r, "objectId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "critical"); err != nil {
		return err
	}
	if err := dcl.Required(r, "value"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ObjectId) {
		if err := r.ObjectId.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId) validate() error {
	if err := dcl.Required(r, "objectIdPath"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityAccessUrls) validate() error {
	return nil
}
func (r *CertificateAuthorityCertificatePolicy) validate() error {
	if !dcl.IsEmptyValueIndirect(r.AllowedConfigList) {
		if err := r.AllowedConfigList.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.OverwriteConfigValues) {
		if err := r.OverwriteConfigValues.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AllowedSans) {
		if err := r.AllowedSans.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AllowedIssuanceModes) {
		if err := r.AllowedIssuanceModes.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyAllowedConfigList) validate() error {
	if err := dcl.Required(r, "allowedConfigValues"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ReusableConfigValues) {
		if err := r.ReusableConfigValues.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues) validate() error {
	if !dcl.IsEmptyValueIndirect(r.KeyUsage) {
		if err := r.KeyUsage.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CaOptions) {
		if err := r.CaOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage) validate() error {
	if !dcl.IsEmptyValueIndirect(r.BaseKeyUsage) {
		if err := r.BaseKeyUsage.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExtendedKeyUsage) {
		if err := r.ExtendedKeyUsage.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) validate() error {
	return nil
}
func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) validate() error {
	return nil
}
func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) validate() error {
	if err := dcl.Required(r, "objectIdPath"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions) validate() error {
	return nil
}
func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds) validate() error {
	if err := dcl.Required(r, "objectIdPath"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions) validate() error {
	if err := dcl.Required(r, "objectId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "critical"); err != nil {
		return err
	}
	if err := dcl.Required(r, "value"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ObjectId) {
		if err := r.ObjectId.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) validate() error {
	if err := dcl.Required(r, "objectIdPath"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValues) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ReusableConfigValues) {
		if err := r.ReusableConfigValues.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues) validate() error {
	if !dcl.IsEmptyValueIndirect(r.KeyUsage) {
		if err := r.KeyUsage.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CaOptions) {
		if err := r.CaOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage) validate() error {
	if !dcl.IsEmptyValueIndirect(r.BaseKeyUsage) {
		if err := r.BaseKeyUsage.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExtendedKeyUsage) {
		if err := r.ExtendedKeyUsage.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) validate() error {
	return nil
}
func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) validate() error {
	return nil
}
func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) validate() error {
	if err := dcl.Required(r, "objectIdPath"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions) validate() error {
	return nil
}
func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds) validate() error {
	if err := dcl.Required(r, "objectIdPath"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions) validate() error {
	if err := dcl.Required(r, "objectId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "critical"); err != nil {
		return err
	}
	if err := dcl.Required(r, "value"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ObjectId) {
		if err := r.ObjectId.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) validate() error {
	if err := dcl.Required(r, "objectIdPath"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations) validate() error {
	return nil
}
func (r *CertificateAuthorityCertificatePolicyAllowedSans) validate() error {
	return nil
}
func (r *CertificateAuthorityCertificatePolicyAllowedIssuanceModes) validate() error {
	if err := dcl.Required(r, "allowCsrBasedIssuance"); err != nil {
		return err
	}
	if err := dcl.Required(r, "allowConfigBasedIssuance"); err != nil {
		return err
	}
	return nil
}
func (r *CertificateAuthorityIssuingOptions) validate() error {
	if err := dcl.Required(r, "includeCaCertUrl"); err != nil {
		return err
	}
	if err := dcl.Required(r, "includeCrlAccessUrl"); err != nil {
		return err
	}
	return nil
}

func certificateAuthorityGetURL(userBasePath string, r *CertificateAuthority) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/certificateAuthorities/{{name}}", "https://privateca.googleapis.com/v1beta1/", userBasePath, params), nil
}

func certificateAuthorityListURL(userBasePath, project, location string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/certificateAuthorities", "https://privateca.googleapis.com/v1beta1/", userBasePath, params), nil

}

func certificateAuthorityCreateURL(userBasePath, project, location, name string) (string, error) {
	params := map[string]interface{}{
		"project":  project,
		"location": location,
		"name":     name,
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/certificateAuthorities?certificateAuthorityId={{name}}", "https://privateca.googleapis.com/v1beta1/", userBasePath, params), nil

}

func certificateAuthorityDeleteURL(userBasePath string, r *CertificateAuthority) (string, error) {
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(r.Project),
		"location": dcl.ValueOrEmptyString(r.Location),
		"name":     dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/certificateAuthorities/{{name}}:scheduleDelete", "https://privateca.googleapis.com/v1beta1/", userBasePath, params), nil
}

// certificateAuthorityApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type certificateAuthorityApiOperation interface {
	do(context.Context, *CertificateAuthority, *Client) error
}

func (c *Client) listCertificateAuthorityRaw(ctx context.Context, project, location, pageToken string, pageSize int32) ([]byte, error) {
	u, err := certificateAuthorityListURL(c.Config.BasePath, project, location)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != CertificateAuthorityMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listCertificateAuthorityOperation struct {
	CertificateAuthorities []map[string]interface{} `json:"certificateAuthorities"`
	Token                  string                   `json:"nextPageToken"`
}

func (c *Client) listCertificateAuthority(ctx context.Context, project, location, pageToken string, pageSize int32) ([]*CertificateAuthority, string, error) {
	b, err := c.listCertificateAuthorityRaw(ctx, project, location, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listCertificateAuthorityOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*CertificateAuthority
	for _, v := range m.CertificateAuthorities {
		res, err := unmarshalMapCertificateAuthority(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		res.Location = &location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllCertificateAuthority(ctx context.Context, f func(*CertificateAuthority) bool, resources []*CertificateAuthority) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteCertificateAuthority(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteCertificateAuthorityOperation struct{}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createCertificateAuthorityOperation struct {
	response map[string]interface{}
}

func (op *createCertificateAuthorityOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) getCertificateAuthorityRaw(ctx context.Context, r *CertificateAuthority) ([]byte, error) {

	u, err := certificateAuthorityGetURL(c.Config.BasePath, r.URLNormalized())
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

func (c *Client) certificateAuthorityDiffsForRawDesired(ctx context.Context, rawDesired *CertificateAuthority, opts ...dcl.ApplyOption) (initial, desired *CertificateAuthority, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *CertificateAuthority
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*CertificateAuthority); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected CertificateAuthority, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetCertificateAuthority(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a CertificateAuthority resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve CertificateAuthority resource: %v", err)
		}
		c.Config.Logger.Info("Found that CertificateAuthority resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeCertificateAuthorityDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for CertificateAuthority: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for CertificateAuthority: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeCertificateAuthorityInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for CertificateAuthority: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeCertificateAuthorityDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for CertificateAuthority: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffCertificateAuthority(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeCertificateAuthorityInitialState(rawInitial, rawDesired *CertificateAuthority) (*CertificateAuthority, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeCertificateAuthorityDesiredState(rawDesired, rawInitial *CertificateAuthority, opts ...dcl.ApplyOption) (*CertificateAuthority, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Config = canonicalizeCertificateAuthorityConfig(rawDesired.Config, nil, opts...)
		rawDesired.KeySpec = canonicalizeCertificateAuthorityKeySpec(rawDesired.KeySpec, nil, opts...)
		rawDesired.SubordinateConfig = canonicalizeCertificateAuthoritySubordinateConfig(rawDesired.SubordinateConfig, nil, opts...)
		rawDesired.AccessUrls = canonicalizeCertificateAuthorityAccessUrls(rawDesired.AccessUrls, nil, opts...)
		rawDesired.CertificatePolicy = canonicalizeCertificateAuthorityCertificatePolicy(rawDesired.CertificatePolicy, nil, opts...)
		rawDesired.IssuingOptions = canonicalizeCertificateAuthorityIssuingOptions(rawDesired.IssuingOptions, nil, opts...)

		return rawDesired, nil
	}

	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		rawDesired.Name = rawInitial.Name
	}
	if dcl.IsZeroValue(rawDesired.Type) {
		rawDesired.Type = rawInitial.Type
	}
	rawDesired.Config = canonicalizeCertificateAuthorityConfig(rawDesired.Config, rawInitial.Config, opts...)
	if dcl.StringCanonicalize(rawDesired.Lifetime, rawInitial.Lifetime) {
		rawDesired.Lifetime = rawInitial.Lifetime
	}
	rawDesired.KeySpec = canonicalizeCertificateAuthorityKeySpec(rawDesired.KeySpec, rawInitial.KeySpec, opts...)
	rawDesired.SubordinateConfig = canonicalizeCertificateAuthoritySubordinateConfig(rawDesired.SubordinateConfig, rawInitial.SubordinateConfig, opts...)
	if dcl.IsZeroValue(rawDesired.Tier) {
		rawDesired.Tier = rawInitial.Tier
	}
	if dcl.StringCanonicalize(rawDesired.GcsBucket, rawInitial.GcsBucket) {
		rawDesired.GcsBucket = rawInitial.GcsBucket
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		rawDesired.Labels = rawInitial.Labels
	}
	rawDesired.CertificatePolicy = canonicalizeCertificateAuthorityCertificatePolicy(rawDesired.CertificatePolicy, rawInitial.CertificatePolicy, opts...)
	rawDesired.IssuingOptions = canonicalizeCertificateAuthorityIssuingOptions(rawDesired.IssuingOptions, rawInitial.IssuingOptions, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		rawDesired.Project = rawInitial.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		rawDesired.Location = rawInitial.Location
	}

	return rawDesired, nil
}

func canonicalizeCertificateAuthorityNewState(c *Client, rawNew, rawDesired *CertificateAuthority) (*CertificateAuthority, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Type) && dcl.IsEmptyValueIndirect(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Config) && dcl.IsEmptyValueIndirect(rawDesired.Config) {
		rawNew.Config = rawDesired.Config
	} else {
		rawNew.Config = canonicalizeNewCertificateAuthorityConfig(c, rawDesired.Config, rawNew.Config)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Lifetime) && dcl.IsEmptyValueIndirect(rawDesired.Lifetime) {
		rawNew.Lifetime = rawDesired.Lifetime
	} else {
		if dcl.StringCanonicalize(rawDesired.Lifetime, rawNew.Lifetime) {
			rawNew.Lifetime = rawDesired.Lifetime
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.KeySpec) && dcl.IsEmptyValueIndirect(rawDesired.KeySpec) {
		rawNew.KeySpec = rawDesired.KeySpec
	} else {
		rawNew.KeySpec = canonicalizeNewCertificateAuthorityKeySpec(c, rawDesired.KeySpec, rawNew.KeySpec)
	}

	if dcl.IsEmptyValueIndirect(rawNew.SubordinateConfig) && dcl.IsEmptyValueIndirect(rawDesired.SubordinateConfig) {
		rawNew.SubordinateConfig = rawDesired.SubordinateConfig
	} else {
		rawNew.SubordinateConfig = canonicalizeNewCertificateAuthoritySubordinateConfig(c, rawDesired.SubordinateConfig, rawNew.SubordinateConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Tier) && dcl.IsEmptyValueIndirect(rawDesired.Tier) {
		rawNew.Tier = rawDesired.Tier
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.PemCaCertificates) && dcl.IsEmptyValueIndirect(rawDesired.PemCaCertificates) {
		rawNew.PemCaCertificates = rawDesired.PemCaCertificates
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CaCertificateDescriptions) && dcl.IsEmptyValueIndirect(rawDesired.CaCertificateDescriptions) {
		rawNew.CaCertificateDescriptions = rawDesired.CaCertificateDescriptions
	} else {
		rawNew.CaCertificateDescriptions = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSlice(c, rawDesired.CaCertificateDescriptions, rawNew.CaCertificateDescriptions)
	}

	if dcl.IsEmptyValueIndirect(rawNew.GcsBucket) && dcl.IsEmptyValueIndirect(rawDesired.GcsBucket) {
		rawNew.GcsBucket = rawDesired.GcsBucket
	} else {
		if dcl.StringCanonicalize(rawDesired.GcsBucket, rawNew.GcsBucket) {
			rawNew.GcsBucket = rawDesired.GcsBucket
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AccessUrls) && dcl.IsEmptyValueIndirect(rawDesired.AccessUrls) {
		rawNew.AccessUrls = rawDesired.AccessUrls
	} else {
		rawNew.AccessUrls = canonicalizeNewCertificateAuthorityAccessUrls(c, rawDesired.AccessUrls, rawNew.AccessUrls)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DeleteTime) && dcl.IsEmptyValueIndirect(rawDesired.DeleteTime) {
		rawNew.DeleteTime = rawDesired.DeleteTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CertificatePolicy) && dcl.IsEmptyValueIndirect(rawDesired.CertificatePolicy) {
		rawNew.CertificatePolicy = rawDesired.CertificatePolicy
	} else {
		rawNew.CertificatePolicy = canonicalizeNewCertificateAuthorityCertificatePolicy(c, rawDesired.CertificatePolicy, rawNew.CertificatePolicy)
	}

	if dcl.IsEmptyValueIndirect(rawNew.IssuingOptions) && dcl.IsEmptyValueIndirect(rawDesired.IssuingOptions) {
		rawNew.IssuingOptions = rawDesired.IssuingOptions
	} else {
		rawNew.IssuingOptions = canonicalizeNewCertificateAuthorityIssuingOptions(c, rawDesired.IssuingOptions, rawNew.IssuingOptions)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeCertificateAuthorityConfig(des, initial *CertificateAuthorityConfig, opts ...dcl.ApplyOption) *CertificateAuthorityConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.SubjectConfig = canonicalizeCertificateAuthorityConfigSubjectConfig(des.SubjectConfig, initial.SubjectConfig, opts...)
	des.PublicKey = canonicalizeCertificateAuthorityConfigPublicKey(des.PublicKey, initial.PublicKey, opts...)
	des.ReusableConfig = canonicalizeCertificateAuthorityConfigReusableConfig(des.ReusableConfig, initial.ReusableConfig, opts...)

	return des
}

func canonicalizeNewCertificateAuthorityConfig(c *Client, des, nw *CertificateAuthorityConfig) *CertificateAuthorityConfig {
	if des == nil || nw == nil {
		return nw
	}

	nw.SubjectConfig = canonicalizeNewCertificateAuthorityConfigSubjectConfig(c, des.SubjectConfig, nw.SubjectConfig)
	nw.PublicKey = canonicalizeNewCertificateAuthorityConfigPublicKey(c, des.PublicKey, nw.PublicKey)
	nw.ReusableConfig = canonicalizeNewCertificateAuthorityConfigReusableConfig(c, des.ReusableConfig, nw.ReusableConfig)

	return nw
}

func canonicalizeNewCertificateAuthorityConfigSet(c *Client, des, nw []CertificateAuthorityConfig) []CertificateAuthorityConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigSlice(c *Client, des, nw []CertificateAuthorityConfig) []CertificateAuthorityConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfig(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigSubjectConfig(des, initial *CertificateAuthorityConfigSubjectConfig, opts ...dcl.ApplyOption) *CertificateAuthorityConfigSubjectConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Subject = canonicalizeCertificateAuthorityConfigSubjectConfigSubject(des.Subject, initial.Subject, opts...)
	if dcl.StringCanonicalize(des.CommonName, initial.CommonName) || dcl.IsZeroValue(des.CommonName) {
		des.CommonName = initial.CommonName
	}
	des.SubjectAltName = canonicalizeCertificateAuthorityConfigSubjectConfigSubjectAltName(des.SubjectAltName, initial.SubjectAltName, opts...)

	return des
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfig(c *Client, des, nw *CertificateAuthorityConfigSubjectConfig) *CertificateAuthorityConfigSubjectConfig {
	if des == nil || nw == nil {
		return nw
	}

	nw.Subject = canonicalizeNewCertificateAuthorityConfigSubjectConfigSubject(c, des.Subject, nw.Subject)
	if dcl.StringCanonicalize(des.CommonName, nw.CommonName) {
		nw.CommonName = des.CommonName
	}
	nw.SubjectAltName = canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltName(c, des.SubjectAltName, nw.SubjectAltName)

	return nw
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfigSet(c *Client, des, nw []CertificateAuthorityConfigSubjectConfig) []CertificateAuthorityConfigSubjectConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigSubjectConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigSubjectConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfigSlice(c *Client, des, nw []CertificateAuthorityConfigSubjectConfig) []CertificateAuthorityConfigSubjectConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigSubjectConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigSubjectConfig(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigSubjectConfigSubject(des, initial *CertificateAuthorityConfigSubjectConfigSubject, opts ...dcl.ApplyOption) *CertificateAuthorityConfigSubjectConfigSubject {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.CountryCode, initial.CountryCode) || dcl.IsZeroValue(des.CountryCode) {
		des.CountryCode = initial.CountryCode
	}
	if dcl.StringCanonicalize(des.Organization, initial.Organization) || dcl.IsZeroValue(des.Organization) {
		des.Organization = initial.Organization
	}
	if dcl.StringCanonicalize(des.OrganizationalUnit, initial.OrganizationalUnit) || dcl.IsZeroValue(des.OrganizationalUnit) {
		des.OrganizationalUnit = initial.OrganizationalUnit
	}
	if dcl.StringCanonicalize(des.Locality, initial.Locality) || dcl.IsZeroValue(des.Locality) {
		des.Locality = initial.Locality
	}
	if dcl.StringCanonicalize(des.Province, initial.Province) || dcl.IsZeroValue(des.Province) {
		des.Province = initial.Province
	}
	if dcl.StringCanonicalize(des.StreetAddress, initial.StreetAddress) || dcl.IsZeroValue(des.StreetAddress) {
		des.StreetAddress = initial.StreetAddress
	}
	if dcl.StringCanonicalize(des.PostalCode, initial.PostalCode) || dcl.IsZeroValue(des.PostalCode) {
		des.PostalCode = initial.PostalCode
	}

	return des
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfigSubject(c *Client, des, nw *CertificateAuthorityConfigSubjectConfigSubject) *CertificateAuthorityConfigSubjectConfigSubject {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.CountryCode, nw.CountryCode) {
		nw.CountryCode = des.CountryCode
	}
	if dcl.StringCanonicalize(des.Organization, nw.Organization) {
		nw.Organization = des.Organization
	}
	if dcl.StringCanonicalize(des.OrganizationalUnit, nw.OrganizationalUnit) {
		nw.OrganizationalUnit = des.OrganizationalUnit
	}
	if dcl.StringCanonicalize(des.Locality, nw.Locality) {
		nw.Locality = des.Locality
	}
	if dcl.StringCanonicalize(des.Province, nw.Province) {
		nw.Province = des.Province
	}
	if dcl.StringCanonicalize(des.StreetAddress, nw.StreetAddress) {
		nw.StreetAddress = des.StreetAddress
	}
	if dcl.StringCanonicalize(des.PostalCode, nw.PostalCode) {
		nw.PostalCode = des.PostalCode
	}

	return nw
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectSet(c *Client, des, nw []CertificateAuthorityConfigSubjectConfigSubject) []CertificateAuthorityConfigSubjectConfigSubject {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigSubjectConfigSubject
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigSubjectConfigSubjectNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectSlice(c *Client, des, nw []CertificateAuthorityConfigSubjectConfigSubject) []CertificateAuthorityConfigSubjectConfigSubject {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigSubjectConfigSubject
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigSubjectConfigSubject(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigSubjectConfigSubjectAltName(des, initial *CertificateAuthorityConfigSubjectConfigSubjectAltName, opts ...dcl.ApplyOption) *CertificateAuthorityConfigSubjectConfigSubjectAltName {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DnsNames) {
		des.DnsNames = initial.DnsNames
	}
	if dcl.IsZeroValue(des.Uris) {
		des.Uris = initial.Uris
	}
	if dcl.IsZeroValue(des.EmailAddresses) {
		des.EmailAddresses = initial.EmailAddresses
	}
	if dcl.IsZeroValue(des.IPAddresses) {
		des.IPAddresses = initial.IPAddresses
	}
	if dcl.IsZeroValue(des.CustomSans) {
		des.CustomSans = initial.CustomSans
	}

	return des
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltName(c *Client, des, nw *CertificateAuthorityConfigSubjectConfigSubjectAltName) *CertificateAuthorityConfigSubjectConfigSubjectAltName {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.DnsNames) {
		nw.DnsNames = des.DnsNames
	}
	if dcl.IsZeroValue(nw.Uris) {
		nw.Uris = des.Uris
	}
	if dcl.IsZeroValue(nw.EmailAddresses) {
		nw.EmailAddresses = des.EmailAddresses
	}
	if dcl.IsZeroValue(nw.IPAddresses) {
		nw.IPAddresses = des.IPAddresses
	}
	nw.CustomSans = canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansSlice(c, des.CustomSans, nw.CustomSans)

	return nw
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltNameSet(c *Client, des, nw []CertificateAuthorityConfigSubjectConfigSubjectAltName) []CertificateAuthorityConfigSubjectConfigSubjectAltName {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigSubjectConfigSubjectAltName
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigSubjectConfigSubjectAltNameNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltNameSlice(c *Client, des, nw []CertificateAuthorityConfigSubjectConfigSubjectAltName) []CertificateAuthorityConfigSubjectConfigSubjectAltName {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigSubjectConfigSubjectAltName
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltName(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(des, initial *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans, opts ...dcl.ApplyOption) *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.ObjectId = canonicalizeCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(des.ObjectId, initial.ObjectId, opts...)
	if dcl.BoolCanonicalize(des.Critical, initial.Critical) || dcl.IsZeroValue(des.Critical) {
		des.Critical = initial.Critical
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(c *Client, des, nw *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	if des == nil || nw == nil {
		return nw
	}

	nw.ObjectId = canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(c, des.ObjectId, nw.ObjectId)
	if dcl.BoolCanonicalize(des.Critical, nw.Critical) {
		nw.Critical = des.Critical
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansSet(c *Client, des, nw []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansSlice(c *Client, des, nw []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(des, initial *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId, opts ...dcl.ApplyOption) *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObjectIdPath) {
		des.ObjectIdPath = initial.ObjectIdPath
	}

	return des
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(c *Client, des, nw *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ObjectIdPath) {
		nw.ObjectIdPath = des.ObjectIdPath
	}

	return nw
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdSet(c *Client, des, nw []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdSlice(c *Client, des, nw []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigPublicKey(des, initial *CertificateAuthorityConfigPublicKey, opts ...dcl.ApplyOption) *CertificateAuthorityConfigPublicKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewCertificateAuthorityConfigPublicKey(c *Client, des, nw *CertificateAuthorityConfigPublicKey) *CertificateAuthorityConfigPublicKey {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.IsZeroValue(nw.Type) {
		nw.Type = des.Type
	}

	return nw
}

func canonicalizeNewCertificateAuthorityConfigPublicKeySet(c *Client, des, nw []CertificateAuthorityConfigPublicKey) []CertificateAuthorityConfigPublicKey {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigPublicKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigPublicKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigPublicKeySlice(c *Client, des, nw []CertificateAuthorityConfigPublicKey) []CertificateAuthorityConfigPublicKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigPublicKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigPublicKey(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigReusableConfig(des, initial *CertificateAuthorityConfigReusableConfig, opts ...dcl.ApplyOption) *CertificateAuthorityConfigReusableConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ReusableConfig, initial.ReusableConfig) || dcl.IsZeroValue(des.ReusableConfig) {
		des.ReusableConfig = initial.ReusableConfig
	}
	des.ReusableConfigValues = canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValues(des.ReusableConfigValues, initial.ReusableConfigValues, opts...)

	return des
}

func canonicalizeNewCertificateAuthorityConfigReusableConfig(c *Client, des, nw *CertificateAuthorityConfigReusableConfig) *CertificateAuthorityConfigReusableConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ReusableConfig, nw.ReusableConfig) {
		nw.ReusableConfig = des.ReusableConfig
	}
	nw.ReusableConfigValues = canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValues(c, des.ReusableConfigValues, nw.ReusableConfigValues)

	return nw
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigSet(c *Client, des, nw []CertificateAuthorityConfigReusableConfig) []CertificateAuthorityConfigReusableConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigReusableConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigReusableConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigSlice(c *Client, des, nw []CertificateAuthorityConfigReusableConfig) []CertificateAuthorityConfigReusableConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigReusableConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigReusableConfig(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValues(des, initial *CertificateAuthorityConfigReusableConfigReusableConfigValues, opts ...dcl.ApplyOption) *CertificateAuthorityConfigReusableConfigReusableConfigValues {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.KeyUsage = canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(des.KeyUsage, initial.KeyUsage, opts...)
	des.CaOptions = canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(des.CaOptions, initial.CaOptions, opts...)
	if dcl.IsZeroValue(des.PolicyIds) {
		des.PolicyIds = initial.PolicyIds
	}
	if dcl.IsZeroValue(des.AiaOcspServers) {
		des.AiaOcspServers = initial.AiaOcspServers
	}
	if dcl.IsZeroValue(des.AdditionalExtensions) {
		des.AdditionalExtensions = initial.AdditionalExtensions
	}

	return des
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValues(c *Client, des, nw *CertificateAuthorityConfigReusableConfigReusableConfigValues) *CertificateAuthorityConfigReusableConfigReusableConfigValues {
	if des == nil || nw == nil {
		return nw
	}

	nw.KeyUsage = canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(c, des.KeyUsage, nw.KeyUsage)
	nw.CaOptions = canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(c, des.CaOptions, nw.CaOptions)
	nw.PolicyIds = canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsSlice(c, des.PolicyIds, nw.PolicyIds)
	if dcl.IsZeroValue(nw.AiaOcspServers) {
		nw.AiaOcspServers = des.AiaOcspServers
	}
	nw.AdditionalExtensions = canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsSlice(c, des.AdditionalExtensions, nw.AdditionalExtensions)

	return nw
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesSet(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValues) []CertificateAuthorityConfigReusableConfigReusableConfigValues {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigReusableConfigReusableConfigValues
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigReusableConfigReusableConfigValuesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesSlice(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValues) []CertificateAuthorityConfigReusableConfigReusableConfigValues {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigReusableConfigReusableConfigValues
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValues(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(des, initial *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage, opts ...dcl.ApplyOption) *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.BaseKeyUsage = canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(des.BaseKeyUsage, initial.BaseKeyUsage, opts...)
	des.ExtendedKeyUsage = canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(des.ExtendedKeyUsage, initial.ExtendedKeyUsage, opts...)
	if dcl.IsZeroValue(des.UnknownExtendedKeyUsages) {
		des.UnknownExtendedKeyUsages = initial.UnknownExtendedKeyUsages
	}

	return des
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(c *Client, des, nw *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage) *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage {
	if des == nil || nw == nil {
		return nw
	}

	nw.BaseKeyUsage = canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(c, des.BaseKeyUsage, nw.BaseKeyUsage)
	nw.ExtendedKeyUsage = canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(c, des.ExtendedKeyUsage, nw.ExtendedKeyUsage)
	nw.UnknownExtendedKeyUsages = canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c, des.UnknownExtendedKeyUsages, nw.UnknownExtendedKeyUsages)

	return nw
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageSet(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage) []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageSlice(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage) []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(des, initial *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage, opts ...dcl.ApplyOption) *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.DigitalSignature, initial.DigitalSignature) || dcl.IsZeroValue(des.DigitalSignature) {
		des.DigitalSignature = initial.DigitalSignature
	}
	if dcl.BoolCanonicalize(des.ContentCommitment, initial.ContentCommitment) || dcl.IsZeroValue(des.ContentCommitment) {
		des.ContentCommitment = initial.ContentCommitment
	}
	if dcl.BoolCanonicalize(des.KeyEncipherment, initial.KeyEncipherment) || dcl.IsZeroValue(des.KeyEncipherment) {
		des.KeyEncipherment = initial.KeyEncipherment
	}
	if dcl.BoolCanonicalize(des.DataEncipherment, initial.DataEncipherment) || dcl.IsZeroValue(des.DataEncipherment) {
		des.DataEncipherment = initial.DataEncipherment
	}
	if dcl.BoolCanonicalize(des.KeyAgreement, initial.KeyAgreement) || dcl.IsZeroValue(des.KeyAgreement) {
		des.KeyAgreement = initial.KeyAgreement
	}
	if dcl.BoolCanonicalize(des.CertSign, initial.CertSign) || dcl.IsZeroValue(des.CertSign) {
		des.CertSign = initial.CertSign
	}
	if dcl.BoolCanonicalize(des.CrlSign, initial.CrlSign) || dcl.IsZeroValue(des.CrlSign) {
		des.CrlSign = initial.CrlSign
	}
	if dcl.BoolCanonicalize(des.EncipherOnly, initial.EncipherOnly) || dcl.IsZeroValue(des.EncipherOnly) {
		des.EncipherOnly = initial.EncipherOnly
	}
	if dcl.BoolCanonicalize(des.DecipherOnly, initial.DecipherOnly) || dcl.IsZeroValue(des.DecipherOnly) {
		des.DecipherOnly = initial.DecipherOnly
	}

	return des
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(c *Client, des, nw *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage) *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.DigitalSignature, nw.DigitalSignature) {
		nw.DigitalSignature = des.DigitalSignature
	}
	if dcl.BoolCanonicalize(des.ContentCommitment, nw.ContentCommitment) {
		nw.ContentCommitment = des.ContentCommitment
	}
	if dcl.BoolCanonicalize(des.KeyEncipherment, nw.KeyEncipherment) {
		nw.KeyEncipherment = des.KeyEncipherment
	}
	if dcl.BoolCanonicalize(des.DataEncipherment, nw.DataEncipherment) {
		nw.DataEncipherment = des.DataEncipherment
	}
	if dcl.BoolCanonicalize(des.KeyAgreement, nw.KeyAgreement) {
		nw.KeyAgreement = des.KeyAgreement
	}
	if dcl.BoolCanonicalize(des.CertSign, nw.CertSign) {
		nw.CertSign = des.CertSign
	}
	if dcl.BoolCanonicalize(des.CrlSign, nw.CrlSign) {
		nw.CrlSign = des.CrlSign
	}
	if dcl.BoolCanonicalize(des.EncipherOnly, nw.EncipherOnly) {
		nw.EncipherOnly = des.EncipherOnly
	}
	if dcl.BoolCanonicalize(des.DecipherOnly, nw.DecipherOnly) {
		nw.DecipherOnly = des.DecipherOnly
	}

	return nw
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageSet(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage) []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageSlice(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage) []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(des, initial *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage, opts ...dcl.ApplyOption) *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.ServerAuth, initial.ServerAuth) || dcl.IsZeroValue(des.ServerAuth) {
		des.ServerAuth = initial.ServerAuth
	}
	if dcl.BoolCanonicalize(des.ClientAuth, initial.ClientAuth) || dcl.IsZeroValue(des.ClientAuth) {
		des.ClientAuth = initial.ClientAuth
	}
	if dcl.BoolCanonicalize(des.CodeSigning, initial.CodeSigning) || dcl.IsZeroValue(des.CodeSigning) {
		des.CodeSigning = initial.CodeSigning
	}
	if dcl.BoolCanonicalize(des.EmailProtection, initial.EmailProtection) || dcl.IsZeroValue(des.EmailProtection) {
		des.EmailProtection = initial.EmailProtection
	}
	if dcl.BoolCanonicalize(des.TimeStamping, initial.TimeStamping) || dcl.IsZeroValue(des.TimeStamping) {
		des.TimeStamping = initial.TimeStamping
	}
	if dcl.BoolCanonicalize(des.OcspSigning, initial.OcspSigning) || dcl.IsZeroValue(des.OcspSigning) {
		des.OcspSigning = initial.OcspSigning
	}

	return des
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(c *Client, des, nw *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage) *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.ServerAuth, nw.ServerAuth) {
		nw.ServerAuth = des.ServerAuth
	}
	if dcl.BoolCanonicalize(des.ClientAuth, nw.ClientAuth) {
		nw.ClientAuth = des.ClientAuth
	}
	if dcl.BoolCanonicalize(des.CodeSigning, nw.CodeSigning) {
		nw.CodeSigning = des.CodeSigning
	}
	if dcl.BoolCanonicalize(des.EmailProtection, nw.EmailProtection) {
		nw.EmailProtection = des.EmailProtection
	}
	if dcl.BoolCanonicalize(des.TimeStamping, nw.TimeStamping) {
		nw.TimeStamping = des.TimeStamping
	}
	if dcl.BoolCanonicalize(des.OcspSigning, nw.OcspSigning) {
		nw.OcspSigning = des.OcspSigning
	}

	return nw
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageSet(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage) []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageSlice(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage) []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(des, initial *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages, opts ...dcl.ApplyOption) *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObjectIdPath) {
		des.ObjectIdPath = initial.ObjectIdPath
	}

	return des
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c *Client, des, nw *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ObjectIdPath) {
		nw.ObjectIdPath = des.ObjectIdPath
	}

	return nw
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSet(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(des, initial *CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions, opts ...dcl.ApplyOption) *CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.IsCa, initial.IsCa) || dcl.IsZeroValue(des.IsCa) {
		des.IsCa = initial.IsCa
	}
	if dcl.IsZeroValue(des.MaxIssuerPathLength) {
		des.MaxIssuerPathLength = initial.MaxIssuerPathLength
	}

	return des
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(c *Client, des, nw *CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions) *CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.IsCa, nw.IsCa) {
		nw.IsCa = des.IsCa
	}
	if dcl.IsZeroValue(nw.MaxIssuerPathLength) {
		nw.MaxIssuerPathLength = des.MaxIssuerPathLength
	}

	return nw
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsSet(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions) []CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsSlice(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions) []CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds(des, initial *CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds, opts ...dcl.ApplyOption) *CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObjectIdPath) {
		des.ObjectIdPath = initial.ObjectIdPath
	}

	return des
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds(c *Client, des, nw *CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds) *CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ObjectIdPath) {
		nw.ObjectIdPath = des.ObjectIdPath
	}

	return nw
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsSet(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds) []CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsSlice(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds) []CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions(des, initial *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions, opts ...dcl.ApplyOption) *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.ObjectId = canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(des.ObjectId, initial.ObjectId, opts...)
	if dcl.BoolCanonicalize(des.Critical, initial.Critical) || dcl.IsZeroValue(des.Critical) {
		des.Critical = initial.Critical
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions(c *Client, des, nw *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions) *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions {
	if des == nil || nw == nil {
		return nw
	}

	nw.ObjectId = canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(c, des.ObjectId, nw.ObjectId)
	if dcl.BoolCanonicalize(des.Critical, nw.Critical) {
		nw.Critical = des.Critical
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsSet(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions) []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsSlice(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions) []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(des, initial *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId, opts ...dcl.ApplyOption) *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObjectIdPath) {
		des.ObjectIdPath = initial.ObjectIdPath
	}

	return des
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(c *Client, des, nw *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId) *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ObjectIdPath) {
		nw.ObjectIdPath = des.ObjectIdPath
	}

	return nw
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdSet(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId) []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdSlice(c *Client, des, nw []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId) []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityKeySpec(des, initial *CertificateAuthorityKeySpec, opts ...dcl.ApplyOption) *CertificateAuthorityKeySpec {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.CloudKmsKeyVersion, initial.CloudKmsKeyVersion) || dcl.IsZeroValue(des.CloudKmsKeyVersion) {
		des.CloudKmsKeyVersion = initial.CloudKmsKeyVersion
	}
	if dcl.IsZeroValue(des.Algorithm) {
		des.Algorithm = initial.Algorithm
	}

	return des
}

func canonicalizeNewCertificateAuthorityKeySpec(c *Client, des, nw *CertificateAuthorityKeySpec) *CertificateAuthorityKeySpec {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.CloudKmsKeyVersion, nw.CloudKmsKeyVersion) {
		nw.CloudKmsKeyVersion = des.CloudKmsKeyVersion
	}
	if dcl.IsZeroValue(nw.Algorithm) {
		nw.Algorithm = des.Algorithm
	}

	return nw
}

func canonicalizeNewCertificateAuthorityKeySpecSet(c *Client, des, nw []CertificateAuthorityKeySpec) []CertificateAuthorityKeySpec {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityKeySpec
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityKeySpecNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityKeySpecSlice(c *Client, des, nw []CertificateAuthorityKeySpec) []CertificateAuthorityKeySpec {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityKeySpec
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityKeySpec(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthoritySubordinateConfig(des, initial *CertificateAuthoritySubordinateConfig, opts ...dcl.ApplyOption) *CertificateAuthoritySubordinateConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.NameToSelfLink(des.CertificateAuthority, initial.CertificateAuthority) || dcl.IsZeroValue(des.CertificateAuthority) {
		des.CertificateAuthority = initial.CertificateAuthority
	}
	des.PemIssuerChain = canonicalizeCertificateAuthoritySubordinateConfigPemIssuerChain(des.PemIssuerChain, initial.PemIssuerChain, opts...)

	return des
}

func canonicalizeNewCertificateAuthoritySubordinateConfig(c *Client, des, nw *CertificateAuthoritySubordinateConfig) *CertificateAuthoritySubordinateConfig {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.NameToSelfLink(des.CertificateAuthority, nw.CertificateAuthority) {
		nw.CertificateAuthority = des.CertificateAuthority
	}
	nw.PemIssuerChain = canonicalizeNewCertificateAuthoritySubordinateConfigPemIssuerChain(c, des.PemIssuerChain, nw.PemIssuerChain)

	return nw
}

func canonicalizeNewCertificateAuthoritySubordinateConfigSet(c *Client, des, nw []CertificateAuthoritySubordinateConfig) []CertificateAuthoritySubordinateConfig {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthoritySubordinateConfig
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthoritySubordinateConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthoritySubordinateConfigSlice(c *Client, des, nw []CertificateAuthoritySubordinateConfig) []CertificateAuthoritySubordinateConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthoritySubordinateConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthoritySubordinateConfig(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthoritySubordinateConfigPemIssuerChain(des, initial *CertificateAuthoritySubordinateConfigPemIssuerChain, opts ...dcl.ApplyOption) *CertificateAuthoritySubordinateConfigPemIssuerChain {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.PemCertificates) {
		des.PemCertificates = initial.PemCertificates
	}

	return des
}

func canonicalizeNewCertificateAuthoritySubordinateConfigPemIssuerChain(c *Client, des, nw *CertificateAuthoritySubordinateConfigPemIssuerChain) *CertificateAuthoritySubordinateConfigPemIssuerChain {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.PemCertificates) {
		nw.PemCertificates = des.PemCertificates
	}

	return nw
}

func canonicalizeNewCertificateAuthoritySubordinateConfigPemIssuerChainSet(c *Client, des, nw []CertificateAuthoritySubordinateConfigPemIssuerChain) []CertificateAuthoritySubordinateConfigPemIssuerChain {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthoritySubordinateConfigPemIssuerChain
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthoritySubordinateConfigPemIssuerChainNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthoritySubordinateConfigPemIssuerChainSlice(c *Client, des, nw []CertificateAuthoritySubordinateConfigPemIssuerChain) []CertificateAuthoritySubordinateConfigPemIssuerChain {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthoritySubordinateConfigPemIssuerChain
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthoritySubordinateConfigPemIssuerChain(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptions(des, initial *CertificateAuthorityCaCertificateDescriptions, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.SubjectDescription = canonicalizeCertificateAuthorityCaCertificateDescriptionsSubjectDescription(des.SubjectDescription, initial.SubjectDescription, opts...)
	des.PublicKey = canonicalizeCertificateAuthorityCaCertificateDescriptionsPublicKey(des.PublicKey, initial.PublicKey, opts...)
	des.SubjectKeyId = canonicalizeCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(des.SubjectKeyId, initial.SubjectKeyId, opts...)
	des.AuthorityKeyId = canonicalizeCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(des.AuthorityKeyId, initial.AuthorityKeyId, opts...)
	if dcl.IsZeroValue(des.CrlDistributionPoints) {
		des.CrlDistributionPoints = initial.CrlDistributionPoints
	}
	if dcl.IsZeroValue(des.AiaIssuingCertificateUrls) {
		des.AiaIssuingCertificateUrls = initial.AiaIssuingCertificateUrls
	}
	des.CertFingerprint = canonicalizeCertificateAuthorityCaCertificateDescriptionsCertFingerprint(des.CertFingerprint, initial.CertFingerprint, opts...)
	des.ConfigValues = canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValues(des.ConfigValues, initial.ConfigValues, opts...)

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptions(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptions) *CertificateAuthorityCaCertificateDescriptions {
	if des == nil || nw == nil {
		return nw
	}

	nw.SubjectDescription = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescription(c, des.SubjectDescription, nw.SubjectDescription)
	nw.PublicKey = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsPublicKey(c, des.PublicKey, nw.PublicKey)
	nw.SubjectKeyId = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(c, des.SubjectKeyId, nw.SubjectKeyId)
	nw.AuthorityKeyId = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(c, des.AuthorityKeyId, nw.AuthorityKeyId)
	if dcl.IsZeroValue(nw.CrlDistributionPoints) {
		nw.CrlDistributionPoints = des.CrlDistributionPoints
	}
	if dcl.IsZeroValue(nw.AiaIssuingCertificateUrls) {
		nw.AiaIssuingCertificateUrls = des.AiaIssuingCertificateUrls
	}
	nw.CertFingerprint = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsCertFingerprint(c, des.CertFingerprint, nw.CertFingerprint)
	nw.ConfigValues = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValues(c, des.ConfigValues, nw.ConfigValues)

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptions) []CertificateAuthorityCaCertificateDescriptions {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptions) []CertificateAuthorityCaCertificateDescriptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptions(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsSubjectDescription(des, initial *CertificateAuthorityCaCertificateDescriptionsSubjectDescription, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.Subject = canonicalizeCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(des.Subject, initial.Subject, opts...)
	des.SubjectAltName = canonicalizeCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(des.SubjectAltName, initial.SubjectAltName, opts...)
	if dcl.StringCanonicalize(des.HexSerialNumber, initial.HexSerialNumber) || dcl.IsZeroValue(des.HexSerialNumber) {
		des.HexSerialNumber = initial.HexSerialNumber
	}
	if dcl.StringCanonicalize(des.Lifetime, initial.Lifetime) || dcl.IsZeroValue(des.Lifetime) {
		des.Lifetime = initial.Lifetime
	}
	if dcl.IsZeroValue(des.NotBeforeTime) {
		des.NotBeforeTime = initial.NotBeforeTime
	}
	if dcl.IsZeroValue(des.NotAfterTime) {
		des.NotAfterTime = initial.NotAfterTime
	}
	if dcl.StringCanonicalize(des.CommonName, initial.CommonName) || dcl.IsZeroValue(des.CommonName) {
		des.CommonName = initial.CommonName
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescription(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsSubjectDescription) *CertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	if des == nil || nw == nil {
		return nw
	}

	nw.Subject = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(c, des.Subject, nw.Subject)
	nw.SubjectAltName = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(c, des.SubjectAltName, nw.SubjectAltName)
	if dcl.StringCanonicalize(des.HexSerialNumber, nw.HexSerialNumber) {
		nw.HexSerialNumber = des.HexSerialNumber
	}
	if dcl.StringCanonicalize(des.Lifetime, nw.Lifetime) {
		nw.Lifetime = des.Lifetime
	}
	if dcl.IsZeroValue(nw.NotBeforeTime) {
		nw.NotBeforeTime = des.NotBeforeTime
	}
	if dcl.IsZeroValue(nw.NotAfterTime) {
		nw.NotAfterTime = des.NotAfterTime
	}
	if dcl.StringCanonicalize(des.CommonName, nw.CommonName) {
		nw.CommonName = des.CommonName
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsSubjectDescription) []CertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsSubjectDescription
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsSubjectDescription) []CertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsSubjectDescription
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescription(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(des, initial *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.CommonName, initial.CommonName) || dcl.IsZeroValue(des.CommonName) {
		des.CommonName = initial.CommonName
	}
	if dcl.StringCanonicalize(des.CountryCode, initial.CountryCode) || dcl.IsZeroValue(des.CountryCode) {
		des.CountryCode = initial.CountryCode
	}
	if dcl.StringCanonicalize(des.Organization, initial.Organization) || dcl.IsZeroValue(des.Organization) {
		des.Organization = initial.Organization
	}
	if dcl.StringCanonicalize(des.OrganizationalUnit, initial.OrganizationalUnit) || dcl.IsZeroValue(des.OrganizationalUnit) {
		des.OrganizationalUnit = initial.OrganizationalUnit
	}
	if dcl.StringCanonicalize(des.Locality, initial.Locality) || dcl.IsZeroValue(des.Locality) {
		des.Locality = initial.Locality
	}
	if dcl.StringCanonicalize(des.Province, initial.Province) || dcl.IsZeroValue(des.Province) {
		des.Province = initial.Province
	}
	if dcl.StringCanonicalize(des.StreetAddress, initial.StreetAddress) || dcl.IsZeroValue(des.StreetAddress) {
		des.StreetAddress = initial.StreetAddress
	}
	if dcl.StringCanonicalize(des.PostalCode, initial.PostalCode) || dcl.IsZeroValue(des.PostalCode) {
		des.PostalCode = initial.PostalCode
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.CommonName, nw.CommonName) {
		nw.CommonName = des.CommonName
	}
	if dcl.StringCanonicalize(des.CountryCode, nw.CountryCode) {
		nw.CountryCode = des.CountryCode
	}
	if dcl.StringCanonicalize(des.Organization, nw.Organization) {
		nw.Organization = des.Organization
	}
	if dcl.StringCanonicalize(des.OrganizationalUnit, nw.OrganizationalUnit) {
		nw.OrganizationalUnit = des.OrganizationalUnit
	}
	if dcl.StringCanonicalize(des.Locality, nw.Locality) {
		nw.Locality = des.Locality
	}
	if dcl.StringCanonicalize(des.Province, nw.Province) {
		nw.Province = des.Province
	}
	if dcl.StringCanonicalize(des.StreetAddress, nw.StreetAddress) {
		nw.StreetAddress = des.StreetAddress
	}
	if dcl.StringCanonicalize(des.PostalCode, nw.PostalCode) {
		nw.PostalCode = des.PostalCode
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(des, initial *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.DnsNames) {
		des.DnsNames = initial.DnsNames
	}
	if dcl.IsZeroValue(des.Uris) {
		des.Uris = initial.Uris
	}
	if dcl.IsZeroValue(des.EmailAddresses) {
		des.EmailAddresses = initial.EmailAddresses
	}
	if dcl.IsZeroValue(des.IPAddresses) {
		des.IPAddresses = initial.IPAddresses
	}
	if dcl.IsZeroValue(des.CustomSans) {
		des.CustomSans = initial.CustomSans
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.DnsNames) {
		nw.DnsNames = des.DnsNames
	}
	if dcl.IsZeroValue(nw.Uris) {
		nw.Uris = des.Uris
	}
	if dcl.IsZeroValue(nw.EmailAddresses) {
		nw.EmailAddresses = des.EmailAddresses
	}
	if dcl.IsZeroValue(nw.IPAddresses) {
		nw.IPAddresses = des.IPAddresses
	}
	nw.CustomSans = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansSlice(c, des.CustomSans, nw.CustomSans)

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(des, initial *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.ObjectId = canonicalizeCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(des.ObjectId, initial.ObjectId, opts...)
	if dcl.BoolCanonicalize(des.Critical, initial.Critical) || dcl.IsZeroValue(des.Critical) {
		des.Critical = initial.Critical
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	if des == nil || nw == nil {
		return nw
	}

	nw.ObjectId = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(c, des.ObjectId, nw.ObjectId)
	if dcl.BoolCanonicalize(des.Critical, nw.Critical) {
		nw.Critical = des.Critical
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(des, initial *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObjectIdPath) {
		des.ObjectIdPath = initial.ObjectIdPath
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ObjectIdPath) {
		nw.ObjectIdPath = des.ObjectIdPath
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsPublicKey(des, initial *CertificateAuthorityCaCertificateDescriptionsPublicKey, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsPublicKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Key, initial.Key) || dcl.IsZeroValue(des.Key) {
		des.Key = initial.Key
	}
	if dcl.IsZeroValue(des.Format) {
		des.Format = initial.Format
	}
	if dcl.IsZeroValue(des.Type) {
		des.Type = initial.Type
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsPublicKey(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsPublicKey) *CertificateAuthorityCaCertificateDescriptionsPublicKey {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Key, nw.Key) {
		nw.Key = des.Key
	}
	if dcl.IsZeroValue(nw.Format) {
		nw.Format = des.Format
	}
	if dcl.IsZeroValue(nw.Type) {
		nw.Type = des.Type
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsPublicKeySet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsPublicKey) []CertificateAuthorityCaCertificateDescriptionsPublicKey {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsPublicKey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsPublicKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsPublicKeySlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsPublicKey) []CertificateAuthorityCaCertificateDescriptionsPublicKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsPublicKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsPublicKey(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(des, initial *CertificateAuthorityCaCertificateDescriptionsSubjectKeyId, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.KeyId, initial.KeyId) || dcl.IsZeroValue(des.KeyId) {
		des.KeyId = initial.KeyId
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) *CertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.KeyId, nw.KeyId) {
		nw.KeyId = des.KeyId
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) []CertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsSubjectKeyId
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) []CertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsSubjectKeyId
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(des, initial *CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.KeyId, initial.KeyId) || dcl.IsZeroValue(des.KeyId) {
		des.KeyId = initial.KeyId
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) *CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.KeyId, nw.KeyId) {
		nw.KeyId = des.KeyId
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) []CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) []CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsCertFingerprint(des, initial *CertificateAuthorityCaCertificateDescriptionsCertFingerprint, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.Sha256Hash, initial.Sha256Hash) || dcl.IsZeroValue(des.Sha256Hash) {
		des.Sha256Hash = initial.Sha256Hash
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsCertFingerprint(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsCertFingerprint) *CertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.Sha256Hash, nw.Sha256Hash) {
		nw.Sha256Hash = des.Sha256Hash
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsCertFingerprintSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsCertFingerprint) []CertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsCertFingerprint
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsCertFingerprintNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsCertFingerprintSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsCertFingerprint) []CertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsCertFingerprint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsCertFingerprint(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValues(des, initial *CertificateAuthorityCaCertificateDescriptionsConfigValues, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsConfigValues {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.KeyUsage = canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(des.KeyUsage, initial.KeyUsage, opts...)
	des.CaOptions = canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(des.CaOptions, initial.CaOptions, opts...)
	if dcl.IsZeroValue(des.PolicyIds) {
		des.PolicyIds = initial.PolicyIds
	}
	if dcl.IsZeroValue(des.AiaOcspServers) {
		des.AiaOcspServers = initial.AiaOcspServers
	}
	if dcl.IsZeroValue(des.AdditionalExtensions) {
		des.AdditionalExtensions = initial.AdditionalExtensions
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValues(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsConfigValues) *CertificateAuthorityCaCertificateDescriptionsConfigValues {
	if des == nil || nw == nil {
		return nw
	}

	nw.KeyUsage = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(c, des.KeyUsage, nw.KeyUsage)
	nw.CaOptions = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(c, des.CaOptions, nw.CaOptions)
	nw.PolicyIds = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsSlice(c, des.PolicyIds, nw.PolicyIds)
	if dcl.IsZeroValue(nw.AiaOcspServers) {
		nw.AiaOcspServers = des.AiaOcspServers
	}
	nw.AdditionalExtensions = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsSlice(c, des.AdditionalExtensions, nw.AdditionalExtensions)

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValues) []CertificateAuthorityCaCertificateDescriptionsConfigValues {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsConfigValues
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsConfigValuesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValues) []CertificateAuthorityCaCertificateDescriptionsConfigValues {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsConfigValues
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValues(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(des, initial *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.BaseKeyUsage = canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(des.BaseKeyUsage, initial.BaseKeyUsage, opts...)
	des.ExtendedKeyUsage = canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(des.ExtendedKeyUsage, initial.ExtendedKeyUsage, opts...)
	if dcl.IsZeroValue(des.UnknownExtendedKeyUsages) {
		des.UnknownExtendedKeyUsages = initial.UnknownExtendedKeyUsages
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage) *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage {
	if des == nil || nw == nil {
		return nw
	}

	nw.BaseKeyUsage = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(c, des.BaseKeyUsage, nw.BaseKeyUsage)
	nw.ExtendedKeyUsage = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(c, des.ExtendedKeyUsage, nw.ExtendedKeyUsage)
	nw.UnknownExtendedKeyUsages = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c, des.UnknownExtendedKeyUsages, nw.UnknownExtendedKeyUsages)

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage) []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage) []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(des, initial *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.DigitalSignature, initial.DigitalSignature) || dcl.IsZeroValue(des.DigitalSignature) {
		des.DigitalSignature = initial.DigitalSignature
	}
	if dcl.BoolCanonicalize(des.ContentCommitment, initial.ContentCommitment) || dcl.IsZeroValue(des.ContentCommitment) {
		des.ContentCommitment = initial.ContentCommitment
	}
	if dcl.BoolCanonicalize(des.KeyEncipherment, initial.KeyEncipherment) || dcl.IsZeroValue(des.KeyEncipherment) {
		des.KeyEncipherment = initial.KeyEncipherment
	}
	if dcl.BoolCanonicalize(des.DataEncipherment, initial.DataEncipherment) || dcl.IsZeroValue(des.DataEncipherment) {
		des.DataEncipherment = initial.DataEncipherment
	}
	if dcl.BoolCanonicalize(des.KeyAgreement, initial.KeyAgreement) || dcl.IsZeroValue(des.KeyAgreement) {
		des.KeyAgreement = initial.KeyAgreement
	}
	if dcl.BoolCanonicalize(des.CertSign, initial.CertSign) || dcl.IsZeroValue(des.CertSign) {
		des.CertSign = initial.CertSign
	}
	if dcl.BoolCanonicalize(des.CrlSign, initial.CrlSign) || dcl.IsZeroValue(des.CrlSign) {
		des.CrlSign = initial.CrlSign
	}
	if dcl.BoolCanonicalize(des.EncipherOnly, initial.EncipherOnly) || dcl.IsZeroValue(des.EncipherOnly) {
		des.EncipherOnly = initial.EncipherOnly
	}
	if dcl.BoolCanonicalize(des.DecipherOnly, initial.DecipherOnly) || dcl.IsZeroValue(des.DecipherOnly) {
		des.DecipherOnly = initial.DecipherOnly
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage) *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.DigitalSignature, nw.DigitalSignature) {
		nw.DigitalSignature = des.DigitalSignature
	}
	if dcl.BoolCanonicalize(des.ContentCommitment, nw.ContentCommitment) {
		nw.ContentCommitment = des.ContentCommitment
	}
	if dcl.BoolCanonicalize(des.KeyEncipherment, nw.KeyEncipherment) {
		nw.KeyEncipherment = des.KeyEncipherment
	}
	if dcl.BoolCanonicalize(des.DataEncipherment, nw.DataEncipherment) {
		nw.DataEncipherment = des.DataEncipherment
	}
	if dcl.BoolCanonicalize(des.KeyAgreement, nw.KeyAgreement) {
		nw.KeyAgreement = des.KeyAgreement
	}
	if dcl.BoolCanonicalize(des.CertSign, nw.CertSign) {
		nw.CertSign = des.CertSign
	}
	if dcl.BoolCanonicalize(des.CrlSign, nw.CrlSign) {
		nw.CrlSign = des.CrlSign
	}
	if dcl.BoolCanonicalize(des.EncipherOnly, nw.EncipherOnly) {
		nw.EncipherOnly = des.EncipherOnly
	}
	if dcl.BoolCanonicalize(des.DecipherOnly, nw.DecipherOnly) {
		nw.DecipherOnly = des.DecipherOnly
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage) []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage) []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(des, initial *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.ServerAuth, initial.ServerAuth) || dcl.IsZeroValue(des.ServerAuth) {
		des.ServerAuth = initial.ServerAuth
	}
	if dcl.BoolCanonicalize(des.ClientAuth, initial.ClientAuth) || dcl.IsZeroValue(des.ClientAuth) {
		des.ClientAuth = initial.ClientAuth
	}
	if dcl.BoolCanonicalize(des.CodeSigning, initial.CodeSigning) || dcl.IsZeroValue(des.CodeSigning) {
		des.CodeSigning = initial.CodeSigning
	}
	if dcl.BoolCanonicalize(des.EmailProtection, initial.EmailProtection) || dcl.IsZeroValue(des.EmailProtection) {
		des.EmailProtection = initial.EmailProtection
	}
	if dcl.BoolCanonicalize(des.TimeStamping, initial.TimeStamping) || dcl.IsZeroValue(des.TimeStamping) {
		des.TimeStamping = initial.TimeStamping
	}
	if dcl.BoolCanonicalize(des.OcspSigning, initial.OcspSigning) || dcl.IsZeroValue(des.OcspSigning) {
		des.OcspSigning = initial.OcspSigning
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage) *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.ServerAuth, nw.ServerAuth) {
		nw.ServerAuth = des.ServerAuth
	}
	if dcl.BoolCanonicalize(des.ClientAuth, nw.ClientAuth) {
		nw.ClientAuth = des.ClientAuth
	}
	if dcl.BoolCanonicalize(des.CodeSigning, nw.CodeSigning) {
		nw.CodeSigning = des.CodeSigning
	}
	if dcl.BoolCanonicalize(des.EmailProtection, nw.EmailProtection) {
		nw.EmailProtection = des.EmailProtection
	}
	if dcl.BoolCanonicalize(des.TimeStamping, nw.TimeStamping) {
		nw.TimeStamping = des.TimeStamping
	}
	if dcl.BoolCanonicalize(des.OcspSigning, nw.OcspSigning) {
		nw.OcspSigning = des.OcspSigning
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage) []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage) []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages(des, initial *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObjectIdPath) {
		des.ObjectIdPath = initial.ObjectIdPath
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages) *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ObjectIdPath) {
		nw.ObjectIdPath = des.ObjectIdPath
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages) []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages) []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(des, initial *CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.IsCa, initial.IsCa) || dcl.IsZeroValue(des.IsCa) {
		des.IsCa = initial.IsCa
	}
	if dcl.IsZeroValue(des.MaxIssuerPathLength) {
		des.MaxIssuerPathLength = initial.MaxIssuerPathLength
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions) *CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.IsCa, nw.IsCa) {
		nw.IsCa = des.IsCa
	}
	if dcl.IsZeroValue(nw.MaxIssuerPathLength) {
		nw.MaxIssuerPathLength = des.MaxIssuerPathLength
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions) []CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions) []CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds(des, initial *CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObjectIdPath) {
		des.ObjectIdPath = initial.ObjectIdPath
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds) *CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ObjectIdPath) {
		nw.ObjectIdPath = des.ObjectIdPath
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds) []CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds) []CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions(des, initial *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.ObjectId = canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(des.ObjectId, initial.ObjectId, opts...)
	if dcl.BoolCanonicalize(des.Critical, initial.Critical) || dcl.IsZeroValue(des.Critical) {
		des.Critical = initial.Critical
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions) *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions {
	if des == nil || nw == nil {
		return nw
	}

	nw.ObjectId = canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(c, des.ObjectId, nw.ObjectId)
	if dcl.BoolCanonicalize(des.Critical, nw.Critical) {
		nw.Critical = des.Critical
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions) []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions) []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(des, initial *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId, opts ...dcl.ApplyOption) *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObjectIdPath) {
		des.ObjectIdPath = initial.ObjectIdPath
	}

	return des
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(c *Client, des, nw *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId) *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ObjectIdPath) {
		nw.ObjectIdPath = des.ObjectIdPath
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdSet(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId) []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdSlice(c *Client, des, nw []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId) []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityAccessUrls(des, initial *CertificateAuthorityAccessUrls, opts ...dcl.ApplyOption) *CertificateAuthorityAccessUrls {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.CaCertificateAccessUrl, initial.CaCertificateAccessUrl) || dcl.IsZeroValue(des.CaCertificateAccessUrl) {
		des.CaCertificateAccessUrl = initial.CaCertificateAccessUrl
	}
	if dcl.IsZeroValue(des.CrlAccessUrls) {
		des.CrlAccessUrls = initial.CrlAccessUrls
	}
	if dcl.StringCanonicalize(des.CrlAccessUrl, initial.CrlAccessUrl) || dcl.IsZeroValue(des.CrlAccessUrl) {
		des.CrlAccessUrl = initial.CrlAccessUrl
	}

	return des
}

func canonicalizeNewCertificateAuthorityAccessUrls(c *Client, des, nw *CertificateAuthorityAccessUrls) *CertificateAuthorityAccessUrls {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.CaCertificateAccessUrl, nw.CaCertificateAccessUrl) {
		nw.CaCertificateAccessUrl = des.CaCertificateAccessUrl
	}
	if dcl.IsZeroValue(nw.CrlAccessUrls) {
		nw.CrlAccessUrls = des.CrlAccessUrls
	}
	if dcl.StringCanonicalize(des.CrlAccessUrl, nw.CrlAccessUrl) {
		nw.CrlAccessUrl = des.CrlAccessUrl
	}

	return nw
}

func canonicalizeNewCertificateAuthorityAccessUrlsSet(c *Client, des, nw []CertificateAuthorityAccessUrls) []CertificateAuthorityAccessUrls {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityAccessUrls
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityAccessUrlsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityAccessUrlsSlice(c *Client, des, nw []CertificateAuthorityAccessUrls) []CertificateAuthorityAccessUrls {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityAccessUrls
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityAccessUrls(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicy(des, initial *CertificateAuthorityCertificatePolicy, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicy {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.AllowedConfigList = canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigList(des.AllowedConfigList, initial.AllowedConfigList, opts...)
	des.OverwriteConfigValues = canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValues(des.OverwriteConfigValues, initial.OverwriteConfigValues, opts...)
	if dcl.IsZeroValue(des.AllowedLocationsAndOrganizations) {
		des.AllowedLocationsAndOrganizations = initial.AllowedLocationsAndOrganizations
	}
	if dcl.IsZeroValue(des.AllowedCommonNames) {
		des.AllowedCommonNames = initial.AllowedCommonNames
	}
	des.AllowedSans = canonicalizeCertificateAuthorityCertificatePolicyAllowedSans(des.AllowedSans, initial.AllowedSans, opts...)
	if dcl.StringCanonicalize(des.MaximumLifetime, initial.MaximumLifetime) || dcl.IsZeroValue(des.MaximumLifetime) {
		des.MaximumLifetime = initial.MaximumLifetime
	}
	des.AllowedIssuanceModes = canonicalizeCertificateAuthorityCertificatePolicyAllowedIssuanceModes(des.AllowedIssuanceModes, initial.AllowedIssuanceModes, opts...)

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicy(c *Client, des, nw *CertificateAuthorityCertificatePolicy) *CertificateAuthorityCertificatePolicy {
	if des == nil || nw == nil {
		return nw
	}

	nw.AllowedConfigList = canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigList(c, des.AllowedConfigList, nw.AllowedConfigList)
	nw.OverwriteConfigValues = canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValues(c, des.OverwriteConfigValues, nw.OverwriteConfigValues)
	nw.AllowedLocationsAndOrganizations = canonicalizeNewCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsSlice(c, des.AllowedLocationsAndOrganizations, nw.AllowedLocationsAndOrganizations)
	if dcl.IsZeroValue(nw.AllowedCommonNames) {
		nw.AllowedCommonNames = des.AllowedCommonNames
	}
	nw.AllowedSans = canonicalizeNewCertificateAuthorityCertificatePolicyAllowedSans(c, des.AllowedSans, nw.AllowedSans)
	if dcl.StringCanonicalize(des.MaximumLifetime, nw.MaximumLifetime) {
		nw.MaximumLifetime = des.MaximumLifetime
	}
	nw.AllowedIssuanceModes = canonicalizeNewCertificateAuthorityCertificatePolicyAllowedIssuanceModes(c, des.AllowedIssuanceModes, nw.AllowedIssuanceModes)

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicySet(c *Client, des, nw []CertificateAuthorityCertificatePolicy) []CertificateAuthorityCertificatePolicy {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicy
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicySlice(c *Client, des, nw []CertificateAuthorityCertificatePolicy) []CertificateAuthorityCertificatePolicy {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicy
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicy(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigList(des, initial *CertificateAuthorityCertificatePolicyAllowedConfigList, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyAllowedConfigList {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AllowedConfigValues) {
		des.AllowedConfigValues = initial.AllowedConfigValues
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigList(c *Client, des, nw *CertificateAuthorityCertificatePolicyAllowedConfigList) *CertificateAuthorityCertificatePolicyAllowedConfigList {
	if des == nil || nw == nil {
		return nw
	}

	nw.AllowedConfigValues = canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesSlice(c, des.AllowedConfigValues, nw.AllowedConfigValues)

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigList) []CertificateAuthorityCertificatePolicyAllowedConfigList {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyAllowedConfigList
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyAllowedConfigListNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigList) []CertificateAuthorityCertificatePolicyAllowedConfigList {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyAllowedConfigList
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigList(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues(des, initial *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ReusableConfig, initial.ReusableConfig) || dcl.IsZeroValue(des.ReusableConfig) {
		des.ReusableConfig = initial.ReusableConfig
	}
	des.ReusableConfigValues = canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(des.ReusableConfigValues, initial.ReusableConfigValues, opts...)

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues(c *Client, des, nw *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ReusableConfig, nw.ReusableConfig) {
		nw.ReusableConfig = des.ReusableConfig
	}
	nw.ReusableConfigValues = canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(c, des.ReusableConfigValues, nw.ReusableConfigValues)

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(des, initial *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.KeyUsage = canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(des.KeyUsage, initial.KeyUsage, opts...)
	des.CaOptions = canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(des.CaOptions, initial.CaOptions, opts...)
	if dcl.IsZeroValue(des.PolicyIds) {
		des.PolicyIds = initial.PolicyIds
	}
	if dcl.IsZeroValue(des.AiaOcspServers) {
		des.AiaOcspServers = initial.AiaOcspServers
	}
	if dcl.IsZeroValue(des.AdditionalExtensions) {
		des.AdditionalExtensions = initial.AdditionalExtensions
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(c *Client, des, nw *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues {
	if des == nil || nw == nil {
		return nw
	}

	nw.KeyUsage = canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(c, des.KeyUsage, nw.KeyUsage)
	nw.CaOptions = canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(c, des.CaOptions, nw.CaOptions)
	nw.PolicyIds = canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsSlice(c, des.PolicyIds, nw.PolicyIds)
	if dcl.IsZeroValue(nw.AiaOcspServers) {
		nw.AiaOcspServers = des.AiaOcspServers
	}
	nw.AdditionalExtensions = canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsSlice(c, des.AdditionalExtensions, nw.AdditionalExtensions)

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(des, initial *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.BaseKeyUsage = canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(des.BaseKeyUsage, initial.BaseKeyUsage, opts...)
	des.ExtendedKeyUsage = canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(des.ExtendedKeyUsage, initial.ExtendedKeyUsage, opts...)
	if dcl.IsZeroValue(des.UnknownExtendedKeyUsages) {
		des.UnknownExtendedKeyUsages = initial.UnknownExtendedKeyUsages
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(c *Client, des, nw *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage {
	if des == nil || nw == nil {
		return nw
	}

	nw.BaseKeyUsage = canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, des.BaseKeyUsage, nw.BaseKeyUsage)
	nw.ExtendedKeyUsage = canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, des.ExtendedKeyUsage, nw.ExtendedKeyUsage)
	nw.UnknownExtendedKeyUsages = canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c, des.UnknownExtendedKeyUsages, nw.UnknownExtendedKeyUsages)

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(des, initial *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.DigitalSignature, initial.DigitalSignature) || dcl.IsZeroValue(des.DigitalSignature) {
		des.DigitalSignature = initial.DigitalSignature
	}
	if dcl.BoolCanonicalize(des.ContentCommitment, initial.ContentCommitment) || dcl.IsZeroValue(des.ContentCommitment) {
		des.ContentCommitment = initial.ContentCommitment
	}
	if dcl.BoolCanonicalize(des.KeyEncipherment, initial.KeyEncipherment) || dcl.IsZeroValue(des.KeyEncipherment) {
		des.KeyEncipherment = initial.KeyEncipherment
	}
	if dcl.BoolCanonicalize(des.DataEncipherment, initial.DataEncipherment) || dcl.IsZeroValue(des.DataEncipherment) {
		des.DataEncipherment = initial.DataEncipherment
	}
	if dcl.BoolCanonicalize(des.KeyAgreement, initial.KeyAgreement) || dcl.IsZeroValue(des.KeyAgreement) {
		des.KeyAgreement = initial.KeyAgreement
	}
	if dcl.BoolCanonicalize(des.CertSign, initial.CertSign) || dcl.IsZeroValue(des.CertSign) {
		des.CertSign = initial.CertSign
	}
	if dcl.BoolCanonicalize(des.CrlSign, initial.CrlSign) || dcl.IsZeroValue(des.CrlSign) {
		des.CrlSign = initial.CrlSign
	}
	if dcl.BoolCanonicalize(des.EncipherOnly, initial.EncipherOnly) || dcl.IsZeroValue(des.EncipherOnly) {
		des.EncipherOnly = initial.EncipherOnly
	}
	if dcl.BoolCanonicalize(des.DecipherOnly, initial.DecipherOnly) || dcl.IsZeroValue(des.DecipherOnly) {
		des.DecipherOnly = initial.DecipherOnly
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c *Client, des, nw *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.DigitalSignature, nw.DigitalSignature) {
		nw.DigitalSignature = des.DigitalSignature
	}
	if dcl.BoolCanonicalize(des.ContentCommitment, nw.ContentCommitment) {
		nw.ContentCommitment = des.ContentCommitment
	}
	if dcl.BoolCanonicalize(des.KeyEncipherment, nw.KeyEncipherment) {
		nw.KeyEncipherment = des.KeyEncipherment
	}
	if dcl.BoolCanonicalize(des.DataEncipherment, nw.DataEncipherment) {
		nw.DataEncipherment = des.DataEncipherment
	}
	if dcl.BoolCanonicalize(des.KeyAgreement, nw.KeyAgreement) {
		nw.KeyAgreement = des.KeyAgreement
	}
	if dcl.BoolCanonicalize(des.CertSign, nw.CertSign) {
		nw.CertSign = des.CertSign
	}
	if dcl.BoolCanonicalize(des.CrlSign, nw.CrlSign) {
		nw.CrlSign = des.CrlSign
	}
	if dcl.BoolCanonicalize(des.EncipherOnly, nw.EncipherOnly) {
		nw.EncipherOnly = des.EncipherOnly
	}
	if dcl.BoolCanonicalize(des.DecipherOnly, nw.DecipherOnly) {
		nw.DecipherOnly = des.DecipherOnly
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(des, initial *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.ServerAuth, initial.ServerAuth) || dcl.IsZeroValue(des.ServerAuth) {
		des.ServerAuth = initial.ServerAuth
	}
	if dcl.BoolCanonicalize(des.ClientAuth, initial.ClientAuth) || dcl.IsZeroValue(des.ClientAuth) {
		des.ClientAuth = initial.ClientAuth
	}
	if dcl.BoolCanonicalize(des.CodeSigning, initial.CodeSigning) || dcl.IsZeroValue(des.CodeSigning) {
		des.CodeSigning = initial.CodeSigning
	}
	if dcl.BoolCanonicalize(des.EmailProtection, initial.EmailProtection) || dcl.IsZeroValue(des.EmailProtection) {
		des.EmailProtection = initial.EmailProtection
	}
	if dcl.BoolCanonicalize(des.TimeStamping, initial.TimeStamping) || dcl.IsZeroValue(des.TimeStamping) {
		des.TimeStamping = initial.TimeStamping
	}
	if dcl.BoolCanonicalize(des.OcspSigning, initial.OcspSigning) || dcl.IsZeroValue(des.OcspSigning) {
		des.OcspSigning = initial.OcspSigning
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c *Client, des, nw *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.ServerAuth, nw.ServerAuth) {
		nw.ServerAuth = des.ServerAuth
	}
	if dcl.BoolCanonicalize(des.ClientAuth, nw.ClientAuth) {
		nw.ClientAuth = des.ClientAuth
	}
	if dcl.BoolCanonicalize(des.CodeSigning, nw.CodeSigning) {
		nw.CodeSigning = des.CodeSigning
	}
	if dcl.BoolCanonicalize(des.EmailProtection, nw.EmailProtection) {
		nw.EmailProtection = des.EmailProtection
	}
	if dcl.BoolCanonicalize(des.TimeStamping, nw.TimeStamping) {
		nw.TimeStamping = des.TimeStamping
	}
	if dcl.BoolCanonicalize(des.OcspSigning, nw.OcspSigning) {
		nw.OcspSigning = des.OcspSigning
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(des, initial *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObjectIdPath) {
		des.ObjectIdPath = initial.ObjectIdPath
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c *Client, des, nw *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ObjectIdPath) {
		nw.ObjectIdPath = des.ObjectIdPath
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(des, initial *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.IsCa, initial.IsCa) || dcl.IsZeroValue(des.IsCa) {
		des.IsCa = initial.IsCa
	}
	if dcl.IsZeroValue(des.MaxIssuerPathLength) {
		des.MaxIssuerPathLength = initial.MaxIssuerPathLength
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(c *Client, des, nw *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.IsCa, nw.IsCa) {
		nw.IsCa = des.IsCa
	}
	if dcl.IsZeroValue(nw.MaxIssuerPathLength) {
		nw.MaxIssuerPathLength = des.MaxIssuerPathLength
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds(des, initial *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObjectIdPath) {
		des.ObjectIdPath = initial.ObjectIdPath
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds(c *Client, des, nw *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ObjectIdPath) {
		nw.ObjectIdPath = des.ObjectIdPath
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions(des, initial *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.ObjectId = canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(des.ObjectId, initial.ObjectId, opts...)
	if dcl.BoolCanonicalize(des.Critical, initial.Critical) || dcl.IsZeroValue(des.Critical) {
		des.Critical = initial.Critical
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions(c *Client, des, nw *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions {
	if des == nil || nw == nil {
		return nw
	}

	nw.ObjectId = canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, des.ObjectId, nw.ObjectId)
	if dcl.BoolCanonicalize(des.Critical, nw.Critical) {
		nw.Critical = des.Critical
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(des, initial *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObjectIdPath) {
		des.ObjectIdPath = initial.ObjectIdPath
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c *Client, des, nw *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ObjectIdPath) {
		nw.ObjectIdPath = des.ObjectIdPath
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValues(des, initial *CertificateAuthorityCertificatePolicyOverwriteConfigValues, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyOverwriteConfigValues {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.ReusableConfig, initial.ReusableConfig) || dcl.IsZeroValue(des.ReusableConfig) {
		des.ReusableConfig = initial.ReusableConfig
	}
	des.ReusableConfigValues = canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(des.ReusableConfigValues, initial.ReusableConfigValues, opts...)

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValues(c *Client, des, nw *CertificateAuthorityCertificatePolicyOverwriteConfigValues) *CertificateAuthorityCertificatePolicyOverwriteConfigValues {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.ReusableConfig, nw.ReusableConfig) {
		nw.ReusableConfig = des.ReusableConfig
	}
	nw.ReusableConfigValues = canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(c, des.ReusableConfigValues, nw.ReusableConfigValues)

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValues) []CertificateAuthorityCertificatePolicyOverwriteConfigValues {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyOverwriteConfigValues
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValues) []CertificateAuthorityCertificatePolicyOverwriteConfigValues {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyOverwriteConfigValues
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValues(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(des, initial *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.KeyUsage = canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(des.KeyUsage, initial.KeyUsage, opts...)
	des.CaOptions = canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(des.CaOptions, initial.CaOptions, opts...)
	if dcl.IsZeroValue(des.PolicyIds) {
		des.PolicyIds = initial.PolicyIds
	}
	if dcl.IsZeroValue(des.AiaOcspServers) {
		des.AiaOcspServers = initial.AiaOcspServers
	}
	if dcl.IsZeroValue(des.AdditionalExtensions) {
		des.AdditionalExtensions = initial.AdditionalExtensions
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(c *Client, des, nw *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues {
	if des == nil || nw == nil {
		return nw
	}

	nw.KeyUsage = canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(c, des.KeyUsage, nw.KeyUsage)
	nw.CaOptions = canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(c, des.CaOptions, nw.CaOptions)
	nw.PolicyIds = canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsSlice(c, des.PolicyIds, nw.PolicyIds)
	if dcl.IsZeroValue(nw.AiaOcspServers) {
		nw.AiaOcspServers = des.AiaOcspServers
	}
	nw.AdditionalExtensions = canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsSlice(c, des.AdditionalExtensions, nw.AdditionalExtensions)

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(des, initial *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.BaseKeyUsage = canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(des.BaseKeyUsage, initial.BaseKeyUsage, opts...)
	des.ExtendedKeyUsage = canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(des.ExtendedKeyUsage, initial.ExtendedKeyUsage, opts...)
	if dcl.IsZeroValue(des.UnknownExtendedKeyUsages) {
		des.UnknownExtendedKeyUsages = initial.UnknownExtendedKeyUsages
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(c *Client, des, nw *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage {
	if des == nil || nw == nil {
		return nw
	}

	nw.BaseKeyUsage = canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, des.BaseKeyUsage, nw.BaseKeyUsage)
	nw.ExtendedKeyUsage = canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, des.ExtendedKeyUsage, nw.ExtendedKeyUsage)
	nw.UnknownExtendedKeyUsages = canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c, des.UnknownExtendedKeyUsages, nw.UnknownExtendedKeyUsages)

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(des, initial *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.DigitalSignature, initial.DigitalSignature) || dcl.IsZeroValue(des.DigitalSignature) {
		des.DigitalSignature = initial.DigitalSignature
	}
	if dcl.BoolCanonicalize(des.ContentCommitment, initial.ContentCommitment) || dcl.IsZeroValue(des.ContentCommitment) {
		des.ContentCommitment = initial.ContentCommitment
	}
	if dcl.BoolCanonicalize(des.KeyEncipherment, initial.KeyEncipherment) || dcl.IsZeroValue(des.KeyEncipherment) {
		des.KeyEncipherment = initial.KeyEncipherment
	}
	if dcl.BoolCanonicalize(des.DataEncipherment, initial.DataEncipherment) || dcl.IsZeroValue(des.DataEncipherment) {
		des.DataEncipherment = initial.DataEncipherment
	}
	if dcl.BoolCanonicalize(des.KeyAgreement, initial.KeyAgreement) || dcl.IsZeroValue(des.KeyAgreement) {
		des.KeyAgreement = initial.KeyAgreement
	}
	if dcl.BoolCanonicalize(des.CertSign, initial.CertSign) || dcl.IsZeroValue(des.CertSign) {
		des.CertSign = initial.CertSign
	}
	if dcl.BoolCanonicalize(des.CrlSign, initial.CrlSign) || dcl.IsZeroValue(des.CrlSign) {
		des.CrlSign = initial.CrlSign
	}
	if dcl.BoolCanonicalize(des.EncipherOnly, initial.EncipherOnly) || dcl.IsZeroValue(des.EncipherOnly) {
		des.EncipherOnly = initial.EncipherOnly
	}
	if dcl.BoolCanonicalize(des.DecipherOnly, initial.DecipherOnly) || dcl.IsZeroValue(des.DecipherOnly) {
		des.DecipherOnly = initial.DecipherOnly
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c *Client, des, nw *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.DigitalSignature, nw.DigitalSignature) {
		nw.DigitalSignature = des.DigitalSignature
	}
	if dcl.BoolCanonicalize(des.ContentCommitment, nw.ContentCommitment) {
		nw.ContentCommitment = des.ContentCommitment
	}
	if dcl.BoolCanonicalize(des.KeyEncipherment, nw.KeyEncipherment) {
		nw.KeyEncipherment = des.KeyEncipherment
	}
	if dcl.BoolCanonicalize(des.DataEncipherment, nw.DataEncipherment) {
		nw.DataEncipherment = des.DataEncipherment
	}
	if dcl.BoolCanonicalize(des.KeyAgreement, nw.KeyAgreement) {
		nw.KeyAgreement = des.KeyAgreement
	}
	if dcl.BoolCanonicalize(des.CertSign, nw.CertSign) {
		nw.CertSign = des.CertSign
	}
	if dcl.BoolCanonicalize(des.CrlSign, nw.CrlSign) {
		nw.CrlSign = des.CrlSign
	}
	if dcl.BoolCanonicalize(des.EncipherOnly, nw.EncipherOnly) {
		nw.EncipherOnly = des.EncipherOnly
	}
	if dcl.BoolCanonicalize(des.DecipherOnly, nw.DecipherOnly) {
		nw.DecipherOnly = des.DecipherOnly
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(des, initial *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.ServerAuth, initial.ServerAuth) || dcl.IsZeroValue(des.ServerAuth) {
		des.ServerAuth = initial.ServerAuth
	}
	if dcl.BoolCanonicalize(des.ClientAuth, initial.ClientAuth) || dcl.IsZeroValue(des.ClientAuth) {
		des.ClientAuth = initial.ClientAuth
	}
	if dcl.BoolCanonicalize(des.CodeSigning, initial.CodeSigning) || dcl.IsZeroValue(des.CodeSigning) {
		des.CodeSigning = initial.CodeSigning
	}
	if dcl.BoolCanonicalize(des.EmailProtection, initial.EmailProtection) || dcl.IsZeroValue(des.EmailProtection) {
		des.EmailProtection = initial.EmailProtection
	}
	if dcl.BoolCanonicalize(des.TimeStamping, initial.TimeStamping) || dcl.IsZeroValue(des.TimeStamping) {
		des.TimeStamping = initial.TimeStamping
	}
	if dcl.BoolCanonicalize(des.OcspSigning, initial.OcspSigning) || dcl.IsZeroValue(des.OcspSigning) {
		des.OcspSigning = initial.OcspSigning
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c *Client, des, nw *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.ServerAuth, nw.ServerAuth) {
		nw.ServerAuth = des.ServerAuth
	}
	if dcl.BoolCanonicalize(des.ClientAuth, nw.ClientAuth) {
		nw.ClientAuth = des.ClientAuth
	}
	if dcl.BoolCanonicalize(des.CodeSigning, nw.CodeSigning) {
		nw.CodeSigning = des.CodeSigning
	}
	if dcl.BoolCanonicalize(des.EmailProtection, nw.EmailProtection) {
		nw.EmailProtection = des.EmailProtection
	}
	if dcl.BoolCanonicalize(des.TimeStamping, nw.TimeStamping) {
		nw.TimeStamping = des.TimeStamping
	}
	if dcl.BoolCanonicalize(des.OcspSigning, nw.OcspSigning) {
		nw.OcspSigning = des.OcspSigning
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(des, initial *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObjectIdPath) {
		des.ObjectIdPath = initial.ObjectIdPath
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c *Client, des, nw *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ObjectIdPath) {
		nw.ObjectIdPath = des.ObjectIdPath
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(des, initial *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.IsCa, initial.IsCa) || dcl.IsZeroValue(des.IsCa) {
		des.IsCa = initial.IsCa
	}
	if dcl.IsZeroValue(des.MaxIssuerPathLength) {
		des.MaxIssuerPathLength = initial.MaxIssuerPathLength
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(c *Client, des, nw *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.IsCa, nw.IsCa) {
		nw.IsCa = des.IsCa
	}
	if dcl.IsZeroValue(nw.MaxIssuerPathLength) {
		nw.MaxIssuerPathLength = des.MaxIssuerPathLength
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds(des, initial *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObjectIdPath) {
		des.ObjectIdPath = initial.ObjectIdPath
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds(c *Client, des, nw *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ObjectIdPath) {
		nw.ObjectIdPath = des.ObjectIdPath
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions(des, initial *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	des.ObjectId = canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(des.ObjectId, initial.ObjectId, opts...)
	if dcl.BoolCanonicalize(des.Critical, initial.Critical) || dcl.IsZeroValue(des.Critical) {
		des.Critical = initial.Critical
	}
	if dcl.StringCanonicalize(des.Value, initial.Value) || dcl.IsZeroValue(des.Value) {
		des.Value = initial.Value
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions(c *Client, des, nw *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions {
	if des == nil || nw == nil {
		return nw
	}

	nw.ObjectId = canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, des.ObjectId, nw.ObjectId)
	if dcl.BoolCanonicalize(des.Critical, nw.Critical) {
		nw.Critical = des.Critical
	}
	if dcl.StringCanonicalize(des.Value, nw.Value) {
		nw.Value = des.Value
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(des, initial *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.ObjectIdPath) {
		des.ObjectIdPath = initial.ObjectIdPath
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c *Client, des, nw *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.ObjectIdPath) {
		nw.ObjectIdPath = des.ObjectIdPath
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations(des, initial *CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.StringCanonicalize(des.CountryCode, initial.CountryCode) || dcl.IsZeroValue(des.CountryCode) {
		des.CountryCode = initial.CountryCode
	}
	if dcl.StringCanonicalize(des.Organization, initial.Organization) || dcl.IsZeroValue(des.Organization) {
		des.Organization = initial.Organization
	}
	if dcl.StringCanonicalize(des.OrganizationalUnit, initial.OrganizationalUnit) || dcl.IsZeroValue(des.OrganizationalUnit) {
		des.OrganizationalUnit = initial.OrganizationalUnit
	}
	if dcl.StringCanonicalize(des.Locality, initial.Locality) || dcl.IsZeroValue(des.Locality) {
		des.Locality = initial.Locality
	}
	if dcl.StringCanonicalize(des.Province, initial.Province) || dcl.IsZeroValue(des.Province) {
		des.Province = initial.Province
	}
	if dcl.StringCanonicalize(des.StreetAddress, initial.StreetAddress) || dcl.IsZeroValue(des.StreetAddress) {
		des.StreetAddress = initial.StreetAddress
	}
	if dcl.StringCanonicalize(des.PostalCode, initial.PostalCode) || dcl.IsZeroValue(des.PostalCode) {
		des.PostalCode = initial.PostalCode
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations(c *Client, des, nw *CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations) *CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.StringCanonicalize(des.CountryCode, nw.CountryCode) {
		nw.CountryCode = des.CountryCode
	}
	if dcl.StringCanonicalize(des.Organization, nw.Organization) {
		nw.Organization = des.Organization
	}
	if dcl.StringCanonicalize(des.OrganizationalUnit, nw.OrganizationalUnit) {
		nw.OrganizationalUnit = des.OrganizationalUnit
	}
	if dcl.StringCanonicalize(des.Locality, nw.Locality) {
		nw.Locality = des.Locality
	}
	if dcl.StringCanonicalize(des.Province, nw.Province) {
		nw.Province = des.Province
	}
	if dcl.StringCanonicalize(des.StreetAddress, nw.StreetAddress) {
		nw.StreetAddress = des.StreetAddress
	}
	if dcl.StringCanonicalize(des.PostalCode, nw.PostalCode) {
		nw.PostalCode = des.PostalCode
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations) []CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations) []CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyAllowedSans(des, initial *CertificateAuthorityCertificatePolicyAllowedSans, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyAllowedSans {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.IsZeroValue(des.AllowedDnsNames) {
		des.AllowedDnsNames = initial.AllowedDnsNames
	}
	if dcl.IsZeroValue(des.AllowedUris) {
		des.AllowedUris = initial.AllowedUris
	}
	if dcl.IsZeroValue(des.AllowedEmailAddresses) {
		des.AllowedEmailAddresses = initial.AllowedEmailAddresses
	}
	if dcl.IsZeroValue(des.AllowedIps) {
		des.AllowedIps = initial.AllowedIps
	}
	if dcl.BoolCanonicalize(des.AllowGlobbingDnsWildcards, initial.AllowGlobbingDnsWildcards) || dcl.IsZeroValue(des.AllowGlobbingDnsWildcards) {
		des.AllowGlobbingDnsWildcards = initial.AllowGlobbingDnsWildcards
	}
	if dcl.BoolCanonicalize(des.AllowCustomSans, initial.AllowCustomSans) || dcl.IsZeroValue(des.AllowCustomSans) {
		des.AllowCustomSans = initial.AllowCustomSans
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedSans(c *Client, des, nw *CertificateAuthorityCertificatePolicyAllowedSans) *CertificateAuthorityCertificatePolicyAllowedSans {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.IsZeroValue(nw.AllowedDnsNames) {
		nw.AllowedDnsNames = des.AllowedDnsNames
	}
	if dcl.IsZeroValue(nw.AllowedUris) {
		nw.AllowedUris = des.AllowedUris
	}
	if dcl.IsZeroValue(nw.AllowedEmailAddresses) {
		nw.AllowedEmailAddresses = des.AllowedEmailAddresses
	}
	if dcl.IsZeroValue(nw.AllowedIps) {
		nw.AllowedIps = des.AllowedIps
	}
	if dcl.BoolCanonicalize(des.AllowGlobbingDnsWildcards, nw.AllowGlobbingDnsWildcards) {
		nw.AllowGlobbingDnsWildcards = des.AllowGlobbingDnsWildcards
	}
	if dcl.BoolCanonicalize(des.AllowCustomSans, nw.AllowCustomSans) {
		nw.AllowCustomSans = des.AllowCustomSans
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedSansSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedSans) []CertificateAuthorityCertificatePolicyAllowedSans {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyAllowedSans
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyAllowedSansNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedSansSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedSans) []CertificateAuthorityCertificatePolicyAllowedSans {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyAllowedSans
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyAllowedSans(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityCertificatePolicyAllowedIssuanceModes(des, initial *CertificateAuthorityCertificatePolicyAllowedIssuanceModes, opts ...dcl.ApplyOption) *CertificateAuthorityCertificatePolicyAllowedIssuanceModes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.AllowCsrBasedIssuance, initial.AllowCsrBasedIssuance) || dcl.IsZeroValue(des.AllowCsrBasedIssuance) {
		des.AllowCsrBasedIssuance = initial.AllowCsrBasedIssuance
	}
	if dcl.BoolCanonicalize(des.AllowConfigBasedIssuance, initial.AllowConfigBasedIssuance) || dcl.IsZeroValue(des.AllowConfigBasedIssuance) {
		des.AllowConfigBasedIssuance = initial.AllowConfigBasedIssuance
	}

	return des
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedIssuanceModes(c *Client, des, nw *CertificateAuthorityCertificatePolicyAllowedIssuanceModes) *CertificateAuthorityCertificatePolicyAllowedIssuanceModes {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.AllowCsrBasedIssuance, nw.AllowCsrBasedIssuance) {
		nw.AllowCsrBasedIssuance = des.AllowCsrBasedIssuance
	}
	if dcl.BoolCanonicalize(des.AllowConfigBasedIssuance, nw.AllowConfigBasedIssuance) {
		nw.AllowConfigBasedIssuance = des.AllowConfigBasedIssuance
	}

	return nw
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedIssuanceModesSet(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedIssuanceModes) []CertificateAuthorityCertificatePolicyAllowedIssuanceModes {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityCertificatePolicyAllowedIssuanceModes
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityCertificatePolicyAllowedIssuanceModesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityCertificatePolicyAllowedIssuanceModesSlice(c *Client, des, nw []CertificateAuthorityCertificatePolicyAllowedIssuanceModes) []CertificateAuthorityCertificatePolicyAllowedIssuanceModes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityCertificatePolicyAllowedIssuanceModes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityCertificatePolicyAllowedIssuanceModes(c, &d, &n))
	}

	return items
}

func canonicalizeCertificateAuthorityIssuingOptions(des, initial *CertificateAuthorityIssuingOptions, opts ...dcl.ApplyOption) *CertificateAuthorityIssuingOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	if dcl.BoolCanonicalize(des.IncludeCaCertUrl, initial.IncludeCaCertUrl) || dcl.IsZeroValue(des.IncludeCaCertUrl) {
		des.IncludeCaCertUrl = initial.IncludeCaCertUrl
	}
	if dcl.BoolCanonicalize(des.IncludeCrlAccessUrl, initial.IncludeCrlAccessUrl) || dcl.IsZeroValue(des.IncludeCrlAccessUrl) {
		des.IncludeCrlAccessUrl = initial.IncludeCrlAccessUrl
	}

	return des
}

func canonicalizeNewCertificateAuthorityIssuingOptions(c *Client, des, nw *CertificateAuthorityIssuingOptions) *CertificateAuthorityIssuingOptions {
	if des == nil || nw == nil {
		return nw
	}

	if dcl.BoolCanonicalize(des.IncludeCaCertUrl, nw.IncludeCaCertUrl) {
		nw.IncludeCaCertUrl = des.IncludeCaCertUrl
	}
	if dcl.BoolCanonicalize(des.IncludeCrlAccessUrl, nw.IncludeCrlAccessUrl) {
		nw.IncludeCrlAccessUrl = des.IncludeCrlAccessUrl
	}

	return nw
}

func canonicalizeNewCertificateAuthorityIssuingOptionsSet(c *Client, des, nw []CertificateAuthorityIssuingOptions) []CertificateAuthorityIssuingOptions {
	if des == nil {
		return nw
	}
	var reorderedNew []CertificateAuthorityIssuingOptions
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareCertificateAuthorityIssuingOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewCertificateAuthorityIssuingOptionsSlice(c *Client, des, nw []CertificateAuthorityIssuingOptions) []CertificateAuthorityIssuingOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []CertificateAuthorityIssuingOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewCertificateAuthorityIssuingOptions(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffCertificateAuthority(c *Client, desired, actual *CertificateAuthority, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Config, actual.Config, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigNewStyle, EmptyObject: EmptyCertificateAuthorityConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Config")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Lifetime, actual.Lifetime, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Lifetime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeySpec, actual.KeySpec, dcl.Info{ObjectFunction: compareCertificateAuthorityKeySpecNewStyle, EmptyObject: EmptyCertificateAuthorityKeySpec, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeySpec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubordinateConfig, actual.SubordinateConfig, dcl.Info{ObjectFunction: compareCertificateAuthoritySubordinateConfigNewStyle, EmptyObject: EmptyCertificateAuthoritySubordinateConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubordinateConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Tier, actual.Tier, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Tier")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.Info{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PemCaCertificates, actual.PemCaCertificates, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PemCaCertificates")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CaCertificateDescriptions, actual.CaCertificateDescriptions, dcl.Info{OutputOnly: true, ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CaCertificateDescriptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GcsBucket, actual.GcsBucket, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GcsBucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AccessUrls, actual.AccessUrls, dcl.Info{OutputOnly: true, ObjectFunction: compareCertificateAuthorityAccessUrlsNewStyle, EmptyObject: EmptyCertificateAuthorityAccessUrls, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AccessUrls")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeleteTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CertificatePolicy, actual.CertificatePolicy, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicy, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CertificatePolicy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IssuingOptions, actual.IssuingOptions, dcl.Info{ObjectFunction: compareCertificateAuthorityIssuingOptionsNewStyle, EmptyObject: EmptyCertificateAuthorityIssuingOptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IssuingOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareCertificateAuthorityConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfig)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfig or *CertificateAuthorityConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfig)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SubjectConfig, actual.SubjectConfig, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigSubjectConfigNewStyle, EmptyObject: EmptyCertificateAuthorityConfigSubjectConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubjectConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PublicKey, actual.PublicKey, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigPublicKeyNewStyle, EmptyObject: EmptyCertificateAuthorityConfigPublicKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PublicKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReusableConfig, actual.ReusableConfig, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigReusableConfigNewStyle, EmptyObject: EmptyCertificateAuthorityConfigReusableConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReusableConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigSubjectConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigSubjectConfig)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigSubjectConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigSubjectConfig or *CertificateAuthorityConfigSubjectConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigSubjectConfig)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigSubjectConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigSubjectConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Subject, actual.Subject, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigSubjectConfigSubjectNewStyle, EmptyObject: EmptyCertificateAuthorityConfigSubjectConfigSubject, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Subject")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CommonName, actual.CommonName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CommonName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubjectAltName, actual.SubjectAltName, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigSubjectConfigSubjectAltNameNewStyle, EmptyObject: EmptyCertificateAuthorityConfigSubjectConfigSubjectAltName, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubjectAltName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigSubjectConfigSubjectNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigSubjectConfigSubject)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigSubjectConfigSubject)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigSubjectConfigSubject or *CertificateAuthorityConfigSubjectConfigSubject", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigSubjectConfigSubject)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigSubjectConfigSubject)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigSubjectConfigSubject", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CountryCode, actual.CountryCode, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CountryCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Organization, actual.Organization, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Organization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OrganizationalUnit, actual.OrganizationalUnit, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OrganizationalUnit")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Locality, actual.Locality, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Locality")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Province, actual.Province, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Province")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StreetAddress, actual.StreetAddress, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StreetAddress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PostalCode, actual.PostalCode, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PostalCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigSubjectConfigSubjectAltNameNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigSubjectConfigSubjectAltName)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigSubjectConfigSubjectAltName)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigSubjectConfigSubjectAltName or *CertificateAuthorityConfigSubjectConfigSubjectAltName", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigSubjectConfigSubjectAltName)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigSubjectConfigSubjectAltName)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigSubjectConfigSubjectAltName", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DnsNames, actual.DnsNames, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DnsNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uris, actual.Uris, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uris")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EmailAddresses, actual.EmailAddresses, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EmailAddresses")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPAddresses, actual.IPAddresses, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IpAddresses")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomSans, actual.CustomSans, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansNewStyle, EmptyObject: EmptyCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CustomSans")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans or *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectId, actual.ObjectId, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdNewStyle, EmptyObject: EmptyCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Critical, actual.Critical, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Critical")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId or *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectIdPath, actual.ObjectIdPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectIdPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigPublicKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigPublicKey)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigPublicKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigPublicKey or *CertificateAuthorityConfigPublicKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigPublicKey)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigPublicKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigPublicKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigReusableConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigReusableConfig)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigReusableConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfig or *CertificateAuthorityConfigReusableConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigReusableConfig)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigReusableConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ReusableConfig, actual.ReusableConfig, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReusableConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReusableConfigValues, actual.ReusableConfigValues, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigReusableConfigReusableConfigValuesNewStyle, EmptyObject: EmptyCertificateAuthorityConfigReusableConfigReusableConfigValues, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReusableConfigValues")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigReusableConfigReusableConfigValuesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigReusableConfigReusableConfigValues)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigReusableConfigReusableConfigValues)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValues or *CertificateAuthorityConfigReusableConfigReusableConfigValues", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigReusableConfigReusableConfigValues)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigReusableConfigReusableConfigValues)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValues", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KeyUsage, actual.KeyUsage, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageNewStyle, EmptyObject: EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyUsage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CaOptions, actual.CaOptions, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsNewStyle, EmptyObject: EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CaOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PolicyIds, actual.PolicyIds, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsNewStyle, EmptyObject: EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PolicyIds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AiaOcspServers, actual.AiaOcspServers, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AiaOcspServers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AdditionalExtensions, actual.AdditionalExtensions, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsNewStyle, EmptyObject: EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AdditionalExtensions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage or *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BaseKeyUsage, actual.BaseKeyUsage, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageNewStyle, EmptyObject: EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BaseKeyUsage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExtendedKeyUsage, actual.ExtendedKeyUsage, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageNewStyle, EmptyObject: EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExtendedKeyUsage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UnknownExtendedKeyUsages, actual.UnknownExtendedKeyUsages, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesNewStyle, EmptyObject: EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UnknownExtendedKeyUsages")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage or *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DigitalSignature, actual.DigitalSignature, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DigitalSignature")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContentCommitment, actual.ContentCommitment, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ContentCommitment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeyEncipherment, actual.KeyEncipherment, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyEncipherment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DataEncipherment, actual.DataEncipherment, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DataEncipherment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeyAgreement, actual.KeyAgreement, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyAgreement")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CertSign, actual.CertSign, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CertSign")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrlSign, actual.CrlSign, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrlSign")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EncipherOnly, actual.EncipherOnly, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EncipherOnly")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DecipherOnly, actual.DecipherOnly, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DecipherOnly")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage or *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ServerAuth, actual.ServerAuth, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServerAuth")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientAuth, actual.ClientAuth, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClientAuth")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CodeSigning, actual.CodeSigning, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CodeSigning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EmailProtection, actual.EmailProtection, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EmailProtection")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeStamping, actual.TimeStamping, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TimeStamping")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OcspSigning, actual.OcspSigning, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OcspSigning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages or *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectIdPath, actual.ObjectIdPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectIdPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions or *CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IsCa, actual.IsCa, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IsCa")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxIssuerPathLength, actual.MaxIssuerPathLength, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxIssuerPathLength")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds or *CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectIdPath, actual.ObjectIdPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectIdPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions or *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectId, actual.ObjectId, dcl.Info{ObjectFunction: compareCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdNewStyle, EmptyObject: EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Critical, actual.Critical, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Critical")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId or *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectIdPath, actual.ObjectIdPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectIdPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityKeySpecNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityKeySpec)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityKeySpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityKeySpec or *CertificateAuthorityKeySpec", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityKeySpec)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityKeySpec)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityKeySpec", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CloudKmsKeyVersion, actual.CloudKmsKeyVersion, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CloudKmsKeyVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Algorithm, actual.Algorithm, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Algorithm")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthoritySubordinateConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthoritySubordinateConfig)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthoritySubordinateConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthoritySubordinateConfig or *CertificateAuthoritySubordinateConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthoritySubordinateConfig)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthoritySubordinateConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthoritySubordinateConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CertificateAuthority, actual.CertificateAuthority, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CertificateAuthority")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PemIssuerChain, actual.PemIssuerChain, dcl.Info{ObjectFunction: compareCertificateAuthoritySubordinateConfigPemIssuerChainNewStyle, EmptyObject: EmptyCertificateAuthoritySubordinateConfigPemIssuerChain, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PemIssuerChain")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthoritySubordinateConfigPemIssuerChainNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthoritySubordinateConfigPemIssuerChain)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthoritySubordinateConfigPemIssuerChain)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthoritySubordinateConfigPemIssuerChain or *CertificateAuthoritySubordinateConfigPemIssuerChain", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthoritySubordinateConfigPemIssuerChain)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthoritySubordinateConfigPemIssuerChain)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthoritySubordinateConfigPemIssuerChain", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PemCertificates, actual.PemCertificates, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PemCertificates")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptions)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptions or *CertificateAuthorityCaCertificateDescriptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptions)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SubjectDescription, actual.SubjectDescription, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescription, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubjectDescription")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PublicKey, actual.PublicKey, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsPublicKeyNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsPublicKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PublicKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubjectKeyId, actual.SubjectKeyId, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsSubjectKeyId, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubjectKeyId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AuthorityKeyId, actual.AuthorityKeyId, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AuthorityKeyId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrlDistributionPoints, actual.CrlDistributionPoints, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrlDistributionPoints")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AiaIssuingCertificateUrls, actual.AiaIssuingCertificateUrls, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AiaIssuingCertificateUrls")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CertFingerprint, actual.CertFingerprint, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsCertFingerprintNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsCertFingerprint, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CertFingerprint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConfigValues, actual.ConfigValues, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsConfigValuesNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsConfigValues, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConfigValues")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsSubjectDescription)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsSubjectDescription)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsSubjectDescription or *CertificateAuthorityCaCertificateDescriptionsSubjectDescription", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsSubjectDescription)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsSubjectDescription)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsSubjectDescription", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Subject, actual.Subject, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Subject")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SubjectAltName, actual.SubjectAltName, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SubjectAltName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HexSerialNumber, actual.HexSerialNumber, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HexSerialNumber")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Lifetime, actual.Lifetime, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Lifetime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NotBeforeTime, actual.NotBeforeTime, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NotBeforeTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NotAfterTime, actual.NotAfterTime, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NotAfterTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CommonName, actual.CommonName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CommonName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject or *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CommonName, actual.CommonName, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CommonName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CountryCode, actual.CountryCode, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CountryCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Organization, actual.Organization, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Organization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OrganizationalUnit, actual.OrganizationalUnit, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OrganizationalUnit")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Locality, actual.Locality, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Locality")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Province, actual.Province, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Province")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StreetAddress, actual.StreetAddress, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StreetAddress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PostalCode, actual.PostalCode, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PostalCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName or *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DnsNames, actual.DnsNames, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DnsNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Uris, actual.Uris, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Uris")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EmailAddresses, actual.EmailAddresses, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EmailAddresses")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPAddresses, actual.IPAddresses, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IpAddresses")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomSans, actual.CustomSans, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CustomSans")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans or *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectId, actual.ObjectId, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Critical, actual.Critical, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Critical")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId or *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectIdPath, actual.ObjectIdPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectIdPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsPublicKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsPublicKey)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsPublicKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsPublicKey or *CertificateAuthorityCaCertificateDescriptionsPublicKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsPublicKey)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsPublicKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsPublicKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Key, actual.Key, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Key")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Format, actual.Format, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Format")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsSubjectKeyId)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsSubjectKeyId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsSubjectKeyId or *CertificateAuthorityCaCertificateDescriptionsSubjectKeyId", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsSubjectKeyId)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsSubjectKeyId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsSubjectKeyId", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KeyId, actual.KeyId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId or *CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KeyId, actual.KeyId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsCertFingerprintNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsCertFingerprint)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsCertFingerprint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsCertFingerprint or *CertificateAuthorityCaCertificateDescriptionsCertFingerprint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsCertFingerprint)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsCertFingerprint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsCertFingerprint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Sha256Hash, actual.Sha256Hash, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Sha256Hash")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsConfigValuesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsConfigValues)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsConfigValues)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValues or *CertificateAuthorityCaCertificateDescriptionsConfigValues", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsConfigValues)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsConfigValues)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValues", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KeyUsage, actual.KeyUsage, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyUsage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CaOptions, actual.CaOptions, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CaOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PolicyIds, actual.PolicyIds, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PolicyIds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AiaOcspServers, actual.AiaOcspServers, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AiaOcspServers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AdditionalExtensions, actual.AdditionalExtensions, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AdditionalExtensions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage or *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BaseKeyUsage, actual.BaseKeyUsage, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BaseKeyUsage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExtendedKeyUsage, actual.ExtendedKeyUsage, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExtendedKeyUsage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UnknownExtendedKeyUsages, actual.UnknownExtendedKeyUsages, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UnknownExtendedKeyUsages")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage or *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DigitalSignature, actual.DigitalSignature, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DigitalSignature")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContentCommitment, actual.ContentCommitment, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ContentCommitment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeyEncipherment, actual.KeyEncipherment, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyEncipherment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DataEncipherment, actual.DataEncipherment, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DataEncipherment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeyAgreement, actual.KeyAgreement, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyAgreement")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CertSign, actual.CertSign, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CertSign")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrlSign, actual.CrlSign, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrlSign")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EncipherOnly, actual.EncipherOnly, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EncipherOnly")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DecipherOnly, actual.DecipherOnly, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DecipherOnly")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage or *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ServerAuth, actual.ServerAuth, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServerAuth")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientAuth, actual.ClientAuth, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClientAuth")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CodeSigning, actual.CodeSigning, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CodeSigning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EmailProtection, actual.EmailProtection, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EmailProtection")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeStamping, actual.TimeStamping, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TimeStamping")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OcspSigning, actual.OcspSigning, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OcspSigning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages or *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectIdPath, actual.ObjectIdPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectIdPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions or *CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IsCa, actual.IsCa, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IsCa")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxIssuerPathLength, actual.MaxIssuerPathLength, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxIssuerPathLength")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds or *CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectIdPath, actual.ObjectIdPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectIdPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions or *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectId, actual.ObjectId, dcl.Info{ObjectFunction: compareCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdNewStyle, EmptyObject: EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Critical, actual.Critical, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Critical")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId or *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectIdPath, actual.ObjectIdPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectIdPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityAccessUrlsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityAccessUrls)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityAccessUrls)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityAccessUrls or *CertificateAuthorityAccessUrls", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityAccessUrls)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityAccessUrls)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityAccessUrls", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CaCertificateAccessUrl, actual.CaCertificateAccessUrl, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CaCertificateAccessUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrlAccessUrls, actual.CrlAccessUrls, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrlAccessUrls")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrlAccessUrl, actual.CrlAccessUrl, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrlAccessUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicy)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicy or *CertificateAuthorityCertificatePolicy", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicy)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicy)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicy", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowedConfigList, actual.AllowedConfigList, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyAllowedConfigListNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyAllowedConfigList, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedConfigList")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OverwriteConfigValues, actual.OverwriteConfigValues, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValues, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OverwriteConfigValues")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowedLocationsAndOrganizations, actual.AllowedLocationsAndOrganizations, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedLocationsAndOrganizations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowedCommonNames, actual.AllowedCommonNames, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedCommonNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowedSans, actual.AllowedSans, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyAllowedSansNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyAllowedSans, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedSans")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaximumLifetime, actual.MaximumLifetime, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaximumLifetime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowedIssuanceModes, actual.AllowedIssuanceModes, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyAllowedIssuanceModesNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyAllowedIssuanceModes, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedIssuanceModes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyAllowedConfigListNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyAllowedConfigList)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyAllowedConfigList)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigList or *CertificateAuthorityCertificatePolicyAllowedConfigList", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyAllowedConfigList)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyAllowedConfigList)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigList", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowedConfigValues, actual.AllowedConfigValues, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedConfigValues")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues or *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ReusableConfig, actual.ReusableConfig, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReusableConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReusableConfigValues, actual.ReusableConfigValues, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReusableConfigValues")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues or *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KeyUsage, actual.KeyUsage, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyUsage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CaOptions, actual.CaOptions, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CaOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PolicyIds, actual.PolicyIds, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PolicyIds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AiaOcspServers, actual.AiaOcspServers, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AiaOcspServers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AdditionalExtensions, actual.AdditionalExtensions, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AdditionalExtensions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage or *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BaseKeyUsage, actual.BaseKeyUsage, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BaseKeyUsage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExtendedKeyUsage, actual.ExtendedKeyUsage, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExtendedKeyUsage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UnknownExtendedKeyUsages, actual.UnknownExtendedKeyUsages, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UnknownExtendedKeyUsages")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage or *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DigitalSignature, actual.DigitalSignature, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DigitalSignature")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContentCommitment, actual.ContentCommitment, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ContentCommitment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeyEncipherment, actual.KeyEncipherment, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyEncipherment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DataEncipherment, actual.DataEncipherment, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DataEncipherment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeyAgreement, actual.KeyAgreement, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyAgreement")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CertSign, actual.CertSign, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CertSign")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrlSign, actual.CrlSign, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrlSign")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EncipherOnly, actual.EncipherOnly, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EncipherOnly")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DecipherOnly, actual.DecipherOnly, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DecipherOnly")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage or *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ServerAuth, actual.ServerAuth, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServerAuth")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientAuth, actual.ClientAuth, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClientAuth")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CodeSigning, actual.CodeSigning, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CodeSigning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EmailProtection, actual.EmailProtection, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EmailProtection")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeStamping, actual.TimeStamping, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TimeStamping")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OcspSigning, actual.OcspSigning, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OcspSigning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages or *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectIdPath, actual.ObjectIdPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectIdPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions or *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IsCa, actual.IsCa, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IsCa")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxIssuerPathLength, actual.MaxIssuerPathLength, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxIssuerPathLength")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds or *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectIdPath, actual.ObjectIdPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectIdPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions or *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectId, actual.ObjectId, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Critical, actual.Critical, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Critical")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId or *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectIdPath, actual.ObjectIdPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectIdPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyOverwriteConfigValues)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyOverwriteConfigValues)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValues or *CertificateAuthorityCertificatePolicyOverwriteConfigValues", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyOverwriteConfigValues)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyOverwriteConfigValues)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValues", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ReusableConfig, actual.ReusableConfig, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReusableConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReusableConfigValues, actual.ReusableConfigValues, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReusableConfigValues")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues or *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KeyUsage, actual.KeyUsage, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyUsage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CaOptions, actual.CaOptions, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CaOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PolicyIds, actual.PolicyIds, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PolicyIds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AiaOcspServers, actual.AiaOcspServers, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AiaOcspServers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AdditionalExtensions, actual.AdditionalExtensions, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AdditionalExtensions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage or *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.BaseKeyUsage, actual.BaseKeyUsage, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BaseKeyUsage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExtendedKeyUsage, actual.ExtendedKeyUsage, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExtendedKeyUsage")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UnknownExtendedKeyUsages, actual.UnknownExtendedKeyUsages, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UnknownExtendedKeyUsages")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage or *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DigitalSignature, actual.DigitalSignature, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DigitalSignature")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContentCommitment, actual.ContentCommitment, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ContentCommitment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeyEncipherment, actual.KeyEncipherment, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyEncipherment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DataEncipherment, actual.DataEncipherment, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DataEncipherment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KeyAgreement, actual.KeyAgreement, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KeyAgreement")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CertSign, actual.CertSign, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CertSign")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CrlSign, actual.CrlSign, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CrlSign")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EncipherOnly, actual.EncipherOnly, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EncipherOnly")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DecipherOnly, actual.DecipherOnly, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DecipherOnly")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage or *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ServerAuth, actual.ServerAuth, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServerAuth")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClientAuth, actual.ClientAuth, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClientAuth")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CodeSigning, actual.CodeSigning, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CodeSigning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EmailProtection, actual.EmailProtection, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EmailProtection")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TimeStamping, actual.TimeStamping, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TimeStamping")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OcspSigning, actual.OcspSigning, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OcspSigning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages or *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectIdPath, actual.ObjectIdPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectIdPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions or *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IsCa, actual.IsCa, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IsCa")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxIssuerPathLength, actual.MaxIssuerPathLength, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxIssuerPathLength")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds or *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectIdPath, actual.ObjectIdPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectIdPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions or *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectId, actual.ObjectId, dcl.Info{ObjectFunction: compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdNewStyle, EmptyObject: EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Critical, actual.Critical, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Critical")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Value, actual.Value, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Value")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId or *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ObjectIdPath, actual.ObjectIdPath, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ObjectIdPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations or *CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CountryCode, actual.CountryCode, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CountryCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Organization, actual.Organization, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Organization")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OrganizationalUnit, actual.OrganizationalUnit, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OrganizationalUnit")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Locality, actual.Locality, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Locality")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Province, actual.Province, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Province")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StreetAddress, actual.StreetAddress, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StreetAddress")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PostalCode, actual.PostalCode, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PostalCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyAllowedSansNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyAllowedSans)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyAllowedSans)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedSans or *CertificateAuthorityCertificatePolicyAllowedSans", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyAllowedSans)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyAllowedSans)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedSans", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowedDnsNames, actual.AllowedDnsNames, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedDnsNames")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowedUris, actual.AllowedUris, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedUris")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowedEmailAddresses, actual.AllowedEmailAddresses, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedEmailAddresses")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowedIps, actual.AllowedIps, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowedIps")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowGlobbingDnsWildcards, actual.AllowGlobbingDnsWildcards, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowGlobbingDnsWildcards")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowCustomSans, actual.AllowCustomSans, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowCustomSans")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityCertificatePolicyAllowedIssuanceModesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityCertificatePolicyAllowedIssuanceModes)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityCertificatePolicyAllowedIssuanceModes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedIssuanceModes or *CertificateAuthorityCertificatePolicyAllowedIssuanceModes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityCertificatePolicyAllowedIssuanceModes)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityCertificatePolicyAllowedIssuanceModes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityCertificatePolicyAllowedIssuanceModes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowCsrBasedIssuance, actual.AllowCsrBasedIssuance, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowCsrBasedIssuance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowConfigBasedIssuance, actual.AllowConfigBasedIssuance, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AllowConfigBasedIssuance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareCertificateAuthorityIssuingOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*CertificateAuthorityIssuingOptions)
	if !ok {
		desiredNotPointer, ok := d.(CertificateAuthorityIssuingOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityIssuingOptions or *CertificateAuthorityIssuingOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*CertificateAuthorityIssuingOptions)
	if !ok {
		actualNotPointer, ok := a.(CertificateAuthorityIssuingOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a CertificateAuthorityIssuingOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.IncludeCaCertUrl, actual.IncludeCaCertUrl, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IncludeCaCertUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IncludeCrlAccessUrl, actual.IncludeCrlAccessUrl, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IncludeCrlAccessUrl")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func (r *CertificateAuthority) getFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *CertificateAuthority) createFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *CertificateAuthority) deleteFields() (string, string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Location), dcl.ValueOrEmptyString(n.Name)
}

func (r *CertificateAuthority) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the CertificateAuthority resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *CertificateAuthority) marshal(c *Client) ([]byte, error) {
	m, err := expandCertificateAuthority(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling CertificateAuthority: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalCertificateAuthority decodes JSON responses into the CertificateAuthority resource schema.
func unmarshalCertificateAuthority(b []byte, c *Client) (*CertificateAuthority, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapCertificateAuthority(m, c)
}

func unmarshalMapCertificateAuthority(m map[string]interface{}, c *Client) (*CertificateAuthority, error) {

	return flattenCertificateAuthority(c, m), nil
}

// expandCertificateAuthority expands CertificateAuthority into a JSON request object.
func expandCertificateAuthority(c *Client, f *CertificateAuthority) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("projects/%s/locations/%s/certificateAuthorities/%s", f.Name, f.Project, f.Location, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}
	if v, err := expandCertificateAuthorityConfig(c, f.Config); err != nil {
		return nil, fmt.Errorf("error expanding Config into config: %w", err)
	} else if v != nil {
		m["config"] = v
	}
	if v := f.Lifetime; !dcl.IsEmptyValueIndirect(v) {
		m["lifetime"] = v
	}
	if v, err := expandCertificateAuthorityKeySpec(c, f.KeySpec); err != nil {
		return nil, fmt.Errorf("error expanding KeySpec into keySpec: %w", err)
	} else if v != nil {
		m["keySpec"] = v
	}
	if v, err := expandCertificateAuthoritySubordinateConfig(c, f.SubordinateConfig); err != nil {
		return nil, fmt.Errorf("error expanding SubordinateConfig into subordinateConfig: %w", err)
	} else if v != nil {
		m["subordinateConfig"] = v
	}
	if v := f.Tier; !dcl.IsEmptyValueIndirect(v) {
		m["tier"] = v
	}
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}
	if v := f.PemCaCertificates; !dcl.IsEmptyValueIndirect(v) {
		m["pemCaCertificates"] = v
	}
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsSlice(c, f.CaCertificateDescriptions); err != nil {
		return nil, fmt.Errorf("error expanding CaCertificateDescriptions into caCertificateDescriptions: %w", err)
	} else if v != nil {
		m["caCertificateDescriptions"] = v
	}
	if v := f.GcsBucket; !dcl.IsEmptyValueIndirect(v) {
		m["gcsBucket"] = v
	}
	if v, err := expandCertificateAuthorityAccessUrls(c, f.AccessUrls); err != nil {
		return nil, fmt.Errorf("error expanding AccessUrls into accessUrls: %w", err)
	} else if v != nil {
		m["accessUrls"] = v
	}
	if v := f.CreateTime; !dcl.IsEmptyValueIndirect(v) {
		m["createTime"] = v
	}
	if v := f.UpdateTime; !dcl.IsEmptyValueIndirect(v) {
		m["updateTime"] = v
	}
	if v := f.DeleteTime; !dcl.IsEmptyValueIndirect(v) {
		m["deleteTime"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicy(c, f.CertificatePolicy); err != nil {
		return nil, fmt.Errorf("error expanding CertificatePolicy into certificatePolicy: %w", err)
	} else if v != nil {
		m["certificatePolicy"] = v
	}
	if v, err := expandCertificateAuthorityIssuingOptions(c, f.IssuingOptions); err != nil {
		return nil, fmt.Errorf("error expanding IssuingOptions into issuingOptions: %w", err)
	} else if v != nil {
		m["issuingOptions"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if v != nil {
		m["location"] = v
	}

	return m, nil
}

// flattenCertificateAuthority flattens CertificateAuthority from a JSON request object into the
// CertificateAuthority type.
func flattenCertificateAuthority(c *Client, i interface{}) *CertificateAuthority {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &CertificateAuthority{}
	res.Name = dcl.FlattenString(m["name"])
	res.Type = flattenCertificateAuthorityTypeEnum(m["type"])
	res.Config = flattenCertificateAuthorityConfig(c, m["config"])
	res.Lifetime = dcl.FlattenString(m["lifetime"])
	res.KeySpec = flattenCertificateAuthorityKeySpec(c, m["keySpec"])
	res.SubordinateConfig = flattenCertificateAuthoritySubordinateConfig(c, m["subordinateConfig"])
	res.Tier = flattenCertificateAuthorityTierEnum(m["tier"])
	res.State = flattenCertificateAuthorityStateEnum(m["state"])
	res.PemCaCertificates = dcl.FlattenStringSlice(m["pemCaCertificates"])
	res.CaCertificateDescriptions = flattenCertificateAuthorityCaCertificateDescriptionsSlice(c, m["caCertificateDescriptions"])
	res.GcsBucket = dcl.FlattenString(m["gcsBucket"])
	res.AccessUrls = flattenCertificateAuthorityAccessUrls(c, m["accessUrls"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.DeleteTime = dcl.FlattenString(m["deleteTime"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.CertificatePolicy = flattenCertificateAuthorityCertificatePolicy(c, m["certificatePolicy"])
	res.IssuingOptions = flattenCertificateAuthorityIssuingOptions(c, m["issuingOptions"])
	res.Project = dcl.FlattenString(m["project"])
	res.Location = dcl.FlattenString(m["location"])

	return res
}

// expandCertificateAuthorityConfigMap expands the contents of CertificateAuthorityConfig into a JSON
// request object.
func expandCertificateAuthorityConfigMap(c *Client, f map[string]CertificateAuthorityConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigSlice expands the contents of CertificateAuthorityConfig into a JSON
// request object.
func expandCertificateAuthorityConfigSlice(c *Client, f []CertificateAuthorityConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigMap flattens the contents of CertificateAuthorityConfig from a JSON
// response object.
func flattenCertificateAuthorityConfigMap(c *Client, i interface{}) map[string]CertificateAuthorityConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfig{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfig{}
	}

	items := make(map[string]CertificateAuthorityConfig)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigSlice flattens the contents of CertificateAuthorityConfig from a JSON
// response object.
func flattenCertificateAuthorityConfigSlice(c *Client, i interface{}) []CertificateAuthorityConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfig{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfig{}
	}

	items := make([]CertificateAuthorityConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfig expands an instance of CertificateAuthorityConfig into a JSON
// request object.
func expandCertificateAuthorityConfig(c *Client, f *CertificateAuthorityConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityConfigSubjectConfig(c, f.SubjectConfig); err != nil {
		return nil, fmt.Errorf("error expanding SubjectConfig into subjectConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["subjectConfig"] = v
	}
	if v, err := expandCertificateAuthorityConfigPublicKey(c, f.PublicKey); err != nil {
		return nil, fmt.Errorf("error expanding PublicKey into publicKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["publicKey"] = v
	}
	if v, err := expandCertificateAuthorityConfigReusableConfig(c, f.ReusableConfig); err != nil {
		return nil, fmt.Errorf("error expanding ReusableConfig into reusableConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reusableConfig"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfig flattens an instance of CertificateAuthorityConfig from a JSON
// response object.
func flattenCertificateAuthorityConfig(c *Client, i interface{}) *CertificateAuthorityConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfig
	}
	r.SubjectConfig = flattenCertificateAuthorityConfigSubjectConfig(c, m["subjectConfig"])
	r.PublicKey = flattenCertificateAuthorityConfigPublicKey(c, m["publicKey"])
	r.ReusableConfig = flattenCertificateAuthorityConfigReusableConfig(c, m["reusableConfig"])

	return r
}

// expandCertificateAuthorityConfigSubjectConfigMap expands the contents of CertificateAuthorityConfigSubjectConfig into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfigMap(c *Client, f map[string]CertificateAuthorityConfigSubjectConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigSubjectConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigSubjectConfigSlice expands the contents of CertificateAuthorityConfigSubjectConfig into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfigSlice(c *Client, f []CertificateAuthorityConfigSubjectConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigSubjectConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigSubjectConfigMap flattens the contents of CertificateAuthorityConfigSubjectConfig from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfigMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigSubjectConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigSubjectConfig{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigSubjectConfig{}
	}

	items := make(map[string]CertificateAuthorityConfigSubjectConfig)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigSubjectConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigSubjectConfigSlice flattens the contents of CertificateAuthorityConfigSubjectConfig from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfigSlice(c *Client, i interface{}) []CertificateAuthorityConfigSubjectConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigSubjectConfig{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigSubjectConfig{}
	}

	items := make([]CertificateAuthorityConfigSubjectConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigSubjectConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigSubjectConfig expands an instance of CertificateAuthorityConfigSubjectConfig into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfig(c *Client, f *CertificateAuthorityConfigSubjectConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityConfigSubjectConfigSubject(c, f.Subject); err != nil {
		return nil, fmt.Errorf("error expanding Subject into subject: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["subject"] = v
	}
	if v := f.CommonName; !dcl.IsEmptyValueIndirect(v) {
		m["commonName"] = v
	}
	if v, err := expandCertificateAuthorityConfigSubjectConfigSubjectAltName(c, f.SubjectAltName); err != nil {
		return nil, fmt.Errorf("error expanding SubjectAltName into subjectAltName: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["subjectAltName"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigSubjectConfig flattens an instance of CertificateAuthorityConfigSubjectConfig from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfig(c *Client, i interface{}) *CertificateAuthorityConfigSubjectConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigSubjectConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigSubjectConfig
	}
	r.Subject = flattenCertificateAuthorityConfigSubjectConfigSubject(c, m["subject"])
	r.CommonName = dcl.FlattenString(m["commonName"])
	r.SubjectAltName = flattenCertificateAuthorityConfigSubjectConfigSubjectAltName(c, m["subjectAltName"])

	return r
}

// expandCertificateAuthorityConfigSubjectConfigSubjectMap expands the contents of CertificateAuthorityConfigSubjectConfigSubject into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfigSubjectMap(c *Client, f map[string]CertificateAuthorityConfigSubjectConfigSubject) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigSubjectConfigSubject(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigSubjectConfigSubjectSlice expands the contents of CertificateAuthorityConfigSubjectConfigSubject into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfigSubjectSlice(c *Client, f []CertificateAuthorityConfigSubjectConfigSubject) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigSubjectConfigSubject(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigSubjectConfigSubjectMap flattens the contents of CertificateAuthorityConfigSubjectConfigSubject from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfigSubjectMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigSubjectConfigSubject {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigSubjectConfigSubject{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigSubjectConfigSubject{}
	}

	items := make(map[string]CertificateAuthorityConfigSubjectConfigSubject)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigSubjectConfigSubject(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigSubjectConfigSubjectSlice flattens the contents of CertificateAuthorityConfigSubjectConfigSubject from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfigSubjectSlice(c *Client, i interface{}) []CertificateAuthorityConfigSubjectConfigSubject {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigSubjectConfigSubject{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigSubjectConfigSubject{}
	}

	items := make([]CertificateAuthorityConfigSubjectConfigSubject, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigSubjectConfigSubject(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigSubjectConfigSubject expands an instance of CertificateAuthorityConfigSubjectConfigSubject into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfigSubject(c *Client, f *CertificateAuthorityConfigSubjectConfigSubject) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CountryCode; !dcl.IsEmptyValueIndirect(v) {
		m["countryCode"] = v
	}
	if v := f.Organization; !dcl.IsEmptyValueIndirect(v) {
		m["organization"] = v
	}
	if v := f.OrganizationalUnit; !dcl.IsEmptyValueIndirect(v) {
		m["organizationalUnit"] = v
	}
	if v := f.Locality; !dcl.IsEmptyValueIndirect(v) {
		m["locality"] = v
	}
	if v := f.Province; !dcl.IsEmptyValueIndirect(v) {
		m["province"] = v
	}
	if v := f.StreetAddress; !dcl.IsEmptyValueIndirect(v) {
		m["streetAddress"] = v
	}
	if v := f.PostalCode; !dcl.IsEmptyValueIndirect(v) {
		m["postalCode"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigSubjectConfigSubject flattens an instance of CertificateAuthorityConfigSubjectConfigSubject from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfigSubject(c *Client, i interface{}) *CertificateAuthorityConfigSubjectConfigSubject {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigSubjectConfigSubject{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigSubjectConfigSubject
	}
	r.CountryCode = dcl.FlattenString(m["countryCode"])
	r.Organization = dcl.FlattenString(m["organization"])
	r.OrganizationalUnit = dcl.FlattenString(m["organizationalUnit"])
	r.Locality = dcl.FlattenString(m["locality"])
	r.Province = dcl.FlattenString(m["province"])
	r.StreetAddress = dcl.FlattenString(m["streetAddress"])
	r.PostalCode = dcl.FlattenString(m["postalCode"])

	return r
}

// expandCertificateAuthorityConfigSubjectConfigSubjectAltNameMap expands the contents of CertificateAuthorityConfigSubjectConfigSubjectAltName into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfigSubjectAltNameMap(c *Client, f map[string]CertificateAuthorityConfigSubjectConfigSubjectAltName) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigSubjectConfigSubjectAltName(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigSubjectConfigSubjectAltNameSlice expands the contents of CertificateAuthorityConfigSubjectConfigSubjectAltName into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfigSubjectAltNameSlice(c *Client, f []CertificateAuthorityConfigSubjectConfigSubjectAltName) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigSubjectConfigSubjectAltName(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameMap flattens the contents of CertificateAuthorityConfigSubjectConfigSubjectAltName from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigSubjectConfigSubjectAltName {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigSubjectConfigSubjectAltName{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigSubjectConfigSubjectAltName{}
	}

	items := make(map[string]CertificateAuthorityConfigSubjectConfigSubjectAltName)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigSubjectConfigSubjectAltName(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameSlice flattens the contents of CertificateAuthorityConfigSubjectConfigSubjectAltName from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameSlice(c *Client, i interface{}) []CertificateAuthorityConfigSubjectConfigSubjectAltName {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigSubjectConfigSubjectAltName{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigSubjectConfigSubjectAltName{}
	}

	items := make([]CertificateAuthorityConfigSubjectConfigSubjectAltName, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigSubjectConfigSubjectAltName(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigSubjectConfigSubjectAltName expands an instance of CertificateAuthorityConfigSubjectConfigSubjectAltName into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfigSubjectAltName(c *Client, f *CertificateAuthorityConfigSubjectConfigSubjectAltName) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DnsNames; !dcl.IsEmptyValueIndirect(v) {
		m["dnsNames"] = v
	}
	if v := f.Uris; !dcl.IsEmptyValueIndirect(v) {
		m["uris"] = v
	}
	if v := f.EmailAddresses; !dcl.IsEmptyValueIndirect(v) {
		m["emailAddresses"] = v
	}
	if v := f.IPAddresses; !dcl.IsEmptyValueIndirect(v) {
		m["ipAddresses"] = v
	}
	if v, err := expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansSlice(c, f.CustomSans); err != nil {
		return nil, fmt.Errorf("error expanding CustomSans into customSans: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["customSans"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigSubjectConfigSubjectAltName flattens an instance of CertificateAuthorityConfigSubjectConfigSubjectAltName from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfigSubjectAltName(c *Client, i interface{}) *CertificateAuthorityConfigSubjectConfigSubjectAltName {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigSubjectConfigSubjectAltName{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigSubjectConfigSubjectAltName
	}
	r.DnsNames = dcl.FlattenStringSlice(m["dnsNames"])
	r.Uris = dcl.FlattenStringSlice(m["uris"])
	r.EmailAddresses = dcl.FlattenStringSlice(m["emailAddresses"])
	r.IPAddresses = dcl.FlattenStringSlice(m["ipAddresses"])
	r.CustomSans = flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansSlice(c, m["customSans"])

	return r
}

// expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansMap expands the contents of CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansMap(c *Client, f map[string]CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansSlice expands the contents of CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansSlice(c *Client, f []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansMap flattens the contents of CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans{}
	}

	items := make(map[string]CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansSlice flattens the contents of CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansSlice(c *Client, i interface{}) []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans{}
	}

	items := make([]CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans expands an instance of CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(c *Client, f *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(c, f.ObjectId); err != nil {
		return nil, fmt.Errorf("error expanding ObjectId into objectId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["objectId"] = v
	}
	if v := f.Critical; !dcl.IsEmptyValueIndirect(v) {
		m["critical"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans flattens an instance of CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(c *Client, i interface{}) *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans
	}
	r.ObjectId = flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(c, m["objectId"])
	r.Critical = dcl.FlattenBool(m["critical"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdMap expands the contents of CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdMap(c *Client, f map[string]CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdSlice expands the contents of CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdSlice(c *Client, f []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdMap flattens the contents of CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{}
	}

	items := make(map[string]CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdSlice flattens the contents of CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdSlice(c *Client, i interface{}) []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{}
	}

	items := make([]CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId expands an instance of CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId into a JSON
// request object.
func expandCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(c *Client, f *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ObjectIdPath; !dcl.IsEmptyValueIndirect(v) {
		m["objectIdPath"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId flattens an instance of CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId from a JSON
// response object.
func flattenCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(c *Client, i interface{}) *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId
	}
	r.ObjectIdPath = dcl.FlattenIntSlice(m["objectIdPath"])

	return r
}

// expandCertificateAuthorityConfigPublicKeyMap expands the contents of CertificateAuthorityConfigPublicKey into a JSON
// request object.
func expandCertificateAuthorityConfigPublicKeyMap(c *Client, f map[string]CertificateAuthorityConfigPublicKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigPublicKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigPublicKeySlice expands the contents of CertificateAuthorityConfigPublicKey into a JSON
// request object.
func expandCertificateAuthorityConfigPublicKeySlice(c *Client, f []CertificateAuthorityConfigPublicKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigPublicKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigPublicKeyMap flattens the contents of CertificateAuthorityConfigPublicKey from a JSON
// response object.
func flattenCertificateAuthorityConfigPublicKeyMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigPublicKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigPublicKey{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigPublicKey{}
	}

	items := make(map[string]CertificateAuthorityConfigPublicKey)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigPublicKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigPublicKeySlice flattens the contents of CertificateAuthorityConfigPublicKey from a JSON
// response object.
func flattenCertificateAuthorityConfigPublicKeySlice(c *Client, i interface{}) []CertificateAuthorityConfigPublicKey {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigPublicKey{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigPublicKey{}
	}

	items := make([]CertificateAuthorityConfigPublicKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigPublicKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigPublicKey expands an instance of CertificateAuthorityConfigPublicKey into a JSON
// request object.
func expandCertificateAuthorityConfigPublicKey(c *Client, f *CertificateAuthorityConfigPublicKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigPublicKey flattens an instance of CertificateAuthorityConfigPublicKey from a JSON
// response object.
func flattenCertificateAuthorityConfigPublicKey(c *Client, i interface{}) *CertificateAuthorityConfigPublicKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigPublicKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigPublicKey
	}
	r.Key = dcl.FlattenString(m["key"])
	r.Type = flattenCertificateAuthorityConfigPublicKeyTypeEnum(m["type"])

	return r
}

// expandCertificateAuthorityConfigReusableConfigMap expands the contents of CertificateAuthorityConfigReusableConfig into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigMap(c *Client, f map[string]CertificateAuthorityConfigReusableConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigReusableConfigSlice expands the contents of CertificateAuthorityConfigReusableConfig into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigSlice(c *Client, f []CertificateAuthorityConfigReusableConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigReusableConfigMap flattens the contents of CertificateAuthorityConfigReusableConfig from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigReusableConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigReusableConfig{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigReusableConfig{}
	}

	items := make(map[string]CertificateAuthorityConfigReusableConfig)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigReusableConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigReusableConfigSlice flattens the contents of CertificateAuthorityConfigReusableConfig from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigSlice(c *Client, i interface{}) []CertificateAuthorityConfigReusableConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigReusableConfig{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigReusableConfig{}
	}

	items := make([]CertificateAuthorityConfigReusableConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigReusableConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigReusableConfig expands an instance of CertificateAuthorityConfigReusableConfig into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfig(c *Client, f *CertificateAuthorityConfigReusableConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ReusableConfig; !dcl.IsEmptyValueIndirect(v) {
		m["reusableConfig"] = v
	}
	if v, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValues(c, f.ReusableConfigValues); err != nil {
		return nil, fmt.Errorf("error expanding ReusableConfigValues into reusableConfigValues: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reusableConfigValues"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigReusableConfig flattens an instance of CertificateAuthorityConfigReusableConfig from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfig(c *Client, i interface{}) *CertificateAuthorityConfigReusableConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigReusableConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigReusableConfig
	}
	r.ReusableConfig = dcl.FlattenString(m["reusableConfig"])
	r.ReusableConfigValues = flattenCertificateAuthorityConfigReusableConfigReusableConfigValues(c, m["reusableConfigValues"])

	return r
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesMap expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValues into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesMap(c *Client, f map[string]CertificateAuthorityConfigReusableConfigReusableConfigValues) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValues(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesSlice expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValues into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesSlice(c *Client, f []CertificateAuthorityConfigReusableConfigReusableConfigValues) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValues(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesMap flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValues from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigReusableConfigReusableConfigValues {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValues{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValues{}
	}

	items := make(map[string]CertificateAuthorityConfigReusableConfigReusableConfigValues)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigReusableConfigReusableConfigValues(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesSlice flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValues from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesSlice(c *Client, i interface{}) []CertificateAuthorityConfigReusableConfigReusableConfigValues {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValues{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValues{}
	}

	items := make([]CertificateAuthorityConfigReusableConfigReusableConfigValues, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigReusableConfigReusableConfigValues(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValues expands an instance of CertificateAuthorityConfigReusableConfigReusableConfigValues into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValues(c *Client, f *CertificateAuthorityConfigReusableConfigReusableConfigValues) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(c, f.KeyUsage); err != nil {
		return nil, fmt.Errorf("error expanding KeyUsage into keyUsage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["keyUsage"] = v
	}
	if v, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(c, f.CaOptions); err != nil {
		return nil, fmt.Errorf("error expanding CaOptions into caOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["caOptions"] = v
	}
	if v, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsSlice(c, f.PolicyIds); err != nil {
		return nil, fmt.Errorf("error expanding PolicyIds into policyIds: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["policyIds"] = v
	}
	if v := f.AiaOcspServers; !dcl.IsEmptyValueIndirect(v) {
		m["aiaOcspServers"] = v
	}
	if v, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsSlice(c, f.AdditionalExtensions); err != nil {
		return nil, fmt.Errorf("error expanding AdditionalExtensions into additionalExtensions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["additionalExtensions"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValues flattens an instance of CertificateAuthorityConfigReusableConfigReusableConfigValues from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValues(c *Client, i interface{}) *CertificateAuthorityConfigReusableConfigReusableConfigValues {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigReusableConfigReusableConfigValues{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigReusableConfigReusableConfigValues
	}
	r.KeyUsage = flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(c, m["keyUsage"])
	r.CaOptions = flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(c, m["caOptions"])
	r.PolicyIds = flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsSlice(c, m["policyIds"])
	r.AiaOcspServers = dcl.FlattenStringSlice(m["aiaOcspServers"])
	r.AdditionalExtensions = flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsSlice(c, m["additionalExtensions"])

	return r
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageMap expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageMap(c *Client, f map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageSlice expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageSlice(c *Client, f []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageMap flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage{}
	}

	items := make(map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageSlice flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageSlice(c *Client, i interface{}) []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage{}
	}

	items := make([]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage expands an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(c *Client, f *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(c, f.BaseKeyUsage); err != nil {
		return nil, fmt.Errorf("error expanding BaseKeyUsage into baseKeyUsage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["baseKeyUsage"] = v
	}
	if v, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(c, f.ExtendedKeyUsage); err != nil {
		return nil, fmt.Errorf("error expanding ExtendedKeyUsage into extendedKeyUsage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["extendedKeyUsage"] = v
	}
	if v, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c, f.UnknownExtendedKeyUsages); err != nil {
		return nil, fmt.Errorf("error expanding UnknownExtendedKeyUsages into unknownExtendedKeyUsages: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["unknownExtendedKeyUsages"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage flattens an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(c *Client, i interface{}) *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage
	}
	r.BaseKeyUsage = flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(c, m["baseKeyUsage"])
	r.ExtendedKeyUsage = flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(c, m["extendedKeyUsage"])
	r.UnknownExtendedKeyUsages = flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c, m["unknownExtendedKeyUsages"])

	return r
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageMap expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageMap(c *Client, f map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageSlice expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageSlice(c *Client, f []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageMap flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage{}
	}

	items := make(map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageSlice flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageSlice(c *Client, i interface{}) []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage{}
	}

	items := make([]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage expands an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(c *Client, f *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DigitalSignature; !dcl.IsEmptyValueIndirect(v) {
		m["digitalSignature"] = v
	}
	if v := f.ContentCommitment; !dcl.IsEmptyValueIndirect(v) {
		m["contentCommitment"] = v
	}
	if v := f.KeyEncipherment; !dcl.IsEmptyValueIndirect(v) {
		m["keyEncipherment"] = v
	}
	if v := f.DataEncipherment; !dcl.IsEmptyValueIndirect(v) {
		m["dataEncipherment"] = v
	}
	if v := f.KeyAgreement; !dcl.IsEmptyValueIndirect(v) {
		m["keyAgreement"] = v
	}
	if v := f.CertSign; !dcl.IsEmptyValueIndirect(v) {
		m["certSign"] = v
	}
	if v := f.CrlSign; !dcl.IsEmptyValueIndirect(v) {
		m["crlSign"] = v
	}
	if v := f.EncipherOnly; !dcl.IsEmptyValueIndirect(v) {
		m["encipherOnly"] = v
	}
	if v := f.DecipherOnly; !dcl.IsEmptyValueIndirect(v) {
		m["decipherOnly"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage flattens an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(c *Client, i interface{}) *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage
	}
	r.DigitalSignature = dcl.FlattenBool(m["digitalSignature"])
	r.ContentCommitment = dcl.FlattenBool(m["contentCommitment"])
	r.KeyEncipherment = dcl.FlattenBool(m["keyEncipherment"])
	r.DataEncipherment = dcl.FlattenBool(m["dataEncipherment"])
	r.KeyAgreement = dcl.FlattenBool(m["keyAgreement"])
	r.CertSign = dcl.FlattenBool(m["certSign"])
	r.CrlSign = dcl.FlattenBool(m["crlSign"])
	r.EncipherOnly = dcl.FlattenBool(m["encipherOnly"])
	r.DecipherOnly = dcl.FlattenBool(m["decipherOnly"])

	return r
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageMap expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageMap(c *Client, f map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageSlice expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageSlice(c *Client, f []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageMap flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage{}
	}

	items := make(map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageSlice flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageSlice(c *Client, i interface{}) []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage{}
	}

	items := make([]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage expands an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(c *Client, f *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ServerAuth; !dcl.IsEmptyValueIndirect(v) {
		m["serverAuth"] = v
	}
	if v := f.ClientAuth; !dcl.IsEmptyValueIndirect(v) {
		m["clientAuth"] = v
	}
	if v := f.CodeSigning; !dcl.IsEmptyValueIndirect(v) {
		m["codeSigning"] = v
	}
	if v := f.EmailProtection; !dcl.IsEmptyValueIndirect(v) {
		m["emailProtection"] = v
	}
	if v := f.TimeStamping; !dcl.IsEmptyValueIndirect(v) {
		m["timeStamping"] = v
	}
	if v := f.OcspSigning; !dcl.IsEmptyValueIndirect(v) {
		m["ocspSigning"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage flattens an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(c *Client, i interface{}) *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage
	}
	r.ServerAuth = dcl.FlattenBool(m["serverAuth"])
	r.ClientAuth = dcl.FlattenBool(m["clientAuth"])
	r.CodeSigning = dcl.FlattenBool(m["codeSigning"])
	r.EmailProtection = dcl.FlattenBool(m["emailProtection"])
	r.TimeStamping = dcl.FlattenBool(m["timeStamping"])
	r.OcspSigning = dcl.FlattenBool(m["ocspSigning"])

	return r
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesMap expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesMap(c *Client, f map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c *Client, f []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesMap flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	items := make(map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c *Client, i interface{}) []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	items := make([]CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages expands an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c *Client, f *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ObjectIdPath; !dcl.IsEmptyValueIndirect(v) {
		m["objectIdPath"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages flattens an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c *Client, i interface{}) *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	}
	r.ObjectIdPath = dcl.FlattenIntSlice(m["objectIdPath"])

	return r
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsMap expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsMap(c *Client, f map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsSlice expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsSlice(c *Client, f []CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsMap flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions{}
	}

	items := make(map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsSlice flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsSlice(c *Client, i interface{}) []CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions{}
	}

	items := make([]CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions expands an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(c *Client, f *CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IsCa; !dcl.IsEmptyValueIndirect(v) {
		m["isCa"] = v
	}
	if v := f.MaxIssuerPathLength; !dcl.IsEmptyValueIndirect(v) {
		m["maxIssuerPathLength"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions flattens an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(c *Client, i interface{}) *CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions
	}
	r.IsCa = dcl.FlattenBool(m["isCa"])
	r.MaxIssuerPathLength = dcl.FlattenInteger(m["maxIssuerPathLength"])

	return r
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsMap expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsMap(c *Client, f map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsSlice expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsSlice(c *Client, f []CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsMap flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds{}
	}

	items := make(map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsSlice flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsSlice(c *Client, i interface{}) []CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds{}
	}

	items := make([]CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds expands an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds(c *Client, f *CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ObjectIdPath; !dcl.IsEmptyValueIndirect(v) {
		m["objectIdPath"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds flattens an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds(c *Client, i interface{}) *CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds
	}
	r.ObjectIdPath = dcl.FlattenIntSlice(m["objectIdPath"])

	return r
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsMap expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsMap(c *Client, f map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsSlice expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsSlice(c *Client, f []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsMap flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions{}
	}

	items := make(map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsSlice flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsSlice(c *Client, i interface{}) []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions{}
	}

	items := make([]CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions expands an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions(c *Client, f *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(c, f.ObjectId); err != nil {
		return nil, fmt.Errorf("error expanding ObjectId into objectId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["objectId"] = v
	}
	if v := f.Critical; !dcl.IsEmptyValueIndirect(v) {
		m["critical"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions flattens an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions(c *Client, i interface{}) *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions
	}
	r.ObjectId = flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(c, m["objectId"])
	r.Critical = dcl.FlattenBool(m["critical"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdMap expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdMap(c *Client, f map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdSlice expands the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdSlice(c *Client, f []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdMap flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdMap(c *Client, i interface{}) map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId{}
	}

	items := make(map[string]CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdSlice flattens the contents of CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdSlice(c *Client, i interface{}) []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId{}
	}

	items := make([]CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId expands an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId into a JSON
// request object.
func expandCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(c *Client, f *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ObjectIdPath; !dcl.IsEmptyValueIndirect(v) {
		m["objectIdPath"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId flattens an instance of CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId from a JSON
// response object.
func flattenCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(c *Client, i interface{}) *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId
	}
	r.ObjectIdPath = dcl.FlattenIntSlice(m["objectIdPath"])

	return r
}

// expandCertificateAuthorityKeySpecMap expands the contents of CertificateAuthorityKeySpec into a JSON
// request object.
func expandCertificateAuthorityKeySpecMap(c *Client, f map[string]CertificateAuthorityKeySpec) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityKeySpec(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityKeySpecSlice expands the contents of CertificateAuthorityKeySpec into a JSON
// request object.
func expandCertificateAuthorityKeySpecSlice(c *Client, f []CertificateAuthorityKeySpec) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityKeySpec(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityKeySpecMap flattens the contents of CertificateAuthorityKeySpec from a JSON
// response object.
func flattenCertificateAuthorityKeySpecMap(c *Client, i interface{}) map[string]CertificateAuthorityKeySpec {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityKeySpec{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityKeySpec{}
	}

	items := make(map[string]CertificateAuthorityKeySpec)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityKeySpec(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityKeySpecSlice flattens the contents of CertificateAuthorityKeySpec from a JSON
// response object.
func flattenCertificateAuthorityKeySpecSlice(c *Client, i interface{}) []CertificateAuthorityKeySpec {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityKeySpec{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityKeySpec{}
	}

	items := make([]CertificateAuthorityKeySpec, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityKeySpec(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityKeySpec expands an instance of CertificateAuthorityKeySpec into a JSON
// request object.
func expandCertificateAuthorityKeySpec(c *Client, f *CertificateAuthorityKeySpec) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CloudKmsKeyVersion; !dcl.IsEmptyValueIndirect(v) {
		m["cloudKmsKeyVersion"] = v
	}
	if v := f.Algorithm; !dcl.IsEmptyValueIndirect(v) {
		m["algorithm"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityKeySpec flattens an instance of CertificateAuthorityKeySpec from a JSON
// response object.
func flattenCertificateAuthorityKeySpec(c *Client, i interface{}) *CertificateAuthorityKeySpec {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityKeySpec{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityKeySpec
	}
	r.CloudKmsKeyVersion = dcl.FlattenString(m["cloudKmsKeyVersion"])
	r.Algorithm = flattenCertificateAuthorityKeySpecAlgorithmEnum(m["algorithm"])

	return r
}

// expandCertificateAuthoritySubordinateConfigMap expands the contents of CertificateAuthoritySubordinateConfig into a JSON
// request object.
func expandCertificateAuthoritySubordinateConfigMap(c *Client, f map[string]CertificateAuthoritySubordinateConfig) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthoritySubordinateConfig(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthoritySubordinateConfigSlice expands the contents of CertificateAuthoritySubordinateConfig into a JSON
// request object.
func expandCertificateAuthoritySubordinateConfigSlice(c *Client, f []CertificateAuthoritySubordinateConfig) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthoritySubordinateConfig(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthoritySubordinateConfigMap flattens the contents of CertificateAuthoritySubordinateConfig from a JSON
// response object.
func flattenCertificateAuthoritySubordinateConfigMap(c *Client, i interface{}) map[string]CertificateAuthoritySubordinateConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthoritySubordinateConfig{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthoritySubordinateConfig{}
	}

	items := make(map[string]CertificateAuthoritySubordinateConfig)
	for k, item := range a {
		items[k] = *flattenCertificateAuthoritySubordinateConfig(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthoritySubordinateConfigSlice flattens the contents of CertificateAuthoritySubordinateConfig from a JSON
// response object.
func flattenCertificateAuthoritySubordinateConfigSlice(c *Client, i interface{}) []CertificateAuthoritySubordinateConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthoritySubordinateConfig{}
	}

	if len(a) == 0 {
		return []CertificateAuthoritySubordinateConfig{}
	}

	items := make([]CertificateAuthoritySubordinateConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthoritySubordinateConfig(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthoritySubordinateConfig expands an instance of CertificateAuthoritySubordinateConfig into a JSON
// request object.
func expandCertificateAuthoritySubordinateConfig(c *Client, f *CertificateAuthoritySubordinateConfig) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CertificateAuthority; !dcl.IsEmptyValueIndirect(v) {
		m["certificateAuthority"] = v
	}
	if v, err := expandCertificateAuthoritySubordinateConfigPemIssuerChain(c, f.PemIssuerChain); err != nil {
		return nil, fmt.Errorf("error expanding PemIssuerChain into pemIssuerChain: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["pemIssuerChain"] = v
	}

	return m, nil
}

// flattenCertificateAuthoritySubordinateConfig flattens an instance of CertificateAuthoritySubordinateConfig from a JSON
// response object.
func flattenCertificateAuthoritySubordinateConfig(c *Client, i interface{}) *CertificateAuthoritySubordinateConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthoritySubordinateConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthoritySubordinateConfig
	}
	r.CertificateAuthority = dcl.FlattenString(m["certificateAuthority"])
	r.PemIssuerChain = flattenCertificateAuthoritySubordinateConfigPemIssuerChain(c, m["pemIssuerChain"])

	return r
}

// expandCertificateAuthoritySubordinateConfigPemIssuerChainMap expands the contents of CertificateAuthoritySubordinateConfigPemIssuerChain into a JSON
// request object.
func expandCertificateAuthoritySubordinateConfigPemIssuerChainMap(c *Client, f map[string]CertificateAuthoritySubordinateConfigPemIssuerChain) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthoritySubordinateConfigPemIssuerChain(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthoritySubordinateConfigPemIssuerChainSlice expands the contents of CertificateAuthoritySubordinateConfigPemIssuerChain into a JSON
// request object.
func expandCertificateAuthoritySubordinateConfigPemIssuerChainSlice(c *Client, f []CertificateAuthoritySubordinateConfigPemIssuerChain) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthoritySubordinateConfigPemIssuerChain(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthoritySubordinateConfigPemIssuerChainMap flattens the contents of CertificateAuthoritySubordinateConfigPemIssuerChain from a JSON
// response object.
func flattenCertificateAuthoritySubordinateConfigPemIssuerChainMap(c *Client, i interface{}) map[string]CertificateAuthoritySubordinateConfigPemIssuerChain {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthoritySubordinateConfigPemIssuerChain{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthoritySubordinateConfigPemIssuerChain{}
	}

	items := make(map[string]CertificateAuthoritySubordinateConfigPemIssuerChain)
	for k, item := range a {
		items[k] = *flattenCertificateAuthoritySubordinateConfigPemIssuerChain(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthoritySubordinateConfigPemIssuerChainSlice flattens the contents of CertificateAuthoritySubordinateConfigPemIssuerChain from a JSON
// response object.
func flattenCertificateAuthoritySubordinateConfigPemIssuerChainSlice(c *Client, i interface{}) []CertificateAuthoritySubordinateConfigPemIssuerChain {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthoritySubordinateConfigPemIssuerChain{}
	}

	if len(a) == 0 {
		return []CertificateAuthoritySubordinateConfigPemIssuerChain{}
	}

	items := make([]CertificateAuthoritySubordinateConfigPemIssuerChain, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthoritySubordinateConfigPemIssuerChain(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthoritySubordinateConfigPemIssuerChain expands an instance of CertificateAuthoritySubordinateConfigPemIssuerChain into a JSON
// request object.
func expandCertificateAuthoritySubordinateConfigPemIssuerChain(c *Client, f *CertificateAuthoritySubordinateConfigPemIssuerChain) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PemCertificates; !dcl.IsEmptyValueIndirect(v) {
		m["pemCertificates"] = v
	}

	return m, nil
}

// flattenCertificateAuthoritySubordinateConfigPemIssuerChain flattens an instance of CertificateAuthoritySubordinateConfigPemIssuerChain from a JSON
// response object.
func flattenCertificateAuthoritySubordinateConfigPemIssuerChain(c *Client, i interface{}) *CertificateAuthoritySubordinateConfigPemIssuerChain {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthoritySubordinateConfigPemIssuerChain{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthoritySubordinateConfigPemIssuerChain
	}
	r.PemCertificates = dcl.FlattenStringSlice(m["pemCertificates"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsMap expands the contents of CertificateAuthorityCaCertificateDescriptions into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsSlice expands the contents of CertificateAuthorityCaCertificateDescriptions into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsMap flattens the contents of CertificateAuthorityCaCertificateDescriptions from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptions{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptions{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptions)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsSlice flattens the contents of CertificateAuthorityCaCertificateDescriptions from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptions {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptions{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptions{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptions expands an instance of CertificateAuthorityCaCertificateDescriptions into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptions(c *Client, f *CertificateAuthorityCaCertificateDescriptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescription(c, f.SubjectDescription); err != nil {
		return nil, fmt.Errorf("error expanding SubjectDescription into subjectDescription: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["subjectDescription"] = v
	}
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsPublicKey(c, f.PublicKey); err != nil {
		return nil, fmt.Errorf("error expanding PublicKey into publicKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["publicKey"] = v
	}
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(c, f.SubjectKeyId); err != nil {
		return nil, fmt.Errorf("error expanding SubjectKeyId into subjectKeyId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["subjectKeyId"] = v
	}
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(c, f.AuthorityKeyId); err != nil {
		return nil, fmt.Errorf("error expanding AuthorityKeyId into authorityKeyId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["authorityKeyId"] = v
	}
	if v := f.CrlDistributionPoints; !dcl.IsEmptyValueIndirect(v) {
		m["crlDistributionPoints"] = v
	}
	if v := f.AiaIssuingCertificateUrls; !dcl.IsEmptyValueIndirect(v) {
		m["aiaIssuingCertificateUrls"] = v
	}
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsCertFingerprint(c, f.CertFingerprint); err != nil {
		return nil, fmt.Errorf("error expanding CertFingerprint into certFingerprint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["certFingerprint"] = v
	}
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValues(c, f.ConfigValues); err != nil {
		return nil, fmt.Errorf("error expanding ConfigValues into configValues: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["configValues"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptions flattens an instance of CertificateAuthorityCaCertificateDescriptions from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptions(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptions
	}
	r.SubjectDescription = flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescription(c, m["subjectDescription"])
	r.PublicKey = flattenCertificateAuthorityCaCertificateDescriptionsPublicKey(c, m["publicKey"])
	r.SubjectKeyId = flattenCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(c, m["subjectKeyId"])
	r.AuthorityKeyId = flattenCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(c, m["authorityKeyId"])
	r.CrlDistributionPoints = dcl.FlattenStringSlice(m["crlDistributionPoints"])
	r.AiaIssuingCertificateUrls = dcl.FlattenStringSlice(m["aiaIssuingCertificateUrls"])
	r.CertFingerprint = flattenCertificateAuthorityCaCertificateDescriptionsCertFingerprint(c, m["certFingerprint"])
	r.ConfigValues = flattenCertificateAuthorityCaCertificateDescriptionsConfigValues(c, m["configValues"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionMap expands the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescription into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescription) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescription(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescription into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsSubjectDescription) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescription(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescription from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescription{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescription{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescription)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescription(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescription from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsSubjectDescription{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsSubjectDescription{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsSubjectDescription, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescription(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescription expands an instance of CertificateAuthorityCaCertificateDescriptionsSubjectDescription into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescription(c *Client, f *CertificateAuthorityCaCertificateDescriptionsSubjectDescription) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(c, f.Subject); err != nil {
		return nil, fmt.Errorf("error expanding Subject into subject: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["subject"] = v
	}
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(c, f.SubjectAltName); err != nil {
		return nil, fmt.Errorf("error expanding SubjectAltName into subjectAltName: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["subjectAltName"] = v
	}
	if v := f.HexSerialNumber; !dcl.IsEmptyValueIndirect(v) {
		m["hexSerialNumber"] = v
	}
	if v := f.Lifetime; !dcl.IsEmptyValueIndirect(v) {
		m["lifetime"] = v
	}
	if v := f.NotBeforeTime; !dcl.IsEmptyValueIndirect(v) {
		m["notBeforeTime"] = v
	}
	if v := f.NotAfterTime; !dcl.IsEmptyValueIndirect(v) {
		m["notAfterTime"] = v
	}
	if v := f.CommonName; !dcl.IsEmptyValueIndirect(v) {
		m["commonName"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescription flattens an instance of CertificateAuthorityCaCertificateDescriptionsSubjectDescription from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescription(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsSubjectDescription {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsSubjectDescription{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescription
	}
	r.Subject = flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(c, m["subject"])
	r.SubjectAltName = flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(c, m["subjectAltName"])
	r.HexSerialNumber = dcl.FlattenString(m["hexSerialNumber"])
	r.Lifetime = dcl.FlattenString(m["lifetime"])
	r.NotBeforeTime = dcl.FlattenString(m["notBeforeTime"])
	r.NotAfterTime = dcl.FlattenString(m["notAfterTime"])
	r.CommonName = dcl.FlattenString(m["commonName"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectMap expands the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject expands an instance of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(c *Client, f *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CommonName; !dcl.IsEmptyValueIndirect(v) {
		m["commonName"] = v
	}
	if v := f.CountryCode; !dcl.IsEmptyValueIndirect(v) {
		m["countryCode"] = v
	}
	if v := f.Organization; !dcl.IsEmptyValueIndirect(v) {
		m["organization"] = v
	}
	if v := f.OrganizationalUnit; !dcl.IsEmptyValueIndirect(v) {
		m["organizationalUnit"] = v
	}
	if v := f.Locality; !dcl.IsEmptyValueIndirect(v) {
		m["locality"] = v
	}
	if v := f.Province; !dcl.IsEmptyValueIndirect(v) {
		m["province"] = v
	}
	if v := f.StreetAddress; !dcl.IsEmptyValueIndirect(v) {
		m["streetAddress"] = v
	}
	if v := f.PostalCode; !dcl.IsEmptyValueIndirect(v) {
		m["postalCode"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject flattens an instance of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject
	}
	r.CommonName = dcl.FlattenString(m["commonName"])
	r.CountryCode = dcl.FlattenString(m["countryCode"])
	r.Organization = dcl.FlattenString(m["organization"])
	r.OrganizationalUnit = dcl.FlattenString(m["organizationalUnit"])
	r.Locality = dcl.FlattenString(m["locality"])
	r.Province = dcl.FlattenString(m["province"])
	r.StreetAddress = dcl.FlattenString(m["streetAddress"])
	r.PostalCode = dcl.FlattenString(m["postalCode"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameMap expands the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName expands an instance of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(c *Client, f *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DnsNames; !dcl.IsEmptyValueIndirect(v) {
		m["dnsNames"] = v
	}
	if v := f.Uris; !dcl.IsEmptyValueIndirect(v) {
		m["uris"] = v
	}
	if v := f.EmailAddresses; !dcl.IsEmptyValueIndirect(v) {
		m["emailAddresses"] = v
	}
	if v := f.IPAddresses; !dcl.IsEmptyValueIndirect(v) {
		m["ipAddresses"] = v
	}
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansSlice(c, f.CustomSans); err != nil {
		return nil, fmt.Errorf("error expanding CustomSans into customSans: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["customSans"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName flattens an instance of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName
	}
	r.DnsNames = dcl.FlattenStringSlice(m["dnsNames"])
	r.Uris = dcl.FlattenStringSlice(m["uris"])
	r.EmailAddresses = dcl.FlattenStringSlice(m["emailAddresses"])
	r.IPAddresses = dcl.FlattenStringSlice(m["ipAddresses"])
	r.CustomSans = flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansSlice(c, m["customSans"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansMap expands the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans expands an instance of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(c *Client, f *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(c, f.ObjectId); err != nil {
		return nil, fmt.Errorf("error expanding ObjectId into objectId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["objectId"] = v
	}
	if v := f.Critical; !dcl.IsEmptyValueIndirect(v) {
		m["critical"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans flattens an instance of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans
	}
	r.ObjectId = flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(c, m["objectId"])
	r.Critical = dcl.FlattenBool(m["critical"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdMap expands the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId expands an instance of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(c *Client, f *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ObjectIdPath; !dcl.IsEmptyValueIndirect(v) {
		m["objectIdPath"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId flattens an instance of CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId
	}
	r.ObjectIdPath = dcl.FlattenIntSlice(m["objectIdPath"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsPublicKeyMap expands the contents of CertificateAuthorityCaCertificateDescriptionsPublicKey into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsPublicKeyMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsPublicKey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsPublicKey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsPublicKeySlice expands the contents of CertificateAuthorityCaCertificateDescriptionsPublicKey into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsPublicKeySlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsPublicKey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsPublicKey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsPublicKeyMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsPublicKey from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsPublicKeyMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsPublicKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsPublicKey{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsPublicKey{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsPublicKey)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsPublicKey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsPublicKeySlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsPublicKey from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsPublicKeySlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsPublicKey {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsPublicKey{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsPublicKey{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsPublicKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsPublicKey(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsPublicKey expands an instance of CertificateAuthorityCaCertificateDescriptionsPublicKey into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsPublicKey(c *Client, f *CertificateAuthorityCaCertificateDescriptionsPublicKey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Key; !dcl.IsEmptyValueIndirect(v) {
		m["key"] = v
	}
	if v := f.Format; !dcl.IsEmptyValueIndirect(v) {
		m["format"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsPublicKey flattens an instance of CertificateAuthorityCaCertificateDescriptionsPublicKey from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsPublicKey(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsPublicKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsPublicKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsPublicKey
	}
	r.Key = dcl.FlattenString(m["key"])
	r.Format = flattenCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(m["format"])
	r.Type = flattenCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum(m["type"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdMap expands the contents of CertificateAuthorityCaCertificateDescriptionsSubjectKeyId into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsSubjectKeyId into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsSubjectKeyId from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsSubjectKeyId{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsSubjectKeyId{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsSubjectKeyId)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsSubjectKeyId from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectKeyIdSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsSubjectKeyId{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsSubjectKeyId{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsSubjectKeyId, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsSubjectKeyId expands an instance of CertificateAuthorityCaCertificateDescriptionsSubjectKeyId into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(c *Client, f *CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.KeyId; !dcl.IsEmptyValueIndirect(v) {
		m["keyId"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsSubjectKeyId flattens an instance of CertificateAuthorityCaCertificateDescriptionsSubjectKeyId from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsSubjectKeyId(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsSubjectKeyId {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsSubjectKeyId{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsSubjectKeyId
	}
	r.KeyId = dcl.FlattenString(m["keyId"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdMap expands the contents of CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId expands an instance of CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(c *Client, f *CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.KeyId; !dcl.IsEmptyValueIndirect(v) {
		m["keyId"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId flattens an instance of CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId
	}
	r.KeyId = dcl.FlattenString(m["keyId"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsCertFingerprintMap expands the contents of CertificateAuthorityCaCertificateDescriptionsCertFingerprint into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsCertFingerprintMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsCertFingerprint) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsCertFingerprint(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsCertFingerprintSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsCertFingerprint into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsCertFingerprintSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsCertFingerprint) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsCertFingerprint(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsCertFingerprintMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsCertFingerprint from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsCertFingerprintMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsCertFingerprint{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsCertFingerprint{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsCertFingerprint)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsCertFingerprint(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsCertFingerprintSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsCertFingerprint from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsCertFingerprintSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsCertFingerprint{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsCertFingerprint{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsCertFingerprint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsCertFingerprint(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsCertFingerprint expands an instance of CertificateAuthorityCaCertificateDescriptionsCertFingerprint into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsCertFingerprint(c *Client, f *CertificateAuthorityCaCertificateDescriptionsCertFingerprint) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Sha256Hash; !dcl.IsEmptyValueIndirect(v) {
		m["sha256Hash"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsCertFingerprint flattens an instance of CertificateAuthorityCaCertificateDescriptionsCertFingerprint from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsCertFingerprint(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsCertFingerprint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsCertFingerprint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsCertFingerprint
	}
	r.Sha256Hash = dcl.FlattenString(m["sha256Hash"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesMap expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValues into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsConfigValues) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValues(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValues into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsConfigValues) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValues(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsConfigValues {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValues{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValues{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsConfigValues)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsConfigValues(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsConfigValues {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValues{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValues{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsConfigValues, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsConfigValues(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValues expands an instance of CertificateAuthorityCaCertificateDescriptionsConfigValues into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValues(c *Client, f *CertificateAuthorityCaCertificateDescriptionsConfigValues) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(c, f.KeyUsage); err != nil {
		return nil, fmt.Errorf("error expanding KeyUsage into keyUsage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["keyUsage"] = v
	}
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(c, f.CaOptions); err != nil {
		return nil, fmt.Errorf("error expanding CaOptions into caOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["caOptions"] = v
	}
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsSlice(c, f.PolicyIds); err != nil {
		return nil, fmt.Errorf("error expanding PolicyIds into policyIds: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["policyIds"] = v
	}
	if v := f.AiaOcspServers; !dcl.IsEmptyValueIndirect(v) {
		m["aiaOcspServers"] = v
	}
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsSlice(c, f.AdditionalExtensions); err != nil {
		return nil, fmt.Errorf("error expanding AdditionalExtensions into additionalExtensions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["additionalExtensions"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValues flattens an instance of CertificateAuthorityCaCertificateDescriptionsConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValues(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsConfigValues {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsConfigValues{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsConfigValues
	}
	r.KeyUsage = flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(c, m["keyUsage"])
	r.CaOptions = flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(c, m["caOptions"])
	r.PolicyIds = flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsSlice(c, m["policyIds"])
	r.AiaOcspServers = dcl.FlattenStringSlice(m["aiaOcspServers"])
	r.AdditionalExtensions = flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsSlice(c, m["additionalExtensions"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageMap expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage expands an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(c *Client, f *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(c, f.BaseKeyUsage); err != nil {
		return nil, fmt.Errorf("error expanding BaseKeyUsage into baseKeyUsage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["baseKeyUsage"] = v
	}
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(c, f.ExtendedKeyUsage); err != nil {
		return nil, fmt.Errorf("error expanding ExtendedKeyUsage into extendedKeyUsage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["extendedKeyUsage"] = v
	}
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c, f.UnknownExtendedKeyUsages); err != nil {
		return nil, fmt.Errorf("error expanding UnknownExtendedKeyUsages into unknownExtendedKeyUsages: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["unknownExtendedKeyUsages"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage flattens an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage
	}
	r.BaseKeyUsage = flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(c, m["baseKeyUsage"])
	r.ExtendedKeyUsage = flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(c, m["extendedKeyUsage"])
	r.UnknownExtendedKeyUsages = flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c, m["unknownExtendedKeyUsages"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageMap expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage expands an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(c *Client, f *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DigitalSignature; !dcl.IsEmptyValueIndirect(v) {
		m["digitalSignature"] = v
	}
	if v := f.ContentCommitment; !dcl.IsEmptyValueIndirect(v) {
		m["contentCommitment"] = v
	}
	if v := f.KeyEncipherment; !dcl.IsEmptyValueIndirect(v) {
		m["keyEncipherment"] = v
	}
	if v := f.DataEncipherment; !dcl.IsEmptyValueIndirect(v) {
		m["dataEncipherment"] = v
	}
	if v := f.KeyAgreement; !dcl.IsEmptyValueIndirect(v) {
		m["keyAgreement"] = v
	}
	if v := f.CertSign; !dcl.IsEmptyValueIndirect(v) {
		m["certSign"] = v
	}
	if v := f.CrlSign; !dcl.IsEmptyValueIndirect(v) {
		m["crlSign"] = v
	}
	if v := f.EncipherOnly; !dcl.IsEmptyValueIndirect(v) {
		m["encipherOnly"] = v
	}
	if v := f.DecipherOnly; !dcl.IsEmptyValueIndirect(v) {
		m["decipherOnly"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage flattens an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage
	}
	r.DigitalSignature = dcl.FlattenBool(m["digitalSignature"])
	r.ContentCommitment = dcl.FlattenBool(m["contentCommitment"])
	r.KeyEncipherment = dcl.FlattenBool(m["keyEncipherment"])
	r.DataEncipherment = dcl.FlattenBool(m["dataEncipherment"])
	r.KeyAgreement = dcl.FlattenBool(m["keyAgreement"])
	r.CertSign = dcl.FlattenBool(m["certSign"])
	r.CrlSign = dcl.FlattenBool(m["crlSign"])
	r.EncipherOnly = dcl.FlattenBool(m["encipherOnly"])
	r.DecipherOnly = dcl.FlattenBool(m["decipherOnly"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageMap expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage expands an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(c *Client, f *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ServerAuth; !dcl.IsEmptyValueIndirect(v) {
		m["serverAuth"] = v
	}
	if v := f.ClientAuth; !dcl.IsEmptyValueIndirect(v) {
		m["clientAuth"] = v
	}
	if v := f.CodeSigning; !dcl.IsEmptyValueIndirect(v) {
		m["codeSigning"] = v
	}
	if v := f.EmailProtection; !dcl.IsEmptyValueIndirect(v) {
		m["emailProtection"] = v
	}
	if v := f.TimeStamping; !dcl.IsEmptyValueIndirect(v) {
		m["timeStamping"] = v
	}
	if v := f.OcspSigning; !dcl.IsEmptyValueIndirect(v) {
		m["ocspSigning"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage flattens an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage
	}
	r.ServerAuth = dcl.FlattenBool(m["serverAuth"])
	r.ClientAuth = dcl.FlattenBool(m["clientAuth"])
	r.CodeSigning = dcl.FlattenBool(m["codeSigning"])
	r.EmailProtection = dcl.FlattenBool(m["emailProtection"])
	r.TimeStamping = dcl.FlattenBool(m["timeStamping"])
	r.OcspSigning = dcl.FlattenBool(m["ocspSigning"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesMap expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages expands an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages(c *Client, f *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ObjectIdPath; !dcl.IsEmptyValueIndirect(v) {
		m["objectIdPath"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages flattens an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages
	}
	r.ObjectIdPath = dcl.FlattenIntSlice(m["objectIdPath"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsMap expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions expands an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(c *Client, f *CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IsCa; !dcl.IsEmptyValueIndirect(v) {
		m["isCa"] = v
	}
	if v := f.MaxIssuerPathLength; !dcl.IsEmptyValueIndirect(v) {
		m["maxIssuerPathLength"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions flattens an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions
	}
	r.IsCa = dcl.FlattenBool(m["isCa"])
	r.MaxIssuerPathLength = dcl.FlattenInteger(m["maxIssuerPathLength"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsMap expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds expands an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds(c *Client, f *CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ObjectIdPath; !dcl.IsEmptyValueIndirect(v) {
		m["objectIdPath"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds flattens an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds
	}
	r.ObjectIdPath = dcl.FlattenIntSlice(m["objectIdPath"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsMap expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions expands an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions(c *Client, f *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(c, f.ObjectId); err != nil {
		return nil, fmt.Errorf("error expanding ObjectId into objectId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["objectId"] = v
	}
	if v := f.Critical; !dcl.IsEmptyValueIndirect(v) {
		m["critical"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions flattens an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions
	}
	r.ObjectId = flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(c, m["objectId"])
	r.Critical = dcl.FlattenBool(m["critical"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdMap expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdMap(c *Client, f map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdSlice expands the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdSlice(c *Client, f []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdMap flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdMap(c *Client, i interface{}) map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId{}
	}

	items := make(map[string]CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId expands an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId into a JSON
// request object.
func expandCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(c *Client, f *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ObjectIdPath; !dcl.IsEmptyValueIndirect(v) {
		m["objectIdPath"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId flattens an instance of CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(c *Client, i interface{}) *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId
	}
	r.ObjectIdPath = dcl.FlattenIntSlice(m["objectIdPath"])

	return r
}

// expandCertificateAuthorityAccessUrlsMap expands the contents of CertificateAuthorityAccessUrls into a JSON
// request object.
func expandCertificateAuthorityAccessUrlsMap(c *Client, f map[string]CertificateAuthorityAccessUrls) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityAccessUrls(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityAccessUrlsSlice expands the contents of CertificateAuthorityAccessUrls into a JSON
// request object.
func expandCertificateAuthorityAccessUrlsSlice(c *Client, f []CertificateAuthorityAccessUrls) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityAccessUrls(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityAccessUrlsMap flattens the contents of CertificateAuthorityAccessUrls from a JSON
// response object.
func flattenCertificateAuthorityAccessUrlsMap(c *Client, i interface{}) map[string]CertificateAuthorityAccessUrls {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityAccessUrls{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityAccessUrls{}
	}

	items := make(map[string]CertificateAuthorityAccessUrls)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityAccessUrls(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityAccessUrlsSlice flattens the contents of CertificateAuthorityAccessUrls from a JSON
// response object.
func flattenCertificateAuthorityAccessUrlsSlice(c *Client, i interface{}) []CertificateAuthorityAccessUrls {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityAccessUrls{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityAccessUrls{}
	}

	items := make([]CertificateAuthorityAccessUrls, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityAccessUrls(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityAccessUrls expands an instance of CertificateAuthorityAccessUrls into a JSON
// request object.
func expandCertificateAuthorityAccessUrls(c *Client, f *CertificateAuthorityAccessUrls) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CaCertificateAccessUrl; !dcl.IsEmptyValueIndirect(v) {
		m["caCertificateAccessUrl"] = v
	}
	if v := f.CrlAccessUrls; !dcl.IsEmptyValueIndirect(v) {
		m["crlAccessUrls"] = v
	}
	if v := f.CrlAccessUrl; !dcl.IsEmptyValueIndirect(v) {
		m["crlAccessUrl"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityAccessUrls flattens an instance of CertificateAuthorityAccessUrls from a JSON
// response object.
func flattenCertificateAuthorityAccessUrls(c *Client, i interface{}) *CertificateAuthorityAccessUrls {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityAccessUrls{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityAccessUrls
	}
	r.CaCertificateAccessUrl = dcl.FlattenString(m["caCertificateAccessUrl"])
	r.CrlAccessUrls = dcl.FlattenStringSlice(m["crlAccessUrls"])
	r.CrlAccessUrl = dcl.FlattenString(m["crlAccessUrl"])

	return r
}

// expandCertificateAuthorityCertificatePolicyMap expands the contents of CertificateAuthorityCertificatePolicy into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyMap(c *Client, f map[string]CertificateAuthorityCertificatePolicy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicy(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicySlice expands the contents of CertificateAuthorityCertificatePolicy into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicySlice(c *Client, f []CertificateAuthorityCertificatePolicy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicy(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyMap flattens the contents of CertificateAuthorityCertificatePolicy from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicy {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicy{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicy{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicy)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicy(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicySlice flattens the contents of CertificateAuthorityCertificatePolicy from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicySlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicy {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicy{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicy{}
	}

	items := make([]CertificateAuthorityCertificatePolicy, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicy(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicy expands an instance of CertificateAuthorityCertificatePolicy into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicy(c *Client, f *CertificateAuthorityCertificatePolicy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityCertificatePolicyAllowedConfigList(c, f.AllowedConfigList); err != nil {
		return nil, fmt.Errorf("error expanding AllowedConfigList into allowedConfigList: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["allowedConfigList"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValues(c, f.OverwriteConfigValues); err != nil {
		return nil, fmt.Errorf("error expanding OverwriteConfigValues into overwriteConfigValues: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["overwriteConfigValues"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsSlice(c, f.AllowedLocationsAndOrganizations); err != nil {
		return nil, fmt.Errorf("error expanding AllowedLocationsAndOrganizations into allowedLocationsAndOrganizations: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["allowedLocationsAndOrganizations"] = v
	}
	if v := f.AllowedCommonNames; !dcl.IsEmptyValueIndirect(v) {
		m["allowedCommonNames"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyAllowedSans(c, f.AllowedSans); err != nil {
		return nil, fmt.Errorf("error expanding AllowedSans into allowedSans: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["allowedSans"] = v
	}
	if v := f.MaximumLifetime; !dcl.IsEmptyValueIndirect(v) {
		m["maximumLifetime"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyAllowedIssuanceModes(c, f.AllowedIssuanceModes); err != nil {
		return nil, fmt.Errorf("error expanding AllowedIssuanceModes into allowedIssuanceModes: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["allowedIssuanceModes"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicy flattens an instance of CertificateAuthorityCertificatePolicy from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicy(c *Client, i interface{}) *CertificateAuthorityCertificatePolicy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicy{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicy
	}
	r.AllowedConfigList = flattenCertificateAuthorityCertificatePolicyAllowedConfigList(c, m["allowedConfigList"])
	r.OverwriteConfigValues = flattenCertificateAuthorityCertificatePolicyOverwriteConfigValues(c, m["overwriteConfigValues"])
	r.AllowedLocationsAndOrganizations = flattenCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsSlice(c, m["allowedLocationsAndOrganizations"])
	r.AllowedCommonNames = dcl.FlattenStringSlice(m["allowedCommonNames"])
	r.AllowedSans = flattenCertificateAuthorityCertificatePolicyAllowedSans(c, m["allowedSans"])
	r.MaximumLifetime = dcl.FlattenString(m["maximumLifetime"])
	r.AllowedIssuanceModes = flattenCertificateAuthorityCertificatePolicyAllowedIssuanceModes(c, m["allowedIssuanceModes"])

	return r
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListMap expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigList into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyAllowedConfigList) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigList(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListSlice expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigList into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListSlice(c *Client, f []CertificateAuthorityCertificatePolicyAllowedConfigList) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigList(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListMap flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigList from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyAllowedConfigList {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigList{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigList{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyAllowedConfigList)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyAllowedConfigList(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListSlice flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigList from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyAllowedConfigList {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyAllowedConfigList{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyAllowedConfigList{}
	}

	items := make([]CertificateAuthorityCertificatePolicyAllowedConfigList, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyAllowedConfigList(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigList expands an instance of CertificateAuthorityCertificatePolicyAllowedConfigList into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigList(c *Client, f *CertificateAuthorityCertificatePolicyAllowedConfigList) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesSlice(c, f.AllowedConfigValues); err != nil {
		return nil, fmt.Errorf("error expanding AllowedConfigValues into allowedConfigValues: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["allowedConfigValues"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigList flattens an instance of CertificateAuthorityCertificatePolicyAllowedConfigList from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigList(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyAllowedConfigList {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyAllowedConfigList{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyAllowedConfigList
	}
	r.AllowedConfigValues = flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesSlice(c, m["allowedConfigValues"])

	return r
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesMap expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesSlice expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesSlice(c *Client, f []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesMap flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesSlice flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues{}
	}

	items := make([]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues expands an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues(c *Client, f *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ReusableConfig; !dcl.IsEmptyValueIndirect(v) {
		m["reusableConfig"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(c, f.ReusableConfigValues); err != nil {
		return nil, fmt.Errorf("error expanding ReusableConfigValues into reusableConfigValues: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reusableConfigValues"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues flattens an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues
	}
	r.ReusableConfig = dcl.FlattenString(m["reusableConfig"])
	r.ReusableConfigValues = flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(c, m["reusableConfigValues"])

	return r
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesMap expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesSlice expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesSlice(c *Client, f []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesMap flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesSlice flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues{}
	}

	items := make([]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues expands an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(c *Client, f *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(c, f.KeyUsage); err != nil {
		return nil, fmt.Errorf("error expanding KeyUsage into keyUsage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["keyUsage"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(c, f.CaOptions); err != nil {
		return nil, fmt.Errorf("error expanding CaOptions into caOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["caOptions"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsSlice(c, f.PolicyIds); err != nil {
		return nil, fmt.Errorf("error expanding PolicyIds into policyIds: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["policyIds"] = v
	}
	if v := f.AiaOcspServers; !dcl.IsEmptyValueIndirect(v) {
		m["aiaOcspServers"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsSlice(c, f.AdditionalExtensions); err != nil {
		return nil, fmt.Errorf("error expanding AdditionalExtensions into additionalExtensions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["additionalExtensions"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues flattens an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues
	}
	r.KeyUsage = flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(c, m["keyUsage"])
	r.CaOptions = flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(c, m["caOptions"])
	r.PolicyIds = flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsSlice(c, m["policyIds"])
	r.AiaOcspServers = dcl.FlattenStringSlice(m["aiaOcspServers"])
	r.AdditionalExtensions = flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsSlice(c, m["additionalExtensions"])

	return r
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageMap expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageSlice expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageSlice(c *Client, f []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageMap flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageSlice flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage{}
	}

	items := make([]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage expands an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(c *Client, f *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, f.BaseKeyUsage); err != nil {
		return nil, fmt.Errorf("error expanding BaseKeyUsage into baseKeyUsage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["baseKeyUsage"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, f.ExtendedKeyUsage); err != nil {
		return nil, fmt.Errorf("error expanding ExtendedKeyUsage into extendedKeyUsage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["extendedKeyUsage"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c, f.UnknownExtendedKeyUsages); err != nil {
		return nil, fmt.Errorf("error expanding UnknownExtendedKeyUsages into unknownExtendedKeyUsages: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["unknownExtendedKeyUsages"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage flattens an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage
	}
	r.BaseKeyUsage = flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, m["baseKeyUsage"])
	r.ExtendedKeyUsage = flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, m["extendedKeyUsage"])
	r.UnknownExtendedKeyUsages = flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c, m["unknownExtendedKeyUsages"])

	return r
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageMap expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageSlice expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageSlice(c *Client, f []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageMap flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageSlice flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{}
	}

	items := make([]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage expands an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c *Client, f *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DigitalSignature; !dcl.IsEmptyValueIndirect(v) {
		m["digitalSignature"] = v
	}
	if v := f.ContentCommitment; !dcl.IsEmptyValueIndirect(v) {
		m["contentCommitment"] = v
	}
	if v := f.KeyEncipherment; !dcl.IsEmptyValueIndirect(v) {
		m["keyEncipherment"] = v
	}
	if v := f.DataEncipherment; !dcl.IsEmptyValueIndirect(v) {
		m["dataEncipherment"] = v
	}
	if v := f.KeyAgreement; !dcl.IsEmptyValueIndirect(v) {
		m["keyAgreement"] = v
	}
	if v := f.CertSign; !dcl.IsEmptyValueIndirect(v) {
		m["certSign"] = v
	}
	if v := f.CrlSign; !dcl.IsEmptyValueIndirect(v) {
		m["crlSign"] = v
	}
	if v := f.EncipherOnly; !dcl.IsEmptyValueIndirect(v) {
		m["encipherOnly"] = v
	}
	if v := f.DecipherOnly; !dcl.IsEmptyValueIndirect(v) {
		m["decipherOnly"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage flattens an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage
	}
	r.DigitalSignature = dcl.FlattenBool(m["digitalSignature"])
	r.ContentCommitment = dcl.FlattenBool(m["contentCommitment"])
	r.KeyEncipherment = dcl.FlattenBool(m["keyEncipherment"])
	r.DataEncipherment = dcl.FlattenBool(m["dataEncipherment"])
	r.KeyAgreement = dcl.FlattenBool(m["keyAgreement"])
	r.CertSign = dcl.FlattenBool(m["certSign"])
	r.CrlSign = dcl.FlattenBool(m["crlSign"])
	r.EncipherOnly = dcl.FlattenBool(m["encipherOnly"])
	r.DecipherOnly = dcl.FlattenBool(m["decipherOnly"])

	return r
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageMap expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageSlice expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageSlice(c *Client, f []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageMap flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageSlice flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{}
	}

	items := make([]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage expands an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c *Client, f *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ServerAuth; !dcl.IsEmptyValueIndirect(v) {
		m["serverAuth"] = v
	}
	if v := f.ClientAuth; !dcl.IsEmptyValueIndirect(v) {
		m["clientAuth"] = v
	}
	if v := f.CodeSigning; !dcl.IsEmptyValueIndirect(v) {
		m["codeSigning"] = v
	}
	if v := f.EmailProtection; !dcl.IsEmptyValueIndirect(v) {
		m["emailProtection"] = v
	}
	if v := f.TimeStamping; !dcl.IsEmptyValueIndirect(v) {
		m["timeStamping"] = v
	}
	if v := f.OcspSigning; !dcl.IsEmptyValueIndirect(v) {
		m["ocspSigning"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage flattens an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage
	}
	r.ServerAuth = dcl.FlattenBool(m["serverAuth"])
	r.ClientAuth = dcl.FlattenBool(m["clientAuth"])
	r.CodeSigning = dcl.FlattenBool(m["codeSigning"])
	r.EmailProtection = dcl.FlattenBool(m["emailProtection"])
	r.TimeStamping = dcl.FlattenBool(m["timeStamping"])
	r.OcspSigning = dcl.FlattenBool(m["ocspSigning"])

	return r
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesMap expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c *Client, f []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesMap flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	items := make([]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages expands an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c *Client, f *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ObjectIdPath; !dcl.IsEmptyValueIndirect(v) {
		m["objectIdPath"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages flattens an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	}
	r.ObjectIdPath = dcl.FlattenIntSlice(m["objectIdPath"])

	return r
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsMap expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsSlice expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsSlice(c *Client, f []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsMap flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsSlice flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions{}
	}

	items := make([]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions expands an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(c *Client, f *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IsCa; !dcl.IsEmptyValueIndirect(v) {
		m["isCa"] = v
	}
	if v := f.MaxIssuerPathLength; !dcl.IsEmptyValueIndirect(v) {
		m["maxIssuerPathLength"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions flattens an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions
	}
	r.IsCa = dcl.FlattenBool(m["isCa"])
	r.MaxIssuerPathLength = dcl.FlattenInteger(m["maxIssuerPathLength"])

	return r
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsMap expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsSlice expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsSlice(c *Client, f []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsMap flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsSlice flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds{}
	}

	items := make([]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds expands an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds(c *Client, f *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ObjectIdPath; !dcl.IsEmptyValueIndirect(v) {
		m["objectIdPath"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds flattens an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds
	}
	r.ObjectIdPath = dcl.FlattenIntSlice(m["objectIdPath"])

	return r
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsMap expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsSlice expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsSlice(c *Client, f []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsMap flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsSlice flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions{}
	}

	items := make([]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions expands an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions(c *Client, f *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, f.ObjectId); err != nil {
		return nil, fmt.Errorf("error expanding ObjectId into objectId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["objectId"] = v
	}
	if v := f.Critical; !dcl.IsEmptyValueIndirect(v) {
		m["critical"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions flattens an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions
	}
	r.ObjectId = flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, m["objectId"])
	r.Critical = dcl.FlattenBool(m["critical"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdMap expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdSlice expands the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdSlice(c *Client, f []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdMap flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdSlice flattens the contents of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{}
	}

	items := make([]CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId expands an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c *Client, f *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ObjectIdPath; !dcl.IsEmptyValueIndirect(v) {
		m["objectIdPath"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId flattens an instance of CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId
	}
	r.ObjectIdPath = dcl.FlattenIntSlice(m["objectIdPath"])

	return r
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesMap expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValues into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValues) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValues(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesSlice expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValues into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesSlice(c *Client, f []CertificateAuthorityCertificatePolicyOverwriteConfigValues) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValues(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesMap flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValues {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValues{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValues{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValues)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValues(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesSlice flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyOverwriteConfigValues {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValues{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValues{}
	}

	items := make([]CertificateAuthorityCertificatePolicyOverwriteConfigValues, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValues(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValues expands an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValues into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValues(c *Client, f *CertificateAuthorityCertificatePolicyOverwriteConfigValues) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ReusableConfig; !dcl.IsEmptyValueIndirect(v) {
		m["reusableConfig"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(c, f.ReusableConfigValues); err != nil {
		return nil, fmt.Errorf("error expanding ReusableConfigValues into reusableConfigValues: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["reusableConfigValues"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValues flattens an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValues(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyOverwriteConfigValues {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyOverwriteConfigValues{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValues
	}
	r.ReusableConfig = dcl.FlattenString(m["reusableConfig"])
	r.ReusableConfigValues = flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(c, m["reusableConfigValues"])

	return r
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesMap expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesSlice expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesSlice(c *Client, f []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesMap flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesSlice flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues{}
	}

	items := make([]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues expands an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(c *Client, f *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(c, f.KeyUsage); err != nil {
		return nil, fmt.Errorf("error expanding KeyUsage into keyUsage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["keyUsage"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(c, f.CaOptions); err != nil {
		return nil, fmt.Errorf("error expanding CaOptions into caOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["caOptions"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsSlice(c, f.PolicyIds); err != nil {
		return nil, fmt.Errorf("error expanding PolicyIds into policyIds: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["policyIds"] = v
	}
	if v := f.AiaOcspServers; !dcl.IsEmptyValueIndirect(v) {
		m["aiaOcspServers"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsSlice(c, f.AdditionalExtensions); err != nil {
		return nil, fmt.Errorf("error expanding AdditionalExtensions into additionalExtensions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["additionalExtensions"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues flattens an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues
	}
	r.KeyUsage = flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(c, m["keyUsage"])
	r.CaOptions = flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(c, m["caOptions"])
	r.PolicyIds = flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsSlice(c, m["policyIds"])
	r.AiaOcspServers = dcl.FlattenStringSlice(m["aiaOcspServers"])
	r.AdditionalExtensions = flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsSlice(c, m["additionalExtensions"])

	return r
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageMap expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageSlice expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageSlice(c *Client, f []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageMap flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageSlice flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage{}
	}

	items := make([]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage expands an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(c *Client, f *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, f.BaseKeyUsage); err != nil {
		return nil, fmt.Errorf("error expanding BaseKeyUsage into baseKeyUsage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["baseKeyUsage"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, f.ExtendedKeyUsage); err != nil {
		return nil, fmt.Errorf("error expanding ExtendedKeyUsage into extendedKeyUsage: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["extendedKeyUsage"] = v
	}
	if v, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c, f.UnknownExtendedKeyUsages); err != nil {
		return nil, fmt.Errorf("error expanding UnknownExtendedKeyUsages into unknownExtendedKeyUsages: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["unknownExtendedKeyUsages"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage flattens an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage
	}
	r.BaseKeyUsage = flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, m["baseKeyUsage"])
	r.ExtendedKeyUsage = flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, m["extendedKeyUsage"])
	r.UnknownExtendedKeyUsages = flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c, m["unknownExtendedKeyUsages"])

	return r
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageMap expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageSlice expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageSlice(c *Client, f []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageMap flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageSlice flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{}
	}

	items := make([]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage expands an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c *Client, f *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DigitalSignature; !dcl.IsEmptyValueIndirect(v) {
		m["digitalSignature"] = v
	}
	if v := f.ContentCommitment; !dcl.IsEmptyValueIndirect(v) {
		m["contentCommitment"] = v
	}
	if v := f.KeyEncipherment; !dcl.IsEmptyValueIndirect(v) {
		m["keyEncipherment"] = v
	}
	if v := f.DataEncipherment; !dcl.IsEmptyValueIndirect(v) {
		m["dataEncipherment"] = v
	}
	if v := f.KeyAgreement; !dcl.IsEmptyValueIndirect(v) {
		m["keyAgreement"] = v
	}
	if v := f.CertSign; !dcl.IsEmptyValueIndirect(v) {
		m["certSign"] = v
	}
	if v := f.CrlSign; !dcl.IsEmptyValueIndirect(v) {
		m["crlSign"] = v
	}
	if v := f.EncipherOnly; !dcl.IsEmptyValueIndirect(v) {
		m["encipherOnly"] = v
	}
	if v := f.DecipherOnly; !dcl.IsEmptyValueIndirect(v) {
		m["decipherOnly"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage flattens an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage
	}
	r.DigitalSignature = dcl.FlattenBool(m["digitalSignature"])
	r.ContentCommitment = dcl.FlattenBool(m["contentCommitment"])
	r.KeyEncipherment = dcl.FlattenBool(m["keyEncipherment"])
	r.DataEncipherment = dcl.FlattenBool(m["dataEncipherment"])
	r.KeyAgreement = dcl.FlattenBool(m["keyAgreement"])
	r.CertSign = dcl.FlattenBool(m["certSign"])
	r.CrlSign = dcl.FlattenBool(m["crlSign"])
	r.EncipherOnly = dcl.FlattenBool(m["encipherOnly"])
	r.DecipherOnly = dcl.FlattenBool(m["decipherOnly"])

	return r
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageMap expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageSlice expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageSlice(c *Client, f []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageMap flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageSlice flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{}
	}

	items := make([]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage expands an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c *Client, f *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ServerAuth; !dcl.IsEmptyValueIndirect(v) {
		m["serverAuth"] = v
	}
	if v := f.ClientAuth; !dcl.IsEmptyValueIndirect(v) {
		m["clientAuth"] = v
	}
	if v := f.CodeSigning; !dcl.IsEmptyValueIndirect(v) {
		m["codeSigning"] = v
	}
	if v := f.EmailProtection; !dcl.IsEmptyValueIndirect(v) {
		m["emailProtection"] = v
	}
	if v := f.TimeStamping; !dcl.IsEmptyValueIndirect(v) {
		m["timeStamping"] = v
	}
	if v := f.OcspSigning; !dcl.IsEmptyValueIndirect(v) {
		m["ocspSigning"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage flattens an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage
	}
	r.ServerAuth = dcl.FlattenBool(m["serverAuth"])
	r.ClientAuth = dcl.FlattenBool(m["clientAuth"])
	r.CodeSigning = dcl.FlattenBool(m["codeSigning"])
	r.EmailProtection = dcl.FlattenBool(m["emailProtection"])
	r.TimeStamping = dcl.FlattenBool(m["timeStamping"])
	r.OcspSigning = dcl.FlattenBool(m["ocspSigning"])

	return r
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesMap expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c *Client, f []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesMap flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}
	}

	items := make([]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages expands an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c *Client, f *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ObjectIdPath; !dcl.IsEmptyValueIndirect(v) {
		m["objectIdPath"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages flattens an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	}
	r.ObjectIdPath = dcl.FlattenIntSlice(m["objectIdPath"])

	return r
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsMap expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsSlice expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsSlice(c *Client, f []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsMap flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsSlice flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions{}
	}

	items := make([]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions expands an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(c *Client, f *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IsCa; !dcl.IsEmptyValueIndirect(v) {
		m["isCa"] = v
	}
	if v := f.MaxIssuerPathLength; !dcl.IsEmptyValueIndirect(v) {
		m["maxIssuerPathLength"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions flattens an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions
	}
	r.IsCa = dcl.FlattenBool(m["isCa"])
	r.MaxIssuerPathLength = dcl.FlattenInteger(m["maxIssuerPathLength"])

	return r
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsMap expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsSlice expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsSlice(c *Client, f []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsMap flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsSlice flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds{}
	}

	items := make([]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds expands an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds(c *Client, f *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ObjectIdPath; !dcl.IsEmptyValueIndirect(v) {
		m["objectIdPath"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds flattens an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds
	}
	r.ObjectIdPath = dcl.FlattenIntSlice(m["objectIdPath"])

	return r
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsMap expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsSlice expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsSlice(c *Client, f []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsMap flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsSlice flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions{}
	}

	items := make([]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions expands an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions(c *Client, f *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, f.ObjectId); err != nil {
		return nil, fmt.Errorf("error expanding ObjectId into objectId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["objectId"] = v
	}
	if v := f.Critical; !dcl.IsEmptyValueIndirect(v) {
		m["critical"] = v
	}
	if v := f.Value; !dcl.IsEmptyValueIndirect(v) {
		m["value"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions flattens an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions
	}
	r.ObjectId = flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, m["objectId"])
	r.Critical = dcl.FlattenBool(m["critical"])
	r.Value = dcl.FlattenString(m["value"])

	return r
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdMap expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdSlice expands the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdSlice(c *Client, f []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdMap flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdSlice flattens the contents of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{}
	}

	items := make([]CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId expands an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c *Client, f *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ObjectIdPath; !dcl.IsEmptyValueIndirect(v) {
		m["objectIdPath"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId flattens an instance of CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId
	}
	r.ObjectIdPath = dcl.FlattenIntSlice(m["objectIdPath"])

	return r
}

// expandCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsMap expands the contents of CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsSlice expands the contents of CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsSlice(c *Client, f []CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsMap flattens the contents of CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsSlice flattens the contents of CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations{}
	}

	items := make([]CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations expands an instance of CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations(c *Client, f *CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CountryCode; !dcl.IsEmptyValueIndirect(v) {
		m["countryCode"] = v
	}
	if v := f.Organization; !dcl.IsEmptyValueIndirect(v) {
		m["organization"] = v
	}
	if v := f.OrganizationalUnit; !dcl.IsEmptyValueIndirect(v) {
		m["organizationalUnit"] = v
	}
	if v := f.Locality; !dcl.IsEmptyValueIndirect(v) {
		m["locality"] = v
	}
	if v := f.Province; !dcl.IsEmptyValueIndirect(v) {
		m["province"] = v
	}
	if v := f.StreetAddress; !dcl.IsEmptyValueIndirect(v) {
		m["streetAddress"] = v
	}
	if v := f.PostalCode; !dcl.IsEmptyValueIndirect(v) {
		m["postalCode"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations flattens an instance of CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations
	}
	r.CountryCode = dcl.FlattenString(m["countryCode"])
	r.Organization = dcl.FlattenString(m["organization"])
	r.OrganizationalUnit = dcl.FlattenString(m["organizationalUnit"])
	r.Locality = dcl.FlattenString(m["locality"])
	r.Province = dcl.FlattenString(m["province"])
	r.StreetAddress = dcl.FlattenString(m["streetAddress"])
	r.PostalCode = dcl.FlattenString(m["postalCode"])

	return r
}

// expandCertificateAuthorityCertificatePolicyAllowedSansMap expands the contents of CertificateAuthorityCertificatePolicyAllowedSans into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedSansMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyAllowedSans) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedSans(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyAllowedSansSlice expands the contents of CertificateAuthorityCertificatePolicyAllowedSans into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedSansSlice(c *Client, f []CertificateAuthorityCertificatePolicyAllowedSans) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedSans(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedSansMap flattens the contents of CertificateAuthorityCertificatePolicyAllowedSans from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedSansMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyAllowedSans {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyAllowedSans{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyAllowedSans{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyAllowedSans)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyAllowedSans(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyAllowedSansSlice flattens the contents of CertificateAuthorityCertificatePolicyAllowedSans from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedSansSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyAllowedSans {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyAllowedSans{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyAllowedSans{}
	}

	items := make([]CertificateAuthorityCertificatePolicyAllowedSans, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyAllowedSans(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyAllowedSans expands an instance of CertificateAuthorityCertificatePolicyAllowedSans into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedSans(c *Client, f *CertificateAuthorityCertificatePolicyAllowedSans) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AllowedDnsNames; !dcl.IsEmptyValueIndirect(v) {
		m["allowedDnsNames"] = v
	}
	if v := f.AllowedUris; !dcl.IsEmptyValueIndirect(v) {
		m["allowedUris"] = v
	}
	if v := f.AllowedEmailAddresses; !dcl.IsEmptyValueIndirect(v) {
		m["allowedEmailAddresses"] = v
	}
	if v := f.AllowedIps; !dcl.IsEmptyValueIndirect(v) {
		m["allowedIps"] = v
	}
	if v := f.AllowGlobbingDnsWildcards; !dcl.IsEmptyValueIndirect(v) {
		m["allowGlobbingDnsWildcards"] = v
	}
	if v := f.AllowCustomSans; !dcl.IsEmptyValueIndirect(v) {
		m["allowCustomSans"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedSans flattens an instance of CertificateAuthorityCertificatePolicyAllowedSans from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedSans(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyAllowedSans {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyAllowedSans{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyAllowedSans
	}
	r.AllowedDnsNames = dcl.FlattenStringSlice(m["allowedDnsNames"])
	r.AllowedUris = dcl.FlattenStringSlice(m["allowedUris"])
	r.AllowedEmailAddresses = dcl.FlattenStringSlice(m["allowedEmailAddresses"])
	r.AllowedIps = dcl.FlattenStringSlice(m["allowedIps"])
	r.AllowGlobbingDnsWildcards = dcl.FlattenBool(m["allowGlobbingDnsWildcards"])
	r.AllowCustomSans = dcl.FlattenBool(m["allowCustomSans"])

	return r
}

// expandCertificateAuthorityCertificatePolicyAllowedIssuanceModesMap expands the contents of CertificateAuthorityCertificatePolicyAllowedIssuanceModes into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedIssuanceModesMap(c *Client, f map[string]CertificateAuthorityCertificatePolicyAllowedIssuanceModes) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedIssuanceModes(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityCertificatePolicyAllowedIssuanceModesSlice expands the contents of CertificateAuthorityCertificatePolicyAllowedIssuanceModes into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedIssuanceModesSlice(c *Client, f []CertificateAuthorityCertificatePolicyAllowedIssuanceModes) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityCertificatePolicyAllowedIssuanceModes(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedIssuanceModesMap flattens the contents of CertificateAuthorityCertificatePolicyAllowedIssuanceModes from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedIssuanceModesMap(c *Client, i interface{}) map[string]CertificateAuthorityCertificatePolicyAllowedIssuanceModes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityCertificatePolicyAllowedIssuanceModes{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityCertificatePolicyAllowedIssuanceModes{}
	}

	items := make(map[string]CertificateAuthorityCertificatePolicyAllowedIssuanceModes)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityCertificatePolicyAllowedIssuanceModes(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityCertificatePolicyAllowedIssuanceModesSlice flattens the contents of CertificateAuthorityCertificatePolicyAllowedIssuanceModes from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedIssuanceModesSlice(c *Client, i interface{}) []CertificateAuthorityCertificatePolicyAllowedIssuanceModes {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCertificatePolicyAllowedIssuanceModes{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCertificatePolicyAllowedIssuanceModes{}
	}

	items := make([]CertificateAuthorityCertificatePolicyAllowedIssuanceModes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCertificatePolicyAllowedIssuanceModes(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityCertificatePolicyAllowedIssuanceModes expands an instance of CertificateAuthorityCertificatePolicyAllowedIssuanceModes into a JSON
// request object.
func expandCertificateAuthorityCertificatePolicyAllowedIssuanceModes(c *Client, f *CertificateAuthorityCertificatePolicyAllowedIssuanceModes) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AllowCsrBasedIssuance; !dcl.IsEmptyValueIndirect(v) {
		m["allowCsrBasedIssuance"] = v
	}
	if v := f.AllowConfigBasedIssuance; !dcl.IsEmptyValueIndirect(v) {
		m["allowConfigBasedIssuance"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityCertificatePolicyAllowedIssuanceModes flattens an instance of CertificateAuthorityCertificatePolicyAllowedIssuanceModes from a JSON
// response object.
func flattenCertificateAuthorityCertificatePolicyAllowedIssuanceModes(c *Client, i interface{}) *CertificateAuthorityCertificatePolicyAllowedIssuanceModes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityCertificatePolicyAllowedIssuanceModes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityCertificatePolicyAllowedIssuanceModes
	}
	r.AllowCsrBasedIssuance = dcl.FlattenBool(m["allowCsrBasedIssuance"])
	r.AllowConfigBasedIssuance = dcl.FlattenBool(m["allowConfigBasedIssuance"])

	return r
}

// expandCertificateAuthorityIssuingOptionsMap expands the contents of CertificateAuthorityIssuingOptions into a JSON
// request object.
func expandCertificateAuthorityIssuingOptionsMap(c *Client, f map[string]CertificateAuthorityIssuingOptions) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandCertificateAuthorityIssuingOptions(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandCertificateAuthorityIssuingOptionsSlice expands the contents of CertificateAuthorityIssuingOptions into a JSON
// request object.
func expandCertificateAuthorityIssuingOptionsSlice(c *Client, f []CertificateAuthorityIssuingOptions) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandCertificateAuthorityIssuingOptions(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenCertificateAuthorityIssuingOptionsMap flattens the contents of CertificateAuthorityIssuingOptions from a JSON
// response object.
func flattenCertificateAuthorityIssuingOptionsMap(c *Client, i interface{}) map[string]CertificateAuthorityIssuingOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]CertificateAuthorityIssuingOptions{}
	}

	if len(a) == 0 {
		return map[string]CertificateAuthorityIssuingOptions{}
	}

	items := make(map[string]CertificateAuthorityIssuingOptions)
	for k, item := range a {
		items[k] = *flattenCertificateAuthorityIssuingOptions(c, item.(map[string]interface{}))
	}

	return items
}

// flattenCertificateAuthorityIssuingOptionsSlice flattens the contents of CertificateAuthorityIssuingOptions from a JSON
// response object.
func flattenCertificateAuthorityIssuingOptionsSlice(c *Client, i interface{}) []CertificateAuthorityIssuingOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityIssuingOptions{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityIssuingOptions{}
	}

	items := make([]CertificateAuthorityIssuingOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityIssuingOptions(c, item.(map[string]interface{})))
	}

	return items
}

// expandCertificateAuthorityIssuingOptions expands an instance of CertificateAuthorityIssuingOptions into a JSON
// request object.
func expandCertificateAuthorityIssuingOptions(c *Client, f *CertificateAuthorityIssuingOptions) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.IncludeCaCertUrl; !dcl.IsEmptyValueIndirect(v) {
		m["includeCaCertUrl"] = v
	}
	if v := f.IncludeCrlAccessUrl; !dcl.IsEmptyValueIndirect(v) {
		m["includeCrlAccessUrl"] = v
	}

	return m, nil
}

// flattenCertificateAuthorityIssuingOptions flattens an instance of CertificateAuthorityIssuingOptions from a JSON
// response object.
func flattenCertificateAuthorityIssuingOptions(c *Client, i interface{}) *CertificateAuthorityIssuingOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &CertificateAuthorityIssuingOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyCertificateAuthorityIssuingOptions
	}
	r.IncludeCaCertUrl = dcl.FlattenBool(m["includeCaCertUrl"])
	r.IncludeCrlAccessUrl = dcl.FlattenBool(m["includeCrlAccessUrl"])

	return r
}

// flattenCertificateAuthorityTypeEnumSlice flattens the contents of CertificateAuthorityTypeEnum from a JSON
// response object.
func flattenCertificateAuthorityTypeEnumSlice(c *Client, i interface{}) []CertificateAuthorityTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityTypeEnum{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityTypeEnum{}
	}

	items := make([]CertificateAuthorityTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityTypeEnum(item.(interface{})))
	}

	return items
}

// flattenCertificateAuthorityTypeEnum asserts that an interface is a string, and returns a
// pointer to a *CertificateAuthorityTypeEnum with the same value as that string.
func flattenCertificateAuthorityTypeEnum(i interface{}) *CertificateAuthorityTypeEnum {
	s, ok := i.(string)
	if !ok {
		return CertificateAuthorityTypeEnumRef("")
	}

	return CertificateAuthorityTypeEnumRef(s)
}

// flattenCertificateAuthorityConfigPublicKeyTypeEnumSlice flattens the contents of CertificateAuthorityConfigPublicKeyTypeEnum from a JSON
// response object.
func flattenCertificateAuthorityConfigPublicKeyTypeEnumSlice(c *Client, i interface{}) []CertificateAuthorityConfigPublicKeyTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityConfigPublicKeyTypeEnum{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityConfigPublicKeyTypeEnum{}
	}

	items := make([]CertificateAuthorityConfigPublicKeyTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityConfigPublicKeyTypeEnum(item.(interface{})))
	}

	return items
}

// flattenCertificateAuthorityConfigPublicKeyTypeEnum asserts that an interface is a string, and returns a
// pointer to a *CertificateAuthorityConfigPublicKeyTypeEnum with the same value as that string.
func flattenCertificateAuthorityConfigPublicKeyTypeEnum(i interface{}) *CertificateAuthorityConfigPublicKeyTypeEnum {
	s, ok := i.(string)
	if !ok {
		return CertificateAuthorityConfigPublicKeyTypeEnumRef("")
	}

	return CertificateAuthorityConfigPublicKeyTypeEnumRef(s)
}

// flattenCertificateAuthorityKeySpecAlgorithmEnumSlice flattens the contents of CertificateAuthorityKeySpecAlgorithmEnum from a JSON
// response object.
func flattenCertificateAuthorityKeySpecAlgorithmEnumSlice(c *Client, i interface{}) []CertificateAuthorityKeySpecAlgorithmEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityKeySpecAlgorithmEnum{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityKeySpecAlgorithmEnum{}
	}

	items := make([]CertificateAuthorityKeySpecAlgorithmEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityKeySpecAlgorithmEnum(item.(interface{})))
	}

	return items
}

// flattenCertificateAuthorityKeySpecAlgorithmEnum asserts that an interface is a string, and returns a
// pointer to a *CertificateAuthorityKeySpecAlgorithmEnum with the same value as that string.
func flattenCertificateAuthorityKeySpecAlgorithmEnum(i interface{}) *CertificateAuthorityKeySpecAlgorithmEnum {
	s, ok := i.(string)
	if !ok {
		return CertificateAuthorityKeySpecAlgorithmEnumRef("")
	}

	return CertificateAuthorityKeySpecAlgorithmEnumRef(s)
}

// flattenCertificateAuthorityTierEnumSlice flattens the contents of CertificateAuthorityTierEnum from a JSON
// response object.
func flattenCertificateAuthorityTierEnumSlice(c *Client, i interface{}) []CertificateAuthorityTierEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityTierEnum{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityTierEnum{}
	}

	items := make([]CertificateAuthorityTierEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityTierEnum(item.(interface{})))
	}

	return items
}

// flattenCertificateAuthorityTierEnum asserts that an interface is a string, and returns a
// pointer to a *CertificateAuthorityTierEnum with the same value as that string.
func flattenCertificateAuthorityTierEnum(i interface{}) *CertificateAuthorityTierEnum {
	s, ok := i.(string)
	if !ok {
		return CertificateAuthorityTierEnumRef("")
	}

	return CertificateAuthorityTierEnumRef(s)
}

// flattenCertificateAuthorityStateEnumSlice flattens the contents of CertificateAuthorityStateEnum from a JSON
// response object.
func flattenCertificateAuthorityStateEnumSlice(c *Client, i interface{}) []CertificateAuthorityStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityStateEnum{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityStateEnum{}
	}

	items := make([]CertificateAuthorityStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityStateEnum(item.(interface{})))
	}

	return items
}

// flattenCertificateAuthorityStateEnum asserts that an interface is a string, and returns a
// pointer to a *CertificateAuthorityStateEnum with the same value as that string.
func flattenCertificateAuthorityStateEnum(i interface{}) *CertificateAuthorityStateEnum {
	s, ok := i.(string)
	if !ok {
		return CertificateAuthorityStateEnumRef("")
	}

	return CertificateAuthorityStateEnumRef(s)
}

// flattenCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(item.(interface{})))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum asserts that an interface is a string, and returns a
// pointer to a *CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum with the same value as that string.
func flattenCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(i interface{}) *CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum {
	s, ok := i.(string)
	if !ok {
		return CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumRef("")
	}

	return CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumRef(s)
}

// flattenCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnumSlice flattens the contents of CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum from a JSON
// response object.
func flattenCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnumSlice(c *Client, i interface{}) []CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum{}
	}

	if len(a) == 0 {
		return []CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum{}
	}

	items := make([]CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum(item.(interface{})))
	}

	return items
}

// flattenCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum asserts that an interface is a string, and returns a
// pointer to a *CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum with the same value as that string.
func flattenCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum(i interface{}) *CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum {
	s, ok := i.(string)
	if !ok {
		return CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnumRef("")
	}

	return CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *CertificateAuthority) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalCertificateAuthority(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type certificateAuthorityDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         certificateAuthorityApiOperation
}

func convertFieldDiffToCertificateAuthorityOp(ops []string, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]certificateAuthorityDiff, error) {
	var diffs []certificateAuthorityDiff
	for _, op := range ops {
		diff := certificateAuthorityDiff{}
		if op == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			op, err := convertOpNameTocertificateAuthorityApiOperation(op, fds, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = op
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameTocertificateAuthorityApiOperation(op string, diffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (certificateAuthorityApiOperation, error) {
	switch op {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", op)
	}
}
