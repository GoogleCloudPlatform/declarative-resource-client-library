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
	bigqueryreservationpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigqueryreservation/bigqueryreservation_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation"
)

// Server implements the gRPC interface for CapacityCommitment.
type CapacityCommitmentServer struct{}

// ProtoToCapacityCommitmentPlanEnum converts a CapacityCommitmentPlanEnum enum from its proto representation.
func ProtoToBigqueryreservationCapacityCommitmentPlanEnum(e bigqueryreservationpb.BigqueryreservationCapacityCommitmentPlanEnum) *bigqueryreservation.CapacityCommitmentPlanEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigqueryreservationpb.BigqueryreservationCapacityCommitmentPlanEnum_name[int32(e)]; ok {
		e := bigqueryreservation.CapacityCommitmentPlanEnum(n[len("BigqueryreservationCapacityCommitmentPlanEnum"):])
		return &e
	}
	return nil
}

// ProtoToCapacityCommitmentStateEnum converts a CapacityCommitmentStateEnum enum from its proto representation.
func ProtoToBigqueryreservationCapacityCommitmentStateEnum(e bigqueryreservationpb.BigqueryreservationCapacityCommitmentStateEnum) *bigqueryreservation.CapacityCommitmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigqueryreservationpb.BigqueryreservationCapacityCommitmentStateEnum_name[int32(e)]; ok {
		e := bigqueryreservation.CapacityCommitmentStateEnum(n[len("BigqueryreservationCapacityCommitmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCapacityCommitmentRenewalPlanEnum converts a CapacityCommitmentRenewalPlanEnum enum from its proto representation.
func ProtoToBigqueryreservationCapacityCommitmentRenewalPlanEnum(e bigqueryreservationpb.BigqueryreservationCapacityCommitmentRenewalPlanEnum) *bigqueryreservation.CapacityCommitmentRenewalPlanEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigqueryreservationpb.BigqueryreservationCapacityCommitmentRenewalPlanEnum_name[int32(e)]; ok {
		e := bigqueryreservation.CapacityCommitmentRenewalPlanEnum(n[len("BigqueryreservationCapacityCommitmentRenewalPlanEnum"):])
		return &e
	}
	return nil
}

// ProtoToCapacityCommitmentFailureStatus converts a CapacityCommitmentFailureStatus resource from its proto representation.
func ProtoToBigqueryreservationCapacityCommitmentFailureStatus(p *bigqueryreservationpb.BigqueryreservationCapacityCommitmentFailureStatus) *bigqueryreservation.CapacityCommitmentFailureStatus {
	if p == nil {
		return nil
	}
	obj := &bigqueryreservation.CapacityCommitmentFailureStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToBigqueryreservationCapacityCommitmentFailureStatusDetails(r))
	}
	return obj
}

// ProtoToCapacityCommitmentFailureStatusDetails converts a CapacityCommitmentFailureStatusDetails resource from its proto representation.
func ProtoToBigqueryreservationCapacityCommitmentFailureStatusDetails(p *bigqueryreservationpb.BigqueryreservationCapacityCommitmentFailureStatusDetails) *bigqueryreservation.CapacityCommitmentFailureStatusDetails {
	if p == nil {
		return nil
	}
	obj := &bigqueryreservation.CapacityCommitmentFailureStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToCapacityCommitment converts a CapacityCommitment resource from its proto representation.
func ProtoToCapacityCommitment(p *bigqueryreservationpb.BigqueryreservationCapacityCommitment) *bigqueryreservation.CapacityCommitment {
	obj := &bigqueryreservation.CapacityCommitment{
		Name:                dcl.StringOrNil(p.Name),
		SlotCount:           dcl.Int64OrNil(p.SlotCount),
		Plan:                ProtoToBigqueryreservationCapacityCommitmentPlanEnum(p.GetPlan()),
		State:               ProtoToBigqueryreservationCapacityCommitmentStateEnum(p.GetState()),
		CommitmentStartTime: dcl.StringOrNil(p.GetCommitmentStartTime()),
		CommitmentEndTime:   dcl.StringOrNil(p.GetCommitmentEndTime()),
		FailureStatus:       ProtoToBigqueryreservationCapacityCommitmentFailureStatus(p.GetFailureStatus()),
		RenewalPlan:         ProtoToBigqueryreservationCapacityCommitmentRenewalPlanEnum(p.GetRenewalPlan()),
		Project:             dcl.StringOrNil(p.Project),
		Location:            dcl.StringOrNil(p.Location),
	}
	return obj
}

// CapacityCommitmentPlanEnumToProto converts a CapacityCommitmentPlanEnum enum to its proto representation.
func BigqueryreservationCapacityCommitmentPlanEnumToProto(e *bigqueryreservation.CapacityCommitmentPlanEnum) bigqueryreservationpb.BigqueryreservationCapacityCommitmentPlanEnum {
	if e == nil {
		return bigqueryreservationpb.BigqueryreservationCapacityCommitmentPlanEnum(0)
	}
	if v, ok := bigqueryreservationpb.BigqueryreservationCapacityCommitmentPlanEnum_value["CapacityCommitmentPlanEnum"+string(*e)]; ok {
		return bigqueryreservationpb.BigqueryreservationCapacityCommitmentPlanEnum(v)
	}
	return bigqueryreservationpb.BigqueryreservationCapacityCommitmentPlanEnum(0)
}

// CapacityCommitmentStateEnumToProto converts a CapacityCommitmentStateEnum enum to its proto representation.
func BigqueryreservationCapacityCommitmentStateEnumToProto(e *bigqueryreservation.CapacityCommitmentStateEnum) bigqueryreservationpb.BigqueryreservationCapacityCommitmentStateEnum {
	if e == nil {
		return bigqueryreservationpb.BigqueryreservationCapacityCommitmentStateEnum(0)
	}
	if v, ok := bigqueryreservationpb.BigqueryreservationCapacityCommitmentStateEnum_value["CapacityCommitmentStateEnum"+string(*e)]; ok {
		return bigqueryreservationpb.BigqueryreservationCapacityCommitmentStateEnum(v)
	}
	return bigqueryreservationpb.BigqueryreservationCapacityCommitmentStateEnum(0)
}

// CapacityCommitmentRenewalPlanEnumToProto converts a CapacityCommitmentRenewalPlanEnum enum to its proto representation.
func BigqueryreservationCapacityCommitmentRenewalPlanEnumToProto(e *bigqueryreservation.CapacityCommitmentRenewalPlanEnum) bigqueryreservationpb.BigqueryreservationCapacityCommitmentRenewalPlanEnum {
	if e == nil {
		return bigqueryreservationpb.BigqueryreservationCapacityCommitmentRenewalPlanEnum(0)
	}
	if v, ok := bigqueryreservationpb.BigqueryreservationCapacityCommitmentRenewalPlanEnum_value["CapacityCommitmentRenewalPlanEnum"+string(*e)]; ok {
		return bigqueryreservationpb.BigqueryreservationCapacityCommitmentRenewalPlanEnum(v)
	}
	return bigqueryreservationpb.BigqueryreservationCapacityCommitmentRenewalPlanEnum(0)
}

// CapacityCommitmentFailureStatusToProto converts a CapacityCommitmentFailureStatus resource to its proto representation.
func BigqueryreservationCapacityCommitmentFailureStatusToProto(o *bigqueryreservation.CapacityCommitmentFailureStatus) *bigqueryreservationpb.BigqueryreservationCapacityCommitmentFailureStatus {
	if o == nil {
		return nil
	}
	p := &bigqueryreservationpb.BigqueryreservationCapacityCommitmentFailureStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, BigqueryreservationCapacityCommitmentFailureStatusDetailsToProto(&r))
	}
	return p
}

