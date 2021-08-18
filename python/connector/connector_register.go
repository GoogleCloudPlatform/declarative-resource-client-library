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
package connector

import (
	connectorpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/connector_go_proto"

	assuredworkloads_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/connector"

	assuredworkloads_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/beta_connector"

	bigqueryreservation_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/connector"

	bigqueryreservation_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/beta_connector"

	binaryauthorization_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/connector"

	binaryauthorization_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/beta_connector"

	cloudbuild_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/beta_connector"

	cloudresourcemanager_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager/connector"

	cloudscheduler_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/connector"

	cloudscheduler_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/beta_connector"

	compute_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/services/redis"

	compute_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta_connector"

	containeranalysis_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/connector"

	containeranalysis_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/beta_connector"

	datafusion_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datafusion/beta_connector"

	dataproc_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/connector"

	dataproc_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta_connector"

	eventarc_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/connector"

	eventarc_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/beta_connector"

	file_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/file/connector"

	file_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/file/beta_connector"

	gameservices_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gameservices/connector"

	gameservices_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gameservices/beta_connector"

	gkehub_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta_connector"

	gkemulticloud_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud/connector"

	gkemulticloud_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud/beta_connector"

	iam_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/connector"

	iam_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/beta_connector"

	iap_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/connector"

	iap_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/beta_connector"

	identitytoolkit_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/connector"

	identitytoolkit_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/beta_connector"

	logging_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/connector"

	logging_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/beta_connector"

	monitoring_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/connector"

	monitoring_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/beta_connector"

	networksecurity_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/alpha_connector"

	networksecurity_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/beta_connector"

	networkservices_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/alpha_connector"

	networkservices_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/beta_connector"

	osconfig_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/alpha_connector"

	osconfig_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/beta_connector"

	pubsub_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub/connector"

	storage_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage/connector"

	privateca_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/connector"

	tier2_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/tier2/alpha_connector"

	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// InitializeServer prepares the server for future RPC requests. It must be called before
// attempting to response to any requests.
func InitializeServer(grpcServer *grpc.Server) *connectorpb.InitializeResponse {

	assuredworkloads_connector.RegisterServers(grpcServer)

	assuredworkloads_beta_connector.RegisterServers(grpcServer)

	bigqueryreservation_connector.RegisterServers(grpcServer)

	bigqueryreservation_beta_connector.RegisterServers(grpcServer)

	binaryauthorization_connector.RegisterServers(grpcServer)

	binaryauthorization_beta_connector.RegisterServers(grpcServer)

	cloudbuild_beta_connector.RegisterServers(grpcServer)

	cloudresourcemanager_connector.RegisterServers(grpcServer)

	cloudscheduler_connector.RegisterServers(grpcServer)

	cloudscheduler_beta_connector.RegisterServers(grpcServer)

	compute_connector.RegisterServers(grpcServer)

	compute_beta_connector.RegisterServers(grpcServer)

	containeranalysis_connector.RegisterServers(grpcServer)

	containeranalysis_beta_connector.RegisterServers(grpcServer)

	datafusion_beta_connector.RegisterServers(grpcServer)

	dataproc_connector.RegisterServers(grpcServer)

	dataproc_beta_connector.RegisterServers(grpcServer)

	eventarc_connector.RegisterServers(grpcServer)

	eventarc_beta_connector.RegisterServers(grpcServer)

	file_connector.RegisterServers(grpcServer)

	file_beta_connector.RegisterServers(grpcServer)

	gameservices_connector.RegisterServers(grpcServer)

	gameservices_beta_connector.RegisterServers(grpcServer)

	gkehub_beta_connector.RegisterServers(grpcServer)

	gkemulticloud_connector.RegisterServers(grpcServer)

	gkemulticloud_beta_connector.RegisterServers(grpcServer)

	iam_connector.RegisterServers(grpcServer)

	iam_beta_connector.RegisterServers(grpcServer)

	iap_connector.RegisterServers(grpcServer)

	iap_beta_connector.RegisterServers(grpcServer)

	identitytoolkit_connector.RegisterServers(grpcServer)

	identitytoolkit_beta_connector.RegisterServers(grpcServer)

	logging_connector.RegisterServers(grpcServer)

	logging_beta_connector.RegisterServers(grpcServer)

	monitoring_connector.RegisterServers(grpcServer)

	monitoring_beta_connector.RegisterServers(grpcServer)

	networksecurity_alpha_connector.RegisterServers(grpcServer)

	networksecurity_beta_connector.RegisterServers(grpcServer)

	networkservices_alpha_connector.RegisterServers(grpcServer)

	networkservices_beta_connector.RegisterServers(grpcServer)

	osconfig_alpha_connector.RegisterServers(grpcServer)

	osconfig_beta_connector.RegisterServers(grpcServer)

	pubsub_connector.RegisterServers(grpcServer)

	storage_connector.RegisterServers(grpcServer)

	privateca_connector.RegisterServers(grpcServer)

	tier2_alpha_connector.RegisterServers(grpcServer)

	return &connectorpb.InitializeResponse{
		Status: &statuspb.Status{
			Code: int32(codes.OK),
		},
	}
}
