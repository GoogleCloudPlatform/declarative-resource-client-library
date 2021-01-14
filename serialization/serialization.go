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
package serialization

import (
	"bytes"
	"encoding/json"
	"fmt"
	fmtcmd "github.com/hashicorp/hcl/hcl/fmtcmd"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/accesscontextmanager"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dns"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/redis"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage"
	"strings"
)

// TFJSONToDCL takes a tfjson plan and returns a list of DCL resources which are in that plan.
// These resources may be more verbose than the config in the HCL - this is a representation
// of what Terraform actually intends to do, rather than a simple reproduction of the
// given HCL.  This provides a better match to the Terraform user experience than the alternative.
func TFJSONToDCL(p tfjson.Plan) ([]dcl.Resource, error) {
	var resources []dcl.Resource
	for _, r := range p.PlannedValues.RootModule.Resources {
		if r.Type == "google_accesscontextmanager_access_level" {
			m := map[string]interface{}{
				"name":        r.AttributeValues["name"],
				"title":       r.AttributeValues["title"],
				"basic":       convertAccesscontextmanagerAccessLevelBasicList(r.AttributeValues["basic"]),
				"description": r.AttributeValues["description"],
				"policy":      r.AttributeValues["policy"],
				"createTime":  r.AttributeValues["create_time"],
				"updateTime":  r.AttributeValues["update_time"],
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &accesscontextmanager.AccessLevel{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
		if r.Type == "google_accesscontextmanager_access_policy" {
			m := map[string]interface{}{
				"parent":     r.AttributeValues["parent"],
				"title":      r.AttributeValues["title"],
				"createTime": r.AttributeValues["create_time"],
				"name":       r.AttributeValues["name"],
				"updateTime": r.AttributeValues["update_time"],
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &accesscontextmanager.AccessPolicy{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
		if r.Type == "google_accesscontextmanager_service_perimeter" {
			m := map[string]interface{}{
				"name":                  r.AttributeValues["name"],
				"policy":                r.AttributeValues["policy"],
				"title":                 r.AttributeValues["title"],
				"description":           r.AttributeValues["description"],
				"perimeterType":         r.AttributeValues["perimeter_type"],
				"spec":                  convertAccesscontextmanagerServicePerimeterSpecList(r.AttributeValues["spec"]),
				"status":                convertAccesscontextmanagerServicePerimeterStatusList(r.AttributeValues["status"]),
				"useExplicitDryRunSpec": r.AttributeValues["use_explicit_dry_run_spec"],
				"createTime":            r.AttributeValues["create_time"],
				"updateTime":            r.AttributeValues["update_time"],
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &accesscontextmanager.ServicePerimeter{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
		if r.Type == "google_compute_backend_bucket" {
			m := map[string]interface{}{
				"bucketName":  r.AttributeValues["bucket_name"],
				"cdnPolicy":   convertComputeBackendBucketCdnPolicyList(r.AttributeValues["cdn_policy"]),
				"description": r.AttributeValues["description"],
				"enableCdn":   r.AttributeValues["enable_cdn"],
				"name":        r.AttributeValues["name"],
				"project":     r.AttributeValues["project"],
				"selfLink":    r.AttributeValues["self_link"],
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &compute.BackendBucket{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
		if r.Type == "google_compute_firewall" {
			m := map[string]interface{}{
				"name":                  r.AttributeValues["name"],
				"network":               r.AttributeValues["network"],
				"allowed":               convertComputeFirewallAllowedList(r.AttributeValues["allow"]),
				"denied":                convertComputeFirewallDeniedList(r.AttributeValues["deny"]),
				"description":           r.AttributeValues["description"],
				"destinationRanges":     r.AttributeValues["destination_ranges"],
				"direction":             r.AttributeValues["direction"],
				"disabled":              r.AttributeValues["disabled"],
				"logConfig":             convertComputeFirewallLogConfigList(r.AttributeValues["log_config"]),
				"priority":              r.AttributeValues["priority"],
				"project":               r.AttributeValues["project"],
				"sourceRanges":          r.AttributeValues["source_ranges"],
				"sourceServiceAccounts": r.AttributeValues["source_service_accounts"],
				"sourceTags":            r.AttributeValues["source_tags"],
				"targetServiceAccounts": r.AttributeValues["target_service_accounts"],
				"targetTags":            r.AttributeValues["target_tags"],
				"creationTimestamp":     r.AttributeValues["creation_timestamp"],
				"selfLink":              r.AttributeValues["self_link"],
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &compute.Firewall{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
		if r.Type == "google_compute_image" {
			m := map[string]interface{}{
				"deprecated":                   convertComputeImageDeprecatedList(r.AttributeValues["deprecated"]),
				"description":                  r.AttributeValues["description"],
				"diskSizeGb":                   r.AttributeValues["disk_size_gb"],
				"family":                       r.AttributeValues["family"],
				"guestOsFeature":               convertComputeImageGuestOsFeatureList(r.AttributeValues["guest_os_features"]),
				"imageEncryptionKey":           convertComputeImageImageEncryptionKeyList(r.AttributeValues["image_encryption_key"]),
				"labels":                       r.AttributeValues["labels"],
				"license":                      r.AttributeValues["licenses"],
				"name":                         r.AttributeValues["name"],
				"project":                      r.AttributeValues["project"],
				"rawDisk":                      convertComputeImageRawDiskList(r.AttributeValues["raw_disk"]),
				"shieldedInstanceInitialState": convertComputeImageShieldedInstanceInitialStateList(r.AttributeValues["shielded_instance_initial_state"]),
				"sourceDisk":                   r.AttributeValues["source_disk"],
				"sourceDiskEncryptionKey":      convertComputeImageSourceDiskEncryptionKeyList(r.AttributeValues["source_disk_encryption_key"]),
				"sourceImage":                  r.AttributeValues["source_image"],
				"sourceImageEncryptionKey":     convertComputeImageSourceImageEncryptionKeyList(r.AttributeValues["source_image_encryption_key"]),
				"sourceImageId":                r.AttributeValues["source_image_id"],
				"sourceSnapshot":               r.AttributeValues["source_snapshot"],
				"sourceSnapshotEncryptionKey":  convertComputeImageSourceSnapshotEncryptionKeyList(r.AttributeValues["source_snapshot_encryption_key"]),
				"sourceSnapshotId":             r.AttributeValues["source_snapshot_id"],
				"sourceType":                   r.AttributeValues["source_type"],
				"storageLocation":              r.AttributeValues["storage_location"],
				"archiveSizeBytes":             r.AttributeValues["archive_size_bytes"],
				"selfLink":                     r.AttributeValues["self_link"],
				"sourceDiskId":                 r.AttributeValues["source_disk_id"],
				"status":                       r.AttributeValues["status"],
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &compute.Image{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
		if r.Type == "google_compute_network" {
			m := map[string]interface{}{
				"name":                  r.AttributeValues["name"],
				"autoCreateSubnetworks": r.AttributeValues["auto_create_subnetworks"],
				"description":           r.AttributeValues["description"],
				"project":               r.AttributeValues["project"],
				"routingConfig":         convertComputeNetworkRoutingConfigList(r.AttributeValues["routing_config"]),
				"gatewayIPv4":           r.AttributeValues["gateway_ipv4"],
				"iPv4Range":             r.AttributeValues["ipv4_range"],
				"selfLink":              r.AttributeValues["self_link"],
				"":                      r.AttributeValues["delete_default_routes_on_create"],
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &compute.Network{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
		if r.Type == "google_compute_route" {
			m := map[string]interface{}{
				"destRange":        r.AttributeValues["dest_range"],
				"name":             r.AttributeValues["name"],
				"network":          r.AttributeValues["network"],
				"description":      r.AttributeValues["description"],
				"nextHopGateway":   r.AttributeValues["next_hop_gateway"],
				"nextHopIlb":       r.AttributeValues["next_hop_ilb"],
				"nextHopInstance":  r.AttributeValues["next_hop_instance"],
				"nextHopIP":        r.AttributeValues["next_hop_ip"],
				"nextHopNetwork":   r.AttributeValues["next_hop_network"],
				"nextHopVpnTunnel": r.AttributeValues["next_hop_vpn_tunnel"],
				"priority":         r.AttributeValues["priority"],
				"project":          r.AttributeValues["project"],
				"tag":              r.AttributeValues["tags"],
				"nextHopPeering":   r.AttributeValues["next_hop_peering"],
				"selfLink":         r.AttributeValues["self_link"],
				"warning":          convertComputeRouteWarningList(r.AttributeValues["warning"]),
				"":                 r.AttributeValues["next_hop_instance_zone"],
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &compute.Route{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
		if r.Type == "google_compute_ssl_policy" {
			m := map[string]interface{}{
				"name":           r.AttributeValues["name"],
				"customFeature":  r.AttributeValues["custom_features"],
				"description":    r.AttributeValues["description"],
				"minTlsVersion":  r.AttributeValues["min_tls_version"],
				"profile":        r.AttributeValues["profile"],
				"project":        r.AttributeValues["project"],
				"enabledFeature": r.AttributeValues["enabled_features"],
				"selfLink":       r.AttributeValues["self_link"],
				"warning":        convertComputeSslPolicyWarningList(r.AttributeValues["warning"]),
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &compute.SslPolicy{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
		if r.Type == "google_compute_subnetwork" {
			m := map[string]interface{}{
				"iPCidrRange":           r.AttributeValues["ip_cidr_range"],
				"name":                  r.AttributeValues["name"],
				"network":               r.AttributeValues["network"],
				"description":           r.AttributeValues["description"],
				"enableFlowLogs":        r.AttributeValues["enable_flow_logs"],
				"logConfig":             convertComputeSubnetworkLogConfigList(r.AttributeValues["log_config"]),
				"privateIPGoogleAccess": r.AttributeValues["private_ip_google_access"],
				"project":               r.AttributeValues["project"],
				"purpose":               r.AttributeValues["purpose"],
				"region":                r.AttributeValues["region"],
				"role":                  r.AttributeValues["role"],
				"secondaryIPRange":      convertComputeSubnetworkSecondaryIPRangeList(r.AttributeValues["secondary_ip_range"]),
				"creationTimestamp":     r.AttributeValues["creation_timestamp"],
				"fingerprint":           r.AttributeValues["fingerprint"],
				"gatewayAddress":        r.AttributeValues["gateway_address"],
				"selfLink":              r.AttributeValues["self_link"],
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &compute.Subnetwork{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
		if r.Type == "google_dns_managed_zone" {
			m := map[string]interface{}{
				"dnsName":                 r.AttributeValues["dns_name"],
				"name":                    r.AttributeValues["name"],
				"description":             r.AttributeValues["description"],
				"dnssecConfig":            convertDnsManagedZoneDnssecConfigList(r.AttributeValues["dnssec_config"]),
				"forwardingConfig":        convertDnsManagedZoneForwardingConfigList(r.AttributeValues["forwarding_config"]),
				"labels":                  r.AttributeValues["labels"],
				"peeringConfig":           convertDnsManagedZonePeeringConfigList(r.AttributeValues["peering_config"]),
				"privateVisibilityConfig": convertDnsManagedZonePrivateVisibilityConfigList(r.AttributeValues["private_visibility_config"]),
				"project":                 r.AttributeValues["project"],
				"reverseLookup":           r.AttributeValues["reverse_lookup"],
				"visibility":              r.AttributeValues["visibility"],
				"nameServers":             r.AttributeValues["name_servers"],
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &dns.ManagedZone{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
		if r.Type == "google_pubsub_subscription" {
			m := map[string]interface{}{
				"name":                     r.AttributeValues["name"],
				"topic":                    r.AttributeValues["topic"],
				"ackDeadlineSeconds":       r.AttributeValues["ack_deadline_seconds"],
				"deadLetterPolicy":         convertPubsubSubscriptionDeadLetterPolicyList(r.AttributeValues["dead_letter_policy"]),
				"expirationPolicy":         convertPubsubSubscriptionExpirationPolicyList(r.AttributeValues["expiration_policy"]),
				"labels":                   r.AttributeValues["labels"],
				"messageRetentionDuration": r.AttributeValues["message_retention_duration"],
				"project":                  r.AttributeValues["project"],
				"pushConfig":               convertPubsubSubscriptionPushConfigList(r.AttributeValues["push_config"]),
				"retainAckedMessages":      r.AttributeValues["retain_acked_messages"],
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &pubsub.Subscription{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
		if r.Type == "google_pubsub_topic" {
			m := map[string]interface{}{
				"name":                 r.AttributeValues["name"],
				"kmsKeyName":           r.AttributeValues["kms_key_name"],
				"labels":               r.AttributeValues["labels"],
				"messageStoragePolicy": convertPubsubTopicMessageStoragePolicyList(r.AttributeValues["message_storage_policy"]),
				"project":              r.AttributeValues["project"],
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &pubsub.Topic{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
		if r.Type == "google_redis_instance" {
			m := map[string]interface{}{
				"memorySizeGb":          r.AttributeValues["memory_size_gb"],
				"name":                  r.AttributeValues["name"],
				"alternativeLocationId": r.AttributeValues["alternative_location_id"],
				"authorizedNetwork":     r.AttributeValues["authorized_network"],
				"displayName":           r.AttributeValues["display_name"],
				"labels":                r.AttributeValues["labels"],
				"locationId":            r.AttributeValues["location_id"],
				"project":               r.AttributeValues["project"],
				"redisConfigs":          r.AttributeValues["redis_configs"],
				"redisVersion":          r.AttributeValues["redis_version"],
				"region":                r.AttributeValues["region"],
				"reservedIPRange":       r.AttributeValues["reserved_ip_range"],
				"tier":                  r.AttributeValues["tier"],
				"currentLocationId":     r.AttributeValues["current_location_id"],
				"host":                  r.AttributeValues["host"],
				"port":                  r.AttributeValues["port"],
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &redis.Instance{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
		if r.Type == "google_storage_bucket" {
			m := map[string]interface{}{
				"location":     r.AttributeValues["location"],
				"name":         r.AttributeValues["name"],
				"cors":         convertStorageBucketCorsList(r.AttributeValues["cors"]),
				"lifecycle":    convertStorageBucketLifecycleList(r.AttributeValues["llifecyle"]),
				"logging":      convertStorageBucketLoggingList(r.AttributeValues["logging"]),
				"project":      r.AttributeValues["project"],
				"storageClass": r.AttributeValues["storage_class"],
				"versioning":   convertStorageBucketVersioningList(r.AttributeValues["versioning"]),
				"website":      convertStorageBucketWebsiteList(r.AttributeValues["website"]),
				"":             r.AttributeValues["force_destroy"],
			}
			b, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			r := &storage.Bucket{}
			if err := json.Unmarshal(b, r); err != nil {
				return nil, err
			}
			resources = append(resources, r)
		}
	}
	return resources, nil
}

// AccesscontextmanagerAccessLevelAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func AccesscontextmanagerAccessLevelAsHCL(r accesscontextmanager.AccessLevel) (string, error) {
	outputConfig := "resource \"google_accesscontextmanager_access_level\" \"output\" {\n"
	if r.Name != nil {
		outputConfig += fmt.Sprintf("\tname = %#v\n", *r.Name)
	}
	if r.Title != nil {
		outputConfig += fmt.Sprintf("\ttitle = %#v\n", *r.Title)
	}
	if v := convertAccesscontextmanagerAccessLevelBasicToHCL(r.Basic); v != "" {
		outputConfig += fmt.Sprintf("\tbasic = %s\n", v)
	}
	if r.Description != nil {
		outputConfig += fmt.Sprintf("\tdescription = %#v\n", *r.Description)
	}
	if r.Policy != nil {
		outputConfig += fmt.Sprintf("\tpolicy = %#v\n", *r.Policy)
	}
	return formatHCL(outputConfig + "}")
}
func convertAccesscontextmanagerAccessLevelBasicToHCL(r *accesscontextmanager.AccessLevelBasic) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.CombiningFunction != nil {
		outputConfig += fmt.Sprintf("\tcombining_function = %#v\n", *r.CombiningFunction)
	}
	if r.Conditions != nil {
		for _, v := range r.Conditions {
			outputConfig += fmt.Sprintf("\tconditions %s\n", convertAccesscontextmanagerAccessLevelBasicConditionsToHCL(&v))
		}
	}
	return outputConfig + "}"
}
func convertAccesscontextmanagerAccessLevelBasicConditionsToHCL(r *accesscontextmanager.AccessLevelBasicConditions) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if v := convertAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyToHCL(r.DevicePolicy); v != "" {
		outputConfig += fmt.Sprintf("\tdevice_policy = %s\n", v)
	}
	if r.IPSubnetworks != nil {
		outputConfig += "\tip_subnetworks = ["
		for _, v := range r.IPSubnetworks {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.Members != nil {
		outputConfig += "\tmembers = ["
		for _, v := range r.Members {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.Negate != nil {
		outputConfig += fmt.Sprintf("\tnegate = %#v\n", *r.Negate)
	}
	if r.Regions != nil {
		outputConfig += "\tregions = ["
		for _, v := range r.Regions {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.RequiredAccessLevels != nil {
		outputConfig += "\trequired_access_levels = ["
		for _, v := range r.RequiredAccessLevels {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	return outputConfig + "}"
}
func convertAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyToHCL(r *accesscontextmanager.AccessLevelBasicConditionsDevicePolicy) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.AllowedDeviceManagementLevels != nil {
		outputConfig += "\tallowed_device_management_levels = ["
		for _, v := range r.AllowedDeviceManagementLevels {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.AllowedEncryptionStatuses != nil {
		outputConfig += "\tallowed_encryption_statuses = ["
		for _, v := range r.AllowedEncryptionStatuses {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.OsConstraints != nil {
		for _, v := range r.OsConstraints {
			outputConfig += fmt.Sprintf("\tos_constraints %s\n", convertAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOsConstraintsToHCL(&v))
		}
	}
	if r.RequireAdminApproval != nil {
		outputConfig += fmt.Sprintf("\trequire_admin_approval = %#v\n", *r.RequireAdminApproval)
	}
	if r.RequireCorpOwned != nil {
		outputConfig += fmt.Sprintf("\trequire_corp_owned = %#v\n", *r.RequireCorpOwned)
	}
	if r.RequireScreenlock != nil {
		outputConfig += fmt.Sprintf("\trequire_screenlock = %#v\n", *r.RequireScreenlock)
	}
	return outputConfig + "}"
}
func convertAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOsConstraintsToHCL(r *accesscontextmanager.AccessLevelBasicConditionsDevicePolicyOsConstraints) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.MinimumVersion != nil {
		outputConfig += fmt.Sprintf("\tminimum_version = %#v\n", *r.MinimumVersion)
	}
	if r.OsType != nil {
		outputConfig += fmt.Sprintf("\tos_type = %#v\n", *r.OsType)
	}
	if r.RequireVerifiedChromeOs != nil {
		outputConfig += fmt.Sprintf("\trequire_verified_chrome_os = %#v\n", *r.RequireVerifiedChromeOs)
	}
	return outputConfig + "}"
}

// AccesscontextmanagerAccessPolicyAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func AccesscontextmanagerAccessPolicyAsHCL(r accesscontextmanager.AccessPolicy) (string, error) {
	outputConfig := "resource \"google_accesscontextmanager_access_policy\" \"output\" {\n"
	if r.Parent != nil {
		outputConfig += fmt.Sprintf("\tparent = %#v\n", *r.Parent)
	}
	if r.Title != nil {
		outputConfig += fmt.Sprintf("\ttitle = %#v\n", *r.Title)
	}
	return formatHCL(outputConfig + "}")
}

// AccesscontextmanagerServicePerimeterAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func AccesscontextmanagerServicePerimeterAsHCL(r accesscontextmanager.ServicePerimeter) (string, error) {
	outputConfig := "resource \"google_accesscontextmanager_service_perimeter\" \"output\" {\n"
	if r.Name != nil {
		outputConfig += fmt.Sprintf("\tname = %#v\n", *r.Name)
	}
	if r.Policy != nil {
		outputConfig += fmt.Sprintf("\tpolicy = %#v\n", *r.Policy)
	}
	if r.Title != nil {
		outputConfig += fmt.Sprintf("\ttitle = %#v\n", *r.Title)
	}
	if r.Description != nil {
		outputConfig += fmt.Sprintf("\tdescription = %#v\n", *r.Description)
	}
	if r.PerimeterType != nil {
		outputConfig += fmt.Sprintf("\tperimeter_type = %#v\n", *r.PerimeterType)
	}
	if v := convertAccesscontextmanagerServicePerimeterSpecToHCL(r.Spec); v != "" {
		outputConfig += fmt.Sprintf("\tspec = %s\n", v)
	}
	if v := convertAccesscontextmanagerServicePerimeterStatusToHCL(r.Status); v != "" {
		outputConfig += fmt.Sprintf("\tstatus = %s\n", v)
	}
	if r.UseExplicitDryRunSpec != nil {
		outputConfig += fmt.Sprintf("\tuse_explicit_dry_run_spec = %#v\n", *r.UseExplicitDryRunSpec)
	}
	return formatHCL(outputConfig + "}")
}
func convertAccesscontextmanagerServicePerimeterSpecToHCL(r *accesscontextmanager.ServicePerimeterSpec) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.AccessLevels != nil {
		outputConfig += "\taccess_levels = ["
		for _, v := range r.AccessLevels {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.Resources != nil {
		outputConfig += "\tresources = ["
		for _, v := range r.Resources {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.RestrictedServices != nil {
		outputConfig += "\trestricted_services = ["
		for _, v := range r.RestrictedServices {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if v := convertAccesscontextmanagerServicePerimeterSpecVPCAccessibleServicesToHCL(r.VPCAccessibleServices); v != "" {
		outputConfig += fmt.Sprintf("\tvpc_accessible_services = %s\n", v)
	}
	return outputConfig + "}"
}
func convertAccesscontextmanagerServicePerimeterSpecVPCAccessibleServicesToHCL(r *accesscontextmanager.ServicePerimeterSpecVPCAccessibleServices) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.AllowedServices != nil {
		outputConfig += "\tallowed_services = ["
		for _, v := range r.AllowedServices {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.EnableRestriction != nil {
		outputConfig += fmt.Sprintf("\tenable_restriction = %#v\n", *r.EnableRestriction)
	}
	return outputConfig + "}"
}
func convertAccesscontextmanagerServicePerimeterStatusToHCL(r *accesscontextmanager.ServicePerimeterStatus) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.AccessLevels != nil {
		outputConfig += "\taccess_levels = ["
		for _, v := range r.AccessLevels {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.Resources != nil {
		outputConfig += "\tresources = ["
		for _, v := range r.Resources {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.RestrictedServices != nil {
		outputConfig += "\trestricted_services = ["
		for _, v := range r.RestrictedServices {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if v := convertAccesscontextmanagerServicePerimeterStatusVPCAccessibleServicesToHCL(r.VPCAccessibleServices); v != "" {
		outputConfig += fmt.Sprintf("\tvpc_accessible_services = %s\n", v)
	}
	return outputConfig + "}"
}
func convertAccesscontextmanagerServicePerimeterStatusVPCAccessibleServicesToHCL(r *accesscontextmanager.ServicePerimeterStatusVPCAccessibleServices) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.AllowedServices != nil {
		outputConfig += "\tallowed_services = ["
		for _, v := range r.AllowedServices {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.EnableRestriction != nil {
		outputConfig += fmt.Sprintf("\tenable_restriction = %#v\n", *r.EnableRestriction)
	}
	return outputConfig + "}"
}

// ComputeBackendBucketAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func ComputeBackendBucketAsHCL(r compute.BackendBucket) (string, error) {
	outputConfig := "resource \"google_compute_backend_bucket\" \"output\" {\n"
	if r.BucketName != nil {
		outputConfig += fmt.Sprintf("\tbucket_name = %#v\n", *r.BucketName)
	}
	if v := convertComputeBackendBucketCdnPolicyToHCL(r.CdnPolicy); v != "" {
		outputConfig += fmt.Sprintf("\tcdn_policy = %s\n", v)
	}
	if r.Description != nil {
		outputConfig += fmt.Sprintf("\tdescription = %#v\n", *r.Description)
	}
	if r.EnableCdn != nil {
		outputConfig += fmt.Sprintf("\tenable_cdn = %#v\n", *r.EnableCdn)
	}
	if r.Name != nil {
		outputConfig += fmt.Sprintf("\tname = %#v\n", *r.Name)
	}
	if r.Project != nil {
		outputConfig += fmt.Sprintf("\tproject = %#v\n", *r.Project)
	}
	return formatHCL(outputConfig + "}")
}
func convertComputeBackendBucketCdnPolicyToHCL(r *compute.BackendBucketCdnPolicy) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.SignedUrlCacheMaxAgeSec != nil {
		outputConfig += fmt.Sprintf("\tsigned_url_cache_max_age_sec = %#v\n", *r.SignedUrlCacheMaxAgeSec)
	}
	return outputConfig + "}"
}

// ComputeFirewallAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func ComputeFirewallAsHCL(r compute.Firewall) (string, error) {
	outputConfig := "resource \"google_compute_firewall\" \"output\" {\n"
	if r.Name != nil {
		outputConfig += fmt.Sprintf("\tname = %#v\n", *r.Name)
	}
	if r.Network != nil {
		outputConfig += fmt.Sprintf("\tnetwork = %#v\n", *r.Network)
	}
	if r.Allowed != nil {
		for _, v := range r.Allowed {
			outputConfig += fmt.Sprintf("\tallow %s\n", convertComputeFirewallAllowedToHCL(&v))
		}
	}
	if r.Denied != nil {
		for _, v := range r.Denied {
			outputConfig += fmt.Sprintf("\tdeny %s\n", convertComputeFirewallDeniedToHCL(&v))
		}
	}
	if r.Description != nil {
		outputConfig += fmt.Sprintf("\tdescription = %#v\n", *r.Description)
	}
	if r.DestinationRanges != nil {
		outputConfig += "\tdestination_ranges = ["
		for _, v := range r.DestinationRanges {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.Direction != nil {
		outputConfig += fmt.Sprintf("\tdirection = %#v\n", *r.Direction)
	}
	if r.Disabled != nil {
		outputConfig += fmt.Sprintf("\tdisabled = %#v\n", *r.Disabled)
	}
	if v := convertComputeFirewallLogConfigToHCL(r.LogConfig); v != "" {
		outputConfig += fmt.Sprintf("\tlog_config = %s\n", v)
	}
	if r.Priority != nil {
		outputConfig += fmt.Sprintf("\tpriority = %#v\n", *r.Priority)
	}
	if r.Project != nil {
		outputConfig += fmt.Sprintf("\tproject = %#v\n", *r.Project)
	}
	if r.SourceRanges != nil {
		outputConfig += "\tsource_ranges = ["
		for _, v := range r.SourceRanges {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.SourceServiceAccounts != nil {
		outputConfig += "\tsource_service_accounts = ["
		for _, v := range r.SourceServiceAccounts {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.SourceTags != nil {
		outputConfig += "\tsource_tags = ["
		for _, v := range r.SourceTags {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.TargetServiceAccounts != nil {
		outputConfig += "\ttarget_service_accounts = ["
		for _, v := range r.TargetServiceAccounts {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.TargetTags != nil {
		outputConfig += "\ttarget_tags = ["
		for _, v := range r.TargetTags {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	return formatHCL(outputConfig + "}")
}
func convertComputeFirewallAllowedToHCL(r *compute.FirewallAllowed) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Ports != nil {
		outputConfig += "\tports = ["
		for _, v := range r.Ports {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.IPProtocol != nil {
		outputConfig += fmt.Sprintf("\tprotocol = %#v\n", *r.IPProtocol)
	}
	return outputConfig + "}"
}
func convertComputeFirewallDeniedToHCL(r *compute.FirewallDenied) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Ports != nil {
		outputConfig += "\tports = ["
		for _, v := range r.Ports {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.IPProtocol != nil {
		outputConfig += fmt.Sprintf("\tprotocol = %#v\n", *r.IPProtocol)
	}
	return outputConfig + "}"
}
func convertComputeFirewallLogConfigToHCL(r *compute.FirewallLogConfig) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Enable != nil {
		outputConfig += fmt.Sprintf("\tenable_logging = %#v\n", *r.Enable)
	}
	return outputConfig + "}"
}

// ComputeImageAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func ComputeImageAsHCL(r compute.Image) (string, error) {
	outputConfig := "resource \"google_compute_image\" \"output\" {\n"
	if v := convertComputeImageDeprecatedToHCL(r.Deprecated); v != "" {
		outputConfig += fmt.Sprintf("\tdeprecated = %s\n", v)
	}
	if r.Description != nil {
		outputConfig += fmt.Sprintf("\tdescription = %#v\n", *r.Description)
	}
	if r.DiskSizeGb != nil {
		outputConfig += fmt.Sprintf("\tdisk_size_gb = %#v\n", *r.DiskSizeGb)
	}
	if r.Family != nil {
		outputConfig += fmt.Sprintf("\tfamily = %#v\n", *r.Family)
	}
	if r.GuestOsFeature != nil {
		for _, v := range r.GuestOsFeature {
			outputConfig += fmt.Sprintf("\tguest_os_features %s\n", convertComputeImageGuestOsFeatureToHCL(&v))
		}
	}
	if v := convertComputeImageImageEncryptionKeyToHCL(r.ImageEncryptionKey); v != "" {
		outputConfig += fmt.Sprintf("\timage_encryption_key = %s\n", v)
	}
	if r.License != nil {
		outputConfig += "\tlicenses = ["
		for _, v := range r.License {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.Name != nil {
		outputConfig += fmt.Sprintf("\tname = %#v\n", *r.Name)
	}
	if r.Project != nil {
		outputConfig += fmt.Sprintf("\tproject = %#v\n", *r.Project)
	}
	if v := convertComputeImageRawDiskToHCL(r.RawDisk); v != "" {
		outputConfig += fmt.Sprintf("\traw_disk = %s\n", v)
	}
	if v := convertComputeImageShieldedInstanceInitialStateToHCL(r.ShieldedInstanceInitialState); v != "" {
		outputConfig += fmt.Sprintf("\tshielded_instance_initial_state = %s\n", v)
	}
	if r.SourceDisk != nil {
		outputConfig += fmt.Sprintf("\tsource_disk = %#v\n", *r.SourceDisk)
	}
	if v := convertComputeImageSourceDiskEncryptionKeyToHCL(r.SourceDiskEncryptionKey); v != "" {
		outputConfig += fmt.Sprintf("\tsource_disk_encryption_key = %s\n", v)
	}
	if r.SourceImage != nil {
		outputConfig += fmt.Sprintf("\tsource_image = %#v\n", *r.SourceImage)
	}
	if v := convertComputeImageSourceImageEncryptionKeyToHCL(r.SourceImageEncryptionKey); v != "" {
		outputConfig += fmt.Sprintf("\tsource_image_encryption_key = %s\n", v)
	}
	if r.SourceImageId != nil {
		outputConfig += fmt.Sprintf("\tsource_image_id = %#v\n", *r.SourceImageId)
	}
	if r.SourceSnapshot != nil {
		outputConfig += fmt.Sprintf("\tsource_snapshot = %#v\n", *r.SourceSnapshot)
	}
	if v := convertComputeImageSourceSnapshotEncryptionKeyToHCL(r.SourceSnapshotEncryptionKey); v != "" {
		outputConfig += fmt.Sprintf("\tsource_snapshot_encryption_key = %s\n", v)
	}
	if r.SourceSnapshotId != nil {
		outputConfig += fmt.Sprintf("\tsource_snapshot_id = %#v\n", *r.SourceSnapshotId)
	}
	if r.SourceType != nil {
		outputConfig += fmt.Sprintf("\tsource_type = %#v\n", *r.SourceType)
	}
	if r.StorageLocation != nil {
		outputConfig += "\tstorage_location = ["
		for _, v := range r.StorageLocation {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	return formatHCL(outputConfig + "}")
}
func convertComputeImageDeprecatedToHCL(r *compute.ImageDeprecated) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Deleted != nil {
		outputConfig += fmt.Sprintf("\tdeleted = %#v\n", *r.Deleted)
	}
	if r.Deprecated != nil {
		outputConfig += fmt.Sprintf("\tdeprecated = %#v\n", *r.Deprecated)
	}
	if r.Obsolete != nil {
		outputConfig += fmt.Sprintf("\tobsolete = %#v\n", *r.Obsolete)
	}
	if r.Replacement != nil {
		outputConfig += fmt.Sprintf("\treplacement = %#v\n", *r.Replacement)
	}
	if r.State != nil {
		outputConfig += fmt.Sprintf("\tstate = %#v\n", *r.State)
	}
	return outputConfig + "}"
}
func convertComputeImageGuestOsFeatureToHCL(r *compute.ImageGuestOsFeature) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Type != nil {
		outputConfig += fmt.Sprintf("\ttype = %#v\n", *r.Type)
	}
	return outputConfig + "}"
}
func convertComputeImageImageEncryptionKeyToHCL(r *compute.ImageImageEncryptionKey) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.KmsKeyName != nil {
		outputConfig += fmt.Sprintf("\tkms_key_name = %#v\n", *r.KmsKeyName)
	}
	if r.KmsKeyServiceAccount != nil {
		outputConfig += fmt.Sprintf("\tkms_key_service_account = %#v\n", *r.KmsKeyServiceAccount)
	}
	if r.RawKey != nil {
		outputConfig += fmt.Sprintf("\traw_key = %#v\n", *r.RawKey)
	}
	if r.Sha256 != nil {
		outputConfig += fmt.Sprintf("\tsha256 = %#v\n", *r.Sha256)
	}
	return outputConfig + "}"
}
func convertComputeImageRawDiskToHCL(r *compute.ImageRawDisk) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.ContainerType != nil {
		outputConfig += fmt.Sprintf("\tcontainer_type = %#v\n", *r.ContainerType)
	}
	if r.Sha1Checksum != nil {
		outputConfig += fmt.Sprintf("\tsha1_checksum = %#v\n", *r.Sha1Checksum)
	}
	if r.Source != nil {
		outputConfig += fmt.Sprintf("\tsource = %#v\n", *r.Source)
	}
	return outputConfig + "}"
}
func convertComputeImageShieldedInstanceInitialStateToHCL(r *compute.ImageShieldedInstanceInitialState) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Db != nil {
		for _, v := range r.Db {
			outputConfig += fmt.Sprintf("\tdb %s\n", convertComputeImageShieldedInstanceInitialStateDbToHCL(&v))
		}
	}
	if r.Dbx != nil {
		for _, v := range r.Dbx {
			outputConfig += fmt.Sprintf("\tdbx %s\n", convertComputeImageShieldedInstanceInitialStateDbxToHCL(&v))
		}
	}
	if r.Kek != nil {
		for _, v := range r.Kek {
			outputConfig += fmt.Sprintf("\tkek %s\n", convertComputeImageShieldedInstanceInitialStateKekToHCL(&v))
		}
	}
	if v := convertComputeImageShieldedInstanceInitialStatePkToHCL(r.Pk); v != "" {
		outputConfig += fmt.Sprintf("\tpk = %s\n", v)
	}
	return outputConfig + "}"
}
func convertComputeImageShieldedInstanceInitialStateDbToHCL(r *compute.ImageShieldedInstanceInitialStateDb) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Content != nil {
		outputConfig += fmt.Sprintf("\tcontent = %#v\n", *r.Content)
	}
	if r.FileType != nil {
		outputConfig += fmt.Sprintf("\tfile_type = %#v\n", *r.FileType)
	}
	return outputConfig + "}"
}
func convertComputeImageShieldedInstanceInitialStateDbxToHCL(r *compute.ImageShieldedInstanceInitialStateDbx) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Content != nil {
		outputConfig += fmt.Sprintf("\tcontent = %#v\n", *r.Content)
	}
	if r.FileType != nil {
		outputConfig += fmt.Sprintf("\tfile_type = %#v\n", *r.FileType)
	}
	return outputConfig + "}"
}
func convertComputeImageShieldedInstanceInitialStateKekToHCL(r *compute.ImageShieldedInstanceInitialStateKek) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Content != nil {
		outputConfig += fmt.Sprintf("\tcontent = %#v\n", *r.Content)
	}
	if r.FileType != nil {
		outputConfig += fmt.Sprintf("\tfile_type = %#v\n", *r.FileType)
	}
	return outputConfig + "}"
}
func convertComputeImageShieldedInstanceInitialStatePkToHCL(r *compute.ImageShieldedInstanceInitialStatePk) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Content != nil {
		outputConfig += fmt.Sprintf("\tcontent = %#v\n", *r.Content)
	}
	if r.FileType != nil {
		outputConfig += fmt.Sprintf("\tfile_type = %#v\n", *r.FileType)
	}
	return outputConfig + "}"
}
func convertComputeImageSourceDiskEncryptionKeyToHCL(r *compute.ImageSourceDiskEncryptionKey) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.KmsKeyName != nil {
		outputConfig += fmt.Sprintf("\tkms_key_name = %#v\n", *r.KmsKeyName)
	}
	if r.KmsKeyServiceAccount != nil {
		outputConfig += fmt.Sprintf("\tkms_key_service_account = %#v\n", *r.KmsKeyServiceAccount)
	}
	if r.RawKey != nil {
		outputConfig += fmt.Sprintf("\traw_key = %#v\n", *r.RawKey)
	}
	if r.Sha256 != nil {
		outputConfig += fmt.Sprintf("\tsha256 = %#v\n", *r.Sha256)
	}
	return outputConfig + "}"
}
func convertComputeImageSourceImageEncryptionKeyToHCL(r *compute.ImageSourceImageEncryptionKey) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.KmsKeyName != nil {
		outputConfig += fmt.Sprintf("\tkms_key_name = %#v\n", *r.KmsKeyName)
	}
	if r.KmsKeyServiceAccount != nil {
		outputConfig += fmt.Sprintf("\tkms_key_service_account = %#v\n", *r.KmsKeyServiceAccount)
	}
	if r.RawKey != nil {
		outputConfig += fmt.Sprintf("\traw_key = %#v\n", *r.RawKey)
	}
	if r.Sha256 != nil {
		outputConfig += fmt.Sprintf("\tsha256 = %#v\n", *r.Sha256)
	}
	return outputConfig + "}"
}
func convertComputeImageSourceSnapshotEncryptionKeyToHCL(r *compute.ImageSourceSnapshotEncryptionKey) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.KmsKeyName != nil {
		outputConfig += fmt.Sprintf("\tkms_key_name = %#v\n", *r.KmsKeyName)
	}
	if r.KmsKeyServiceAccount != nil {
		outputConfig += fmt.Sprintf("\tkms_key_service_account = %#v\n", *r.KmsKeyServiceAccount)
	}
	if r.RawKey != nil {
		outputConfig += fmt.Sprintf("\traw_key = %#v\n", *r.RawKey)
	}
	if r.Sha256 != nil {
		outputConfig += fmt.Sprintf("\tsha256 = %#v\n", *r.Sha256)
	}
	return outputConfig + "}"
}

// ComputeNetworkAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func ComputeNetworkAsHCL(r compute.Network) (string, error) {
	outputConfig := "resource \"google_compute_network\" \"output\" {\n"
	if r.Name != nil {
		outputConfig += fmt.Sprintf("\tname = %#v\n", *r.Name)
	}
	if r.AutoCreateSubnetworks != nil {
		outputConfig += fmt.Sprintf("\tauto_create_subnetworks = %#v\n", *r.AutoCreateSubnetworks)
	}
	if r.Description != nil {
		outputConfig += fmt.Sprintf("\tdescription = %#v\n", *r.Description)
	}
	if r.Project != nil {
		outputConfig += fmt.Sprintf("\tproject = %#v\n", *r.Project)
	}
	if v := convertComputeNetworkRoutingConfigToHCL(r.RoutingConfig); v != "" {
		outputConfig += fmt.Sprintf("\trouting_config = %s\n", v)
	}
	return formatHCL(outputConfig + "}")
}
func convertComputeNetworkRoutingConfigToHCL(r *compute.NetworkRoutingConfig) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.RoutingMode != nil {
		outputConfig += fmt.Sprintf("\trouting_mode = %#v\n", *r.RoutingMode)
	}
	return outputConfig + "}"
}

// ComputeRouteAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func ComputeRouteAsHCL(r compute.Route) (string, error) {
	outputConfig := "resource \"google_compute_route\" \"output\" {\n"
	if r.DestRange != nil {
		outputConfig += fmt.Sprintf("\tdest_range = %#v\n", *r.DestRange)
	}
	if r.Name != nil {
		outputConfig += fmt.Sprintf("\tname = %#v\n", *r.Name)
	}
	if r.Network != nil {
		outputConfig += fmt.Sprintf("\tnetwork = %#v\n", *r.Network)
	}
	if r.Description != nil {
		outputConfig += fmt.Sprintf("\tdescription = %#v\n", *r.Description)
	}
	if r.NextHopGateway != nil {
		outputConfig += fmt.Sprintf("\tnext_hop_gateway = %#v\n", *r.NextHopGateway)
	}
	if r.NextHopIlb != nil {
		outputConfig += fmt.Sprintf("\tnext_hop_ilb = %#v\n", *r.NextHopIlb)
	}
	if r.NextHopInstance != nil {
		outputConfig += fmt.Sprintf("\tnext_hop_instance = %#v\n", *r.NextHopInstance)
	}
	if r.NextHopIP != nil {
		outputConfig += fmt.Sprintf("\tnext_hop_ip = %#v\n", *r.NextHopIP)
	}
	if r.NextHopNetwork != nil {
		outputConfig += fmt.Sprintf("\tnext_hop_network = %#v\n", *r.NextHopNetwork)
	}
	if r.NextHopVpnTunnel != nil {
		outputConfig += fmt.Sprintf("\tnext_hop_vpn_tunnel = %#v\n", *r.NextHopVpnTunnel)
	}
	if r.Priority != nil {
		outputConfig += fmt.Sprintf("\tpriority = %#v\n", *r.Priority)
	}
	if r.Project != nil {
		outputConfig += fmt.Sprintf("\tproject = %#v\n", *r.Project)
	}
	if r.Tag != nil {
		outputConfig += "\ttags = ["
		for _, v := range r.Tag {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	return formatHCL(outputConfig + "}")
}
func convertComputeRouteWarningToHCL(r *compute.RouteWarning) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	return outputConfig + "}"
}

// ComputeSslPolicyAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func ComputeSslPolicyAsHCL(r compute.SslPolicy) (string, error) {
	outputConfig := "resource \"google_compute_ssl_policy\" \"output\" {\n"
	if r.Name != nil {
		outputConfig += fmt.Sprintf("\tname = %#v\n", *r.Name)
	}
	if r.CustomFeature != nil {
		outputConfig += "\tcustom_features = ["
		for _, v := range r.CustomFeature {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.Description != nil {
		outputConfig += fmt.Sprintf("\tdescription = %#v\n", *r.Description)
	}
	if r.MinTlsVersion != nil {
		outputConfig += fmt.Sprintf("\tmin_tls_version = %#v\n", *r.MinTlsVersion)
	}
	if r.Profile != nil {
		outputConfig += fmt.Sprintf("\tprofile = %#v\n", *r.Profile)
	}
	if r.Project != nil {
		outputConfig += fmt.Sprintf("\tproject = %#v\n", *r.Project)
	}
	return formatHCL(outputConfig + "}")
}
func convertComputeSslPolicyWarningToHCL(r *compute.SslPolicyWarning) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Code != nil {
		outputConfig += fmt.Sprintf("\tcode = %#v\n", *r.Code)
	}
	if r.Data != nil {
		for _, v := range r.Data {
			outputConfig += fmt.Sprintf("\tdata %s\n", convertComputeSslPolicyWarningDataToHCL(&v))
		}
	}
	if r.Message != nil {
		outputConfig += fmt.Sprintf("\tmessage = %#v\n", *r.Message)
	}
	return outputConfig + "}"
}
func convertComputeSslPolicyWarningDataToHCL(r *compute.SslPolicyWarningData) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Key != nil {
		outputConfig += fmt.Sprintf("\tkey = %#v\n", *r.Key)
	}
	if r.Value != nil {
		outputConfig += fmt.Sprintf("\tvalue = %#v\n", *r.Value)
	}
	return outputConfig + "}"
}

// ComputeSubnetworkAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func ComputeSubnetworkAsHCL(r compute.Subnetwork) (string, error) {
	outputConfig := "resource \"google_compute_subnetwork\" \"output\" {\n"
	if r.IPCidrRange != nil {
		outputConfig += fmt.Sprintf("\tip_cidr_range = %#v\n", *r.IPCidrRange)
	}
	if r.Name != nil {
		outputConfig += fmt.Sprintf("\tname = %#v\n", *r.Name)
	}
	if r.Network != nil {
		outputConfig += fmt.Sprintf("\tnetwork = %#v\n", *r.Network)
	}
	if r.Description != nil {
		outputConfig += fmt.Sprintf("\tdescription = %#v\n", *r.Description)
	}
	if r.EnableFlowLogs != nil {
		outputConfig += fmt.Sprintf("\tenable_flow_logs = %#v\n", *r.EnableFlowLogs)
	}
	if v := convertComputeSubnetworkLogConfigToHCL(r.LogConfig); v != "" {
		outputConfig += fmt.Sprintf("\tlog_config = %s\n", v)
	}
	if r.PrivateIPGoogleAccess != nil {
		outputConfig += fmt.Sprintf("\tprivate_ip_google_access = %#v\n", *r.PrivateIPGoogleAccess)
	}
	if r.Project != nil {
		outputConfig += fmt.Sprintf("\tproject = %#v\n", *r.Project)
	}
	if r.Purpose != nil {
		outputConfig += fmt.Sprintf("\tpurpose = %#v\n", *r.Purpose)
	}
	if r.Region != nil {
		outputConfig += fmt.Sprintf("\tregion = %#v\n", *r.Region)
	}
	if r.Role != nil {
		outputConfig += fmt.Sprintf("\trole = %#v\n", *r.Role)
	}
	if r.SecondaryIPRange != nil {
		for _, v := range r.SecondaryIPRange {
			outputConfig += fmt.Sprintf("\tsecondary_ip_range %s\n", convertComputeSubnetworkSecondaryIPRangeToHCL(&v))
		}
	}
	return formatHCL(outputConfig + "}")
}
func convertComputeSubnetworkLogConfigToHCL(r *compute.SubnetworkLogConfig) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.AggregationInterval != nil {
		outputConfig += fmt.Sprintf("\taggregation_interval = %#v\n", *r.AggregationInterval)
	}
	if r.FlowSampling != nil {
		outputConfig += fmt.Sprintf("\tflow_sampling = %#v\n", *r.FlowSampling)
	}
	if r.Metadata != nil {
		outputConfig += fmt.Sprintf("\tmetadata = %#v\n", *r.Metadata)
	}
	return outputConfig + "}"
}
func convertComputeSubnetworkSecondaryIPRangeToHCL(r *compute.SubnetworkSecondaryIPRange) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.IPCidrRange != nil {
		outputConfig += fmt.Sprintf("\tip_cidr_range = %#v\n", *r.IPCidrRange)
	}
	if r.RangeName != nil {
		outputConfig += fmt.Sprintf("\trange_name = %#v\n", *r.RangeName)
	}
	return outputConfig + "}"
}

// DnsManagedZoneAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func DnsManagedZoneAsHCL(r dns.ManagedZone) (string, error) {
	outputConfig := "resource \"google_dns_managed_zone\" \"output\" {\n"
	if r.DnsName != nil {
		outputConfig += fmt.Sprintf("\tdns_name = %#v\n", *r.DnsName)
	}
	if r.Name != nil {
		outputConfig += fmt.Sprintf("\tname = %#v\n", *r.Name)
	}
	if r.Description != nil {
		outputConfig += fmt.Sprintf("\tdescription = %#v\n", *r.Description)
	}
	if v := convertDnsManagedZoneDnssecConfigToHCL(r.DnssecConfig); v != "" {
		outputConfig += fmt.Sprintf("\tdnssec_config = %s\n", v)
	}
	if v := convertDnsManagedZoneForwardingConfigToHCL(r.ForwardingConfig); v != "" {
		outputConfig += fmt.Sprintf("\tforwarding_config = %s\n", v)
	}
	if v := convertDnsManagedZonePeeringConfigToHCL(r.PeeringConfig); v != "" {
		outputConfig += fmt.Sprintf("\tpeering_config = %s\n", v)
	}
	if v := convertDnsManagedZonePrivateVisibilityConfigToHCL(r.PrivateVisibilityConfig); v != "" {
		outputConfig += fmt.Sprintf("\tprivate_visibility_config = %s\n", v)
	}
	if r.Project != nil {
		outputConfig += fmt.Sprintf("\tproject = %#v\n", *r.Project)
	}
	if r.ReverseLookup != nil {
		outputConfig += fmt.Sprintf("\treverse_lookup = %#v\n", *r.ReverseLookup)
	}
	if r.Visibility != nil {
		outputConfig += fmt.Sprintf("\tvisibility = %#v\n", *r.Visibility)
	}
	return formatHCL(outputConfig + "}")
}
func convertDnsManagedZoneDnssecConfigToHCL(r *dns.ManagedZoneDnssecConfig) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.DefaultKeySpecs != nil {
		for _, v := range r.DefaultKeySpecs {
			outputConfig += fmt.Sprintf("\tdefault_key_specs %s\n", convertDnsManagedZoneDnssecConfigDefaultKeySpecsToHCL(&v))
		}
	}
	if r.NonExistence != nil {
		outputConfig += fmt.Sprintf("\tnon_existence = %#v\n", *r.NonExistence)
	}
	if r.State != nil {
		outputConfig += fmt.Sprintf("\tstate = %#v\n", *r.State)
	}
	return outputConfig + "}"
}
func convertDnsManagedZoneDnssecConfigDefaultKeySpecsToHCL(r *dns.ManagedZoneDnssecConfigDefaultKeySpecs) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Algorithm != nil {
		outputConfig += fmt.Sprintf("\talgorithm = %#v\n", *r.Algorithm)
	}
	if r.KeyLength != nil {
		outputConfig += fmt.Sprintf("\tkey_length = %#v\n", *r.KeyLength)
	}
	if r.KeyType != nil {
		outputConfig += fmt.Sprintf("\tkey_type = %#v\n", *r.KeyType)
	}
	return outputConfig + "}"
}
func convertDnsManagedZoneForwardingConfigToHCL(r *dns.ManagedZoneForwardingConfig) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.TargetNameServers != nil {
		for _, v := range r.TargetNameServers {
			outputConfig += fmt.Sprintf("\ttarget_name_servers %s\n", convertDnsManagedZoneForwardingConfigTargetNameServersToHCL(&v))
		}
	}
	return outputConfig + "}"
}
func convertDnsManagedZoneForwardingConfigTargetNameServersToHCL(r *dns.ManagedZoneForwardingConfigTargetNameServers) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.ForwardingPath != nil {
		outputConfig += fmt.Sprintf("\tforwarding_path = %#v\n", *r.ForwardingPath)
	}
	if r.IPv4Address != nil {
		outputConfig += fmt.Sprintf("\tipv4_address = %#v\n", *r.IPv4Address)
	}
	return outputConfig + "}"
}
func convertDnsManagedZonePeeringConfigToHCL(r *dns.ManagedZonePeeringConfig) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if v := convertDnsManagedZonePeeringConfigTargetNetworkToHCL(r.TargetNetwork); v != "" {
		outputConfig += fmt.Sprintf("\ttarget_network = %s\n", v)
	}
	return outputConfig + "}"
}
func convertDnsManagedZonePeeringConfigTargetNetworkToHCL(r *dns.ManagedZonePeeringConfigTargetNetwork) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.NetworkUrl != nil {
		outputConfig += fmt.Sprintf("\tnetwork_url = %#v\n", *r.NetworkUrl)
	}
	return outputConfig + "}"
}
func convertDnsManagedZonePrivateVisibilityConfigToHCL(r *dns.ManagedZonePrivateVisibilityConfig) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Networks != nil {
		for _, v := range r.Networks {
			outputConfig += fmt.Sprintf("\tnetworks %s\n", convertDnsManagedZonePrivateVisibilityConfigNetworksToHCL(&v))
		}
	}
	return outputConfig + "}"
}
func convertDnsManagedZonePrivateVisibilityConfigNetworksToHCL(r *dns.ManagedZonePrivateVisibilityConfigNetworks) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.NetworkUrl != nil {
		outputConfig += fmt.Sprintf("\tnetwork_url = %#v\n", *r.NetworkUrl)
	}
	return outputConfig + "}"
}

// PubsubSubscriptionAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func PubsubSubscriptionAsHCL(r pubsub.Subscription) (string, error) {
	outputConfig := "resource \"google_pubsub_subscription\" \"output\" {\n"
	if r.Name != nil {
		outputConfig += fmt.Sprintf("\tname = %#v\n", *r.Name)
	}
	if r.Topic != nil {
		outputConfig += fmt.Sprintf("\ttopic = %#v\n", *r.Topic)
	}
	if r.AckDeadlineSeconds != nil {
		outputConfig += fmt.Sprintf("\tack_deadline_seconds = %#v\n", *r.AckDeadlineSeconds)
	}
	if v := convertPubsubSubscriptionDeadLetterPolicyToHCL(r.DeadLetterPolicy); v != "" {
		outputConfig += fmt.Sprintf("\tdead_letter_policy = %s\n", v)
	}
	if v := convertPubsubSubscriptionExpirationPolicyToHCL(r.ExpirationPolicy); v != "" {
		outputConfig += fmt.Sprintf("\texpiration_policy = %s\n", v)
	}
	if r.MessageRetentionDuration != nil {
		outputConfig += fmt.Sprintf("\tmessage_retention_duration = %#v\n", *r.MessageRetentionDuration)
	}
	if r.Project != nil {
		outputConfig += fmt.Sprintf("\tproject = %#v\n", *r.Project)
	}
	if v := convertPubsubSubscriptionPushConfigToHCL(r.PushConfig); v != "" {
		outputConfig += fmt.Sprintf("\tpush_config = %s\n", v)
	}
	if r.RetainAckedMessages != nil {
		outputConfig += fmt.Sprintf("\tretain_acked_messages = %#v\n", *r.RetainAckedMessages)
	}
	return formatHCL(outputConfig + "}")
}
func convertPubsubSubscriptionDeadLetterPolicyToHCL(r *pubsub.SubscriptionDeadLetterPolicy) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.DeadLetterTopic != nil {
		outputConfig += fmt.Sprintf("\tdead_letter_topic = %#v\n", *r.DeadLetterTopic)
	}
	if r.MaxDeliveryAttempts != nil {
		outputConfig += fmt.Sprintf("\tmax_delivery_attempts = %#v\n", *r.MaxDeliveryAttempts)
	}
	return outputConfig + "}"
}
func convertPubsubSubscriptionExpirationPolicyToHCL(r *pubsub.SubscriptionExpirationPolicy) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Ttl != nil {
		outputConfig += fmt.Sprintf("\tttl = %#v\n", *r.Ttl)
	}
	return outputConfig + "}"
}
func convertPubsubSubscriptionPushConfigToHCL(r *pubsub.SubscriptionPushConfig) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if v := convertPubsubSubscriptionPushConfigOidcTokenToHCL(r.OidcToken); v != "" {
		outputConfig += fmt.Sprintf("\toidc_token = %s\n", v)
	}
	if r.PushEndpoint != nil {
		outputConfig += fmt.Sprintf("\tpush_endpoint = %#v\n", *r.PushEndpoint)
	}
	return outputConfig + "}"
}
func convertPubsubSubscriptionPushConfigOidcTokenToHCL(r *pubsub.SubscriptionPushConfigOidcToken) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Audience != nil {
		outputConfig += fmt.Sprintf("\taudience = %#v\n", *r.Audience)
	}
	if r.ServiceAccountEmail != nil {
		outputConfig += fmt.Sprintf("\tservice_account_email = %#v\n", *r.ServiceAccountEmail)
	}
	return outputConfig + "}"
}

// PubsubTopicAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func PubsubTopicAsHCL(r pubsub.Topic) (string, error) {
	outputConfig := "resource \"google_pubsub_topic\" \"output\" {\n"
	if r.Name != nil {
		outputConfig += fmt.Sprintf("\tname = %#v\n", *r.Name)
	}
	if r.KmsKeyName != nil {
		outputConfig += fmt.Sprintf("\tkms_key_name = %#v\n", *r.KmsKeyName)
	}
	if v := convertPubsubTopicMessageStoragePolicyToHCL(r.MessageStoragePolicy); v != "" {
		outputConfig += fmt.Sprintf("\tmessage_storage_policy = %s\n", v)
	}
	if r.Project != nil {
		outputConfig += fmt.Sprintf("\tproject = %#v\n", *r.Project)
	}
	return formatHCL(outputConfig + "}")
}
func convertPubsubTopicMessageStoragePolicyToHCL(r *pubsub.TopicMessageStoragePolicy) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.AllowedPersistenceRegions != nil {
		outputConfig += "\tallowed_persistence_regions = ["
		for _, v := range r.AllowedPersistenceRegions {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	return outputConfig + "}"
}

// RedisInstanceAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func RedisInstanceAsHCL(r redis.Instance) (string, error) {
	outputConfig := "resource \"google_redis_instance\" \"output\" {\n"
	if r.MemorySizeGb != nil {
		outputConfig += fmt.Sprintf("\tmemory_size_gb = %#v\n", *r.MemorySizeGb)
	}
	if r.Name != nil {
		outputConfig += fmt.Sprintf("\tname = %#v\n", *r.Name)
	}
	if r.AlternativeLocationId != nil {
		outputConfig += fmt.Sprintf("\talternative_location_id = %#v\n", *r.AlternativeLocationId)
	}
	if r.AuthorizedNetwork != nil {
		outputConfig += fmt.Sprintf("\tauthorized_network = %#v\n", *r.AuthorizedNetwork)
	}
	if r.DisplayName != nil {
		outputConfig += fmt.Sprintf("\tdisplay_name = %#v\n", *r.DisplayName)
	}
	if r.LocationId != nil {
		outputConfig += fmt.Sprintf("\tlocation_id = %#v\n", *r.LocationId)
	}
	if r.Project != nil {
		outputConfig += fmt.Sprintf("\tproject = %#v\n", *r.Project)
	}
	if r.RedisVersion != nil {
		outputConfig += fmt.Sprintf("\tredis_version = %#v\n", *r.RedisVersion)
	}
	if r.Region != nil {
		outputConfig += fmt.Sprintf("\tregion = %#v\n", *r.Region)
	}
	if r.ReservedIPRange != nil {
		outputConfig += fmt.Sprintf("\treserved_ip_range = %#v\n", *r.ReservedIPRange)
	}
	if r.Tier != nil {
		outputConfig += fmt.Sprintf("\ttier = %#v\n", *r.Tier)
	}
	return formatHCL(outputConfig + "}")
}

// StorageBucketAsHCL returns a string representation of the specified resource in HCL.
// The generated HCL will include every settable field as a literal - that is, no
// variables, no references.  This may not be the best possible representation, but
// the crucial point is that `terraform import; terraform apply` will not produce
// any changes.  We do not validate that the resource specified will pass terraform
// validation unless is an object returned from the API after an Apply.
func StorageBucketAsHCL(r storage.Bucket) (string, error) {
	outputConfig := "resource \"google_storage_bucket\" \"output\" {\n"
	if r.Location != nil {
		outputConfig += fmt.Sprintf("\tlocation = %#v\n", *r.Location)
	}
	if r.Name != nil {
		outputConfig += fmt.Sprintf("\tname = %#v\n", *r.Name)
	}
	if r.Cors != nil {
		for _, v := range r.Cors {
			outputConfig += fmt.Sprintf("\tcors %s\n", convertStorageBucketCorsToHCL(&v))
		}
	}
	if v := convertStorageBucketLifecycleToHCL(r.Lifecycle); v != "" {
		outputConfig += fmt.Sprintf("\tllifecyle = %s\n", v)
	}
	if v := convertStorageBucketLoggingToHCL(r.Logging); v != "" {
		outputConfig += fmt.Sprintf("\tlogging = %s\n", v)
	}
	if r.Project != nil {
		outputConfig += fmt.Sprintf("\tproject = %#v\n", *r.Project)
	}
	if r.StorageClass != nil {
		outputConfig += fmt.Sprintf("\tstorage_class = %#v\n", *r.StorageClass)
	}
	if v := convertStorageBucketVersioningToHCL(r.Versioning); v != "" {
		outputConfig += fmt.Sprintf("\tversioning = %s\n", v)
	}
	if v := convertStorageBucketWebsiteToHCL(r.Website); v != "" {
		outputConfig += fmt.Sprintf("\twebsite = %s\n", v)
	}
	return formatHCL(outputConfig + "}")
}
func convertStorageBucketCorsToHCL(r *storage.BucketCors) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.MaxAgeSeconds != nil {
		outputConfig += fmt.Sprintf("\tmax_age_seconds = %#v\n", *r.MaxAgeSeconds)
	}
	if r.Method != nil {
		outputConfig += "\tmethod = ["
		for _, v := range r.Method {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.Origin != nil {
		outputConfig += "\torigin = ["
		for _, v := range r.Origin {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.ResponseHeader != nil {
		outputConfig += "\tresponse_header = ["
		for _, v := range r.ResponseHeader {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	return outputConfig + "}"
}
func convertStorageBucketLifecycleToHCL(r *storage.BucketLifecycle) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Rule != nil {
		for _, v := range r.Rule {
			outputConfig += fmt.Sprintf("\trule %s\n", convertStorageBucketLifecycleRuleToHCL(&v))
		}
	}
	return outputConfig + "}"
}
func convertStorageBucketLifecycleRuleToHCL(r *storage.BucketLifecycleRule) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if v := convertStorageBucketLifecycleRuleActionToHCL(r.Action); v != "" {
		outputConfig += fmt.Sprintf("\taction = %s\n", v)
	}
	if v := convertStorageBucketLifecycleRuleConditionToHCL(r.Condition); v != "" {
		outputConfig += fmt.Sprintf("\tcondition = %s\n", v)
	}
	return outputConfig + "}"
}
func convertStorageBucketLifecycleRuleActionToHCL(r *storage.BucketLifecycleRuleAction) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.StorageClass != nil {
		outputConfig += fmt.Sprintf("\tstorage_class = %#v\n", *r.StorageClass)
	}
	if r.Type != nil {
		outputConfig += fmt.Sprintf("\ttype = %#v\n", *r.Type)
	}
	return outputConfig + "}"
}
func convertStorageBucketLifecycleRuleConditionToHCL(r *storage.BucketLifecycleRuleCondition) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Age != nil {
		outputConfig += fmt.Sprintf("\tage = %#v\n", *r.Age)
	}
	if r.CreatedBefore != nil {
		outputConfig += fmt.Sprintf("\tcreated_before = %#v\n", *r.CreatedBefore)
	}
	if r.MatchesStorageClass != nil {
		outputConfig += "\tmatches_storage_class = ["
		for _, v := range r.MatchesStorageClass {
			outputConfig += fmt.Sprintf("%#v, ", v)
		}
		outputConfig += "]\n"
	}
	if r.NumNewerVersions != nil {
		outputConfig += fmt.Sprintf("\tnum_newer_versions = %#v\n", *r.NumNewerVersions)
	}
	if r.WithState != nil {
		outputConfig += fmt.Sprintf("\twith_state = %#v\n", *r.WithState)
	}
	return outputConfig + "}"
}
func convertStorageBucketLoggingToHCL(r *storage.BucketLogging) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.LogBucket != nil {
		outputConfig += fmt.Sprintf("\tlog_bucket = %#v\n", *r.LogBucket)
	}
	if r.LogObjectPrefix != nil {
		outputConfig += fmt.Sprintf("\tlog_object_prefix = %#v\n", *r.LogObjectPrefix)
	}
	return outputConfig + "}"
}
func convertStorageBucketVersioningToHCL(r *storage.BucketVersioning) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.Enabled != nil {
		outputConfig += fmt.Sprintf("\tenabled = %#v\n", *r.Enabled)
	}
	return outputConfig + "}"
}
func convertStorageBucketWebsiteToHCL(r *storage.BucketWebsite) string {
	if r == nil {
		return ""
	}
	outputConfig := "{\n"
	if r.MainPageSuffix != nil {
		outputConfig += fmt.Sprintf("\tmain_page_suffix = %#v\n", *r.MainPageSuffix)
	}
	if r.NotFoundPage != nil {
		outputConfig += fmt.Sprintf("\tnot_found_page = %#v\n", *r.NotFoundPage)
	}
	return outputConfig + "}"
}
func convertAccesscontextmanagerAccessLevelBasic(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"combiningFunction": in["combining_function"],
		"conditions":        in["conditions"],
	}
}

func convertAccesscontextmanagerAccessLevelBasicList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertAccesscontextmanagerAccessLevelBasic(v))
	}
	return out
}
func convertAccesscontextmanagerAccessLevelBasicConditions(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"devicePolicy":         convertAccesscontextmanagerAccessLevelBasicConditionsDevicePolicy(in["device_policy"]),
		"iPSubnetworks":        in["ip_subnetworks"],
		"members":              in["members"],
		"negate":               in["negate"],
		"regions":              in["regions"],
		"requiredAccessLevels": in["required_access_levels"],
	}
}

func convertAccesscontextmanagerAccessLevelBasicConditionsList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertAccesscontextmanagerAccessLevelBasicConditions(v))
	}
	return out
}
func convertAccesscontextmanagerAccessLevelBasicConditionsDevicePolicy(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"allowedDeviceManagementLevels": in["allowed_device_management_levels"],
		"allowedEncryptionStatuses":     in["allowed_encryption_statuses"],
		"osConstraints":                 in["os_constraints"],
		"requireAdminApproval":          in["require_admin_approval"],
		"requireCorpOwned":              in["require_corp_owned"],
		"requireScreenlock":             in["require_screenlock"],
	}
}

func convertAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertAccesscontextmanagerAccessLevelBasicConditionsDevicePolicy(v))
	}
	return out
}
func convertAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOsConstraints(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"minimumVersion":          in["minimum_version"],
		"osType":                  in["os_type"],
		"requireVerifiedChromeOs": in["require_verified_chrome_os"],
	}
}

func convertAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOsConstraintsList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertAccesscontextmanagerAccessLevelBasicConditionsDevicePolicyOsConstraints(v))
	}
	return out
}
func convertAccesscontextmanagerServicePerimeterSpec(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"accessLevels":          in["access_levels"],
		"resources":             in["resources"],
		"restrictedServices":    in["restricted_services"],
		"vPCAccessibleServices": convertAccesscontextmanagerServicePerimeterSpecVPCAccessibleServices(in["vpc_accessible_services"]),
	}
}

func convertAccesscontextmanagerServicePerimeterSpecList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertAccesscontextmanagerServicePerimeterSpec(v))
	}
	return out
}
func convertAccesscontextmanagerServicePerimeterSpecVPCAccessibleServices(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"allowedServices":   in["allowed_services"],
		"enableRestriction": in["enable_restriction"],
	}
}

func convertAccesscontextmanagerServicePerimeterSpecVPCAccessibleServicesList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertAccesscontextmanagerServicePerimeterSpecVPCAccessibleServices(v))
	}
	return out
}
func convertAccesscontextmanagerServicePerimeterStatus(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"accessLevels":          in["access_levels"],
		"resources":             in["resources"],
		"restrictedServices":    in["restricted_services"],
		"vPCAccessibleServices": convertAccesscontextmanagerServicePerimeterStatusVPCAccessibleServices(in["vpc_accessible_services"]),
	}
}

func convertAccesscontextmanagerServicePerimeterStatusList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertAccesscontextmanagerServicePerimeterStatus(v))
	}
	return out
}
func convertAccesscontextmanagerServicePerimeterStatusVPCAccessibleServices(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"allowedServices":   in["allowed_services"],
		"enableRestriction": in["enable_restriction"],
	}
}

func convertAccesscontextmanagerServicePerimeterStatusVPCAccessibleServicesList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertAccesscontextmanagerServicePerimeterStatusVPCAccessibleServices(v))
	}
	return out
}
func convertComputeBackendBucketCdnPolicy(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"signedUrlCacheMaxAgeSec": in["signed_url_cache_max_age_sec"],
	}
}

func convertComputeBackendBucketCdnPolicyList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeBackendBucketCdnPolicy(v))
	}
	return out
}
func convertComputeFirewallAllowed(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"ports":      in["ports"],
		"iPProtocol": in["protocol"],
	}
}

func convertComputeFirewallAllowedList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeFirewallAllowed(v))
	}
	return out
}
func convertComputeFirewallDenied(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"ports":      in["ports"],
		"iPProtocol": in["protocol"],
	}
}

