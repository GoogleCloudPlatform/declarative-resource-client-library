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
from google3.cloud.graphite.mmv2.services.google.runtimeconfig import config_pb2
from google3.cloud.graphite.mmv2.services.google.runtimeconfig import config_pb2_grpc

from typing import List


class Config(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = config_pb2_grpc.RuntimeconfigConfigServiceStub(channel.Channel())
        request = config_pb2.ApplyRuntimeconfigConfigRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyRuntimeconfigConfig(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.project = Primitive.from_proto(response.project)

    @classmethod
    def delete(self, project, name, service_account_file=""):
        stub = config_pb2_grpc.RuntimeconfigConfigServiceStub(channel.Channel())
        request = config_pb2.DeleteRuntimeconfigConfigRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Name = name

        response = stub.DeleteRuntimeconfigConfig(request)

    @classmethod
    def list(self, project, name, service_account_file=""):
        stub = config_pb2_grpc.RuntimeconfigConfigServiceStub(channel.Channel())
        request = config_pb2.ListRuntimeconfigConfigRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Name = name

        return stub.ListRuntimeconfigConfig(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = config_pb2.RuntimeconfigConfig()
        any_proto.Unpack(res_proto)

        res = Config()
        res.name = Primitive.from_proto(res_proto.name)
        res.description = Primitive.from_proto(res_proto.description)
        res.project = Primitive.from_proto(res_proto.project)
        return res


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
