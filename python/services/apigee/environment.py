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
from google3.cloud.graphite.mmv2.services.google.apigee import environment_pb2
from google3.cloud.graphite.mmv2.services.google.apigee import environment_pb2_grpc

from typing import List


class Environment(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        created_at: int = None,
        last_modified_at: int = None,
        properties: dict = None,
        display_name: str = None,
        state: str = None,
        organization: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.properties = properties
        self.display_name = display_name
        self.organization = organization
        self.service_account_file = service_account_file

    def apply(self):
        stub = environment_pb2_grpc.ApigeeEnvironmentServiceStub(channel.Channel())
        request = environment_pb2.ApplyApigeeEnvironmentRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if EnvironmentProperties.to_proto(self.properties):
            request.resource.properties.CopyFrom(
                EnvironmentProperties.to_proto(self.properties)
            )
        else:
            request.resource.ClearField("properties")
        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.organization):
            request.resource.organization = Primitive.to_proto(self.organization)

        request.service_account_file = self.service_account_file

        response = stub.ApplyApigeeEnvironment(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.created_at = Primitive.from_proto(response.created_at)
        self.last_modified_at = Primitive.from_proto(response.last_modified_at)
        self.properties = EnvironmentProperties.from_proto(response.properties)
        self.display_name = Primitive.from_proto(response.display_name)
        self.state = EnvironmentStateEnum.from_proto(response.state)
        self.organization = Primitive.from_proto(response.organization)

    def delete(self):
        stub = environment_pb2_grpc.ApigeeEnvironmentServiceStub(channel.Channel())
        request = environment_pb2.DeleteApigeeEnvironmentRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if EnvironmentProperties.to_proto(self.properties):
            request.resource.properties.CopyFrom(
                EnvironmentProperties.to_proto(self.properties)
            )
        else:
            request.resource.ClearField("properties")
        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.organization):
            request.resource.organization = Primitive.to_proto(self.organization)

        response = stub.DeleteApigeeEnvironment(request)

    @classmethod
    def list(self, organization, service_account_file=""):
        stub = environment_pb2_grpc.ApigeeEnvironmentServiceStub(channel.Channel())
        request = environment_pb2.ListApigeeEnvironmentRequest()
        request.service_account_file = service_account_file
        request.Organization = organization

        return stub.ListApigeeEnvironment(request).items

    def to_proto(self):
        resource = environment_pb2.ApigeeEnvironment()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if EnvironmentProperties.to_proto(self.properties):
            resource.properties.CopyFrom(
                EnvironmentProperties.to_proto(self.properties)
            )
        else:
            resource.ClearField("properties")
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.organization):
            resource.organization = Primitive.to_proto(self.organization)
        return resource


class EnvironmentProperties(object):
    def __init__(self, property: list = None):
        self.property = property

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = environment_pb2.ApigeeEnvironmentProperties()
        if EnvironmentPropertiesPropertyArray.to_proto(resource.property):
            res.property.extend(
                EnvironmentPropertiesPropertyArray.to_proto(resource.property)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EnvironmentProperties(property=resource.property,)


class EnvironmentPropertiesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EnvironmentProperties.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EnvironmentProperties.from_proto(i) for i in resources]


class EnvironmentPropertiesProperty(object):
    def __init__(self, name: str = None, value: str = None):
        self.name = name
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = environment_pb2.ApigeeEnvironmentPropertiesProperty()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return EnvironmentPropertiesProperty(name=resource.name, value=resource.value,)


class EnvironmentPropertiesPropertyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [EnvironmentPropertiesProperty.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [EnvironmentPropertiesProperty.from_proto(i) for i in resources]


class EnvironmentStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return environment_pb2.ApigeeEnvironmentStateEnum.Value(
            "ApigeeEnvironmentStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return environment_pb2.ApigeeEnvironmentStateEnum.Name(resource)[
            len("ApigeeEnvironmentStateEnum") :
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