func convertComputeFirewallDeniedList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeFirewallDenied(v))
	}
	return out
}
func convertComputeFirewallLogConfig(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"enable": in["enable_logging"],
	}
}

func convertComputeFirewallLogConfigList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeFirewallLogConfig(v))
	}
	return out
}
func convertComputeImageDeprecated(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"deleted":     in["deleted"],
		"deprecated":  in["deprecated"],
		"obsolete":    in["obsolete"],
		"replacement": in["replacement"],
		"state":       in["state"],
	}
}

func convertComputeImageDeprecatedList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeImageDeprecated(v))
	}
	return out
}
func convertComputeImageGuestOsFeature(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"type": in["type"],
	}
}

func convertComputeImageGuestOsFeatureList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeImageGuestOsFeature(v))
	}
	return out
}
func convertComputeImageImageEncryptionKey(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"kmsKeyName":           in["kms_key_name"],
		"kmsKeyServiceAccount": in["kms_key_service_account"],
		"rawKey":               in["raw_key"],
		"sha256":               in["sha256"],
	}
}

func convertComputeImageImageEncryptionKeyList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeImageImageEncryptionKey(v))
	}
	return out
}
func convertComputeImageRawDisk(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"containerType": in["container_type"],
		"sha1Checksum":  in["sha1_checksum"],
		"source":        in["source"],
	}
}

