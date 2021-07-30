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
package directory

import (
	"bytes"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/accesscontextmanager"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apikeys"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/appengine"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads"
	assuredworkloads_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryconnection"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization"
	binaryauthorization_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbilling"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild"
	cloudbuild_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudfunctions"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/composer"
	composer_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/composer/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
	compute_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/container"
	container_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/container/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis"
	containeranalysis_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/beta"
	datafusion_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datafusion/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc"
	dataproc_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datastore"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dns"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc"
	eventarc_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/file"
	file_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/file/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gameservices"
	gkehub_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud"
	gkemulticloud_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit"
	krmapihosting_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/krmapihosting/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
	networksecurity_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/alpha"
	networksecurity_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/beta"
	networkservices_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/alpha"
	networkservices_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/beta"
	osconfig_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/alpha"
	osconfig_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/beta"
	privateca_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsublite"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/redis"
	redis_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/redis/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/run"
	runtimeconfig_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/runtimeconfig/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/servicemanagement"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/servicenetworking"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/serviceusage"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/sourcerepo"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/spanner"
	sql_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/sql/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage"
	tier2_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/tier2/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/tpu"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vpcaccess"
)

type Directory struct {
	versions []*version
}

type version struct {
	name     string
	services []*service
}

type service struct {
	name      string
	resources map[string]*bytes.Buffer
}

func (d *Directory) findOrCreateVersion(vName string) *version {
	for _, v := range d.versions {
		if v.name == vName {
			return v
		}
	}
	newV := &version{name: vName}
	d.versions = append(d.versions, newV)
	return newV
}

func (v *version) findOrCreateService(sName string) *service {
	for _, s := range v.services {
		if s.name == sName {
			return s
		}
	}
	newS := &service{name: sName, resources: make(map[string]*bytes.Buffer)}
	v.services = append(v.services, newS)
	return newS
}

func (d *Directory) AddResource(version, service, resource string, yaml []byte) {
	v := d.findOrCreateVersion(version)
	s := v.findOrCreateService(service)
	s.loadResource(resource, yaml)
}

func (s *service) loadResource(resourceName string, resource []byte) {
	s.resources[resourceName] = bytes.NewBuffer(resource)
}

func (d *Directory) GetResource(version, service, resource string) *bytes.Buffer {
	v := d.findOrCreateVersion(version)
	s := v.findOrCreateService(service)
	return s.resources[resource]
}

