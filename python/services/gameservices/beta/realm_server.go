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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gameservices/beta/gameservices_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gameservices/beta"
)

// Server implements the gRPC interface for Realm.
type RealmServer struct{}

// ProtoToRealm converts a Realm resource from its proto representation.
func ProtoToRealm(p *betapb.GameservicesBetaRealm) *beta.Realm {
	obj := &beta.Realm{
		Name:        dcl.StringOrNil(p.Name),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		TimeZone:    dcl.StringOrNil(p.TimeZone),
		Description: dcl.StringOrNil(p.Description),
		Location:    dcl.StringOrNil(p.Location),
		Project:     dcl.StringOrNil(p.Project),
	}
	return obj
}

// RealmToProto converts a Realm resource to its proto representation.
func RealmToProto(resource *beta.Realm) *betapb.GameservicesBetaRealm {
	p := &betapb.GameservicesBetaRealm{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		CreateTime:  dcl.ValueOrEmptyString(resource.CreateTime),
		TimeZone:    dcl.ValueOrEmptyString(resource.TimeZone),
		Description: dcl.ValueOrEmptyString(resource.Description),
		Location:    dcl.ValueOrEmptyString(resource.Location),
		Project:     dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyRealm handles the gRPC request by passing it to the underlying Realm Apply() method.
func (s *RealmServer) applyRealm(ctx context.Context, c *beta.Client, request *betapb.ApplyGameservicesBetaRealmRequest) (*betapb.GameservicesBetaRealm, error) {
	p := ProtoToRealm(request.GetResource())
	res, err := c.ApplyRealm(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RealmToProto(res)
	return r, nil
}

// ApplyRealm handles the gRPC request by passing it to the underlying Realm Apply() method.
func (s *RealmServer) ApplyGameservicesBetaRealm(ctx context.Context, request *betapb.ApplyGameservicesBetaRealmRequest) (*betapb.GameservicesBetaRealm, error) {
	cl, err := createConfigRealm(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyRealm(ctx, cl, request)
}

// DeleteRealm handles the gRPC request by passing it to the underlying Realm Delete() method.
func (s *RealmServer) DeleteGameservicesBetaRealm(ctx context.Context, request *betapb.DeleteGameservicesBetaRealmRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRealm(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRealm(ctx, ProtoToRealm(request.GetResource()))

}

// ListGameservicesBetaRealm handles the gRPC request by passing it to the underlying RealmList() method.
func (s *RealmServer) ListGameservicesBetaRealm(ctx context.Context, request *betapb.ListGameservicesBetaRealmRequest) (*betapb.ListGameservicesBetaRealmResponse, error) {
	cl, err := createConfigRealm(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRealm(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.GameservicesBetaRealm
	for _, r := range resources.Items {
		rp := RealmToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListGameservicesBetaRealmResponse{Items: protos}, nil
}

func createConfigRealm(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
