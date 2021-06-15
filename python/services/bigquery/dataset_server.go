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
	bigquerypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigquery/bigquery_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery"
)

// Server implements the gRPC interface for Dataset.
type DatasetServer struct{}

// ProtoToDatasetAccess converts a DatasetAccess resource from its proto representation.
func ProtoToBigqueryDatasetAccess(p *bigquerypb.BigqueryDatasetAccess) *bigquery.DatasetAccess {
	if p == nil {
		return nil
	}
	obj := &bigquery.DatasetAccess{
		Role:         dcl.StringOrNil(p.Role),
		UserByEmail:  dcl.StringOrNil(p.UserByEmail),
		GroupByEmail: dcl.StringOrNil(p.GroupByEmail),
		Domain:       dcl.StringOrNil(p.Domain),
		SpecialGroup: dcl.StringOrNil(p.SpecialGroup),
		IamMember:    dcl.StringOrNil(p.IamMember),
		View:         ProtoToBigqueryDatasetAccessView(p.GetView()),
		Routine:      ProtoToBigqueryDatasetAccessRoutine(p.GetRoutine()),
	}
	return obj
}

// ProtoToDatasetAccessView converts a DatasetAccessView resource from its proto representation.
func ProtoToBigqueryDatasetAccessView(p *bigquerypb.BigqueryDatasetAccessView) *bigquery.DatasetAccessView {
	if p == nil {
		return nil
	}
	obj := &bigquery.DatasetAccessView{
		ProjectId: dcl.StringOrNil(p.ProjectId),
		DatasetId: dcl.StringOrNil(p.DatasetId),
		TableId:   dcl.StringOrNil(p.TableId),
	}
	return obj
}

// ProtoToDatasetAccessRoutine converts a DatasetAccessRoutine resource from its proto representation.
func ProtoToBigqueryDatasetAccessRoutine(p *bigquerypb.BigqueryDatasetAccessRoutine) *bigquery.DatasetAccessRoutine {
	if p == nil {
		return nil
	}
	obj := &bigquery.DatasetAccessRoutine{
		ProjectId: dcl.StringOrNil(p.ProjectId),
		DatasetId: dcl.StringOrNil(p.DatasetId),
		RoutineId: dcl.StringOrNil(p.RoutineId),
	}
	return obj
}

// ProtoToDatasetDefaultEncryptionConfiguration converts a DatasetDefaultEncryptionConfiguration resource from its proto representation.
func ProtoToBigqueryDatasetDefaultEncryptionConfiguration(p *bigquerypb.BigqueryDatasetDefaultEncryptionConfiguration) *bigquery.DatasetDefaultEncryptionConfiguration {
	if p == nil {
		return nil
	}
	obj := &bigquery.DatasetDefaultEncryptionConfiguration{
		KmsKeyName: dcl.StringOrNil(p.KmsKeyName),
	}
	return obj
}

// ProtoToDataset converts a Dataset resource from its proto representation.
func ProtoToDataset(p *bigquerypb.BigqueryDataset) *bigquery.Dataset {
	obj := &bigquery.Dataset{
		Etag:                           dcl.StringOrNil(p.Etag),
		Id:                             dcl.StringOrNil(p.Id),
		SelfLink:                       dcl.StringOrNil(p.SelfLink),
		Name:                           dcl.StringOrNil(p.Name),
		Project:                        dcl.StringOrNil(p.Project),
		FriendlyName:                   dcl.StringOrNil(p.FriendlyName),
		Description:                    dcl.StringOrNil(p.Description),
		DefaultTableExpirationMs:       dcl.StringOrNil(p.DefaultTableExpirationMs),
		DefaultPartitionExpirationMs:   dcl.StringOrNil(p.DefaultPartitionExpirationMs),
		CreationTime:                   dcl.Int64OrNil(p.CreationTime),
		LastModifiedTime:               dcl.Int64OrNil(p.LastModifiedTime),
		Location:                       dcl.StringOrNil(p.Location),
		Published:                      dcl.Bool(p.Published),
		DefaultEncryptionConfiguration: ProtoToBigqueryDatasetDefaultEncryptionConfiguration(p.GetDefaultEncryptionConfiguration()),
	}
	for _, r := range p.GetAccess() {
		obj.Access = append(obj.Access, *ProtoToBigqueryDatasetAccess(r))
	}
	return obj
}

// DatasetAccessToProto converts a DatasetAccess resource to its proto representation.
func BigqueryDatasetAccessToProto(o *bigquery.DatasetAccess) *bigquerypb.BigqueryDatasetAccess {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryDatasetAccess{
		Role:         dcl.ValueOrEmptyString(o.Role),
		UserByEmail:  dcl.ValueOrEmptyString(o.UserByEmail),
		GroupByEmail: dcl.ValueOrEmptyString(o.GroupByEmail),
		Domain:       dcl.ValueOrEmptyString(o.Domain),
		SpecialGroup: dcl.ValueOrEmptyString(o.SpecialGroup),
		IamMember:    dcl.ValueOrEmptyString(o.IamMember),
		View:         BigqueryDatasetAccessViewToProto(o.View),
		Routine:      BigqueryDatasetAccessRoutineToProto(o.Routine),
	}
	return p
}