func convertComputeImageRawDiskList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeImageRawDisk(v))
	}
	return out
}
func convertComputeImageShieldedInstanceInitialState(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"db":  in["db"],
		"dbx": in["dbx"],
		"kek": in["kek"],
		"pk":  convertComputeImageShieldedInstanceInitialStatePk(in["pk"]),
	}
}

func convertComputeImageShieldedInstanceInitialStateList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeImageShieldedInstanceInitialState(v))
	}
	return out
}
func convertComputeImageShieldedInstanceInitialStateDb(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"content":  in["content"],
		"fileType": in["file_type"],
	}
}

func convertComputeImageShieldedInstanceInitialStateDbList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeImageShieldedInstanceInitialStateDb(v))
	}
	return out
}
func convertComputeImageShieldedInstanceInitialStateDbx(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"content":  in["content"],
		"fileType": in["file_type"],
	}
}

func convertComputeImageShieldedInstanceInitialStateDbxList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeImageShieldedInstanceInitialStateDbx(v))
	}
	return out
}
func convertComputeImageShieldedInstanceInitialStateKek(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"content":  in["content"],
		"fileType": in["file_type"],
	}
}

func convertComputeImageShieldedInstanceInitialStateKekList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeImageShieldedInstanceInitialStateKek(v))
	}
	return out
}
func convertComputeImageShieldedInstanceInitialStatePk(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"content":  in["content"],
		"fileType": in["file_type"],
	}
}

