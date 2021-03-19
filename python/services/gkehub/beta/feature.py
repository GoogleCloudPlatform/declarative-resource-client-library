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
from google3.cloud.graphite.mmv2.services.google.gkehub import feature_pb2
from google3.cloud.graphite.mmv2.services.google.gkehub import feature_pb2_grpc

from typing import List


class Feature(object):
    def __init__(
        self,
        name: str = None,
        labels: dict = None,
        resource_state: dict = None,
        spec: dict = None,
        membership_specs: dict = None,
        state: dict = None,
        membership_states: dict = None,
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
        self.membership_specs = membership_specs
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
        if Primitive.to_proto(self.membership_specs):
            request.resource.membership_specs = Primitive.to_proto(
                self.membership_specs
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyGkehubBetaFeature(request)
        self.name = Primitive.from_proto(response.name)
        self.labels = Primitive.from_proto(response.labels)
        self.resource_state = FeatureResourceState.from_proto(response.resource_state)
        self.spec = FeatureSpec.from_proto(response.spec)
        self.membership_specs = Primitive.from_proto(response.membership_specs)
        self.state = FeatureState.from_proto(response.state)
        self.membership_states = Primitive.from_proto(response.membership_states)
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
        if Primitive.to_proto(self.membership_specs):
            request.resource.membership_specs = Primitive.to_proto(
                self.membership_specs
            )

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

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = feature_pb2.GkehubBetaFeature()
        any_proto.Unpack(res_proto)

        res = Feature()
        res.name = Primitive.from_proto(res_proto.name)
        res.labels = Primitive.from_proto(res_proto.labels)
        res.resource_state = FeatureResourceState.from_proto(res_proto.resource_state)
        res.spec = FeatureSpec.from_proto(res_proto.spec)
        res.membership_specs = Primitive.from_proto(res_proto.membership_specs)
        res.state = FeatureState.from_proto(res_proto.state)
        res.membership_states = Primitive.from_proto(res_proto.membership_states)
        res.create_time = Primitive.from_proto(res_proto.create_time)
        res.update_time = Primitive.from_proto(res_proto.update_time)
        res.delete_time = Primitive.from_proto(res_proto.delete_time)
        res.project = Primitive.from_proto(res_proto.project)
        res.location = Primitive.from_proto(res_proto.location)
        return res

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
        if Primitive.to_proto(self.membership_specs):
            resource.membership_specs = Primitive.to_proto(self.membership_specs)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class FeatureResourceState(object):
    def __init__(self, state: str = None):
        self.state = state

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubBetaFeatureResourceState()
        if FeatureResourceStateStateEnum.to_proto(resource.state):
            res.state = FeatureResourceStateStateEnum.to_proto(resource.state)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureResourceState(state=resource.state,)


class FeatureResourceStateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureResourceState.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureResourceState.from_proto(i) for i in resources]


class FeatureSpec(object):
    def __init__(self, multiclusteringress: dict = None, helloworld: dict = None):
        self.multiclusteringress = multiclusteringress
        self.helloworld = helloworld

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
        if FeatureSpecHelloworld.to_proto(resource.helloworld):
            res.helloworld.CopyFrom(FeatureSpecHelloworld.to_proto(resource.helloworld))
        else:
            res.ClearField("helloworld")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureSpec(
            multiclusteringress=resource.multiclusteringress,
            helloworld=resource.helloworld,
        )


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


class FeatureSpecHelloworld(object):
    def __init__(self, feature_test: dict = None, custom_config: str = None):
        self.feature_test = feature_test
        self.custom_config = custom_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubBetaFeatureSpecHelloworld()
        if FeatureSpecHelloworldFeatureTest.to_proto(resource.feature_test):
            res.feature_test.CopyFrom(
                FeatureSpecHelloworldFeatureTest.to_proto(resource.feature_test)
            )
        else:
            res.ClearField("feature_test")
        if Primitive.to_proto(resource.custom_config):
            res.custom_config = Primitive.to_proto(resource.custom_config)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureSpecHelloworld(
            feature_test=resource.feature_test, custom_config=resource.custom_config,
        )


class FeatureSpecHelloworldArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureSpecHelloworld.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureSpecHelloworld.from_proto(i) for i in resources]


