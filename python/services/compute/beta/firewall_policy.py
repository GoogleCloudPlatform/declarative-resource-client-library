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
from google3.cloud.graphite.mmv2.services.google.compute import firewall_policy_pb2
from google3.cloud.graphite.mmv2.services.google.compute import firewall_policy_pb2_grpc

from typing import List


class FirewallPolicy(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        fingerprint: str = None,
        self_link: str = None,
        self_link_with_id: str = None,
        rule_tuple_count: int = None,
        display_name: str = None,
        parent: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.fingerprint = fingerprint
        self.self_link = self_link
        self.self_link_with_id = self_link_with_id
        self.display_name = display_name
        self.parent = parent
        self.service_account_file = service_account_file

    def apply(self):
        stub = firewall_policy_pb2_grpc.ComputeBetaFirewallPolicyServiceStub(
            channel.Channel()
        )
        request = firewall_policy_pb2.ApplyComputeBetaFirewallPolicyRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.fingerprint):
            request.resource.fingerprint = Primitive.to_proto(self.fingerprint)

        if Primitive.to_proto(self.self_link):
            request.resource.self_link = Primitive.to_proto(self.self_link)

        if Primitive.to_proto(self.self_link_with_id):
            request.resource.self_link_with_id = Primitive.to_proto(
                self.self_link_with_id
            )

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeBetaFirewallPolicy(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.fingerprint = Primitive.from_proto(response.fingerprint)
        self.self_link = Primitive.from_proto(response.self_link)
        self.self_link_with_id = Primitive.from_proto(response.self_link_with_id)
        self.rule_tuple_count = Primitive.from_proto(response.rule_tuple_count)
        self.display_name = Primitive.from_proto(response.display_name)
        self.parent = Primitive.from_proto(response.parent)

    def delete(self):
        stub = firewall_policy_pb2_grpc.ComputeBetaFirewallPolicyServiceStub(
            channel.Channel()
        )
        request = firewall_policy_pb2.DeleteComputeBetaFirewallPolicyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.fingerprint):
            request.resource.fingerprint = Primitive.to_proto(self.fingerprint)

        if Primitive.to_proto(self.self_link):
            request.resource.self_link = Primitive.to_proto(self.self_link)

        if Primitive.to_proto(self.self_link_with_id):
            request.resource.self_link_with_id = Primitive.to_proto(
                self.self_link_with_id
            )

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        response = stub.DeleteComputeBetaFirewallPolicy(request)

    @classmethod
    def list(self, parent, service_account_file=""):
        stub = firewall_policy_pb2_grpc.ComputeBetaFirewallPolicyServiceStub(
            channel.Channel()
        )
        request = firewall_policy_pb2.ListComputeBetaFirewallPolicyRequest()
        request.service_account_file = service_account_file
        request.Parent = parent

        return stub.ListComputeBetaFirewallPolicy(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = firewall_policy_pb2.ComputeBetaFirewallPolicy()
        any_proto.Unpack(res_proto)

        res = FirewallPolicy()
        res.name = Primitive.from_proto(res_proto.name)
        res.description = Primitive.from_proto(res_proto.description)
        res.fingerprint = Primitive.from_proto(res_proto.fingerprint)
        res.self_link = Primitive.from_proto(res_proto.self_link)
        res.self_link_with_id = Primitive.from_proto(res_proto.self_link_with_id)
        res.rule_tuple_count = Primitive.from_proto(res_proto.rule_tuple_count)
        res.display_name = Primitive.from_proto(res_proto.display_name)
        res.parent = Primitive.from_proto(res_proto.parent)
        return res

    def to_proto(self):
        resource = firewall_policy_pb2.ComputeBetaFirewallPolicy()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.fingerprint):
            resource.fingerprint = Primitive.to_proto(self.fingerprint)
        if Primitive.to_proto(self.self_link):
            resource.self_link = Primitive.to_proto(self.self_link)
        if Primitive.to_proto(self.self_link_with_id):
            resource.self_link_with_id = Primitive.to_proto(self.self_link_with_id)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
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