func convertComputeImageShieldedInstanceInitialStatePkList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeImageShieldedInstanceInitialStatePk(v))
	}
	return out
}
func convertComputeImageSourceDiskEncryptionKey(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"kmsKeyName":           in["kms_key_name"],
		"kmsKeyServiceAccount": in["kms_key_service_account"],
		"rawKey":               in["raw_key"],
		"sha256":               in["sha256"],
	}
}

func convertComputeImageSourceDiskEncryptionKeyList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeImageSourceDiskEncryptionKey(v))
	}
	return out
}
func convertComputeImageSourceImageEncryptionKey(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"kmsKeyName":           in["kms_key_name"],
		"kmsKeyServiceAccount": in["kms_key_service_account"],
		"rawKey":               in["raw_key"],
		"sha256":               in["sha256"],
	}
}

func convertComputeImageSourceImageEncryptionKeyList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeImageSourceImageEncryptionKey(v))
	}
	return out
}
func convertComputeImageSourceSnapshotEncryptionKey(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"kmsKeyName":           in["kms_key_name"],
		"kmsKeyServiceAccount": in["kms_key_service_account"],
		"rawKey":               in["raw_key"],
		"sha256":               in["sha256"],
	}
}

func convertComputeImageSourceSnapshotEncryptionKeyList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeImageSourceSnapshotEncryptionKey(v))
	}
	return out
}
func convertComputeNetworkRoutingConfig(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"routingMode": in["routing_mode"],
	}
}

