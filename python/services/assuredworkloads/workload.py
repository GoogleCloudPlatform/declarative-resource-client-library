# Copyright 2022 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.assured_workloads import workload_pb2
from google3.cloud.graphite.mmv2.services.google.assured_workloads import (
    workload_pb2_grpc,
)

from typing import List


class Workload(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        resources: list = None,
        compliance_regime: str = None,
        create_time: str = None,
        billing_account: str = None,
        labels: dict = None,
        provisioned_resources_parent: str = None,
        kms_settings: dict = None,
        resource_settings: list = None,
        organization: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.compliance_regime = compliance_regime
        self.billing_account = billing_account
        self.labels = labels
        self.provisioned_resources_parent = provisioned_resources_parent
        self.kms_settings = kms_settings
        self.resource_settings = resource_settings
        self.organization = organization
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = workload_pb2_grpc.AssuredworkloadsWorkloadServiceStub(channel.Channel())
        request = workload_pb2.ApplyAssuredworkloadsWorkloadRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if WorkloadComplianceRegimeEnum.to_proto(self.compliance_regime):
            request.resource.compliance_regime = WorkloadComplianceRegimeEnum.to_proto(
                self.compliance_regime
            )

        if Primitive.to_proto(self.billing_account):
            request.resource.billing_account = Primitive.to_proto(self.billing_account)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.provisioned_resources_parent):
            request.resource.provisioned_resources_parent = Primitive.to_proto(
                self.provisioned_resources_parent
            )

        if WorkloadKmsSettings.to_proto(self.kms_settings):
            request.resource.kms_settings.CopyFrom(
                WorkloadKmsSettings.to_proto(self.kms_settings)
            )
        else:
            request.resource.ClearField("kms_settings")
        if WorkloadResourceSettingsArray.to_proto(self.resource_settings):
            request.resource.resource_settings.extend(
                WorkloadResourceSettingsArray.to_proto(self.resource_settings)
            )
        if Primitive.to_proto(self.organization):
            request.resource.organization = Primitive.to_proto(self.organization)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyAssuredworkloadsWorkload(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.resources = WorkloadResourcesArray.from_proto(response.resources)
        self.compliance_regime = WorkloadComplianceRegimeEnum.from_proto(
            response.compliance_regime
        )
        self.create_time = Primitive.from_proto(response.create_time)
        self.billing_account = Primitive.from_proto(response.billing_account)
        self.labels = Primitive.from_proto(response.labels)
        self.provisioned_resources_parent = Primitive.from_proto(
            response.provisioned_resources_parent
        )
        self.kms_settings = WorkloadKmsSettings.from_proto(response.kms_settings)
        self.resource_settings = WorkloadResourceSettingsArray.from_proto(
            response.resource_settings
        )
        self.organization = Primitive.from_proto(response.organization)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = workload_pb2_grpc.AssuredworkloadsWorkloadServiceStub(channel.Channel())
        request = workload_pb2.DeleteAssuredworkloadsWorkloadRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if WorkloadComplianceRegimeEnum.to_proto(self.compliance_regime):
            request.resource.compliance_regime = WorkloadComplianceRegimeEnum.to_proto(
                self.compliance_regime
            )

        if Primitive.to_proto(self.billing_account):
            request.resource.billing_account = Primitive.to_proto(self.billing_account)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.provisioned_resources_parent):
            request.resource.provisioned_resources_parent = Primitive.to_proto(
                self.provisioned_resources_parent
            )

        if WorkloadKmsSettings.to_proto(self.kms_settings):
            request.resource.kms_settings.CopyFrom(
                WorkloadKmsSettings.to_proto(self.kms_settings)
            )
        else:
            request.resource.ClearField("kms_settings")
        if WorkloadResourceSettingsArray.to_proto(self.resource_settings):
            request.resource.resource_settings.extend(
                WorkloadResourceSettingsArray.to_proto(self.resource_settings)
            )
        if Primitive.to_proto(self.organization):
            request.resource.organization = Primitive.to_proto(self.organization)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteAssuredworkloadsWorkload(request)

    @classmethod
    def list(self, organization, location, service_account_file=""):
        stub = workload_pb2_grpc.AssuredworkloadsWorkloadServiceStub(channel.Channel())
        request = workload_pb2.ListAssuredworkloadsWorkloadRequest()
        request.service_account_file = service_account_file
        request.Organization = organization

        request.Location = location

        return stub.ListAssuredworkloadsWorkload(request).items

    def to_proto(self):
        resource = workload_pb2.AssuredworkloadsWorkload()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if WorkloadComplianceRegimeEnum.to_proto(self.compliance_regime):
            resource.compliance_regime = WorkloadComplianceRegimeEnum.to_proto(
                self.compliance_regime
            )
        if Primitive.to_proto(self.billing_account):
            resource.billing_account = Primitive.to_proto(self.billing_account)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.provisioned_resources_parent):
            resource.provisioned_resources_parent = Primitive.to_proto(
                self.provisioned_resources_parent
            )
        if WorkloadKmsSettings.to_proto(self.kms_settings):
            resource.kms_settings.CopyFrom(
                WorkloadKmsSettings.to_proto(self.kms_settings)
            )
        else:
            resource.ClearField("kms_settings")
        if WorkloadResourceSettingsArray.to_proto(self.resource_settings):
            resource.resource_settings.extend(
                WorkloadResourceSettingsArray.to_proto(self.resource_settings)
            )
        if Primitive.to_proto(self.organization):
            resource.organization = Primitive.to_proto(self.organization)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class WorkloadResources(object):
    def __init__(self, resource_id: int = None, resource_type: str = None):
        self.resource_id = resource_id
        self.resource_type = resource_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workload_pb2.AssuredworkloadsWorkloadResources()
        if Primitive.to_proto(resource.resource_id):
            res.resource_id = Primitive.to_proto(resource.resource_id)
        if WorkloadResourcesResourceTypeEnum.to_proto(resource.resource_type):
            res.resource_type = WorkloadResourcesResourceTypeEnum.to_proto(
                resource.resource_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkloadResources(
            resource_id=Primitive.from_proto(resource.resource_id),
            resource_type=WorkloadResourcesResourceTypeEnum.from_proto(
                resource.resource_type
            ),
        )


class WorkloadResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkloadResources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkloadResources.from_proto(i) for i in resources]


