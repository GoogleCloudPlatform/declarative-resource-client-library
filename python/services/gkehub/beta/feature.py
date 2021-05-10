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
from google3.cloud.graphite.mmv2.services.google.gke_hub import feature_pb2
from google3.cloud.graphite.mmv2.services.google.gke_hub import feature_pb2_grpc

from typing import List


class Feature(object):
    def __init__(
        self,
        name: str = None,
        labels: dict = None,
        spec: dict = None,
        create_time: str = None,
        update_time: str = None,
        delete_time: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.labels = labels
        self.spec = spec
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = feature_pb2_grpc.GkehubBetaFeatureServiceStub(channel.Channel())
        request = feature_pb2.ApplyGkehubBetaFeatureRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if FeatureSpec.to_proto(self.spec):
            request.resource.spec.CopyFrom(FeatureSpec.to_proto(self.spec))
        else:
            request.resource.ClearField("spec")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyGkehubBetaFeature(request)
        self.name = Primitive.from_proto(response.name)
        self.labels = Primitive.from_proto(response.labels)
        self.spec = FeatureSpec.from_proto(response.spec)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.delete_time = Primitive.from_proto(response.delete_time)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = feature_pb2_grpc.GkehubBetaFeatureServiceStub(channel.Channel())
        request = feature_pb2.DeleteGkehubBetaFeatureRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if FeatureSpec.to_proto(self.spec):
            request.resource.spec.CopyFrom(FeatureSpec.to_proto(self.spec))
        else:
            request.resource.ClearField("spec")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteGkehubBetaFeature(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = feature_pb2_grpc.GkehubBetaFeatureServiceStub(channel.Channel())
        request = feature_pb2.ListGkehubBetaFeatureRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListGkehubBetaFeature(request).items

    def to_proto(self):
        resource = feature_pb2.GkehubBetaFeature()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if FeatureSpec.to_proto(self.spec):
            resource.spec.CopyFrom(FeatureSpec.to_proto(self.spec))
        else:
            resource.ClearField("spec")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class FeatureSpec(object):
    def __init__(self, multiclusteringress: dict = None):
        self.multiclusteringress = multiclusteringress

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubBetaFeatureSpec()
        if FeatureSpecMulticlusteringress.to_proto(resource.multiclusteringress):
            res.multiclusteringress.CopyFrom(
                FeatureSpecMulticlusteringress.to_proto(resource.multiclusteringress)
            )
        else:
            res.ClearField("multiclusteringress")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureSpec(multiclusteringress=resource.multiclusteringress,)


class FeatureSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureSpec.from_proto(i) for i in resources]


class FeatureSpecMulticlusteringress(object):
    def __init__(self, config_membership: str = None, billing: str = None):
        self.config_membership = config_membership
        self.billing = billing

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubBetaFeatureSpecMulticlusteringress()
        if Primitive.to_proto(resource.config_membership):
            res.config_membership = Primitive.to_proto(resource.config_membership)
        if FeatureSpecMulticlusteringressBillingEnum.to_proto(resource.billing):
            res.billing = FeatureSpecMulticlusteringressBillingEnum.to_proto(
                resource.billing
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureSpecMulticlusteringress(
            config_membership=resource.config_membership, billing=resource.billing,
        )


class FeatureSpecMulticlusteringressArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureSpecMulticlusteringress.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureSpecMulticlusteringress.from_proto(i) for i in resources]


class FeatureSpecMulticlusteringressBillingEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return feature_pb2.GkehubBetaFeatureSpecMulticlusteringressBillingEnum.Value(
            "GkehubBetaFeatureSpecMulticlusteringressBillingEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return feature_pb2.GkehubBetaFeatureSpecMulticlusteringressBillingEnum.Name(
            resource
        )[len("GkehubBetaFeatureSpecMulticlusteringressBillingEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
