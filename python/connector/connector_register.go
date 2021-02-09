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

	accesscontextmanager_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/accesscontextmanager/connector"

	appengine_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/appengine/connector"

	cloudbilling_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbilling/connector"

	cloudbuild_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/connector"

	cloudfunctions_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudfunctions/connector"

	cloudresourcemanager_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager/connector"

	cloudscheduler_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/connector"

	compute_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/services/redis"

	container_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/container/connector"

	containeranalysis_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/connector"

	dataproc_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/connector"

	datastore_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datastore/connector"

	dns_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dns/connector"

	file_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/file/connector"

	eventarc_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/beta_connector"

	gameservices_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gameservices/connector"

	logging_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/connector"

	iam_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/connector"

	iap_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/connector"

	monitoring_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/connector"

	pubsub_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub/connector"

	redis_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/redis/connector"

	run_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/run/connector"

	runtimeconfig_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/runtimeconfig/connector"

	servicenetworking_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/servicenetworking/connector"

	sourcerepo_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/sourcerepo/connector"

	serviceusage_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/serviceusage/connector"

	spanner_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/spanner/connector"

	sql_beta_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/sql/beta_connector"

	storage_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage/connector"

	tpu_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/tpu/connector"

	vpcaccess_connector "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vpcaccess/connector"

	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	tfjson "github.com/GoogleCloudPlatform/declarative-resource-client-library/tfjson"
	tfjsongrpc "github.com/GoogleCloudPlatform/declarative-resource-client-library/tfjson/go_grpc"
)

// InitializeServer prepares the server for future RPC requests. It must be called before
// attempting to response to any requests.
func InitializeServer(grpcServer *grpc.Server) *connectorpb.InitializeResponse {

	accesscontextmanager_connector.RegisterServers(grpcServer)

	appengine_connector.RegisterServers(grpcServer)

	cloudbilling_connector.RegisterServers(grpcServer)

	cloudbuild_connector.RegisterServers(grpcServer)

	cloudfunctions_connector.RegisterServers(grpcServer)

	cloudresourcemanager_connector.RegisterServers(grpcServer)

	cloudscheduler_connector.RegisterServers(grpcServer)

	compute_connector.RegisterServers(grpcServer)

	container_connector.RegisterServers(grpcServer)

	containeranalysis_connector.RegisterServers(grpcServer)

	dataproc_connector.RegisterServers(grpcServer)

	datastore_connector.RegisterServers(grpcServer)

	dns_connector.RegisterServers(grpcServer)

	file_connector.RegisterServers(grpcServer)

	eventarc_beta_connector.RegisterServers(grpcServer)

	gameservices_connector.RegisterServers(grpcServer)

	logging_connector.RegisterServers(grpcServer)

	iam_connector.RegisterServers(grpcServer)

	iap_connector.RegisterServers(grpcServer)

	monitoring_connector.RegisterServers(grpcServer)

	pubsub_connector.RegisterServers(grpcServer)

	redis_connector.RegisterServers(grpcServer)

	run_connector.RegisterServers(grpcServer)

	runtimeconfig_connector.RegisterServers(grpcServer)

	servicenetworking_connector.RegisterServers(grpcServer)

	sourcerepo_connector.RegisterServers(grpcServer)

	serviceusage_connector.RegisterServers(grpcServer)

	spanner_connector.RegisterServers(grpcServer)

	sql_beta_connector.RegisterServers(grpcServer)

	storage_connector.RegisterServers(grpcServer)

	tpu_connector.RegisterServers(grpcServer)

	vpcaccess_connector.RegisterServers(grpcServer)

	tfjsongrpc.RegisterTFJSONtoDCLServiceServer(grpcServer, &tfjson.TFJSONtoDCLService{})

	return &connectorpb.InitializeResponse{
		Status: &statuspb.Status{
			Code: int32(codes.OK),
		},
	}
}
