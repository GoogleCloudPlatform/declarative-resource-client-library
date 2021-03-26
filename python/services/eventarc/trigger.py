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
from google3.cloud.graphite.mmv2.services.google.eventarc import trigger_pb2
from google3.cloud.graphite.mmv2.services.google.eventarc import trigger_pb2_grpc

from typing import List


class Trigger(object):
    def __init__(
        self,
        name: str = None,
        uid: str = None,
        create_time: str = None,
        update_time: str = None,
        event_filters: list = None,
        service_account: str = None,
        destination: dict = None,
        transport: dict = None,
        labels: dict = None,
        etag: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.event_filters = event_filters
        self.service_account = service_account
        self.destination = destination
        self.transport = transport
        self.labels = labels
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = trigger_pb2_grpc.EventarcTriggerServiceStub(channel.Channel())
        request = trigger_pb2.ApplyEventarcTriggerRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if TriggerEventFiltersArray.to_proto(self.event_filters):
            request.resource.event_filters.extend(
                TriggerEventFiltersArray.to_proto(self.event_filters)
            )
        if Primitive.to_proto(self.service_account):
            request.resource.service_account = Primitive.to_proto(self.service_account)

        if TriggerDestination.to_proto(self.destination):
            request.resource.destination.CopyFrom(
                TriggerDestination.to_proto(self.destination)
            )
        else:
            request.resource.ClearField("destination")
        if TriggerTransport.to_proto(self.transport):
            request.resource.transport.CopyFrom(
                TriggerTransport.to_proto(self.transport)
            )
        else:
            request.resource.ClearField("transport")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyEventarcTrigger(request)
        self.name = Primitive.from_proto(response.name)
        self.uid = Primitive.from_proto(response.uid)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.event_filters = TriggerEventFiltersArray.from_proto(response.event_filters)
        self.service_account = Primitive.from_proto(response.service_account)
        self.destination = TriggerDestination.from_proto(response.destination)
        self.transport = TriggerTransport.from_proto(response.transport)
        self.labels = Primitive.from_proto(response.labels)
        self.etag = Primitive.from_proto(response.etag)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = trigger_pb2_grpc.EventarcTriggerServiceStub(channel.Channel())
        request = trigger_pb2.DeleteEventarcTriggerRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if TriggerEventFiltersArray.to_proto(self.event_filters):
            request.resource.event_filters.extend(
                TriggerEventFiltersArray.to_proto(self.event_filters)
            )
        if Primitive.to_proto(self.service_account):
            request.resource.service_account = Primitive.to_proto(self.service_account)

        if TriggerDestination.to_proto(self.destination):
            request.resource.destination.CopyFrom(
                TriggerDestination.to_proto(self.destination)
            )
        else:
            request.resource.ClearField("destination")
        if TriggerTransport.to_proto(self.transport):
            request.resource.transport.CopyFrom(
                TriggerTransport.to_proto(self.transport)
            )
        else:
            request.resource.ClearField("transport")
        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteEventarcTrigger(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = trigger_pb2_grpc.EventarcTriggerServiceStub(channel.Channel())
        request = trigger_pb2.ListEventarcTriggerRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListEventarcTrigger(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = trigger_pb2.EventarcTrigger()
        any_proto.Unpack(res_proto)

        res = Trigger()
        res.name = Primitive.from_proto(res_proto.name)
        res.uid = Primitive.from_proto(res_proto.uid)
        res.create_time = Primitive.from_proto(res_proto.create_time)
        res.update_time = Primitive.from_proto(res_proto.update_time)
        res.event_filters = TriggerEventFiltersArray.from_proto(res_proto.event_filters)
        res.service_account = Primitive.from_proto(res_proto.service_account)
        res.destination = TriggerDestination.from_proto(res_proto.destination)
        res.transport = TriggerTransport.from_proto(res_proto.transport)
        res.labels = Primitive.from_proto(res_proto.labels)
        res.etag = Primitive.from_proto(res_proto.etag)
        res.project = Primitive.from_proto(res_proto.project)
        res.location = Primitive.from_proto(res_proto.location)
        return res

    def to_proto(self):
        resource = trigger_pb2.EventarcTrigger()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if TriggerEventFiltersArray.to_proto(self.event_filters):
            resource.event_filters.extend(
                TriggerEventFiltersArray.to_proto(self.event_filters)
            )
        if Primitive.to_proto(self.service_account):
            resource.service_account = Primitive.to_proto(self.service_account)
        if TriggerDestination.to_proto(self.destination):
            resource.destination.CopyFrom(TriggerDestination.to_proto(self.destination))
        else:
            resource.ClearField("destination")
        if TriggerTransport.to_proto(self.transport):
            resource.transport.CopyFrom(TriggerTransport.to_proto(self.transport))
        else:
            resource.ClearField("transport")
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class TriggerEventFilters(object):
    def __init__(self, attribute: str = None, value: str = None):
        self.attribute = attribute
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = trigger_pb2.EventarcTriggerEventFilters()
        if Primitive.to_proto(resource.attribute):
            res.attribute = Primitive.to_proto(resource.attribute)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TriggerEventFilters(attribute=resource.attribute, value=resource.value,)


class TriggerEventFiltersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TriggerEventFilters.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TriggerEventFilters.from_proto(i) for i in resources]


class TriggerDestination(object):
    def __init__(self, cloud_run: dict = None, cloud_function: str = None):
        self.cloud_run = cloud_run
        self.cloud_function = cloud_function

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = trigger_pb2.EventarcTriggerDestination()
        if TriggerDestinationCloudRun.to_proto(resource.cloud_run):
            res.cloud_run.CopyFrom(
                TriggerDestinationCloudRun.to_proto(resource.cloud_run)
            )
        else:
            res.ClearField("cloud_run")
        if Primitive.to_proto(resource.cloud_function):
            res.cloud_function = Primitive.to_proto(resource.cloud_function)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TriggerDestination(
            cloud_run=resource.cloud_run, cloud_function=resource.cloud_function,
        )


class TriggerDestinationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TriggerDestination.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TriggerDestination.from_proto(i) for i in resources]


class TriggerDestinationCloudRun(object):
    def __init__(self, service: str = None, path: str = None, region: str = None):
        self.service = service
        self.path = path
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = trigger_pb2.EventarcTriggerDestinationCloudRun()
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TriggerDestinationCloudRun(
            service=resource.service, path=resource.path, region=resource.region,
        )


class TriggerDestinationCloudRunArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TriggerDestinationCloudRun.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TriggerDestinationCloudRun.from_proto(i) for i in resources]


class TriggerTransport(object):
    def __init__(self, pubsub: dict = None):
        self.pubsub = pubsub

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = trigger_pb2.EventarcTriggerTransport()
        if TriggerTransportPubsub.to_proto(resource.pubsub):
            res.pubsub.CopyFrom(TriggerTransportPubsub.to_proto(resource.pubsub))
        else:
            res.ClearField("pubsub")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TriggerTransport(pubsub=resource.pubsub,)


class TriggerTransportArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TriggerTransport.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TriggerTransport.from_proto(i) for i in resources]


class TriggerTransportPubsub(object):
    def __init__(self, topic: str = None, subscription: str = None):
        self.topic = topic
        self.subscription = subscription

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = trigger_pb2.EventarcTriggerTransportPubsub()
        if Primitive.to_proto(resource.topic):
            res.topic = Primitive.to_proto(resource.topic)
        if Primitive.to_proto(resource.subscription):
            res.subscription = Primitive.to_proto(resource.subscription)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TriggerTransportPubsub(
            topic=resource.topic, subscription=resource.subscription,
        )


class TriggerTransportPubsubArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TriggerTransportPubsub.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TriggerTransportPubsub.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
