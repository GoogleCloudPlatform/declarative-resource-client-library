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
from google3.cloud.graphite.mmv2.services.google.runtimeconfig import variable_pb2
from google3.cloud.graphite.mmv2.services.google.runtimeconfig import variable_pb2_grpc

from typing import List


class Variable(object):
    def __init__(
        self,
        name: str = None,
        runtime_config: str = None,
        text: str = None,
        value: str = None,
        update_time: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.runtime_config = runtime_config
        self.text = text
        self.value = value
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = variable_pb2_grpc.RuntimeconfigVariableServiceStub(channel.Channel())
        request = variable_pb2.ApplyRuntimeconfigVariableRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.runtime_config):
            request.resource.runtime_config = Primitive.to_proto(self.runtime_config)

        if Primitive.to_proto(self.text):
            request.resource.text = Primitive.to_proto(self.text)

        if Primitive.to_proto(self.value):
            request.resource.value = Primitive.to_proto(self.value)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyRuntimeconfigVariable(request)
        self.name = Primitive.from_proto(response.name)
        self.runtime_config = Primitive.from_proto(response.runtime_config)
        self.text = Primitive.from_proto(response.text)
        self.value = Primitive.from_proto(response.value)
        self.update_time = Primitive.from_proto(response.update_time)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = variable_pb2_grpc.RuntimeconfigVariableServiceStub(channel.Channel())
        request = variable_pb2.DeleteRuntimeconfigVariableRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.runtime_config):
            request.resource.runtime_config = Primitive.to_proto(self.runtime_config)

        if Primitive.to_proto(self.text):
            request.resource.text = Primitive.to_proto(self.text)

        if Primitive.to_proto(self.value):
            request.resource.value = Primitive.to_proto(self.value)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteRuntimeconfigVariable(request)

    @classmethod
    def list(self, project, runtimeConfig, service_account_file=""):
        stub = variable_pb2_grpc.RuntimeconfigVariableServiceStub(channel.Channel())
        request = variable_pb2.ListRuntimeconfigVariableRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.RuntimeConfig = runtimeConfig

        return stub.ListRuntimeconfigVariable(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = variable_pb2.RuntimeconfigVariable()
        any_proto.Unpack(res_proto)

        res = Variable()
        res.name = Primitive.from_proto(res_proto.name)
        res.runtime_config = Primitive.from_proto(res_proto.runtime_config)
        res.text = Primitive.from_proto(res_proto.text)
        res.value = Primitive.from_proto(res_proto.value)
        res.update_time = Primitive.from_proto(res_proto.update_time)
        res.project = Primitive.from_proto(res_proto.project)
        return res

    def to_proto(self):
        resource = variable_pb2.RuntimeconfigVariable()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.runtime_config):
            resource.runtime_config = Primitive.to_proto(self.runtime_config)
        if Primitive.to_proto(self.text):
            resource.text = Primitive.to_proto(self.text)
        if Primitive.to_proto(self.value):
            resource.value = Primitive.to_proto(self.value)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
