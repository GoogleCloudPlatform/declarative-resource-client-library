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
from google3.cloud.graphite.mmv2.services.google.dataproc import cluster_pb2
from google3.cloud.graphite.mmv2.services.google.dataproc import cluster_pb2_grpc

from typing import List


class Cluster(object):
    def __init__(
        self,
        project: str = None,
        name: str = None,
        config: dict = None,
        labels: dict = None,
        status: dict = None,
        status_history: list = None,
        cluster_uuid: str = None,
        metrics: dict = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.project = project
        self.name = name
        self.config = config
        self.labels = labels
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = cluster_pb2_grpc.DataprocClusterServiceStub(channel.Channel())
        request = cluster_pb2.ApplyDataprocClusterRequest()
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ClusterClusterConfig.to_proto(self.config):
            request.resource.config.CopyFrom(ClusterClusterConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDataprocCluster(request)
        self.project = Primitive.from_proto(response.project)
        self.name = Primitive.from_proto(response.name)
        self.config = ClusterClusterConfig.from_proto(response.config)
        self.labels = Primitive.from_proto(response.labels)
        self.status = ClusterStatus.from_proto(response.status)
        self.status_history = ClusterStatusHistoryArray.from_proto(
            response.status_history
        )
        self.cluster_uuid = Primitive.from_proto(response.cluster_uuid)
        self.metrics = ClusterMetrics.from_proto(response.metrics)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = cluster_pb2_grpc.DataprocClusterServiceStub(channel.Channel())
        request = cluster_pb2.DeleteDataprocClusterRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if ClusterClusterConfig.to_proto(self.config):
            request.resource.config.CopyFrom(ClusterClusterConfig.to_proto(self.config))
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteDataprocCluster(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = cluster_pb2_grpc.DataprocClusterServiceStub(channel.Channel())
        request = cluster_pb2.ListDataprocClusterRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListDataprocCluster(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = cluster_pb2.DataprocCluster()
        any_proto.Unpack(res_proto)

        res = Cluster()
        res.project = Primitive.from_proto(res_proto.project)
        res.name = Primitive.from_proto(res_proto.name)
        res.config = ClusterClusterConfig.from_proto(res_proto.config)
        res.labels = Primitive.from_proto(res_proto.labels)
        res.status = ClusterStatus.from_proto(res_proto.status)
        res.status_history = ClusterStatusHistoryArray.from_proto(
            res_proto.status_history
        )
        res.cluster_uuid = Primitive.from_proto(res_proto.cluster_uuid)
        res.metrics = ClusterMetrics.from_proto(res_proto.metrics)
        res.location = Primitive.from_proto(res_proto.location)
        return res

    def to_proto(self):
        resource = cluster_pb2.DataprocCluster()
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if ClusterClusterConfig.to_proto(self.config):
            resource.config.CopyFrom(ClusterClusterConfig.to_proto(self.config))
        else:
            resource.ClearField("config")
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class ClusterClusterConfig(object):
    def __init__(
        self,
        staging_bucket: str = None,
        temp_bucket: str = None,
        gce_cluster_config: dict = None,
        master_config: dict = None,
        worker_config: dict = None,
        secondary_worker_config: dict = None,
        software_config: dict = None,
        initialization_actions: list = None,
        encryption_config: dict = None,
        autoscaling_config: dict = None,
        security_config: dict = None,
        lifecycle_config: dict = None,
        endpoint_config: dict = None,
    ):
        self.staging_bucket = staging_bucket
        self.temp_bucket = temp_bucket
        self.gce_cluster_config = gce_cluster_config
        self.master_config = master_config
        self.worker_config = worker_config
        self.secondary_worker_config = secondary_worker_config
        self.software_config = software_config
        self.initialization_actions = initialization_actions
        self.encryption_config = encryption_config
        self.autoscaling_config = autoscaling_config
        self.security_config = security_config
        self.lifecycle_config = lifecycle_config
        self.endpoint_config = endpoint_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterClusterConfig()
        if Primitive.to_proto(resource.staging_bucket):
            res.staging_bucket = Primitive.to_proto(resource.staging_bucket)
        if Primitive.to_proto(resource.temp_bucket):
            res.temp_bucket = Primitive.to_proto(resource.temp_bucket)
        if ClusterClusterConfigGceClusterConfig.to_proto(resource.gce_cluster_config):
            res.gce_cluster_config.CopyFrom(
                ClusterClusterConfigGceClusterConfig.to_proto(
                    resource.gce_cluster_config
                )
            )
        else:
            res.ClearField("gce_cluster_config")
        if ClusterInstanceGroupConfig.to_proto(resource.master_config):
            res.master_config.CopyFrom(
                ClusterInstanceGroupConfig.to_proto(resource.master_config)
            )
        else:
            res.ClearField("master_config")
        if ClusterInstanceGroupConfig.to_proto(resource.worker_config):
            res.worker_config.CopyFrom(
                ClusterInstanceGroupConfig.to_proto(resource.worker_config)
            )
        else:
            res.ClearField("worker_config")
        if ClusterInstanceGroupConfig.to_proto(resource.secondary_worker_config):
            res.secondary_worker_config.CopyFrom(
                ClusterInstanceGroupConfig.to_proto(resource.secondary_worker_config)
            )
        else:
            res.ClearField("secondary_worker_config")
        if ClusterClusterConfigSoftwareConfig.to_proto(resource.software_config):
            res.software_config.CopyFrom(
                ClusterClusterConfigSoftwareConfig.to_proto(resource.software_config)
            )
        else:
            res.ClearField("software_config")
        if ClusterClusterConfigInitializationActionsArray.to_proto(
            resource.initialization_actions
        ):
            res.initialization_actions.extend(
                ClusterClusterConfigInitializationActionsArray.to_proto(
                    resource.initialization_actions
                )
            )
        if ClusterClusterConfigEncryptionConfig.to_proto(resource.encryption_config):
            res.encryption_config.CopyFrom(
                ClusterClusterConfigEncryptionConfig.to_proto(
                    resource.encryption_config
                )
            )
        else:
            res.ClearField("encryption_config")
        if ClusterClusterConfigAutoscalingConfig.to_proto(resource.autoscaling_config):
            res.autoscaling_config.CopyFrom(
                ClusterClusterConfigAutoscalingConfig.to_proto(
                    resource.autoscaling_config
                )
            )
        else:
            res.ClearField("autoscaling_config")
        if ClusterClusterConfigSecurityConfig.to_proto(resource.security_config):
            res.security_config.CopyFrom(
                ClusterClusterConfigSecurityConfig.to_proto(resource.security_config)
            )
        else:
            res.ClearField("security_config")
        if ClusterClusterConfigLifecycleConfig.to_proto(resource.lifecycle_config):
            res.lifecycle_config.CopyFrom(
                ClusterClusterConfigLifecycleConfig.to_proto(resource.lifecycle_config)
            )
        else:
            res.ClearField("lifecycle_config")
        if ClusterClusterConfigEndpointConfig.to_proto(resource.endpoint_config):
            res.endpoint_config.CopyFrom(
                ClusterClusterConfigEndpointConfig.to_proto(resource.endpoint_config)
            )
        else:
            res.ClearField("endpoint_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterClusterConfig(
            staging_bucket=resource.staging_bucket,
            temp_bucket=resource.temp_bucket,
            gce_cluster_config=resource.gce_cluster_config,
            master_config=resource.master_config,
            worker_config=resource.worker_config,
            secondary_worker_config=resource.secondary_worker_config,
            software_config=resource.software_config,
            initialization_actions=resource.initialization_actions,
            encryption_config=resource.encryption_config,
            autoscaling_config=resource.autoscaling_config,
            security_config=resource.security_config,
            lifecycle_config=resource.lifecycle_config,
            endpoint_config=resource.endpoint_config,
        )


class ClusterClusterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterClusterConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterClusterConfig.from_proto(i) for i in resources]


class ClusterClusterConfigGceClusterConfig(object):
    def __init__(
        self,
        zone: str = None,
        network: str = None,
        subnetwork: str = None,
        internal_ip_only: bool = None,
        private_ipv6_google_access: str = None,
        service_account: str = None,
        service_account_scopes: list = None,
        tags: list = None,
        metadata: dict = None,
        reservation_affinity: dict = None,
        node_group_affinity: dict = None,
    ):
        self.zone = zone
        self.network = network
        self.subnetwork = subnetwork
        self.internal_ip_only = internal_ip_only
        self.private_ipv6_google_access = private_ipv6_google_access
        self.service_account = service_account
        self.service_account_scopes = service_account_scopes
        self.tags = tags
        self.metadata = metadata
        self.reservation_affinity = reservation_affinity
        self.node_group_affinity = node_group_affinity

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterClusterConfigGceClusterConfig()
        if Primitive.to_proto(resource.zone):
            res.zone = Primitive.to_proto(resource.zone)
        if Primitive.to_proto(resource.network):
            res.network = Primitive.to_proto(resource.network)
        if Primitive.to_proto(resource.subnetwork):
            res.subnetwork = Primitive.to_proto(resource.subnetwork)
        if Primitive.to_proto(resource.internal_ip_only):
            res.internal_ip_only = Primitive.to_proto(resource.internal_ip_only)
        if ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.to_proto(
            resource.private_ipv6_google_access
        ):
            res.private_ipv6_google_access = ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.to_proto(
                resource.private_ipv6_google_access
            )
        if Primitive.to_proto(resource.service_account):
            res.service_account = Primitive.to_proto(resource.service_account)
        if Primitive.to_proto(resource.service_account_scopes):
            res.service_account_scopes.extend(
                Primitive.to_proto(resource.service_account_scopes)
            )
        if Primitive.to_proto(resource.tags):
            res.tags.extend(Primitive.to_proto(resource.tags))
        if Primitive.to_proto(resource.metadata):
            res.metadata = Primitive.to_proto(resource.metadata)
        if ClusterClusterConfigGceClusterConfigReservationAffinity.to_proto(
            resource.reservation_affinity
        ):
            res.reservation_affinity.CopyFrom(
                ClusterClusterConfigGceClusterConfigReservationAffinity.to_proto(
                    resource.reservation_affinity
                )
            )
        else:
            res.ClearField("reservation_affinity")
        if ClusterClusterConfigGceClusterConfigNodeGroupAffinity.to_proto(
            resource.node_group_affinity
        ):
            res.node_group_affinity.CopyFrom(
                ClusterClusterConfigGceClusterConfigNodeGroupAffinity.to_proto(
                    resource.node_group_affinity
                )
            )
        else:
            res.ClearField("node_group_affinity")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterClusterConfigGceClusterConfig(
            zone=resource.zone,
            network=resource.network,
            subnetwork=resource.subnetwork,
            internal_ip_only=resource.internal_ip_only,
            private_ipv6_google_access=resource.private_ipv6_google_access,
            service_account=resource.service_account,
            service_account_scopes=resource.service_account_scopes,
            tags=resource.tags,
            metadata=resource.metadata,
            reservation_affinity=resource.reservation_affinity,
            node_group_affinity=resource.node_group_affinity,
        )


class ClusterClusterConfigGceClusterConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterClusterConfigGceClusterConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterClusterConfigGceClusterConfig.from_proto(i) for i in resources]


class ClusterClusterConfigGceClusterConfigReservationAffinity(object):
    def __init__(
        self, consume_reservation_type: str = None, key: str = None, values: list = None
    ):
        self.consume_reservation_type = consume_reservation_type
        self.key = key
        self.values = values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocClusterClusterConfigGceClusterConfigReservationAffinity()
        )
        if ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.to_proto(
            resource.consume_reservation_type
        ):
            res.consume_reservation_type = ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.to_proto(
                resource.consume_reservation_type
            )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.values):
            res.values.extend(Primitive.to_proto(resource.values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterClusterConfigGceClusterConfigReservationAffinity(
            consume_reservation_type=resource.consume_reservation_type,
            key=resource.key,
            values=resource.values,
        )


class ClusterClusterConfigGceClusterConfigReservationAffinityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterClusterConfigGceClusterConfigReservationAffinity.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterClusterConfigGceClusterConfigReservationAffinity.from_proto(i)
            for i in resources
        ]


class ClusterClusterConfigGceClusterConfigNodeGroupAffinity(object):
    def __init__(self, node_group: str = None):
        self.node_group = node_group

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            cluster_pb2.DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity()
        )
        if Primitive.to_proto(resource.node_group):
            res.node_group = Primitive.to_proto(resource.node_group)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterClusterConfigGceClusterConfigNodeGroupAffinity(
            node_group=resource.node_group,
        )


