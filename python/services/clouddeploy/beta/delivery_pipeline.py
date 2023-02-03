# Copyright 2023 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.clouddeploy import (
    delivery_pipeline_pb2,
)
from google3.cloud.graphite.mmv2.services.google.clouddeploy import (
    delivery_pipeline_pb2_grpc,
)

from typing import List


class DeliveryPipeline(object):
    def __init__(
        self,
        name: str = None,
        uid: str = None,
        description: str = None,
        annotations: dict = None,
        labels: dict = None,
        create_time: str = None,
        update_time: str = None,
        serial_pipeline: dict = None,
        condition: dict = None,
        etag: str = None,
        project: str = None,
        location: str = None,
        suspended: bool = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.annotations = annotations
        self.labels = labels
        self.serial_pipeline = serial_pipeline
        self.project = project
        self.location = location
        self.suspended = suspended
        self.service_account_file = service_account_file

    def apply(self):
        stub = delivery_pipeline_pb2_grpc.ClouddeployBetaDeliveryPipelineServiceStub(
            channel.Channel()
        )
        request = delivery_pipeline_pb2.ApplyClouddeployBetaDeliveryPipelineRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if DeliveryPipelineSerialPipeline.to_proto(self.serial_pipeline):
            request.resource.serial_pipeline.CopyFrom(
                DeliveryPipelineSerialPipeline.to_proto(self.serial_pipeline)
            )
        else:
            request.resource.ClearField("serial_pipeline")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.suspended):
            request.resource.suspended = Primitive.to_proto(self.suspended)

        request.service_account_file = self.service_account_file

        response = stub.ApplyClouddeployBetaDeliveryPipeline(request)
        self.name = Primitive.from_proto(response.name)
        self.uid = Primitive.from_proto(response.uid)
        self.description = Primitive.from_proto(response.description)
        self.annotations = Primitive.from_proto(response.annotations)
        self.labels = Primitive.from_proto(response.labels)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.serial_pipeline = DeliveryPipelineSerialPipeline.from_proto(
            response.serial_pipeline
        )
        self.condition = DeliveryPipelineCondition.from_proto(response.condition)
        self.etag = Primitive.from_proto(response.etag)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.suspended = Primitive.from_proto(response.suspended)

    def delete(self):
        stub = delivery_pipeline_pb2_grpc.ClouddeployBetaDeliveryPipelineServiceStub(
            channel.Channel()
        )
        request = delivery_pipeline_pb2.DeleteClouddeployBetaDeliveryPipelineRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if DeliveryPipelineSerialPipeline.to_proto(self.serial_pipeline):
            request.resource.serial_pipeline.CopyFrom(
                DeliveryPipelineSerialPipeline.to_proto(self.serial_pipeline)
            )
        else:
            request.resource.ClearField("serial_pipeline")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.suspended):
            request.resource.suspended = Primitive.to_proto(self.suspended)

        response = stub.DeleteClouddeployBetaDeliveryPipeline(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = delivery_pipeline_pb2_grpc.ClouddeployBetaDeliveryPipelineServiceStub(
            channel.Channel()
        )
        request = delivery_pipeline_pb2.ListClouddeployBetaDeliveryPipelineRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListClouddeployBetaDeliveryPipeline(request).items

    def to_proto(self):
        resource = delivery_pipeline_pb2.ClouddeployBetaDeliveryPipeline()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if DeliveryPipelineSerialPipeline.to_proto(self.serial_pipeline):
            resource.serial_pipeline.CopyFrom(
                DeliveryPipelineSerialPipeline.to_proto(self.serial_pipeline)
            )
        else:
            resource.ClearField("serial_pipeline")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.suspended):
            resource.suspended = Primitive.to_proto(self.suspended)
        return resource


class DeliveryPipelineSerialPipeline(object):
    def __init__(self, stages: list = None):
        self.stages = stages

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = delivery_pipeline_pb2.ClouddeployBetaDeliveryPipelineSerialPipeline()
        if DeliveryPipelineSerialPipelineStagesArray.to_proto(resource.stages):
            res.stages.extend(
                DeliveryPipelineSerialPipelineStagesArray.to_proto(resource.stages)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipeline(
            stages=DeliveryPipelineSerialPipelineStagesArray.from_proto(
                resource.stages
            ),
        )


class DeliveryPipelineSerialPipelineArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DeliveryPipelineSerialPipeline.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DeliveryPipelineSerialPipeline.from_proto(i) for i in resources]


class DeliveryPipelineSerialPipelineStages(object):
    def __init__(
        self, target_id: str = None, profiles: list = None, strategy: dict = None
    ):
        self.target_id = target_id
        self.profiles = profiles
        self.strategy = strategy

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployBetaDeliveryPipelineSerialPipelineStages()
        )
        if Primitive.to_proto(resource.target_id):
            res.target_id = Primitive.to_proto(resource.target_id)
        if Primitive.to_proto(resource.profiles):
            res.profiles.extend(Primitive.to_proto(resource.profiles))
        if DeliveryPipelineSerialPipelineStagesStrategy.to_proto(resource.strategy):
            res.strategy.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategy.to_proto(resource.strategy)
            )
        else:
            res.ClearField("strategy")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStages(
            target_id=Primitive.from_proto(resource.target_id),
            profiles=Primitive.from_proto(resource.profiles),
            strategy=DeliveryPipelineSerialPipelineStagesStrategy.from_proto(
                resource.strategy
            ),
        )


class DeliveryPipelineSerialPipelineStagesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DeliveryPipelineSerialPipelineStages.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DeliveryPipelineSerialPipelineStages.from_proto(i) for i in resources]


class DeliveryPipelineSerialPipelineStagesStrategy(object):
    def __init__(self, standard: dict = None):
        self.standard = standard

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategy()
        )
        if DeliveryPipelineSerialPipelineStagesStrategyStandard.to_proto(
            resource.standard
        ):
            res.standard.CopyFrom(
                DeliveryPipelineSerialPipelineStagesStrategyStandard.to_proto(
                    resource.standard
                )
            )
        else:
            res.ClearField("standard")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategy(
            standard=DeliveryPipelineSerialPipelineStagesStrategyStandard.from_proto(
                resource.standard
            ),
        )


class DeliveryPipelineSerialPipelineStagesStrategyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategy.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategy.from_proto(i)
            for i in resources
        ]


class DeliveryPipelineSerialPipelineStagesStrategyStandard(object):
    def __init__(self, verify: bool = None):
        self.verify = verify

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandard()
        )
        if Primitive.to_proto(resource.verify):
            res.verify = Primitive.to_proto(resource.verify)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineSerialPipelineStagesStrategyStandard(
            verify=Primitive.from_proto(resource.verify),
        )


class DeliveryPipelineSerialPipelineStagesStrategyStandardArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineSerialPipelineStagesStrategyStandard.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineSerialPipelineStagesStrategyStandard.from_proto(i)
            for i in resources
        ]


class DeliveryPipelineCondition(object):
    def __init__(
        self,
        pipeline_ready_condition: dict = None,
        targets_present_condition: dict = None,
    ):
        self.pipeline_ready_condition = pipeline_ready_condition
        self.targets_present_condition = targets_present_condition

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = delivery_pipeline_pb2.ClouddeployBetaDeliveryPipelineCondition()
        if DeliveryPipelineConditionPipelineReadyCondition.to_proto(
            resource.pipeline_ready_condition
        ):
            res.pipeline_ready_condition.CopyFrom(
                DeliveryPipelineConditionPipelineReadyCondition.to_proto(
                    resource.pipeline_ready_condition
                )
            )
        else:
            res.ClearField("pipeline_ready_condition")
        if DeliveryPipelineConditionTargetsPresentCondition.to_proto(
            resource.targets_present_condition
        ):
            res.targets_present_condition.CopyFrom(
                DeliveryPipelineConditionTargetsPresentCondition.to_proto(
                    resource.targets_present_condition
                )
            )
        else:
            res.ClearField("targets_present_condition")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineCondition(
            pipeline_ready_condition=DeliveryPipelineConditionPipelineReadyCondition.from_proto(
                resource.pipeline_ready_condition
            ),
            targets_present_condition=DeliveryPipelineConditionTargetsPresentCondition.from_proto(
                resource.targets_present_condition
            ),
        )


class DeliveryPipelineConditionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [DeliveryPipelineCondition.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [DeliveryPipelineCondition.from_proto(i) for i in resources]


class DeliveryPipelineConditionPipelineReadyCondition(object):
    def __init__(self, status: bool = None, update_time: str = None):
        self.status = status
        self.update_time = update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployBetaDeliveryPipelineConditionPipelineReadyCondition()
        )
        if Primitive.to_proto(resource.status):
            res.status = Primitive.to_proto(resource.status)
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineConditionPipelineReadyCondition(
            status=Primitive.from_proto(resource.status),
            update_time=Primitive.from_proto(resource.update_time),
        )


class DeliveryPipelineConditionPipelineReadyConditionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineConditionPipelineReadyCondition.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineConditionPipelineReadyCondition.from_proto(i)
            for i in resources
        ]


class DeliveryPipelineConditionTargetsPresentCondition(object):
    def __init__(
        self, status: bool = None, missing_targets: list = None, update_time: str = None
    ):
        self.status = status
        self.missing_targets = missing_targets
        self.update_time = update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            delivery_pipeline_pb2.ClouddeployBetaDeliveryPipelineConditionTargetsPresentCondition()
        )
        if Primitive.to_proto(resource.status):
            res.status = Primitive.to_proto(resource.status)
        if Primitive.to_proto(resource.missing_targets):
            res.missing_targets.extend(Primitive.to_proto(resource.missing_targets))
        if Primitive.to_proto(resource.update_time):
            res.update_time = Primitive.to_proto(resource.update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return DeliveryPipelineConditionTargetsPresentCondition(
            status=Primitive.from_proto(resource.status),
            missing_targets=Primitive.from_proto(resource.missing_targets),
            update_time=Primitive.from_proto(resource.update_time),
        )


class DeliveryPipelineConditionTargetsPresentConditionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            DeliveryPipelineConditionTargetsPresentCondition.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            DeliveryPipelineConditionTargetsPresentCondition.from_proto(i)
            for i in resources
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
