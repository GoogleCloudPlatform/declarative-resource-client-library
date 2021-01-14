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
	"google.golang.org/grpc"
	sdkgrpc "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
)

// RegisterServers registers each resource with the gRPC server.
func RegisterServers(s *grpc.Server) {
	sdkgrpc.RegisterComputeAddressServiceServer(s, &AddressServer{})
	sdkgrpc.RegisterComputeAutoscalerServiceServer(s, &AutoscalerServer{})
	sdkgrpc.RegisterComputeBackendBucketServiceServer(s, &BackendBucketServer{})
	sdkgrpc.RegisterComputeDiskServiceServer(s, &DiskServer{})
	sdkgrpc.RegisterComputeFirewallServiceServer(s, &FirewallServer{})
	sdkgrpc.RegisterComputeForwardingRuleServiceServer(s, &ForwardingRuleServer{})
	sdkgrpc.RegisterComputeHealthCheckServiceServer(s, &HealthCheckServer{})
	sdkgrpc.RegisterComputeHttpHealthCheckServiceServer(s, &HttpHealthCheckServer{})
	sdkgrpc.RegisterComputeHttpsHealthCheckServiceServer(s, &HttpsHealthCheckServer{})
	sdkgrpc.RegisterComputeImageServiceServer(s, &ImageServer{})
	sdkgrpc.RegisterComputeInstanceServiceServer(s, &InstanceServer{})
	sdkgrpc.RegisterComputeInstanceGroupManagerServiceServer(s, &InstanceGroupManagerServer{})
	sdkgrpc.RegisterComputeInstanceTemplateServiceServer(s, &InstanceTemplateServer{})
	sdkgrpc.RegisterComputeInterconnectServiceServer(s, &InterconnectServer{})
	sdkgrpc.RegisterComputeManagedSslCertificateServiceServer(s, &ManagedSslCertificateServer{})
	sdkgrpc.RegisterComputeNetworkServiceServer(s, &NetworkServer{})
	sdkgrpc.RegisterComputeReservationServiceServer(s, &ReservationServer{})
	sdkgrpc.RegisterComputeRouteServiceServer(s, &RouteServer{})
	sdkgrpc.RegisterComputeRouterServiceServer(s, &RouterServer{})
	sdkgrpc.RegisterComputeRouterPeerServiceServer(s, &RouterPeerServer{})
	sdkgrpc.RegisterComputeSnapshotServiceServer(s, &SnapshotServer{})
	sdkgrpc.RegisterComputeSslCertificateServiceServer(s, &SslCertificateServer{})
	sdkgrpc.RegisterComputeSslPolicyServiceServer(s, &SslPolicyServer{})
	sdkgrpc.RegisterComputeSubnetworkServiceServer(s, &SubnetworkServer{})
	sdkgrpc.RegisterComputeTargetHttpProxyServiceServer(s, &TargetHttpProxyServer{})
	sdkgrpc.RegisterComputeTargetPoolServiceServer(s, &TargetPoolServer{})
	sdkgrpc.RegisterComputeTargetSslProxyServiceServer(s, &TargetSslProxyServer{})
	sdkgrpc.RegisterComputeTargetVpnGatewayServiceServer(s, &TargetVpnGatewayServer{})
	sdkgrpc.RegisterComputeUrlMapServiceServer(s, &UrlMapServer{})
	sdkgrpc.RegisterComputeVpnGatewayServiceServer(s, &VpnGatewayServer{})
	sdkgrpc.RegisterComputeVpnTunnelServiceServer(s, &VpnTunnelServer{})
}