class ClusterClusterConfigGceClusterConfigNodeGroupAffinityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterClusterConfigGceClusterConfigNodeGroupAffinity.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterClusterConfigGceClusterConfigNodeGroupAffinity.from_proto(i)
            for i in resources
        ]


class ClusterInstanceGroupConfig(object):
    def __init__(
        self,
        num_instances: int = None,
        instance_names: list = None,
        image: str = None,
        machine_type: str = None,
        disk_config: dict = None,
        is_preemptible: bool = None,
        preemptibility: str = None,
        managed_group_config: dict = None,
        accelerators: list = None,
        min_cpu_platform: str = None,
    ):
        self.num_instances = num_instances
        self.instance_names = instance_names
        self.image = image
        self.machine_type = machine_type
        self.disk_config = disk_config
        self.is_preemptible = is_preemptible
        self.preemptibility = preemptibility
        self.managed_group_config = managed_group_config
        self.accelerators = accelerators
        self.min_cpu_platform = min_cpu_platform

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterInstanceGroupConfig()
        if Primitive.to_proto(resource.num_instances):
            res.num_instances = Primitive.to_proto(resource.num_instances)
        if Primitive.to_proto(resource.instance_names):
            res.instance_names.extend(Primitive.to_proto(resource.instance_names))
        if Primitive.to_proto(resource.image):
            res.image = Primitive.to_proto(resource.image)
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if ClusterInstanceGroupConfigDiskConfig.to_proto(resource.disk_config):
            res.disk_config.CopyFrom(
                ClusterInstanceGroupConfigDiskConfig.to_proto(resource.disk_config)
            )
        else:
            res.ClearField("disk_config")
        if Primitive.to_proto(resource.is_preemptible):
            res.is_preemptible = Primitive.to_proto(resource.is_preemptible)
        if ClusterInstanceGroupConfigPreemptibilityEnum.to_proto(
            resource.preemptibility
        ):
            res.preemptibility = ClusterInstanceGroupConfigPreemptibilityEnum.to_proto(
                resource.preemptibility
            )
        if ClusterInstanceGroupConfigManagedGroupConfig.to_proto(
            resource.managed_group_config
        ):
            res.managed_group_config.CopyFrom(
                ClusterInstanceGroupConfigManagedGroupConfig.to_proto(
                    resource.managed_group_config
                )
            )
        else:
            res.ClearField("managed_group_config")
        if ClusterInstanceGroupConfigAcceleratorsArray.to_proto(resource.accelerators):
            res.accelerators.extend(
                ClusterInstanceGroupConfigAcceleratorsArray.to_proto(
                    resource.accelerators
                )
            )
        if Primitive.to_proto(resource.min_cpu_platform):
            res.min_cpu_platform = Primitive.to_proto(resource.min_cpu_platform)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterInstanceGroupConfig(
            num_instances=resource.num_instances,
            instance_names=resource.instance_names,
            image=resource.image,
            machine_type=resource.machine_type,
            disk_config=resource.disk_config,
            is_preemptible=resource.is_preemptible,
            preemptibility=resource.preemptibility,
            managed_group_config=resource.managed_group_config,
            accelerators=resource.accelerators,
            min_cpu_platform=resource.min_cpu_platform,
        )


class ClusterInstanceGroupConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterInstanceGroupConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterInstanceGroupConfig.from_proto(i) for i in resources]


class ClusterInstanceGroupConfigDiskConfig(object):
    def __init__(
        self,
        boot_disk_type: str = None,
        boot_disk_size_gb: int = None,
        num_local_ssds: int = None,
    ):
        self.boot_disk_type = boot_disk_type
        self.boot_disk_size_gb = boot_disk_size_gb
        self.num_local_ssds = num_local_ssds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterInstanceGroupConfigDiskConfig()
        if Primitive.to_proto(resource.boot_disk_type):
            res.boot_disk_type = Primitive.to_proto(resource.boot_disk_type)
        if Primitive.to_proto(resource.boot_disk_size_gb):
            res.boot_disk_size_gb = Primitive.to_proto(resource.boot_disk_size_gb)
        if Primitive.to_proto(resource.num_local_ssds):
            res.num_local_ssds = Primitive.to_proto(resource.num_local_ssds)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterInstanceGroupConfigDiskConfig(
            boot_disk_type=resource.boot_disk_type,
            boot_disk_size_gb=resource.boot_disk_size_gb,
            num_local_ssds=resource.num_local_ssds,
        )


class ClusterInstanceGroupConfigDiskConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterInstanceGroupConfigDiskConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterInstanceGroupConfigDiskConfig.from_proto(i) for i in resources]