func convertComputeNetworkRoutingConfigList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeNetworkRoutingConfig(v))
	}
	return out
}
func convertComputeRouteWarning(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"code":    in["code"],
		"data":    in["data"],
		"message": in["message"],
	}
}

func convertComputeRouteWarningList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeRouteWarning(v))
	}
	return out
}
func convertComputeSslPolicyWarning(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"code":    in["code"],
		"data":    in["data"],
		"message": in["message"],
	}
}

func convertComputeSslPolicyWarningList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeSslPolicyWarning(v))
	}
	return out
}
func convertComputeSslPolicyWarningData(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"key":   in["key"],
		"value": in["value"],
	}
}

func convertComputeSslPolicyWarningDataList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeSslPolicyWarningData(v))
	}
	return out
}
func convertComputeSubnetworkLogConfig(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"aggregationInterval": in["aggregation_interval"],
		"flowSampling":        in["flow_sampling"],
		"metadata":            in["metadata"],
	}
}

func convertComputeSubnetworkLogConfigList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeSubnetworkLogConfig(v))
	}
	return out
}
func convertComputeSubnetworkSecondaryIPRange(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"iPCidrRange": in["ip_cidr_range"],
		"rangeName":   in["range_name"],
	}
}

func convertComputeSubnetworkSecondaryIPRangeList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertComputeSubnetworkSecondaryIPRange(v))
	}
	return out
}
func convertDnsManagedZoneDnssecConfig(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"defaultKeySpecs": in["default_key_specs"],
		"nonExistence":    in["non_existence"],
		"state":           in["state"],
	}
}