// CapacityCommitmentFailureStatusDetailsToProto converts a CapacityCommitmentFailureStatusDetails resource to its proto representation.
func BigqueryreservationCapacityCommitmentFailureStatusDetailsToProto(o *bigqueryreservation.CapacityCommitmentFailureStatusDetails) *bigqueryreservationpb.BigqueryreservationCapacityCommitmentFailureStatusDetails {
	if o == nil {
		return nil
	}
	p := &bigqueryreservationpb.BigqueryreservationCapacityCommitmentFailureStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// CapacityCommitmentToProto converts a CapacityCommitment resource to its proto representation.
func CapacityCommitmentToProto(resource *bigqueryreservation.CapacityCommitment) *bigqueryreservationpb.BigqueryreservationCapacityCommitment {
	p := &bigqueryreservationpb.BigqueryreservationCapacityCommitment{
		Name:                dcl.ValueOrEmptyString(resource.Name),
		SlotCount:           dcl.ValueOrEmptyInt64(resource.SlotCount),
		Plan:                BigqueryreservationCapacityCommitmentPlanEnumToProto(resource.Plan),
		State:               BigqueryreservationCapacityCommitmentStateEnumToProto(resource.State),
		CommitmentStartTime: dcl.ValueOrEmptyString(resource.CommitmentStartTime),
		CommitmentEndTime:   dcl.ValueOrEmptyString(resource.CommitmentEndTime),
		FailureStatus:       BigqueryreservationCapacityCommitmentFailureStatusToProto(resource.FailureStatus),
		RenewalPlan:         BigqueryreservationCapacityCommitmentRenewalPlanEnumToProto(resource.RenewalPlan),
		Project:             dcl.ValueOrEmptyString(resource.Project),
		Location:            dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitment Apply() method.
func (s *CapacityCommitmentServer) applyCapacityCommitment(ctx context.Context, c *bigqueryreservation.Client, request *bigqueryreservationpb.ApplyBigqueryreservationCapacityCommitmentRequest) (*bigqueryreservationpb.BigqueryreservationCapacityCommitment, error) {
	p := ProtoToCapacityCommitment(request.GetResource())
	res, err := c.ApplyCapacityCommitment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CapacityCommitmentToProto(res)
	return r, nil
}

// ApplyCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitment Apply() method.
func (s *CapacityCommitmentServer) ApplyBigqueryreservationCapacityCommitment(ctx context.Context, request *bigqueryreservationpb.ApplyBigqueryreservationCapacityCommitmentRequest) (*bigqueryreservationpb.BigqueryreservationCapacityCommitment, error) {
	cl, err := createConfigCapacityCommitment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCapacityCommitment(ctx, cl, request)
}

// DeleteCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitment Delete() method.
func (s *CapacityCommitmentServer) DeleteBigqueryreservationCapacityCommitment(ctx context.Context, request *bigqueryreservationpb.DeleteBigqueryreservationCapacityCommitmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCapacityCommitment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCapacityCommitment(ctx, ProtoToCapacityCommitment(request.GetResource()))

}

// ListBigqueryreservationCapacityCommitment handles the gRPC request by passing it to the underlying CapacityCommitmentList() method.
func (s *CapacityCommitmentServer) ListBigqueryreservationCapacityCommitment(ctx context.Context, request *bigqueryreservationpb.ListBigqueryreservationCapacityCommitmentRequest) (*bigqueryreservationpb.ListBigqueryreservationCapacityCommitmentResponse, error) {
	cl, err := createConfigCapacityCommitment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCapacityCommitment(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*bigqueryreservationpb.BigqueryreservationCapacityCommitment
	for _, r := range resources.Items {
		rp := CapacityCommitmentToProto(r)
		protos = append(protos, rp)
	}
	return &bigqueryreservationpb.ListBigqueryreservationCapacityCommitmentResponse{Items: protos}, nil
}

func createConfigCapacityCommitment(ctx context.Context, service_account_file string) (*bigqueryreservation.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return bigqueryreservation.NewClient(conf), nil
}