class ClusterInstanceGroupConfigManagedGroupConfig(object):
    def __init__(
        self,
        instance_template_name: str = None,
        instance_group_manager_name: str = None,
    ):
        self.instance_template_name = instance_template_name
        self.instance_group_manager_name = instance_group_manager_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterInstanceGroupConfigManagedGroupConfig()
        if Primitive.to_proto(resource.instance_template_name):
            res.instance_template_name = Primitive.to_proto(
                resource.instance_template_name
            )
        if Primitive.to_proto(resource.instance_group_manager_name):
            res.instance_group_manager_name = Primitive.to_proto(
                resource.instance_group_manager_name
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterInstanceGroupConfigManagedGroupConfig(
            instance_template_name=resource.instance_template_name,
            instance_group_manager_name=resource.instance_group_manager_name,
        )


class ClusterInstanceGroupConfigManagedGroupConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterInstanceGroupConfigManagedGroupConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterInstanceGroupConfigManagedGroupConfig.from_proto(i)
            for i in resources
        ]


class ClusterInstanceGroupConfigAccelerators(object):
    def __init__(self, accelerator_type: str = None, accelerator_count: int = None):
        self.accelerator_type = accelerator_type
        self.accelerator_count = accelerator_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterInstanceGroupConfigAccelerators()
        if Primitive.to_proto(resource.accelerator_type):
            res.accelerator_type = Primitive.to_proto(resource.accelerator_type)
        if Primitive.to_proto(resource.accelerator_count):
            res.accelerator_count = Primitive.to_proto(resource.accelerator_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterInstanceGroupConfigAccelerators(
            accelerator_type=resource.accelerator_type,
            accelerator_count=resource.accelerator_count,
        )


class ClusterInstanceGroupConfigAcceleratorsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterInstanceGroupConfigAccelerators.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterInstanceGroupConfigAccelerators.from_proto(i) for i in resources]