func convertDnsManagedZoneDnssecConfigList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertDnsManagedZoneDnssecConfig(v))
	}
	return out
}
func convertDnsManagedZoneDnssecConfigDefaultKeySpecs(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"algorithm": in["algorithm"],
		"keyLength": in["key_length"],
		"keyType":   in["key_type"],
	}
}

func convertDnsManagedZoneDnssecConfigDefaultKeySpecsList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertDnsManagedZoneDnssecConfigDefaultKeySpecs(v))
	}
	return out
}
func convertDnsManagedZoneForwardingConfig(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"targetNameServers": in["target_name_servers"],
	}
}

func convertDnsManagedZoneForwardingConfigList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertDnsManagedZoneForwardingConfig(v))
	}
	return out
}
func convertDnsManagedZoneForwardingConfigTargetNameServers(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"forwardingPath": in["forwarding_path"],
		"iPv4Address":    in["ipv4_address"],
	}
}

func convertDnsManagedZoneForwardingConfigTargetNameServersList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertDnsManagedZoneForwardingConfigTargetNameServers(v))
	}
	return out
}
func convertDnsManagedZonePeeringConfig(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"targetNetwork": convertDnsManagedZonePeeringConfigTargetNetwork(in["target_network"]),
	}
}

func convertDnsManagedZonePeeringConfigList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertDnsManagedZonePeeringConfig(v))
	}
	return out
}
func convertDnsManagedZonePeeringConfigTargetNetwork(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"networkUrl": in["network_url"],
	}
}

