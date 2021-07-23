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

	cloudbuild_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/beta_connector"

	accesscontextmanager_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/accesscontextmanager/connector"

	apigee_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/connector"

	apikeys_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apikeys/connector"

	appengine_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/appengine/connector"

	assuredworkloads_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/connector"

	bigquery_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery/connector"

	bigqueryconnection_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryconnection/connector"

	bigqueryreservation_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/connector"

	binaryauthorization_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/connector"

	cloudbilling_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbilling/connector"

	cloudfunctions_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudfunctions/connector"

	cloudresourcemanager_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager/connector"

	cloudscheduler_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/connector"

	composer_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/composer/connector"

	compute_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/services/redis"

	container_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/container/connector"

	containeranalysis_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/connector"

	datafusion_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datafusion/beta_connector"

	dataproc_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/connector"

	datastore_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datastore/connector"

	dns_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dns/connector"

	file_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/file/connector"

	eventarc_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/beta_connector"

	gameservices_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gameservices/connector"

	gkehub_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta_connector"

	gkemulticloud_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud/connector"

	logging_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/connector"

	iam_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/connector"

	iap_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/connector"

	identitytoolkit_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/connector"

	krmapihosting_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/krmapihosting/alpha_connector"

	monitoring_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/connector"

	networksecurity_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/alpha_connector"

	networkservices_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/alpha_connector"

	osconfig_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/alpha_connector"

	privateca_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/beta_connector"

	pubsub_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub/connector"

	pubsublite_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsublite/connector"

	redis_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/redis/connector"

	run_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/run/connector"

	runtimeconfig_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/runtimeconfig/beta_connector"

	servicemanagement_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/servicemanagement/connector"

	servicenetworking_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/servicenetworking/connector"

	sourcerepo_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/sourcerepo/connector"

	serviceusage_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/serviceusage/connector"

	spanner_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/spanner/connector"

	sql_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/sql/beta_connector"

	storage_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage/connector"

	tier2_alpha_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/tier2/alpha_connector"

	tpu_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/tpu/connector"

	vpcaccess_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vpcaccess/connector"

	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// InitializeServer prepares the server for future RPC requests. It must be called before
// attempting to response to any requests.
func InitializeServer(grpcServer *grpc.Server) *connectorpb.InitializeResponse {

	cloudbuild_beta_connector.RegisterServers(grpcServer)

	accesscontextmanager_connector.RegisterServers(grpcServer)

	apigee_connector.RegisterServers(grpcServer)

	apikeys_connector.RegisterServers(grpcServer)

	appengine_connector.RegisterServers(grpcServer)

	assuredworkloads_connector.RegisterServers(grpcServer)

	bigquery_connector.RegisterServers(grpcServer)

	bigqueryconnection_connector.RegisterServers(grpcServer)

	bigqueryreservation_connector.RegisterServers(grpcServer)

	binaryauthorization_connector.RegisterServers(grpcServer)

	cloudbilling_connector.RegisterServers(grpcServer)

	cloudfunctions_connector.RegisterServers(grpcServer)

	cloudresourcemanager_connector.RegisterServers(grpcServer)

	cloudscheduler_connector.RegisterServers(grpcServer)

	composer_connector.RegisterServers(grpcServer)

	compute_connector.RegisterServers(grpcServer)

	container_connector.RegisterServers(grpcServer)

	containeranalysis_connector.RegisterServers(grpcServer)

	datafusion_beta_connector.RegisterServers(grpcServer)

	dataproc_connector.RegisterServers(grpcServer)

	datastore_connector.RegisterServers(grpcServer)

	dns_connector.RegisterServers(grpcServer)

	file_connector.RegisterServers(grpcServer)

	eventarc_beta_connector.RegisterServers(grpcServer)

	gameservices_connector.RegisterServers(grpcServer)

	gkehub_beta_connector.RegisterServers(grpcServer)

	gkemulticloud_connector.RegisterServers(grpcServer)

	logging_connector.RegisterServers(grpcServer)

	iam_connector.RegisterServers(grpcServer)

	iap_connector.RegisterServers(grpcServer)

	identitytoolkit_connector.RegisterServers(grpcServer)

	krmapihosting_alpha_connector.RegisterServers(grpcServer)

	monitoring_connector.RegisterServers(grpcServer)

	networksecurity_alpha_connector.RegisterServers(grpcServer)

	networkservices_alpha_connector.RegisterServers(grpcServer)

	osconfig_alpha_connector.RegisterServers(grpcServer)

	privateca_beta_connector.RegisterServers(grpcServer)

	pubsub_connector.RegisterServers(grpcServer)

	pubsublite_connector.RegisterServers(grpcServer)

	redis_connector.RegisterServers(grpcServer)

	run_connector.RegisterServers(grpcServer)

	runtimeconfig_beta_connector.RegisterServers(grpcServer)

	servicemanagement_connector.RegisterServers(grpcServer)

	servicenetworking_connector.RegisterServers(grpcServer)

	sourcerepo_connector.RegisterServers(grpcServer)

	serviceusage_connector.RegisterServers(grpcServer)

	spanner_connector.RegisterServers(grpcServer)

	sql_beta_connector.RegisterServers(grpcServer)

	storage_connector.RegisterServers(grpcServer)

	tier2_alpha_connector.RegisterServers(grpcServer)

	tpu_connector.RegisterServers(grpcServer)

	vpcaccess_connector.RegisterServers(grpcServer)

	return &connectorpb.InitializeResponse{
		Status: &statuspb.Status{
			Code: int32(codes.OK),
		},
	}
}
