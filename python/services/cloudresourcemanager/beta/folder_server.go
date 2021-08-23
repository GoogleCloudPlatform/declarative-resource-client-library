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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudresourcemanager/beta/cloudresourcemanager_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager/beta"
)

// Server implements the gRPC interface for Folder.
type FolderServer struct{}

// ProtoToFolderStateEnum converts a FolderStateEnum enum from its proto representation.
func ProtoToCloudresourcemanagerBetaFolderStateEnum(e betapb.CloudresourcemanagerBetaFolderStateEnum) *beta.FolderStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudresourcemanagerBetaFolderStateEnum_name[int32(e)]; ok {
		e := beta.FolderStateEnum(n[len("CloudresourcemanagerBetaFolderStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToFolder converts a Folder resource from its proto representation.
func ProtoToFolder(p *betapb.CloudresourcemanagerBetaFolder) *beta.Folder {
	obj := &beta.Folder{
		Name:        dcl.StringOrNil(p.Name),
		Parent:      dcl.StringOrNil(p.Parent),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		State:       ProtoToCloudresourcemanagerBetaFolderStateEnum(p.GetState()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:  dcl.StringOrNil(p.GetDeleteTime()),
		Etag:        dcl.StringOrNil(p.Etag),
	}
	return obj
}

// FolderStateEnumToProto converts a FolderStateEnum enum to its proto representation.
func CloudresourcemanagerBetaFolderStateEnumToProto(e *beta.FolderStateEnum) betapb.CloudresourcemanagerBetaFolderStateEnum {
	if e == nil {
		return betapb.CloudresourcemanagerBetaFolderStateEnum(0)
	}
	if v, ok := betapb.CloudresourcemanagerBetaFolderStateEnum_value["FolderStateEnum"+string(*e)]; ok {
		return betapb.CloudresourcemanagerBetaFolderStateEnum(v)
	}
	return betapb.CloudresourcemanagerBetaFolderStateEnum(0)
}

// FolderToProto converts a Folder resource to its proto representation.
func FolderToProto(resource *beta.Folder) *betapb.CloudresourcemanagerBetaFolder {
	p := &betapb.CloudresourcemanagerBetaFolder{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Parent:      dcl.ValueOrEmptyString(resource.Parent),
		DisplayName: dcl.ValueOrEmptyString(resource.DisplayName),
		State:       CloudresourcemanagerBetaFolderStateEnumToProto(resource.State),
		CreateTime:  dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:  dcl.ValueOrEmptyString(resource.UpdateTime),
		DeleteTime:  dcl.ValueOrEmptyString(resource.DeleteTime),
		Etag:        dcl.ValueOrEmptyString(resource.Etag),
	}

	return p
}

// ApplyFolder handles the gRPC request by passing it to the underlying Folder Apply() method.
func (s *FolderServer) applyFolder(ctx context.Context, c *beta.Client, request *betapb.ApplyCloudresourcemanagerBetaFolderRequest) (*betapb.CloudresourcemanagerBetaFolder, error) {
	p := ProtoToFolder(request.GetResource())
	res, err := c.ApplyFolder(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FolderToProto(res)
	return r, nil
}

// ApplyFolder handles the gRPC request by passing it to the underlying Folder Apply() method.
func (s *FolderServer) ApplyCloudresourcemanagerBetaFolder(ctx context.Context, request *betapb.ApplyCloudresourcemanagerBetaFolderRequest) (*betapb.CloudresourcemanagerBetaFolder, error) {
	cl, err := createConfigFolder(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyFolder(ctx, cl, request)
}

// DeleteFolder handles the gRPC request by passing it to the underlying Folder Delete() method.
func (s *FolderServer) DeleteCloudresourcemanagerBetaFolder(ctx context.Context, request *betapb.DeleteCloudresourcemanagerBetaFolderRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFolder(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFolder(ctx, ProtoToFolder(request.GetResource()))

}

// ListCloudresourcemanagerBetaFolder handles the gRPC request by passing it to the underlying FolderList() method.
func (s *FolderServer) ListCloudresourcemanagerBetaFolder(ctx context.Context, request *betapb.ListCloudresourcemanagerBetaFolderRequest) (*betapb.ListCloudresourcemanagerBetaFolderResponse, error) {
	cl, err := createConfigFolder(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFolder(ctx, ProtoToFolder(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.CloudresourcemanagerBetaFolder
	for _, r := range resources.Items {
		rp := FolderToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListCloudresourcemanagerBetaFolderResponse{Items: protos}, nil
}

func createConfigFolder(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
