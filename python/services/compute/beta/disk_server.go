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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for Disk.
type DiskServer struct{}

// ProtoToDiskGuestOsFeatureTypeEnum converts a DiskGuestOsFeatureTypeEnum enum from its proto representation.
func ProtoToComputeBetaDiskGuestOsFeatureTypeEnum(e betapb.ComputeBetaDiskGuestOsFeatureTypeEnum) *beta.DiskGuestOsFeatureTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaDiskGuestOsFeatureTypeEnum_name[int32(e)]; ok {
		e := beta.DiskGuestOsFeatureTypeEnum(n[len("ComputeBetaDiskGuestOsFeatureTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskGuestOsFeatureTypeAltEnum converts a DiskGuestOsFeatureTypeAltEnum enum from its proto representation.
func ProtoToComputeBetaDiskGuestOsFeatureTypeAltEnum(e betapb.ComputeBetaDiskGuestOsFeatureTypeAltEnum) *beta.DiskGuestOsFeatureTypeAltEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaDiskGuestOsFeatureTypeAltEnum_name[int32(e)]; ok {
		e := beta.DiskGuestOsFeatureTypeAltEnum(n[len("ComputeBetaDiskGuestOsFeatureTypeAltEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskStatusEnum converts a DiskStatusEnum enum from its proto representation.
func ProtoToComputeBetaDiskStatusEnum(e betapb.ComputeBetaDiskStatusEnum) *beta.DiskStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaDiskStatusEnum_name[int32(e)]; ok {
		e := beta.DiskStatusEnum(n[len("ComputeBetaDiskStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskGuestOsFeaturesTypeEnum converts a DiskGuestOsFeaturesTypeEnum enum from its proto representation.
func ProtoToComputeBetaDiskGuestOsFeaturesTypeEnum(e betapb.ComputeBetaDiskGuestOsFeaturesTypeEnum) *beta.DiskGuestOsFeaturesTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaDiskGuestOsFeaturesTypeEnum_name[int32(e)]; ok {
		e := beta.DiskGuestOsFeaturesTypeEnum(n[len("ComputeBetaDiskGuestOsFeaturesTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskGuestOsFeaturesTypeAltsEnum converts a DiskGuestOsFeaturesTypeAltsEnum enum from its proto representation.
func ProtoToComputeBetaDiskGuestOsFeaturesTypeAltsEnum(e betapb.ComputeBetaDiskGuestOsFeaturesTypeAltsEnum) *beta.DiskGuestOsFeaturesTypeAltsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaDiskGuestOsFeaturesTypeAltsEnum_name[int32(e)]; ok {
		e := beta.DiskGuestOsFeaturesTypeAltsEnum(n[len("ComputeBetaDiskGuestOsFeaturesTypeAltsEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskGuestOsFeature converts a DiskGuestOsFeature resource from its proto representation.
func ProtoToComputeBetaDiskGuestOsFeature(p *betapb.ComputeBetaDiskGuestOsFeature) *beta.DiskGuestOsFeature {
	if p == nil {
		return nil
	}
	obj := &beta.DiskGuestOsFeature{
		Type: ProtoToComputeBetaDiskGuestOsFeatureTypeEnum(p.GetType()),
	}
	for _, r := range p.GetTypeAlt() {
		obj.TypeAlt = append(obj.TypeAlt, *ProtoToComputeBetaDiskGuestOsFeatureTypeAltEnum(r))
	}
	return obj
}

// ProtoToDiskEncryptionKey converts a DiskEncryptionKey resource from its proto representation.
func ProtoToComputeBetaDiskEncryptionKey(p *betapb.ComputeBetaDiskEncryptionKey) *beta.DiskEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.DiskEncryptionKey{
		RawKey:               dcl.StringOrNil(p.RawKey),
		KmsKeyName:           dcl.StringOrNil(p.KmsKeyName),
		Sha256:               dcl.StringOrNil(p.Sha256),
		KmsKeyServiceAccount: dcl.StringOrNil(p.KmsKeyServiceAccount),
	}
	return obj
}

// ProtoToDiskGuestOsFeatures converts a DiskGuestOsFeatures resource from its proto representation.
func ProtoToComputeBetaDiskGuestOsFeatures(p *betapb.ComputeBetaDiskGuestOsFeatures) *beta.DiskGuestOsFeatures {
	if p == nil {
		return nil
	}
	obj := &beta.DiskGuestOsFeatures{
		Type: ProtoToComputeBetaDiskGuestOsFeaturesTypeEnum(p.GetType()),
	}
	for _, r := range p.GetTypeAlts() {
		obj.TypeAlts = append(obj.TypeAlts, *ProtoToComputeBetaDiskGuestOsFeaturesTypeAltsEnum(r))
	}
	return obj
}

// ProtoToDisk converts a Disk resource from its proto representation.
func ProtoToDisk(p *betapb.ComputeBetaDisk) *beta.Disk {
	obj := &beta.Disk{
		SelfLink:                    dcl.StringOrNil(p.SelfLink),
		Description:                 dcl.StringOrNil(p.Description),
		DiskEncryptionKey:           ProtoToComputeBetaDiskEncryptionKey(p.GetDiskEncryptionKey()),
		LabelFingerprint:            dcl.StringOrNil(p.LabelFingerprint),
		Name:                        dcl.StringOrNil(p.Name),
		Region:                      dcl.StringOrNil(p.Region),
		SizeGb:                      dcl.Int64OrNil(p.SizeGb),
		SourceImage:                 dcl.StringOrNil(p.SourceImage),
		SourceImageEncryptionKey:    ProtoToComputeBetaDiskEncryptionKey(p.GetSourceImageEncryptionKey()),
		SourceImageId:               dcl.StringOrNil(p.SourceImageId),
		SourceSnapshot:              dcl.StringOrNil(p.SourceSnapshot),
		SourceSnapshotEncryptionKey: ProtoToComputeBetaDiskEncryptionKey(p.GetSourceSnapshotEncryptionKey()),
		SourceSnapshotId:            dcl.StringOrNil(p.SourceSnapshotId),
		Type:                        dcl.StringOrNil(p.Type),
		Zone:                        dcl.StringOrNil(p.Zone),
		Project:                     dcl.StringOrNil(p.Project),
		Id:                          dcl.Int64OrNil(p.Id),
		Status:                      ProtoToComputeBetaDiskStatusEnum(p.GetStatus()),
		Options:                     dcl.StringOrNil(p.Options),
		LastAttachTimestamp:         dcl.StringOrNil(p.LastAttachTimestamp),
		LastDetachTimestamp:         dcl.StringOrNil(p.LastDetachTimestamp),
		PhysicalBlockSizeBytes:      dcl.Int64OrNil(p.PhysicalBlockSizeBytes),
		SourceDisk:                  dcl.StringOrNil(p.SourceDisk),
		SourceDiskId:                dcl.StringOrNil(p.SourceDiskId),
		Location:                    dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetGuestOsFeature() {
		obj.GuestOsFeature = append(obj.GuestOsFeature, *ProtoToComputeBetaDiskGuestOsFeature(r))
	}
	for _, r := range p.GetLicense() {
		obj.License = append(obj.License, r)
	}
	for _, r := range p.GetReplicaZones() {
		obj.ReplicaZones = append(obj.ReplicaZones, r)
	}
	for _, r := range p.GetResourcePolicy() {
		obj.ResourcePolicy = append(obj.ResourcePolicy, r)
	}
	for _, r := range p.GetLicenses() {
		obj.Licenses = append(obj.Licenses, r)
	}
	for _, r := range p.GetGuestOsFeatures() {
		obj.GuestOsFeatures = append(obj.GuestOsFeatures, *ProtoToComputeBetaDiskGuestOsFeatures(r))
	}
	for _, r := range p.GetUsers() {
		obj.Users = append(obj.Users, r)
	}
	for _, r := range p.GetLicenseCodes() {
		obj.LicenseCodes = append(obj.LicenseCodes, r)
	}
	for _, r := range p.GetResourcePolicies() {
		obj.ResourcePolicies = append(obj.ResourcePolicies, r)
	}
	return obj
}

// DiskGuestOsFeatureTypeEnumToProto converts a DiskGuestOsFeatureTypeEnum enum to its proto representation.
func ComputeBetaDiskGuestOsFeatureTypeEnumToProto(e *beta.DiskGuestOsFeatureTypeEnum) betapb.ComputeBetaDiskGuestOsFeatureTypeEnum {
	if e == nil {
		return betapb.ComputeBetaDiskGuestOsFeatureTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaDiskGuestOsFeatureTypeEnum_value["DiskGuestOsFeatureTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaDiskGuestOsFeatureTypeEnum(v)
	}
	return betapb.ComputeBetaDiskGuestOsFeatureTypeEnum(0)
}

// DiskGuestOsFeatureTypeAltEnumToProto converts a DiskGuestOsFeatureTypeAltEnum enum to its proto representation.
func ComputeBetaDiskGuestOsFeatureTypeAltEnumToProto(e *beta.DiskGuestOsFeatureTypeAltEnum) betapb.ComputeBetaDiskGuestOsFeatureTypeAltEnum {
	if e == nil {
		return betapb.ComputeBetaDiskGuestOsFeatureTypeAltEnum(0)
	}
	if v, ok := betapb.ComputeBetaDiskGuestOsFeatureTypeAltEnum_value["DiskGuestOsFeatureTypeAltEnum"+string(*e)]; ok {
		return betapb.ComputeBetaDiskGuestOsFeatureTypeAltEnum(v)
	}
	return betapb.ComputeBetaDiskGuestOsFeatureTypeAltEnum(0)
}

// DiskStatusEnumToProto converts a DiskStatusEnum enum to its proto representation.
func ComputeBetaDiskStatusEnumToProto(e *beta.DiskStatusEnum) betapb.ComputeBetaDiskStatusEnum {
	if e == nil {
		return betapb.ComputeBetaDiskStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaDiskStatusEnum_value["DiskStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaDiskStatusEnum(v)
	}
	return betapb.ComputeBetaDiskStatusEnum(0)
}

// DiskGuestOsFeaturesTypeEnumToProto converts a DiskGuestOsFeaturesTypeEnum enum to its proto representation.
func ComputeBetaDiskGuestOsFeaturesTypeEnumToProto(e *beta.DiskGuestOsFeaturesTypeEnum) betapb.ComputeBetaDiskGuestOsFeaturesTypeEnum {
	if e == nil {
		return betapb.ComputeBetaDiskGuestOsFeaturesTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaDiskGuestOsFeaturesTypeEnum_value["DiskGuestOsFeaturesTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaDiskGuestOsFeaturesTypeEnum(v)
	}
	return betapb.ComputeBetaDiskGuestOsFeaturesTypeEnum(0)
}

// DiskGuestOsFeaturesTypeAltsEnumToProto converts a DiskGuestOsFeaturesTypeAltsEnum enum to its proto representation.
func ComputeBetaDiskGuestOsFeaturesTypeAltsEnumToProto(e *beta.DiskGuestOsFeaturesTypeAltsEnum) betapb.ComputeBetaDiskGuestOsFeaturesTypeAltsEnum {
	if e == nil {
		return betapb.ComputeBetaDiskGuestOsFeaturesTypeAltsEnum(0)
	}
	if v, ok := betapb.ComputeBetaDiskGuestOsFeaturesTypeAltsEnum_value["DiskGuestOsFeaturesTypeAltsEnum"+string(*e)]; ok {
		return betapb.ComputeBetaDiskGuestOsFeaturesTypeAltsEnum(v)
	}
	return betapb.ComputeBetaDiskGuestOsFeaturesTypeAltsEnum(0)
}

// DiskGuestOsFeatureToProto converts a DiskGuestOsFeature resource to its proto representation.
func ComputeBetaDiskGuestOsFeatureToProto(o *beta.DiskGuestOsFeature) *betapb.ComputeBetaDiskGuestOsFeature {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaDiskGuestOsFeature{
		Type: ComputeBetaDiskGuestOsFeatureTypeEnumToProto(o.Type),
	}
	for _, r := range o.TypeAlt {
		p.TypeAlt = append(p.TypeAlt, betapb.ComputeBetaDiskGuestOsFeatureTypeAltEnum(betapb.ComputeBetaDiskGuestOsFeatureTypeAltEnum_value[string(r)]))
	}
	return p
}

// DiskEncryptionKeyToProto converts a DiskEncryptionKey resource to its proto representation.
func ComputeBetaDiskEncryptionKeyToProto(o *beta.DiskEncryptionKey) *betapb.ComputeBetaDiskEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaDiskEncryptionKey{
		RawKey:               dcl.ValueOrEmptyString(o.RawKey),
		KmsKeyName:           dcl.ValueOrEmptyString(o.KmsKeyName),
		Sha256:               dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyServiceAccount: dcl.ValueOrEmptyString(o.KmsKeyServiceAccount),
	}
	return p
}

// DiskGuestOsFeaturesToProto converts a DiskGuestOsFeatures resource to its proto representation.
func ComputeBetaDiskGuestOsFeaturesToProto(o *beta.DiskGuestOsFeatures) *betapb.ComputeBetaDiskGuestOsFeatures {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaDiskGuestOsFeatures{
		Type: ComputeBetaDiskGuestOsFeaturesTypeEnumToProto(o.Type),
	}
	for _, r := range o.TypeAlts {
		p.TypeAlts = append(p.TypeAlts, betapb.ComputeBetaDiskGuestOsFeaturesTypeAltsEnum(betapb.ComputeBetaDiskGuestOsFeaturesTypeAltsEnum_value[string(r)]))
	}
	return p
}

// DiskToProto converts a Disk resource to its proto representation.
func DiskToProto(resource *beta.Disk) *betapb.ComputeBetaDisk {
	p := &betapb.ComputeBetaDisk{
		SelfLink:                    dcl.ValueOrEmptyString(resource.SelfLink),
		Description:                 dcl.ValueOrEmptyString(resource.Description),
		DiskEncryptionKey:           ComputeBetaDiskEncryptionKeyToProto(resource.DiskEncryptionKey),
		LabelFingerprint:            dcl.ValueOrEmptyString(resource.LabelFingerprint),
		Name:                        dcl.ValueOrEmptyString(resource.Name),
		Region:                      dcl.ValueOrEmptyString(resource.Region),
		SizeGb:                      dcl.ValueOrEmptyInt64(resource.SizeGb),
		SourceImage:                 dcl.ValueOrEmptyString(resource.SourceImage),
		SourceImageEncryptionKey:    ComputeBetaDiskEncryptionKeyToProto(resource.SourceImageEncryptionKey),
		SourceImageId:               dcl.ValueOrEmptyString(resource.SourceImageId),
		SourceSnapshot:              dcl.ValueOrEmptyString(resource.SourceSnapshot),
		SourceSnapshotEncryptionKey: ComputeBetaDiskEncryptionKeyToProto(resource.SourceSnapshotEncryptionKey),
		SourceSnapshotId:            dcl.ValueOrEmptyString(resource.SourceSnapshotId),
		Type:                        dcl.ValueOrEmptyString(resource.Type),
		Zone:                        dcl.ValueOrEmptyString(resource.Zone),
		Project:                     dcl.ValueOrEmptyString(resource.Project),
		Id:                          dcl.ValueOrEmptyInt64(resource.Id),
		Status:                      ComputeBetaDiskStatusEnumToProto(resource.Status),
		Options:                     dcl.ValueOrEmptyString(resource.Options),
		LastAttachTimestamp:         dcl.ValueOrEmptyString(resource.LastAttachTimestamp),
		LastDetachTimestamp:         dcl.ValueOrEmptyString(resource.LastDetachTimestamp),
		PhysicalBlockSizeBytes:      dcl.ValueOrEmptyInt64(resource.PhysicalBlockSizeBytes),
		SourceDisk:                  dcl.ValueOrEmptyString(resource.SourceDisk),
		SourceDiskId:                dcl.ValueOrEmptyString(resource.SourceDiskId),
		Location:                    dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.GuestOsFeature {
		p.GuestOsFeature = append(p.GuestOsFeature, ComputeBetaDiskGuestOsFeatureToProto(&r))
	}
	for _, r := range resource.License {
		p.License = append(p.License, r)
	}
	for _, r := range resource.ReplicaZones {
		p.ReplicaZones = append(p.ReplicaZones, r)
	}
	for _, r := range resource.ResourcePolicy {
		p.ResourcePolicy = append(p.ResourcePolicy, r)
	}
	for _, r := range resource.Licenses {
		p.Licenses = append(p.Licenses, r)
	}
	for _, r := range resource.GuestOsFeatures {
		p.GuestOsFeatures = append(p.GuestOsFeatures, ComputeBetaDiskGuestOsFeaturesToProto(&r))
	}
	for _, r := range resource.Users {
		p.Users = append(p.Users, r)
	}
	for _, r := range resource.LicenseCodes {
		p.LicenseCodes = append(p.LicenseCodes, r)
	}
	for _, r := range resource.ResourcePolicies {
		p.ResourcePolicies = append(p.ResourcePolicies, r)
	}

	return p
}

// ApplyDisk handles the gRPC request by passing it to the underlying Disk Apply() method.
func (s *DiskServer) applyDisk(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaDiskRequest) (*betapb.ComputeBetaDisk, error) {
	p := ProtoToDisk(request.GetResource())
	res, err := c.ApplyDisk(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DiskToProto(res)
	return r, nil
}

// ApplyDisk handles the gRPC request by passing it to the underlying Disk Apply() method.
func (s *DiskServer) ApplyComputeBetaDisk(ctx context.Context, request *betapb.ApplyComputeBetaDiskRequest) (*betapb.ComputeBetaDisk, error) {
	cl, err := createConfigDisk(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyDisk(ctx, cl, request)
}

// DeleteDisk handles the gRPC request by passing it to the underlying Disk Delete() method.
func (s *DiskServer) DeleteComputeBetaDisk(ctx context.Context, request *betapb.DeleteComputeBetaDiskRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDisk(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDisk(ctx, ProtoToDisk(request.GetResource()))

}

// ListComputeBetaDisk handles the gRPC request by passing it to the underlying DiskList() method.
func (s *DiskServer) ListComputeBetaDisk(ctx context.Context, request *betapb.ListComputeBetaDiskRequest) (*betapb.ListComputeBetaDiskResponse, error) {
	cl, err := createConfigDisk(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDisk(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaDisk
	for _, r := range resources.Items {
		rp := DiskToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaDiskResponse{Items: protos}, nil
}

func createConfigDisk(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