class ClusterClusterConfigSoftwareConfig(object):
    def __init__(
        self,
        image_version: str = None,
        properties: dict = None,
        optional_components: list = None,
    ):
        self.image_version = image_version
        self.properties = properties
        self.optional_components = optional_components

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterClusterConfigSoftwareConfig()
        if Primitive.to_proto(resource.image_version):
            res.image_version = Primitive.to_proto(resource.image_version)
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if ClusterClusterConfigSoftwareConfigOptionalComponentsEnumArray.to_proto(
            resource.optional_components
        ):
            res.optional_components.extend(
                ClusterClusterConfigSoftwareConfigOptionalComponentsEnumArray.to_proto(
                    resource.optional_components
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterClusterConfigSoftwareConfig(
            image_version=resource.image_version,
            properties=resource.properties,
            optional_components=resource.optional_components,
        )


class ClusterClusterConfigSoftwareConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterClusterConfigSoftwareConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterClusterConfigSoftwareConfig.from_proto(i) for i in resources]


class ClusterClusterConfigInitializationActions(object):
    def __init__(self, executable_file: str = None, execution_timeout: str = None):
        self.executable_file = executable_file
        self.execution_timeout = execution_timeout

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterClusterConfigInitializationActions()
        if Primitive.to_proto(resource.executable_file):
            res.executable_file = Primitive.to_proto(resource.executable_file)
        if Primitive.to_proto(resource.execution_timeout):
            res.execution_timeout = Primitive.to_proto(resource.execution_timeout)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterClusterConfigInitializationActions(
            executable_file=resource.executable_file,
            execution_timeout=resource.execution_timeout,
        )


class ClusterClusterConfigInitializationActionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterClusterConfigInitializationActions.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterClusterConfigInitializationActions.from_proto(i) for i in resources
        ]


class ClusterClusterConfigEncryptionConfig(object):
    def __init__(self, gce_pd_kms_key_name: str = None):
        self.gce_pd_kms_key_name = gce_pd_kms_key_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterClusterConfigEncryptionConfig()
        if Primitive.to_proto(resource.gce_pd_kms_key_name):
            res.gce_pd_kms_key_name = Primitive.to_proto(resource.gce_pd_kms_key_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterClusterConfigEncryptionConfig(
            gce_pd_kms_key_name=resource.gce_pd_kms_key_name,
        )


class ClusterClusterConfigEncryptionConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterClusterConfigEncryptionConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterClusterConfigEncryptionConfig.from_proto(i) for i in resources]


