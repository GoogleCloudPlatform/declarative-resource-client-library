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
from google3.cloud.graphite.mmv2.services.google.compute import target_pool_pb2
from google3.cloud.graphite.mmv2.services.google.compute import target_pool_pb2_grpc

from typing import List


class TargetPool(object):
    def __init__(
        self,
        backup_pool: str = None,
        description: str = None,
        failover_ratio: float = None,
        health_check: list = None,
        instance: list = None,
        name: str = None,
        region: str = None,
        self_link: str = None,
        session_affinity: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.backup_pool = backup_pool
        self.description = description
        self.failover_ratio = failover_ratio
        self.health_check = health_check
        self.instance = instance
        self.name = name
        self.region = region
        self.self_link = self_link
        self.session_affinity = session_affinity
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = target_pool_pb2_grpc.ComputeBetaTargetPoolServiceStub(channel.Channel())
        request = target_pool_pb2.ApplyComputeBetaTargetPoolRequest()
        if Primitive.to_proto(self.backup_pool):
            request.resource.backup_pool = Primitive.to_proto(self.backup_pool)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.failover_ratio):
            request.resource.failover_ratio = Primitive.to_proto(self.failover_ratio)

        if Primitive.to_proto(self.health_check):
            request.resource.health_check.extend(Primitive.to_proto(self.health_check))
        if Primitive.to_proto(self.instance):
            request.resource.instance.extend(Primitive.to_proto(self.instance))
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.region):
            request.resource.region = Primitive.to_proto(self.region)

        if Primitive.to_proto(self.self_link):
            request.resource.self_link = Primitive.to_proto(self.self_link)

        if TargetPoolSessionAffinityEnum.to_proto(self.session_affinity):
            request.resource.session_affinity = TargetPoolSessionAffinityEnum.to_proto(
                self.session_affinity
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaTargetPool(request)
        self.backup_pool = Primitive.from_proto(response.backup_pool)
        self.description = Primitive.from_proto(response.description)
        self.failover_ratio = Primitive.from_proto(response.failover_ratio)
        self.health_check = Primitive.from_proto(response.health_check)
        self.instance = Primitive.from_proto(response.instance)
        self.name = Primitive.from_proto(response.name)
        self.region = Primitive.from_proto(response.region)
        self.self_link = Primitive.from_proto(response.self_link)
        self.session_affinity = TargetPoolSessionAffinityEnum.from_proto(
            response.session_affinity
        )
        self.project = Primitive.from_proto(response.project)

    @classmethod
    def delete(self, project, region, name, service_account_file=""):
        stub = target_pool_pb2_grpc.ComputeBetaTargetPoolServiceStub(channel.Channel())
        request = target_pool_pb2.DeleteComputeBetaTargetPoolRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Region = region

        request.Name = name

        response = stub.DeleteComputeBetaTargetPool(request)

    @classmethod
    def list(self, project, region, service_account_file=""):
        stub = target_pool_pb2_grpc.ComputeBetaTargetPoolServiceStub(channel.Channel())
        request = target_pool_pb2.ListComputeBetaTargetPoolRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Region = region

        return stub.ListComputeBetaTargetPool(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = target_pool_pb2.ComputeBetaTargetPool()
        any_proto.Unpack(res_proto)

        res = TargetPool()
        res.backup_pool = Primitive.from_proto(res_proto.backup_pool)
        res.description = Primitive.from_proto(res_proto.description)
        res.failover_ratio = Primitive.from_proto(res_proto.failover_ratio)
        res.health_check = Primitive.from_proto(res_proto.health_check)
        res.instance = Primitive.from_proto(res_proto.instance)
        res.name = Primitive.from_proto(res_proto.name)
        res.region = Primitive.from_proto(res_proto.region)
        res.self_link = Primitive.from_proto(res_proto.self_link)
        res.session_affinity = TargetPoolSessionAffinityEnum.from_proto(
            res_proto.session_affinity
        )
        res.project = Primitive.from_proto(res_proto.project)
        return res


class TargetPoolSessionAffinityEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return target_pool_pb2.ComputeBetaTargetPoolSessionAffinityEnum.Value(
            "ComputeBetaTargetPoolSessionAffinityEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return target_pool_pb2.ComputeBetaTargetPoolSessionAffinityEnum.Name(resource)[
            len("ComputeBetaTargetPoolSessionAffinityEnum") :
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
