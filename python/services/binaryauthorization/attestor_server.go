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

// ProtoToAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum converts a AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum enum from its proto representation.
func ProtoToBinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(e binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) *binaryauthorization.AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum_name[int32(e)]; ok {
		e := binaryauthorization.AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(n[len("BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToAttestorUserOwnedGrafeasNote converts a AttestorUserOwnedGrafeasNote resource from its proto representation.
func ProtoToBinaryauthorizationAttestorUserOwnedGrafeasNote(p *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNote) *binaryauthorization.AttestorUserOwnedGrafeasNote {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.AttestorUserOwnedGrafeasNote{
		NoteReference:                 dcl.StringOrNil(p.NoteReference),
		DelegationServiceAccountEmail: dcl.StringOrNil(p.DelegationServiceAccountEmail),
	}
	for _, r := range p.GetPublicKeys() {
		obj.PublicKeys = append(obj.PublicKeys, *ProtoToBinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeys(r))
	}
	return obj
}

// ProtoToAttestorUserOwnedGrafeasNotePublicKeys converts a AttestorUserOwnedGrafeasNotePublicKeys resource from its proto representation.
func ProtoToBinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeys(p *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeys) *binaryauthorization.AttestorUserOwnedGrafeasNotePublicKeys {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.AttestorUserOwnedGrafeasNotePublicKeys{
		Comment:                  dcl.StringOrNil(p.Comment),
		Id:                       dcl.StringOrNil(p.Id),
		AsciiArmoredPgpPublicKey: dcl.StringOrNil(p.AsciiArmoredPgpPublicKey),
		PkixPublicKey:            ProtoToBinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(p.GetPkixPublicKey()),
	}
	return obj
}

// ProtoToAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey converts a AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey resource from its proto representation.
func ProtoToBinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey(p *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey) *binaryauthorization.AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey {
	if p == nil {
		return nil
	}
	obj := &binaryauthorization.AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey{
		PublicKeyPem:       dcl.StringOrNil(p.PublicKeyPem),
		SignatureAlgorithm: ProtoToBinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(p.GetSignatureAlgorithm()),
	}
	return obj
}

// ProtoToAttestor converts a Attestor resource from its proto representation.
func ProtoToAttestor(p *binaryauthorizationpb.BinaryauthorizationAttestor) *binaryauthorization.Attestor {
	obj := &binaryauthorization.Attestor{
		Name:                 dcl.StringOrNil(p.Name),
		Description:          dcl.StringOrNil(p.Description),
		UserOwnedGrafeasNote: ProtoToBinaryauthorizationAttestorUserOwnedGrafeasNote(p.GetUserOwnedGrafeasNote()),
		UpdateTime:           dcl.StringOrNil(p.GetUpdateTime()),
		Project:              dcl.StringOrNil(p.Project),
	}
	return obj
}

// AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto converts a AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum enum to its proto representation.
func BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto(e *binaryauthorization.AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	if e == nil {
		return binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(0)
	}
	if v, ok := binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum_value["AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum"+string(*e)]; ok {
		return binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(v)
	}
	return binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(0)
}

// AttestorUserOwnedGrafeasNoteToProto converts a AttestorUserOwnedGrafeasNote resource to its proto representation.
func BinaryauthorizationAttestorUserOwnedGrafeasNoteToProto(o *binaryauthorization.AttestorUserOwnedGrafeasNote) *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNote {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNote{
		NoteReference:                 dcl.ValueOrEmptyString(o.NoteReference),
		DelegationServiceAccountEmail: dcl.ValueOrEmptyString(o.DelegationServiceAccountEmail),
	}
	for _, r := range o.PublicKeys {
		p.PublicKeys = append(p.PublicKeys, BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysToProto(&r))
	}
	return p
}

// AttestorUserOwnedGrafeasNotePublicKeysToProto converts a AttestorUserOwnedGrafeasNotePublicKeys resource to its proto representation.
func BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysToProto(o *binaryauthorization.AttestorUserOwnedGrafeasNotePublicKeys) *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeys {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeys{
		Comment:                  dcl.ValueOrEmptyString(o.Comment),
		Id:                       dcl.ValueOrEmptyString(o.Id),
		AsciiArmoredPgpPublicKey: dcl.ValueOrEmptyString(o.AsciiArmoredPgpPublicKey),
		PkixPublicKey:            BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeyToProto(o.PkixPublicKey),
	}
	return p
}

// AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeyToProto converts a AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey resource to its proto representation.
func BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeyToProto(o *binaryauthorization.AttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey) *binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey {
	if o == nil {
		return nil
	}
	p := &binaryauthorizationpb.BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKey{
		PublicKeyPem:       dcl.ValueOrEmptyString(o.PublicKeyPem),
		SignatureAlgorithm: BinaryauthorizationAttestorUserOwnedGrafeasNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto(o.SignatureAlgorithm),
	}
	return p
}

// AttestorToProto converts a Attestor resource to its proto representation.
func AttestorToProto(resource *binaryauthorization.Attestor) *binaryauthorizationpb.BinaryauthorizationAttestor {
	p := &binaryauthorizationpb.BinaryauthorizationAttestor{
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		Description:          dcl.ValueOrEmptyString(resource.Description),
		UserOwnedGrafeasNote: BinaryauthorizationAttestorUserOwnedGrafeasNoteToProto(resource.UserOwnedGrafeasNote),
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

// ListAttestor handles the gRPC request by passing it to the underlying AttestorList() method.
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
