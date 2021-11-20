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
	assuredworkloads_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/alpha"
	assuredworkloads_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation"
	bigqueryreservation_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/alpha"
	bigqueryreservation_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryreservation/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/billingbudgets"
	billingbudgets_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/billingbudgets/alpha"
	billingbudgets_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/billingbudgets/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization"
	binaryauthorization_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/alpha"
	binaryauthorization_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/binaryauthorization/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild"
	cloudbuild_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/alpha"
	cloudbuild_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuild/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudfunctions"
	cloudfunctions_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudfunctions/alpha"
	cloudfunctions_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudfunctions/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudidentity"
	cloudidentity_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudidentity/alpha"
	cloudidentity_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudidentity/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms"
	cloudkms_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms/alpha"
	cloudkms_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager"
	cloudresourcemanager_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager/alpha"
	cloudresourcemanager_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler"
	cloudscheduler_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/alpha"
	cloudscheduler_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudscheduler/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
	compute_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/alpha"
	compute_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
	configcontroller_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/configcontroller/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis"
	containeranalysis_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/alpha"
	containeranalysis_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeranalysis/beta"
	containeraws_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeraws/alpha"
	containerazure_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure/alpha"
	datafusion_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datafusion/alpha"
	datafusion_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datafusion/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc"
	dataproc_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/alpha"
	dataproc_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc"
	eventarc_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/alpha"
	eventarc_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore"
	filestore_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore/alpha"
	filestore_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gameservices"
	gameservices_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gameservices/alpha"
	gameservices_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gameservices/beta"
	gkehub_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/alpha"
	gkehub_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	iam_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/alpha"
	iam_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap"
	iap_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/alpha"
	iap_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iap/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit"
	identitytoolkit_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/alpha"
	identitytoolkit_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging"
	logging_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/alpha"
	logging_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
	monitoring_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/alpha"
	monitoring_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkconnectivity"
	networkconnectivity_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkconnectivity/alpha"
	networkconnectivity_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkconnectivity/beta"
	networksecurity_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/alpha"
	networksecurity_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/beta"
	networkservices_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/alpha"
	networkservices_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy"
	orgpolicy_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy/alpha"
	orgpolicy_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/orgpolicy/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig"
	osconfig_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/alpha"
	osconfig_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca"
	privateca_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/alpha"
	privateca_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub"
	pubsub_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub/alpha"
	pubsub_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/recaptchaenterprise"
	recaptchaenterprise_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/recaptchaenterprise/alpha"
	recaptchaenterprise_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/recaptchaenterprise/beta"
	run_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/run/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage"
	storage_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage/alpha"
	storage_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage/beta"
	vmware_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vmware/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vpcaccess"
	vpcaccess_alpha "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vpcaccess/alpha"
	vpcaccess_beta "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vpcaccess/beta"
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
	d.AddResource("ga", "billingbudgets", dcl.TitleToSnakeCase("Budget"), billingbudgets.YAML_budget)
	d.AddResource("ga", "billingbudgets", "Budget", billingbudgets.YAML_budget)
	d.AddResource("ga", "bigqueryreservation", dcl.TitleToSnakeCase("Assignment"), bigqueryreservation.YAML_assignment)
	d.AddResource("ga", "bigqueryreservation", "Assignment", bigqueryreservation.YAML_assignment)
	d.AddResource("ga", "bigqueryreservation", dcl.TitleToSnakeCase("CapacityCommitment"), bigqueryreservation.YAML_capacity_commitment)
	d.AddResource("ga", "bigqueryreservation", "CapacityCommitment", bigqueryreservation.YAML_capacity_commitment)
	d.AddResource("ga", "bigqueryreservation", dcl.TitleToSnakeCase("Reservation"), bigqueryreservation.YAML_reservation)
	d.AddResource("ga", "bigqueryreservation", "Reservation", bigqueryreservation.YAML_reservation)
	d.AddResource("ga", "binaryauthorization", dcl.TitleToSnakeCase("Attestor"), binaryauthorization.YAML_attestor)
	d.AddResource("ga", "binaryauthorization", "Attestor", binaryauthorization.YAML_attestor)
	d.AddResource("ga", "binaryauthorization", dcl.TitleToSnakeCase("Policy"), binaryauthorization.YAML_policy)
	d.AddResource("ga", "binaryauthorization", "Policy", binaryauthorization.YAML_policy)
	d.AddResource("ga", "cloudbuild", dcl.TitleToSnakeCase("WorkerPool"), cloudbuild.YAML_worker_pool)
	d.AddResource("ga", "cloudbuild", "WorkerPool", cloudbuild.YAML_worker_pool)
	d.AddResource("ga", "cloudfunctions", dcl.TitleToSnakeCase("Function"), cloudfunctions.YAML_function)
	d.AddResource("ga", "cloudfunctions", "Function", cloudfunctions.YAML_function)
	d.AddResource("ga", "cloudidentity", dcl.TitleToSnakeCase("Group"), cloudidentity.YAML_group)
	d.AddResource("ga", "cloudidentity", "Group", cloudidentity.YAML_group)
	d.AddResource("ga", "cloudidentity", dcl.TitleToSnakeCase("Membership"), cloudidentity.YAML_membership)
	d.AddResource("ga", "cloudidentity", "Membership", cloudidentity.YAML_membership)
	d.AddResource("ga", "cloudkms", dcl.TitleToSnakeCase("CryptoKey"), cloudkms.YAML_crypto_key)
	d.AddResource("ga", "cloudkms", "CryptoKey", cloudkms.YAML_crypto_key)
	d.AddResource("ga", "cloudkms", dcl.TitleToSnakeCase("KeyRing"), cloudkms.YAML_key_ring)
	d.AddResource("ga", "cloudkms", "KeyRing", cloudkms.YAML_key_ring)
	d.AddResource("ga", "cloudresourcemanager", dcl.TitleToSnakeCase("Folder"), cloudresourcemanager.YAML_folder)
	d.AddResource("ga", "cloudresourcemanager", "Folder", cloudresourcemanager.YAML_folder)
	d.AddResource("ga", "cloudresourcemanager", dcl.TitleToSnakeCase("Project"), cloudresourcemanager.YAML_project)
	d.AddResource("ga", "cloudresourcemanager", "Project", cloudresourcemanager.YAML_project)
	d.AddResource("ga", "cloudresourcemanager", dcl.TitleToSnakeCase("TagKey"), cloudresourcemanager.YAML_tag_key)
	d.AddResource("ga", "cloudresourcemanager", "TagKey", cloudresourcemanager.YAML_tag_key)
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
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Instance"), compute.YAML_instance)
	d.AddResource("ga", "compute", "Instance", compute.YAML_instance)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("InstanceGroupManager"), compute.YAML_instance_group_manager)
	d.AddResource("ga", "compute", "InstanceGroupManager", compute.YAML_instance_group_manager)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("InterconnectAttachment"), compute.YAML_interconnect_attachment)
	d.AddResource("ga", "compute", "InterconnectAttachment", compute.YAML_interconnect_attachment)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Network"), compute.YAML_network)
	d.AddResource("ga", "compute", "Network", compute.YAML_network)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("PacketMirroring"), compute.YAML_packet_mirroring)
	d.AddResource("ga", "compute", "PacketMirroring", compute.YAML_packet_mirroring)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Route"), compute.YAML_route)
	d.AddResource("ga", "compute", "Route", compute.YAML_route)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("ServiceAttachment"), compute.YAML_service_attachment)
	d.AddResource("ga", "compute", "ServiceAttachment", compute.YAML_service_attachment)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("Subnetwork"), compute.YAML_subnetwork)
	d.AddResource("ga", "compute", "Subnetwork", compute.YAML_subnetwork)
	d.AddResource("ga", "compute", dcl.TitleToSnakeCase("VpnTunnel"), compute.YAML_vpn_tunnel)
	d.AddResource("ga", "compute", "VpnTunnel", compute.YAML_vpn_tunnel)
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
	d.AddResource("ga", "filestore", dcl.TitleToSnakeCase("Instance"), filestore.YAML_instance)
	d.AddResource("ga", "filestore", "Instance", filestore.YAML_instance)
	d.AddResource("ga", "gameservices", dcl.TitleToSnakeCase("Realm"), gameservices.YAML_realm)
	d.AddResource("ga", "gameservices", "Realm", gameservices.YAML_realm)
	d.AddResource("ga", "iam", dcl.TitleToSnakeCase("Role"), iam.YAML_role)
	d.AddResource("ga", "iam", "Role", iam.YAML_role)
	d.AddResource("ga", "iam", dcl.TitleToSnakeCase("ServiceAccount"), iam.YAML_service_account)
	d.AddResource("ga", "iam", "ServiceAccount", iam.YAML_service_account)
	d.AddResource("ga", "iam", dcl.TitleToSnakeCase("WorkloadIdentityPool"), iam.YAML_workload_identity_pool)
	d.AddResource("ga", "iam", "WorkloadIdentityPool", iam.YAML_workload_identity_pool)
	d.AddResource("ga", "iam", dcl.TitleToSnakeCase("WorkloadIdentityPoolProvider"), iam.YAML_workload_identity_pool_provider)
	d.AddResource("ga", "iam", "WorkloadIdentityPoolProvider", iam.YAML_workload_identity_pool_provider)
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
	d.AddResource("ga", "monitoring", dcl.TitleToSnakeCase("MetricDescriptor"), monitoring.YAML_metric_descriptor)
	d.AddResource("ga", "monitoring", "MetricDescriptor", monitoring.YAML_metric_descriptor)
	d.AddResource("ga", "monitoring", dcl.TitleToSnakeCase("NotificationChannel"), monitoring.YAML_notification_channel)
	d.AddResource("ga", "monitoring", "NotificationChannel", monitoring.YAML_notification_channel)
	d.AddResource("ga", "monitoring", dcl.TitleToSnakeCase("UptimeCheckConfig"), monitoring.YAML_uptime_check_config)
	d.AddResource("ga", "monitoring", "UptimeCheckConfig", monitoring.YAML_uptime_check_config)
	d.AddResource("ga", "monitoring", dcl.TitleToSnakeCase("Service"), monitoring.YAML_service)
	d.AddResource("ga", "monitoring", "Service", monitoring.YAML_service)
	d.AddResource("ga", "monitoring", dcl.TitleToSnakeCase("ServiceLevelObjective"), monitoring.YAML_service_level_objective)
	d.AddResource("ga", "monitoring", "ServiceLevelObjective", monitoring.YAML_service_level_objective)
	d.AddResource("ga", "networkconnectivity", dcl.TitleToSnakeCase("Hub"), networkconnectivity.YAML_hub)
	d.AddResource("ga", "networkconnectivity", "Hub", networkconnectivity.YAML_hub)
	d.AddResource("ga", "networkconnectivity", dcl.TitleToSnakeCase("Spoke"), networkconnectivity.YAML_spoke)
	d.AddResource("ga", "networkconnectivity", "Spoke", networkconnectivity.YAML_spoke)
	d.AddResource("ga", "orgpolicy", dcl.TitleToSnakeCase("Policy"), orgpolicy.YAML_policy)
	d.AddResource("ga", "orgpolicy", "Policy", orgpolicy.YAML_policy)
	d.AddResource("ga", "osconfig", dcl.TitleToSnakeCase("OSPolicyAssignment"), osconfig.YAML_os_policy_assignment)
	d.AddResource("ga", "osconfig", "OSPolicyAssignment", osconfig.YAML_os_policy_assignment)
	d.AddResource("ga", "pubsub", dcl.TitleToSnakeCase("Topic"), pubsub.YAML_topic)
	d.AddResource("ga", "pubsub", "Topic", pubsub.YAML_topic)
	d.AddResource("ga", "storage", dcl.TitleToSnakeCase("Bucket"), storage.YAML_bucket)
	d.AddResource("ga", "storage", "Bucket", storage.YAML_bucket)
	d.AddResource("ga", "privateca", dcl.TitleToSnakeCase("CertificateTemplate"), privateca.YAML_certificate_template)
	d.AddResource("ga", "privateca", "CertificateTemplate", privateca.YAML_certificate_template)
	d.AddResource("ga", "privateca", dcl.TitleToSnakeCase("CaPool"), privateca.YAML_ca_pool)
	d.AddResource("ga", "privateca", "CaPool", privateca.YAML_ca_pool)
	d.AddResource("ga", "privateca", dcl.TitleToSnakeCase("Certificate"), privateca.YAML_certificate)
	d.AddResource("ga", "privateca", "Certificate", privateca.YAML_certificate)
	d.AddResource("ga", "privateca", dcl.TitleToSnakeCase("CertificateAuthority"), privateca.YAML_certificate_authority)
	d.AddResource("ga", "privateca", "CertificateAuthority", privateca.YAML_certificate_authority)
	d.AddResource("ga", "vpcaccess", dcl.TitleToSnakeCase("Connector"), vpcaccess.YAML_connector)
	d.AddResource("ga", "vpcaccess", "Connector", vpcaccess.YAML_connector)
	d.AddResource("ga", "recaptchaenterprise", dcl.TitleToSnakeCase("Key"), recaptchaenterprise.YAML_key)
	d.AddResource("ga", "recaptchaenterprise", "Key", recaptchaenterprise.YAML_key)
	d.AddResource("beta", "assuredworkloads", dcl.TitleToSnakeCase("Workload"), assuredworkloads_beta.YAML_workload)
	d.AddResource("beta", "assuredworkloads", "Workload", assuredworkloads_beta.YAML_workload)
	d.AddResource("beta", "billingbudgets", dcl.TitleToSnakeCase("Budget"), billingbudgets_beta.YAML_budget)
	d.AddResource("beta", "billingbudgets", "Budget", billingbudgets_beta.YAML_budget)
	d.AddResource("beta", "bigqueryreservation", dcl.TitleToSnakeCase("Assignment"), bigqueryreservation_beta.YAML_assignment)
	d.AddResource("beta", "bigqueryreservation", "Assignment", bigqueryreservation_beta.YAML_assignment)
	d.AddResource("beta", "bigqueryreservation", dcl.TitleToSnakeCase("CapacityCommitment"), bigqueryreservation_beta.YAML_capacity_commitment)
	d.AddResource("beta", "bigqueryreservation", "CapacityCommitment", bigqueryreservation_beta.YAML_capacity_commitment)
	d.AddResource("beta", "bigqueryreservation", dcl.TitleToSnakeCase("Reservation"), bigqueryreservation_beta.YAML_reservation)
	d.AddResource("beta", "bigqueryreservation", "Reservation", bigqueryreservation_beta.YAML_reservation)
	d.AddResource("beta", "binaryauthorization", dcl.TitleToSnakeCase("Attestor"), binaryauthorization_beta.YAML_attestor)
	d.AddResource("beta", "binaryauthorization", "Attestor", binaryauthorization_beta.YAML_attestor)
	d.AddResource("beta", "binaryauthorization", dcl.TitleToSnakeCase("Policy"), binaryauthorization_beta.YAML_policy)
	d.AddResource("beta", "binaryauthorization", "Policy", binaryauthorization_beta.YAML_policy)
	d.AddResource("beta", "cloudbuild", dcl.TitleToSnakeCase("WorkerPool"), cloudbuild_beta.YAML_worker_pool)
	d.AddResource("beta", "cloudbuild", "WorkerPool", cloudbuild_beta.YAML_worker_pool)
	d.AddResource("beta", "cloudfunctions", dcl.TitleToSnakeCase("Function"), cloudfunctions_beta.YAML_function)
	d.AddResource("beta", "cloudfunctions", "Function", cloudfunctions_beta.YAML_function)
	d.AddResource("beta", "cloudidentity", dcl.TitleToSnakeCase("Group"), cloudidentity_beta.YAML_group)
	d.AddResource("beta", "cloudidentity", "Group", cloudidentity_beta.YAML_group)
	d.AddResource("beta", "cloudidentity", dcl.TitleToSnakeCase("Membership"), cloudidentity_beta.YAML_membership)
	d.AddResource("beta", "cloudidentity", "Membership", cloudidentity_beta.YAML_membership)
	d.AddResource("beta", "cloudkms", dcl.TitleToSnakeCase("CryptoKey"), cloudkms_beta.YAML_crypto_key)
	d.AddResource("beta", "cloudkms", "CryptoKey", cloudkms_beta.YAML_crypto_key)
	d.AddResource("beta", "cloudkms", dcl.TitleToSnakeCase("KeyRing"), cloudkms_beta.YAML_key_ring)
	d.AddResource("beta", "cloudkms", "KeyRing", cloudkms_beta.YAML_key_ring)
	d.AddResource("beta", "cloudresourcemanager", dcl.TitleToSnakeCase("Folder"), cloudresourcemanager_beta.YAML_folder)
	d.AddResource("beta", "cloudresourcemanager", "Folder", cloudresourcemanager_beta.YAML_folder)
	d.AddResource("beta", "cloudresourcemanager", dcl.TitleToSnakeCase("Project"), cloudresourcemanager_beta.YAML_project)
	d.AddResource("beta", "cloudresourcemanager", "Project", cloudresourcemanager_beta.YAML_project)
	d.AddResource("beta", "cloudresourcemanager", dcl.TitleToSnakeCase("TagKey"), cloudresourcemanager_beta.YAML_tag_key)
	d.AddResource("beta", "cloudresourcemanager", "TagKey", cloudresourcemanager_beta.YAML_tag_key)
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
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Instance"), compute_beta.YAML_instance)
	d.AddResource("beta", "compute", "Instance", compute_beta.YAML_instance)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("InstanceGroupManager"), compute_beta.YAML_instance_group_manager)
	d.AddResource("beta", "compute", "InstanceGroupManager", compute_beta.YAML_instance_group_manager)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("InterconnectAttachment"), compute_beta.YAML_interconnect_attachment)
	d.AddResource("beta", "compute", "InterconnectAttachment", compute_beta.YAML_interconnect_attachment)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Network"), compute_beta.YAML_network)
	d.AddResource("beta", "compute", "Network", compute_beta.YAML_network)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("PacketMirroring"), compute_beta.YAML_packet_mirroring)
	d.AddResource("beta", "compute", "PacketMirroring", compute_beta.YAML_packet_mirroring)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Route"), compute_beta.YAML_route)
	d.AddResource("beta", "compute", "Route", compute_beta.YAML_route)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("ServiceAttachment"), compute_beta.YAML_service_attachment)
	d.AddResource("beta", "compute", "ServiceAttachment", compute_beta.YAML_service_attachment)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("Subnetwork"), compute_beta.YAML_subnetwork)
	d.AddResource("beta", "compute", "Subnetwork", compute_beta.YAML_subnetwork)
	d.AddResource("beta", "compute", dcl.TitleToSnakeCase("VpnTunnel"), compute_beta.YAML_vpn_tunnel)
	d.AddResource("beta", "compute", "VpnTunnel", compute_beta.YAML_vpn_tunnel)
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
	d.AddResource("beta", "filestore", dcl.TitleToSnakeCase("Instance"), filestore_beta.YAML_instance)
	d.AddResource("beta", "filestore", "Instance", filestore_beta.YAML_instance)
	d.AddResource("beta", "filestore", dcl.TitleToSnakeCase("Backup"), filestore_beta.YAML_backup)
	d.AddResource("beta", "filestore", "Backup", filestore_beta.YAML_backup)
	d.AddResource("beta", "gameservices", dcl.TitleToSnakeCase("Realm"), gameservices_beta.YAML_realm)
	d.AddResource("beta", "gameservices", "Realm", gameservices_beta.YAML_realm)
	d.AddResource("beta", "gkehub", dcl.TitleToSnakeCase("Feature"), gkehub_beta.YAML_feature)
	d.AddResource("beta", "gkehub", "Feature", gkehub_beta.YAML_feature)
	d.AddResource("beta", "gkehub", dcl.TitleToSnakeCase("FeatureMembership"), gkehub_beta.YAML_feature_membership)
	d.AddResource("beta", "gkehub", "FeatureMembership", gkehub_beta.YAML_feature_membership)
	d.AddResource("beta", "gkehub", dcl.TitleToSnakeCase("Membership"), gkehub_beta.YAML_membership)
	d.AddResource("beta", "gkehub", "Membership", gkehub_beta.YAML_membership)
	d.AddResource("beta", "iam", dcl.TitleToSnakeCase("Role"), iam_beta.YAML_role)
	d.AddResource("beta", "iam", "Role", iam_beta.YAML_role)
	d.AddResource("beta", "iam", dcl.TitleToSnakeCase("ServiceAccount"), iam_beta.YAML_service_account)
	d.AddResource("beta", "iam", "ServiceAccount", iam_beta.YAML_service_account)
	d.AddResource("beta", "iam", dcl.TitleToSnakeCase("WorkloadIdentityPool"), iam_beta.YAML_workload_identity_pool)
	d.AddResource("beta", "iam", "WorkloadIdentityPool", iam_beta.YAML_workload_identity_pool)
	d.AddResource("beta", "iam", dcl.TitleToSnakeCase("WorkloadIdentityPoolProvider"), iam_beta.YAML_workload_identity_pool_provider)
	d.AddResource("beta", "iam", "WorkloadIdentityPoolProvider", iam_beta.YAML_workload_identity_pool_provider)
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
	d.AddResource("beta", "monitoring", dcl.TitleToSnakeCase("MetricDescriptor"), monitoring_beta.YAML_metric_descriptor)
	d.AddResource("beta", "monitoring", "MetricDescriptor", monitoring_beta.YAML_metric_descriptor)
	d.AddResource("beta", "monitoring", dcl.TitleToSnakeCase("NotificationChannel"), monitoring_beta.YAML_notification_channel)
	d.AddResource("beta", "monitoring", "NotificationChannel", monitoring_beta.YAML_notification_channel)
	d.AddResource("beta", "monitoring", dcl.TitleToSnakeCase("UptimeCheckConfig"), monitoring_beta.YAML_uptime_check_config)
	d.AddResource("beta", "monitoring", "UptimeCheckConfig", monitoring_beta.YAML_uptime_check_config)
	d.AddResource("beta", "monitoring", dcl.TitleToSnakeCase("Service"), monitoring_beta.YAML_service)
	d.AddResource("beta", "monitoring", "Service", monitoring_beta.YAML_service)
	d.AddResource("beta", "monitoring", dcl.TitleToSnakeCase("ServiceLevelObjective"), monitoring_beta.YAML_service_level_objective)
	d.AddResource("beta", "monitoring", "ServiceLevelObjective", monitoring_beta.YAML_service_level_objective)
	d.AddResource("beta", "monitoring", dcl.TitleToSnakeCase("MetricsScope"), monitoring_beta.YAML_metrics_scope)
	d.AddResource("beta", "monitoring", "MetricsScope", monitoring_beta.YAML_metrics_scope)
	d.AddResource("beta", "monitoring", dcl.TitleToSnakeCase("MonitoredProject"), monitoring_beta.YAML_monitored_project)
	d.AddResource("beta", "monitoring", "MonitoredProject", monitoring_beta.YAML_monitored_project)
	d.AddResource("beta", "networkconnectivity", dcl.TitleToSnakeCase("Hub"), networkconnectivity_beta.YAML_hub)
	d.AddResource("beta", "networkconnectivity", "Hub", networkconnectivity_beta.YAML_hub)
	d.AddResource("beta", "networkconnectivity", dcl.TitleToSnakeCase("Spoke"), networkconnectivity_beta.YAML_spoke)
	d.AddResource("beta", "networkconnectivity", "Spoke", networkconnectivity_beta.YAML_spoke)
	d.AddResource("beta", "networksecurity", dcl.TitleToSnakeCase("AuthorizationPolicy"), networksecurity_beta.YAML_authorization_policy)
	d.AddResource("beta", "networksecurity", "AuthorizationPolicy", networksecurity_beta.YAML_authorization_policy)
	d.AddResource("beta", "networksecurity", dcl.TitleToSnakeCase("ClientTlsPolicy"), networksecurity_beta.YAML_client_tls_policy)
	d.AddResource("beta", "networksecurity", "ClientTlsPolicy", networksecurity_beta.YAML_client_tls_policy)
	d.AddResource("beta", "networksecurity", dcl.TitleToSnakeCase("ServerTlsPolicy"), networksecurity_beta.YAML_server_tls_policy)
	d.AddResource("beta", "networksecurity", "ServerTlsPolicy", networksecurity_beta.YAML_server_tls_policy)
	d.AddResource("beta", "networkservices", dcl.TitleToSnakeCase("EndpointPolicy"), networkservices_beta.YAML_endpoint_policy)
	d.AddResource("beta", "networkservices", "EndpointPolicy", networkservices_beta.YAML_endpoint_policy)
	d.AddResource("beta", "orgpolicy", dcl.TitleToSnakeCase("Policy"), orgpolicy_beta.YAML_policy)
	d.AddResource("beta", "orgpolicy", "Policy", orgpolicy_beta.YAML_policy)
	d.AddResource("beta", "osconfig", dcl.TitleToSnakeCase("OSPolicyAssignment"), osconfig_beta.YAML_os_policy_assignment)
	d.AddResource("beta", "osconfig", "OSPolicyAssignment", osconfig_beta.YAML_os_policy_assignment)
	d.AddResource("beta", "osconfig", dcl.TitleToSnakeCase("GuestPolicy"), osconfig_beta.YAML_guest_policy)
	d.AddResource("beta", "osconfig", "GuestPolicy", osconfig_beta.YAML_guest_policy)
	d.AddResource("beta", "pubsub", dcl.TitleToSnakeCase("Topic"), pubsub_beta.YAML_topic)
	d.AddResource("beta", "pubsub", "Topic", pubsub_beta.YAML_topic)
	d.AddResource("beta", "storage", dcl.TitleToSnakeCase("Bucket"), storage_beta.YAML_bucket)
	d.AddResource("beta", "storage", "Bucket", storage_beta.YAML_bucket)
	d.AddResource("beta", "privateca", dcl.TitleToSnakeCase("CertificateTemplate"), privateca_beta.YAML_certificate_template)
	d.AddResource("beta", "privateca", "CertificateTemplate", privateca_beta.YAML_certificate_template)
	d.AddResource("beta", "privateca", dcl.TitleToSnakeCase("CaPool"), privateca_beta.YAML_ca_pool)
	d.AddResource("beta", "privateca", "CaPool", privateca_beta.YAML_ca_pool)
	d.AddResource("beta", "privateca", dcl.TitleToSnakeCase("Certificate"), privateca_beta.YAML_certificate)
	d.AddResource("beta", "privateca", "Certificate", privateca_beta.YAML_certificate)
	d.AddResource("beta", "privateca", dcl.TitleToSnakeCase("CertificateAuthority"), privateca_beta.YAML_certificate_authority)
	d.AddResource("beta", "privateca", "CertificateAuthority", privateca_beta.YAML_certificate_authority)
	d.AddResource("beta", "vpcaccess", dcl.TitleToSnakeCase("Connector"), vpcaccess_beta.YAML_connector)
	d.AddResource("beta", "vpcaccess", "Connector", vpcaccess_beta.YAML_connector)
	d.AddResource("beta", "recaptchaenterprise", dcl.TitleToSnakeCase("Key"), recaptchaenterprise_beta.YAML_key)
	d.AddResource("beta", "recaptchaenterprise", "Key", recaptchaenterprise_beta.YAML_key)
	d.AddResource("alpha", "assuredworkloads", dcl.TitleToSnakeCase("Workload"), assuredworkloads_alpha.YAML_workload)
	d.AddResource("alpha", "assuredworkloads", "Workload", assuredworkloads_alpha.YAML_workload)
	d.AddResource("alpha", "billingbudgets", dcl.TitleToSnakeCase("Budget"), billingbudgets_alpha.YAML_budget)
	d.AddResource("alpha", "billingbudgets", "Budget", billingbudgets_alpha.YAML_budget)
	d.AddResource("alpha", "bigqueryreservation", dcl.TitleToSnakeCase("Assignment"), bigqueryreservation_alpha.YAML_assignment)
	d.AddResource("alpha", "bigqueryreservation", "Assignment", bigqueryreservation_alpha.YAML_assignment)
	d.AddResource("alpha", "bigqueryreservation", dcl.TitleToSnakeCase("CapacityCommitment"), bigqueryreservation_alpha.YAML_capacity_commitment)
	d.AddResource("alpha", "bigqueryreservation", "CapacityCommitment", bigqueryreservation_alpha.YAML_capacity_commitment)
	d.AddResource("alpha", "bigqueryreservation", dcl.TitleToSnakeCase("Reservation"), bigqueryreservation_alpha.YAML_reservation)
	d.AddResource("alpha", "bigqueryreservation", "Reservation", bigqueryreservation_alpha.YAML_reservation)
	d.AddResource("alpha", "binaryauthorization", dcl.TitleToSnakeCase("Attestor"), binaryauthorization_alpha.YAML_attestor)
	d.AddResource("alpha", "binaryauthorization", "Attestor", binaryauthorization_alpha.YAML_attestor)
	d.AddResource("alpha", "binaryauthorization", dcl.TitleToSnakeCase("Policy"), binaryauthorization_alpha.YAML_policy)
	d.AddResource("alpha", "binaryauthorization", "Policy", binaryauthorization_alpha.YAML_policy)
	d.AddResource("alpha", "cloudbuild", dcl.TitleToSnakeCase("WorkerPool"), cloudbuild_alpha.YAML_worker_pool)
	d.AddResource("alpha", "cloudbuild", "WorkerPool", cloudbuild_alpha.YAML_worker_pool)
	d.AddResource("alpha", "cloudfunctions", dcl.TitleToSnakeCase("Function"), cloudfunctions_alpha.YAML_function)
	d.AddResource("alpha", "cloudfunctions", "Function", cloudfunctions_alpha.YAML_function)
	d.AddResource("alpha", "cloudidentity", dcl.TitleToSnakeCase("Group"), cloudidentity_alpha.YAML_group)
	d.AddResource("alpha", "cloudidentity", "Group", cloudidentity_alpha.YAML_group)
	d.AddResource("alpha", "cloudidentity", dcl.TitleToSnakeCase("Membership"), cloudidentity_alpha.YAML_membership)
	d.AddResource("alpha", "cloudidentity", "Membership", cloudidentity_alpha.YAML_membership)
	d.AddResource("alpha", "cloudkms", dcl.TitleToSnakeCase("CryptoKey"), cloudkms_alpha.YAML_crypto_key)
	d.AddResource("alpha", "cloudkms", "CryptoKey", cloudkms_alpha.YAML_crypto_key)
	d.AddResource("alpha", "cloudkms", dcl.TitleToSnakeCase("KeyRing"), cloudkms_alpha.YAML_key_ring)
	d.AddResource("alpha", "cloudkms", "KeyRing", cloudkms_alpha.YAML_key_ring)
	d.AddResource("alpha", "cloudresourcemanager", dcl.TitleToSnakeCase("Folder"), cloudresourcemanager_alpha.YAML_folder)
	d.AddResource("alpha", "cloudresourcemanager", "Folder", cloudresourcemanager_alpha.YAML_folder)
	d.AddResource("alpha", "cloudresourcemanager", dcl.TitleToSnakeCase("Project"), cloudresourcemanager_alpha.YAML_project)
	d.AddResource("alpha", "cloudresourcemanager", "Project", cloudresourcemanager_alpha.YAML_project)
	d.AddResource("alpha", "cloudresourcemanager", dcl.TitleToSnakeCase("TagKey"), cloudresourcemanager_alpha.YAML_tag_key)
	d.AddResource("alpha", "cloudresourcemanager", "TagKey", cloudresourcemanager_alpha.YAML_tag_key)
	d.AddResource("alpha", "cloudscheduler", dcl.TitleToSnakeCase("Job"), cloudscheduler_alpha.YAML_job)
	d.AddResource("alpha", "cloudscheduler", "Job", cloudscheduler_alpha.YAML_job)
	d.AddResource("alpha", "compute", dcl.TitleToSnakeCase("FirewallPolicy"), compute_alpha.YAML_firewall_policy)
	d.AddResource("alpha", "compute", "FirewallPolicy", compute_alpha.YAML_firewall_policy)
	d.AddResource("alpha", "compute", dcl.TitleToSnakeCase("FirewallPolicyAssociation"), compute_alpha.YAML_firewall_policy_association)
	d.AddResource("alpha", "compute", "FirewallPolicyAssociation", compute_alpha.YAML_firewall_policy_association)
	d.AddResource("alpha", "compute", dcl.TitleToSnakeCase("FirewallPolicyRule"), compute_alpha.YAML_firewall_policy_rule)
	d.AddResource("alpha", "compute", "FirewallPolicyRule", compute_alpha.YAML_firewall_policy_rule)
	d.AddResource("alpha", "compute", dcl.TitleToSnakeCase("ForwardingRule"), compute_alpha.YAML_forwarding_rule)
	d.AddResource("alpha", "compute", "ForwardingRule", compute_alpha.YAML_forwarding_rule)
	d.AddResource("alpha", "compute", dcl.TitleToSnakeCase("Instance"), compute_alpha.YAML_instance)
	d.AddResource("alpha", "compute", "Instance", compute_alpha.YAML_instance)
	d.AddResource("alpha", "compute", dcl.TitleToSnakeCase("InstanceGroupManager"), compute_alpha.YAML_instance_group_manager)
	d.AddResource("alpha", "compute", "InstanceGroupManager", compute_alpha.YAML_instance_group_manager)
	d.AddResource("alpha", "compute", dcl.TitleToSnakeCase("InterconnectAttachment"), compute_alpha.YAML_interconnect_attachment)
	d.AddResource("alpha", "compute", "InterconnectAttachment", compute_alpha.YAML_interconnect_attachment)
	d.AddResource("alpha", "compute", dcl.TitleToSnakeCase("Network"), compute_alpha.YAML_network)
	d.AddResource("alpha", "compute", "Network", compute_alpha.YAML_network)
	d.AddResource("alpha", "compute", dcl.TitleToSnakeCase("PacketMirroring"), compute_alpha.YAML_packet_mirroring)
	d.AddResource("alpha", "compute", "PacketMirroring", compute_alpha.YAML_packet_mirroring)
	d.AddResource("alpha", "compute", dcl.TitleToSnakeCase("Route"), compute_alpha.YAML_route)
	d.AddResource("alpha", "compute", "Route", compute_alpha.YAML_route)
	d.AddResource("alpha", "compute", dcl.TitleToSnakeCase("ServiceAttachment"), compute_alpha.YAML_service_attachment)
	d.AddResource("alpha", "compute", "ServiceAttachment", compute_alpha.YAML_service_attachment)
	d.AddResource("alpha", "compute", dcl.TitleToSnakeCase("Subnetwork"), compute_alpha.YAML_subnetwork)
	d.AddResource("alpha", "compute", "Subnetwork", compute_alpha.YAML_subnetwork)
	d.AddResource("alpha", "compute", dcl.TitleToSnakeCase("VpnTunnel"), compute_alpha.YAML_vpn_tunnel)
	d.AddResource("alpha", "compute", "VpnTunnel", compute_alpha.YAML_vpn_tunnel)
	d.AddResource("alpha", "configcontroller", dcl.TitleToSnakeCase("Instance"), configcontroller_alpha.YAML_instance)
	d.AddResource("alpha", "configcontroller", "Instance", configcontroller_alpha.YAML_instance)
	d.AddResource("alpha", "containeranalysis", dcl.TitleToSnakeCase("Note"), containeranalysis_alpha.YAML_note)
	d.AddResource("alpha", "containeranalysis", "Note", containeranalysis_alpha.YAML_note)
	d.AddResource("alpha", "containeraws", dcl.TitleToSnakeCase("Cluster"), containeraws_alpha.YAML_cluster)
	d.AddResource("alpha", "containeraws", "Cluster", containeraws_alpha.YAML_cluster)
	d.AddResource("alpha", "containeraws", dcl.TitleToSnakeCase("NodePool"), containeraws_alpha.YAML_node_pool)
	d.AddResource("alpha", "containeraws", "NodePool", containeraws_alpha.YAML_node_pool)
	d.AddResource("alpha", "containerazure", dcl.TitleToSnakeCase("AzureClient"), containerazure_alpha.YAML_azure_client)
	d.AddResource("alpha", "containerazure", "AzureClient", containerazure_alpha.YAML_azure_client)
	d.AddResource("alpha", "containerazure", dcl.TitleToSnakeCase("Cluster"), containerazure_alpha.YAML_cluster)
	d.AddResource("alpha", "containerazure", "Cluster", containerazure_alpha.YAML_cluster)
	d.AddResource("alpha", "containerazure", dcl.TitleToSnakeCase("NodePool"), containerazure_alpha.YAML_node_pool)
	d.AddResource("alpha", "containerazure", "NodePool", containerazure_alpha.YAML_node_pool)
	d.AddResource("alpha", "datafusion", dcl.TitleToSnakeCase("Instance"), datafusion_alpha.YAML_instance)
	d.AddResource("alpha", "datafusion", "Instance", datafusion_alpha.YAML_instance)
	d.AddResource("alpha", "dataproc", dcl.TitleToSnakeCase("AutoscalingPolicy"), dataproc_alpha.YAML_autoscaling_policy)
	d.AddResource("alpha", "dataproc", "AutoscalingPolicy", dataproc_alpha.YAML_autoscaling_policy)
	d.AddResource("alpha", "dataproc", dcl.TitleToSnakeCase("Cluster"), dataproc_alpha.YAML_cluster)
	d.AddResource("alpha", "dataproc", "Cluster", dataproc_alpha.YAML_cluster)
	d.AddResource("alpha", "dataproc", dcl.TitleToSnakeCase("WorkflowTemplate"), dataproc_alpha.YAML_workflow_template)
	d.AddResource("alpha", "dataproc", "WorkflowTemplate", dataproc_alpha.YAML_workflow_template)
	d.AddResource("alpha", "eventarc", dcl.TitleToSnakeCase("Trigger"), eventarc_alpha.YAML_trigger)
	d.AddResource("alpha", "eventarc", "Trigger", eventarc_alpha.YAML_trigger)
	d.AddResource("alpha", "filestore", dcl.TitleToSnakeCase("Instance"), filestore_alpha.YAML_instance)
	d.AddResource("alpha", "filestore", "Instance", filestore_alpha.YAML_instance)
	d.AddResource("alpha", "filestore", dcl.TitleToSnakeCase("Backup"), filestore_alpha.YAML_backup)
	d.AddResource("alpha", "filestore", "Backup", filestore_alpha.YAML_backup)
	d.AddResource("alpha", "gameservices", dcl.TitleToSnakeCase("Realm"), gameservices_alpha.YAML_realm)
	d.AddResource("alpha", "gameservices", "Realm", gameservices_alpha.YAML_realm)
	d.AddResource("alpha", "gkehub", dcl.TitleToSnakeCase("Feature"), gkehub_alpha.YAML_feature)
	d.AddResource("alpha", "gkehub", "Feature", gkehub_alpha.YAML_feature)
	d.AddResource("alpha", "gkehub", dcl.TitleToSnakeCase("FeatureMembership"), gkehub_alpha.YAML_feature_membership)
	d.AddResource("alpha", "gkehub", "FeatureMembership", gkehub_alpha.YAML_feature_membership)
	d.AddResource("alpha", "gkehub", dcl.TitleToSnakeCase("Membership"), gkehub_alpha.YAML_membership)
	d.AddResource("alpha", "gkehub", "Membership", gkehub_alpha.YAML_membership)
	d.AddResource("alpha", "iam", dcl.TitleToSnakeCase("Role"), iam_alpha.YAML_role)
	d.AddResource("alpha", "iam", "Role", iam_alpha.YAML_role)
	d.AddResource("alpha", "iam", dcl.TitleToSnakeCase("ServiceAccount"), iam_alpha.YAML_service_account)
	d.AddResource("alpha", "iam", "ServiceAccount", iam_alpha.YAML_service_account)
	d.AddResource("alpha", "iam", dcl.TitleToSnakeCase("WorkloadIdentityPool"), iam_alpha.YAML_workload_identity_pool)
	d.AddResource("alpha", "iam", "WorkloadIdentityPool", iam_alpha.YAML_workload_identity_pool)
	d.AddResource("alpha", "iam", dcl.TitleToSnakeCase("WorkloadIdentityPoolProvider"), iam_alpha.YAML_workload_identity_pool_provider)
	d.AddResource("alpha", "iam", "WorkloadIdentityPoolProvider", iam_alpha.YAML_workload_identity_pool_provider)
	d.AddResource("alpha", "iap", dcl.TitleToSnakeCase("Brand"), iap_alpha.YAML_brand)
	d.AddResource("alpha", "iap", "Brand", iap_alpha.YAML_brand)
	d.AddResource("alpha", "iap", dcl.TitleToSnakeCase("IdentityAwareProxyClient"), iap_alpha.YAML_identity_aware_proxy_client)
	d.AddResource("alpha", "iap", "IdentityAwareProxyClient", iap_alpha.YAML_identity_aware_proxy_client)
	d.AddResource("alpha", "identitytoolkit", dcl.TitleToSnakeCase("OAuthIdpConfig"), identitytoolkit_alpha.YAML_oauth_idp_config)
	d.AddResource("alpha", "identitytoolkit", "OAuthIdpConfig", identitytoolkit_alpha.YAML_oauth_idp_config)
	d.AddResource("alpha", "identitytoolkit", dcl.TitleToSnakeCase("Tenant"), identitytoolkit_alpha.YAML_tenant)
	d.AddResource("alpha", "identitytoolkit", "Tenant", identitytoolkit_alpha.YAML_tenant)
	d.AddResource("alpha", "identitytoolkit", dcl.TitleToSnakeCase("TenantOAuthIdpConfig"), identitytoolkit_alpha.YAML_tenant_oauth_idp_config)
	d.AddResource("alpha", "identitytoolkit", "TenantOAuthIdpConfig", identitytoolkit_alpha.YAML_tenant_oauth_idp_config)
	d.AddResource("alpha", "logging", dcl.TitleToSnakeCase("LogExclusion"), logging_alpha.YAML_log_exclusion)
	d.AddResource("alpha", "logging", "LogExclusion", logging_alpha.YAML_log_exclusion)
	d.AddResource("alpha", "monitoring", dcl.TitleToSnakeCase("Dashboard"), monitoring_alpha.YAML_dashboard)
	d.AddResource("alpha", "monitoring", "Dashboard", monitoring_alpha.YAML_dashboard)
	d.AddResource("alpha", "monitoring", dcl.TitleToSnakeCase("Group"), monitoring_alpha.YAML_group)
	d.AddResource("alpha", "monitoring", "Group", monitoring_alpha.YAML_group)
	d.AddResource("alpha", "monitoring", dcl.TitleToSnakeCase("MetricDescriptor"), monitoring_alpha.YAML_metric_descriptor)
	d.AddResource("alpha", "monitoring", "MetricDescriptor", monitoring_alpha.YAML_metric_descriptor)
	d.AddResource("alpha", "monitoring", dcl.TitleToSnakeCase("NotificationChannel"), monitoring_alpha.YAML_notification_channel)
	d.AddResource("alpha", "monitoring", "NotificationChannel", monitoring_alpha.YAML_notification_channel)
	d.AddResource("alpha", "monitoring", dcl.TitleToSnakeCase("UptimeCheckConfig"), monitoring_alpha.YAML_uptime_check_config)
	d.AddResource("alpha", "monitoring", "UptimeCheckConfig", monitoring_alpha.YAML_uptime_check_config)
	d.AddResource("alpha", "monitoring", dcl.TitleToSnakeCase("Service"), monitoring_alpha.YAML_service)
	d.AddResource("alpha", "monitoring", "Service", monitoring_alpha.YAML_service)
	d.AddResource("alpha", "monitoring", dcl.TitleToSnakeCase("ServiceLevelObjective"), monitoring_alpha.YAML_service_level_objective)
	d.AddResource("alpha", "monitoring", "ServiceLevelObjective", monitoring_alpha.YAML_service_level_objective)
	d.AddResource("alpha", "monitoring", dcl.TitleToSnakeCase("MetricsScope"), monitoring_alpha.YAML_metrics_scope)
	d.AddResource("alpha", "monitoring", "MetricsScope", monitoring_alpha.YAML_metrics_scope)
	d.AddResource("alpha", "monitoring", dcl.TitleToSnakeCase("MonitoredProject"), monitoring_alpha.YAML_monitored_project)
	d.AddResource("alpha", "monitoring", "MonitoredProject", monitoring_alpha.YAML_monitored_project)
	d.AddResource("alpha", "networkconnectivity", dcl.TitleToSnakeCase("Hub"), networkconnectivity_alpha.YAML_hub)
	d.AddResource("alpha", "networkconnectivity", "Hub", networkconnectivity_alpha.YAML_hub)
	d.AddResource("alpha", "networkconnectivity", dcl.TitleToSnakeCase("Spoke"), networkconnectivity_alpha.YAML_spoke)
	d.AddResource("alpha", "networkconnectivity", "Spoke", networkconnectivity_alpha.YAML_spoke)
	d.AddResource("alpha", "networksecurity", dcl.TitleToSnakeCase("AuthorizationPolicy"), networksecurity_alpha.YAML_authorization_policy)
	d.AddResource("alpha", "networksecurity", "AuthorizationPolicy", networksecurity_alpha.YAML_authorization_policy)
	d.AddResource("alpha", "networksecurity", dcl.TitleToSnakeCase("ClientTlsPolicy"), networksecurity_alpha.YAML_client_tls_policy)
	d.AddResource("alpha", "networksecurity", "ClientTlsPolicy", networksecurity_alpha.YAML_client_tls_policy)
	d.AddResource("alpha", "networksecurity", dcl.TitleToSnakeCase("ServerTlsPolicy"), networksecurity_alpha.YAML_server_tls_policy)
	d.AddResource("alpha", "networksecurity", "ServerTlsPolicy", networksecurity_alpha.YAML_server_tls_policy)
	d.AddResource("alpha", "networkservices", dcl.TitleToSnakeCase("EndpointPolicy"), networkservices_alpha.YAML_endpoint_policy)
	d.AddResource("alpha", "networkservices", "EndpointPolicy", networkservices_alpha.YAML_endpoint_policy)
	d.AddResource("alpha", "orgpolicy", dcl.TitleToSnakeCase("Policy"), orgpolicy_alpha.YAML_policy)
	d.AddResource("alpha", "orgpolicy", "Policy", orgpolicy_alpha.YAML_policy)
	d.AddResource("alpha", "osconfig", dcl.TitleToSnakeCase("OSPolicyAssignment"), osconfig_alpha.YAML_os_policy_assignment)
	d.AddResource("alpha", "osconfig", "OSPolicyAssignment", osconfig_alpha.YAML_os_policy_assignment)
	d.AddResource("alpha", "osconfig", dcl.TitleToSnakeCase("GuestPolicy"), osconfig_alpha.YAML_guest_policy)
	d.AddResource("alpha", "osconfig", "GuestPolicy", osconfig_alpha.YAML_guest_policy)
	d.AddResource("alpha", "pubsub", dcl.TitleToSnakeCase("Topic"), pubsub_alpha.YAML_topic)
	d.AddResource("alpha", "pubsub", "Topic", pubsub_alpha.YAML_topic)
	d.AddResource("alpha", "run", dcl.TitleToSnakeCase("Service"), run_alpha.YAML_service)
	d.AddResource("alpha", "run", "Service", run_alpha.YAML_service)
	d.AddResource("alpha", "storage", dcl.TitleToSnakeCase("Bucket"), storage_alpha.YAML_bucket)
	d.AddResource("alpha", "storage", "Bucket", storage_alpha.YAML_bucket)
	d.AddResource("alpha", "privateca", dcl.TitleToSnakeCase("CertificateTemplate"), privateca_alpha.YAML_certificate_template)
	d.AddResource("alpha", "privateca", "CertificateTemplate", privateca_alpha.YAML_certificate_template)
	d.AddResource("alpha", "privateca", dcl.TitleToSnakeCase("CaPool"), privateca_alpha.YAML_ca_pool)
	d.AddResource("alpha", "privateca", "CaPool", privateca_alpha.YAML_ca_pool)
	d.AddResource("alpha", "privateca", dcl.TitleToSnakeCase("Certificate"), privateca_alpha.YAML_certificate)
	d.AddResource("alpha", "privateca", "Certificate", privateca_alpha.YAML_certificate)
	d.AddResource("alpha", "privateca", dcl.TitleToSnakeCase("CertificateAuthority"), privateca_alpha.YAML_certificate_authority)
	d.AddResource("alpha", "privateca", "CertificateAuthority", privateca_alpha.YAML_certificate_authority)
	d.AddResource("alpha", "vmware", dcl.TitleToSnakeCase("PrivateCloud"), vmware_alpha.YAML_private_cloud)
	d.AddResource("alpha", "vmware", "PrivateCloud", vmware_alpha.YAML_private_cloud)
	d.AddResource("alpha", "vmware", dcl.TitleToSnakeCase("Cluster"), vmware_alpha.YAML_cluster)
	d.AddResource("alpha", "vmware", "Cluster", vmware_alpha.YAML_cluster)
	d.AddResource("alpha", "vpcaccess", dcl.TitleToSnakeCase("Connector"), vpcaccess_alpha.YAML_connector)
	d.AddResource("alpha", "vpcaccess", "Connector", vpcaccess_alpha.YAML_connector)
	d.AddResource("alpha", "recaptchaenterprise", dcl.TitleToSnakeCase("Key"), recaptchaenterprise_alpha.YAML_key)
	d.AddResource("alpha", "recaptchaenterprise", "Key", recaptchaenterprise_alpha.YAML_key)
	return d
}
