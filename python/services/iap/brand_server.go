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
	"errors"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	iappb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iap/iap_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap"
)

// Server implements the gRPC interface for Brand.
type BrandServer struct{}

// ProtoToBrand converts a Brand resource from its proto representation.
func ProtoToBrand(p *iappb.IapBrand) *iap.Brand {
	obj := &iap.Brand{
		ApplicationTitle: dcl.StringOrNil(p.ApplicationTitle),
		Name:             dcl.StringOrNil(p.Name),
		OrgInternalOnly:  dcl.Bool(p.OrgInternalOnly),
		SupportEmail:     dcl.StringOrNil(p.SupportEmail),
		Project:          dcl.StringOrNil(p.Project),
	}
	return obj
}

// BrandToProto converts a Brand resource to its proto representation.
func BrandToProto(resource *iap.Brand) *iappb.IapBrand {
	p := &iappb.IapBrand{
		ApplicationTitle: dcl.ValueOrEmptyString(resource.ApplicationTitle),
		Name:             dcl.ValueOrEmptyString(resource.Name),
		OrgInternalOnly:  dcl.ValueOrEmptyBool(resource.OrgInternalOnly),
		SupportEmail:     dcl.ValueOrEmptyString(resource.SupportEmail),
		Project:          dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyBrand handles the gRPC request by passing it to the underlying Brand Apply() method.
func (s *BrandServer) applyBrand(ctx context.Context, c *iap.Client, request *iappb.ApplyIapBrandRequest) (*iappb.IapBrand, error) {
	p := ProtoToBrand(request.GetResource())
	res, err := c.ApplyBrand(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BrandToProto(res)
	return r, nil
}

// ApplyBrand handles the gRPC request by passing it to the underlying Brand Apply() method.
func (s *BrandServer) ApplyIapBrand(ctx context.Context, request *iappb.ApplyIapBrandRequest) (*iappb.IapBrand, error) {
	cl, err := createConfigBrand(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyBrand(ctx, cl, request)
}

// DeleteBrand handles the gRPC request by passing it to the underlying Brand Delete() method.
func (s *BrandServer) DeleteIapBrand(ctx context.Context, request *iappb.DeleteIapBrandRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for Brand")

}

// ListIapBrand handles the gRPC request by passing it to the underlying BrandList() method.
func (s *BrandServer) ListIapBrand(ctx context.Context, request *iappb.ListIapBrandRequest) (*iappb.ListIapBrandResponse, error) {
	cl, err := createConfigBrand(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBrand(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*iappb.IapBrand
	for _, r := range resources.Items {
		rp := BrandToProto(r)
		protos = append(protos, rp)
	}
	return &iappb.ListIapBrandResponse{Items: protos}, nil
}

func createConfigBrand(ctx context.Context, service_account_file string) (*iap.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return iap.NewClient(conf), nil
}
