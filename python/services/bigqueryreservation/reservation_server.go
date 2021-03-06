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

// Server implements the gRPC interface for Reservation.
type ReservationServer struct{}

// ProtoToReservation converts a Reservation resource from its proto representation.
func ProtoToReservation(p *bigqueryreservationpb.BigqueryreservationReservation) *bigqueryreservation.Reservation {
	obj := &bigqueryreservation.Reservation{
		Name:            dcl.StringOrNil(p.Name),
		SlotCapacity:    dcl.Int64OrNil(p.SlotCapacity),
		IgnoreIdleSlots: dcl.Bool(p.IgnoreIdleSlots),
		CreationTime:    dcl.StringOrNil(p.GetCreationTime()),
		UpdateTime:      dcl.StringOrNil(p.GetUpdateTime()),
		Project:         dcl.StringOrNil(p.Project),
		Location:        dcl.StringOrNil(p.Location),
	}
	return obj
}

// ReservationToProto converts a Reservation resource to its proto representation.
func ReservationToProto(resource *bigqueryreservation.Reservation) *bigqueryreservationpb.BigqueryreservationReservation {
	p := &bigqueryreservationpb.BigqueryreservationReservation{
		Name:            dcl.ValueOrEmptyString(resource.Name),
		SlotCapacity:    dcl.ValueOrEmptyInt64(resource.SlotCapacity),
		IgnoreIdleSlots: dcl.ValueOrEmptyBool(resource.IgnoreIdleSlots),
		CreationTime:    dcl.ValueOrEmptyString(resource.CreationTime),
		UpdateTime:      dcl.ValueOrEmptyString(resource.UpdateTime),
		Project:         dcl.ValueOrEmptyString(resource.Project),
		Location:        dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyReservation handles the gRPC request by passing it to the underlying Reservation Apply() method.
func (s *ReservationServer) applyReservation(ctx context.Context, c *bigqueryreservation.Client, request *bigqueryreservationpb.ApplyBigqueryreservationReservationRequest) (*bigqueryreservationpb.BigqueryreservationReservation, error) {
	p := ProtoToReservation(request.GetResource())
	res, err := c.ApplyReservation(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ReservationToProto(res)
	return r, nil
}

// ApplyReservation handles the gRPC request by passing it to the underlying Reservation Apply() method.
func (s *ReservationServer) ApplyBigqueryreservationReservation(ctx context.Context, request *bigqueryreservationpb.ApplyBigqueryreservationReservationRequest) (*bigqueryreservationpb.BigqueryreservationReservation, error) {
	cl, err := createConfigReservation(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyReservation(ctx, cl, request)
}

// DeleteReservation handles the gRPC request by passing it to the underlying Reservation Delete() method.
func (s *ReservationServer) DeleteBigqueryreservationReservation(ctx context.Context, request *bigqueryreservationpb.DeleteBigqueryreservationReservationRequest) (*emptypb.Empty, error) {

	cl, err := createConfigReservation(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteReservation(ctx, ProtoToReservation(request.GetResource()))

}

// ListBigqueryreservationReservation handles the gRPC request by passing it to the underlying ReservationList() method.
func (s *ReservationServer) ListBigqueryreservationReservation(ctx context.Context, request *bigqueryreservationpb.ListBigqueryreservationReservationRequest) (*bigqueryreservationpb.ListBigqueryreservationReservationResponse, error) {
	cl, err := createConfigReservation(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListReservation(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*bigqueryreservationpb.BigqueryreservationReservation
	for _, r := range resources.Items {
		rp := ReservationToProto(r)
		protos = append(protos, rp)
	}
	return &bigqueryreservationpb.ListBigqueryreservationReservationResponse{Items: protos}, nil
}

func createConfigReservation(ctx context.Context, service_account_file string) (*bigqueryreservation.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return bigqueryreservation.NewClient(conf), nil
}
