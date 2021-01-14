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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for Interconnect.
type InterconnectServer struct{}

// ProtoToInterconnectLinkTypeEnum converts a InterconnectLinkTypeEnum enum from its proto representation.
func ProtoToComputeInterconnectLinkTypeEnum(e computepb.ComputeInterconnectLinkTypeEnum) *compute.InterconnectLinkTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectLinkTypeEnum_name[int32(e)]; ok {
		e := compute.InterconnectLinkTypeEnum(n[len("InterconnectLinkTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectInterconnectTypeEnum converts a InterconnectInterconnectTypeEnum enum from its proto representation.
func ProtoToComputeInterconnectInterconnectTypeEnum(e computepb.ComputeInterconnectInterconnectTypeEnum) *compute.InterconnectInterconnectTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectInterconnectTypeEnum_name[int32(e)]; ok {
		e := compute.InterconnectInterconnectTypeEnum(n[len("InterconnectInterconnectTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectOperationalStatusEnum converts a InterconnectOperationalStatusEnum enum from its proto representation.
func ProtoToComputeInterconnectOperationalStatusEnum(e computepb.ComputeInterconnectOperationalStatusEnum) *compute.InterconnectOperationalStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectOperationalStatusEnum_name[int32(e)]; ok {
		e := compute.InterconnectOperationalStatusEnum(n[len("InterconnectOperationalStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutagesSourceEnum converts a InterconnectExpectedOutagesSourceEnum enum from its proto representation.
func ProtoToComputeInterconnectExpectedOutagesSourceEnum(e computepb.ComputeInterconnectExpectedOutagesSourceEnum) *compute.InterconnectExpectedOutagesSourceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectExpectedOutagesSourceEnum_name[int32(e)]; ok {
		e := compute.InterconnectExpectedOutagesSourceEnum(n[len("InterconnectExpectedOutagesSourceEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutagesStateEnum converts a InterconnectExpectedOutagesStateEnum enum from its proto representation.
func ProtoToComputeInterconnectExpectedOutagesStateEnum(e computepb.ComputeInterconnectExpectedOutagesStateEnum) *compute.InterconnectExpectedOutagesStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectExpectedOutagesStateEnum_name[int32(e)]; ok {
		e := compute.InterconnectExpectedOutagesStateEnum(n[len("InterconnectExpectedOutagesStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutagesIssueTypeEnum converts a InterconnectExpectedOutagesIssueTypeEnum enum from its proto representation.
func ProtoToComputeInterconnectExpectedOutagesIssueTypeEnum(e computepb.ComputeInterconnectExpectedOutagesIssueTypeEnum) *compute.InterconnectExpectedOutagesIssueTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectExpectedOutagesIssueTypeEnum_name[int32(e)]; ok {
		e := compute.InterconnectExpectedOutagesIssueTypeEnum(n[len("InterconnectExpectedOutagesIssueTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectStateEnum converts a InterconnectStateEnum enum from its proto representation.
func ProtoToComputeInterconnectStateEnum(e computepb.ComputeInterconnectStateEnum) *compute.InterconnectStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInterconnectStateEnum_name[int32(e)]; ok {
		e := compute.InterconnectStateEnum(n[len("InterconnectStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInterconnectExpectedOutages converts a InterconnectExpectedOutages resource from its proto representation.
func ProtoToComputeInterconnectExpectedOutages(p *computepb.ComputeInterconnectExpectedOutages) *compute.InterconnectExpectedOutages {
	if p == nil {
		return nil
	}
	obj := &compute.InterconnectExpectedOutages{
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		Source:      ProtoToComputeInterconnectExpectedOutagesSourceEnum(p.GetSource()),
		State:       ProtoToComputeInterconnectExpectedOutagesStateEnum(p.GetState()),
		IssueType:   ProtoToComputeInterconnectExpectedOutagesIssueTypeEnum(p.GetIssueType()),
		StartTime:   dcl.Int64OrNil(p.StartTime),
		EndTime:     dcl.Int64OrNil(p.EndTime),
	}
	for _, r := range p.GetAffectedCircuits() {
		obj.AffectedCircuits = append(obj.AffectedCircuits, r)
	}
	return obj
}

// ProtoToInterconnectCircuitInfos converts a InterconnectCircuitInfos resource from its proto representation.
func ProtoToComputeInterconnectCircuitInfos(p *computepb.ComputeInterconnectCircuitInfos) *compute.InterconnectCircuitInfos {
	if p == nil {
		return nil
	}
	obj := &compute.InterconnectCircuitInfos{
		GoogleCircuitId:  dcl.StringOrNil(p.GoogleCircuitId),
		GoogleDemarcId:   dcl.StringOrNil(p.GoogleDemarcId),
		CustomerDemarcId: dcl.StringOrNil(p.CustomerDemarcId),
	}
	return obj
}

// ProtoToInterconnect converts a Interconnect resource from its proto representation.
func ProtoToInterconnect(p *computepb.ComputeInterconnect) *compute.Interconnect {
	obj := &compute.Interconnect{
		Description:          dcl.StringOrNil(p.Description),
		SelfLink:             dcl.StringOrNil(p.SelfLink),
		Id:                   dcl.Int64OrNil(p.Id),
		Name:                 dcl.StringOrNil(p.Name),
		Location:             dcl.StringOrNil(p.Location),
		LinkType:             ProtoToComputeInterconnectLinkTypeEnum(p.GetLinkType()),
		RequestedLinkCount:   dcl.Int64OrNil(p.RequestedLinkCount),
		InterconnectType:     ProtoToComputeInterconnectInterconnectTypeEnum(p.GetInterconnectType()),
		AdminEnabled:         dcl.Bool(p.AdminEnabled),
		NocContactEmail:      dcl.StringOrNil(p.NocContactEmail),
		CustomerName:         dcl.StringOrNil(p.CustomerName),
		OperationalStatus:    ProtoToComputeInterconnectOperationalStatusEnum(p.GetOperationalStatus()),
		ProvisionedLinkCount: dcl.Int64OrNil(p.ProvisionedLinkCount),
		PeerIPAddress:        dcl.StringOrNil(p.PeerIpAddress),
		GoogleIPAddress:      dcl.StringOrNil(p.GoogleIpAddress),
		GoogleReferenceId:    dcl.StringOrNil(p.GoogleReferenceId),
		State:                ProtoToComputeInterconnectStateEnum(p.GetState()),
		Project:              dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetInterconnectAttachments() {
		obj.InterconnectAttachments = append(obj.InterconnectAttachments, r)
	}
	for _, r := range p.GetExpectedOutages() {
		obj.ExpectedOutages = append(obj.ExpectedOutages, *ProtoToComputeInterconnectExpectedOutages(r))
	}
	for _, r := range p.GetCircuitInfos() {
		obj.CircuitInfos = append(obj.CircuitInfos, *ProtoToComputeInterconnectCircuitInfos(r))
	}
	return obj
}

// InterconnectLinkTypeEnumToProto converts a InterconnectLinkTypeEnum enum to its proto representation.
func ComputeInterconnectLinkTypeEnumToProto(e *compute.InterconnectLinkTypeEnum) computepb.ComputeInterconnectLinkTypeEnum {
	if e == nil {
		return computepb.ComputeInterconnectLinkTypeEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectLinkTypeEnum_value["InterconnectLinkTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectLinkTypeEnum(v)
	}
	return computepb.ComputeInterconnectLinkTypeEnum(0)
}

// InterconnectInterconnectTypeEnumToProto converts a InterconnectInterconnectTypeEnum enum to its proto representation.
func ComputeInterconnectInterconnectTypeEnumToProto(e *compute.InterconnectInterconnectTypeEnum) computepb.ComputeInterconnectInterconnectTypeEnum {
	if e == nil {
		return computepb.ComputeInterconnectInterconnectTypeEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectInterconnectTypeEnum_value["InterconnectInterconnectTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectInterconnectTypeEnum(v)
	}
	return computepb.ComputeInterconnectInterconnectTypeEnum(0)
}

// InterconnectOperationalStatusEnumToProto converts a InterconnectOperationalStatusEnum enum to its proto representation.
func ComputeInterconnectOperationalStatusEnumToProto(e *compute.InterconnectOperationalStatusEnum) computepb.ComputeInterconnectOperationalStatusEnum {
	if e == nil {
		return computepb.ComputeInterconnectOperationalStatusEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectOperationalStatusEnum_value["InterconnectOperationalStatusEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectOperationalStatusEnum(v)
	}
	return computepb.ComputeInterconnectOperationalStatusEnum(0)
}

// InterconnectExpectedOutagesSourceEnumToProto converts a InterconnectExpectedOutagesSourceEnum enum to its proto representation.
func ComputeInterconnectExpectedOutagesSourceEnumToProto(e *compute.InterconnectExpectedOutagesSourceEnum) computepb.ComputeInterconnectExpectedOutagesSourceEnum {
	if e == nil {
		return computepb.ComputeInterconnectExpectedOutagesSourceEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectExpectedOutagesSourceEnum_value["InterconnectExpectedOutagesSourceEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectExpectedOutagesSourceEnum(v)
	}
	return computepb.ComputeInterconnectExpectedOutagesSourceEnum(0)
}

// InterconnectExpectedOutagesStateEnumToProto converts a InterconnectExpectedOutagesStateEnum enum to its proto representation.
func ComputeInterconnectExpectedOutagesStateEnumToProto(e *compute.InterconnectExpectedOutagesStateEnum) computepb.ComputeInterconnectExpectedOutagesStateEnum {
	if e == nil {
		return computepb.ComputeInterconnectExpectedOutagesStateEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectExpectedOutagesStateEnum_value["InterconnectExpectedOutagesStateEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectExpectedOutagesStateEnum(v)
	}
	return computepb.ComputeInterconnectExpectedOutagesStateEnum(0)
}

// InterconnectExpectedOutagesIssueTypeEnumToProto converts a InterconnectExpectedOutagesIssueTypeEnum enum to its proto representation.
func ComputeInterconnectExpectedOutagesIssueTypeEnumToProto(e *compute.InterconnectExpectedOutagesIssueTypeEnum) computepb.ComputeInterconnectExpectedOutagesIssueTypeEnum {
	if e == nil {
		return computepb.ComputeInterconnectExpectedOutagesIssueTypeEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectExpectedOutagesIssueTypeEnum_value["InterconnectExpectedOutagesIssueTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectExpectedOutagesIssueTypeEnum(v)
	}
	return computepb.ComputeInterconnectExpectedOutagesIssueTypeEnum(0)
}

// InterconnectStateEnumToProto converts a InterconnectStateEnum enum to its proto representation.
func ComputeInterconnectStateEnumToProto(e *compute.InterconnectStateEnum) computepb.ComputeInterconnectStateEnum {
	if e == nil {
		return computepb.ComputeInterconnectStateEnum(0)
	}
	if v, ok := computepb.ComputeInterconnectStateEnum_value["InterconnectStateEnum"+string(*e)]; ok {
		return computepb.ComputeInterconnectStateEnum(v)
	}
	return computepb.ComputeInterconnectStateEnum(0)
}

// InterconnectExpectedOutagesToProto converts a InterconnectExpectedOutages resource to its proto representation.
func ComputeInterconnectExpectedOutagesToProto(o *compute.InterconnectExpectedOutages) *computepb.ComputeInterconnectExpectedOutages {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInterconnectExpectedOutages{
		Name:        dcl.ValueOrEmptyString(o.Name),
		Description: dcl.ValueOrEmptyString(o.Description),
		Source:      ComputeInterconnectExpectedOutagesSourceEnumToProto(o.Source),
		State:       ComputeInterconnectExpectedOutagesStateEnumToProto(o.State),
		IssueType:   ComputeInterconnectExpectedOutagesIssueTypeEnumToProto(o.IssueType),
		StartTime:   dcl.ValueOrEmptyInt64(o.StartTime),
		EndTime:     dcl.ValueOrEmptyInt64(o.EndTime),
	}
	for _, r := range o.AffectedCircuits {
		p.AffectedCircuits = append(p.AffectedCircuits, r)
	}
	return p
}

// InterconnectCircuitInfosToProto converts a InterconnectCircuitInfos resource to its proto representation.
func ComputeInterconnectCircuitInfosToProto(o *compute.InterconnectCircuitInfos) *computepb.ComputeInterconnectCircuitInfos {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInterconnectCircuitInfos{
		GoogleCircuitId:  dcl.ValueOrEmptyString(o.GoogleCircuitId),
		GoogleDemarcId:   dcl.ValueOrEmptyString(o.GoogleDemarcId),
		CustomerDemarcId: dcl.ValueOrEmptyString(o.CustomerDemarcId),
	}
	return p
}

// InterconnectToProto converts a Interconnect resource to its proto representation.
func InterconnectToProto(resource *compute.Interconnect) *computepb.ComputeInterconnect {
	p := &computepb.ComputeInterconnect{
		Description:          dcl.ValueOrEmptyString(resource.Description),
		SelfLink:             dcl.ValueOrEmptyString(resource.SelfLink),
		Id:                   dcl.ValueOrEmptyInt64(resource.Id),
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		Location:             dcl.ValueOrEmptyString(resource.Location),
		LinkType:             ComputeInterconnectLinkTypeEnumToProto(resource.LinkType),
		RequestedLinkCount:   dcl.ValueOrEmptyInt64(resource.RequestedLinkCount),
		InterconnectType:     ComputeInterconnectInterconnectTypeEnumToProto(resource.InterconnectType),
		AdminEnabled:         dcl.ValueOrEmptyBool(resource.AdminEnabled),
		NocContactEmail:      dcl.ValueOrEmptyString(resource.NocContactEmail),
		CustomerName:         dcl.ValueOrEmptyString(resource.CustomerName),
		OperationalStatus:    ComputeInterconnectOperationalStatusEnumToProto(resource.OperationalStatus),
		ProvisionedLinkCount: dcl.ValueOrEmptyInt64(resource.ProvisionedLinkCount),
		PeerIpAddress:        dcl.ValueOrEmptyString(resource.PeerIPAddress),
		GoogleIpAddress:      dcl.ValueOrEmptyString(resource.GoogleIPAddress),
		GoogleReferenceId:    dcl.ValueOrEmptyString(resource.GoogleReferenceId),
		State:                ComputeInterconnectStateEnumToProto(resource.State),
		Project:              dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.InterconnectAttachments {
		p.InterconnectAttachments = append(p.InterconnectAttachments, r)
	}
	for _, r := range resource.ExpectedOutages {
		p.ExpectedOutages = append(p.ExpectedOutages, ComputeInterconnectExpectedOutagesToProto(&r))
	}
	for _, r := range resource.CircuitInfos {
		p.CircuitInfos = append(p.CircuitInfos, ComputeInterconnectCircuitInfosToProto(&r))
	}

	return p
}

// ApplyInterconnect handles the gRPC request by passing it to the underlying Interconnect Apply() method.
func (s *InterconnectServer) applyInterconnect(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeInterconnectRequest) (*computepb.ComputeInterconnect, error) {
	p := ProtoToInterconnect(request.GetResource())
	res, err := c.ApplyInterconnect(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InterconnectToProto(res)
	return r, nil
}

// ApplyInterconnect handles the gRPC request by passing it to the underlying Interconnect Apply() method.
func (s *InterconnectServer) ApplyComputeInterconnect(ctx context.Context, request *computepb.ApplyComputeInterconnectRequest) (*computepb.ComputeInterconnect, error) {
	cl, err := createConfigInterconnect(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInterconnect(ctx, cl, request)
}

// DeleteInterconnect handles the gRPC request by passing it to the underlying Interconnect Delete() method.
func (s *InterconnectServer) DeleteComputeInterconnect(ctx context.Context, request *computepb.DeleteComputeInterconnectRequest) (*emptypb.Empty, error) {
	cl, err := createConfigInterconnect(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInterconnect(ctx, ProtoToInterconnect(request.GetResource()))
}

// ListInterconnect handles the gRPC request by passing it to the underlying InterconnectList() method.
func (s *InterconnectServer) ListComputeInterconnect(ctx context.Context, request *computepb.ListComputeInterconnectRequest) (*computepb.ListComputeInterconnectResponse, error) {
	cl, err := createConfigInterconnect(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInterconnect(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeInterconnect
	for _, r := range resources.Items {
		rp := InterconnectToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeInterconnectResponse{Items: protos}, nil
}

func createConfigInterconnect(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
