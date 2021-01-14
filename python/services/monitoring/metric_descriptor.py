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
from google3.cloud.graphite.mmv2.services.google.monitoring import metric_descriptor_pb2
from google3.cloud.graphite.mmv2.services.google.monitoring import (
    metric_descriptor_pb2_grpc,
)

from typing import List


class MetricDescriptor(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        labels: list = None,
        metric_kind: str = None,
        value_type: str = None,
        unit: str = None,
        description: str = None,
        display_name: str = None,
        metadata: dict = None,
        launch_stage: str = None,
        monitored_resource_types: list = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.type = type
        self.labels = labels
        self.metric_kind = metric_kind
        self.value_type = value_type
        self.unit = unit
        self.description = description
        self.display_name = display_name
        self.metadata = metadata
        self.launch_stage = launch_stage
        self.monitored_resource_types = monitored_resource_types
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = metric_descriptor_pb2_grpc.MonitoringMetricDescriptorServiceStub(
            channel.Channel()
        )
        request = metric_descriptor_pb2.ApplyMonitoringMetricDescriptorRequest()
        if Primitive.to_proto(self.type):
            request.resource.type = Primitive.to_proto(self.type)

        if MetricDescriptorLabelsArray.to_proto(self.labels):
            request.resource.labels.extend(
                MetricDescriptorLabelsArray.to_proto(self.labels)
            )
        if MetricDescriptorMetricKindEnum.to_proto(self.metric_kind):
            request.resource.metric_kind = MetricDescriptorMetricKindEnum.to_proto(
                self.metric_kind
            )

        if MetricDescriptorValueTypeEnum.to_proto(self.value_type):
            request.resource.value_type = MetricDescriptorValueTypeEnum.to_proto(
                self.value_type
            )

        if Primitive.to_proto(self.unit):
            request.resource.unit = Primitive.to_proto(self.unit)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if MetricDescriptorMetadata.to_proto(self.metadata):
            request.resource.metadata.CopyFrom(
                MetricDescriptorMetadata.to_proto(self.metadata)
            )
        else:
            request.resource.ClearField("metadata")
        if MetricDescriptorLaunchStageEnum.to_proto(self.launch_stage):
            request.resource.launch_stage = MetricDescriptorLaunchStageEnum.to_proto(
                self.launch_stage
            )

        if Primitive.to_proto(self.monitored_resource_types):
            request.resource.monitored_resource_types.extend(
                Primitive.to_proto(self.monitored_resource_types)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyMonitoringMetricDescriptor(request)
        self.name = Primitive.from_proto(response.name)
        self.type = Primitive.from_proto(response.type)
        self.labels = MetricDescriptorLabelsArray.from_proto(response.labels)
        self.metric_kind = MetricDescriptorMetricKindEnum.from_proto(
            response.metric_kind
        )
        self.value_type = MetricDescriptorValueTypeEnum.from_proto(response.value_type)
        self.unit = Primitive.from_proto(response.unit)
        self.description = Primitive.from_proto(response.description)
        self.display_name = Primitive.from_proto(response.display_name)
        self.metadata = MetricDescriptorMetadata.from_proto(response.metadata)
        self.launch_stage = MetricDescriptorLaunchStageEnum.from_proto(
            response.launch_stage
        )
        self.monitored_resource_types = Primitive.from_proto(
            response.monitored_resource_types
        )
        self.project = Primitive.from_proto(response.project)

    @classmethod
    def delete(self, project, type, service_account_file=""):
        stub = metric_descriptor_pb2_grpc.MonitoringMetricDescriptorServiceStub(
            channel.Channel()
        )
        request = metric_descriptor_pb2.DeleteMonitoringMetricDescriptorRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Type = type

        response = stub.DeleteMonitoringMetricDescriptor(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = metric_descriptor_pb2_grpc.MonitoringMetricDescriptorServiceStub(
            channel.Channel()
        )
        request = metric_descriptor_pb2.ListMonitoringMetricDescriptorRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListMonitoringMetricDescriptor(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = metric_descriptor_pb2.MonitoringMetricDescriptor()
        any_proto.Unpack(res_proto)

        res = MetricDescriptor()
        res.name = Primitive.from_proto(res_proto.name)
        res.type = Primitive.from_proto(res_proto.type)
        res.labels = MetricDescriptorLabelsArray.from_proto(res_proto.labels)
        res.metric_kind = MetricDescriptorMetricKindEnum.from_proto(
            res_proto.metric_kind
        )
        res.value_type = MetricDescriptorValueTypeEnum.from_proto(res_proto.value_type)
        res.unit = Primitive.from_proto(res_proto.unit)
        res.description = Primitive.from_proto(res_proto.description)
        res.display_name = Primitive.from_proto(res_proto.display_name)
        res.metadata = MetricDescriptorMetadata.from_proto(res_proto.metadata)
        res.launch_stage = MetricDescriptorLaunchStageEnum.from_proto(
            res_proto.launch_stage
        )
        res.monitored_resource_types = Primitive.from_proto(
            res_proto.monitored_resource_types
        )
        res.project = Primitive.from_proto(res_proto.project)
        return res


class MetricDescriptorLabels(object):
    def __init__(
        self, key: str = None, value_type: str = None, description: str = None
    ):
        self.key = key
        self.value_type = value_type
        self.description = description

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = metric_descriptor_pb2.MonitoringMetricDescriptorLabels()
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if MetricDescriptorLabelsValueTypeEnum.to_proto(resource.value_type):
            res.value_type = MetricDescriptorLabelsValueTypeEnum.to_proto(
                resource.value_type
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MetricDescriptorLabels(
            key=resource.key,
            value_type=resource.value_type,
            description=resource.description,
        )


class MetricDescriptorLabelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MetricDescriptorLabels.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MetricDescriptorLabels.from_proto(i) for i in resources]


class MetricDescriptorMetadata(object):
    def __init__(
        self,
        launch_stage: str = None,
        sample_period: dict = None,
        ingest_delay: dict = None,
    ):
        self.launch_stage = launch_stage
        self.sample_period = sample_period
        self.ingest_delay = ingest_delay

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = metric_descriptor_pb2.MonitoringMetricDescriptorMetadata()
        if MetricDescriptorMetadataLaunchStageEnum.to_proto(resource.launch_stage):
            res.launch_stage = MetricDescriptorMetadataLaunchStageEnum.to_proto(
                resource.launch_stage
            )
        if MetricDescriptorMetadataSamplePeriod.to_proto(resource.sample_period):
            res.sample_period.CopyFrom(
                MetricDescriptorMetadataSamplePeriod.to_proto(resource.sample_period)
            )
        else:
            res.ClearField("sample_period")
        if MetricDescriptorMetadataIngestDelay.to_proto(resource.ingest_delay):
            res.ingest_delay.CopyFrom(
                MetricDescriptorMetadataIngestDelay.to_proto(resource.ingest_delay)
            )
        else:
            res.ClearField("ingest_delay")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MetricDescriptorMetadata(
            launch_stage=resource.launch_stage,
            sample_period=resource.sample_period,
            ingest_delay=resource.ingest_delay,
        )


class MetricDescriptorMetadataArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MetricDescriptorMetadata.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MetricDescriptorMetadata.from_proto(i) for i in resources]


class MetricDescriptorMetadataSamplePeriod(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = metric_descriptor_pb2.MonitoringMetricDescriptorMetadataSamplePeriod()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MetricDescriptorMetadataSamplePeriod(
            seconds=resource.seconds, nanos=resource.nanos,
        )


class MetricDescriptorMetadataSamplePeriodArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MetricDescriptorMetadataSamplePeriod.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MetricDescriptorMetadataSamplePeriod.from_proto(i) for i in resources]


class MetricDescriptorMetadataIngestDelay(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = metric_descriptor_pb2.MonitoringMetricDescriptorMetadataIngestDelay()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return MetricDescriptorMetadataIngestDelay(
            seconds=resource.seconds, nanos=resource.nanos,
        )


class MetricDescriptorMetadataIngestDelayArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [MetricDescriptorMetadataIngestDelay.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [MetricDescriptorMetadataIngestDelay.from_proto(i) for i in resources]


class MetricDescriptorLabelsValueTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return metric_descriptor_pb2.MonitoringMetricDescriptorLabelsValueTypeEnum.Value(
            "MonitoringMetricDescriptorLabelsValueTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return metric_descriptor_pb2.MonitoringMetricDescriptorLabelsValueTypeEnum.Name(
            resource
        )[len("MonitoringMetricDescriptorLabelsValueTypeEnum") :]


class MetricDescriptorMetricKindEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return metric_descriptor_pb2.MonitoringMetricDescriptorMetricKindEnum.Value(
            "MonitoringMetricDescriptorMetricKindEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return metric_descriptor_pb2.MonitoringMetricDescriptorMetricKindEnum.Name(
            resource
        )[len("MonitoringMetricDescriptorMetricKindEnum") :]


class MetricDescriptorValueTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return metric_descriptor_pb2.MonitoringMetricDescriptorValueTypeEnum.Value(
            "MonitoringMetricDescriptorValueTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return metric_descriptor_pb2.MonitoringMetricDescriptorValueTypeEnum.Name(
            resource
        )[len("MonitoringMetricDescriptorValueTypeEnum") :]


class MetricDescriptorMetadataLaunchStageEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return metric_descriptor_pb2.MonitoringMetricDescriptorMetadataLaunchStageEnum.Value(
            "MonitoringMetricDescriptorMetadataLaunchStageEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return metric_descriptor_pb2.MonitoringMetricDescriptorMetadataLaunchStageEnum.Name(
            resource
        )[
            len("MonitoringMetricDescriptorMetadataLaunchStageEnum") :
        ]


class MetricDescriptorLaunchStageEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return metric_descriptor_pb2.MonitoringMetricDescriptorLaunchStageEnum.Value(
            "MonitoringMetricDescriptorLaunchStageEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return metric_descriptor_pb2.MonitoringMetricDescriptorLaunchStageEnum.Name(
            resource
        )[len("MonitoringMetricDescriptorLaunchStageEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
