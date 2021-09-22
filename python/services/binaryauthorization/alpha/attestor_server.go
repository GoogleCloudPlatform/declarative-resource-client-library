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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/binaryauthorization/alpha/binaryauthorization_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/alpha"
)

// Server implements the gRPC interface for Attestor.
type AttestorServer struct{}

// ProtoToAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum enum from its proto representation.
func ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(e alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) *alpha.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum_name[int32(e)]; ok {
		e := alpha.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(n[len("BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToAttestorUserOwnedDrydockNote converts a AttestorUserOwnedDrydockNote resource from its proto representation.
func ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNote(p *alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNote) *alpha.AttestorUserOwnedDrydockNote {
	if p == nil {
		return nil
	}
	obj := &alpha.AttestorUserOwnedDrydockNote{
		NoteReference:                 dcl.StringOrNil(p.NoteReference),
		DelegationServiceAccountEmail: dcl.StringOrNil(p.DelegationServiceAccountEmail),
	}
	for _, r := range p.GetPublicKeys() {
		obj.PublicKeys = append(obj.PublicKeys, *ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeys(r))
	}
	return obj
}

// ProtoToAttestorUserOwnedDrydockNotePublicKeys converts a AttestorUserOwnedDrydockNotePublicKeys resource from its proto representation.
func ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeys(p *alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeys) *alpha.AttestorUserOwnedDrydockNotePublicKeys {
	if p == nil {
		return nil
	}
	obj := &alpha.AttestorUserOwnedDrydockNotePublicKeys{
		Comment:                  dcl.StringOrNil(p.Comment),
		Id:                       dcl.StringOrNil(p.Id),
		AsciiArmoredPgpPublicKey: dcl.StringOrNil(p.AsciiArmoredPgpPublicKey),
		PkixPublicKey:            ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(p.GetPkixPublicKey()),
	}
	return obj
}

// ProtoToAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey resource from its proto representation.
func ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey(p *alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) *alpha.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if p == nil {
		return nil
	}
	obj := &alpha.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{
		PublicKeyPem:       dcl.StringOrNil(p.PublicKeyPem),
		SignatureAlgorithm: ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(p.GetSignatureAlgorithm()),
	}
	return obj
}

// ProtoToAttestor converts a Attestor resource from its proto representation.
func ProtoToAttestor(p *alphapb.BinaryauthorizationAlphaAttestor) *alpha.Attestor {
	obj := &alpha.Attestor{
		Name:                 dcl.StringOrNil(p.Name),
		Description:          dcl.StringOrNil(p.Description),
		UserOwnedDrydockNote: ProtoToBinaryauthorizationAlphaAttestorUserOwnedDrydockNote(p.GetUserOwnedDrydockNote()),
		UpdateTime:           dcl.StringOrNil(p.GetUpdateTime()),
		Project:              dcl.StringOrNil(p.Project),
	}
	return obj
}

// AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum enum to its proto representation.
func BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto(e *alpha.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum) alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum {
	if e == nil {
		return alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(0)
	}
	if v, ok := alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum_value["AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum"+string(*e)]; ok {
		return alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(v)
	}
	return alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnum(0)
}

// AttestorUserOwnedDrydockNoteToProto converts a AttestorUserOwnedDrydockNote resource to its proto representation.
func BinaryauthorizationAlphaAttestorUserOwnedDrydockNoteToProto(o *alpha.AttestorUserOwnedDrydockNote) *alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNote {
	if o == nil {
		return nil
	}
	p := &alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNote{
		NoteReference:                 dcl.ValueOrEmptyString(o.NoteReference),
		DelegationServiceAccountEmail: dcl.ValueOrEmptyString(o.DelegationServiceAccountEmail),
	}
	for _, r := range o.PublicKeys {
		p.PublicKeys = append(p.PublicKeys, BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysToProto(&r))
	}
	return p
}

// AttestorUserOwnedDrydockNotePublicKeysToProto converts a AttestorUserOwnedDrydockNotePublicKeys resource to its proto representation.
func BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysToProto(o *alpha.AttestorUserOwnedDrydockNotePublicKeys) *alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeys {
	if o == nil {
		return nil
	}
	p := &alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeys{
		Comment:                  dcl.ValueOrEmptyString(o.Comment),
		Id:                       dcl.ValueOrEmptyString(o.Id),
		AsciiArmoredPgpPublicKey: dcl.ValueOrEmptyString(o.AsciiArmoredPgpPublicKey),
		PkixPublicKey:            BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto(o.PkixPublicKey),
	}
	return p
}

// AttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto converts a AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey resource to its proto representation.
func BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeyToProto(o *alpha.AttestorUserOwnedDrydockNotePublicKeysPkixPublicKey) *alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey {
	if o == nil {
		return nil
	}
	p := &alphapb.BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKey{
		PublicKeyPem:       dcl.ValueOrEmptyString(o.PublicKeyPem),
		SignatureAlgorithm: BinaryauthorizationAlphaAttestorUserOwnedDrydockNotePublicKeysPkixPublicKeySignatureAlgorithmEnumToProto(o.SignatureAlgorithm),
	}
	return p
}

// AttestorToProto converts a Attestor resource to its proto representation.
func AttestorToProto(resource *alpha.Attestor) *alphapb.BinaryauthorizationAlphaAttestor {
	p := &alphapb.BinaryauthorizationAlphaAttestor{
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		Description:          dcl.ValueOrEmptyString(resource.Description),
		UserOwnedDrydockNote: BinaryauthorizationAlphaAttestorUserOwnedDrydockNoteToProto(resource.UserOwnedDrydockNote),
		UpdateTime:           dcl.ValueOrEmptyString(resource.UpdateTime),
		Project:              dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyAttestor handles the gRPC request by passing it to the underlying Attestor Apply() method.
func (s *AttestorServer) applyAttestor(ctx context.Context, c *alpha.Client, request *alphapb.ApplyBinaryauthorizationAlphaAttestorRequest) (*alphapb.BinaryauthorizationAlphaAttestor, error) {
	p := ProtoToAttestor(request.GetResource())
	res, err := c.ApplyAttestor(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AttestorToProto(res)
	return r, nil
}

// ApplyAttestor handles the gRPC request by passing it to the underlying Attestor Apply() method.
func (s *AttestorServer) ApplyBinaryauthorizationAlphaAttestor(ctx context.Context, request *alphapb.ApplyBinaryauthorizationAlphaAttestorRequest) (*alphapb.BinaryauthorizationAlphaAttestor, error) {
	cl, err := createConfigAttestor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAttestor(ctx, cl, request)
}

// DeleteAttestor handles the gRPC request by passing it to the underlying Attestor Delete() method.
func (s *AttestorServer) DeleteBinaryauthorizationAlphaAttestor(ctx context.Context, request *alphapb.DeleteBinaryauthorizationAlphaAttestorRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAttestor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAttestor(ctx, ProtoToAttestor(request.GetResource()))

}

// ListBinaryauthorizationAlphaAttestor handles the gRPC request by passing it to the underlying AttestorList() method.
func (s *AttestorServer) ListBinaryauthorizationAlphaAttestor(ctx context.Context, request *alphapb.ListBinaryauthorizationAlphaAttestorRequest) (*alphapb.ListBinaryauthorizationAlphaAttestorResponse, error) {
	cl, err := createConfigAttestor(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAttestor(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.BinaryauthorizationAlphaAttestor
	for _, r := range resources.Items {
		rp := AttestorToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListBinaryauthorizationAlphaAttestorResponse{Items: protos}, nil
}

func createConfigAttestor(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
