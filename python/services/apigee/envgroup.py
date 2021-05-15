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
from google3.cloud.graphite.mmv2.services.google.apigee import envgroup_pb2
from google3.cloud.graphite.mmv2.services.google.apigee import envgroup_pb2_grpc

from typing import List


class Envgroup(object):
    def __init__(
        self,
        name: str = None,
        hostnames: list = None,
        created_at: int = None,
        last_modified_at: int = None,
        state: str = None,
        organization: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.hostnames = hostnames
        self.organization = organization
        self.service_account_file = service_account_file

    def apply(self):
        stub = envgroup_pb2_grpc.ApigeeEnvgroupServiceStub(channel.Channel())
        request = envgroup_pb2.ApplyApigeeEnvgroupRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.hostnames):
            request.resource.hostnames.extend(Primitive.to_proto(self.hostnames))
        if Primitive.to_proto(self.organization):
            request.resource.organization = Primitive.to_proto(self.organization)

        request.service_account_file = self.service_account_file

        response = stub.ApplyApigeeEnvgroup(request)
        self.name = Primitive.from_proto(response.name)
        self.hostnames = Primitive.from_proto(response.hostnames)
        self.created_at = Primitive.from_proto(response.created_at)
        self.last_modified_at = Primitive.from_proto(response.last_modified_at)
        self.state = EnvgroupStateEnum.from_proto(response.state)
        self.organization = Primitive.from_proto(response.organization)

    def delete(self):
        stub = envgroup_pb2_grpc.ApigeeEnvgroupServiceStub(channel.Channel())
        request = envgroup_pb2.DeleteApigeeEnvgroupRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.hostnames):
            request.resource.hostnames.extend(Primitive.to_proto(self.hostnames))
        if Primitive.to_proto(self.organization):
            request.resource.organization = Primitive.to_proto(self.organization)

        response = stub.DeleteApigeeEnvgroup(request)

    @classmethod
    def list(self, organization, service_account_file=""):
        stub = envgroup_pb2_grpc.ApigeeEnvgroupServiceStub(channel.Channel())
        request = envgroup_pb2.ListApigeeEnvgroupRequest()
        request.service_account_file = service_account_file
        request.Organization = organization

        return stub.ListApigeeEnvgroup(request).items

    def to_proto(self):
        resource = envgroup_pb2.ApigeeEnvgroup()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.hostnames):
            resource.hostnames.extend(Primitive.to_proto(self.hostnames))
        if Primitive.to_proto(self.organization):
            resource.organization = Primitive.to_proto(self.organization)
        return resource


class EnvgroupStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return envgroup_pb2.ApigeeEnvgroupStateEnum.Value(
            "ApigeeEnvgroupStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return envgroup_pb2.ApigeeEnvgroupStateEnum.Name(resource)[
            len("ApigeeEnvgroupStateEnum") :
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
