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
from google3.cloud.graphite.mmv2.services.google.cloudbilling import (
    project_billing_info_pb2,
)
from google3.cloud.graphite.mmv2.services.google.cloudbilling import (
    project_billing_info_pb2_grpc,
)

from typing import List


class ProjectBillingInfo(object):
    def __init__(
        self,
        name: str = None,
        billing_account_name: str = None,
        billing_enabled: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.billing_account_name = billing_account_name
        self.service_account_file = service_account_file

    def apply(self):
        stub = project_billing_info_pb2_grpc.CloudbillingProjectBillingInfoServiceStub(
            channel.Channel()
        )
        request = project_billing_info_pb2.ApplyCloudbillingProjectBillingInfoRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.billing_account_name):
            request.resource.billing_account_name = Primitive.to_proto(
                self.billing_account_name
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyCloudbillingProjectBillingInfo(request)
        self.name = Primitive.from_proto(response.name)
        self.billing_account_name = Primitive.from_proto(response.billing_account_name)
        self.billing_enabled = Primitive.from_proto(response.billing_enabled)

    @classmethod
    def delete(self, name, service_account_file=""):
        stub = project_billing_info_pb2_grpc.CloudbillingProjectBillingInfoServiceStub(
            channel.Channel()
        )
        request = project_billing_info_pb2.DeleteCloudbillingProjectBillingInfoRequest()
        request.service_account_file = service_account_file
        request.Name = name

        response = stub.DeleteCloudbillingProjectBillingInfo(request)

    @classmethod
    def list(self, name, service_account_file=""):
        stub = project_billing_info_pb2_grpc.CloudbillingProjectBillingInfoServiceStub(
            channel.Channel()
        )
        request = project_billing_info_pb2.ListCloudbillingProjectBillingInfoRequest()
        request.service_account_file = service_account_file
        request.Name = name

        return stub.ListCloudbillingProjectBillingInfo(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = project_billing_info_pb2.CloudbillingProjectBillingInfo()
        any_proto.Unpack(res_proto)

        res = ProjectBillingInfo()
        res.name = Primitive.from_proto(res_proto.name)
        res.billing_account_name = Primitive.from_proto(res_proto.billing_account_name)
        res.billing_enabled = Primitive.from_proto(res_proto.billing_enabled)
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