func convertDnsManagedZonePeeringConfigTargetNetworkList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertDnsManagedZonePeeringConfigTargetNetwork(v))
	}
	return out
}
func convertDnsManagedZonePrivateVisibilityConfig(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"networks": in["networks"],
	}
}

func convertDnsManagedZonePrivateVisibilityConfigList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertDnsManagedZonePrivateVisibilityConfig(v))
	}
	return out
}
func convertDnsManagedZonePrivateVisibilityConfigNetworks(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"networkUrl": in["network_url"],
	}
}

func convertDnsManagedZonePrivateVisibilityConfigNetworksList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertDnsManagedZonePrivateVisibilityConfigNetworks(v))
	}
	return out
}
func convertPubsubSubscriptionDeadLetterPolicy(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"deadLetterTopic":     in["dead_letter_topic"],
		"maxDeliveryAttempts": in["max_delivery_attempts"],
	}
}

func convertPubsubSubscriptionDeadLetterPolicyList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertPubsubSubscriptionDeadLetterPolicy(v))
	}
	return out
}
func convertPubsubSubscriptionExpirationPolicy(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"ttl": in["ttl"],
	}
}

func convertPubsubSubscriptionExpirationPolicyList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertPubsubSubscriptionExpirationPolicy(v))
	}
	return out
}
func convertPubsubSubscriptionPushConfig(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"attributes":   in["attributes"],
		"oidcToken":    convertPubsubSubscriptionPushConfigOidcToken(in["oidc_token"]),
		"pushEndpoint": in["push_endpoint"],
	}
}

