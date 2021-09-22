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
	binaryauthorizationpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/binaryauthorization/binaryauthorization_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization"
)

// Server implements the gRPC interface for Attestor.
type AttestorServer struct{}

// ProtoToAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum enum from its proto representation.
func ProtoToBinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(e binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) *binaryauthorization.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum_name[int32(e)]; ok {
		e := binaryauthorization.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(n[len("BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToAttestorUserOwnedDrydockNote converts a AttestorUserOwnedDrydockNote resource from its proto representation.
func ProtoToBinaryauthorizationAttestorUserOwnedDrydockNote(p *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNote) *binaryauthorization.AttestorUserOwnedDrydockNote {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.AttestorUserOwnedDrydockNote{
		NoteReference:                 dcl.StringOrNil(p.NoteReference),
		DelegationServiceAccountEmail: dcl.StringOrNil(p.DelegationServiceAccountEmail),
	}
	for _, r := range p.GetPublicKeys() {
		obj.PublicKeys = append(obj.PublicKeys, *ProtoToBinaryauthorizationAttestorUserOwnedDrydockNotePublicKeys(r))
	}
	return obj
}

// ProtoToAttestorUserOwnedDrydockNotePublicKeys converts a AttestorUserOwnedDrydockNotePublicKeys resource from its proto representation.
func ProtoToBinaryauthorizationAttestorUserOwnedDrydockNotePublicKeys(p *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeys) *binaryauthorization.AttestorUserOwnedDrydockNotePublicKeys {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.AttestorUserOwnedDrydockNotePublicKeys{
		Comment:                  dcl.StringOrNil(p.Comment),
		Id:                       dcl.StringOrNil(p.Id),
		AsciiArmoredPgpPublicKey: dcl.StringOrNil(p.AsciiArmoredPgpPublicKey),
		PkixPublicKey:            ProtoToBinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(p.GetPkixPublicKey()),
	}
	return obj
}

// ProtoToAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey resource from its proto representation.
func ProtoToBinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(p *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) *binaryauthorization.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{
		PublicKeyPem:       dcl.StringOrNil(p.PublicKeyPem),
		SignatureAlgorithm: ProtoToBinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(p.GetSignatureAlgorithm()),
	}
	return obj
}

// ProtoToAttestor converts a Attestor resource from its proto representation.
func ProtoToAttestor(p *binaryauthorizationpb.BinaryauthorizationAttestor) *binaryauthorization.Attestor {
	obj := &binaryauthorization.Attestor{
		Name:                 dcl.StringOrNil(p.Name),
		Description:          dcl.StringOrNil(p.Description),
		UserOwnedDrydockNote: ProtoToBinaryauthorizationAttestorUserOwnedDrydockNote(p.GetUserOwnedDrydockNote()),
		UpdateTime:           dcl.StringOrNil(p.GetUpdateTime()),
		Project:              dcl.StringOrNil(p.Project),
	}
	return obj
}

// AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum enum to its proto representation.
func BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto(e *binaryauthorization.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum_value["AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(0)
}

// AttestorUserOwnedDrydockNoteToProto converts a AttestorUserOwnedDrydockNote resource to its proto representation.
func BinaryauthorizationAttestorUserOwnedDrydockNoteToProto(o *binaryauthorization.AttestorUserOwnedDrydockNote) *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNote {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNote{
		NoteReference:                 dcl.ValueOrEmptyString(o.NoteReference),
		DelegationServiceAccountEmail: dcl.ValueOrEmptyString(o.DelegationServiceAccountEmail),
	}
	for _, r := range o.PublicKeys {
		p.PublicKeys = append(p.PublicKeys, BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysToProto(&r))
	}
	return p
}

// AttestorUserOwnedDrydockNotePublicKeysToProto converts a AttestorUserOwnedDrydockNotePublicKeys resource to its proto representation.
func BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysToProto(o *binaryauthorization.AttestorUserOwnedDrydockNotePublicKeys) *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeys {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeys{
		Comment:                  dcl.ValueOrEmptyString(o.Comment),
		Id:                       dcl.ValueOrEmptyString(o.Id),
		AsciiArmoredPgpPublicKey: dcl.ValueOrEmptyString(o.AsciiArmoredPgpPublicKey),
		PkixPublicKey:            BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto(o.PkixPublicKey),
	}
	return p
}

// AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey resource to its proto representation.
func BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto(o *binaryauthorization.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{
		PublicKeyPem:       dcl.ValueOrEmptyString(o.PublicKeyPem),
		SignatureAlgorithm: BinaryauthorizationAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto(o.SignatureAlgorithm),
	}
	return p
}

// AttestorToProto converts a Attestor resource to its proto representation.
func AttestorToProto(resource *binaryauthorization.Attestor) *binaryauthorizationpb.BinaryauthorizationAttestor {
	p := &binaryauthorizationpb.BinaryauthorizationAttestor{
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		Description:          dcl.ValueOrEmptyString(resource.Description),
		UserOwnedDrydockNote: BinaryauthorizationAttestorUserOwnedDrydockNoteToProto(resource.UserOwnedDrydockNote),
		UpdateTime:           dcl.ValueOrEmptyString(resource.UpdateTime),
		Project:              dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyAttestor handles the gRPC request by passing it to the underlying Attestor Apply() method.
func (s *AttestorServer) applyAttestor(ctx context.Context, c *binaryauthorization.Client, request *binaryauthorizationpb.ApplyBinaryauthorizationAttestorRequest) (*binaryauthorizationpb.BinaryauthorizationAttestor, error) {
	p := ProtoToAttestor(request.GetResource())
	res, err := c.ApplyAttestor(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AttestorToProto(res)
	return r, nil
}

// ApplyAttestor handles the gRPC request by passing it to the underlying Attestor Apply() method.
func (s *AttestorServer) ApplyBinaryauthorizationAttestor(ctx context.Context, request *binaryauthorizationpb.ApplyBinaryauthorizationAttestorRequest) (*binaryauthorizationpb.BinaryauthorizationAttestor, error) {
	cl, err := createConfigAttestor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAttestor(ctx, cl, request)
}

// DeleteAttestor handles the gRPC request by passing it to the underlying Attestor Delete() method.
func (s *AttestorServer) DeleteBinaryauthorizationAttestor(ctx context.Context, request *binaryauthorizationpb.DeleteBinaryauthorizationAttestorRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAttestor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAttestor(ctx, ProtoToAttestor(request.GetResource()))

}

// ListBinaryauthorizationAttestor handles the gRPC request by passing it to the underlying AttestorList() method.
func (s *AttestorServer) ListBinaryauthorizationAttestor(ctx context.Context, request *binaryauthorizationpb.ListBinaryauthorizationAttestorRequest) (*binaryauthorizationpb.ListBinaryauthorizationAttestorResponse, error) {
	cl, err := createConfigAttestor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAttestor(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*binaryauthorizationpb.BinaryauthorizationAttestor
	for _, r := range resources.Items {
		rp := AttestorToProto(r)
		protos = append(protos, rp)
	}
	return &binaryauthorizationpb.ListBinaryauthorizationAttestorResponse{Items: protos}, nil
}

func createConfigAttestor(ctx context.Context, service_account_file string) (*binaryauthorization.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return binaryauthorization.NewClient(conf), nil
}
