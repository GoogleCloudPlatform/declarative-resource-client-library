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
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	apikeyspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apikeys/apikeys_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apikeys"
)

// Server implements the gRPC interface for Key.
type KeyServer struct{}

// ProtoToKeyRestrictions converts a KeyRestrictions resource from its proto representation.
func ProtoToApikeysKeyRestrictions(p *apikeyspb.ApikeysKeyRestrictions) *apikeys.KeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &apikeys.KeyRestrictions{
		BrowserKeyRestrictions: ProtoToApikeysKeyRestrictionsBrowserKeyRestrictions(p.GetBrowserKeyRestrictions()),
		ServerKeyRestrictions:  ProtoToApikeysKeyRestrictionsServerKeyRestrictions(p.GetServerKeyRestrictions()),
		AndroidKeyRestrictions: ProtoToApikeysKeyRestrictionsAndroidKeyRestrictions(p.GetAndroidKeyRestrictions()),
		IosKeyRestrictions:     ProtoToApikeysKeyRestrictionsIosKeyRestrictions(p.GetIosKeyRestrictions()),
	}
	for _, r := range p.GetApiTargets() {
		obj.ApiTargets = append(obj.ApiTargets, *ProtoToApikeysKeyRestrictionsApiTargets(r))
	}
	return obj
}

// ProtoToKeyRestrictionsBrowserKeyRestrictions converts a KeyRestrictionsBrowserKeyRestrictions resource from its proto representation.
func ProtoToApikeysKeyRestrictionsBrowserKeyRestrictions(p *apikeyspb.ApikeysKeyRestrictionsBrowserKeyRestrictions) *apikeys.KeyRestrictionsBrowserKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &apikeys.KeyRestrictionsBrowserKeyRestrictions{}
	for _, r := range p.GetAllowedReferrers() {
		obj.AllowedReferrers = append(obj.AllowedReferrers, r)
	}
	return obj
}

// ProtoToKeyRestrictionsServerKeyRestrictions converts a KeyRestrictionsServerKeyRestrictions resource from its proto representation.
func ProtoToApikeysKeyRestrictionsServerKeyRestrictions(p *apikeyspb.ApikeysKeyRestrictionsServerKeyRestrictions) *apikeys.KeyRestrictionsServerKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &apikeys.KeyRestrictionsServerKeyRestrictions{}
	for _, r := range p.GetAllowedIps() {
		obj.AllowedIps = append(obj.AllowedIps, r)
	}
	return obj
}

// ProtoToKeyRestrictionsAndroidKeyRestrictions converts a KeyRestrictionsAndroidKeyRestrictions resource from its proto representation.
func ProtoToApikeysKeyRestrictionsAndroidKeyRestrictions(p *apikeyspb.ApikeysKeyRestrictionsAndroidKeyRestrictions) *apikeys.KeyRestrictionsAndroidKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &apikeys.KeyRestrictionsAndroidKeyRestrictions{}
	for _, r := range p.GetAllowedApplications() {
		obj.AllowedApplications = append(obj.AllowedApplications, *ProtoToApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(r))
	}
	return obj
}

// ProtoToKeyRestrictionsAndroidKeyRestrictionsAllowedApplications converts a KeyRestrictionsAndroidKeyRestrictionsAllowedApplications resource from its proto representation.
func ProtoToApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications(p *apikeyspb.ApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications) *apikeys.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	if p == nil {
		return nil
	}
	obj := &apikeys.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications{
		Sha1Fingerprint: dcl.StringOrNil(p.Sha1Fingerprint),
		PackageName:     dcl.StringOrNil(p.PackageName),
	}
	return obj
}

// ProtoToKeyRestrictionsIosKeyRestrictions converts a KeyRestrictionsIosKeyRestrictions resource from its proto representation.
func ProtoToApikeysKeyRestrictionsIosKeyRestrictions(p *apikeyspb.ApikeysKeyRestrictionsIosKeyRestrictions) *apikeys.KeyRestrictionsIosKeyRestrictions {
	if p == nil {
		return nil
	}
	obj := &apikeys.KeyRestrictionsIosKeyRestrictions{}
	for _, r := range p.GetAllowedBundleIds() {
		obj.AllowedBundleIds = append(obj.AllowedBundleIds, r)
	}
	return obj
}

