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
from google3.cloud.graphite.mmv2.services.google.network_connectivity import spoke_pb2
from google3.cloud.graphite.mmv2.services.google.network_connectivity import (
    spoke_pb2_grpc,
)

from typing import List


class Spoke(object):
    def __init__(
        self,
        name: str = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        description: str = None,
        hub: str = None,
        linked_vpn_tunnels: list = None,
        linked_interconnect_attachments: list = None,
        linked_router_appliance_instances: list = None,
        unique_id: str = None,
        state: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.labels = labels
        self.description = description
        self.hub = hub
        self.linked_vpn_tunnels = linked_vpn_tunnels
        self.linked_interconnect_attachments = linked_interconnect_attachments
        self.linked_router_appliance_instances = linked_router_appliance_instances
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = spoke_pb2_grpc.NetworkconnectivityAlphaSpokeServiceStub(
            channel.Channel()
        )
        request = spoke_pb2.ApplyNetworkconnectivityAlphaSpokeRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.hub):
            request.resource.hub = Primitive.to_proto(self.hub)

        if Primitive.to_proto(self.linked_vpn_tunnels):
            request.resource.linked_vpn_tunnels.extend(
                Primitive.to_proto(self.linked_vpn_tunnels)
            )
        if Primitive.to_proto(self.linked_interconnect_attachments):
            request.resource.linked_interconnect_attachments.extend(
                Primitive.to_proto(self.linked_interconnect_attachments)
            )
        if SpokeLinkedRouterApplianceInstancesArray.to_proto(
            self.linked_router_appliance_instances
        ):
            request.resource.linked_router_appliance_instances.extend(
                SpokeLinkedRouterApplianceInstancesArray.to_proto(
                    self.linked_router_appliance_instances
                )
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyNetworkconnectivityAlphaSpoke(request)
        self.name = Primitive.from_proto(response.name)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.description = Primitive.from_proto(response.description)
        self.hub = Primitive.from_proto(response.hub)
        self.linked_vpn_tunnels = Primitive.from_proto(response.linked_vpn_tunnels)
        self.linked_interconnect_attachments = Primitive.from_proto(
            response.linked_interconnect_attachments
        )
        self.linked_router_appliance_instances = SpokeLinkedRouterApplianceInstancesArray.from_proto(
            response.linked_router_appliance_instances
        )
        self.unique_id = Primitive.from_proto(response.unique_id)
        self.state = SpokeStateEnum.from_proto(response.state)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = spoke_pb2_grpc.NetworkconnectivityAlphaSpokeServiceStub(
            channel.Channel()
        )
        request = spoke_pb2.DeleteNetworkconnectivityAlphaSpokeRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.hub):
            request.resource.hub = Primitive.to_proto(self.hub)

        if Primitive.to_proto(self.linked_vpn_tunnels):
            request.resource.linked_vpn_tunnels.extend(
                Primitive.to_proto(self.linked_vpn_tunnels)
            )
        if Primitive.to_proto(self.linked_interconnect_attachments):
            request.resource.linked_interconnect_attachments.extend(
                Primitive.to_proto(self.linked_interconnect_attachments)
            )
        if SpokeLinkedRouterApplianceInstancesArray.to_proto(
            self.linked_router_appliance_instances
        ):
            request.resource.linked_router_appliance_instances.extend(
                SpokeLinkedRouterApplianceInstancesArray.to_proto(
                    self.linked_router_appliance_instances
                )
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteNetworkconnectivityAlphaSpoke(request)

    def list(self):
        stub = spoke_pb2_grpc.NetworkconnectivityAlphaSpokeServiceStub(
            channel.Channel()
        )
        request = spoke_pb2.ListNetworkconnectivityAlphaSpokeRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.hub):
            request.resource.hub = Primitive.to_proto(self.hub)

        if Primitive.to_proto(self.linked_vpn_tunnels):
            request.resource.linked_vpn_tunnels.extend(
                Primitive.to_proto(self.linked_vpn_tunnels)
            )
        if Primitive.to_proto(self.linked_interconnect_attachments):
            request.resource.linked_interconnect_attachments.extend(
                Primitive.to_proto(self.linked_interconnect_attachments)
            )
        if SpokeLinkedRouterApplianceInstancesArray.to_proto(
            self.linked_router_appliance_instances
        ):
            request.resource.linked_router_appliance_instances.extend(
                SpokeLinkedRouterApplianceInstancesArray.to_proto(
                    self.linked_router_appliance_instances
                )
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        return stub.ListNetworkconnectivityAlphaSpoke(request).items

    def to_proto(self):
        resource = spoke_pb2.NetworkconnectivityAlphaSpoke()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.hub):
            resource.hub = Primitive.to_proto(self.hub)
        if Primitive.to_proto(self.linked_vpn_tunnels):
            resource.linked_vpn_tunnels.extend(
                Primitive.to_proto(self.linked_vpn_tunnels)
            )
        if Primitive.to_proto(self.linked_interconnect_attachments):
            resource.linked_interconnect_attachments.extend(
                Primitive.to_proto(self.linked_interconnect_attachments)
            )
        if SpokeLinkedRouterApplianceInstancesArray.to_proto(
            self.linked_router_appliance_instances
        ):
            resource.linked_router_appliance_instances.extend(
                SpokeLinkedRouterApplianceInstancesArray.to_proto(
                    self.linked_router_appliance_instances
                )
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class SpokeLinkedRouterApplianceInstances(object):
    def __init__(self, virtual_machine: str = None, ip_address: str = None):
        self.virtual_machine = virtual_machine
        self.ip_address = ip_address

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = spoke_pb2.NetworkconnectivityAlphaSpokeLinkedRouterApplianceInstances()
        if Primitive.to_proto(resource.virtual_machine):
            res.virtual_machine = Primitive.to_proto(resource.virtual_machine)
        if Primitive.to_proto(resource.ip_address):
            res.ip_address = Primitive.to_proto(resource.ip_address)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SpokeLinkedRouterApplianceInstances(
            virtual_machine=Primitive.from_proto(resource.virtual_machine),
            ip_address=Primitive.from_proto(resource.ip_address),
        )


class SpokeLinkedRouterApplianceInstancesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SpokeLinkedRouterApplianceInstances.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SpokeLinkedRouterApplianceInstances.from_proto(i) for i in resources]


class SpokeStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return spoke_pb2.NetworkconnectivityAlphaSpokeStateEnum.Value(
            "NetworkconnectivityAlphaSpokeStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return spoke_pb2.NetworkconnectivityAlphaSpokeStateEnum.Name(resource)[
            len("NetworkconnectivityAlphaSpokeStateEnum") :
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