class FeatureSpecHelloworldFeatureTest(object):
    def __init__(
        self,
        first: str = None,
        second: int = None,
        third: str = None,
        fourth: str = None,
        fifth: dict = None,
        sixth: int = None,
        seventh: str = None,
        eighth: list = None,
        ninth: dict = None,
    ):
        self.first = first
        self.second = second
        self.third = third
        self.fourth = fourth
        self.fifth = fifth
        self.sixth = sixth
        self.seventh = seventh
        self.eighth = eighth
        self.ninth = ninth

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubBetaFeatureSpecHelloworldFeatureTest()
        if Primitive.to_proto(resource.first):
            res.first = Primitive.to_proto(resource.first)
        if Primitive.to_proto(resource.second):
            res.second = Primitive.to_proto(resource.second)
        if FeatureSpecHelloworldFeatureTestThirdEnum.to_proto(resource.third):
            res.third = FeatureSpecHelloworldFeatureTestThirdEnum.to_proto(
                resource.third
            )
        if Primitive.to_proto(resource.fourth):
            res.fourth = Primitive.to_proto(resource.fourth)
        if FeatureSpecHelloworldFeatureTestFifth.to_proto(resource.fifth):
            res.fifth.CopyFrom(
                FeatureSpecHelloworldFeatureTestFifth.to_proto(resource.fifth)
            )
        else:
            res.ClearField("fifth")
        if Primitive.to_proto(resource.sixth):
            res.sixth = Primitive.to_proto(resource.sixth)
        if Primitive.to_proto(resource.seventh):
            res.seventh = Primitive.to_proto(resource.seventh)
        if FeatureSpecHelloworldFeatureTestEighthArray.to_proto(resource.eighth):
            res.eighth.extend(
                FeatureSpecHelloworldFeatureTestEighthArray.to_proto(resource.eighth)
            )
        if Primitive.to_proto(resource.ninth):
            res.ninth = Primitive.to_proto(resource.ninth)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureSpecHelloworldFeatureTest(
            first=resource.first,
            second=resource.second,
            third=resource.third,
            fourth=resource.fourth,
            fifth=resource.fifth,
            sixth=resource.sixth,
            seventh=resource.seventh,
            eighth=resource.eighth,
            ninth=resource.ninth,
        )


class FeatureSpecHelloworldFeatureTestArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureSpecHelloworldFeatureTest.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureSpecHelloworldFeatureTest.from_proto(i) for i in resources]


class FeatureSpecHelloworldFeatureTestFifth(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubBetaFeatureSpecHelloworldFeatureTestFifth()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureSpecHelloworldFeatureTestFifth(
            type_url=resource.type_url, value=resource.value,
        )


class FeatureSpecHelloworldFeatureTestFifthArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureSpecHelloworldFeatureTestFifth.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureSpecHelloworldFeatureTestFifth.from_proto(i) for i in resources]


class FeatureSpecHelloworldFeatureTestEighth(object):
    def __init__(self, first: str = None, second: int = None):
        self.first = first
        self.second = second

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubBetaFeatureSpecHelloworldFeatureTestEighth()
        if Primitive.to_proto(resource.first):
            res.first = Primitive.to_proto(resource.first)
        if Primitive.to_proto(resource.second):
            res.second = Primitive.to_proto(resource.second)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureSpecHelloworldFeatureTestEighth(
            first=resource.first, second=resource.second,
        )


class FeatureSpecHelloworldFeatureTestEighthArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureSpecHelloworldFeatureTestEighth.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureSpecHelloworldFeatureTestEighth.from_proto(i) for i in resources]


class FeatureState(object):
    def __init__(self, state: dict = None, helloworld: dict = None):
        self.state = state
        self.helloworld = helloworld

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubBetaFeatureState()
        if FeatureStateState.to_proto(resource.state):
            res.state.CopyFrom(FeatureStateState.to_proto(resource.state))
        else:
            res.ClearField("state")
        if FeatureStateHelloworld.to_proto(resource.helloworld):
            res.helloworld.CopyFrom(
                FeatureStateHelloworld.to_proto(resource.helloworld)
            )
        else:
            res.ClearField("helloworld")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureState(state=resource.state, helloworld=resource.helloworld,)


class FeatureStateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureState.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureState.from_proto(i) for i in resources]


class FeatureStateState(object):
    def __init__(
        self, code: str = None, description: str = None, update_time: str = None
    ):
        self.code = code
        self.description = description
        self.update_time = update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubBetaFeatureStateState()
        if FeatureStateStateCodeEnum.to_proto(resource.code):
            res.code = FeatureStateStateCodeEnum.to_proto(resource.code)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureStateState(
            code=resource.code,
            description=resource.description,
            update_time=resource.update_time,
        )


class FeatureStateStateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureStateState.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureStateState.from_proto(i) for i in resources]


class FeatureStateHelloworld(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_pb2.GkehubBetaFeatureStateHelloworld()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureStateHelloworld()


class FeatureStateHelloworldArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureStateHelloworld.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureStateHelloworld.from_proto(i) for i in resources]


class FeatureResourceStateStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return feature_pb2.GkehubBetaFeatureResourceStateStateEnum.Value(
            "GkehubBetaFeatureResourceStateStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return feature_pb2.GkehubBetaFeatureResourceStateStateEnum.Name(resource)[
            len("GkehubBetaFeatureResourceStateStateEnum") :
        ]


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


class FeatureSpecHelloworldFeatureTestThirdEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return feature_pb2.GkehubBetaFeatureSpecHelloworldFeatureTestThirdEnum.Value(
            "GkehubBetaFeatureSpecHelloworldFeatureTestThirdEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return feature_pb2.GkehubBetaFeatureSpecHelloworldFeatureTestThirdEnum.Name(
            resource
        )[len("GkehubBetaFeatureSpecHelloworldFeatureTestThirdEnum") :]


class FeatureStateStateCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return feature_pb2.GkehubBetaFeatureStateStateCodeEnum.Value(
            "GkehubBetaFeatureStateStateCodeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return feature_pb2.GkehubBetaFeatureStateStateCodeEnum.Name(resource)[
            len("GkehubBetaFeatureStateStateCodeEnum") :
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
