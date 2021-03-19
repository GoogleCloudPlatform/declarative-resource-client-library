# Copyright 2021 Google LLC. All Rights Reserved.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
from connector import channel
from google3.cloud.graphite.mmv2.services.google.container import cluster_pb2
from google3.cloud.graphite.mmv2.services.google.container import cluster_pb2_grpc

from typing import List


class Cluster(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        initial_node_count: int = None,
        master_auth: dict = None,
        logging_service: str = None,
        monitoring_service: str = None,
        network: str = None,
        cluster_ipv4_cidr: str = None,
        addons_config: dict = None,
        subnetwork: str = None,
        node_pools: list = None,
        locations: list = None,
        enable_kubernetes_alpha: bool = None,
        resource_labels: dict = None,
        label_fingerprint: str = None,
        legacy_abac: dict = None,
        network_policy: dict = None,
        ip_allocation_policy: dict = None,
        master_authorized_networks_config: dict = None,
        binary_authorization: dict = None,
        autoscaling: dict = None,
        network_config: dict = None,
        maintenance_policy: dict = None,
        default_max_pods_constraint: dict = None,
        resource_usage_export_config: dict = None,
        authenticator_groups_config: dict = None,
        private_cluster_config: dict = None,
        database_encryption: dict = None,
        vertical_pod_autoscaling: dict = None,
        shielded_nodes: dict = None,
        endpoint: str = None,
        master_version: str = None,
        create_time: str = None,
        status: str = None,
        status_message: str = None,
        node_ipv4_cidr_size: int = None,
        services_ipv4_cidr: str = None,
        expire_time: str = None,
        location: str = None,
        enable_tpu: bool = None,
        tpu_ipv4_cidr_block: str = None,
        conditions: list = None,
        autopilot: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.initial_node_count = initial_node_count
        self.master_auth = master_auth
        self.logging_service = logging_service
        self.monitoring_service = monitoring_service
        self.network = network
        self.cluster_ipv4_cidr = cluster_ipv4_cidr
        self.addons_config = addons_config
        self.subnetwork = subnetwork
        self.locations = locations
        self.enable_kubernetes_alpha = enable_kubernetes_alpha
        self.resource_labels = resource_labels
        self.legacy_abac = legacy_abac
        self.network_policy = network_policy
        self.ip_allocation_policy = ip_allocation_policy
        self.master_authorized_networks_config = master_authorized_networks_config
        self.binary_authorization = binary_authorization
        self.autoscaling = autoscaling
        self.network_config = network_config
        self.maintenance_policy = maintenance_policy
        self.default_max_pods_constraint = default_max_pods_constraint
        self.resource_usage_export_config = resource_usage_export_config
        self.authenticator_groups_config = authenticator_groups_config
        self.private_cluster_config = private_cluster_config
        self.database_encryption = database_encryption
        self.vertical_pod_autoscaling = vertical_pod_autoscaling
        self.shielded_nodes = shielded_nodes
        self.master_version = master_version
        self.location = location
        self.enable_tpu = enable_tpu
        self.autopilot = autopilot
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = cluster_pb2_grpc.ContainerClusterServiceStub(channel.Channel())
        request = cluster_pb2.ApplyContainerClusterRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.initial_node_count):
            request.resource.initial_node_count = Primitive.to_proto(
                self.initial_node_count
            )

        if ClusterMasterAuth.to_proto(self.master_auth):
            request.resource.master_auth.CopyFrom(
                ClusterMasterAuth.to_proto(self.master_auth)
            )
        else:
            request.resource.ClearField("master_auth")
        if Primitive.to_proto(self.logging_service):
            request.resource.logging_service = Primitive.to_proto(self.logging_service)

        if Primitive.to_proto(self.monitoring_service):
            request.resource.monitoring_service = Primitive.to_proto(
                self.monitoring_service
            )

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.cluster_ipv4_cidr):
            request.resource.cluster_ipv4_cidr = Primitive.to_proto(
                self.cluster_ipv4_cidr
            )

        if ClusterAddonsConfig.to_proto(self.addons_config):
            request.resource.addons_config.CopyFrom(
                ClusterAddonsConfig.to_proto(self.addons_config)
            )
        else:
            request.resource.ClearField("addons_config")
        if Primitive.to_proto(self.subnetwork):
            request.resource.subnetwork = Primitive.to_proto(self.subnetwork)

        if Primitive.to_proto(self.locations):
            request.resource.locations.extend(Primitive.to_proto(self.locations))
        if Primitive.to_proto(self.enable_kubernetes_alpha):
            request.resource.enable_kubernetes_alpha = Primitive.to_proto(
                self.enable_kubernetes_alpha
            )

        if Primitive.to_proto(self.resource_labels):
            request.resource.resource_labels = Primitive.to_proto(self.resource_labels)

        if ClusterLegacyAbac.to_proto(self.legacy_abac):
            request.resource.legacy_abac.CopyFrom(
                ClusterLegacyAbac.to_proto(self.legacy_abac)
            )
        else:
            request.resource.ClearField("legacy_abac")
        if ClusterNetworkPolicy.to_proto(self.network_policy):
            request.resource.network_policy.CopyFrom(
                ClusterNetworkPolicy.to_proto(self.network_policy)
            )
        else:
            request.resource.ClearField("network_policy")
        if ClusterIPAllocationPolicy.to_proto(self.ip_allocation_policy):
            request.resource.ip_allocation_policy.CopyFrom(
                ClusterIPAllocationPolicy.to_proto(self.ip_allocation_policy)
            )
        else:
            request.resource.ClearField("ip_allocation_policy")
        if ClusterMasterAuthorizedNetworksConfig.to_proto(
            self.master_authorized_networks_config
        ):
            request.resource.master_authorized_networks_config.CopyFrom(
                ClusterMasterAuthorizedNetworksConfig.to_proto(
                    self.master_authorized_networks_config
                )
            )
        else:
            request.resource.ClearField("master_authorized_networks_config")
        if ClusterBinaryAuthorization.to_proto(self.binary_authorization):
            request.resource.binary_authorization.CopyFrom(
                ClusterBinaryAuthorization.to_proto(self.binary_authorization)
            )
        else:
            request.resource.ClearField("binary_authorization")
        if ClusterAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                ClusterAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if ClusterNetworkConfig.to_proto(self.network_config):
            request.resource.network_config.CopyFrom(
                ClusterNetworkConfig.to_proto(self.network_config)
            )
        else:
            request.resource.ClearField("network_config")
        if ClusterMaintenancePolicy.to_proto(self.maintenance_policy):
            request.resource.maintenance_policy.CopyFrom(
                ClusterMaintenancePolicy.to_proto(self.maintenance_policy)
            )
        else:
            request.resource.ClearField("maintenance_policy")
        if ClusterDefaultMaxPodsConstraint.to_proto(self.default_max_pods_constraint):
            request.resource.default_max_pods_constraint.CopyFrom(
                ClusterDefaultMaxPodsConstraint.to_proto(
                    self.default_max_pods_constraint
                )
            )
        else:
            request.resource.ClearField("default_max_pods_constraint")
        if ClusterResourceUsageExportConfig.to_proto(self.resource_usage_export_config):
            request.resource.resource_usage_export_config.CopyFrom(
                ClusterResourceUsageExportConfig.to_proto(
                    self.resource_usage_export_config
                )
            )
        else:
            request.resource.ClearField("resource_usage_export_config")
        if ClusterAuthenticatorGroupsConfig.to_proto(self.authenticator_groups_config):
            request.resource.authenticator_groups_config.CopyFrom(
                ClusterAuthenticatorGroupsConfig.to_proto(
                    self.authenticator_groups_config
                )
            )
        else:
            request.resource.ClearField("authenticator_groups_config")
        if ClusterPrivateClusterConfig.to_proto(self.private_cluster_config):
            request.resource.private_cluster_config.CopyFrom(
                ClusterPrivateClusterConfig.to_proto(self.private_cluster_config)
            )
        else:
            request.resource.ClearField("private_cluster_config")
        if ClusterDatabaseEncryption.to_proto(self.database_encryption):
            request.resource.database_encryption.CopyFrom(
                ClusterDatabaseEncryption.to_proto(self.database_encryption)
            )
        else:
            request.resource.ClearField("database_encryption")
        if ClusterVerticalPodAutoscaling.to_proto(self.vertical_pod_autoscaling):
            request.resource.vertical_pod_autoscaling.CopyFrom(
                ClusterVerticalPodAutoscaling.to_proto(self.vertical_pod_autoscaling)
            )
        else:
            request.resource.ClearField("vertical_pod_autoscaling")
        if ClusterShieldedNodes.to_proto(self.shielded_nodes):
            request.resource.shielded_nodes.CopyFrom(
                ClusterShieldedNodes.to_proto(self.shielded_nodes)
            )
        else:
            request.resource.ClearField("shielded_nodes")
        if Primitive.to_proto(self.master_version):
            request.resource.master_version = Primitive.to_proto(self.master_version)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.enable_tpu):
            request.resource.enable_tpu = Primitive.to_proto(self.enable_tpu)

        if ClusterAutopilot.to_proto(self.autopilot):
            request.resource.autopilot.CopyFrom(
                ClusterAutopilot.to_proto(self.autopilot)
            )
        else:
            request.resource.ClearField("autopilot")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyContainerCluster(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.initial_node_count = Primitive.from_proto(response.initial_node_count)
        self.master_auth = ClusterMasterAuth.from_proto(response.master_auth)
        self.logging_service = Primitive.from_proto(response.logging_service)
        self.monitoring_service = Primitive.from_proto(response.monitoring_service)
        self.network = Primitive.from_proto(response.network)
        self.cluster_ipv4_cidr = Primitive.from_proto(response.cluster_ipv4_cidr)
        self.addons_config = ClusterAddonsConfig.from_proto(response.addons_config)
        self.subnetwork = Primitive.from_proto(response.subnetwork)
        self.node_pools = ClusterNodePoolsArray.from_proto(response.node_pools)
        self.locations = Primitive.from_proto(response.locations)
        self.enable_kubernetes_alpha = Primitive.from_proto(
            response.enable_kubernetes_alpha
        )
        self.resource_labels = Primitive.from_proto(response.resource_labels)
        self.label_fingerprint = Primitive.from_proto(response.label_fingerprint)
        self.legacy_abac = ClusterLegacyAbac.from_proto(response.legacy_abac)
        self.network_policy = ClusterNetworkPolicy.from_proto(response.network_policy)
        self.ip_allocation_policy = ClusterIPAllocationPolicy.from_proto(
            response.ip_allocation_policy
        )
        self.master_authorized_networks_config = ClusterMasterAuthorizedNetworksConfig.from_proto(
            response.master_authorized_networks_config
        )
        self.binary_authorization = ClusterBinaryAuthorization.from_proto(
            response.binary_authorization
        )
        self.autoscaling = ClusterAutoscaling.from_proto(response.autoscaling)
        self.network_config = ClusterNetworkConfig.from_proto(response.network_config)
        self.maintenance_policy = ClusterMaintenancePolicy.from_proto(
            response.maintenance_policy
        )
        self.default_max_pods_constraint = ClusterDefaultMaxPodsConstraint.from_proto(
            response.default_max_pods_constraint
        )
        self.resource_usage_export_config = ClusterResourceUsageExportConfig.from_proto(
            response.resource_usage_export_config
        )
        self.authenticator_groups_config = ClusterAuthenticatorGroupsConfig.from_proto(
            response.authenticator_groups_config
        )
        self.private_cluster_config = ClusterPrivateClusterConfig.from_proto(
            response.private_cluster_config
        )
        self.database_encryption = ClusterDatabaseEncryption.from_proto(
            response.database_encryption
        )
        self.vertical_pod_autoscaling = ClusterVerticalPodAutoscaling.from_proto(
            response.vertical_pod_autoscaling
        )
        self.shielded_nodes = ClusterShieldedNodes.from_proto(response.shielded_nodes)
        self.endpoint = Primitive.from_proto(response.endpoint)
        self.master_version = Primitive.from_proto(response.master_version)
        self.create_time = Primitive.from_proto(response.create_time)
        self.status = Primitive.from_proto(response.status)
        self.status_message = Primitive.from_proto(response.status_message)
        self.node_ipv4_cidr_size = Primitive.from_proto(response.node_ipv4_cidr_size)
        self.services_ipv4_cidr = Primitive.from_proto(response.services_ipv4_cidr)
        self.expire_time = Primitive.from_proto(response.expire_time)
        self.location = Primitive.from_proto(response.location)
        self.enable_tpu = Primitive.from_proto(response.enable_tpu)
        self.tpu_ipv4_cidr_block = Primitive.from_proto(response.tpu_ipv4_cidr_block)
        self.conditions = ClusterConditionsArray.from_proto(response.conditions)
        self.autopilot = ClusterAutopilot.from_proto(response.autopilot)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = cluster_pb2_grpc.ContainerClusterServiceStub(channel.Channel())
        request = cluster_pb2.DeleteContainerClusterRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.initial_node_count):
            request.resource.initial_node_count = Primitive.to_proto(
                self.initial_node_count
            )

        if ClusterMasterAuth.to_proto(self.master_auth):
            request.resource.master_auth.CopyFrom(
                ClusterMasterAuth.to_proto(self.master_auth)
            )
        else:
            request.resource.ClearField("master_auth")
        if Primitive.to_proto(self.logging_service):
            request.resource.logging_service = Primitive.to_proto(self.logging_service)

        if Primitive.to_proto(self.monitoring_service):
            request.resource.monitoring_service = Primitive.to_proto(
                self.monitoring_service
            )

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.cluster_ipv4_cidr):
            request.resource.cluster_ipv4_cidr = Primitive.to_proto(
                self.cluster_ipv4_cidr
            )

        if ClusterAddonsConfig.to_proto(self.addons_config):
            request.resource.addons_config.CopyFrom(
                ClusterAddonsConfig.to_proto(self.addons_config)
            )
        else:
            request.resource.ClearField("addons_config")
        if Primitive.to_proto(self.subnetwork):
            request.resource.subnetwork = Primitive.to_proto(self.subnetwork)

        if Primitive.to_proto(self.locations):
            request.resource.locations.extend(Primitive.to_proto(self.locations))
        if Primitive.to_proto(self.enable_kubernetes_alpha):
            request.resource.enable_kubernetes_alpha = Primitive.to_proto(
                self.enable_kubernetes_alpha
            )

        if Primitive.to_proto(self.resource_labels):
            request.resource.resource_labels = Primitive.to_proto(self.resource_labels)

        if ClusterLegacyAbac.to_proto(self.legacy_abac):
            request.resource.legacy_abac.CopyFrom(
                ClusterLegacyAbac.to_proto(self.legacy_abac)
            )
        else:
            request.resource.ClearField("legacy_abac")
        if ClusterNetworkPolicy.to_proto(self.network_policy):
            request.resource.network_policy.CopyFrom(
                ClusterNetworkPolicy.to_proto(self.network_policy)
            )
        else:
            request.resource.ClearField("network_policy")
        if ClusterIPAllocationPolicy.to_proto(self.ip_allocation_policy):
            request.resource.ip_allocation_policy.CopyFrom(
                ClusterIPAllocationPolicy.to_proto(self.ip_allocation_policy)
            )
        else:
            request.resource.ClearField("ip_allocation_policy")
        if ClusterMasterAuthorizedNetworksConfig.to_proto(
            self.master_authorized_networks_config
        ):
            request.resource.master_authorized_networks_config.CopyFrom(
                ClusterMasterAuthorizedNetworksConfig.to_proto(
                    self.master_authorized_networks_config
                )
            )
        else:
            request.resource.ClearField("master_authorized_networks_config")
        if ClusterBinaryAuthorization.to_proto(self.binary_authorization):
            request.resource.binary_authorization.CopyFrom(
                ClusterBinaryAuthorization.to_proto(self.binary_authorization)
            )
        else:
            request.resource.ClearField("binary_authorization")
        if ClusterAutoscaling.to_proto(self.autoscaling):
            request.resource.autoscaling.CopyFrom(
                ClusterAutoscaling.to_proto(self.autoscaling)
            )
        else:
            request.resource.ClearField("autoscaling")
        if ClusterNetworkConfig.to_proto(self.network_config):
            request.resource.network_config.CopyFrom(
                ClusterNetworkConfig.to_proto(self.network_config)
            )
        else:
            request.resource.ClearField("network_config")
        if ClusterMaintenancePolicy.to_proto(self.maintenance_policy):
            request.resource.maintenance_policy.CopyFrom(
                ClusterMaintenancePolicy.to_proto(self.maintenance_policy)
            )
        else:
            request.resource.ClearField("maintenance_policy")
        if ClusterDefaultMaxPodsConstraint.to_proto(self.default_max_pods_constraint):
            request.resource.default_max_pods_constraint.CopyFrom(
                ClusterDefaultMaxPodsConstraint.to_proto(
                    self.default_max_pods_constraint
                )
            )
        else:
            request.resource.ClearField("default_max_pods_constraint")
        if ClusterResourceUsageExportConfig.to_proto(self.resource_usage_export_config):
            request.resource.resource_usage_export_config.CopyFrom(
                ClusterResourceUsageExportConfig.to_proto(
                    self.resource_usage_export_config
                )
            )
        else:
            request.resource.ClearField("resource_usage_export_config")
        if ClusterAuthenticatorGroupsConfig.to_proto(self.authenticator_groups_config):
            request.resource.authenticator_groups_config.CopyFrom(
                ClusterAuthenticatorGroupsConfig.to_proto(
                    self.authenticator_groups_config
                )
            )
        else:
            request.resource.ClearField("authenticator_groups_config")
        if ClusterPrivateClusterConfig.to_proto(self.private_cluster_config):
            request.resource.private_cluster_config.CopyFrom(
                ClusterPrivateClusterConfig.to_proto(self.private_cluster_config)
            )
        else:
            request.resource.ClearField("private_cluster_config")
        if ClusterDatabaseEncryption.to_proto(self.database_encryption):
            request.resource.database_encryption.CopyFrom(
                ClusterDatabaseEncryption.to_proto(self.database_encryption)
            )
        else:
            request.resource.ClearField("database_encryption")
        if ClusterVerticalPodAutoscaling.to_proto(self.vertical_pod_autoscaling):
            request.resource.vertical_pod_autoscaling.CopyFrom(
                ClusterVerticalPodAutoscaling.to_proto(self.vertical_pod_autoscaling)
            )
        else:
            request.resource.ClearField("vertical_pod_autoscaling")
        if ClusterShieldedNodes.to_proto(self.shielded_nodes):
            request.resource.shielded_nodes.CopyFrom(
                ClusterShieldedNodes.to_proto(self.shielded_nodes)
            )
        else:
            request.resource.ClearField("shielded_nodes")
        if Primitive.to_proto(self.master_version):
            request.resource.master_version = Primitive.to_proto(self.master_version)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.enable_tpu):
            request.resource.enable_tpu = Primitive.to_proto(self.enable_tpu)

        if ClusterAutopilot.to_proto(self.autopilot):
            request.resource.autopilot.CopyFrom(
                ClusterAutopilot.to_proto(self.autopilot)
            )
        else:
            request.resource.ClearField("autopilot")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteContainerCluster(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = cluster_pb2_grpc.ContainerClusterServiceStub(channel.Channel())
        request = cluster_pb2.ListContainerClusterRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListContainerCluster(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = cluster_pb2.ContainerCluster()
        any_proto.Unpack(res_proto)

        res = Cluster()
        res.name = Primitive.from_proto(res_proto.name)
        res.description = Primitive.from_proto(res_proto.description)
        res.initial_node_count = Primitive.from_proto(res_proto.initial_node_count)
        res.master_auth = ClusterMasterAuth.from_proto(res_proto.master_auth)
        res.logging_service = Primitive.from_proto(res_proto.logging_service)
        res.monitoring_service = Primitive.from_proto(res_proto.monitoring_service)
        res.network = Primitive.from_proto(res_proto.network)
        res.cluster_ipv4_cidr = Primitive.from_proto(res_proto.cluster_ipv4_cidr)
        res.addons_config = ClusterAddonsConfig.from_proto(res_proto.addons_config)
        res.subnetwork = Primitive.from_proto(res_proto.subnetwork)
        res.node_pools = ClusterNodePoolsArray.from_proto(res_proto.node_pools)
        res.locations = Primitive.from_proto(res_proto.locations)
        res.enable_kubernetes_alpha = Primitive.from_proto(
            res_proto.enable_kubernetes_alpha
        )
        res.resource_labels = Primitive.from_proto(res_proto.resource_labels)
        res.label_fingerprint = Primitive.from_proto(res_proto.label_fingerprint)
        res.legacy_abac = ClusterLegacyAbac.from_proto(res_proto.legacy_abac)
        res.network_policy = ClusterNetworkPolicy.from_proto(res_proto.network_policy)
        res.ip_allocation_policy = ClusterIPAllocationPolicy.from_proto(
            res_proto.ip_allocation_policy
        )
        res.master_authorized_networks_config = ClusterMasterAuthorizedNetworksConfig.from_proto(
            res_proto.master_authorized_networks_config
        )
        res.binary_authorization = ClusterBinaryAuthorization.from_proto(
            res_proto.binary_authorization
        )
        res.autoscaling = ClusterAutoscaling.from_proto(res_proto.autoscaling)
        res.network_config = ClusterNetworkConfig.from_proto(res_proto.network_config)
        res.maintenance_policy = ClusterMaintenancePolicy.from_proto(
            res_proto.maintenance_policy
        )
        res.default_max_pods_constraint = ClusterDefaultMaxPodsConstraint.from_proto(
            res_proto.default_max_pods_constraint
        )
        res.resource_usage_export_config = ClusterResourceUsageExportConfig.from_proto(
            res_proto.resource_usage_export_config
        )
        res.authenticator_groups_config = ClusterAuthenticatorGroupsConfig.from_proto(
            res_proto.authenticator_groups_config
        )
        res.private_cluster_config = ClusterPrivateClusterConfig.from_proto(
            res_proto.private_cluster_config
        )
        res.database_encryption = ClusterDatabaseEncryption.from_proto(
            res_proto.database_encryption
        )
        res.vertical_pod_autoscaling = ClusterVerticalPodAutoscaling.from_proto(
            res_proto.vertical_pod_autoscaling
        )
        res.shielded_nodes = ClusterShieldedNodes.from_proto(res_proto.shielded_nodes)
        res.endpoint = Primitive.from_proto(res_proto.endpoint)
        res.master_version = Primitive.from_proto(res_proto.master_version)
        res.create_time = Primitive.from_proto(res_proto.create_time)
        res.status = Primitive.from_proto(res_proto.status)
        res.status_message = Primitive.from_proto(res_proto.status_message)
        res.node_ipv4_cidr_size = Primitive.from_proto(res_proto.node_ipv4_cidr_size)
        res.services_ipv4_cidr = Primitive.from_proto(res_proto.services_ipv4_cidr)
        res.expire_time = Primitive.from_proto(res_proto.expire_time)
        res.location = Primitive.from_proto(res_proto.location)
        res.enable_tpu = Primitive.from_proto(res_proto.enable_tpu)
        res.tpu_ipv4_cidr_block = Primitive.from_proto(res_proto.tpu_ipv4_cidr_block)
        res.conditions = ClusterConditionsArray.from_proto(res_proto.conditions)
        res.autopilot = ClusterAutopilot.from_proto(res_proto.autopilot)
        res.project = Primitive.from_proto(res_proto.project)
        return res

    def to_proto(self):
        resource = cluster_pb2.ContainerCluster()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.initial_node_count):
            resource.initial_node_count = Primitive.to_proto(self.initial_node_count)
        if ClusterMasterAuth.to_proto(self.master_auth):
            resource.master_auth.CopyFrom(ClusterMasterAuth.to_proto(self.master_auth))
        else:
            resource.ClearField("master_auth")
        if Primitive.to_proto(self.logging_service):
            resource.logging_service = Primitive.to_proto(self.logging_service)
        if Primitive.to_proto(self.monitoring_service):
            resource.monitoring_service = Primitive.to_proto(self.monitoring_service)
        if Primitive.to_proto(self.network):
            resource.network = Primitive.to_proto(self.network)
        if Primitive.to_proto(self.cluster_ipv4_cidr):
            resource.cluster_ipv4_cidr = Primitive.to_proto(self.cluster_ipv4_cidr)
        if ClusterAddonsConfig.to_proto(self.addons_config):
            resource.addons_config.CopyFrom(
                ClusterAddonsConfig.to_proto(self.addons_config)
            )
        else:
            resource.ClearField("addons_config")
        if Primitive.to_proto(self.subnetwork):
            resource.subnetwork = Primitive.to_proto(self.subnetwork)
        if Primitive.to_proto(self.locations):
            resource.locations.extend(Primitive.to_proto(self.locations))
        if Primitive.to_proto(self.enable_kubernetes_alpha):
            resource.enable_kubernetes_alpha = Primitive.to_proto(
                self.enable_kubernetes_alpha
            )
        if Primitive.to_proto(self.resource_labels):
            resource.resource_labels = Primitive.to_proto(self.resource_labels)
        if ClusterLegacyAbac.to_proto(self.legacy_abac):
            resource.legacy_abac.CopyFrom(ClusterLegacyAbac.to_proto(self.legacy_abac))
        else:
            resource.ClearField("legacy_abac")
        if ClusterNetworkPolicy.to_proto(self.network_policy):
            resource.network_policy.CopyFrom(
                ClusterNetworkPolicy.to_proto(self.network_policy)
            )
        else:
            resource.ClearField("network_policy")
        if ClusterIPAllocationPolicy.to_proto(self.ip_allocation_policy):
            resource.ip_allocation_policy.CopyFrom(
                ClusterIPAllocationPolicy.to_proto(self.ip_allocation_policy)
            )
        else:
            resource.ClearField("ip_allocation_policy")
        if ClusterMasterAuthorizedNetworksConfig.to_proto(
            self.master_authorized_networks_config
        ):
            resource.master_authorized_networks_config.CopyFrom(
                ClusterMasterAuthorizedNetworksConfig.to_proto(
                    self.master_authorized_networks_config
                )
            )
        else:
            resource.ClearField("master_authorized_networks_config")
        if ClusterBinaryAuthorization.to_proto(self.binary_authorization):
            resource.binary_authorization.CopyFrom(
                ClusterBinaryAuthorization.to_proto(self.binary_authorization)
            )
        else:
            resource.ClearField("binary_authorization")
        if ClusterAutoscaling.to_proto(self.autoscaling):
            resource.autoscaling.CopyFrom(ClusterAutoscaling.to_proto(self.autoscaling))
        else:
            resource.ClearField("autoscaling")
        if ClusterNetworkConfig.to_proto(self.network_config):
            resource.network_config.CopyFrom(
                ClusterNetworkConfig.to_proto(self.network_config)
            )
        else:
            resource.ClearField("network_config")
        if ClusterMaintenancePolicy.to_proto(self.maintenance_policy):
            resource.maintenance_policy.CopyFrom(
                ClusterMaintenancePolicy.to_proto(self.maintenance_policy)
            )
        else:
            resource.ClearField("maintenance_policy")
        if ClusterDefaultMaxPodsConstraint.to_proto(self.default_max_pods_constraint):
            resource.default_max_pods_constraint.CopyFrom(
                ClusterDefaultMaxPodsConstraint.to_proto(
                    self.default_max_pods_constraint
                )
            )
        else:
            resource.ClearField("default_max_pods_constraint")
        if ClusterResourceUsageExportConfig.to_proto(self.resource_usage_export_config):
            resource.resource_usage_export_config.CopyFrom(
                ClusterResourceUsageExportConfig.to_proto(
                    self.resource_usage_export_config
                )
            )
        else:
            resource.ClearField("resource_usage_export_config")
        if ClusterAuthenticatorGroupsConfig.to_proto(self.authenticator_groups_config):
            resource.authenticator_groups_config.CopyFrom(
                ClusterAuthenticatorGroupsConfig.to_proto(
                    self.authenticator_groups_config
                )
            )
        else:
            resource.ClearField("authenticator_groups_config")
        if ClusterPrivateClusterConfig.to_proto(self.private_cluster_config):
            resource.private_cluster_config.CopyFrom(
                ClusterPrivateClusterConfig.to_proto(self.private_cluster_config)
            )
        else:
            resource.ClearField("private_cluster_config")
        if ClusterDatabaseEncryption.to_proto(self.database_encryption):
            resource.database_encryption.CopyFrom(
                ClusterDatabaseEncryption.to_proto(self.database_encryption)
            )
        else:
            resource.ClearField("database_encryption")
        if ClusterVerticalPodAutoscaling.to_proto(self.vertical_pod_autoscaling):
            resource.vertical_pod_autoscaling.CopyFrom(
                ClusterVerticalPodAutoscaling.to_proto(self.vertical_pod_autoscaling)
            )
        else:
            resource.ClearField("vertical_pod_autoscaling")
        if ClusterShieldedNodes.to_proto(self.shielded_nodes):
            resource.shielded_nodes.CopyFrom(
                ClusterShieldedNodes.to_proto(self.shielded_nodes)
            )
        else:
            resource.ClearField("shielded_nodes")
        if Primitive.to_proto(self.master_version):
            resource.master_version = Primitive.to_proto(self.master_version)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.enable_tpu):
            resource.enable_tpu = Primitive.to_proto(self.enable_tpu)
        if ClusterAutopilot.to_proto(self.autopilot):
            resource.autopilot.CopyFrom(ClusterAutopilot.to_proto(self.autopilot))
        else:
            resource.ClearField("autopilot")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class ClusterMasterAuth(object):
    def __init__(
        self,
        username: str = None,
        password: str = None,
        client_certificate_config: dict = None,
        cluster_ca_certificate: str = None,
        client_certificate: str = None,
        client_key: str = None,
    ):
        self.username = username
        self.password = password
        self.client_certificate_config = client_certificate_config
        self.cluster_ca_certificate = cluster_ca_certificate
        self.client_certificate = client_certificate
        self.client_key = client_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMasterAuth()
        if Primitive.to_proto(resource.username):
            res.username = Primitive.to_proto(resource.username)
        if Primitive.to_proto(resource.password):
            res.password = Primitive.to_proto(resource.password)
        if ClusterMasterAuthClientCertificateConfig.to_proto(
            resource.client_certificate_config
        ):
            res.client_certificate_config.CopyFrom(
                ClusterMasterAuthClientCertificateConfig.to_proto(
                    resource.client_certificate_config
                )
            )
        else:
            res.ClearField("client_certificate_config")
        if Primitive.to_proto(resource.cluster_ca_certificate):
            res.cluster_ca_certificate = Primitive.to_proto(
                resource.cluster_ca_certificate
            )
        if Primitive.to_proto(resource.client_certificate):
            res.client_certificate = Primitive.to_proto(resource.client_certificate)
        if Primitive.to_proto(resource.client_key):
            res.client_key = Primitive.to_proto(resource.client_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMasterAuth(
            username=resource.username,
            password=resource.password,
            client_certificate_config=resource.client_certificate_config,
            cluster_ca_certificate=resource.cluster_ca_certificate,
            client_certificate=resource.client_certificate,
            client_key=resource.client_key,
        )


class ClusterMasterAuthArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterMasterAuth.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterMasterAuth.from_proto(i) for i in resources]