class ClusterClusterConfigAutoscalingConfig(object):
    def __init__(self, policy: str = None):
        self.policy = policy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterClusterConfigAutoscalingConfig()
        if Primitive.to_proto(resource.policy):
            res.policy = Primitive.to_proto(resource.policy)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterClusterConfigAutoscalingConfig(policy=resource.policy,)


class ClusterClusterConfigAutoscalingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterClusterConfigAutoscalingConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterClusterConfigAutoscalingConfig.from_proto(i) for i in resources]


class ClusterClusterConfigSecurityConfig(object):
    def __init__(self, kerberos_config: dict = None):
        self.kerberos_config = kerberos_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterClusterConfigSecurityConfig()
        if ClusterClusterConfigSecurityConfigKerberosConfig.to_proto(
            resource.kerberos_config
        ):
            res.kerberos_config.CopyFrom(
                ClusterClusterConfigSecurityConfigKerberosConfig.to_proto(
                    resource.kerberos_config
                )
            )
        else:
            res.ClearField("kerberos_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterClusterConfigSecurityConfig(
            kerberos_config=resource.kerberos_config,
        )


class ClusterClusterConfigSecurityConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterClusterConfigSecurityConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterClusterConfigSecurityConfig.from_proto(i) for i in resources]


class ClusterClusterConfigSecurityConfigKerberosConfig(object):
    def __init__(
        self,
        enable_kerberos: bool = None,
        root_principal_password: str = None,
        kms_key: str = None,
        keystore: str = None,
        truststore: str = None,
        keystore_password: str = None,
        key_password: str = None,
        truststore_password: str = None,
        cross_realm_trust_realm: str = None,
        cross_realm_trust_kdc: str = None,
        cross_realm_trust_admin_server: str = None,
        cross_realm_trust_shared_password: str = None,
        kdc_db_key: str = None,
        tgt_lifetime_hours: int = None,
        realm: str = None,
    ):
        self.enable_kerberos = enable_kerberos
        self.root_principal_password = root_principal_password
        self.kms_key = kms_key
        self.keystore = keystore
        self.truststore = truststore
        self.keystore_password = keystore_password
        self.key_password = key_password
        self.truststore_password = truststore_password
        self.cross_realm_trust_realm = cross_realm_trust_realm
        self.cross_realm_trust_kdc = cross_realm_trust_kdc
        self.cross_realm_trust_admin_server = cross_realm_trust_admin_server
        self.cross_realm_trust_shared_password = cross_realm_trust_shared_password
        self.kdc_db_key = kdc_db_key
        self.tgt_lifetime_hours = tgt_lifetime_hours
        self.realm = realm

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterClusterConfigSecurityConfigKerberosConfig()
        if Primitive.to_proto(resource.enable_kerberos):
            res.enable_kerberos = Primitive.to_proto(resource.enable_kerberos)
        if Primitive.to_proto(resource.root_principal_password):
            res.root_principal_password = Primitive.to_proto(
                resource.root_principal_password
            )
        if Primitive.to_proto(resource.kms_key):
            res.kms_key = Primitive.to_proto(resource.kms_key)
        if Primitive.to_proto(resource.keystore):
            res.keystore = Primitive.to_proto(resource.keystore)
        if Primitive.to_proto(resource.truststore):
            res.truststore = Primitive.to_proto(resource.truststore)
        if Primitive.to_proto(resource.keystore_password):
            res.keystore_password = Primitive.to_proto(resource.keystore_password)
        if Primitive.to_proto(resource.key_password):
            res.key_password = Primitive.to_proto(resource.key_password)
        if Primitive.to_proto(resource.truststore_password):
            res.truststore_password = Primitive.to_proto(resource.truststore_password)
        if Primitive.to_proto(resource.cross_realm_trust_realm):
            res.cross_realm_trust_realm = Primitive.to_proto(
                resource.cross_realm_trust_realm
            )
        if Primitive.to_proto(resource.cross_realm_trust_kdc):
            res.cross_realm_trust_kdc = Primitive.to_proto(
                resource.cross_realm_trust_kdc
            )
        if Primitive.to_proto(resource.cross_realm_trust_admin_server):
            res.cross_realm_trust_admin_server = Primitive.to_proto(
                resource.cross_realm_trust_admin_server
            )
        if Primitive.to_proto(resource.cross_realm_trust_shared_password):
            res.cross_realm_trust_shared_password = Primitive.to_proto(
                resource.cross_realm_trust_shared_password
            )
        if Primitive.to_proto(resource.kdc_db_key):
            res.kdc_db_key = Primitive.to_proto(resource.kdc_db_key)
        if Primitive.to_proto(resource.tgt_lifetime_hours):
            res.tgt_lifetime_hours = Primitive.to_proto(resource.tgt_lifetime_hours)
        if Primitive.to_proto(resource.realm):
            res.realm = Primitive.to_proto(resource.realm)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterClusterConfigSecurityConfigKerberosConfig(
            enable_kerberos=resource.enable_kerberos,
            root_principal_password=resource.root_principal_password,
            kms_key=resource.kms_key,
            keystore=resource.keystore,
            truststore=resource.truststore,
            keystore_password=resource.keystore_password,
            key_password=resource.key_password,
            truststore_password=resource.truststore_password,
            cross_realm_trust_realm=resource.cross_realm_trust_realm,
            cross_realm_trust_kdc=resource.cross_realm_trust_kdc,
            cross_realm_trust_admin_server=resource.cross_realm_trust_admin_server,
            cross_realm_trust_shared_password=resource.cross_realm_trust_shared_password,
            kdc_db_key=resource.kdc_db_key,
            tgt_lifetime_hours=resource.tgt_lifetime_hours,
            realm=resource.realm,
        )


