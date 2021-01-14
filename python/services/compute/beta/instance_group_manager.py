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
from google3.cloud.graphite.mmv2.services.google.compute import (
    instance_group_manager_pb2,
)
from google3.cloud.graphite.mmv2.services.google.compute import (
    instance_group_manager_pb2_grpc,
)

from typing import List


class InstanceGroupManager(object):
    def __init__(
        self,
        base_instance_name: str = None,
        creation_timestamp: str = None,
        distribution_policy: dict = None,
        current_actions: dict = None,
        description: str = None,
        versions: list = None,
        id: int = None,
        instance_group: str = None,
        instance_template: str = None,
        name: str = None,
        named_ports: list = None,
        status: dict = None,
        target_pools: list = None,
        autohealing_policies: list = None,
        update_policy: dict = None,
        target_size: int = None,
        zone: str = None,
        region: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.base_instance_name = base_instance_name
        self.distribution_policy = distribution_policy
        self.description = description
        self.versions = versions
        self.instance_template = instance_template
        self.name = name
        self.named_ports = named_ports
        self.target_pools = target_pools
        self.autohealing_policies = autohealing_policies
        self.update_policy = update_policy
        self.target_size = target_size
        self.zone = zone
        self.region = region
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = instance_group_manager_pb2_grpc.ComputeBetaInstanceGroupManagerServiceStub(
            channel.Channel()
        )
        request = (
            instance_group_manager_pb2.ApplyComputeBetaInstanceGroupManagerRequest()
        )
        if Primitive.to_proto(self.base_instance_name):
            request.resource.base_instance_name = Primitive.to_proto(
                self.base_instance_name
            )

        if InstanceGroupManagerDistributionPolicy.to_proto(self.distribution_policy):
            request.resource.distribution_policy.CopyFrom(
                InstanceGroupManagerDistributionPolicy.to_proto(
                    self.distribution_policy
                )
            )
        else:
            request.resource.ClearField("distribution_policy")
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if InstanceGroupManagerVersionsArray.to_proto(self.versions):
            request.resource.versions.extend(
                InstanceGroupManagerVersionsArray.to_proto(self.versions)
            )
        if Primitive.to_proto(self.instance_template):
            request.resource.instance_template = Primitive.to_proto(
                self.instance_template
            )

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if InstanceGroupManagerNamedPortsArray.to_proto(self.named_ports):
            request.resource.named_ports.extend(
                InstanceGroupManagerNamedPortsArray.to_proto(self.named_ports)
            )
        if Primitive.to_proto(self.target_pools):
            request.resource.target_pools.extend(Primitive.to_proto(self.target_pools))
        if InstanceGroupManagerAutohealingPoliciesArray.to_proto(
            self.autohealing_policies
        ):
            request.resource.autohealing_policies.extend(
                InstanceGroupManagerAutohealingPoliciesArray.to_proto(
                    self.autohealing_policies
                )
            )
        if InstanceGroupManagerUpdatePolicy.to_proto(self.update_policy):
            request.resource.update_policy.CopyFrom(
                InstanceGroupManagerUpdatePolicy.to_proto(self.update_policy)
            )
        else:
            request.resource.ClearField("update_policy")
        if Primitive.to_proto(self.target_size):
            request.resource.target_size = Primitive.to_proto(self.target_size)

        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaInstanceGroupManager(request)
        self.base_instance_name = Primitive.from_proto(response.base_instance_name)
        self.creation_timestamp = Primitive.from_proto(response.creation_timestamp)
        self.distribution_policy = InstanceGroupManagerDistributionPolicy.from_proto(
            response.distribution_policy
        )
        self.current_actions = InstanceGroupManagerCurrentActions.from_proto(
            response.current_actions
        )
        self.description = Primitive.from_proto(response.description)
        self.versions = InstanceGroupManagerVersionsArray.from_proto(response.versions)
        self.id = Primitive.from_proto(response.id)
        self.instance_group = Primitive.from_proto(response.instance_group)
        self.instance_template = Primitive.from_proto(response.instance_template)
        self.name = Primitive.from_proto(response.name)
        self.named_ports = InstanceGroupManagerNamedPortsArray.from_proto(
            response.named_ports
        )
        self.status = InstanceGroupManagerStatus.from_proto(response.status)
        self.target_pools = Primitive.from_proto(response.target_pools)
        self.autohealing_policies = InstanceGroupManagerAutohealingPoliciesArray.from_proto(
            response.autohealing_policies
        )
        self.update_policy = InstanceGroupManagerUpdatePolicy.from_proto(
            response.update_policy
        )
        self.target_size = Primitive.from_proto(response.target_size)
        self.zone = Primitive.from_proto(response.zone)
        self.region = Primitive.from_proto(response.region)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    @classmethod
    def delete(self, project, location, name, service_account_file=""):
        stub = instance_group_manager_pb2_grpc.ComputeBetaInstanceGroupManagerServiceStub(
            channel.Channel()
        )
        request = (
            instance_group_manager_pb2.DeleteComputeBetaInstanceGroupManagerRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.Name = name

        response = stub.DeleteComputeBetaInstanceGroupManager(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = instance_group_manager_pb2_grpc.ComputeBetaInstanceGroupManagerServiceStub(
            channel.Channel()
        )
        request = (
            instance_group_manager_pb2.ListComputeBetaInstanceGroupManagerRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListComputeBetaInstanceGroupManager(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = instance_group_manager_pb2.ComputeBetaInstanceGroupManager()
        any_proto.Unpack(res_proto)

        res = InstanceGroupManager()
        res.base_instance_name = Primitive.from_proto(res_proto.base_instance_name)
        res.creation_timestamp = Primitive.from_proto(res_proto.creation_timestamp)
        res.distribution_policy = InstanceGroupManagerDistributionPolicy.from_proto(
            res_proto.distribution_policy
        )
        res.current_actions = InstanceGroupManagerCurrentActions.from_proto(
            res_proto.current_actions
        )
        res.description = Primitive.from_proto(res_proto.description)
        res.versions = InstanceGroupManagerVersionsArray.from_proto(res_proto.versions)
        res.id = Primitive.from_proto(res_proto.id)
        res.instance_group = Primitive.from_proto(res_proto.instance_group)
        res.instance_template = Primitive.from_proto(res_proto.instance_template)
        res.name = Primitive.from_proto(res_proto.name)
        res.named_ports = InstanceGroupManagerNamedPortsArray.from_proto(
            res_proto.named_ports
        )
        res.status = InstanceGroupManagerStatus.from_proto(res_proto.status)
        res.target_pools = Primitive.from_proto(res_proto.target_pools)
        res.autohealing_policies = InstanceGroupManagerAutohealingPoliciesArray.from_proto(
            res_proto.autohealing_policies
        )
        res.update_policy = InstanceGroupManagerUpdatePolicy.from_proto(
            res_proto.update_policy
        )
        res.target_size = Primitive.from_proto(res_proto.target_size)
        res.zone = Primitive.from_proto(res_proto.zone)
        res.region = Primitive.from_proto(res_proto.region)
        res.project = Primitive.from_proto(res_proto.project)
        res.location = Primitive.from_proto(res_proto.location)
        return res


class InstanceGroupManagerDistributionPolicy(object):
    def __init__(self, zones: list = None):
        self.zones = zones

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeBetaInstanceGroupManagerDistributionPolicy()
        )
        if InstanceGroupManagerDistributionPolicyZonesArray.to_proto(resource.zones):
            res.zones.extend(
                InstanceGroupManagerDistributionPolicyZonesArray.to_proto(
                    resource.zones
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerDistributionPolicy(zones=resource.zones,)


class InstanceGroupManagerDistributionPolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerDistributionPolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerDistributionPolicy.from_proto(i) for i in resources]


class InstanceGroupManagerDistributionPolicyZones(object):
    def __init__(self, zone: str = None):
        self.zone = zone

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeBetaInstanceGroupManagerDistributionPolicyZones()
        )
        if Primitive.to_proto(resource.zone):
            res.zone = Primitive.to_proto(resource.zone)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerDistributionPolicyZones(zone=resource.zone,)


class InstanceGroupManagerDistributionPolicyZonesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceGroupManagerDistributionPolicyZones.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerDistributionPolicyZones.from_proto(i) for i in resources
        ]


class InstanceGroupManagerCurrentActions(object):
    def __init__(
        self,
        abandoning: int = None,
        creating: int = None,
        creating_without_retries: int = None,
        deleting: int = None,
        none: int = None,
        recreating: int = None,
        refreshing: int = None,
        restarting: int = None,
    ):
        self.abandoning = abandoning
        self.creating = creating
        self.creating_without_retries = creating_without_retries
        self.deleting = deleting
        self.none = none
        self.recreating = recreating
        self.refreshing = refreshing
        self.restarting = restarting

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_group_manager_pb2.ComputeBetaInstanceGroupManagerCurrentActions()
        if Primitive.to_proto(resource.abandoning):
            res.abandoning = Primitive.to_proto(resource.abandoning)
        if Primitive.to_proto(resource.creating):
            res.creating = Primitive.to_proto(resource.creating)
        if Primitive.to_proto(resource.creating_without_retries):
            res.creating_without_retries = Primitive.to_proto(
                resource.creating_without_retries
            )
        if Primitive.to_proto(resource.deleting):
            res.deleting = Primitive.to_proto(resource.deleting)
        if Primitive.to_proto(resource.none):
            res.none = Primitive.to_proto(resource.none)
        if Primitive.to_proto(resource.recreating):
            res.recreating = Primitive.to_proto(resource.recreating)
        if Primitive.to_proto(resource.refreshing):
            res.refreshing = Primitive.to_proto(resource.refreshing)
        if Primitive.to_proto(resource.restarting):
            res.restarting = Primitive.to_proto(resource.restarting)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerCurrentActions(
            abandoning=resource.abandoning,
            creating=resource.creating,
            creating_without_retries=resource.creating_without_retries,
            deleting=resource.deleting,
            none=resource.none,
            recreating=resource.recreating,
            refreshing=resource.refreshing,
            restarting=resource.restarting,
        )


class InstanceGroupManagerCurrentActionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerCurrentActions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerCurrentActions.from_proto(i) for i in resources]


class InstanceGroupManagerVersions(object):
    def __init__(
        self, name: str = None, instance_template: str = None, target_size: dict = None
    ):
        self.name = name
        self.instance_template = instance_template
        self.target_size = target_size

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_group_manager_pb2.ComputeBetaInstanceGroupManagerVersions()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.instance_template):
            res.instance_template = Primitive.to_proto(resource.instance_template)
        if InstanceGroupManagerVersionsTargetSize.to_proto(resource.target_size):
            res.target_size.CopyFrom(
                InstanceGroupManagerVersionsTargetSize.to_proto(resource.target_size)
            )
        else:
            res.ClearField("target_size")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerVersions(
            name=resource.name,
            instance_template=resource.instance_template,
            target_size=resource.target_size,
        )


class InstanceGroupManagerVersionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerVersions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerVersions.from_proto(i) for i in resources]