func convertPubsubSubscriptionPushConfigList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertPubsubSubscriptionPushConfig(v))
	}
	return out
}
func convertPubsubSubscriptionPushConfigOidcToken(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"audience":            in["audience"],
		"serviceAccountEmail": in["service_account_email"],
	}
}

func convertPubsubSubscriptionPushConfigOidcTokenList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertPubsubSubscriptionPushConfigOidcToken(v))
	}
	return out
}
func convertPubsubTopicMessageStoragePolicy(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"allowedPersistenceRegions": in["allowed_persistence_regions"],
	}
}

func convertPubsubTopicMessageStoragePolicyList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertPubsubTopicMessageStoragePolicy(v))
	}
	return out
}
func convertStorageBucketCors(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"maxAgeSeconds":  in["max_age_seconds"],
		"method":         in["method"],
		"origin":         in["origin"],
		"responseHeader": in["response_header"],
	}
}

func convertStorageBucketCorsList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertStorageBucketCors(v))
	}
	return out
}
func convertStorageBucketLifecycle(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"rule": in["rule"],
	}
}

func convertStorageBucketLifecycleList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertStorageBucketLifecycle(v))
	}
	return out
}
func convertStorageBucketLifecycleRule(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"action":    convertStorageBucketLifecycleRuleAction(in["action"]),
		"condition": convertStorageBucketLifecycleRuleCondition(in["condition"]),
	}
}

func convertStorageBucketLifecycleRuleList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertStorageBucketLifecycleRule(v))
	}
	return out
}
func convertStorageBucketLifecycleRuleAction(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"storageClass": in["storage_class"],
		"type":         in["type"],
	}
}

func convertStorageBucketLifecycleRuleActionList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertStorageBucketLifecycleRuleAction(v))
	}
	return out
}
func convertStorageBucketLifecycleRuleCondition(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"age":                 in["age"],
		"createdBefore":       in["created_before"],
		"matchesStorageClass": in["matches_storage_class"],
		"numNewerVersions":    in["num_newer_versions"],
		"withState":           in["with_state"],
	}
}

func convertStorageBucketLifecycleRuleConditionList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertStorageBucketLifecycleRuleCondition(v))
	}
	return out
}
func convertStorageBucketLogging(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"logBucket":       in["log_bucket"],
		"logObjectPrefix": in["log_object_prefix"],
	}
}

func convertStorageBucketLoggingList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertStorageBucketLogging(v))
	}
	return out
}
func convertStorageBucketVersioning(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"enabled": in["enabled"],
	}
}

func convertStorageBucketVersioningList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertStorageBucketVersioning(v))
	}
	return out
}
func convertStorageBucketWebsite(i interface{}) map[string]interface{} {
	if i == nil {
		return nil
	}
	in := i.(map[string]interface{})
	return map[string]interface{}{
		"mainPageSuffix": in["main_page_suffix"],
		"notFoundPage":   in["not_found_page"],
	}
}

func convertStorageBucketWebsiteList(i interface{}) (out []map[string]interface{}) {
	if i == nil {
		return nil
	}

	for _, v := range i.([]interface{}) {
		out = append(out, convertStorageBucketWebsite(v))
	}
	return out
}

func formatHCL(hcl string) (string, error) {
	b := bytes.Buffer{}
	r := strings.NewReader(hcl)
	if err := fmtcmd.Run(nil, nil, r, &b, fmtcmd.Options{}); err != nil {
		return "", err
	}
	return b.String(), nil
}