class ClusterClusterConfigSecurityConfigKerberosConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ClusterClusterConfigSecurityConfigKerberosConfig.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ClusterClusterConfigSecurityConfigKerberosConfig.from_proto(i)
            for i in resources
        ]


class ClusterClusterConfigLifecycleConfig(object):
    def __init__(
        self,
        idle_delete_ttl: str = None,
        auto_delete_time: str = None,
        auto_delete_ttl: str = None,
        idle_start_time: str = None,
    ):
        self.idle_delete_ttl = idle_delete_ttl
        self.auto_delete_time = auto_delete_time
        self.auto_delete_ttl = auto_delete_ttl
        self.idle_start_time = idle_start_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterClusterConfigLifecycleConfig()
        if Primitive.to_proto(resource.idle_delete_ttl):
            res.idle_delete_ttl = Primitive.to_proto(resource.idle_delete_ttl)
        if Primitive.to_proto(resource.auto_delete_time):
            res.auto_delete_time = Primitive.to_proto(resource.auto_delete_time)
        if Primitive.to_proto(resource.auto_delete_ttl):
            res.auto_delete_ttl = Primitive.to_proto(resource.auto_delete_ttl)
        if Primitive.to_proto(resource.idle_start_time):
            res.idle_start_time = Primitive.to_proto(resource.idle_start_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterClusterConfigLifecycleConfig(
            idle_delete_ttl=resource.idle_delete_ttl,
            auto_delete_time=resource.auto_delete_time,
            auto_delete_ttl=resource.auto_delete_ttl,
            idle_start_time=resource.idle_start_time,
        )


class ClusterClusterConfigLifecycleConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterClusterConfigLifecycleConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterClusterConfigLifecycleConfig.from_proto(i) for i in resources]


class ClusterClusterConfigEndpointConfig(object):
    def __init__(self, http_ports: dict = None, enable_http_port_access: bool = None):
        self.http_ports = http_ports
        self.enable_http_port_access = enable_http_port_access

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterClusterConfigEndpointConfig()
        if Primitive.to_proto(resource.http_ports):
            res.http_ports = Primitive.to_proto(resource.http_ports)
        if Primitive.to_proto(resource.enable_http_port_access):
            res.enable_http_port_access = Primitive.to_proto(
                resource.enable_http_port_access
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterClusterConfigEndpointConfig(
            http_ports=resource.http_ports,
            enable_http_port_access=resource.enable_http_port_access,
        )


class ClusterClusterConfigEndpointConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterClusterConfigEndpointConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterClusterConfigEndpointConfig.from_proto(i) for i in resources]