class InstanceGroupManagerVersionsTargetSize(object):
    def __init__(self, fixed: int = None, percent: int = None, calculated: int = None):
        self.fixed = fixed
        self.percent = percent
        self.calculated = calculated

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeBetaInstanceGroupManagerVersionsTargetSize()
        )
        if Primitive.to_proto(resource.fixed):
            res.fixed = Primitive.to_proto(resource.fixed)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        if Primitive.to_proto(resource.calculated):
            res.calculated = Primitive.to_proto(resource.calculated)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerVersionsTargetSize(
            fixed=resource.fixed,
            percent=resource.percent,
            calculated=resource.calculated,
        )


class InstanceGroupManagerVersionsTargetSizeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerVersionsTargetSize.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerVersionsTargetSize.from_proto(i) for i in resources]


class InstanceGroupManagerNamedPorts(object):
    def __init__(self, name: str = None, port: int = None):
        self.name = name
        self.port = port

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_group_manager_pb2.ComputeBetaInstanceGroupManagerNamedPorts()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerNamedPorts(name=resource.name, port=resource.port,)


class InstanceGroupManagerNamedPortsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerNamedPorts.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerNamedPorts.from_proto(i) for i in resources]


class InstanceGroupManagerStatus(object):
    def __init__(
        self,
        is_stable: bool = None,
        version_target: dict = None,
        autoscalar: str = None,
    ):
        self.is_stable = is_stable
        self.version_target = version_target
        self.autoscalar = autoscalar

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_group_manager_pb2.ComputeBetaInstanceGroupManagerStatus()
        if Primitive.to_proto(resource.is_stable):
            res.is_stable = Primitive.to_proto(resource.is_stable)
        if InstanceGroupManagerStatusVersionTarget.to_proto(resource.version_target):
            res.version_target.CopyFrom(
                InstanceGroupManagerStatusVersionTarget.to_proto(
                    resource.version_target
                )
            )
        else:
            res.ClearField("version_target")
        if Primitive.to_proto(resource.autoscalar):
            res.autoscalar = Primitive.to_proto(resource.autoscalar)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerStatus(
            is_stable=resource.is_stable,
            version_target=resource.version_target,
            autoscalar=resource.autoscalar,
        )


class InstanceGroupManagerStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerStatus.from_proto(i) for i in resources]


class InstanceGroupManagerStatusVersionTarget(object):
    def __init__(self, is_reached: bool = None):
        self.is_reached = is_reached

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeBetaInstanceGroupManagerStatusVersionTarget()
        )
        if Primitive.to_proto(resource.is_reached):
            res.is_reached = Primitive.to_proto(resource.is_reached)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerStatusVersionTarget(is_reached=resource.is_reached,)


class InstanceGroupManagerStatusVersionTargetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerStatusVersionTarget.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerStatusVersionTarget.from_proto(i) for i in resources
        ]


class InstanceGroupManagerAutohealingPolicies(object):
    def __init__(self, health_check: str = None, initial_delay_sec: int = None):
        self.health_check = health_check
        self.initial_delay_sec = initial_delay_sec

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeBetaInstanceGroupManagerAutohealingPolicies()
        )
        if Primitive.to_proto(resource.health_check):
            res.health_check = Primitive.to_proto(resource.health_check)
        if Primitive.to_proto(resource.initial_delay_sec):
            res.initial_delay_sec = Primitive.to_proto(resource.initial_delay_sec)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerAutohealingPolicies(
            health_check=resource.health_check,
            initial_delay_sec=resource.initial_delay_sec,
        )


class InstanceGroupManagerAutohealingPoliciesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerAutohealingPolicies.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerAutohealingPolicies.from_proto(i) for i in resources
        ]


class InstanceGroupManagerUpdatePolicy(object):
    def __init__(
        self,
        instance_redistribution_type: str = None,
        minimal_action: str = None,
        max_surge: dict = None,
    ):
        self.instance_redistribution_type = instance_redistribution_type
        self.minimal_action = minimal_action
        self.max_surge = max_surge

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_group_manager_pb2.ComputeBetaInstanceGroupManagerUpdatePolicy()
        if InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum.to_proto(
            resource.instance_redistribution_type
        ):
            res.instance_redistribution_type = InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum.to_proto(
                resource.instance_redistribution_type
            )
        if InstanceGroupManagerUpdatePolicyMinimalActionEnum.to_proto(
            resource.minimal_action
        ):
            res.minimal_action = InstanceGroupManagerUpdatePolicyMinimalActionEnum.to_proto(
                resource.minimal_action
            )
        if InstanceGroupManagerUpdatePolicyMaxSurge.to_proto(resource.max_surge):
            res.max_surge.CopyFrom(
                InstanceGroupManagerUpdatePolicyMaxSurge.to_proto(resource.max_surge)
            )
        else:
            res.ClearField("max_surge")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerUpdatePolicy(
            instance_redistribution_type=resource.instance_redistribution_type,
            minimal_action=resource.minimal_action,
            max_surge=resource.max_surge,
        )


class InstanceGroupManagerUpdatePolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerUpdatePolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceGroupManagerUpdatePolicy.from_proto(i) for i in resources]


class InstanceGroupManagerUpdatePolicyMaxSurge(object):
    def __init__(
        self,
        fixed: int = None,
        percent: int = None,
        calculated: int = None,
        max_unavailable: dict = None,
    ):
        self.fixed = fixed
        self.percent = percent
        self.calculated = calculated
        self.max_unavailable = max_unavailable

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurge()
        )
        if Primitive.to_proto(resource.fixed):
            res.fixed = Primitive.to_proto(resource.fixed)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        if Primitive.to_proto(resource.calculated):
            res.calculated = Primitive.to_proto(resource.calculated)
        if InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable.to_proto(
            resource.max_unavailable
        ):
            res.max_unavailable.CopyFrom(
                InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable.to_proto(
                    resource.max_unavailable
                )
            )
        else:
            res.ClearField("max_unavailable")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerUpdatePolicyMaxSurge(
            fixed=resource.fixed,
            percent=resource.percent,
            calculated=resource.calculated,
            max_unavailable=resource.max_unavailable,
        )


class InstanceGroupManagerUpdatePolicyMaxSurgeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceGroupManagerUpdatePolicyMaxSurge.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerUpdatePolicyMaxSurge.from_proto(i) for i in resources
        ]


class InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(object):
    def __init__(self, fixed: int = None, percent: int = None, calculated: int = None):
        self.fixed = fixed
        self.percent = percent
        self.calculated = calculated

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_group_manager_pb2.ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable()
        )
        if Primitive.to_proto(resource.fixed):
            res.fixed = Primitive.to_proto(resource.fixed)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        if Primitive.to_proto(resource.calculated):
            res.calculated = Primitive.to_proto(resource.calculated)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable(
            fixed=resource.fixed,
            percent=resource.percent,
            calculated=resource.calculated,
        )


class InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailableArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceGroupManagerUpdatePolicyMaxSurgeMaxUnavailable.from_proto(i)
            for i in resources
        ]


class InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum.Value(
            "ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum.Name(
            resource
        )[
            len(
                "ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"
            ) :
        ]


class InstanceGroupManagerUpdatePolicyMinimalActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum.Value(
            "ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_group_manager_pb2.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum.Name(
            resource
        )[
            len("ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum") :
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