class ClusterMasterAuthClientCertificateConfig(object):
    def __init__(self, issue_client_certificate: bool = None):
        self.issue_client_certificate = issue_client_certificate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMasterAuthClientCertificateConfig()
        if Primitive.to_proto(resource.issue_client_certificate):
            res.issue_client_certificate = Primitive.to_proto(
                resource.issue_client_certificate
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMasterAuthClientCertificateConfig(
            issue_client_certificate=resource.issue_client_certificate,
        )


class ClusterMasterAuthClientCertificateConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterMasterAuthClientCertificateConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterMasterAuthClientCertificateConfig.from_proto(i) for i in resources
        ]


class ClusterAddonsConfig(object):
    def __init__(
        self,
        http_load_balancing: dict = None,
        horizontal_pod_autoscaling: dict = None,
        kubernetes_dashboard: dict = None,
        network_policy_config: dict = None,
        cloud_run_config: dict = None,
    ):
        self.http_load_balancing = http_load_balancing
        self.horizontal_pod_autoscaling = horizontal_pod_autoscaling
        self.kubernetes_dashboard = kubernetes_dashboard
        self.network_policy_config = network_policy_config
        self.cloud_run_config = cloud_run_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfig()
        if ClusterAddonsConfigHttpLoadBalancing.to_proto(resource.http_load_balancing):
            res.http_load_balancing.CopyFrom(
                ClusterAddonsConfigHttpLoadBalancing.to_proto(
                    resource.http_load_balancing
                )
            )
        else:
            res.ClearField("http_load_balancing")
        if ClusterAddonsConfigHorizontalPodAutoscaling.to_proto(
            resource.horizontal_pod_autoscaling
        ):
            res.horizontal_pod_autoscaling.CopyFrom(
                ClusterAddonsConfigHorizontalPodAutoscaling.to_proto(
                    resource.horizontal_pod_autoscaling
                )
            )
        else:
            res.ClearField("horizontal_pod_autoscaling")
        if ClusterAddonsConfigKubernetesDashboard.to_proto(
            resource.kubernetes_dashboard
        ):
            res.kubernetes_dashboard.CopyFrom(
                ClusterAddonsConfigKubernetesDashboard.to_proto(
                    resource.kubernetes_dashboard
                )
            )
        else:
            res.ClearField("kubernetes_dashboard")
        if ClusterAddonsConfigNetworkPolicyConfig.to_proto(
            resource.network_policy_config
        ):
            res.network_policy_config.CopyFrom(
                ClusterAddonsConfigNetworkPolicyConfig.to_proto(
                    resource.network_policy_config
                )
            )
        else:
            res.ClearField("network_policy_config")
        if ClusterAddonsConfigCloudRunConfig.to_proto(resource.cloud_run_config):
            res.cloud_run_config.CopyFrom(
                ClusterAddonsConfigCloudRunConfig.to_proto(resource.cloud_run_config)
            )
        else:
            res.ClearField("cloud_run_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfig(
            http_load_balancing=resource.http_load_balancing,
            horizontal_pod_autoscaling=resource.horizontal_pod_autoscaling,
            kubernetes_dashboard=resource.kubernetes_dashboard,
            network_policy_config=resource.network_policy_config,
            cloud_run_config=resource.cloud_run_config,
        )


class ClusterAddonsConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAddonsConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAddonsConfig.from_proto(i) for i in resources]


class ClusterAddonsConfigHttpLoadBalancing(object):
    def __init__(self, disabled: bool = None):
        self.disabled = disabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfigHttpLoadBalancing()
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfigHttpLoadBalancing(disabled=resource.disabled,)


class ClusterAddonsConfigHttpLoadBalancingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAddonsConfigHttpLoadBalancing.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAddonsConfigHttpLoadBalancing.from_proto(i) for i in resources]


class ClusterAddonsConfigHorizontalPodAutoscaling(object):
    def __init__(self, disabled: bool = None):
        self.disabled = disabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfigHorizontalPodAutoscaling()
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfigHorizontalPodAutoscaling(disabled=resource.disabled,)


class ClusterAddonsConfigHorizontalPodAutoscalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterAddonsConfigHorizontalPodAutoscaling.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterAddonsConfigHorizontalPodAutoscaling.from_proto(i) for i in resources
        ]


class ClusterAddonsConfigKubernetesDashboard(object):
    def __init__(self, disabled: bool = None):
        self.disabled = disabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfigKubernetesDashboard()
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfigKubernetesDashboard(disabled=resource.disabled,)


class ClusterAddonsConfigKubernetesDashboardArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAddonsConfigKubernetesDashboard.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAddonsConfigKubernetesDashboard.from_proto(i) for i in resources]