class ClusterStatus(object):
    def __init__(
        self,
        state: str = None,
        detail: str = None,
        state_start_time: str = None,
        substate: str = None,
    ):
        self.state = state
        self.detail = detail
        self.state_start_time = state_start_time
        self.substate = substate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterStatus()
        if ClusterStatusStateEnum.to_proto(resource.state):
            res.state = ClusterStatusStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.detail):
            res.detail = Primitive.to_proto(resource.detail)
        if Primitive.to_proto(resource.state_start_time):
            res.state_start_time = Primitive.to_proto(resource.state_start_time)
        if ClusterStatusSubstateEnum.to_proto(resource.substate):
            res.substate = ClusterStatusSubstateEnum.to_proto(resource.substate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterStatus(
            state=resource.state,
            detail=resource.detail,
            state_start_time=resource.state_start_time,
            substate=resource.substate,
        )


class ClusterStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterStatus.from_proto(i) for i in resources]


class ClusterStatusHistory(object):
    def __init__(
        self,
        state: str = None,
        detail: str = None,
        state_start_time: str = None,
        substate: str = None,
    ):
        self.state = state
        self.detail = detail
        self.state_start_time = state_start_time
        self.substate = substate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterStatusHistory()
        if ClusterStatusHistoryStateEnum.to_proto(resource.state):
            res.state = ClusterStatusHistoryStateEnum.to_proto(resource.state)
        if Primitive.to_proto(resource.detail):
            res.detail = Primitive.to_proto(resource.detail)
        if Primitive.to_proto(resource.state_start_time):
            res.state_start_time = Primitive.to_proto(resource.state_start_time)
        if ClusterStatusHistorySubstateEnum.to_proto(resource.substate):
            res.substate = ClusterStatusHistorySubstateEnum.to_proto(resource.substate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterStatusHistory(
            state=resource.state,
            detail=resource.detail,
            state_start_time=resource.state_start_time,
            substate=resource.substate,
        )


class ClusterStatusHistoryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterStatusHistory.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterStatusHistory.from_proto(i) for i in resources]


class ClusterMetrics(object):
    def __init__(self, hdfs_metrics: dict = None, yarn_metrics: dict = None):
        self.hdfs_metrics = hdfs_metrics
        self.yarn_metrics = yarn_metrics

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.DataprocClusterMetrics()
        if Primitive.to_proto(resource.hdfs_metrics):
            res.hdfs_metrics = Primitive.to_proto(resource.hdfs_metrics)
        if Primitive.to_proto(resource.yarn_metrics):
            res.yarn_metrics = Primitive.to_proto(resource.yarn_metrics)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterMetrics(
            hdfs_metrics=resource.hdfs_metrics, yarn_metrics=resource.yarn_metrics,
        )


class ClusterMetricsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterMetrics.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterMetrics.from_proto(i) for i in resources]


class ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.Value(
            "DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum.Name(
            resource
        )[
            len(
                "DataprocClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"
            ) :
        ]


class ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.Value(
            "DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum.Name(
            resource
        )[
            len(
                "DataprocClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"
            ) :
        ]


class ClusterInstanceGroupConfigPreemptibilityEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterInstanceGroupConfigPreemptibilityEnum.Value(
            "DataprocClusterInstanceGroupConfigPreemptibilityEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterInstanceGroupConfigPreemptibilityEnum.Name(
            resource
        )[len("DataprocClusterInstanceGroupConfigPreemptibilityEnum") :]


class ClusterClusterConfigSoftwareConfigOptionalComponentsEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum.Value(
            "DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum.Name(
            resource
        )[
            len("DataprocClusterClusterConfigSoftwareConfigOptionalComponentsEnum") :
        ]


class ClusterStatusStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterStatusStateEnum.Value(
            "DataprocClusterStatusStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterStatusStateEnum.Name(resource)[
            len("DataprocClusterStatusStateEnum") :
        ]


class ClusterStatusSubstateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterStatusSubstateEnum.Value(
            "DataprocClusterStatusSubstateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterStatusSubstateEnum.Name(resource)[
            len("DataprocClusterStatusSubstateEnum") :
        ]


class ClusterStatusHistoryStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterStatusHistoryStateEnum.Value(
            "DataprocClusterStatusHistoryStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterStatusHistoryStateEnum.Name(resource)[
            len("DataprocClusterStatusHistoryStateEnum") :
        ]


class ClusterStatusHistorySubstateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterStatusHistorySubstateEnum.Value(
            "DataprocClusterStatusHistorySubstateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.DataprocClusterStatusHistorySubstateEnum.Name(resource)[
            len("DataprocClusterStatusHistorySubstateEnum") :
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