func Services() *Directory {
	d := &Directory{}
	d.AddResource("beta", "cloudbuild", dcl.TitleToSnakeCase("BuildTrigger"), cloudbuild_beta.YAML_build_trigger)
	d.AddResource("beta", "cloudbuild", "BuildTrigger", cloudbuild_beta.YAML_build_trigger)
	d.AddResource("beta", "cloudbuild", dcl.TitleToSnakeCase("WorkerPool"), cloudbuild_beta.YAML_worker_pool)
	d.AddResource("beta", "cloudbuild", "WorkerPool", cloudbuild_beta.YAML_worker_pool)
	d.AddResource("beta", "assuredworkloads", dcl.TitleToSnakeCase("Workload"), assuredworkloads_beta.YAML_workload)
	d.AddResource("beta", "assuredworkloads", "Workload", assuredworkloads_beta.YAML_workload)
	d.AddResource("beta", "binaryauthorization", dcl.TitleToSnakeCase("Attestor"), binaryauthorization_beta.YAML_attestor)
	d.AddResource("beta", "binaryauthorization", "Attestor", binaryauthorization_beta.YAML_attestor)
	d.AddResource("beta", "binaryauthorization", dcl.TitleToSnakeCase("Policy"), binaryauthorization_beta.YAML_policy)
	d.AddResource("beta", "binaryauthorization", "Policy", binaryauthorization_beta.YAML_policy)
	d.AddResource("beta", "composer", dcl.TitleToSnakeCase("Environment"), composer_beta.YAML_environment)
	d.AddResource("beta", "composer", "Environment", composer_beta.YAML_environment)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Address"), compute_beta.YAML_address)
	d.AddResource("beta", "compute", "Address", compute_beta.YAML_address)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Autoscaler"), compute_beta.YAML_autoscaler)
	d.AddResource("beta", "compute", "Autoscaler", compute_beta.YAML_autoscaler)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("BackendBucket"), compute_beta.YAML_backend_bucket)
	d.AddResource("beta", "compute", "BackendBucket", compute_beta.YAML_backend_bucket)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("BackendService"), compute_beta.YAML_backend_service)
	d.AddResource("beta", "compute", "BackendService", compute_beta.YAML_backend_service)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Disk"), compute_beta.YAML_disk)
	d.AddResource("beta", "compute", "Disk", compute_beta.YAML_disk)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Firewall"), compute_beta.YAML_firewall)
	d.AddResource("beta", "compute", "Firewall", compute_beta.YAML_firewall)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("FirewallPolicy"), compute_beta.YAML_firewall_policy)
	d.AddResource("beta", "compute", "FirewallPolicy", compute_beta.YAML_firewall_policy)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("FirewallPolicyAssociation"), compute_beta.YAML_firewall_policy_association)
	d.AddResource("beta", "compute", "FirewallPolicyAssociation", compute_beta.YAML_firewall_policy_association)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("FirewallPolicyRule"), compute_beta.YAML_firewall_policy_rule)
	d.AddResource("beta", "compute", "FirewallPolicyRule", compute_beta.YAML_firewall_policy_rule)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("ForwardingRule"), compute_beta.YAML_forwarding_rule)
	d.AddResource("beta", "compute", "ForwardingRule", compute_beta.YAML_forwarding_rule)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("HealthCheck"), compute_beta.YAML_health_check)
	d.AddResource("beta", "compute", "HealthCheck", compute_beta.YAML_health_check)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("HttpHealthCheck"), compute_beta.YAML_http_health_check)
	d.AddResource("beta", "compute", "HttpHealthCheck", compute_beta.YAML_http_health_check)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("HttpsHealthCheck"), compute_beta.YAML_https_health_check)
	d.AddResource("beta", "compute", "HttpsHealthCheck", compute_beta.YAML_https_health_check)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Image"), compute_beta.YAML_image)
	d.AddResource("beta", "compute", "Image", compute_beta.YAML_image)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Instance"), compute_beta.YAML_instance)
	d.AddResource("beta", "compute", "Instance", compute_beta.YAML_instance)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("InstanceGroupManager"), compute_beta.YAML_instance_group_manager)
	d.AddResource("beta", "compute", "InstanceGroupManager", compute_beta.YAML_instance_group_manager)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("InstanceTemplate"), compute_beta.YAML_instance_template)
	d.AddResource("beta", "compute", "InstanceTemplate", compute_beta.YAML_instance_template)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("ManagedSslCertificate"), compute_beta.YAML_managed_ssl_certificate)
	d.AddResource("beta", "compute", "ManagedSslCertificate", compute_beta.YAML_managed_ssl_certificate)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Network"), compute_beta.YAML_network)
	d.AddResource("beta", "compute", "Network", compute_beta.YAML_network)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("NetworkEndpoint"), compute_beta.YAML_network_endpoint)
	d.AddResource("beta", "compute", "NetworkEndpoint", compute_beta.YAML_network_endpoint)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("NetworkEndpointGroup"), compute_beta.YAML_network_endpoint_group)
	d.AddResource("beta", "compute", "NetworkEndpointGroup", compute_beta.YAML_network_endpoint_group)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Reservation"), compute_beta.YAML_reservation)
	d.AddResource("beta", "compute", "Reservation", compute_beta.YAML_reservation)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Route"), compute_beta.YAML_route)
	d.AddResource("beta", "compute", "Route", compute_beta.YAML_route)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Router"), compute_beta.YAML_router)
	d.AddResource("beta", "compute", "Router", compute_beta.YAML_router)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("RouterPeer"), compute_beta.YAML_router_peer)
	d.AddResource("beta", "compute", "RouterPeer", compute_beta.YAML_router_peer)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Snapshot"), compute_beta.YAML_snapshot)
	d.AddResource("beta", "compute", "Snapshot", compute_beta.YAML_snapshot)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("SslCertificate"), compute_beta.YAML_ssl_certificate)
	d.AddResource("beta", "compute", "SslCertificate", compute_beta.YAML_ssl_certificate)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("SslPolicy"), compute_beta.YAML_ssl_policy)
	d.AddResource("beta", "compute", "SslPolicy", compute_beta.YAML_ssl_policy)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Subnetwork"), compute_beta.YAML_subnetwork)
	d.AddResource("beta", "compute", "Subnetwork", compute_beta.YAML_subnetwork)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("TargetHttpProxy"), compute_beta.YAML_target_http_proxy)
	d.AddResource("beta", "compute", "TargetHttpProxy", compute_beta.YAML_target_http_proxy)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("TargetPool"), compute_beta.YAML_target_pool)
	d.AddResource("beta", "compute", "TargetPool", compute_beta.YAML_target_pool)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("TargetSslProxy"), compute_beta.YAML_target_ssl_proxy)
	d.AddResource("beta", "compute", "TargetSslProxy", compute_beta.YAML_target_ssl_proxy)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("TargetVpnGateway"), compute_beta.YAML_target_vpn_gateway)
	d.AddResource("beta", "compute", "TargetVpnGateway", compute_beta.YAML_target_vpn_gateway)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("UrlMap"), compute_beta.YAML_url_map)
	d.AddResource("beta", "compute", "UrlMap", compute_beta.YAML_url_map)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("VpnGateway"), compute_beta.YAML_vpn_gateway)
	d.AddResource("beta", "compute", "VpnGateway", compute_beta.YAML_vpn_gateway)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("VpnTunnel"), compute_beta.YAML_vpn_tunnel)
	d.AddResource("beta", "compute", "VpnTunnel", compute_beta.YAML_vpn_tunnel)
	d.AddResource("beta", "container", dcl.TitleToSnakeCase("Cluster"), container_beta.YAML_cluster)
	d.AddResource("beta", "container", "Cluster", container_beta.YAML_cluster)
	d.AddResource("beta", "container", dcl.TitleToSnakeCase("NodePool"), container_beta.YAML_node_pool)
	d.AddResource("beta", "container", "NodePool", container_beta.YAML_node_pool)
	d.AddResource("beta", "containeranalysis", dcl.TitleToSnakeCase("Note"), containeranalysis_beta.YAML_note)
	d.AddResource("beta", "containeranalysis", "Note", containeranalysis_beta.YAML_note)
	d.AddResource("beta", "datafusion", dcl.TitleToSnakeCase("Instance"), datafusion_beta.YAML_instance)
	d.AddResource("beta", "datafusion", "Instance", datafusion_beta.YAML_instance)
	d.AddResource("beta", "dataproc", dcl.TitleToSnakeCase("AutoscalingPolicy"), dataproc_beta.YAML_autoscaling_policy)
	d.AddResource("beta", "dataproc", "AutoscalingPolicy", dataproc_beta.YAML_autoscaling_policy)
	d.AddResource("beta", "dataproc", dcl.TitleToSnakeCase("Cluster"), dataproc_beta.YAML_cluster)
	d.AddResource("beta", "dataproc", "Cluster", dataproc_beta.YAML_cluster)
	d.AddResource("beta", "dataproc", dcl.TitleToSnakeCase("Job"), dataproc_beta.YAML_job)
	d.AddResource("beta", "dataproc", "Job", dataproc_beta.YAML_job)
	d.AddResource("beta", "dataproc", dcl.TitleToSnakeCase("WorkflowTemplate"), dataproc_beta.YAML_workflow_template)
	d.AddResource("beta", "dataproc", "WorkflowTemplate", dataproc_beta.YAML_workflow_template)
	d.AddResource("beta", "file", dcl.TitleToSnakeCase("Backup"), file_beta.YAML_backup)
	d.AddResource("beta", "file", "Backup", file_beta.YAML_backup)
	d.AddResource("beta", "file", dcl.TitleToSnakeCase("Instance"), file_beta.YAML_instance)
	d.AddResource("beta", "file", "Instance", file_beta.YAML_instance)
	d.AddResource("beta", "eventarc", dcl.TitleToSnakeCase("Trigger"), eventarc_beta.YAML_trigger)
	d.AddResource("beta", "eventarc", "Trigger", eventarc_beta.YAML_trigger)
	d.AddResource("beta", "gkehub", dcl.TitleToSnakeCase("Feature"), gkehub_beta.YAML_feature)
	d.AddResource("beta", "gkehub", "Feature", gkehub_beta.YAML_feature)
	d.AddResource("beta", "gkehub", dcl.TitleToSnakeCase("FeatureMembership"), gkehub_beta.YAML_feature_membership)
	d.AddResource("beta", "gkehub", "FeatureMembership", gkehub_beta.YAML_feature_membership)
	d.AddResource("beta", "gkehub", dcl.TitleToSnakeCase("Membership"), gkehub_beta.YAML_membership)
	d.AddResource("beta", "gkehub", "Membership", gkehub_beta.YAML_membership)
	d.AddResource("beta", "gkemulticloud", dcl.TitleToSnakeCase("AwsCluster"), gkemulticloud_beta.YAML_aws_cluster)
	d.AddResource("beta", "gkemulticloud", "AwsCluster", gkemulticloud_beta.YAML_aws_cluster)
	d.AddResource("beta", "gkemulticloud", dcl.TitleToSnakeCase("AwsNodePool"), gkemulticloud_beta.YAML_aws_node_pool)
	d.AddResource("beta", "gkemulticloud", "AwsNodePool", gkemulticloud_beta.YAML_aws_node_pool)
	d.AddResource("beta", "gkemulticloud", dcl.TitleToSnakeCase("AzureClient"), gkemulticloud_beta.YAML_azure_client)
	d.AddResource("beta", "gkemulticloud", "AzureClient", gkemulticloud_beta.YAML_azure_client)
	d.AddResource("beta", "gkemulticloud", dcl.TitleToSnakeCase("AzureCluster"), gkemulticloud_beta.YAML_azure_cluster)
	d.AddResource("beta", "gkemulticloud", "AzureCluster", gkemulticloud_beta.YAML_azure_cluster)
	d.AddResource("beta", "gkemulticloud", dcl.TitleToSnakeCase("AzureNodePool"), gkemulticloud_beta.YAML_azure_node_pool)
	d.AddResource("beta", "gkemulticloud", "AzureNodePool", gkemulticloud_beta.YAML_azure_node_pool)
	d.AddResource("beta", "networksecurity", dcl.TitleToSnakeCase("AuthorizationPolicy"), networksecurity_beta.YAML_authorization_policy)
	d.AddResource("beta", "networksecurity", "AuthorizationPolicy", networksecurity_beta.YAML_authorization_policy)
	d.AddResource("beta", "networksecurity", dcl.TitleToSnakeCase("ClientTlsPolicy"), networksecurity_beta.YAML_client_tls_policy)
	d.AddResource("beta", "networksecurity", "ClientTlsPolicy", networksecurity_beta.YAML_client_tls_policy)
	d.AddResource("beta", "networksecurity", dcl.TitleToSnakeCase("ServerTlsPolicy"), networksecurity_beta.YAML_server_tls_policy)
	d.AddResource("beta", "networksecurity", "ServerTlsPolicy", networksecurity_beta.YAML_server_tls_policy)
	d.AddResource("beta", "networkservices", dcl.TitleToSnakeCase("EndpointPolicy"), networkservices_beta.YAML_endpoint_policy)
	d.AddResource("beta", "networkservices", "EndpointPolicy", networkservices_beta.YAML_endpoint_policy)
	d.AddResource("beta", "osconfig", dcl.TitleToSnakeCase("GuestPolicy"), osconfig_beta.YAML_guest_policy)
	d.AddResource("beta", "osconfig", "GuestPolicy", osconfig_beta.YAML_guest_policy)
	d.AddResource("beta", "privateca", dcl.TitleToSnakeCase("CaPool"), privateca_beta.YAML_ca_pool)
	d.AddResource("beta", "privateca", "CaPool", privateca_beta.YAML_ca_pool)
	d.AddResource("beta", "privateca", dcl.TitleToSnakeCase("CertificateAuthority"), privateca_beta.YAML_certificate_authority)
	d.AddResource("beta", "privateca", "CertificateAuthority", privateca_beta.YAML_certificate_authority)
	d.AddResource("beta", "redis", dcl.TitleToSnakeCase("Instance"), redis_beta.YAML_instance)
	d.AddResource("beta", "redis", "Instance", redis_beta.YAML_instance)
	d.AddResource("beta", "runtimeconfig", dcl.TitleToSnakeCase("Config"), runtimeconfig_beta.YAML_config)
	d.AddResource("beta", "runtimeconfig", "Config", runtimeconfig_beta.YAML_config)
	d.AddResource("beta", "runtimeconfig", dcl.TitleToSnakeCase("Variable"), runtimeconfig_beta.YAML_variable)
	d.AddResource("beta", "runtimeconfig", "Variable", runtimeconfig_beta.YAML_variable)
	d.AddResource("beta", "sql", dcl.TitleToSnakeCase("Database"), sql_beta.YAML_database)
	d.AddResource("beta", "sql", "Database", sql_beta.YAML_database)
	d.AddResource("beta", "sql", dcl.TitleToSnakeCase("Instance"), sql_beta.YAML_instance)
	d.AddResource("beta", "sql", "Instance", sql_beta.YAML_instance)
	d.AddResource("beta", "sql", dcl.TitleToSnakeCase("SslCert"), sql_beta.YAML_ssl_cert)
	d.AddResource("beta", "sql", "SslCert", sql_beta.YAML_ssl_cert)
	d.AddResource("beta", "sql", dcl.TitleToSnakeCase("User"), sql_beta.YAML_user)
	d.AddResource("beta", "sql", "User", sql_beta.YAML_user)
	d.AddResource("ga", "accesscontextmanager", dcl.TitleToSnakeCase("AccessLevel"), accesscontextmanager.YAML_access_level)
	d.AddResource("ga", "accesscontextmanager", "AccessLevel", accesscontextmanager.YAML_access_level)
	d.AddResource("ga", "accesscontextmanager", dcl.TitleToSnakeCase("AccessPolicy"), accesscontextmanager.YAML_access_policy)
	d.AddResource("ga", "accesscontextmanager", "AccessPolicy", accesscontextmanager.YAML_access_policy)
	d.AddResource("ga", "accesscontextmanager", dcl.TitleToSnakeCase("ServicePerimeter"), accesscontextmanager.YAML_service_perimeter)
	d.AddResource("ga", "accesscontextmanager", "ServicePerimeter", accesscontextmanager.YAML_service_perimeter)
	d.AddResource("ga", "apigee", dcl.TitleToSnakeCase("Attachment"), apigee.YAML_attachment)
	d.AddResource("ga", "apigee", "Attachment", apigee.YAML_attachment)
	d.AddResource("ga", "apigee", dcl.TitleToSnakeCase("Envgroup"), apigee.YAML_envgroup)
	d.AddResource("ga", "apigee", "Envgroup", apigee.YAML_envgroup)
	d.AddResource("ga", "apigee", dcl.TitleToSnakeCase("Environment"), apigee.YAML_environment)
	d.AddResource("ga", "apigee", "Environment", apigee.YAML_environment)
	d.AddResource("ga", "apigee", dcl.TitleToSnakeCase("Organization"), apigee.YAML_organization)
	d.AddResource("ga", "apigee", "Organization", apigee.YAML_organization)
	d.AddResource("ga", "apikeys", dcl.TitleToSnakeCase("Key"), apikeys.YAML_key)
	d.AddResource("ga", "apikeys", "Key", apikeys.YAML_key)
	d.AddResource("ga", "appengine", dcl.TitleToSnakeCase("Application"), appengine.YAML_application)
	d.AddResource("ga", "appengine", "Application", appengine.YAML_application)
	d.AddResource("ga", "appengine", dcl.TitleToSnakeCase("DomainMapping"), appengine.YAML_domain_mapping)
	d.AddResource("ga", "appengine", "DomainMapping", appengine.YAML_domain_mapping)
	d.AddResource("ga", "appengine", dcl.TitleToSnakeCase("FirewallRule"), appengine.YAML_firewall_rule)
	d.AddResource("ga", "appengine", "FirewallRule", appengine.YAML_firewall_rule)
	d.AddResource("ga", "appengine", dcl.TitleToSnakeCase("Version"), appengine.YAML_version)
	d.AddResource("ga", "appengine", "Version", appengine.YAML_version)
	d.AddResource("ga", "assuredworkloads", dcl.TitleToSnakeCase("Workload"), assuredworkloads.YAML_workload)
	d.AddResource("ga", "assuredworkloads", "Workload", assuredworkloads.YAML_workload)
	d.AddResource("ga", "bigquery", dcl.TitleToSnakeCase("Dataset"), bigquery.YAML_dataset)
	d.AddResource("ga", "bigquery", "Dataset", bigquery.YAML_dataset)
	d.AddResource("ga", "bigqueryconnection", dcl.TitleToSnakeCase("Connection"), bigqueryconnection.YAML_connection)
	d.AddResource("ga", "bigqueryconnection", "Connection", bigqueryconnection.YAML_connection)
	d.AddResource("ga", "bigqueryreservation", dcl.TitleToSnakeCase("Assignment"), bigqueryreservation.YAML_assignment)
	d.AddResource("ga", "bigqueryreservation", "Assignment", bigqueryreservation.YAML_assignment)
	d.AddResource("ga", "bigqueryreservation", dcl.TitleToSnakeCase("Reservation"), bigqueryreservation.YAML_reservation)
	d.AddResource("ga", "bigqueryreservation", "Reservation", bigqueryreservation.YAML_reservation)
	d.AddResource("ga", "binaryauthorization", dcl.TitleToSnakeCase("Attestor"), binaryauthorization.YAML_attestor)
	d.AddResource("ga", "binaryauthorization", "Attestor", binaryauthorization.YAML_attestor)
	d.AddResource("ga", "binaryauthorization", dcl.TitleToSnakeCase("Policy"), binaryauthorization.YAML_policy)
	d.AddResource("ga", "binaryauthorization", "Policy", binaryauthorization.YAML_policy)
	d.AddResource("ga", "cloudbilling", dcl.TitleToSnakeCase("ProjectBillingInfo"), cloudbilling.YAML_project_billing_info)
	d.AddResource("ga", "cloudbilling", "ProjectBillingInfo", cloudbilling.YAML_project_billing_info)
	d.AddResource("ga", "cloudbuild", dcl.TitleToSnakeCase("BuildTrigger"), cloudbuild.YAML_build_trigger)
	d.AddResource("ga", "cloudbuild", "BuildTrigger", cloudbuild.YAML_build_trigger)
	d.AddResource("ga", "cloudfunctions", dcl.TitleToSnakeCase("Function"), cloudfunctions.YAML_function)
	d.AddResource("ga", "cloudfunctions", "Function", cloudfunctions.YAML_function)
	d.AddResource("ga", "cloudresourcemanager", dcl.TitleToSnakeCase("Folder"), cloudresourcemanager.YAML_folder)
	d.AddResource("ga", "cloudresourcemanager", "Folder", cloudresourcemanager.YAML_folder)
	d.AddResource("ga", "cloudresourcemanager", dcl.TitleToSnakeCase("Project"), cloudresourcemanager.YAML_project)
	d.AddResource("ga", "cloudresourcemanager", "Project", cloudresourcemanager.YAML_project)
	d.AddResource("ga", "cloudscheduler", dcl.TitleToSnakeCase("Job"), cloudscheduler.YAML_job)
	d.AddResource("ga", "cloudscheduler", "Job", cloudscheduler.YAML_job)
	d.AddResource("ga", "composer", dcl.TitleToSnakeCase("Environment"), composer.YAML_environment)
	d.AddResource("ga", "composer", "Environment", composer.YAML_environment)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Address"), compute.YAML_address)
	d.AddResource("ga", "compute", "Address", compute.YAML_address)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Autoscaler"), compute.YAML_autoscaler)
	d.AddResource("ga", "compute", "Autoscaler", compute.YAML_autoscaler)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("BackendBucket"), compute.YAML_backend_bucket)
	d.AddResource("ga", "compute", "BackendBucket", compute.YAML_backend_bucket)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("BackendService"), compute.YAML_backend_service)
	d.AddResource("ga", "compute", "BackendService", compute.YAML_backend_service)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Disk"), compute.YAML_disk)
	d.AddResource("ga", "compute", "Disk", compute.YAML_disk)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Firewall"), compute.YAML_firewall)
	d.AddResource("ga", "compute", "Firewall", compute.YAML_firewall)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("FirewallPolicy"), compute.YAML_firewall_policy)
	d.AddResource("ga", "compute", "FirewallPolicy", compute.YAML_firewall_policy)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("FirewallPolicyAssociation"), compute.YAML_firewall_policy_association)
	d.AddResource("ga", "compute", "FirewallPolicyAssociation", compute.YAML_firewall_policy_association)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("FirewallPolicyRule"), compute.YAML_firewall_policy_rule)
	d.AddResource("ga", "compute", "FirewallPolicyRule", compute.YAML_firewall_policy_rule)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("ForwardingRule"), compute.YAML_forwarding_rule)
	d.AddResource("ga", "compute", "ForwardingRule", compute.YAML_forwarding_rule)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("HealthCheck"), compute.YAML_health_check)
	d.AddResource("ga", "compute", "HealthCheck", compute.YAML_health_check)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("HttpHealthCheck"), compute.YAML_http_health_check)
	d.AddResource("ga", "compute", "HttpHealthCheck", compute.YAML_http_health_check)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("HttpsHealthCheck"), compute.YAML_https_health_check)
	d.AddResource("ga", "compute", "HttpsHealthCheck", compute.YAML_https_health_check)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Image"), compute.YAML_image)
	d.AddResource("ga", "compute", "Image", compute.YAML_image)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Instance"), compute.YAML_instance)
	d.AddResource("ga", "compute", "Instance", compute.YAML_instance)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("InstanceGroupManager"), compute.YAML_instance_group_manager)
	d.AddResource("ga", "compute", "InstanceGroupManager", compute.YAML_instance_group_manager)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("InstanceTemplate"), compute.YAML_instance_template)
	d.AddResource("ga", "compute", "InstanceTemplate", compute.YAML_instance_template)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Interconnect"), compute.YAML_interconnect)
	d.AddResource("ga", "compute", "Interconnect", compute.YAML_interconnect)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("ManagedSslCertificate"), compute.YAML_managed_ssl_certificate)
	d.AddResource("ga", "compute", "ManagedSslCertificate", compute.YAML_managed_ssl_certificate)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Network"), compute.YAML_network)
	d.AddResource("ga", "compute", "Network", compute.YAML_network)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("NetworkEndpoint"), compute.YAML_network_endpoint)
	d.AddResource("ga", "compute", "NetworkEndpoint", compute.YAML_network_endpoint)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("NetworkEndpointGroup"), compute.YAML_network_endpoint_group)
	d.AddResource("ga", "compute", "NetworkEndpointGroup", compute.YAML_network_endpoint_group)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Reservation"), compute.YAML_reservation)
	d.AddResource("ga", "compute", "Reservation", compute.YAML_reservation)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Route"), compute.YAML_route)
	d.AddResource("ga", "compute", "Route", compute.YAML_route)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Router"), compute.YAML_router)
	d.AddResource("ga", "compute", "Router", compute.YAML_router)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("RouterPeer"), compute.YAML_router_peer)
	d.AddResource("ga", "compute", "RouterPeer", compute.YAML_router_peer)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Snapshot"), compute.YAML_snapshot)
	d.AddResource("ga", "compute", "Snapshot", compute.YAML_snapshot)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("SslCertificate"), compute.YAML_ssl_certificate)
	d.AddResource("ga", "compute", "SslCertificate", compute.YAML_ssl_certificate)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("SslPolicy"), compute.YAML_ssl_policy)
	d.AddResource("ga", "compute", "SslPolicy", compute.YAML_ssl_policy)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Subnetwork"), compute.YAML_subnetwork)
	d.AddResource("ga", "compute", "Subnetwork", compute.YAML_subnetwork)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("TargetHttpProxy"), compute.YAML_target_http_proxy)
	d.AddResource("ga", "compute", "TargetHttpProxy", compute.YAML_target_http_proxy)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("TargetPool"), compute.YAML_target_pool)
	d.AddResource("ga", "compute", "TargetPool", compute.YAML_target_pool)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("TargetSslProxy"), compute.YAML_target_ssl_proxy)
	d.AddResource("ga", "compute", "TargetSslProxy", compute.YAML_target_ssl_proxy)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("TargetVpnGateway"), compute.YAML_target_vpn_gateway)
	d.AddResource("ga", "compute", "TargetVpnGateway", compute.YAML_target_vpn_gateway)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("UrlMap"), compute.YAML_url_map)
	d.AddResource("ga", "compute", "UrlMap", compute.YAML_url_map)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("VpnGateway"), compute.YAML_vpn_gateway)
	d.AddResource("ga", "compute", "VpnGateway", compute.YAML_vpn_gateway)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("VpnTunnel"), compute.YAML_vpn_tunnel)
	d.AddResource("ga", "compute", "VpnTunnel", compute.YAML_vpn_tunnel)
	d.AddResource("ga", "container", dcl.TitleToSnakeCase("Cluster"), container.YAML_cluster)
	d.AddResource("ga", "container", "Cluster", container.YAML_cluster)
	d.AddResource("ga", "container", dcl.TitleToSnakeCase("NodePool"), container.YAML_node_pool)
	d.AddResource("ga", "container", "NodePool", container.YAML_node_pool)
	d.AddResource("ga", "containeranalysis", dcl.TitleToSnakeCase("Note"), containeranalysis.YAML_note)
	d.AddResource("ga", "containeranalysis", "Note", containeranalysis.YAML_note)
	d.AddResource("ga", "dataproc", dcl.TitleToSnakeCase("AutoscalingPolicy"), dataproc.YAML_autoscaling_policy)
	d.AddResource("ga", "dataproc", "AutoscalingPolicy", dataproc.YAML_autoscaling_policy)
	d.AddResource("ga", "dataproc", dcl.TitleToSnakeCase("Cluster"), dataproc.YAML_cluster)
	d.AddResource("ga", "dataproc", "Cluster", dataproc.YAML_cluster)
	d.AddResource("ga", "dataproc", dcl.TitleToSnakeCase("Job"), dataproc.YAML_job)
	d.AddResource("ga", "dataproc", "Job", dataproc.YAML_job)
	d.AddResource("ga", "dataproc", dcl.TitleToSnakeCase("WorkflowTemplate"), dataproc.YAML_workflow_template)
	d.AddResource("ga", "dataproc", "WorkflowTemplate", dataproc.YAML_workflow_template)
	d.AddResource("ga", "datastore", dcl.TitleToSnakeCase("Index"), datastore.YAML_index)
	d.AddResource("ga", "datastore", "Index", datastore.YAML_index)
	d.AddResource("ga", "dns", dcl.TitleToSnakeCase("ManagedZone"), dns.YAML_managed_zone)
	d.AddResource("ga", "dns", "ManagedZone", dns.YAML_managed_zone)
	d.AddResource("ga", "dns", dcl.TitleToSnakeCase("ResourceRecordSet"), dns.YAML_resource_record_set)
	d.AddResource("ga", "dns", "ResourceRecordSet", dns.YAML_resource_record_set)
	d.AddResource("ga", "file", dcl.TitleToSnakeCase("Instance"), file.YAML_instance)
	d.AddResource("ga", "file", "Instance", file.YAML_instance)
	d.AddResource("ga", "eventarc", dcl.TitleToSnakeCase("Trigger"), eventarc.YAML_trigger)
	d.AddResource("ga", "eventarc", "Trigger", eventarc.YAML_trigger)
	d.AddResource("ga", "gameservices", dcl.TitleToSnakeCase("Realm"), gameservices.YAML_realm)
	d.AddResource("ga", "gameservices", "Realm", gameservices.YAML_realm)
	d.AddResource("ga", "gkemulticloud", dcl.TitleToSnakeCase("AwsCluster"), gkemulticloud.YAML_aws_cluster)
	d.AddResource("ga", "gkemulticloud", "AwsCluster", gkemulticloud.YAML_aws_cluster)
	d.AddResource("ga", "gkemulticloud", dcl.TitleToSnakeCase("AwsNodePool"), gkemulticloud.YAML_aws_node_pool)
	d.AddResource("ga", "gkemulticloud", "AwsNodePool", gkemulticloud.YAML_aws_node_pool)
	d.AddResource("ga", "gkemulticloud", dcl.TitleToSnakeCase("AzureClient"), gkemulticloud.YAML_azure_client)
	d.AddResource("ga", "gkemulticloud", "AzureClient", gkemulticloud.YAML_azure_client)
	d.AddResource("ga", "gkemulticloud", dcl.TitleToSnakeCase("AzureCluster"), gkemulticloud.YAML_azure_cluster)
	d.AddResource("ga", "gkemulticloud", "AzureCluster", gkemulticloud.YAML_azure_cluster)
	d.AddResource("ga", "gkemulticloud", dcl.TitleToSnakeCase("AzureNodePool"), gkemulticloud.YAML_azure_node_pool)
	d.AddResource("ga", "gkemulticloud", "AzureNodePool", gkemulticloud.YAML_azure_node_pool)
	d.AddResource("ga", "logging", dcl.TitleToSnakeCase("LogBucket"), logging.YAML_log_bucket)
	d.AddResource("ga", "logging", "LogBucket", logging.YAML_log_bucket)
	d.AddResource("ga", "logging", dcl.TitleToSnakeCase("LogExclusion"), logging.YAML_log_exclusion)
	d.AddResource("ga", "logging", "LogExclusion", logging.YAML_log_exclusion)
	d.AddResource("ga", "logging", dcl.TitleToSnakeCase("LogMetric"), logging.YAML_log_metric)
	d.AddResource("ga", "logging", "LogMetric", logging.YAML_log_metric)
	d.AddResource("ga", "iam", dcl.TitleToSnakeCase("Role"), iam.YAML_role)
	d.AddResource("ga", "iam", "Role", iam.YAML_role)
	d.AddResource("ga", "iam", dcl.TitleToSnakeCase("ServiceAccount"), iam.YAML_service_account)
	d.AddResource("ga", "iam", "ServiceAccount", iam.YAML_service_account)
	d.AddResource("ga", "iap", dcl.TitleToSnakeCase("Brand"), iap.YAML_brand)
	d.AddResource("ga", "iap", "Brand", iap.YAML_brand)
	d.AddResource("ga", "iap", dcl.TitleToSnakeCase("IdentityAwareProxyClient"), iap.YAML_identity_aware_proxy_client)
	d.AddResource("ga", "iap", "IdentityAwareProxyClient", iap.YAML_identity_aware_proxy_client)
	d.AddResource("ga", "identitytoolkit", dcl.TitleToSnakeCase("OAuthIdpConfig"), identitytoolkit.YAML_oauth_idp_config)
	d.AddResource("ga", "identitytoolkit", "OAuthIdpConfig", identitytoolkit.YAML_oauth_idp_config)
	d.AddResource("ga", "identitytoolkit", dcl.TitleToSnakeCase("Tenant"), identitytoolkit.YAML_tenant)
	d.AddResource("ga", "identitytoolkit", "Tenant", identitytoolkit.YAML_tenant)
	d.AddResource("ga", "identitytoolkit", dcl.TitleToSnakeCase("TenantOAuthIdpConfig"), identitytoolkit.YAML_tenant_oauth_idp_config)
	d.AddResource("ga", "identitytoolkit", "TenantOAuthIdpConfig", identitytoolkit.YAML_tenant_oauth_idp_config)
	d.AddResource("ga", "monitoring", dcl.TitleToSnakeCase("AlertPolicy"), monitoring.YAML_alert_policy)
	d.AddResource("ga", "monitoring", "AlertPolicy", monitoring.YAML_alert_policy)
	d.AddResource("ga", "monitoring", dcl.TitleToSnakeCase("Dashboard"), monitoring.YAML_dashboard)
	d.AddResource("ga", "monitoring", "Dashboard", monitoring.YAML_dashboard)
	d.AddResource("ga", "monitoring", dcl.TitleToSnakeCase("Group"), monitoring.YAML_group)
	d.AddResource("ga", "monitoring", "Group", monitoring.YAML_group)
	d.AddResource("ga", "monitoring", dcl.TitleToSnakeCase("MetricDescriptor"), monitoring.YAML_metric_descriptor)
	d.AddResource("ga", "monitoring", "MetricDescriptor", monitoring.YAML_metric_descriptor)
	d.AddResource("ga", "monitoring", dcl.TitleToSnakeCase("NotificationChannel"), monitoring.YAML_notification_channel)
	d.AddResource("ga", "monitoring", "NotificationChannel", monitoring.YAML_notification_channel)
	d.AddResource("ga", "monitoring", dcl.TitleToSnakeCase("UptimeCheckConfig"), monitoring.YAML_uptime_check_config)
	d.AddResource("ga", "monitoring", "UptimeCheckConfig", monitoring.YAML_uptime_check_config)
	d.AddResource("ga", "pubsub", dcl.TitleToSnakeCase("Subscription"), pubsub.YAML_subscription)
	d.AddResource("ga", "pubsub", "Subscription", pubsub.YAML_subscription)
	d.AddResource("ga", "pubsub", dcl.TitleToSnakeCase("Topic"), pubsub.YAML_topic)
	d.AddResource("ga", "pubsub", "Topic", pubsub.YAML_topic)
	d.AddResource("ga", "pubsublite", dcl.TitleToSnakeCase("Subscription"), pubsublite.YAML_subscription)
	d.AddResource("ga", "pubsublite", "Subscription", pubsublite.YAML_subscription)
	d.AddResource("ga", "pubsublite", dcl.TitleToSnakeCase("Topic"), pubsublite.YAML_topic)
	d.AddResource("ga", "pubsublite", "Topic", pubsublite.YAML_topic)
	d.AddResource("ga", "redis", dcl.TitleToSnakeCase("Instance"), redis.YAML_instance)
	d.AddResource("ga", "redis", "Instance", redis.YAML_instance)
	d.AddResource("ga", "run", dcl.TitleToSnakeCase("Service"), run.YAML_service)
	d.AddResource("ga", "run", "Service", run.YAML_service)
	d.AddResource("ga", "servicemanagement", dcl.TitleToSnakeCase("Service"), servicemanagement.YAML_service)
	d.AddResource("ga", "servicemanagement", "Service", servicemanagement.YAML_service)
	d.AddResource("ga", "servicenetworking", dcl.TitleToSnakeCase("Connection"), servicenetworking.YAML_connection)
	d.AddResource("ga", "servicenetworking", "Connection", servicenetworking.YAML_connection)
	d.AddResource("ga", "sourcerepo", dcl.TitleToSnakeCase("Repo"), sourcerepo.YAML_repo)
	d.AddResource("ga", "sourcerepo", "Repo", sourcerepo.YAML_repo)
	d.AddResource("ga", "serviceusage", dcl.TitleToSnakeCase("Service"), serviceusage.YAML_service)
	d.AddResource("ga", "serviceusage", "Service", serviceusage.YAML_service)
	d.AddResource("ga", "spanner", dcl.TitleToSnakeCase("Database"), spanner.YAML_database)
	d.AddResource("ga", "spanner", "Database", spanner.YAML_database)
	d.AddResource("ga", "spanner", dcl.TitleToSnakeCase("Instance"), spanner.YAML_instance)
	d.AddResource("ga", "spanner", "Instance", spanner.YAML_instance)
	d.AddResource("ga", "storage", dcl.TitleToSnakeCase("Bucket"), storage.YAML_bucket)
	d.AddResource("ga", "storage", "Bucket", storage.YAML_bucket)
	d.AddResource("ga", "storage", dcl.TitleToSnakeCase("DefaultObjectAccessControl"), storage.YAML_default_object_access_control)
	d.AddResource("ga", "storage", "DefaultObjectAccessControl", storage.YAML_default_object_access_control)
	d.AddResource("ga", "storage", dcl.TitleToSnakeCase("HmacKey"), storage.YAML_hmac_key)
	d.AddResource("ga", "storage", "HmacKey", storage.YAML_hmac_key)
	d.AddResource("ga", "storage", dcl.TitleToSnakeCase("Object"), storage.YAML_object)
	d.AddResource("ga", "storage", "Object", storage.YAML_object)
	d.AddResource("ga", "storage", dcl.TitleToSnakeCase("ObjectAccessControl"), storage.YAML_object_access_control)
	d.AddResource("ga", "storage", "ObjectAccessControl", storage.YAML_object_access_control)
	d.AddResource("ga", "tpu", dcl.TitleToSnakeCase("Node"), tpu.YAML_node)
	d.AddResource("ga", "tpu", "Node", tpu.YAML_node)
	d.AddResource("ga", "vpcaccess", dcl.TitleToSnakeCase("Connector"), vpcaccess.YAML_connector)
	d.AddResource("ga", "vpcaccess", "Connector", vpcaccess.YAML_connector)
	d.AddResource("alpha", "krmapihosting", dcl.TitleToSnakeCase("KrmApiHost"), krmapihosting_alpha.YAML_krm_api_host)
	d.AddResource("alpha", "krmapihosting", "KrmApiHost", krmapihosting_alpha.YAML_krm_api_host)
	d.AddResource("alpha", "networksecurity", dcl.TitleToSnakeCase("AuthorizationPolicy"), networksecurity_alpha.YAML_authorization_policy)
	d.AddResource("alpha", "networksecurity", "AuthorizationPolicy", networksecurity_alpha.YAML_authorization_policy)
	d.AddResource("alpha", "networksecurity", dcl.TitleToSnakeCase("ClientTlsPolicy"), networksecurity_alpha.YAML_client_tls_policy)
	d.AddResource("alpha", "networksecurity", "ClientTlsPolicy", networksecurity_alpha.YAML_client_tls_policy)
	d.AddResource("alpha", "networksecurity", dcl.TitleToSnakeCase("ServerTlsPolicy"), networksecurity_alpha.YAML_server_tls_policy)
	d.AddResource("alpha", "networksecurity", "ServerTlsPolicy", networksecurity_alpha.YAML_server_tls_policy)
	d.AddResource("alpha", "networkservices", dcl.TitleToSnakeCase("EndpointConfigSelector"), networkservices_alpha.YAML_endpoint_config_selector)
	d.AddResource("alpha", "networkservices", "EndpointConfigSelector", networkservices_alpha.YAML_endpoint_config_selector)
	d.AddResource("alpha", "networkservices", dcl.TitleToSnakeCase("HttpFilter"), networkservices_alpha.YAML_http_filter)
	d.AddResource("alpha", "networkservices", "HttpFilter", networkservices_alpha.YAML_http_filter)
	d.AddResource("alpha", "osconfig", dcl.TitleToSnakeCase("OSPolicyAssignment"), osconfig_alpha.YAML_os_policy_assignment)
	d.AddResource("alpha", "osconfig", "OSPolicyAssignment", osconfig_alpha.YAML_os_policy_assignment)
	d.AddResource("alpha", "tier2", dcl.TitleToSnakeCase("Instance"), tier2_alpha.YAML_instance)
	d.AddResource("alpha", "tier2", "Instance", tier2_alpha.YAML_instance)
	return d
}
