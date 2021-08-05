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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/binaryauthorization/beta/binaryauthorization_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/beta"
)

// Server implements the gRPC interface for Attestor.
type AttestorServer struct{}

// ProtoToAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum enum from its proto representation.
func ProtoToBinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(e betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) *beta.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum_name[int32(e)]; ok {
		e := beta.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(n[len("BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToAttestorUserOwnedDrydockNote converts a AttestorUserOwnedDrydockNote resource from its proto representation.
func ProtoToBinaryauthorizationBetaAttestorUserOwnedDrydockNote(p *betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNote) *beta.AttestorUserOwnedDrydockNote {
	if p == nil {
		return nil
	}
	obj := &beta.AttestorUserOwnedDrydockNote{
		NoteReference:                 dcl.StringOrNil(p.NoteReference),
		DelegationServiceAccountEmail: dcl.StringOrNil(p.DelegationServiceAccountEmail),
	}
	for _, r := range p.GetPublicKeys() {
		obj.PublicKeys = append(obj.PublicKeys, *ProtoToBinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeys(r))
	}
	return obj
}

// ProtoToAttestorUserOwnedDrydockNotePublicKeys converts a AttestorUserOwnedDrydockNotePublicKeys resource from its proto representation.
func ProtoToBinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeys(p *betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeys) *beta.AttestorUserOwnedDrydockNotePublicKeys {
	if p == nil {
		return nil
	}
	obj := &beta.AttestorUserOwnedDrydockNotePublicKeys{
		Comment:                  dcl.StringOrNil(p.Comment),
		Id:                       dcl.StringOrNil(p.Id),
		AsciiArmoredPgpPublicKey: dcl.StringOrNil(p.AsciiArmoredPgpPublicKey),
		PkixPublicKey:            ProtoToBinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(p.GetPkixPublicKey()),
	}
	return obj
}

// ProtoToAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey resource from its proto representation.
func ProtoToBinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(p *betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) *beta.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if p == nil {
		return nil
	}
	obj := &beta.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{
		PublicKeyPem:       dcl.StringOrNil(p.PublicKeyPem),
		SignatureAlgorithm: ProtoToBinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(p.GetSignatureAlgorithm()),
	}
	return obj
}

// ProtoToAttestor converts a Attestor resource from its proto representation.
func ProtoToAttestor(p *betapb.BinaryauthorizationBetaAttestor) *beta.Attestor {
	obj := &beta.Attestor{
		Name:                 dcl.StringOrNil(p.Name),
		Description:          dcl.StringOrNil(p.Description),
		UserOwnedDrydockNote: ProtoToBinaryauthorizationBetaAttestorUserOwnedDrydockNote(p.GetUserOwnedDrydockNote()),
		UpdateTime:           dcl.StringOrNil(p.GetUpdateTime()),
		Project:              dcl.StringOrNil(p.Project),
	}
	return obj
}

// AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum enum to its proto representation.
func BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto(e *beta.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	if e == nil {
		return betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(0)
	}
	if v, ok := betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum_value["AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum"+string(*e)]; ok {
		return betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(v)
	}
	return betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(0)
}

// AttestorUserOwnedDrydockNoteToProto converts a AttestorUserOwnedDrydockNote resource to its proto representation.
func BinaryauthorizationBetaAttestorUserOwnedDrydockNoteToProto(o *beta.AttestorUserOwnedDrydockNote) *betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNote {
	if o == nil {
		return nil
	}
	p := &betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNote{
		NoteReference:                 dcl.ValueOrEmptyString(o.NoteReference),
		DelegationServiceAccountEmail: dcl.ValueOrEmptyString(o.DelegationServiceAccountEmail),
	}
	for _, r := range o.PublicKeys {
		p.PublicKeys = append(p.PublicKeys, BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysToProto(&r))
	}
	return p
}

// AttestorUserOwnedDrydockNotePublicKeysToProto converts a AttestorUserOwnedDrydockNotePublicKeys resource to its proto representation.
func BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysToProto(o *beta.AttestorUserOwnedDrydockNotePublicKeys) *betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeys {
	if o == nil {
		return nil
	}
	p := &betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeys{
		Comment:                  dcl.ValueOrEmptyString(o.Comment),
		Id:                       dcl.ValueOrEmptyString(o.Id),
		AsciiArmoredPgpPublicKey: dcl.ValueOrEmptyString(o.AsciiArmoredPgpPublicKey),
		PkixPublicKey:            BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto(o.PkixPublicKey),
	}
	return p
}

// AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey resource to its proto representation.
func BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto(o *beta.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) *betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if o == nil {
		return nil
	}
	p := &betapb.BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{
		PublicKeyPem:       dcl.ValueOrEmptyString(o.PublicKeyPem),
		SignatureAlgorithm: BinaryauthorizationBetaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto(o.SignatureAlgorithm),
	}
	return p
}

// AttestorToProto converts a Attestor resource to its proto representation.
func AttestorToProto(resource *beta.Attestor) *betapb.BinaryauthorizationBetaAttestor {
	p := &betapb.BinaryauthorizationBetaAttestor{
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		Description:          dcl.ValueOrEmptyString(resource.Description),
		UserOwnedDrydockNote: BinaryauthorizationBetaAttestorUserOwnedDrydockNoteToProto(resource.UserOwnedDrydockNote),
		UpdateTime:           dcl.ValueOrEmptyString(resource.UpdateTime),
		Project:              dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyAttestor handles the gRPC request by passing it to the underlying Attestor Apply() method.
func (s *AttestorServer) applyAttestor(ctx context.Context, c *beta.Client, request *betapb.ApplyBinaryauthorizationBetaAttestorRequest) (*betapb.BinaryauthorizationBetaAttestor, error) {
	p := ProtoToAttestor(request.GetResource())
	res, err := c.ApplyAttestor(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AttestorToProto(res)
	return r, nil
}

// ApplyAttestor handles the gRPC request by passing it to the underlying Attestor Apply() method.
func (s *AttestorServer) ApplyBinaryauthorizationBetaAttestor(ctx context.Context, request *betapb.ApplyBinaryauthorizationBetaAttestorRequest) (*betapb.BinaryauthorizationBetaAttestor, error) {
	cl, err := createConfigAttestor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAttestor(ctx, cl, request)
}

// DeleteAttestor handles the gRPC request by passing it to the underlying Attestor Delete() method.
func (s *AttestorServer) DeleteBinaryauthorizationBetaAttestor(ctx context.Context, request *betapb.DeleteBinaryauthorizationBetaAttestorRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAttestor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAttestor(ctx, ProtoToAttestor(request.GetResource()))

}

// ListBinaryauthorizationBetaAttestor handles the gRPC request by passing it to the underlying AttestorList() method.
func (s *AttestorServer) ListBinaryauthorizationBetaAttestor(ctx context.Context, request *betapb.ListBinaryauthorizationBetaAttestorRequest) (*betapb.ListBinaryauthorizationBetaAttestorResponse, error) {
	cl, err := createConfigAttestor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAttestor(ctx, ProtoToAttestor(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.BinaryauthorizationBetaAttestor
	for _, r := range resources.Items {
		rp := AttestorToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListBinaryauthorizationBetaAttestorResponse{Items: protos}, nil
}

func createConfigAttestor(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
