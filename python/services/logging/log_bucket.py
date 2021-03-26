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
from google3.cloud.graphite.mmv2.services.google.logging import log_bucket_pb2
from google3.cloud.graphite.mmv2.services.google.logging import log_bucket_pb2_grpc

from typing import List


class LogBucket(object):
    def __init__(
        self,
        self_link: str = None,
        description: str = None,
        create_time: str = None,
        update_time: str = None,
        retention_days: int = None,
        locked: bool = None,
        lifecycle_state: str = None,
        parent: str = None,
        location: str = None,
        name: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.description = description
        self.retention_days = retention_days
        self.locked = locked
        self.parent = parent
        self.location = location
        self.name = name
        self.service_account_file = service_account_file

    def apply(self):
        stub = log_bucket_pb2_grpc.LoggingLogBucketServiceStub(channel.Channel())
        request = log_bucket_pb2.ApplyLoggingLogBucketRequest()
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.retention_days):
            request.resource.retention_days = Primitive.to_proto(self.retention_days)

        if Primitive.to_proto(self.locked):
            request.resource.locked = Primitive.to_proto(self.locked)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        request.service_account_file = self.service_account_file

        response = stub.ApplyLoggingLogBucket(request)
        self.self_link = Primitive.from_proto(response.self_link)
        self.description = Primitive.from_proto(response.description)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.retention_days = Primitive.from_proto(response.retention_days)
        self.locked = Primitive.from_proto(response.locked)
        self.lifecycle_state = LogBucketLifecycleStateEnum.from_proto(
            response.lifecycle_state
        )
        self.parent = Primitive.from_proto(response.parent)
        self.location = Primitive.from_proto(response.location)
        self.name = Primitive.from_proto(response.name)

    def delete(self):
        stub = log_bucket_pb2_grpc.LoggingLogBucketServiceStub(channel.Channel())
        request = log_bucket_pb2.DeleteLoggingLogBucketRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.retention_days):
            request.resource.retention_days = Primitive.to_proto(self.retention_days)

        if Primitive.to_proto(self.locked):
            request.resource.locked = Primitive.to_proto(self.locked)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        response = stub.DeleteLoggingLogBucket(request)

    @classmethod
    def list(self, location, parent, service_account_file=""):
        stub = log_bucket_pb2_grpc.LoggingLogBucketServiceStub(channel.Channel())
        request = log_bucket_pb2.ListLoggingLogBucketRequest()
        request.service_account_file = service_account_file
        request.Location = location

        request.Parent = parent

        return stub.ListLoggingLogBucket(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = log_bucket_pb2.LoggingLogBucket()
        any_proto.Unpack(res_proto)

        res = LogBucket()
        res.self_link = Primitive.from_proto(res_proto.self_link)
        res.description = Primitive.from_proto(res_proto.description)
        res.create_time = Primitive.from_proto(res_proto.create_time)
        res.update_time = Primitive.from_proto(res_proto.update_time)
        res.retention_days = Primitive.from_proto(res_proto.retention_days)
        res.locked = Primitive.from_proto(res_proto.locked)
        res.lifecycle_state = LogBucketLifecycleStateEnum.from_proto(
            res_proto.lifecycle_state
        )
        res.parent = Primitive.from_proto(res_proto.parent)
        res.location = Primitive.from_proto(res_proto.location)
        res.name = Primitive.from_proto(res_proto.name)
        return res

    def to_proto(self):
        resource = log_bucket_pb2.LoggingLogBucket()
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.retention_days):
            resource.retention_days = Primitive.to_proto(self.retention_days)
        if Primitive.to_proto(self.locked):
            resource.locked = Primitive.to_proto(self.locked)
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        return resource


class LogBucketLifecycleStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return log_bucket_pb2.LoggingLogBucketLifecycleStateEnum.Value(
            "LoggingLogBucketLifecycleStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return log_bucket_pb2.LoggingLogBucketLifecycleStateEnum.Name(resource)[
            len("LoggingLogBucketLifecycleStateEnum") :
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
