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
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/recaptchaenterprise/alpha/recaptchaenterprise_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/recaptchaenterprise/alpha"
)

// Server implements the gRPC interface for Key.
type KeyServer struct{}

// ProtoToKeyWebSettingsIntegrationTypeEnum converts a KeyWebSettingsIntegrationTypeEnum enum from its proto representation.
func ProtoToRecaptchaenterpriseAlphaKeyWebSettingsIntegrationTypeEnum(e alphapb.RecaptchaenterpriseAlphaKeyWebSettingsIntegrationTypeEnum) *alpha.KeyWebSettingsIntegrationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RecaptchaenterpriseAlphaKeyWebSettingsIntegrationTypeEnum_name[int32(e)]; ok {
		e := alpha.KeyWebSettingsIntegrationTypeEnum(n[len("RecaptchaenterpriseAlphaKeyWebSettingsIntegrationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToKeyWebSettingsChallengeSecurityPreferenceEnum converts a KeyWebSettingsChallengeSecurityPreferenceEnum enum from its proto representation.
func ProtoToRecaptchaenterpriseAlphaKeyWebSettingsChallengeSecurityPreferenceEnum(e alphapb.RecaptchaenterpriseAlphaKeyWebSettingsChallengeSecurityPreferenceEnum) *alpha.KeyWebSettingsChallengeSecurityPreferenceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RecaptchaenterpriseAlphaKeyWebSettingsChallengeSecurityPreferenceEnum_name[int32(e)]; ok {
		e := alpha.KeyWebSettingsChallengeSecurityPreferenceEnum(n[len("RecaptchaenterpriseAlphaKeyWebSettingsChallengeSecurityPreferenceEnum"):])
		return &e
	}
	return nil
}

// ProtoToKeyTestingOptionsTestingChallengeEnum converts a KeyTestingOptionsTestingChallengeEnum enum from its proto representation.
func ProtoToRecaptchaenterpriseAlphaKeyTestingOptionsTestingChallengeEnum(e alphapb.RecaptchaenterpriseAlphaKeyTestingOptionsTestingChallengeEnum) *alpha.KeyTestingOptionsTestingChallengeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RecaptchaenterpriseAlphaKeyTestingOptionsTestingChallengeEnum_name[int32(e)]; ok {
		e := alpha.KeyTestingOptionsTestingChallengeEnum(n[len("RecaptchaenterpriseAlphaKeyTestingOptionsTestingChallengeEnum"):])
		return &e
	}
	return nil
}

// ProtoToKeyWebSettings converts a KeyWebSettings object from its proto representation.
func ProtoToRecaptchaenterpriseAlphaKeyWebSettings(p *alphapb.RecaptchaenterpriseAlphaKeyWebSettings) *alpha.KeyWebSettings {
	if p == nil {
		return nil
	}
	obj := &alpha.KeyWebSettings{
		AllowAllDomains:             dcl.Bool(p.GetAllowAllDomains()),
		AllowAmpTraffic:             dcl.Bool(p.GetAllowAmpTraffic()),
		IntegrationType:             ProtoToRecaptchaenterpriseAlphaKeyWebSettingsIntegrationTypeEnum(p.GetIntegrationType()),
		ChallengeSecurityPreference: ProtoToRecaptchaenterpriseAlphaKeyWebSettingsChallengeSecurityPreferenceEnum(p.GetChallengeSecurityPreference()),
	}
	for _, r := range p.GetAllowedDomains() {
		obj.AllowedDomains = append(obj.AllowedDomains, r)
	}
	return obj
}

// ProtoToKeyAndroidSettings converts a KeyAndroidSettings object from its proto representation.
func ProtoToRecaptchaenterpriseAlphaKeyAndroidSettings(p *alphapb.RecaptchaenterpriseAlphaKeyAndroidSettings) *alpha.KeyAndroidSettings {
	if p == nil {
		return nil
	}
	obj := &alpha.KeyAndroidSettings{
		AllowAllPackageNames: dcl.Bool(p.GetAllowAllPackageNames()),
	}
	for _, r := range p.GetAllowedPackageNames() {
		obj.AllowedPackageNames = append(obj.AllowedPackageNames, r)
	}
	return obj
}

// ProtoToKeyIosSettings converts a KeyIosSettings object from its proto representation.
func ProtoToRecaptchaenterpriseAlphaKeyIosSettings(p *alphapb.RecaptchaenterpriseAlphaKeyIosSettings) *alpha.KeyIosSettings {
	if p == nil {
		return nil
	}
	obj := &alpha.KeyIosSettings{
		AllowAllBundleIds: dcl.Bool(p.GetAllowAllBundleIds()),
	}
	for _, r := range p.GetAllowedBundleIds() {
		obj.AllowedBundleIds = append(obj.AllowedBundleIds, r)
	}
	return obj
}

// ProtoToKeyTestingOptions converts a KeyTestingOptions object from its proto representation.
func ProtoToRecaptchaenterpriseAlphaKeyTestingOptions(p *alphapb.RecaptchaenterpriseAlphaKeyTestingOptions) *alpha.KeyTestingOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.KeyTestingOptions{
		TestingScore:     dcl.Float64OrNil(p.GetTestingScore()),
		TestingChallenge: ProtoToRecaptchaenterpriseAlphaKeyTestingOptionsTestingChallengeEnum(p.GetTestingChallenge()),
	}
	return obj
}

// ProtoToKey converts a Key resource from its proto representation.
func ProtoToKey(p *alphapb.RecaptchaenterpriseAlphaKey) *alpha.Key {
	obj := &alpha.Key{
		Name:            dcl.StringOrNil(p.GetName()),
		DisplayName:     dcl.StringOrNil(p.GetDisplayName()),
		WebSettings:     ProtoToRecaptchaenterpriseAlphaKeyWebSettings(p.GetWebSettings()),
		AndroidSettings: ProtoToRecaptchaenterpriseAlphaKeyAndroidSettings(p.GetAndroidSettings()),
		IosSettings:     ProtoToRecaptchaenterpriseAlphaKeyIosSettings(p.GetIosSettings()),
		CreateTime:      dcl.StringOrNil(p.GetCreateTime()),
		TestingOptions:  ProtoToRecaptchaenterpriseAlphaKeyTestingOptions(p.GetTestingOptions()),
		Project:         dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// KeyWebSettingsIntegrationTypeEnumToProto converts a KeyWebSettingsIntegrationTypeEnum enum to its proto representation.
func RecaptchaenterpriseAlphaKeyWebSettingsIntegrationTypeEnumToProto(e *alpha.KeyWebSettingsIntegrationTypeEnum) alphapb.RecaptchaenterpriseAlphaKeyWebSettingsIntegrationTypeEnum {
	if e == nil {
		return alphapb.RecaptchaenterpriseAlphaKeyWebSettingsIntegrationTypeEnum(0)
	}
	if v, ok := alphapb.RecaptchaenterpriseAlphaKeyWebSettingsIntegrationTypeEnum_value["KeyWebSettingsIntegrationTypeEnum"+string(*e)]; ok {
		return alphapb.RecaptchaenterpriseAlphaKeyWebSettingsIntegrationTypeEnum(v)
	}
	return alphapb.RecaptchaenterpriseAlphaKeyWebSettingsIntegrationTypeEnum(0)
}

// KeyWebSettingsChallengeSecurityPreferenceEnumToProto converts a KeyWebSettingsChallengeSecurityPreferenceEnum enum to its proto representation.
func RecaptchaenterpriseAlphaKeyWebSettingsChallengeSecurityPreferenceEnumToProto(e *alpha.KeyWebSettingsChallengeSecurityPreferenceEnum) alphapb.RecaptchaenterpriseAlphaKeyWebSettingsChallengeSecurityPreferenceEnum {
	if e == nil {
		return alphapb.RecaptchaenterpriseAlphaKeyWebSettingsChallengeSecurityPreferenceEnum(0)
	}
	if v, ok := alphapb.RecaptchaenterpriseAlphaKeyWebSettingsChallengeSecurityPreferenceEnum_value["KeyWebSettingsChallengeSecurityPreferenceEnum"+string(*e)]; ok {
		return alphapb.RecaptchaenterpriseAlphaKeyWebSettingsChallengeSecurityPreferenceEnum(v)
	}
	return alphapb.RecaptchaenterpriseAlphaKeyWebSettingsChallengeSecurityPreferenceEnum(0)
}

// KeyTestingOptionsTestingChallengeEnumToProto converts a KeyTestingOptionsTestingChallengeEnum enum to its proto representation.
func RecaptchaenterpriseAlphaKeyTestingOptionsTestingChallengeEnumToProto(e *alpha.KeyTestingOptionsTestingChallengeEnum) alphapb.RecaptchaenterpriseAlphaKeyTestingOptionsTestingChallengeEnum {
	if e == nil {
		return alphapb.RecaptchaenterpriseAlphaKeyTestingOptionsTestingChallengeEnum(0)
	}
	if v, ok := alphapb.RecaptchaenterpriseAlphaKeyTestingOptionsTestingChallengeEnum_value["KeyTestingOptionsTestingChallengeEnum"+string(*e)]; ok {
		return alphapb.RecaptchaenterpriseAlphaKeyTestingOptionsTestingChallengeEnum(v)
	}
	return alphapb.RecaptchaenterpriseAlphaKeyTestingOptionsTestingChallengeEnum(0)
}

// KeyWebSettingsToProto converts a KeyWebSettings object to its proto representation.
func RecaptchaenterpriseAlphaKeyWebSettingsToProto(o *alpha.KeyWebSettings) *alphapb.RecaptchaenterpriseAlphaKeyWebSettings {
	if o == nil {
		return nil
	}
	p := &alphapb.RecaptchaenterpriseAlphaKeyWebSettings{}
	p.SetAllowAllDomains(dcl.ValueOrEmptyBool(o.AllowAllDomains))
	p.SetAllowAmpTraffic(dcl.ValueOrEmptyBool(o.AllowAmpTraffic))
	p.SetIntegrationType(RecaptchaenterpriseAlphaKeyWebSettingsIntegrationTypeEnumToProto(o.IntegrationType))
	p.SetChallengeSecurityPreference(RecaptchaenterpriseAlphaKeyWebSettingsChallengeSecurityPreferenceEnumToProto(o.ChallengeSecurityPreference))
	sAllowedDomains := make([]string, len(o.AllowedDomains))
	for i, r := range o.AllowedDomains {
		sAllowedDomains[i] = r
	}
	p.SetAllowedDomains(sAllowedDomains)
	return p
}

// KeyAndroidSettingsToProto converts a KeyAndroidSettings object to its proto representation.
func RecaptchaenterpriseAlphaKeyAndroidSettingsToProto(o *alpha.KeyAndroidSettings) *alphapb.RecaptchaenterpriseAlphaKeyAndroidSettings {
	if o == nil {
		return nil
	}
	p := &alphapb.RecaptchaenterpriseAlphaKeyAndroidSettings{}
	p.SetAllowAllPackageNames(dcl.ValueOrEmptyBool(o.AllowAllPackageNames))
	sAllowedPackageNames := make([]string, len(o.AllowedPackageNames))
	for i, r := range o.AllowedPackageNames {
		sAllowedPackageNames[i] = r
	}
	p.SetAllowedPackageNames(sAllowedPackageNames)
	return p
}

// KeyIosSettingsToProto converts a KeyIosSettings object to its proto representation.
func RecaptchaenterpriseAlphaKeyIosSettingsToProto(o *alpha.KeyIosSettings) *alphapb.RecaptchaenterpriseAlphaKeyIosSettings {
	if o == nil {
		return nil
	}
	p := &alphapb.RecaptchaenterpriseAlphaKeyIosSettings{}
	p.SetAllowAllBundleIds(dcl.ValueOrEmptyBool(o.AllowAllBundleIds))
	sAllowedBundleIds := make([]string, len(o.AllowedBundleIds))
	for i, r := range o.AllowedBundleIds {
		sAllowedBundleIds[i] = r
	}
	p.SetAllowedBundleIds(sAllowedBundleIds)
	return p
}

// KeyTestingOptionsToProto converts a KeyTestingOptions object to its proto representation.
func RecaptchaenterpriseAlphaKeyTestingOptionsToProto(o *alpha.KeyTestingOptions) *alphapb.RecaptchaenterpriseAlphaKeyTestingOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.RecaptchaenterpriseAlphaKeyTestingOptions{}
	p.SetTestingScore(dcl.ValueOrEmptyDouble(o.TestingScore))
	p.SetTestingChallenge(RecaptchaenterpriseAlphaKeyTestingOptionsTestingChallengeEnumToProto(o.TestingChallenge))
	return p
}

// KeyToProto converts a Key resource to its proto representation.
func KeyToProto(resource *alpha.Key) *alphapb.RecaptchaenterpriseAlphaKey {
	p := &alphapb.RecaptchaenterpriseAlphaKey{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetWebSettings(RecaptchaenterpriseAlphaKeyWebSettingsToProto(resource.WebSettings))
	p.SetAndroidSettings(RecaptchaenterpriseAlphaKeyAndroidSettingsToProto(resource.AndroidSettings))
	p.SetIosSettings(RecaptchaenterpriseAlphaKeyIosSettingsToProto(resource.IosSettings))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetTestingOptions(RecaptchaenterpriseAlphaKeyTestingOptionsToProto(resource.TestingOptions))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyKey handles the gRPC request by passing it to the underlying Key Apply() method.
func (s *KeyServer) applyKey(ctx context.Context, c *alpha.Client, request *alphapb.ApplyRecaptchaenterpriseAlphaKeyRequest) (*alphapb.RecaptchaenterpriseAlphaKey, error) {
	p := ProtoToKey(request.GetResource())
	res, err := c.ApplyKey(ctx, p)
	if err != nil {
		return nil, err
	}
	r := KeyToProto(res)
	return r, nil
}

// applyRecaptchaenterpriseAlphaKey handles the gRPC request by passing it to the underlying Key Apply() method.
func (s *KeyServer) ApplyRecaptchaenterpriseAlphaKey(ctx context.Context, request *alphapb.ApplyRecaptchaenterpriseAlphaKeyRequest) (*alphapb.RecaptchaenterpriseAlphaKey, error) {
	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyKey(ctx, cl, request)
}

// DeleteKey handles the gRPC request by passing it to the underlying Key Delete() method.
func (s *KeyServer) DeleteRecaptchaenterpriseAlphaKey(ctx context.Context, request *alphapb.DeleteRecaptchaenterpriseAlphaKeyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteKey(ctx, ProtoToKey(request.GetResource()))

}

// ListRecaptchaenterpriseAlphaKey handles the gRPC request by passing it to the underlying KeyList() method.
func (s *KeyServer) ListRecaptchaenterpriseAlphaKey(ctx context.Context, request *alphapb.ListRecaptchaenterpriseAlphaKeyRequest) (*alphapb.ListRecaptchaenterpriseAlphaKeyResponse, error) {
	cl, err := createConfigKey(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListKey(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.RecaptchaenterpriseAlphaKey
	for _, r := range resources.Items {
		rp := KeyToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListRecaptchaenterpriseAlphaKeyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigKey(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
