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
from google3.cloud.graphite.mmv2.services.google.cloud_build import worker_pool_pb2
from google3.cloud.graphite.mmv2.services.google.cloud_build import worker_pool_pb2_grpc

from typing import List


class WorkerPool(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        uid: str = None,
        annotations: dict = None,
        create_time: str = None,
        update_time: str = None,
        delete_time: str = None,
        state: str = None,
        etag: str = None,
        worker_config: dict = None,
        network_config: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.annotations = annotations
        self.worker_config = worker_config
        self.network_config = network_config
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = worker_pool_pb2_grpc.CloudbuildAlphaWorkerPoolServiceStub(
            channel.Channel()
        )
        request = worker_pool_pb2.ApplyCloudbuildAlphaWorkerPoolRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if WorkerPoolWorkerConfig.to_proto(self.worker_config):
            request.resource.worker_config.CopyFrom(
                WorkerPoolWorkerConfig.to_proto(self.worker_config)
            )
        else:
            request.resource.ClearField("worker_config")
        if WorkerPoolNetworkConfig.to_proto(self.network_config):
            request.resource.network_config.CopyFrom(
                WorkerPoolNetworkConfig.to_proto(self.network_config)
            )
        else:
            request.resource.ClearField("network_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudbuildAlphaWorkerPool(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.uid = Primitive.from_proto(response.uid)
        self.annotations = Primitive.from_proto(response.annotations)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.delete_time = Primitive.from_proto(response.delete_time)
        self.state = WorkerPoolStateEnum.from_proto(response.state)
        self.etag = Primitive.from_proto(response.etag)
        self.worker_config = WorkerPoolWorkerConfig.from_proto(response.worker_config)
        self.network_config = WorkerPoolNetworkConfig.from_proto(
            response.network_config
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = worker_pool_pb2_grpc.CloudbuildAlphaWorkerPoolServiceStub(
            channel.Channel()
        )
        request = worker_pool_pb2.DeleteCloudbuildAlphaWorkerPoolRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if WorkerPoolWorkerConfig.to_proto(self.worker_config):
            request.resource.worker_config.CopyFrom(
                WorkerPoolWorkerConfig.to_proto(self.worker_config)
            )
        else:
            request.resource.ClearField("worker_config")
        if WorkerPoolNetworkConfig.to_proto(self.network_config):
            request.resource.network_config.CopyFrom(
                WorkerPoolNetworkConfig.to_proto(self.network_config)
            )
        else:
            request.resource.ClearField("network_config")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteCloudbuildAlphaWorkerPool(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = worker_pool_pb2_grpc.CloudbuildAlphaWorkerPoolServiceStub(
            channel.Channel()
        )
        request = worker_pool_pb2.ListCloudbuildAlphaWorkerPoolRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListCloudbuildAlphaWorkerPool(request).items

    def to_proto(self):
        resource = worker_pool_pb2.CloudbuildAlphaWorkerPool()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if WorkerPoolWorkerConfig.to_proto(self.worker_config):
            resource.worker_config.CopyFrom(
                WorkerPoolWorkerConfig.to_proto(self.worker_config)
            )
        else:
            resource.ClearField("worker_config")
        if WorkerPoolNetworkConfig.to_proto(self.network_config):
            resource.network_config.CopyFrom(
                WorkerPoolNetworkConfig.to_proto(self.network_config)
            )
        else:
            resource.ClearField("network_config")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class WorkerPoolWorkerConfig(object):
    def __init__(
        self,
        machine_type: str = None,
        disk_size_gb: int = None,
        no_external_ip: bool = None,
    ):
        self.machine_type = machine_type
        self.disk_size_gb = disk_size_gb
        self.no_external_ip = no_external_ip

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = worker_pool_pb2.CloudbuildAlphaWorkerPoolWorkerConfig()
        if Primitive.to_proto(resource.machine_type):
            res.machine_type = Primitive.to_proto(resource.machine_type)
        if Primitive.to_proto(resource.disk_size_gb):
            res.disk_size_gb = Primitive.to_proto(resource.disk_size_gb)
        if Primitive.to_proto(resource.no_external_ip):
            res.no_external_ip = Primitive.to_proto(resource.no_external_ip)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkerPoolWorkerConfig(
            machine_type=Primitive.from_proto(resource.machine_type),
            disk_size_gb=Primitive.from_proto(resource.disk_size_gb),
            no_external_ip=Primitive.from_proto(resource.no_external_ip),
        )


class WorkerPoolWorkerConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkerPoolWorkerConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkerPoolWorkerConfig.from_proto(i) for i in resources]


class WorkerPoolNetworkConfig(object):
    def __init__(self, peered_network: str = None):
        self.peered_network = peered_network

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = worker_pool_pb2.CloudbuildAlphaWorkerPoolNetworkConfig()
        if Primitive.to_proto(resource.peered_network):
            res.peered_network = Primitive.to_proto(resource.peered_network)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkerPoolNetworkConfig(
            peered_network=Primitive.from_proto(resource.peered_network),
        )


class WorkerPoolNetworkConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkerPoolNetworkConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkerPoolNetworkConfig.from_proto(i) for i in resources]


class WorkerPoolStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return worker_pool_pb2.CloudbuildAlphaWorkerPoolStateEnum.Value(
            "CloudbuildAlphaWorkerPoolStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return worker_pool_pb2.CloudbuildAlphaWorkerPoolStateEnum.Name(resource)[
            len("CloudbuildAlphaWorkerPoolStateEnum") :
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
