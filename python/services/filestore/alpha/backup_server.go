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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/filestore/alpha/filestore_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore/alpha"
)

// Server implements the gRPC interface for Backup.
type BackupServer struct{}

// ProtoToBackupStateEnum converts a BackupStateEnum enum from its proto representation.
func ProtoToFilestoreAlphaBackupStateEnum(e alphapb.FilestoreAlphaBackupStateEnum) *alpha.BackupStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FilestoreAlphaBackupStateEnum_name[int32(e)]; ok {
		e := alpha.BackupStateEnum(n[len("FilestoreAlphaBackupStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackupSourceInstanceTierEnum converts a BackupSourceInstanceTierEnum enum from its proto representation.
func ProtoToFilestoreAlphaBackupSourceInstanceTierEnum(e alphapb.FilestoreAlphaBackupSourceInstanceTierEnum) *alpha.BackupSourceInstanceTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FilestoreAlphaBackupSourceInstanceTierEnum_name[int32(e)]; ok {
		e := alpha.BackupSourceInstanceTierEnum(n[len("FilestoreAlphaBackupSourceInstanceTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackup converts a Backup resource from its proto representation.
func ProtoToBackup(p *alphapb.FilestoreAlphaBackup) *alpha.Backup {
	obj := &alpha.Backup{
		Name:               dcl.StringOrNil(p.Name),
		Description:        dcl.StringOrNil(p.Description),
		State:              ProtoToFilestoreAlphaBackupStateEnum(p.GetState()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		CapacityGb:         dcl.Int64OrNil(p.CapacityGb),
		StorageBytes:       dcl.Int64OrNil(p.StorageBytes),
		SourceInstance:     dcl.StringOrNil(p.SourceInstance),
		SourceFileShare:    dcl.StringOrNil(p.SourceFileShare),
		SourceInstanceTier: ProtoToFilestoreAlphaBackupSourceInstanceTierEnum(p.GetSourceInstanceTier()),
		DownloadBytes:      dcl.Int64OrNil(p.DownloadBytes),
		Project:            dcl.StringOrNil(p.Project),
		Location:           dcl.StringOrNil(p.Location),
	}
	return obj
}

// BackupStateEnumToProto converts a BackupStateEnum enum to its proto representation.
func FilestoreAlphaBackupStateEnumToProto(e *alpha.BackupStateEnum) alphapb.FilestoreAlphaBackupStateEnum {
	if e == nil {
		return alphapb.FilestoreAlphaBackupStateEnum(0)
	}
	if v, ok := alphapb.FilestoreAlphaBackupStateEnum_value["BackupStateEnum"+string(*e)]; ok {
		return alphapb.FilestoreAlphaBackupStateEnum(v)
	}
	return alphapb.FilestoreAlphaBackupStateEnum(0)
}

// BackupSourceInstanceTierEnumToProto converts a BackupSourceInstanceTierEnum enum to its proto representation.
func FilestoreAlphaBackupSourceInstanceTierEnumToProto(e *alpha.BackupSourceInstanceTierEnum) alphapb.FilestoreAlphaBackupSourceInstanceTierEnum {
	if e == nil {
		return alphapb.FilestoreAlphaBackupSourceInstanceTierEnum(0)
	}
	if v, ok := alphapb.FilestoreAlphaBackupSourceInstanceTierEnum_value["BackupSourceInstanceTierEnum"+string(*e)]; ok {
		return alphapb.FilestoreAlphaBackupSourceInstanceTierEnum(v)
	}
	return alphapb.FilestoreAlphaBackupSourceInstanceTierEnum(0)
}

// BackupToProto converts a Backup resource to its proto representation.
func BackupToProto(resource *alpha.Backup) *alphapb.FilestoreAlphaBackup {
	p := &alphapb.FilestoreAlphaBackup{
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		State:              FilestoreAlphaBackupStateEnumToProto(resource.State),
		CreateTime:         dcl.ValueOrEmptyString(resource.CreateTime),
		CapacityGb:         dcl.ValueOrEmptyInt64(resource.CapacityGb),
		StorageBytes:       dcl.ValueOrEmptyInt64(resource.StorageBytes),
		SourceInstance:     dcl.ValueOrEmptyString(resource.SourceInstance),
		SourceFileShare:    dcl.ValueOrEmptyString(resource.SourceFileShare),
		SourceInstanceTier: FilestoreAlphaBackupSourceInstanceTierEnumToProto(resource.SourceInstanceTier),
		DownloadBytes:      dcl.ValueOrEmptyInt64(resource.DownloadBytes),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		Location:           dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyBackup handles the gRPC request by passing it to the underlying Backup Apply() method.
func (s *BackupServer) applyBackup(ctx context.Context, c *alpha.Client, request *alphapb.ApplyFilestoreAlphaBackupRequest) (*alphapb.FilestoreAlphaBackup, error) {
	p := ProtoToBackup(request.GetResource())
	res, err := c.ApplyBackup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BackupToProto(res)
	return r, nil
}

// ApplyBackup handles the gRPC request by passing it to the underlying Backup Apply() method.
func (s *BackupServer) ApplyFilestoreAlphaBackup(ctx context.Context, request *alphapb.ApplyFilestoreAlphaBackupRequest) (*alphapb.FilestoreAlphaBackup, error) {
	cl, err := createConfigBackup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyBackup(ctx, cl, request)
}

// DeleteBackup handles the gRPC request by passing it to the underlying Backup Delete() method.
func (s *BackupServer) DeleteFilestoreAlphaBackup(ctx context.Context, request *alphapb.DeleteFilestoreAlphaBackupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBackup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBackup(ctx, ProtoToBackup(request.GetResource()))

}

// ListFilestoreAlphaBackup handles the gRPC request by passing it to the underlying BackupList() method.
func (s *BackupServer) ListFilestoreAlphaBackup(ctx context.Context, request *alphapb.ListFilestoreAlphaBackupRequest) (*alphapb.ListFilestoreAlphaBackupResponse, error) {
	cl, err := createConfigBackup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBackup(ctx, ProtoToBackup(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.FilestoreAlphaBackup
	for _, r := range resources.Items {
		rp := BackupToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListFilestoreAlphaBackupResponse{Items: protos}, nil
}

func createConfigBackup(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