// ProtoToKeyRestrictionsApiTargets converts a KeyRestrictionsApiTargets resource from its proto representation.
func ProtoToApikeysKeyRestrictionsApiTargets(p *apikeyspb.ApikeysKeyRestrictionsApiTargets) *apikeys.KeyRestrictionsApiTargets {
	if p == nil {
		return nil
	}
	obj := &apikeys.KeyRestrictionsApiTargets{
		Service: dcl.StringOrNil(p.Service),
	}
	for _, r := range p.GetMethods() {
		obj.Methods = append(obj.Methods, r)
	}
	return obj
}

// ProtoToKey converts a Key resource from its proto representation.
func ProtoToKey(p *apikeyspb.ApikeysKey) *apikeys.Key {
	obj := &apikeys.Key{
		Name:         dcl.StringOrNil(p.Name),
		Uid:          dcl.StringOrNil(p.Uid),
		DisplayName:  dcl.StringOrNil(p.DisplayName),
		KeyString:    dcl.StringOrNil(p.KeyString),
		CreateTime:   dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:   dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:   dcl.StringOrNil(p.GetDeleteTime()),
		Restrictions: ProtoToApikeysKeyRestrictions(p.GetRestrictions()),
		Etag:         dcl.StringOrNil(p.Etag),
		Project:      dcl.StringOrNil(p.Project),
	}
	return obj
}

// KeyRestrictionsToProto converts a KeyRestrictions resource to its proto representation.
func ApikeysKeyRestrictionsToProto(o *apikeys.KeyRestrictions) *apikeyspb.ApikeysKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &apikeyspb.ApikeysKeyRestrictions{
		BrowserKeyRestrictions: ApikeysKeyRestrictionsBrowserKeyRestrictionsToProto(o.BrowserKeyRestrictions),
		ServerKeyRestrictions:  ApikeysKeyRestrictionsServerKeyRestrictionsToProto(o.ServerKeyRestrictions),
		AndroidKeyRestrictions: ApikeysKeyRestrictionsAndroidKeyRestrictionsToProto(o.AndroidKeyRestrictions),
		IosKeyRestrictions:     ApikeysKeyRestrictionsIosKeyRestrictionsToProto(o.IosKeyRestrictions),
	}
	for _, r := range o.ApiTargets {
		p.ApiTargets = append(p.ApiTargets, ApikeysKeyRestrictionsApiTargetsToProto(&r))
	}
	return p
}

// KeyRestrictionsBrowserKeyRestrictionsToProto converts a KeyRestrictionsBrowserKeyRestrictions resource to its proto representation.
func ApikeysKeyRestrictionsBrowserKeyRestrictionsToProto(o *apikeys.KeyRestrictionsBrowserKeyRestrictions) *apikeyspb.ApikeysKeyRestrictionsBrowserKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &apikeyspb.ApikeysKeyRestrictionsBrowserKeyRestrictions{}
	for _, r := range o.AllowedReferrers {
		p.AllowedReferrers = append(p.AllowedReferrers, r)
	}
	return p
}

// KeyRestrictionsServerKeyRestrictionsToProto converts a KeyRestrictionsServerKeyRestrictions resource to its proto representation.
func ApikeysKeyRestrictionsServerKeyRestrictionsToProto(o *apikeys.KeyRestrictionsServerKeyRestrictions) *apikeyspb.ApikeysKeyRestrictionsServerKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &apikeyspb.ApikeysKeyRestrictionsServerKeyRestrictions{}
	for _, r := range o.AllowedIps {
		p.AllowedIps = append(p.AllowedIps, r)
	}
	return p
}

// KeyRestrictionsAndroidKeyRestrictionsToProto converts a KeyRestrictionsAndroidKeyRestrictions resource to its proto representation.
func ApikeysKeyRestrictionsAndroidKeyRestrictionsToProto(o *apikeys.KeyRestrictionsAndroidKeyRestrictions) *apikeyspb.ApikeysKeyRestrictionsAndroidKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &apikeyspb.ApikeysKeyRestrictionsAndroidKeyRestrictions{}
	for _, r := range o.AllowedApplications {
		p.AllowedApplications = append(p.AllowedApplications, ApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsToProto(&r))
	}
	return p
}

// KeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsToProto converts a KeyRestrictionsAndroidKeyRestrictionsAllowedApplications resource to its proto representation.
func ApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplicationsToProto(o *apikeys.KeyRestrictionsAndroidKeyRestrictionsAllowedApplications) *apikeyspb.ApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications {
	if o == nil {
		return nil
	}
	p := &apikeyspb.ApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications{
		Sha1Fingerprint: dcl.ValueOrEmptyString(o.Sha1Fingerprint),
		PackageName:     dcl.ValueOrEmptyString(o.PackageName),
	}
	return p
}

// KeyRestrictionsIosKeyRestrictionsToProto converts a KeyRestrictionsIosKeyRestrictions resource to its proto representation.
func ApikeysKeyRestrictionsIosKeyRestrictionsToProto(o *apikeys.KeyRestrictionsIosKeyRestrictions) *apikeyspb.ApikeysKeyRestrictionsIosKeyRestrictions {
	if o == nil {
		return nil
	}
	p := &apikeyspb.ApikeysKeyRestrictionsIosKeyRestrictions{}
	for _, r := range o.AllowedBundleIds {
		p.AllowedBundleIds = append(p.AllowedBundleIds, r)
	}
	return p
}

// KeyRestrictionsApiTargetsToProto converts a KeyRestrictionsApiTargets resource to its proto representation.
func ApikeysKeyRestrictionsApiTargetsToProto(o *apikeys.KeyRestrictionsApiTargets) *apikeyspb.ApikeysKeyRestrictionsApiTargets {
	if o == nil {
		return nil
	}
	p := &apikeyspb.ApikeysKeyRestrictionsApiTargets{
		Service: dcl.ValueOrEmptyString(o.Service),
	}
	for _, r := range o.Methods {
		p.Methods = append(p.Methods, r)
	}
	return p
}

// KeyToProto converts a Key resource to its proto representation.
func KeyToProto(resource *apikeys.Key) *apikeyspb.ApikeysKey {
	p := &apikeyspb.ApikeysKey{
		Name:         dcl.ValueOrEmptyString(resource.Name),
		Uid:          dcl.ValueOrEmptyString(resource.Uid),
		DisplayName:  dcl.ValueOrEmptyString(resource.DisplayName),
		KeyString:    dcl.ValueOrEmptyString(resource.KeyString),
		CreateTime:   dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:   dcl.ValueOrEmptyString(resource.UpdateTime),
		DeleteTime:   dcl.ValueOrEmptyString(resource.DeleteTime),
		Restrictions: ApikeysKeyRestrictionsToProto(resource.Restrictions),
		Etag:         dcl.ValueOrEmptyString(resource.Etag),
		Project:      dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyKey handles the gRPC request by passing it to the underlying Key Apply() method.
func (s *KeyServer) applyKey(ctx context.Context, c *apikeys.Client, request *apikeyspb.ApplyApikeysKeyRequest) (*apikeyspb.ApikeysKey, error) {
	p := ProtoToKey(request.GetResource())
	res, err := c.ApplyKey(ctx, p)
	if err != nil {
		return nil, err
	}
	r := KeyToProto(res)
	return r, nil
}

// ApplyKey handles the gRPC request by passing it to the underlying Key Apply() method.
func (s *KeyServer) ApplyApikeysKey(ctx context.Context, request *apikeyspb.ApplyApikeysKeyRequest) (*apikeyspb.ApikeysKey, error) {
	cl, err := createConfigKey(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyKey(ctx, cl, request)
}

// DeleteKey handles the gRPC request by passing it to the underlying Key Delete() method.
func (s *KeyServer) DeleteApikeysKey(ctx context.Context, request *apikeyspb.DeleteApikeysKeyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigKey(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteKey(ctx, ProtoToKey(request.GetResource()))

}

// ListApikeysKey handles the gRPC request by passing it to the underlying KeyList() method.
func (s *KeyServer) ListApikeysKey(ctx context.Context, request *apikeyspb.ListApikeysKeyRequest) (*apikeyspb.ListApikeysKeyResponse, error) {
	cl, err := createConfigKey(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListKey(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*apikeyspb.ApikeysKey
	for _, r := range resources.Items {
		rp := KeyToProto(r)
		protos = append(protos, rp)
	}
	return &apikeyspb.ListApikeysKeyResponse{Items: protos}, nil
}

func createConfigKey(ctx context.Context, service_account_file string) (*apikeys.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return apikeys.NewClient(conf), nil
}
