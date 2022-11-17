// Copyright 2022 Google LLC. All Rights Reserved.
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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vmware/alpha/vmware_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vmware/alpha"
)

// PrivateCloudServer implements the gRPC interface for PrivateCloud.
type PrivateCloudServer struct{}

// ProtoToPrivateCloudStateEnum converts a PrivateCloudStateEnum enum from its proto representation.
func ProtoToVmwareAlphaPrivateCloudStateEnum(e alphapb.VmwareAlphaPrivateCloudStateEnum) *alpha.PrivateCloudStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VmwareAlphaPrivateCloudStateEnum_name[int32(e)]; ok {
		e := alpha.PrivateCloudStateEnum(n[len("VmwareAlphaPrivateCloudStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToPrivateCloudNetworkConfigIPAddressLayoutVersionEnum converts a PrivateCloudNetworkConfigIPAddressLayoutVersionEnum enum from its proto representation.
func ProtoToVmwareAlphaPrivateCloudNetworkConfigIPAddressLayoutVersionEnum(e alphapb.VmwareAlphaPrivateCloudNetworkConfigIPAddressLayoutVersionEnum) *alpha.PrivateCloudNetworkConfigIPAddressLayoutVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VmwareAlphaPrivateCloudNetworkConfigIPAddressLayoutVersionEnum_name[int32(e)]; ok {
		e := alpha.PrivateCloudNetworkConfigIPAddressLayoutVersionEnum(n[len("VmwareAlphaPrivateCloudNetworkConfigIPAddressLayoutVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToPrivateCloudHcxStateEnum converts a PrivateCloudHcxStateEnum enum from its proto representation.
func ProtoToVmwareAlphaPrivateCloudHcxStateEnum(e alphapb.VmwareAlphaPrivateCloudHcxStateEnum) *alpha.PrivateCloudHcxStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VmwareAlphaPrivateCloudHcxStateEnum_name[int32(e)]; ok {
		e := alpha.PrivateCloudHcxStateEnum(n[len("VmwareAlphaPrivateCloudHcxStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToPrivateCloudNsxStateEnum converts a PrivateCloudNsxStateEnum enum from its proto representation.
func ProtoToVmwareAlphaPrivateCloudNsxStateEnum(e alphapb.VmwareAlphaPrivateCloudNsxStateEnum) *alpha.PrivateCloudNsxStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VmwareAlphaPrivateCloudNsxStateEnum_name[int32(e)]; ok {
		e := alpha.PrivateCloudNsxStateEnum(n[len("VmwareAlphaPrivateCloudNsxStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToPrivateCloudVcenterStateEnum converts a PrivateCloudVcenterStateEnum enum from its proto representation.
func ProtoToVmwareAlphaPrivateCloudVcenterStateEnum(e alphapb.VmwareAlphaPrivateCloudVcenterStateEnum) *alpha.PrivateCloudVcenterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VmwareAlphaPrivateCloudVcenterStateEnum_name[int32(e)]; ok {
		e := alpha.PrivateCloudVcenterStateEnum(n[len("VmwareAlphaPrivateCloudVcenterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToPrivateCloudNetworkConfig converts a PrivateCloudNetworkConfig object from its proto representation.
func ProtoToVmwareAlphaPrivateCloudNetworkConfig(p *alphapb.VmwareAlphaPrivateCloudNetworkConfig) *alpha.PrivateCloudNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.PrivateCloudNetworkConfig{
		Network:                      dcl.StringOrNil(p.GetNetwork()),
		ServiceNetwork:               dcl.StringOrNil(p.GetServiceNetwork()),
		ManagementCidr:               dcl.StringOrNil(p.GetManagementCidr()),
		VmwareEngineNetwork:          dcl.StringOrNil(p.GetVmwareEngineNetwork()),
		VmwareEngineNetworkCanonical: dcl.StringOrNil(p.GetVmwareEngineNetworkCanonical()),
		IPAddressLayoutVersion:       ProtoToVmwareAlphaPrivateCloudNetworkConfigIPAddressLayoutVersionEnum(p.GetIpAddressLayoutVersion()),
	}
	return obj
}

// ProtoToPrivateCloudManagementCluster converts a PrivateCloudManagementCluster object from its proto representation.
func ProtoToVmwareAlphaPrivateCloudManagementCluster(p *alphapb.VmwareAlphaPrivateCloudManagementCluster) *alpha.PrivateCloudManagementCluster {
	if p == nil {
		return nil
	}
	obj := &alpha.PrivateCloudManagementCluster{
		ClusterId:           dcl.StringOrNil(p.GetClusterId()),
		NodeTypeId:          dcl.StringOrNil(p.GetNodeTypeId()),
		NodeCount:           dcl.Int64OrNil(p.GetNodeCount()),
		NodeCustomCoreCount: dcl.Int64OrNil(p.GetNodeCustomCoreCount()),
	}
	return obj
}

// ProtoToPrivateCloudConditions converts a PrivateCloudConditions object from its proto representation.
func ProtoToVmwareAlphaPrivateCloudConditions(p *alphapb.VmwareAlphaPrivateCloudConditions) *alpha.PrivateCloudConditions {
	if p == nil {
		return nil
	}
	obj := &alpha.PrivateCloudConditions{
		Code:    dcl.StringOrNil(p.GetCode()),
		Message: dcl.StringOrNil(p.GetMessage()),
	}
	return obj
}

// ProtoToPrivateCloudHcx converts a PrivateCloudHcx object from its proto representation.
func ProtoToVmwareAlphaPrivateCloudHcx(p *alphapb.VmwareAlphaPrivateCloudHcx) *alpha.PrivateCloudHcx {
	if p == nil {
		return nil
	}
	obj := &alpha.PrivateCloudHcx{
		Fdqn:       dcl.StringOrNil(p.GetFdqn()),
		InternalIP: dcl.StringOrNil(p.GetInternalIp()),
		ExternalIP: dcl.StringOrNil(p.GetExternalIp()),
		Version:    dcl.StringOrNil(p.GetVersion()),
		State:      ProtoToVmwareAlphaPrivateCloudHcxStateEnum(p.GetState()),
		Fqdn:       dcl.StringOrNil(p.GetFqdn()),
	}
	return obj
}

// ProtoToPrivateCloudNsx converts a PrivateCloudNsx object from its proto representation.
func ProtoToVmwareAlphaPrivateCloudNsx(p *alphapb.VmwareAlphaPrivateCloudNsx) *alpha.PrivateCloudNsx {
	if p == nil {
		return nil
	}
	obj := &alpha.PrivateCloudNsx{
		Fdqn:       dcl.StringOrNil(p.GetFdqn()),
		InternalIP: dcl.StringOrNil(p.GetInternalIp()),
		ExternalIP: dcl.StringOrNil(p.GetExternalIp()),
		Version:    dcl.StringOrNil(p.GetVersion()),
		State:      ProtoToVmwareAlphaPrivateCloudNsxStateEnum(p.GetState()),
		Fqdn:       dcl.StringOrNil(p.GetFqdn()),
	}
	return obj
}

// ProtoToPrivateCloudVcenter converts a PrivateCloudVcenter object from its proto representation.
func ProtoToVmwareAlphaPrivateCloudVcenter(p *alphapb.VmwareAlphaPrivateCloudVcenter) *alpha.PrivateCloudVcenter {
	if p == nil {
		return nil
	}
	obj := &alpha.PrivateCloudVcenter{
		Fdqn:       dcl.StringOrNil(p.GetFdqn()),
		InternalIP: dcl.StringOrNil(p.GetInternalIp()),
		ExternalIP: dcl.StringOrNil(p.GetExternalIp()),
		Version:    dcl.StringOrNil(p.GetVersion()),
		State:      ProtoToVmwareAlphaPrivateCloudVcenterStateEnum(p.GetState()),
		Fqdn:       dcl.StringOrNil(p.GetFqdn()),
	}
	return obj
}

// ProtoToPrivateCloud converts a PrivateCloud resource from its proto representation.
func ProtoToPrivateCloud(p *alphapb.VmwareAlphaPrivateCloud) *alpha.PrivateCloud {
	obj := &alpha.PrivateCloud{
		Name:              dcl.StringOrNil(p.GetName()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:        dcl.StringOrNil(p.GetDeleteTime()),
		ExpireTime:        dcl.StringOrNil(p.GetExpireTime()),
		State:             ProtoToVmwareAlphaPrivateCloudStateEnum(p.GetState()),
		NetworkConfig:     ProtoToVmwareAlphaPrivateCloudNetworkConfig(p.GetNetworkConfig()),
		ManagementCluster: ProtoToVmwareAlphaPrivateCloudManagementCluster(p.GetManagementCluster()),
		Description:       dcl.StringOrNil(p.GetDescription()),
		Hcx:               ProtoToVmwareAlphaPrivateCloudHcx(p.GetHcx()),
		Nsx:               ProtoToVmwareAlphaPrivateCloudNsx(p.GetNsx()),
		Vcenter:           ProtoToVmwareAlphaPrivateCloudVcenter(p.GetVcenter()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
		Uid:               dcl.StringOrNil(p.GetUid()),
		Etag:              dcl.StringOrNil(p.GetEtag()),
	}
	for _, r := range p.GetConditions() {
		obj.Conditions = append(obj.Conditions, *ProtoToVmwareAlphaPrivateCloudConditions(r))
	}
	return obj
}

// PrivateCloudStateEnumToProto converts a PrivateCloudStateEnum enum to its proto representation.
func VmwareAlphaPrivateCloudStateEnumToProto(e *alpha.PrivateCloudStateEnum) alphapb.VmwareAlphaPrivateCloudStateEnum {
	if e == nil {
		return alphapb.VmwareAlphaPrivateCloudStateEnum(0)
	}
	if v, ok := alphapb.VmwareAlphaPrivateCloudStateEnum_value["PrivateCloudStateEnum"+string(*e)]; ok {
		return alphapb.VmwareAlphaPrivateCloudStateEnum(v)
	}
	return alphapb.VmwareAlphaPrivateCloudStateEnum(0)
}

// PrivateCloudNetworkConfigIPAddressLayoutVersionEnumToProto converts a PrivateCloudNetworkConfigIPAddressLayoutVersionEnum enum to its proto representation.
func VmwareAlphaPrivateCloudNetworkConfigIPAddressLayoutVersionEnumToProto(e *alpha.PrivateCloudNetworkConfigIPAddressLayoutVersionEnum) alphapb.VmwareAlphaPrivateCloudNetworkConfigIPAddressLayoutVersionEnum {
	if e == nil {
		return alphapb.VmwareAlphaPrivateCloudNetworkConfigIPAddressLayoutVersionEnum(0)
	}
	if v, ok := alphapb.VmwareAlphaPrivateCloudNetworkConfigIPAddressLayoutVersionEnum_value["PrivateCloudNetworkConfigIPAddressLayoutVersionEnum"+string(*e)]; ok {
		return alphapb.VmwareAlphaPrivateCloudNetworkConfigIPAddressLayoutVersionEnum(v)
	}
	return alphapb.VmwareAlphaPrivateCloudNetworkConfigIPAddressLayoutVersionEnum(0)
}

// PrivateCloudHcxStateEnumToProto converts a PrivateCloudHcxStateEnum enum to its proto representation.
func VmwareAlphaPrivateCloudHcxStateEnumToProto(e *alpha.PrivateCloudHcxStateEnum) alphapb.VmwareAlphaPrivateCloudHcxStateEnum {
	if e == nil {
		return alphapb.VmwareAlphaPrivateCloudHcxStateEnum(0)
	}
	if v, ok := alphapb.VmwareAlphaPrivateCloudHcxStateEnum_value["PrivateCloudHcxStateEnum"+string(*e)]; ok {
		return alphapb.VmwareAlphaPrivateCloudHcxStateEnum(v)
	}
	return alphapb.VmwareAlphaPrivateCloudHcxStateEnum(0)
}

// PrivateCloudNsxStateEnumToProto converts a PrivateCloudNsxStateEnum enum to its proto representation.
func VmwareAlphaPrivateCloudNsxStateEnumToProto(e *alpha.PrivateCloudNsxStateEnum) alphapb.VmwareAlphaPrivateCloudNsxStateEnum {
	if e == nil {
		return alphapb.VmwareAlphaPrivateCloudNsxStateEnum(0)
	}
	if v, ok := alphapb.VmwareAlphaPrivateCloudNsxStateEnum_value["PrivateCloudNsxStateEnum"+string(*e)]; ok {
		return alphapb.VmwareAlphaPrivateCloudNsxStateEnum(v)
	}
	return alphapb.VmwareAlphaPrivateCloudNsxStateEnum(0)
}

// PrivateCloudVcenterStateEnumToProto converts a PrivateCloudVcenterStateEnum enum to its proto representation.
func VmwareAlphaPrivateCloudVcenterStateEnumToProto(e *alpha.PrivateCloudVcenterStateEnum) alphapb.VmwareAlphaPrivateCloudVcenterStateEnum {
	if e == nil {
		return alphapb.VmwareAlphaPrivateCloudVcenterStateEnum(0)
	}
	if v, ok := alphapb.VmwareAlphaPrivateCloudVcenterStateEnum_value["PrivateCloudVcenterStateEnum"+string(*e)]; ok {
		return alphapb.VmwareAlphaPrivateCloudVcenterStateEnum(v)
	}
	return alphapb.VmwareAlphaPrivateCloudVcenterStateEnum(0)
}

// PrivateCloudNetworkConfigToProto converts a PrivateCloudNetworkConfig object to its proto representation.
func VmwareAlphaPrivateCloudNetworkConfigToProto(o *alpha.PrivateCloudNetworkConfig) *alphapb.VmwareAlphaPrivateCloudNetworkConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.VmwareAlphaPrivateCloudNetworkConfig{}
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetServiceNetwork(dcl.ValueOrEmptyString(o.ServiceNetwork))
	p.SetManagementCidr(dcl.ValueOrEmptyString(o.ManagementCidr))
	p.SetVmwareEngineNetwork(dcl.ValueOrEmptyString(o.VmwareEngineNetwork))
	p.SetVmwareEngineNetworkCanonical(dcl.ValueOrEmptyString(o.VmwareEngineNetworkCanonical))
	p.SetIpAddressLayoutVersion(VmwareAlphaPrivateCloudNetworkConfigIPAddressLayoutVersionEnumToProto(o.IPAddressLayoutVersion))
	return p
}

// PrivateCloudManagementClusterToProto converts a PrivateCloudManagementCluster object to its proto representation.
func VmwareAlphaPrivateCloudManagementClusterToProto(o *alpha.PrivateCloudManagementCluster) *alphapb.VmwareAlphaPrivateCloudManagementCluster {
	if o == nil {
		return nil
	}
	p := &alphapb.VmwareAlphaPrivateCloudManagementCluster{}
	p.SetClusterId(dcl.ValueOrEmptyString(o.ClusterId))
	p.SetNodeTypeId(dcl.ValueOrEmptyString(o.NodeTypeId))
	p.SetNodeCount(dcl.ValueOrEmptyInt64(o.NodeCount))
	p.SetNodeCustomCoreCount(dcl.ValueOrEmptyInt64(o.NodeCustomCoreCount))
	return p
}

// PrivateCloudConditionsToProto converts a PrivateCloudConditions object to its proto representation.
func VmwareAlphaPrivateCloudConditionsToProto(o *alpha.PrivateCloudConditions) *alphapb.VmwareAlphaPrivateCloudConditions {
	if o == nil {
		return nil
	}
	p := &alphapb.VmwareAlphaPrivateCloudConditions{}
	p.SetCode(dcl.ValueOrEmptyString(o.Code))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	return p
}

// PrivateCloudHcxToProto converts a PrivateCloudHcx object to its proto representation.
func VmwareAlphaPrivateCloudHcxToProto(o *alpha.PrivateCloudHcx) *alphapb.VmwareAlphaPrivateCloudHcx {
	if o == nil {
		return nil
	}
	p := &alphapb.VmwareAlphaPrivateCloudHcx{}
	p.SetFdqn(dcl.ValueOrEmptyString(o.Fdqn))
	p.SetInternalIp(dcl.ValueOrEmptyString(o.InternalIP))
	p.SetExternalIp(dcl.ValueOrEmptyString(o.ExternalIP))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetState(VmwareAlphaPrivateCloudHcxStateEnumToProto(o.State))
	p.SetFqdn(dcl.ValueOrEmptyString(o.Fqdn))
	return p
}

// PrivateCloudNsxToProto converts a PrivateCloudNsx object to its proto representation.
func VmwareAlphaPrivateCloudNsxToProto(o *alpha.PrivateCloudNsx) *alphapb.VmwareAlphaPrivateCloudNsx {
	if o == nil {
		return nil
	}
	p := &alphapb.VmwareAlphaPrivateCloudNsx{}
	p.SetFdqn(dcl.ValueOrEmptyString(o.Fdqn))
	p.SetInternalIp(dcl.ValueOrEmptyString(o.InternalIP))
	p.SetExternalIp(dcl.ValueOrEmptyString(o.ExternalIP))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetState(VmwareAlphaPrivateCloudNsxStateEnumToProto(o.State))
	p.SetFqdn(dcl.ValueOrEmptyString(o.Fqdn))
	return p
}

// PrivateCloudVcenterToProto converts a PrivateCloudVcenter object to its proto representation.
func VmwareAlphaPrivateCloudVcenterToProto(o *alpha.PrivateCloudVcenter) *alphapb.VmwareAlphaPrivateCloudVcenter {
	if o == nil {
		return nil
	}
	p := &alphapb.VmwareAlphaPrivateCloudVcenter{}
	p.SetFdqn(dcl.ValueOrEmptyString(o.Fdqn))
	p.SetInternalIp(dcl.ValueOrEmptyString(o.InternalIP))
	p.SetExternalIp(dcl.ValueOrEmptyString(o.ExternalIP))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetState(VmwareAlphaPrivateCloudVcenterStateEnumToProto(o.State))
	p.SetFqdn(dcl.ValueOrEmptyString(o.Fqdn))
	return p
}

// PrivateCloudToProto converts a PrivateCloud resource to its proto representation.
func PrivateCloudToProto(resource *alpha.PrivateCloud) *alphapb.VmwareAlphaPrivateCloud {
	p := &alphapb.VmwareAlphaPrivateCloud{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDeleteTime(dcl.ValueOrEmptyString(resource.DeleteTime))
	p.SetExpireTime(dcl.ValueOrEmptyString(resource.ExpireTime))
	p.SetState(VmwareAlphaPrivateCloudStateEnumToProto(resource.State))
	p.SetNetworkConfig(VmwareAlphaPrivateCloudNetworkConfigToProto(resource.NetworkConfig))
	p.SetManagementCluster(VmwareAlphaPrivateCloudManagementClusterToProto(resource.ManagementCluster))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetHcx(VmwareAlphaPrivateCloudHcxToProto(resource.Hcx))
	p.SetNsx(VmwareAlphaPrivateCloudNsxToProto(resource.Nsx))
	p.SetVcenter(VmwareAlphaPrivateCloudVcenterToProto(resource.Vcenter))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	sConditions := make([]*alphapb.VmwareAlphaPrivateCloudConditions, len(resource.Conditions))
	for i, r := range resource.Conditions {
		sConditions[i] = VmwareAlphaPrivateCloudConditionsToProto(&r)
	}
	p.SetConditions(sConditions)

	return p
}

// applyPrivateCloud handles the gRPC request by passing it to the underlying PrivateCloud Apply() method.
func (s *PrivateCloudServer) applyPrivateCloud(ctx context.Context, c *alpha.Client, request *alphapb.ApplyVmwareAlphaPrivateCloudRequest) (*alphapb.VmwareAlphaPrivateCloud, error) {
	p := ProtoToPrivateCloud(request.GetResource())
	res, err := c.ApplyPrivateCloud(ctx, p)
	if err != nil {
		return nil, err
	}
	r := PrivateCloudToProto(res)
	return r, nil
}

// applyVmwareAlphaPrivateCloud handles the gRPC request by passing it to the underlying PrivateCloud Apply() method.
func (s *PrivateCloudServer) ApplyVmwareAlphaPrivateCloud(ctx context.Context, request *alphapb.ApplyVmwareAlphaPrivateCloudRequest) (*alphapb.VmwareAlphaPrivateCloud, error) {
	cl, err := createConfigPrivateCloud(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyPrivateCloud(ctx, cl, request)
}

// DeletePrivateCloud handles the gRPC request by passing it to the underlying PrivateCloud Delete() method.
func (s *PrivateCloudServer) DeleteVmwareAlphaPrivateCloud(ctx context.Context, request *alphapb.DeleteVmwareAlphaPrivateCloudRequest) (*emptypb.Empty, error) {

	cl, err := createConfigPrivateCloud(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeletePrivateCloud(ctx, ProtoToPrivateCloud(request.GetResource()))

}

// ListVmwareAlphaPrivateCloud handles the gRPC request by passing it to the underlying PrivateCloudList() method.
func (s *PrivateCloudServer) ListVmwareAlphaPrivateCloud(ctx context.Context, request *alphapb.ListVmwareAlphaPrivateCloudRequest) (*alphapb.ListVmwareAlphaPrivateCloudResponse, error) {
	cl, err := createConfigPrivateCloud(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListPrivateCloud(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.VmwareAlphaPrivateCloud
	for _, r := range resources.Items {
		rp := PrivateCloudToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListVmwareAlphaPrivateCloudResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigPrivateCloud(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
