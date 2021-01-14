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
from google3.cloud.graphite.mmv2.services.google.dns import resource_record_set_pb2
from google3.cloud.graphite.mmv2.services.google.dns import resource_record_set_pb2_grpc

from typing import List


class ResourceRecordSet(object):
    def __init__(
        self,
        dns_name: str = None,
        dns_type: str = None,
        ttl: int = None,
        target: list = None,
        managed_zone: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.dns_name = dns_name
        self.dns_type = dns_type
        self.ttl = ttl
        self.target = target
        self.managed_zone = managed_zone
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = resource_record_set_pb2_grpc.DnsResourceRecordSetServiceStub(
            channel.Channel()
        )
        request = resource_record_set_pb2.ApplyDnsResourceRecordSetRequest()
        if Primitive.to_proto(self.dns_name):
            request.resource.dns_name = Primitive.to_proto(self.dns_name)

        if Primitive.to_proto(self.dns_type):
            request.resource.dns_type = Primitive.to_proto(self.dns_type)

        if Primitive.to_proto(self.ttl):
            request.resource.ttl = Primitive.to_proto(self.ttl)

        if Primitive.to_proto(self.target):
            request.resource.target.extend(Primitive.to_proto(self.target))
        if Primitive.to_proto(self.managed_zone):
            request.resource.managed_zone = Primitive.to_proto(self.managed_zone)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDnsResourceRecordSet(request)
        self.dns_name = Primitive.from_proto(response.dns_name)
        self.dns_type = Primitive.from_proto(response.dns_type)
        self.ttl = Primitive.from_proto(response.ttl)
        self.target = Primitive.from_proto(response.target)
        self.managed_zone = Primitive.from_proto(response.managed_zone)
        self.project = Primitive.from_proto(response.project)

    @classmethod
    def delete(self, project, dnsName, dnsType, managedZone, service_account_file=""):
        stub = resource_record_set_pb2_grpc.DnsResourceRecordSetServiceStub(
            channel.Channel()
        )
        request = resource_record_set_pb2.DeleteDnsResourceRecordSetRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.DnsName = dnsName

        request.DnsType = dnsType

        request.ManagedZone = managedZone

        response = stub.DeleteDnsResourceRecordSet(request)

    @classmethod
    def list(self, project, managedZone, service_account_file=""):
        stub = resource_record_set_pb2_grpc.DnsResourceRecordSetServiceStub(
            channel.Channel()
        )
        request = resource_record_set_pb2.ListDnsResourceRecordSetRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.ManagedZone = managedZone

        return stub.ListDnsResourceRecordSet(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = resource_record_set_pb2.DnsResourceRecordSet()
        any_proto.Unpack(res_proto)

        res = ResourceRecordSet()
        res.dns_name = Primitive.from_proto(res_proto.dns_name)
        res.dns_type = Primitive.from_proto(res_proto.dns_type)
        res.ttl = Primitive.from_proto(res_proto.ttl)
        res.target = Primitive.from_proto(res_proto.target)
        res.managed_zone = Primitive.from_proto(res_proto.managed_zone)
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