class ClusterAddonsConfigNetworkPolicyConfig(object):
    def __init__(self, disabled: bool = None):
        self.disabled = disabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfigNetworkPolicyConfig()
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfigNetworkPolicyConfig(disabled=resource.disabled,)


class ClusterAddonsConfigNetworkPolicyConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAddonsConfigNetworkPolicyConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAddonsConfigNetworkPolicyConfig.from_proto(i) for i in resources]


class ClusterAddonsConfigCloudRunConfig(object):
    def __init__(self, disabled: bool = None):
        self.disabled = disabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAddonsConfigCloudRunConfig()
        if Primitive.to_proto(resource.disabled):
            res.disabled = Primitive.to_proto(resource.disabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAddonsConfigCloudRunConfig(disabled=resource.disabled,)


class ClusterAddonsConfigCloudRunConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAddonsConfigCloudRunConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAddonsConfigCloudRunConfig.from_proto(i) for i in resources]


class ClusterNodePools(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNodePools()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNodePools(name=resource.name,)


class ClusterNodePoolsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNodePools.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNodePools.from_proto(i) for i in resources]


class ClusterLegacyAbac(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterLegacyAbac()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterLegacyAbac(enabled=resource.enabled,)


class ClusterLegacyAbacArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterLegacyAbac.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterLegacyAbac.from_proto(i) for i in resources]