class WorkloadKmsSettings(object):
    def __init__(self, next_rotation_time: str = None, rotation_period: str = None):
        self.next_rotation_time = next_rotation_time
        self.rotation_period = rotation_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workload_pb2.AssuredworkloadsWorkloadKmsSettings()
        if Primitive.to_proto(resource.next_rotation_time):
            res.next_rotation_time = Primitive.to_proto(resource.next_rotation_time)
        if Primitive.to_proto(resource.rotation_period):
            res.rotation_period = Primitive.to_proto(resource.rotation_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkloadKmsSettings(
            next_rotation_time=Primitive.from_proto(resource.next_rotation_time),
            rotation_period=Primitive.from_proto(resource.rotation_period),
        )


class WorkloadKmsSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkloadKmsSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkloadKmsSettings.from_proto(i) for i in resources]


class WorkloadResourceSettings(object):
    def __init__(self, resource_id: str = None, resource_type: str = None):
        self.resource_id = resource_id
        self.resource_type = resource_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workload_pb2.AssuredworkloadsWorkloadResourceSettings()
        if Primitive.to_proto(resource.resource_id):
            res.resource_id = Primitive.to_proto(resource.resource_id)
        if WorkloadResourceSettingsResourceTypeEnum.to_proto(resource.resource_type):
            res.resource_type = WorkloadResourceSettingsResourceTypeEnum.to_proto(
                resource.resource_type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkloadResourceSettings(
            resource_id=Primitive.from_proto(resource.resource_id),
            resource_type=WorkloadResourceSettingsResourceTypeEnum.from_proto(
                resource.resource_type
            ),
        )


class WorkloadResourceSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkloadResourceSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkloadResourceSettings.from_proto(i) for i in resources]


class WorkloadResourcesResourceTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsWorkloadResourcesResourceTypeEnum.Value(
            "AssuredworkloadsWorkloadResourcesResourceTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsWorkloadResourcesResourceTypeEnum.Name(
            resource
        )[len("AssuredworkloadsWorkloadResourcesResourceTypeEnum") :]


class WorkloadComplianceRegimeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsWorkloadComplianceRegimeEnum.Value(
            "AssuredworkloadsWorkloadComplianceRegimeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return workload_pb2.AssuredworkloadsWorkloadComplianceRegimeEnum.Name(resource)[
            len("AssuredworkloadsWorkloadComplianceRegimeEnum") :
        ]


class WorkloadResourceSettingsResourceTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return (
            workload_pb2.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum.Value(
                "AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum%s" % resource
            )
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            workload_pb2.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum.Name(
                resource
            )[len("AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum") :]
        )


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
