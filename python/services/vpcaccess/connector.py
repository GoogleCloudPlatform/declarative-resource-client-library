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
from google3.cloud.graphite.mmv2.services.google.vpc_access import connector_pb2
from google3.cloud.graphite.mmv2.services.google.vpc_access import connector_pb2_grpc

from typing import List


class Connector(object):
    def __init__(
        self,
        name: str = None,
        network: str = None,
        ip_cidr_range: str = None,
        min_throughput: int = None,
        max_throughput: int = None,
        project: str = None,
        location: str = None,
        state: str = None,
        self_link: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.network = network
        self.ip_cidr_range = ip_cidr_range
        self.min_throughput = min_throughput
        self.max_throughput = max_throughput
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = connector_pb2_grpc.VpcaccessConnectorServiceStub(channel.Channel())
        request = connector_pb2.ApplyVpcaccessConnectorRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.ip_cidr_range):
            request.resource.ip_cidr_range = Primitive.to_proto(self.ip_cidr_range)

        if Primitive.to_proto(self.min_throughput):
            request.resource.min_throughput = Primitive.to_proto(self.min_throughput)

        if Primitive.to_proto(self.max_throughput):
            request.resource.max_throughput = Primitive.to_proto(self.max_throughput)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyVpcaccessConnector(request)
        self.name = Primitive.from_proto(response.name)
        self.network = Primitive.from_proto(response.network)
        self.ip_cidr_range = Primitive.from_proto(response.ip_cidr_range)
        self.min_throughput = Primitive.from_proto(response.min_throughput)
        self.max_throughput = Primitive.from_proto(response.max_throughput)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.state = ConnectorStateEnum.from_proto(response.state)
        self.self_link = Primitive.from_proto(response.self_link)

    def delete(self):
        stub = connector_pb2_grpc.VpcaccessConnectorServiceStub(channel.Channel())
        request = connector_pb2.DeleteVpcaccessConnectorRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.network):
            request.resource.network = Primitive.to_proto(self.network)

        if Primitive.to_proto(self.ip_cidr_range):
            request.resource.ip_cidr_range = Primitive.to_proto(self.ip_cidr_range)

        if Primitive.to_proto(self.min_throughput):
            request.resource.min_throughput = Primitive.to_proto(self.min_throughput)

        if Primitive.to_proto(self.max_throughput):
            request.resource.max_throughput = Primitive.to_proto(self.max_throughput)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteVpcaccessConnector(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = connector_pb2_grpc.VpcaccessConnectorServiceStub(channel.Channel())
        request = connector_pb2.ListVpcaccessConnectorRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListVpcaccessConnector(request).items

    def to_proto(self):
        resource = connector_pb2.VpcaccessConnector()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.network):
            resource.network = Primitive.to_proto(self.network)
        if Primitive.to_proto(self.ip_cidr_range):
            resource.ip_cidr_range = Primitive.to_proto(self.ip_cidr_range)
        if Primitive.to_proto(self.min_throughput):
            resource.min_throughput = Primitive.to_proto(self.min_throughput)
        if Primitive.to_proto(self.max_throughput):
            resource.max_throughput = Primitive.to_proto(self.max_throughput)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class ConnectorStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return connector_pb2.VpcaccessConnectorStateEnum.Value(
            "VpcaccessConnectorStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return connector_pb2.VpcaccessConnectorStateEnum.Name(resource)[
            len("VpcaccessConnectorStateEnum") :
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