class ClusterNetworkPolicy(object):
    def __init__(self, provider: str = None, enabled: bool = None):
        self.provider = provider
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNetworkPolicy()
        if ClusterNetworkPolicyProviderEnum.to_proto(resource.provider):
            res.provider = ClusterNetworkPolicyProviderEnum.to_proto(resource.provider)
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNetworkPolicy(
            provider=resource.provider, enabled=resource.enabled,
        )


class ClusterNetworkPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNetworkPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNetworkPolicy.from_proto(i) for i in resources]


class ClusterIPAllocationPolicy(object):
    def __init__(
        self,
        use_ip_aliases: bool = None,
        create_subnetwork: bool = None,
        subnetwork_name: str = None,
        cluster_secondary_range_name: str = None,
        services_secondary_range_name: str = None,
        cluster_ipv4_cidr_block: str = None,
        node_ipv4_cidr_block: str = None,
        services_ipv4_cidr_block: str = None,
        tpu_ipv4_cidr_block: str = None,
    ):
        self.use_ip_aliases = use_ip_aliases
        self.create_subnetwork = create_subnetwork
        self.subnetwork_name = subnetwork_name
        self.cluster_secondary_range_name = cluster_secondary_range_name
        self.services_secondary_range_name = services_secondary_range_name
        self.cluster_ipv4_cidr_block = cluster_ipv4_cidr_block
        self.node_ipv4_cidr_block = node_ipv4_cidr_block
        self.services_ipv4_cidr_block = services_ipv4_cidr_block
        self.tpu_ipv4_cidr_block = tpu_ipv4_cidr_block

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterIPAllocationPolicy()
        if Primitive.to_proto(resource.use_ip_aliases):
            res.use_ip_aliases = Primitive.to_proto(resource.use_ip_aliases)
        if Primitive.to_proto(resource.create_subnetwork):
            res.create_subnetwork = Primitive.to_proto(resource.create_subnetwork)
        if Primitive.to_proto(resource.subnetwork_name):
            res.subnetwork_name = Primitive.to_proto(resource.subnetwork_name)
        if Primitive.to_proto(resource.cluster_secondary_range_name):
            res.cluster_secondary_range_name = Primitive.to_proto(
                resource.cluster_secondary_range_name
            )
        if Primitive.to_proto(resource.services_secondary_range_name):
            res.services_secondary_range_name = Primitive.to_proto(
                resource.services_secondary_range_name
            )
        if Primitive.to_proto(resource.cluster_ipv4_cidr_block):
            res.cluster_ipv4_cidr_block = Primitive.to_proto(
                resource.cluster_ipv4_cidr_block
            )
        if Primitive.to_proto(resource.node_ipv4_cidr_block):
            res.node_ipv4_cidr_block = Primitive.to_proto(resource.node_ipv4_cidr_block)
        if Primitive.to_proto(resource.services_ipv4_cidr_block):
            res.services_ipv4_cidr_block = Primitive.to_proto(
                resource.services_ipv4_cidr_block
            )
        if Primitive.to_proto(resource.tpu_ipv4_cidr_block):
            res.tpu_ipv4_cidr_block = Primitive.to_proto(resource.tpu_ipv4_cidr_block)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterIPAllocationPolicy(
            use_ip_aliases=resource.use_ip_aliases,
            create_subnetwork=resource.create_subnetwork,
            subnetwork_name=resource.subnetwork_name,
            cluster_secondary_range_name=resource.cluster_secondary_range_name,
            services_secondary_range_name=resource.services_secondary_range_name,
            cluster_ipv4_cidr_block=resource.cluster_ipv4_cidr_block,
            node_ipv4_cidr_block=resource.node_ipv4_cidr_block,
            services_ipv4_cidr_block=resource.services_ipv4_cidr_block,
            tpu_ipv4_cidr_block=resource.tpu_ipv4_cidr_block,
        )


class ClusterIPAllocationPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterIPAllocationPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterIPAllocationPolicy.from_proto(i) for i in resources]


class ClusterMasterAuthorizedNetworksConfig(object):
    def __init__(self, enabled: bool = None, cidr_blocks: list = None):
        self.enabled = enabled
        self.cidr_blocks = cidr_blocks

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMasterAuthorizedNetworksConfig()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if ClusterMasterAuthorizedNetworksConfigCidrBlocksArray.to_proto(
            resource.cidr_blocks
        ):
            res.cidr_blocks.extend(
                ClusterMasterAuthorizedNetworksConfigCidrBlocksArray.to_proto(
                    resource.cidr_blocks
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMasterAuthorizedNetworksConfig(
            enabled=resource.enabled, cidr_blocks=resource.cidr_blocks,
        )


class ClusterMasterAuthorizedNetworksConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterMasterAuthorizedNetworksConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterMasterAuthorizedNetworksConfig.from_proto(i) for i in resources]


class ClusterMasterAuthorizedNetworksConfigCidrBlocks(object):
    def __init__(self, display_name: str = None, cidr_block: str = None):
        self.display_name = display_name
        self.cidr_block = cidr_block

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMasterAuthorizedNetworksConfigCidrBlocks()
        if Primitive.to_proto(resource.display_name):
            res.display_name = Primitive.to_proto(resource.display_name)
        if Primitive.to_proto(resource.cidr_block):
            res.cidr_block = Primitive.to_proto(resource.cidr_block)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMasterAuthorizedNetworksConfigCidrBlocks(
            display_name=resource.display_name, cidr_block=resource.cidr_block,
        )


class ClusterMasterAuthorizedNetworksConfigCidrBlocksArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterMasterAuthorizedNetworksConfigCidrBlocks.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterMasterAuthorizedNetworksConfigCidrBlocks.from_proto(i)
            for i in resources
        ]


class ClusterBinaryAuthorization(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterBinaryAuthorization()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterBinaryAuthorization(enabled=resource.enabled,)


class ClusterBinaryAuthorizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterBinaryAuthorization.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterBinaryAuthorization.from_proto(i) for i in resources]


class ClusterAutoscaling(object):
    def __init__(
        self,
        enable_node_autoprovisioning: bool = None,
        resource_limits: list = None,
        autoprovisioning_node_pool_defaults: dict = None,
    ):
        self.enable_node_autoprovisioning = enable_node_autoprovisioning
        self.resource_limits = resource_limits
        self.autoprovisioning_node_pool_defaults = autoprovisioning_node_pool_defaults

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAutoscaling()
        if Primitive.to_proto(resource.enable_node_autoprovisioning):
            res.enable_node_autoprovisioning = Primitive.to_proto(
                resource.enable_node_autoprovisioning
            )
        if ClusterAutoscalingResourceLimitsArray.to_proto(resource.resource_limits):
            res.resource_limits.extend(
                ClusterAutoscalingResourceLimitsArray.to_proto(resource.resource_limits)
            )
        if ClusterAutoscalingAutoprovisioningNodePoolDefaults.to_proto(
            resource.autoprovisioning_node_pool_defaults
        ):
            res.autoprovisioning_node_pool_defaults.CopyFrom(
                ClusterAutoscalingAutoprovisioningNodePoolDefaults.to_proto(
                    resource.autoprovisioning_node_pool_defaults
                )
            )
        else:
            res.ClearField("autoprovisioning_node_pool_defaults")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAutoscaling(
            enable_node_autoprovisioning=resource.enable_node_autoprovisioning,
            resource_limits=resource.resource_limits,
            autoprovisioning_node_pool_defaults=resource.autoprovisioning_node_pool_defaults,
        )


class ClusterAutoscalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAutoscaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAutoscaling.from_proto(i) for i in resources]


class ClusterAutoscalingResourceLimits(object):
    def __init__(
        self, resource_type: str = None, minimum: int = None, maximum: int = None
    ):
        self.resource_type = resource_type
        self.minimum = minimum
        self.maximum = maximum

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAutoscalingResourceLimits()
        if Primitive.to_proto(resource.resource_type):
            res.resource_type = Primitive.to_proto(resource.resource_type)
        if Primitive.to_proto(resource.minimum):
            res.minimum = Primitive.to_proto(resource.minimum)
        if Primitive.to_proto(resource.maximum):
            res.maximum = Primitive.to_proto(resource.maximum)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAutoscalingResourceLimits(
            resource_type=resource.resource_type,
            minimum=resource.minimum,
            maximum=resource.maximum,
        )


class ClusterAutoscalingResourceLimitsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAutoscalingResourceLimits.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAutoscalingResourceLimits.from_proto(i) for i in resources]


class ClusterAutoscalingAutoprovisioningNodePoolDefaults(object):
    def __init__(
        self,
        oauth_scopes: list = None,
        service_account: str = None,
        upgrade_settings: dict = None,
        management: dict = None,
    ):
        self.oauth_scopes = oauth_scopes
        self.service_account = service_account
        self.upgrade_settings = upgrade_settings
        self.management = management

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaults()
        if Primitive.to_proto(resource.oauth_scopes):
            res.oauth_scopes.extend(Primitive.to_proto(resource.oauth_scopes))
        if Primitive.to_proto(resource.service_account):
            res.service_account = Primitive.to_proto(resource.service_account)
        if ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings.to_proto(
            resource.upgrade_settings
        ):
            res.upgrade_settings.CopyFrom(
                ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings.to_proto(
                    resource.upgrade_settings
                )
            )
        else:
            res.ClearField("upgrade_settings")
        if ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement.to_proto(
            resource.management
        ):
            res.management.CopyFrom(
                ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement.to_proto(
                    resource.management
                )
            )
        else:
            res.ClearField("management")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAutoscalingAutoprovisioningNodePoolDefaults(
            oauth_scopes=resource.oauth_scopes,
            service_account=resource.service_account,
            upgrade_settings=resource.upgrade_settings,
            management=resource.management,
        )