// DatasetAccessViewToProto converts a DatasetAccessView resource to its proto representation.
func BigqueryDatasetAccessViewToProto(o *bigquery.DatasetAccessView) *bigquerypb.BigqueryDatasetAccessView {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryDatasetAccessView{
		ProjectId: dcl.ValueOrEmptyString(o.ProjectId),
		DatasetId: dcl.ValueOrEmptyString(o.DatasetId),
		TableId:   dcl.ValueOrEmptyString(o.TableId),
	}
	return p
}

// DatasetAccessRoutineToProto converts a DatasetAccessRoutine resource to its proto representation.
func BigqueryDatasetAccessRoutineToProto(o *bigquery.DatasetAccessRoutine) *bigquerypb.BigqueryDatasetAccessRoutine {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryDatasetAccessRoutine{
		ProjectId: dcl.ValueOrEmptyString(o.ProjectId),
		DatasetId: dcl.ValueOrEmptyString(o.DatasetId),
		RoutineId: dcl.ValueOrEmptyString(o.RoutineId),
	}
	return p
}

// DatasetDefaultEncryptionConfigurationToProto converts a DatasetDefaultEncryptionConfiguration resource to its proto representation.
func BigqueryDatasetDefaultEncryptionConfigurationToProto(o *bigquery.DatasetDefaultEncryptionConfiguration) *bigquerypb.BigqueryDatasetDefaultEncryptionConfiguration {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryDatasetDefaultEncryptionConfiguration{
		KmsKeyName: dcl.ValueOrEmptyString(o.KmsKeyName),
	}
	return p
}

// DatasetToProto converts a Dataset resource to its proto representation.
func DatasetToProto(resource *bigquery.Dataset) *bigquerypb.BigqueryDataset {
	p := &bigquerypb.BigqueryDataset{
		Etag:                           dcl.ValueOrEmptyString(resource.Etag),
		Id:                             dcl.ValueOrEmptyString(resource.Id),
		SelfLink:                       dcl.ValueOrEmptyString(resource.SelfLink),
		Name:                           dcl.ValueOrEmptyString(resource.Name),
		Project:                        dcl.ValueOrEmptyString(resource.Project),
		FriendlyName:                   dcl.ValueOrEmptyString(resource.FriendlyName),
		Description:                    dcl.ValueOrEmptyString(resource.Description),
		DefaultTableExpirationMs:       dcl.ValueOrEmptyString(resource.DefaultTableExpirationMs),
		DefaultPartitionExpirationMs:   dcl.ValueOrEmptyString(resource.DefaultPartitionExpirationMs),
		CreationTime:                   dcl.ValueOrEmptyInt64(resource.CreationTime),
		LastModifiedTime:               dcl.ValueOrEmptyInt64(resource.LastModifiedTime),
		Location:                       dcl.ValueOrEmptyString(resource.Location),
		Published:                      dcl.ValueOrEmptyBool(resource.Published),
		DefaultEncryptionConfiguration: BigqueryDatasetDefaultEncryptionConfigurationToProto(resource.DefaultEncryptionConfiguration),
	}
	for _, r := range resource.Access {
		p.Access = append(p.Access, BigqueryDatasetAccessToProto(&r))
	}

	return p
}

// ApplyDataset handles the gRPC request by passing it to the underlying Dataset Apply() method.
func (s *DatasetServer) applyDataset(ctx context.Context, c *bigquery.Client, request *bigquerypb.ApplyBigqueryDatasetRequest) (*bigquerypb.BigqueryDataset, error) {
	p := ProtoToDataset(request.GetResource())
	res, err := c.ApplyDataset(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DatasetToProto(res)
	return r, nil
}

// ApplyDataset handles the gRPC request by passing it to the underlying Dataset Apply() method.
func (s *DatasetServer) ApplyBigqueryDataset(ctx context.Context, request *bigquerypb.ApplyBigqueryDatasetRequest) (*bigquerypb.BigqueryDataset, error) {
	cl, err := createConfigDataset(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyDataset(ctx, cl, request)
}

// DeleteDataset handles the gRPC request by passing it to the underlying Dataset Delete() method.
func (s *DatasetServer) DeleteBigqueryDataset(ctx context.Context, request *bigquerypb.DeleteBigqueryDatasetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDataset(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDataset(ctx, ProtoToDataset(request.GetResource()))

}

// ListBigqueryDataset handles the gRPC request by passing it to the underlying DatasetList() method.
func (s *DatasetServer) ListBigqueryDataset(ctx context.Context, request *bigquerypb.ListBigqueryDatasetRequest) (*bigquerypb.ListBigqueryDatasetResponse, error) {
	cl, err := createConfigDataset(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDataset(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*bigquerypb.BigqueryDataset
	for _, r := range resources.Items {
		rp := DatasetToProto(r)
		protos = append(protos, rp)
	}
	return &bigquerypb.ListBigqueryDatasetResponse{Items: protos}, nil
}

func createConfigDataset(ctx context.Context, service_account_file string) (*bigquery.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return bigquery.NewClient(conf), nil
}
