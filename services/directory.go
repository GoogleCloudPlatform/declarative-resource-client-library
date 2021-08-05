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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads"
	assuredworkloads_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation"
	bigqueryreservation_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization"
	binaryauthorization_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/beta"
	cloudbuild_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler"
	cloudscheduler_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
	compute_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis"
	containeranalysis_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/beta"
	datafusion_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datafusion/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc"
	dataproc_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc"
	eventarc_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/file"
	file_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/file/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gameservices"
	gameservices_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gameservices/beta"
	gkehub_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud"
	gkemulticloud_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	iam_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap"
	iap_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit"
	identitytoolkit_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging"
	logging_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
	monitoring_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/beta"
	networksecurity_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/alpha"
	networksecurity_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/beta"
	networkservices_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/alpha"
	networkservices_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/beta"
	osconfig_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/alpha"
	osconfig_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/beta"
	tier2_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/tier2/alpha"
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
	d.AddResource("ga", "assuredworkloads", dcl.TitleToSnakeCase("Workload"), assuredworkloads.YAML_workload)
	d.AddResource("ga", "assuredworkloads", "Workload", assuredworkloads.YAML_workload)
	d.AddResource("ga", "bigqueryreservation", dcl.TitleToSnakeCase("Assignment"), bigqueryreservation.YAML_assignment)
	d.AddResource("ga", "bigqueryreservation", "Assignment", bigqueryreservation.YAML_assignment)
	d.AddResource("ga", "binaryauthorization", dcl.TitleToSnakeCase("Attestor"), binaryauthorization.YAML_attestor)
	d.AddResource("ga", "binaryauthorization", "Attestor", binaryauthorization.YAML_attestor)
	d.AddResource("ga", "binaryauthorization", dcl.TitleToSnakeCase("Policy"), binaryauthorization.YAML_policy)
	d.AddResource("ga", "binaryauthorization", "Policy", binaryauthorization.YAML_policy)
	d.AddResource("ga", "cloudscheduler", dcl.TitleToSnakeCase("Job"), cloudscheduler.YAML_job)
	d.AddResource("ga", "cloudscheduler", "Job", cloudscheduler.YAML_job)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("FirewallPolicy"), compute.YAML_firewall_policy)
	d.AddResource("ga", "compute", "FirewallPolicy", compute.YAML_firewall_policy)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("FirewallPolicyAssociation"), compute.YAML_firewall_policy_association)
	d.AddResource("ga", "compute", "FirewallPolicyAssociation", compute.YAML_firewall_policy_association)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("FirewallPolicyRule"), compute.YAML_firewall_policy_rule)
	d.AddResource("ga", "compute", "FirewallPolicyRule", compute.YAML_firewall_policy_rule)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("ForwardingRule"), compute.YAML_forwarding_rule)
	d.AddResource("ga", "compute", "ForwardingRule", compute.YAML_forwarding_rule)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("InstanceGroupManager"), compute.YAML_instance_group_manager)
	d.AddResource("ga", "compute", "InstanceGroupManager", compute.YAML_instance_group_manager)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Route"), compute.YAML_route)
	d.AddResource("ga", "compute", "Route", compute.YAML_route)
	d.AddResource("ga", "containeranalysis", dcl.TitleToSnakeCase("Note"), containeranalysis.YAML_note)
	d.AddResource("ga", "containeranalysis", "Note", containeranalysis.YAML_note)
	d.AddResource("ga", "dataproc", dcl.TitleToSnakeCase("AutoscalingPolicy"), dataproc.YAML_autoscaling_policy)
	d.AddResource("ga", "dataproc", "AutoscalingPolicy", dataproc.YAML_autoscaling_policy)
	d.AddResource("ga", "dataproc", dcl.TitleToSnakeCase("Cluster"), dataproc.YAML_cluster)
	d.AddResource("ga", "dataproc", "Cluster", dataproc.YAML_cluster)
	d.AddResource("ga", "dataproc", dcl.TitleToSnakeCase("WorkflowTemplate"), dataproc.YAML_workflow_template)
	d.AddResource("ga", "dataproc", "WorkflowTemplate", dataproc.YAML_workflow_template)
	d.AddResource("ga", "eventarc", dcl.TitleToSnakeCase("Trigger"), eventarc.YAML_trigger)
	d.AddResource("ga", "eventarc", "Trigger", eventarc.YAML_trigger)
	d.AddResource("ga", "file", dcl.TitleToSnakeCase("Instance"), file.YAML_instance)
	d.AddResource("ga", "file", "Instance", file.YAML_instance)
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
	d.AddResource("ga", "logging", dcl.TitleToSnakeCase("LogExclusion"), logging.YAML_log_exclusion)
	d.AddResource("ga", "logging", "LogExclusion", logging.YAML_log_exclusion)
	d.AddResource("ga", "monitoring", dcl.TitleToSnakeCase("Dashboard"), monitoring.YAML_dashboard)
	d.AddResource("ga", "monitoring", "Dashboard", monitoring.YAML_dashboard)
	d.AddResource("ga", "monitoring", dcl.TitleToSnakeCase("Group"), monitoring.YAML_group)
	d.AddResource("ga", "monitoring", "Group", monitoring.YAML_group)
	d.AddResource("beta", "assuredworkloads", dcl.TitleToSnakeCase("Workload"), assuredworkloads_beta.YAML_workload)
	d.AddResource("beta", "assuredworkloads", "Workload", assuredworkloads_beta.YAML_workload)
	d.AddResource("beta", "bigqueryreservation", dcl.TitleToSnakeCase("Assignment"), bigqueryreservation_beta.YAML_assignment)
	d.AddResource("beta", "bigqueryreservation", "Assignment", bigqueryreservation_beta.YAML_assignment)
	d.AddResource("beta", "binaryauthorization", dcl.TitleToSnakeCase("Attestor"), binaryauthorization_beta.YAML_attestor)
	d.AddResource("beta", "binaryauthorization", "Attestor", binaryauthorization_beta.YAML_attestor)
	d.AddResource("beta", "binaryauthorization", dcl.TitleToSnakeCase("Policy"), binaryauthorization_beta.YAML_policy)
	d.AddResource("beta", "binaryauthorization", "Policy", binaryauthorization_beta.YAML_policy)
	d.AddResource("beta", "cloudbuild", dcl.TitleToSnakeCase("WorkerPool"), cloudbuild_beta.YAML_worker_pool)
	d.AddResource("beta", "cloudbuild", "WorkerPool", cloudbuild_beta.YAML_worker_pool)
	d.AddResource("beta", "cloudscheduler", dcl.TitleToSnakeCase("Job"), cloudscheduler_beta.YAML_job)
	d.AddResource("beta", "cloudscheduler", "Job", cloudscheduler_beta.YAML_job)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("FirewallPolicy"), compute_beta.YAML_firewall_policy)
	d.AddResource("beta", "compute", "FirewallPolicy", compute_beta.YAML_firewall_policy)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("FirewallPolicyAssociation"), compute_beta.YAML_firewall_policy_association)
	d.AddResource("beta", "compute", "FirewallPolicyAssociation", compute_beta.YAML_firewall_policy_association)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("FirewallPolicyRule"), compute_beta.YAML_firewall_policy_rule)
	d.AddResource("beta", "compute", "FirewallPolicyRule", compute_beta.YAML_firewall_policy_rule)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("ForwardingRule"), compute_beta.YAML_forwarding_rule)
	d.AddResource("beta", "compute", "ForwardingRule", compute_beta.YAML_forwarding_rule)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("InstanceGroupManager"), compute_beta.YAML_instance_group_manager)
	d.AddResource("beta", "compute", "InstanceGroupManager", compute_beta.YAML_instance_group_manager)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Route"), compute_beta.YAML_route)
	d.AddResource("beta", "compute", "Route", compute_beta.YAML_route)
	d.AddResource("beta", "containeranalysis", dcl.TitleToSnakeCase("Note"), containeranalysis_beta.YAML_note)
	d.AddResource("beta", "containeranalysis", "Note", containeranalysis_beta.YAML_note)
	d.AddResource("beta", "datafusion", dcl.TitleToSnakeCase("Instance"), datafusion_beta.YAML_instance)
	d.AddResource("beta", "datafusion", "Instance", datafusion_beta.YAML_instance)
	d.AddResource("beta", "dataproc", dcl.TitleToSnakeCase("AutoscalingPolicy"), dataproc_beta.YAML_autoscaling_policy)
	d.AddResource("beta", "dataproc", "AutoscalingPolicy", dataproc_beta.YAML_autoscaling_policy)
	d.AddResource("beta", "dataproc", dcl.TitleToSnakeCase("Cluster"), dataproc_beta.YAML_cluster)
	d.AddResource("beta", "dataproc", "Cluster", dataproc_beta.YAML_cluster)
	d.AddResource("beta", "dataproc", dcl.TitleToSnakeCase("WorkflowTemplate"), dataproc_beta.YAML_workflow_template)
	d.AddResource("beta", "dataproc", "WorkflowTemplate", dataproc_beta.YAML_workflow_template)
	d.AddResource("beta", "eventarc", dcl.TitleToSnakeCase("Trigger"), eventarc_beta.YAML_trigger)
	d.AddResource("beta", "eventarc", "Trigger", eventarc_beta.YAML_trigger)
	d.AddResource("beta", "file", dcl.TitleToSnakeCase("Backup"), file_beta.YAML_backup)
	d.AddResource("beta", "file", "Backup", file_beta.YAML_backup)
	d.AddResource("beta", "file", dcl.TitleToSnakeCase("Instance"), file_beta.YAML_instance)
	d.AddResource("beta", "file", "Instance", file_beta.YAML_instance)
	d.AddResource("beta", "gameservices", dcl.TitleToSnakeCase("Realm"), gameservices_beta.YAML_realm)
	d.AddResource("beta", "gameservices", "Realm", gameservices_beta.YAML_realm)
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
	d.AddResource("beta", "iam", dcl.TitleToSnakeCase("Role"), iam_beta.YAML_role)
	d.AddResource("beta", "iam", "Role", iam_beta.YAML_role)
	d.AddResource("beta", "iam", dcl.TitleToSnakeCase("ServiceAccount"), iam_beta.YAML_service_account)
	d.AddResource("beta", "iam", "ServiceAccount", iam_beta.YAML_service_account)
	d.AddResource("beta", "iap", dcl.TitleToSnakeCase("Brand"), iap_beta.YAML_brand)
	d.AddResource("beta", "iap", "Brand", iap_beta.YAML_brand)
	d.AddResource("beta", "iap", dcl.TitleToSnakeCase("IdentityAwareProxyClient"), iap_beta.YAML_identity_aware_proxy_client)
	d.AddResource("beta", "iap", "IdentityAwareProxyClient", iap_beta.YAML_identity_aware_proxy_client)
	d.AddResource("beta", "identitytoolkit", dcl.TitleToSnakeCase("OAuthIdpConfig"), identitytoolkit_beta.YAML_oauth_idp_config)
	d.AddResource("beta", "identitytoolkit", "OAuthIdpConfig", identitytoolkit_beta.YAML_oauth_idp_config)
	d.AddResource("beta", "identitytoolkit", dcl.TitleToSnakeCase("Tenant"), identitytoolkit_beta.YAML_tenant)
	d.AddResource("beta", "identitytoolkit", "Tenant", identitytoolkit_beta.YAML_tenant)
	d.AddResource("beta", "identitytoolkit", dcl.TitleToSnakeCase("TenantOAuthIdpConfig"), identitytoolkit_beta.YAML_tenant_oauth_idp_config)
	d.AddResource("beta", "identitytoolkit", "TenantOAuthIdpConfig", identitytoolkit_beta.YAML_tenant_oauth_idp_config)
	d.AddResource("beta", "logging", dcl.TitleToSnakeCase("LogExclusion"), logging_beta.YAML_log_exclusion)
	d.AddResource("beta", "logging", "LogExclusion", logging_beta.YAML_log_exclusion)
	d.AddResource("beta", "monitoring", dcl.TitleToSnakeCase("Dashboard"), monitoring_beta.YAML_dashboard)
	d.AddResource("beta", "monitoring", "Dashboard", monitoring_beta.YAML_dashboard)
	d.AddResource("beta", "monitoring", dcl.TitleToSnakeCase("Group"), monitoring_beta.YAML_group)
	d.AddResource("beta", "monitoring", "Group", monitoring_beta.YAML_group)
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
	d.AddResource("alpha", "networksecurity", dcl.TitleToSnakeCase("AuthorizationPolicy"), networksecurity_alpha.YAML_authorization_policy)
	d.AddResource("alpha", "networksecurity", "AuthorizationPolicy", networksecurity_alpha.YAML_authorization_policy)
	d.AddResource("alpha", "networksecurity", dcl.TitleToSnakeCase("ClientTlsPolicy"), networksecurity_alpha.YAML_client_tls_policy)
	d.AddResource("alpha", "networksecurity", "ClientTlsPolicy", networksecurity_alpha.YAML_client_tls_policy)
	d.AddResource("alpha", "networksecurity", dcl.TitleToSnakeCase("ServerTlsPolicy"), networksecurity_alpha.YAML_server_tls_policy)
	d.AddResource("alpha", "networksecurity", "ServerTlsPolicy", networksecurity_alpha.YAML_server_tls_policy)
	d.AddResource("alpha", "networkservices", dcl.TitleToSnakeCase("EndpointConfigSelector"), networkservices_alpha.YAML_endpoint_config_selector)
	d.AddResource("alpha", "networkservices", "EndpointConfigSelector", networkservices_alpha.YAML_endpoint_config_selector)
	d.AddResource("alpha", "osconfig", dcl.TitleToSnakeCase("OSPolicyAssignment"), osconfig_alpha.YAML_os_policy_assignment)
	d.AddResource("alpha", "osconfig", "OSPolicyAssignment", osconfig_alpha.YAML_os_policy_assignment)
	d.AddResource("alpha", "tier2", dcl.TitleToSnakeCase("Instance"), tier2_alpha.YAML_instance)
	d.AddResource("alpha", "tier2", "Instance", tier2_alpha.YAML_instance)
	return d
}