class ClusterAutoscalingAutoprovisioningNodePoolDefaultsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaults.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaults.from_proto(i)
            for i in resources
        ]


class ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(object):
    def __init__(self, max_surge: int = None, max_unavailable: int = None):
        self.max_surge = max_surge
        self.max_unavailable = max_unavailable

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings()
        )
        if Primitive.to_proto(resource.max_surge):
            res.max_surge = Primitive.to_proto(resource.max_surge)
        if Primitive.to_proto(resource.max_unavailable):
            res.max_unavailable = Primitive.to_proto(resource.max_unavailable)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings(
            max_surge=resource.max_surge, max_unavailable=resource.max_unavailable,
        )


class ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaultsUpgradeSettings.from_proto(
                i
            )
            for i in resources
        ]


class ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(object):
    def __init__(self, auto_upgrade: bool = None, auto_repair: bool = None):
        self.auto_upgrade = auto_upgrade
        self.auto_repair = auto_repair

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.ContainerClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement()
        )
        if Primitive.to_proto(resource.auto_upgrade):
            res.auto_upgrade = Primitive.to_proto(resource.auto_upgrade)
        if Primitive.to_proto(resource.auto_repair):
            res.auto_repair = Primitive.to_proto(resource.auto_repair)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement(
            auto_upgrade=resource.auto_upgrade, auto_repair=resource.auto_repair,
        )


class ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagementArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterAutoscalingAutoprovisioningNodePoolDefaultsManagement.from_proto(i)
            for i in resources
        ]


class ClusterNetworkConfig(object):
    def __init__(
        self,
        network: str = None,
        subnetwork: str = None,
        enable_intra_node_visibility: bool = None,
    ):
        self.network = network
        self.subnetwork = subnetwork
        self.enable_intra_node_visibility = enable_intra_node_visibility

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterNetworkConfig()
        if Primitive.to_proto(resource.network):
            res.network = Primitive.to_proto(resource.network)
        if Primitive.to_proto(resource.subnetwork):
            res.subnetwork = Primitive.to_proto(resource.subnetwork)
        if Primitive.to_proto(resource.enable_intra_node_visibility):
            res.enable_intra_node_visibility = Primitive.to_proto(
                resource.enable_intra_node_visibility
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNetworkConfig(
            network=resource.network,
            subnetwork=resource.subnetwork,
            enable_intra_node_visibility=resource.enable_intra_node_visibility,
        )


class ClusterNetworkConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNetworkConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNetworkConfig.from_proto(i) for i in resources]


class ClusterMaintenancePolicy(object):
    def __init__(self, window: dict = None, resource_version: str = None):
        self.window = window
        self.resource_version = resource_version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMaintenancePolicy()
        if ClusterMaintenancePolicyWindow.to_proto(resource.window):
            res.window.CopyFrom(
                ClusterMaintenancePolicyWindow.to_proto(resource.window)
            )
        else:
            res.ClearField("window")
        if Primitive.to_proto(resource.resource_version):
            res.resource_version = Primitive.to_proto(resource.resource_version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMaintenancePolicy(
            window=resource.window, resource_version=resource.resource_version,
        )


class ClusterMaintenancePolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterMaintenancePolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterMaintenancePolicy.from_proto(i) for i in resources]


class ClusterMaintenancePolicyWindow(object):
    def __init__(
        self, daily_maintenance_window: dict = None, recurring_window: dict = None
    ):
        self.daily_maintenance_window = daily_maintenance_window
        self.recurring_window = recurring_window

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMaintenancePolicyWindow()
        if ClusterMaintenancePolicyWindowDailyMaintenanceWindow.to_proto(
            resource.daily_maintenance_window
        ):
            res.daily_maintenance_window.CopyFrom(
                ClusterMaintenancePolicyWindowDailyMaintenanceWindow.to_proto(
                    resource.daily_maintenance_window
                )
            )
        else:
            res.ClearField("daily_maintenance_window")
        if ClusterMaintenancePolicyWindowRecurringWindow.to_proto(
            resource.recurring_window
        ):
            res.recurring_window.CopyFrom(
                ClusterMaintenancePolicyWindowRecurringWindow.to_proto(
                    resource.recurring_window
                )
            )
        else:
            res.ClearField("recurring_window")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMaintenancePolicyWindow(
            daily_maintenance_window=resource.daily_maintenance_window,
            recurring_window=resource.recurring_window,
        )


class ClusterMaintenancePolicyWindowArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterMaintenancePolicyWindow.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterMaintenancePolicyWindow.from_proto(i) for i in resources]


class ClusterMaintenancePolicyWindowDailyMaintenanceWindow(object):
    def __init__(self, start_time: str = None, duration: str = None):
        self.start_time = start_time
        self.duration = duration

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.ContainerClusterMaintenancePolicyWindowDailyMaintenanceWindow()
        )
        if Primitive.to_proto(resource.start_time):
            res.start_time = Primitive.to_proto(resource.start_time)
        if Primitive.to_proto(resource.duration):
            res.duration = Primitive.to_proto(resource.duration)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMaintenancePolicyWindowDailyMaintenanceWindow(
            start_time=resource.start_time, duration=resource.duration,
        )


class ClusterMaintenancePolicyWindowDailyMaintenanceWindowArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterMaintenancePolicyWindowDailyMaintenanceWindow.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterMaintenancePolicyWindowDailyMaintenanceWindow.from_proto(i)
            for i in resources
        ]


class ClusterMaintenancePolicyWindowRecurringWindow(object):
    def __init__(self, window: dict = None, recurrence: str = None):
        self.window = window
        self.recurrence = recurrence

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMaintenancePolicyWindowRecurringWindow()
        if ClusterMaintenancePolicyWindowRecurringWindowWindow.to_proto(
            resource.window
        ):
            res.window.CopyFrom(
                ClusterMaintenancePolicyWindowRecurringWindowWindow.to_proto(
                    resource.window
                )
            )
        else:
            res.ClearField("window")
        if Primitive.to_proto(resource.recurrence):
            res.recurrence = Primitive.to_proto(resource.recurrence)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMaintenancePolicyWindowRecurringWindow(
            window=resource.window, recurrence=resource.recurrence,
        )


class ClusterMaintenancePolicyWindowRecurringWindowArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterMaintenancePolicyWindowRecurringWindow.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterMaintenancePolicyWindowRecurringWindow.from_proto(i)
            for i in resources
        ]


class ClusterMaintenancePolicyWindowRecurringWindowWindow(object):
    def __init__(self, start_time: str = None, end_time: str = None):
        self.start_time = start_time
        self.end_time = end_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterMaintenancePolicyWindowRecurringWindowWindow()
        if Primitive.to_proto(resource.start_time):
            res.start_time = Primitive.to_proto(resource.start_time)
        if Primitive.to_proto(resource.end_time):
            res.end_time = Primitive.to_proto(resource.end_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMaintenancePolicyWindowRecurringWindowWindow(
            start_time=resource.start_time, end_time=resource.end_time,
        )


class ClusterMaintenancePolicyWindowRecurringWindowWindowArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterMaintenancePolicyWindowRecurringWindowWindow.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterMaintenancePolicyWindowRecurringWindowWindow.from_proto(i)
            for i in resources
        ]


class ClusterDefaultMaxPodsConstraint(object):
    def __init__(self, max_pods_per_node: str = None):
        self.max_pods_per_node = max_pods_per_node

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterDefaultMaxPodsConstraint()
        if Primitive.to_proto(resource.max_pods_per_node):
            res.max_pods_per_node = Primitive.to_proto(resource.max_pods_per_node)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterDefaultMaxPodsConstraint(
            max_pods_per_node=resource.max_pods_per_node,
        )


class ClusterDefaultMaxPodsConstraintArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterDefaultMaxPodsConstraint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterDefaultMaxPodsConstraint.from_proto(i) for i in resources]


class ClusterResourceUsageExportConfig(object):
    def __init__(
        self,
        bigquery_destination: dict = None,
        enable_network_egress_monitoring: bool = None,
        consumption_metering_config: dict = None,
    ):
        self.bigquery_destination = bigquery_destination
        self.enable_network_egress_monitoring = enable_network_egress_monitoring
        self.consumption_metering_config = consumption_metering_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterResourceUsageExportConfig()
        if ClusterResourceUsageExportConfigBigqueryDestination.to_proto(
            resource.bigquery_destination
        ):
            res.bigquery_destination.CopyFrom(
                ClusterResourceUsageExportConfigBigqueryDestination.to_proto(
                    resource.bigquery_destination
                )
            )
        else:
            res.ClearField("bigquery_destination")
        if Primitive.to_proto(resource.enable_network_egress_monitoring):
            res.enable_network_egress_monitoring = Primitive.to_proto(
                resource.enable_network_egress_monitoring
            )
        if ClusterResourceUsageExportConfigConsumptionMeteringConfig.to_proto(
            resource.consumption_metering_config
        ):
            res.consumption_metering_config.CopyFrom(
                ClusterResourceUsageExportConfigConsumptionMeteringConfig.to_proto(
                    resource.consumption_metering_config
                )
            )
        else:
            res.ClearField("consumption_metering_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterResourceUsageExportConfig(
            bigquery_destination=resource.bigquery_destination,
            enable_network_egress_monitoring=resource.enable_network_egress_monitoring,
            consumption_metering_config=resource.consumption_metering_config,
        )


class ClusterResourceUsageExportConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterResourceUsageExportConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterResourceUsageExportConfig.from_proto(i) for i in resources]


class ClusterResourceUsageExportConfigBigqueryDestination(object):
    def __init__(self, dataset_id: str = None):
        self.dataset_id = dataset_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterResourceUsageExportConfigBigqueryDestination()
        if Primitive.to_proto(resource.dataset_id):
            res.dataset_id = Primitive.to_proto(resource.dataset_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterResourceUsageExportConfigBigqueryDestination(
            dataset_id=resource.dataset_id,
        )


class ClusterResourceUsageExportConfigBigqueryDestinationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterResourceUsageExportConfigBigqueryDestination.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterResourceUsageExportConfigBigqueryDestination.from_proto(i)
            for i in resources
        ]


class ClusterResourceUsageExportConfigConsumptionMeteringConfig(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.ContainerClusterResourceUsageExportConfigConsumptionMeteringConfig()
        )
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterResourceUsageExportConfigConsumptionMeteringConfig(
            enabled=resource.enabled,
        )


class ClusterResourceUsageExportConfigConsumptionMeteringConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterResourceUsageExportConfigConsumptionMeteringConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterResourceUsageExportConfigConsumptionMeteringConfig.from_proto(i)
            for i in resources
        ]


class ClusterAuthenticatorGroupsConfig(object):
    def __init__(self, enabled: bool = None, security_group: str = None):
        self.enabled = enabled
        self.security_group = security_group

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAuthenticatorGroupsConfig()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if Primitive.to_proto(resource.security_group):
            res.security_group = Primitive.to_proto(resource.security_group)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAuthenticatorGroupsConfig(
            enabled=resource.enabled, security_group=resource.security_group,
        )


class ClusterAuthenticatorGroupsConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAuthenticatorGroupsConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAuthenticatorGroupsConfig.from_proto(i) for i in resources]


class ClusterPrivateClusterConfig(object):
    def __init__(
        self,
        enable_private_nodes: bool = None,
        enable_private_endpoint: bool = None,
        master_ipv4_cidr_block: str = None,
        private_endpoint: str = None,
        public_endpoint: str = None,
        peering_name: str = None,
    ):
        self.enable_private_nodes = enable_private_nodes
        self.enable_private_endpoint = enable_private_endpoint
        self.master_ipv4_cidr_block = master_ipv4_cidr_block
        self.private_endpoint = private_endpoint
        self.public_endpoint = public_endpoint
        self.peering_name = peering_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterPrivateClusterConfig()
        if Primitive.to_proto(resource.enable_private_nodes):
            res.enable_private_nodes = Primitive.to_proto(resource.enable_private_nodes)
        if Primitive.to_proto(resource.enable_private_endpoint):
            res.enable_private_endpoint = Primitive.to_proto(
                resource.enable_private_endpoint
            )
        if Primitive.to_proto(resource.master_ipv4_cidr_block):
            res.master_ipv4_cidr_block = Primitive.to_proto(
                resource.master_ipv4_cidr_block
            )
        if Primitive.to_proto(resource.private_endpoint):
            res.private_endpoint = Primitive.to_proto(resource.private_endpoint)
        if Primitive.to_proto(resource.public_endpoint):
            res.public_endpoint = Primitive.to_proto(resource.public_endpoint)
        if Primitive.to_proto(resource.peering_name):
            res.peering_name = Primitive.to_proto(resource.peering_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterPrivateClusterConfig(
            enable_private_nodes=resource.enable_private_nodes,
            enable_private_endpoint=resource.enable_private_endpoint,
            master_ipv4_cidr_block=resource.master_ipv4_cidr_block,
            private_endpoint=resource.private_endpoint,
            public_endpoint=resource.public_endpoint,
            peering_name=resource.peering_name,
        )


class ClusterPrivateClusterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterPrivateClusterConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterPrivateClusterConfig.from_proto(i) for i in resources]


class ClusterDatabaseEncryption(object):
    def __init__(self, state: str = None, key_name: str = None):
        self.state = state
        self.key_name = key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterDatabaseEncryption()
        if ClusterDatabaseEncryptionStateEnum.to_proto(resource.state):
            res.state = ClusterDatabaseEncryptionStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.key_name):
            res.key_name = Primitive.to_proto(resource.key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterDatabaseEncryption(
            state=resource.state, key_name=resource.key_name,
        )


class ClusterDatabaseEncryptionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterDatabaseEncryption.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterDatabaseEncryption.from_proto(i) for i in resources]


class ClusterVerticalPodAutoscaling(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterVerticalPodAutoscaling()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterVerticalPodAutoscaling(enabled=resource.enabled,)


class ClusterVerticalPodAutoscalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterVerticalPodAutoscaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterVerticalPodAutoscaling.from_proto(i) for i in resources]


class ClusterShieldedNodes(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterShieldedNodes()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterShieldedNodes(enabled=resource.enabled,)


class ClusterShieldedNodesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterShieldedNodes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterShieldedNodes.from_proto(i) for i in resources]


class ClusterConditions(object):
    def __init__(self, code: str = None, message: str = None):
        self.code = code
        self.message = message

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterConditions()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterConditions(code=resource.code, message=resource.message,)


class ClusterConditionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterConditions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterConditions.from_proto(i) for i in resources]


class ClusterAutopilot(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerClusterAutopilot()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAutopilot(enabled=resource.enabled,)


class ClusterAutopilotArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAutopilot.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAutopilot.from_proto(i) for i in resources]


class ClusterNetworkPolicyProviderEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNetworkPolicyProviderEnum.Value(
            "ContainerClusterNetworkPolicyProviderEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterNetworkPolicyProviderEnum.Name(resource)[
            len("ContainerClusterNetworkPolicyProviderEnum") :
        ]


class ClusterDatabaseEncryptionStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterDatabaseEncryptionStateEnum.Value(
            "ContainerClusterDatabaseEncryptionStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerClusterDatabaseEncryptionStateEnum.Name(resource)[
            len("ContainerClusterDatabaseEncryptionStateEnum") :
        ]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
