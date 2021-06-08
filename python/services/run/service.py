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
from google3.cloud.graphite.mmv2.services.google.run import service_pb2
from google3.cloud.graphite.mmv2.services.google.run import service_pb2_grpc

from typing import List


class Service(object):
    def __init__(
        self,
        api_version: str = None,
        kind: str = None,
        metadata: dict = None,
        spec: dict = None,
        status: dict = None,
        project: str = None,
        location: str = None,
        name: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.api_version = api_version
        self.kind = kind
        self.metadata = metadata
        self.spec = spec
        self.project = project
        self.location = location
        self.name = name
        self.service_account_file = service_account_file

    def apply(self):
        stub = service_pb2_grpc.RunServiceServiceStub(channel.Channel())
        request = service_pb2.ApplyRunServiceRequest()
        if Primitive.to_proto(self.api_version):
            request.resource.api_version = Primitive.to_proto(self.api_version)

        if Primitive.to_proto(self.kind):
            request.resource.kind = Primitive.to_proto(self.kind)

        if ServiceMetadata.to_proto(self.metadata):
            request.resource.metadata.CopyFrom(ServiceMetadata.to_proto(self.metadata))
        else:
            request.resource.ClearField("metadata")
        if ServiceSpec.to_proto(self.spec):
            request.resource.spec.CopyFrom(ServiceSpec.to_proto(self.spec))
        else:
            request.resource.ClearField("spec")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        request.service_account_file = self.service_account_file

        response = stub.ApplyRunService(request)
        self.api_version = Primitive.from_proto(response.api_version)
        self.kind = Primitive.from_proto(response.kind)
        self.metadata = ServiceMetadata.from_proto(response.metadata)
        self.spec = ServiceSpec.from_proto(response.spec)
        self.status = ServiceStatus.from_proto(response.status)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.name = Primitive.from_proto(response.name)

    def delete(self):
        stub = service_pb2_grpc.RunServiceServiceStub(channel.Channel())
        request = service_pb2.DeleteRunServiceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.api_version):
            request.resource.api_version = Primitive.to_proto(self.api_version)

        if Primitive.to_proto(self.kind):
            request.resource.kind = Primitive.to_proto(self.kind)

        if ServiceMetadata.to_proto(self.metadata):
            request.resource.metadata.CopyFrom(ServiceMetadata.to_proto(self.metadata))
        else:
            request.resource.ClearField("metadata")
        if ServiceSpec.to_proto(self.spec):
            request.resource.spec.CopyFrom(ServiceSpec.to_proto(self.spec))
        else:
            request.resource.ClearField("spec")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        response = stub.DeleteRunService(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = service_pb2_grpc.RunServiceServiceStub(channel.Channel())
        request = service_pb2.ListRunServiceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListRunService(request).items

    def to_proto(self):
        resource = service_pb2.RunService()
        if Primitive.to_proto(self.api_version):
            resource.api_version = Primitive.to_proto(self.api_version)
        if Primitive.to_proto(self.kind):
            resource.kind = Primitive.to_proto(self.kind)
        if ServiceMetadata.to_proto(self.metadata):
            resource.metadata.CopyFrom(ServiceMetadata.to_proto(self.metadata))
        else:
            resource.ClearField("metadata")
        if ServiceSpec.to_proto(self.spec):
            resource.spec.CopyFrom(ServiceSpec.to_proto(self.spec))
        else:
            resource.ClearField("spec")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        return resource


class ServiceMetadata(object):
    def __init__(
        self,
        name: str = None,
        generate_name: str = None,
        namespace: str = None,
        self_link: str = None,
        uid: str = None,
        resource_version: str = None,
        generation: int = None,
        create_time: dict = None,
        labels: dict = None,
        annotations: dict = None,
        owner_references: list = None,
        delete_time: dict = None,
        deletion_grace_period_seconds: int = None,
        finalizers: list = None,
        cluster_name: str = None,
    ):
        self.name = name
        self.generate_name = generate_name
        self.namespace = namespace
        self.self_link = self_link
        self.uid = uid
        self.resource_version = resource_version
        self.generation = generation
        self.create_time = create_time
        self.labels = labels
        self.annotations = annotations
        self.owner_references = owner_references
        self.delete_time = delete_time
        self.deletion_grace_period_seconds = deletion_grace_period_seconds
        self.finalizers = finalizers
        self.cluster_name = cluster_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceMetadata()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.generate_name):
            res.generate_name = Primitive.to_proto(resource.generate_name)
        if Primitive.to_proto(resource.namespace):
            res.namespace = Primitive.to_proto(resource.namespace)
        if Primitive.to_proto(resource.self_link):
            res.self_link = Primitive.to_proto(resource.self_link)
        if Primitive.to_proto(resource.uid):
            res.uid = Primitive.to_proto(resource.uid)
        if Primitive.to_proto(resource.resource_version):
            res.resource_version = Primitive.to_proto(resource.resource_version)
        if Primitive.to_proto(resource.generation):
            res.generation = Primitive.to_proto(resource.generation)
        if ServiceMetadataCreateTime.to_proto(resource.create_time):
            res.create_time.CopyFrom(
                ServiceMetadataCreateTime.to_proto(resource.create_time)
            )
        else:
            res.ClearField("create_time")
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if Primitive.to_proto(resource.annotations):
            res.annotations = Primitive.to_proto(resource.annotations)
        if ServiceMetadataOwnerReferencesArray.to_proto(resource.owner_references):
            res.owner_references.extend(
                ServiceMetadataOwnerReferencesArray.to_proto(resource.owner_references)
            )
        if ServiceMetadataDeleteTime.to_proto(resource.delete_time):
            res.delete_time.CopyFrom(
                ServiceMetadataDeleteTime.to_proto(resource.delete_time)
            )
        else:
            res.ClearField("delete_time")
        if Primitive.to_proto(resource.deletion_grace_period_seconds):
            res.deletion_grace_period_seconds = Primitive.to_proto(
                resource.deletion_grace_period_seconds
            )
        if Primitive.to_proto(resource.finalizers):
            res.finalizers.extend(Primitive.to_proto(resource.finalizers))
        if Primitive.to_proto(resource.cluster_name):
            res.cluster_name = Primitive.to_proto(resource.cluster_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceMetadata(
            name=Primitive.from_proto(resource.name),
            generate_name=Primitive.from_proto(resource.generate_name),
            namespace=Primitive.from_proto(resource.namespace),
            self_link=Primitive.from_proto(resource.self_link),
            uid=Primitive.from_proto(resource.uid),
            resource_version=Primitive.from_proto(resource.resource_version),
            generation=Primitive.from_proto(resource.generation),
            create_time=ServiceMetadataCreateTime.from_proto(resource.create_time),
            labels=Primitive.from_proto(resource.labels),
            annotations=Primitive.from_proto(resource.annotations),
            owner_references=ServiceMetadataOwnerReferencesArray.from_proto(
                resource.owner_references
            ),
            delete_time=ServiceMetadataDeleteTime.from_proto(resource.delete_time),
            deletion_grace_period_seconds=Primitive.from_proto(
                resource.deletion_grace_period_seconds
            ),
            finalizers=Primitive.from_proto(resource.finalizers),
            cluster_name=Primitive.from_proto(resource.cluster_name),
        )


class ServiceMetadataArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceMetadata.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceMetadata.from_proto(i) for i in resources]


class ServiceMetadataCreateTime(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceMetadataCreateTime()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceMetadataCreateTime(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class ServiceMetadataCreateTimeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceMetadataCreateTime.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceMetadataCreateTime.from_proto(i) for i in resources]


class ServiceMetadataOwnerReferences(object):
    def __init__(
        self,
        api_version: str = None,
        kind: str = None,
        name: str = None,
        uid: str = None,
        controller: bool = None,
        block_owner_deletion: bool = None,
    ):
        self.api_version = api_version
        self.kind = kind
        self.name = name
        self.uid = uid
        self.controller = controller
        self.block_owner_deletion = block_owner_deletion

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceMetadataOwnerReferences()
        if Primitive.to_proto(resource.api_version):
            res.api_version = Primitive.to_proto(resource.api_version)
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.uid):
            res.uid = Primitive.to_proto(resource.uid)
        if Primitive.to_proto(resource.controller):
            res.controller = Primitive.to_proto(resource.controller)
        if Primitive.to_proto(resource.block_owner_deletion):
            res.block_owner_deletion = Primitive.to_proto(resource.block_owner_deletion)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceMetadataOwnerReferences(
            api_version=Primitive.from_proto(resource.api_version),
            kind=Primitive.from_proto(resource.kind),
            name=Primitive.from_proto(resource.name),
            uid=Primitive.from_proto(resource.uid),
            controller=Primitive.from_proto(resource.controller),
            block_owner_deletion=Primitive.from_proto(resource.block_owner_deletion),
        )


class ServiceMetadataOwnerReferencesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceMetadataOwnerReferences.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceMetadataOwnerReferences.from_proto(i) for i in resources]


class ServiceMetadataDeleteTime(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceMetadataDeleteTime()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceMetadataDeleteTime(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class ServiceMetadataDeleteTimeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceMetadataDeleteTime.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceMetadataDeleteTime.from_proto(i) for i in resources]


class ServiceSpec(object):
    def __init__(self, template: dict = None, traffic: list = None):
        self.template = template
        self.traffic = traffic

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpec()
        if ServiceSpecTemplate.to_proto(resource.template):
            res.template.CopyFrom(ServiceSpecTemplate.to_proto(resource.template))
        else:
            res.ClearField("template")
        if ServiceSpecTrafficArray.to_proto(resource.traffic):
            res.traffic.extend(ServiceSpecTrafficArray.to_proto(resource.traffic))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpec(
            template=ServiceSpecTemplate.from_proto(resource.template),
            traffic=ServiceSpecTrafficArray.from_proto(resource.traffic),
        )


class ServiceSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceSpec.from_proto(i) for i in resources]


class ServiceSpecTemplate(object):
    def __init__(self, metadata: dict = None, spec: dict = None):
        self.metadata = metadata
        self.spec = spec

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplate()
        if ServiceSpecTemplateMetadata.to_proto(resource.metadata):
            res.metadata.CopyFrom(
                ServiceSpecTemplateMetadata.to_proto(resource.metadata)
            )
        else:
            res.ClearField("metadata")
        if ServiceSpecTemplateSpec.to_proto(resource.spec):
            res.spec.CopyFrom(ServiceSpecTemplateSpec.to_proto(resource.spec))
        else:
            res.ClearField("spec")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplate(
            metadata=ServiceSpecTemplateMetadata.from_proto(resource.metadata),
            spec=ServiceSpecTemplateSpec.from_proto(resource.spec),
        )


class ServiceSpecTemplateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceSpecTemplate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceSpecTemplate.from_proto(i) for i in resources]


class ServiceSpecTemplateMetadata(object):
    def __init__(
        self,
        name: str = None,
        generate_name: str = None,
        namespace: str = None,
        self_link: str = None,
        uid: str = None,
        resource_version: str = None,
        generation: int = None,
        create_time: dict = None,
        labels: dict = None,
        annotations: dict = None,
        owner_references: list = None,
        delete_time: dict = None,
        deletion_grace_period_seconds: int = None,
        finalizers: list = None,
        cluster_name: str = None,
    ):
        self.name = name
        self.generate_name = generate_name
        self.namespace = namespace
        self.self_link = self_link
        self.uid = uid
        self.resource_version = resource_version
        self.generation = generation
        self.create_time = create_time
        self.labels = labels
        self.annotations = annotations
        self.owner_references = owner_references
        self.delete_time = delete_time
        self.deletion_grace_period_seconds = deletion_grace_period_seconds
        self.finalizers = finalizers
        self.cluster_name = cluster_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateMetadata()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.generate_name):
            res.generate_name = Primitive.to_proto(resource.generate_name)
        if Primitive.to_proto(resource.namespace):
            res.namespace = Primitive.to_proto(resource.namespace)
        if Primitive.to_proto(resource.self_link):
            res.self_link = Primitive.to_proto(resource.self_link)
        if Primitive.to_proto(resource.uid):
            res.uid = Primitive.to_proto(resource.uid)
        if Primitive.to_proto(resource.resource_version):
            res.resource_version = Primitive.to_proto(resource.resource_version)
        if Primitive.to_proto(resource.generation):
            res.generation = Primitive.to_proto(resource.generation)
        if ServiceSpecTemplateMetadataCreateTime.to_proto(resource.create_time):
            res.create_time.CopyFrom(
                ServiceSpecTemplateMetadataCreateTime.to_proto(resource.create_time)
            )
        else:
            res.ClearField("create_time")
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if Primitive.to_proto(resource.annotations):
            res.annotations = Primitive.to_proto(resource.annotations)
        if ServiceSpecTemplateMetadataOwnerReferencesArray.to_proto(
            resource.owner_references
        ):
            res.owner_references.extend(
                ServiceSpecTemplateMetadataOwnerReferencesArray.to_proto(
                    resource.owner_references
                )
            )
        if ServiceSpecTemplateMetadataDeleteTime.to_proto(resource.delete_time):
            res.delete_time.CopyFrom(
                ServiceSpecTemplateMetadataDeleteTime.to_proto(resource.delete_time)
            )
        else:
            res.ClearField("delete_time")
        if Primitive.to_proto(resource.deletion_grace_period_seconds):
            res.deletion_grace_period_seconds = Primitive.to_proto(
                resource.deletion_grace_period_seconds
            )
        if Primitive.to_proto(resource.finalizers):
            res.finalizers.extend(Primitive.to_proto(resource.finalizers))
        if Primitive.to_proto(resource.cluster_name):
            res.cluster_name = Primitive.to_proto(resource.cluster_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateMetadata(
            name=Primitive.from_proto(resource.name),
            generate_name=Primitive.from_proto(resource.generate_name),
            namespace=Primitive.from_proto(resource.namespace),
            self_link=Primitive.from_proto(resource.self_link),
            uid=Primitive.from_proto(resource.uid),
            resource_version=Primitive.from_proto(resource.resource_version),
            generation=Primitive.from_proto(resource.generation),
            create_time=ServiceSpecTemplateMetadataCreateTime.from_proto(
                resource.create_time
            ),
            labels=Primitive.from_proto(resource.labels),
            annotations=Primitive.from_proto(resource.annotations),
            owner_references=ServiceSpecTemplateMetadataOwnerReferencesArray.from_proto(
                resource.owner_references
            ),
            delete_time=ServiceSpecTemplateMetadataDeleteTime.from_proto(
                resource.delete_time
            ),
            deletion_grace_period_seconds=Primitive.from_proto(
                resource.deletion_grace_period_seconds
            ),
            finalizers=Primitive.from_proto(resource.finalizers),
            cluster_name=Primitive.from_proto(resource.cluster_name),
        )


class ServiceSpecTemplateMetadataArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceSpecTemplateMetadata.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceSpecTemplateMetadata.from_proto(i) for i in resources]


class ServiceSpecTemplateMetadataCreateTime(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateMetadataCreateTime()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateMetadataCreateTime(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class ServiceSpecTemplateMetadataCreateTimeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceSpecTemplateMetadataCreateTime.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceSpecTemplateMetadataCreateTime.from_proto(i) for i in resources]


class ServiceSpecTemplateMetadataOwnerReferences(object):
    def __init__(
        self,
        api_version: str = None,
        kind: str = None,
        name: str = None,
        uid: str = None,
        controller: bool = None,
        block_owner_deletion: bool = None,
    ):
        self.api_version = api_version
        self.kind = kind
        self.name = name
        self.uid = uid
        self.controller = controller
        self.block_owner_deletion = block_owner_deletion

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateMetadataOwnerReferences()
        if Primitive.to_proto(resource.api_version):
            res.api_version = Primitive.to_proto(resource.api_version)
        if Primitive.to_proto(resource.kind):
            res.kind = Primitive.to_proto(resource.kind)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.uid):
            res.uid = Primitive.to_proto(resource.uid)
        if Primitive.to_proto(resource.controller):
            res.controller = Primitive.to_proto(resource.controller)
        if Primitive.to_proto(resource.block_owner_deletion):
            res.block_owner_deletion = Primitive.to_proto(resource.block_owner_deletion)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateMetadataOwnerReferences(
            api_version=Primitive.from_proto(resource.api_version),
            kind=Primitive.from_proto(resource.kind),
            name=Primitive.from_proto(resource.name),
            uid=Primitive.from_proto(resource.uid),
            controller=Primitive.from_proto(resource.controller),
            block_owner_deletion=Primitive.from_proto(resource.block_owner_deletion),
        )


class ServiceSpecTemplateMetadataOwnerReferencesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateMetadataOwnerReferences.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateMetadataOwnerReferences.from_proto(i) for i in resources
        ]


class ServiceSpecTemplateMetadataDeleteTime(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateMetadataDeleteTime()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateMetadataDeleteTime(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class ServiceSpecTemplateMetadataDeleteTimeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceSpecTemplateMetadataDeleteTime.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceSpecTemplateMetadataDeleteTime.from_proto(i) for i in resources]


class ServiceSpecTemplateSpec(object):
    def __init__(
        self,
        container_concurrency: int = None,
        timeout_seconds: int = None,
        service_account_name: str = None,
        containers: list = None,
        volumes: list = None,
    ):
        self.container_concurrency = container_concurrency
        self.timeout_seconds = timeout_seconds
        self.service_account_name = service_account_name
        self.containers = containers
        self.volumes = volumes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpec()
        if Primitive.to_proto(resource.container_concurrency):
            res.container_concurrency = Primitive.to_proto(
                resource.container_concurrency
            )
        if Primitive.to_proto(resource.timeout_seconds):
            res.timeout_seconds = Primitive.to_proto(resource.timeout_seconds)
        if Primitive.to_proto(resource.service_account_name):
            res.service_account_name = Primitive.to_proto(resource.service_account_name)
        if ServiceSpecTemplateSpecContainersArray.to_proto(resource.containers):
            res.containers.extend(
                ServiceSpecTemplateSpecContainersArray.to_proto(resource.containers)
            )
        if ServiceSpecTemplateSpecVolumesArray.to_proto(resource.volumes):
            res.volumes.extend(
                ServiceSpecTemplateSpecVolumesArray.to_proto(resource.volumes)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpec(
            container_concurrency=Primitive.from_proto(resource.container_concurrency),
            timeout_seconds=Primitive.from_proto(resource.timeout_seconds),
            service_account_name=Primitive.from_proto(resource.service_account_name),
            containers=ServiceSpecTemplateSpecContainersArray.from_proto(
                resource.containers
            ),
            volumes=ServiceSpecTemplateSpecVolumesArray.from_proto(resource.volumes),
        )


class ServiceSpecTemplateSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceSpecTemplateSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceSpecTemplateSpec.from_proto(i) for i in resources]


class ServiceSpecTemplateSpecContainers(object):
    def __init__(
        self,
        name: str = None,
        image: str = None,
        command: list = None,
        args: list = None,
        env: list = None,
        resources: dict = None,
        working_dir: str = None,
        ports: list = None,
        env_from: list = None,
        volume_mounts: list = None,
        liveness_probe: dict = None,
        readiness_probe: dict = None,
        termination_message_path: str = None,
        termination_message_policy: str = None,
        image_pull_policy: str = None,
        security_context: dict = None,
    ):
        self.name = name
        self.image = image
        self.command = command
        self.args = args
        self.env = env
        self.resources = resources
        self.working_dir = working_dir
        self.ports = ports
        self.env_from = env_from
        self.volume_mounts = volume_mounts
        self.liveness_probe = liveness_probe
        self.readiness_probe = readiness_probe
        self.termination_message_path = termination_message_path
        self.termination_message_policy = termination_message_policy
        self.image_pull_policy = image_pull_policy
        self.security_context = security_context

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainers()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.image):
            res.image = Primitive.to_proto(resource.image)
        if Primitive.to_proto(resource.command):
            res.command.extend(Primitive.to_proto(resource.command))
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if ServiceSpecTemplateSpecContainersEnvArray.to_proto(resource.env):
            res.env.extend(
                ServiceSpecTemplateSpecContainersEnvArray.to_proto(resource.env)
            )
        if ServiceSpecTemplateSpecContainersResources.to_proto(resource.resources):
            res.resources.CopyFrom(
                ServiceSpecTemplateSpecContainersResources.to_proto(resource.resources)
            )
        else:
            res.ClearField("resources")
        if Primitive.to_proto(resource.working_dir):
            res.working_dir = Primitive.to_proto(resource.working_dir)
        if ServiceSpecTemplateSpecContainersPortsArray.to_proto(resource.ports):
            res.ports.extend(
                ServiceSpecTemplateSpecContainersPortsArray.to_proto(resource.ports)
            )
        if ServiceSpecTemplateSpecContainersEnvFromArray.to_proto(resource.env_from):
            res.env_from.extend(
                ServiceSpecTemplateSpecContainersEnvFromArray.to_proto(
                    resource.env_from
                )
            )
        if ServiceSpecTemplateSpecContainersVolumeMountsArray.to_proto(
            resource.volume_mounts
        ):
            res.volume_mounts.extend(
                ServiceSpecTemplateSpecContainersVolumeMountsArray.to_proto(
                    resource.volume_mounts
                )
            )
        if ServiceSpecTemplateSpecContainersLivenessProbe.to_proto(
            resource.liveness_probe
        ):
            res.liveness_probe.CopyFrom(
                ServiceSpecTemplateSpecContainersLivenessProbe.to_proto(
                    resource.liveness_probe
                )
            )
        else:
            res.ClearField("liveness_probe")
        if ServiceSpecTemplateSpecContainersReadinessProbe.to_proto(
            resource.readiness_probe
        ):
            res.readiness_probe.CopyFrom(
                ServiceSpecTemplateSpecContainersReadinessProbe.to_proto(
                    resource.readiness_probe
                )
            )
        else:
            res.ClearField("readiness_probe")
        if Primitive.to_proto(resource.termination_message_path):
            res.termination_message_path = Primitive.to_proto(
                resource.termination_message_path
            )
        if Primitive.to_proto(resource.termination_message_policy):
            res.termination_message_policy = Primitive.to_proto(
                resource.termination_message_policy
            )
        if Primitive.to_proto(resource.image_pull_policy):
            res.image_pull_policy = Primitive.to_proto(resource.image_pull_policy)
        if ServiceSpecTemplateSpecContainersSecurityContext.to_proto(
            resource.security_context
        ):
            res.security_context.CopyFrom(
                ServiceSpecTemplateSpecContainersSecurityContext.to_proto(
                    resource.security_context
                )
            )
        else:
            res.ClearField("security_context")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainers(
            name=Primitive.from_proto(resource.name),
            image=Primitive.from_proto(resource.image),
            command=Primitive.from_proto(resource.command),
            args=Primitive.from_proto(resource.args),
            env=ServiceSpecTemplateSpecContainersEnvArray.from_proto(resource.env),
            resources=ServiceSpecTemplateSpecContainersResources.from_proto(
                resource.resources
            ),
            working_dir=Primitive.from_proto(resource.working_dir),
            ports=ServiceSpecTemplateSpecContainersPortsArray.from_proto(
                resource.ports
            ),
            env_from=ServiceSpecTemplateSpecContainersEnvFromArray.from_proto(
                resource.env_from
            ),
            volume_mounts=ServiceSpecTemplateSpecContainersVolumeMountsArray.from_proto(
                resource.volume_mounts
            ),
            liveness_probe=ServiceSpecTemplateSpecContainersLivenessProbe.from_proto(
                resource.liveness_probe
            ),
            readiness_probe=ServiceSpecTemplateSpecContainersReadinessProbe.from_proto(
                resource.readiness_probe
            ),
            termination_message_path=Primitive.from_proto(
                resource.termination_message_path
            ),
            termination_message_policy=Primitive.from_proto(
                resource.termination_message_policy
            ),
            image_pull_policy=Primitive.from_proto(resource.image_pull_policy),
            security_context=ServiceSpecTemplateSpecContainersSecurityContext.from_proto(
                resource.security_context
            ),
        )


class ServiceSpecTemplateSpecContainersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceSpecTemplateSpecContainers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceSpecTemplateSpecContainers.from_proto(i) for i in resources]


class ServiceSpecTemplateSpecContainersEnv(object):
    def __init__(self, name: str = None, value: str = None, value_from: dict = None):
        self.name = name
        self.value = value
        self.value_from = value_from

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersEnv()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        if ServiceSpecTemplateSpecContainersEnvValueFrom.to_proto(resource.value_from):
            res.value_from.CopyFrom(
                ServiceSpecTemplateSpecContainersEnvValueFrom.to_proto(
                    resource.value_from
                )
            )
        else:
            res.ClearField("value_from")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersEnv(
            name=Primitive.from_proto(resource.name),
            value=Primitive.from_proto(resource.value),
            value_from=ServiceSpecTemplateSpecContainersEnvValueFrom.from_proto(
                resource.value_from
            ),
        )


class ServiceSpecTemplateSpecContainersEnvArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceSpecTemplateSpecContainersEnv.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceSpecTemplateSpecContainersEnv.from_proto(i) for i in resources]


class ServiceSpecTemplateSpecContainersEnvValueFrom(object):
    def __init__(self, config_map_key_ref: dict = None, secret_key_ref: dict = None):
        self.config_map_key_ref = config_map_key_ref
        self.secret_key_ref = secret_key_ref

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersEnvValueFrom()
        if ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef.to_proto(
            resource.config_map_key_ref
        ):
            res.config_map_key_ref.CopyFrom(
                ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef.to_proto(
                    resource.config_map_key_ref
                )
            )
        else:
            res.ClearField("config_map_key_ref")
        if ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef.to_proto(
            resource.secret_key_ref
        ):
            res.secret_key_ref.CopyFrom(
                ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef.to_proto(
                    resource.secret_key_ref
                )
            )
        else:
            res.ClearField("secret_key_ref")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersEnvValueFrom(
            config_map_key_ref=ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef.from_proto(
                resource.config_map_key_ref
            ),
            secret_key_ref=ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef.from_proto(
                resource.secret_key_ref
            ),
        )


class ServiceSpecTemplateSpecContainersEnvValueFromArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersEnvValueFrom.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersEnvValueFrom.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(object):
    def __init__(
        self,
        local_object_reference: dict = None,
        key: str = None,
        optional: bool = None,
        name: str = None,
    ):
        self.local_object_reference = local_object_reference
        self.key = key
        self.optional = optional
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_pb2.RunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef()
        )
        if ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference.to_proto(
            resource.local_object_reference
        ):
            res.local_object_reference.CopyFrom(
                ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference.to_proto(
                    resource.local_object_reference
                )
            )
        else:
            res.ClearField("local_object_reference")
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.optional):
            res.optional = Primitive.to_proto(resource.optional)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef(
            local_object_reference=ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference.from_proto(
                resource.local_object_reference
            ),
            key=Primitive.from_proto(resource.key),
            optional=Primitive.from_proto(resource.optional),
            name=Primitive.from_proto(resource.name),
        )


class ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_pb2.RunServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference(
            name=Primitive.from_proto(resource.name),
        )


class ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReferenceArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersEnvValueFromConfigMapKeyRefLocalObjectReference.from_proto(
                i
            )
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(object):
    def __init__(
        self,
        local_object_reference: dict = None,
        key: str = None,
        optional: bool = None,
        name: str = None,
    ):
        self.local_object_reference = local_object_reference
        self.key = key
        self.optional = optional
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef()
        if ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference.to_proto(
            resource.local_object_reference
        ):
            res.local_object_reference.CopyFrom(
                ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference.to_proto(
                    resource.local_object_reference
                )
            )
        else:
            res.ClearField("local_object_reference")
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.optional):
            res.optional = Primitive.to_proto(resource.optional)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef(
            local_object_reference=ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference.from_proto(
                resource.local_object_reference
            ),
            key=Primitive.from_proto(resource.key),
            optional=Primitive.from_proto(resource.optional),
            name=Primitive.from_proto(resource.name),
        )


class ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(
    object
):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_pb2.RunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference(
            name=Primitive.from_proto(resource.name),
        )


class ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReferenceArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRefLocalObjectReference.from_proto(
                i
            )
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersResources(object):
    def __init__(self, limits: dict = None, requests: dict = None):
        self.limits = limits
        self.requests = requests

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersResources()
        if Primitive.to_proto(resource.limits):
            res.limits = Primitive.to_proto(resource.limits)
        if Primitive.to_proto(resource.requests):
            res.requests = Primitive.to_proto(resource.requests)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersResources(
            limits=Primitive.from_proto(resource.limits),
            requests=Primitive.from_proto(resource.requests),
        )


class ServiceSpecTemplateSpecContainersResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersResources.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersResources.from_proto(i) for i in resources
        ]


class ServiceSpecTemplateSpecContainersPorts(object):
    def __init__(
        self, name: str = None, container_port: int = None, protocol: str = None
    ):
        self.name = name
        self.container_port = container_port
        self.protocol = protocol

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersPorts()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.container_port):
            res.container_port = Primitive.to_proto(resource.container_port)
        if Primitive.to_proto(resource.protocol):
            res.protocol = Primitive.to_proto(resource.protocol)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersPorts(
            name=Primitive.from_proto(resource.name),
            container_port=Primitive.from_proto(resource.container_port),
            protocol=Primitive.from_proto(resource.protocol),
        )


class ServiceSpecTemplateSpecContainersPortsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceSpecTemplateSpecContainersPorts.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceSpecTemplateSpecContainersPorts.from_proto(i) for i in resources]


class ServiceSpecTemplateSpecContainersEnvFrom(object):
    def __init__(
        self, prefix: str = None, config_map_ref: dict = None, secret_ref: dict = None
    ):
        self.prefix = prefix
        self.config_map_ref = config_map_ref
        self.secret_ref = secret_ref

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersEnvFrom()
        if Primitive.to_proto(resource.prefix):
            res.prefix = Primitive.to_proto(resource.prefix)
        if ServiceSpecTemplateSpecContainersEnvFromConfigMapRef.to_proto(
            resource.config_map_ref
        ):
            res.config_map_ref.CopyFrom(
                ServiceSpecTemplateSpecContainersEnvFromConfigMapRef.to_proto(
                    resource.config_map_ref
                )
            )
        else:
            res.ClearField("config_map_ref")
        if ServiceSpecTemplateSpecContainersEnvFromSecretRef.to_proto(
            resource.secret_ref
        ):
            res.secret_ref.CopyFrom(
                ServiceSpecTemplateSpecContainersEnvFromSecretRef.to_proto(
                    resource.secret_ref
                )
            )
        else:
            res.ClearField("secret_ref")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersEnvFrom(
            prefix=Primitive.from_proto(resource.prefix),
            config_map_ref=ServiceSpecTemplateSpecContainersEnvFromConfigMapRef.from_proto(
                resource.config_map_ref
            ),
            secret_ref=ServiceSpecTemplateSpecContainersEnvFromSecretRef.from_proto(
                resource.secret_ref
            ),
        )


class ServiceSpecTemplateSpecContainersEnvFromArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceSpecTemplateSpecContainersEnvFrom.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersEnvFrom.from_proto(i) for i in resources
        ]


class ServiceSpecTemplateSpecContainersEnvFromConfigMapRef(object):
    def __init__(
        self,
        local_object_reference: dict = None,
        optional: bool = None,
        name: str = None,
    ):
        self.local_object_reference = local_object_reference
        self.optional = optional
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersEnvFromConfigMapRef()
        if ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference.to_proto(
            resource.local_object_reference
        ):
            res.local_object_reference.CopyFrom(
                ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference.to_proto(
                    resource.local_object_reference
                )
            )
        else:
            res.ClearField("local_object_reference")
        if Primitive.to_proto(resource.optional):
            res.optional = Primitive.to_proto(resource.optional)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersEnvFromConfigMapRef(
            local_object_reference=ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference.from_proto(
                resource.local_object_reference
            ),
            optional=Primitive.from_proto(resource.optional),
            name=Primitive.from_proto(resource.name),
        )


class ServiceSpecTemplateSpecContainersEnvFromConfigMapRefArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersEnvFromConfigMapRef.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersEnvFromConfigMapRef.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_pb2.RunServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference(
            name=Primitive.from_proto(resource.name),
        )


class ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReferenceArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference.from_proto(
                i
            )
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersEnvFromSecretRef(object):
    def __init__(
        self,
        local_object_reference: dict = None,
        optional: bool = None,
        name: str = None,
    ):
        self.local_object_reference = local_object_reference
        self.optional = optional
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersEnvFromSecretRef()
        if ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference.to_proto(
            resource.local_object_reference
        ):
            res.local_object_reference.CopyFrom(
                ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference.to_proto(
                    resource.local_object_reference
                )
            )
        else:
            res.ClearField("local_object_reference")
        if Primitive.to_proto(resource.optional):
            res.optional = Primitive.to_proto(resource.optional)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersEnvFromSecretRef(
            local_object_reference=ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference.from_proto(
                resource.local_object_reference
            ),
            optional=Primitive.from_proto(resource.optional),
            name=Primitive.from_proto(resource.name),
        )


class ServiceSpecTemplateSpecContainersEnvFromSecretRefArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersEnvFromSecretRef.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersEnvFromSecretRef.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_pb2.RunServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference(
            name=Primitive.from_proto(resource.name),
        )


class ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReferenceArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference.from_proto(
                i
            )
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersVolumeMounts(object):
    def __init__(
        self,
        name: str = None,
        read_only: bool = None,
        mount_path: str = None,
        sub_path: str = None,
    ):
        self.name = name
        self.read_only = read_only
        self.mount_path = mount_path
        self.sub_path = sub_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersVolumeMounts()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.read_only):
            res.read_only = Primitive.to_proto(resource.read_only)
        if Primitive.to_proto(resource.mount_path):
            res.mount_path = Primitive.to_proto(resource.mount_path)
        if Primitive.to_proto(resource.sub_path):
            res.sub_path = Primitive.to_proto(resource.sub_path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersVolumeMounts(
            name=Primitive.from_proto(resource.name),
            read_only=Primitive.from_proto(resource.read_only),
            mount_path=Primitive.from_proto(resource.mount_path),
            sub_path=Primitive.from_proto(resource.sub_path),
        )


class ServiceSpecTemplateSpecContainersVolumeMountsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersVolumeMounts.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersVolumeMounts.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersLivenessProbe(object):
    def __init__(
        self,
        initial_delay_seconds: int = None,
        timeout_seconds: int = None,
        period_seconds: int = None,
        success_threshold: int = None,
        failure_threshold: int = None,
        exec: dict = None,
        http_get: dict = None,
        tcp_socket: dict = None,
    ):
        self.initial_delay_seconds = initial_delay_seconds
        self.timeout_seconds = timeout_seconds
        self.period_seconds = period_seconds
        self.success_threshold = success_threshold
        self.failure_threshold = failure_threshold
        self.exec = exec
        self.http_get = http_get
        self.tcp_socket = tcp_socket

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersLivenessProbe()
        if Primitive.to_proto(resource.initial_delay_seconds):
            res.initial_delay_seconds = Primitive.to_proto(
                resource.initial_delay_seconds
            )
        if Primitive.to_proto(resource.timeout_seconds):
            res.timeout_seconds = Primitive.to_proto(resource.timeout_seconds)
        if Primitive.to_proto(resource.period_seconds):
            res.period_seconds = Primitive.to_proto(resource.period_seconds)
        if Primitive.to_proto(resource.success_threshold):
            res.success_threshold = Primitive.to_proto(resource.success_threshold)
        if Primitive.to_proto(resource.failure_threshold):
            res.failure_threshold = Primitive.to_proto(resource.failure_threshold)
        if ServiceSpecTemplateSpecContainersLivenessProbeExec.to_proto(resource.exec):
            res.exec.CopyFrom(
                ServiceSpecTemplateSpecContainersLivenessProbeExec.to_proto(
                    resource.exec
                )
            )
        else:
            res.ClearField("exec")
        if ServiceSpecTemplateSpecContainersLivenessProbeHttpGet.to_proto(
            resource.http_get
        ):
            res.http_get.CopyFrom(
                ServiceSpecTemplateSpecContainersLivenessProbeHttpGet.to_proto(
                    resource.http_get
                )
            )
        else:
            res.ClearField("http_get")
        if ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket.to_proto(
            resource.tcp_socket
        ):
            res.tcp_socket.CopyFrom(
                ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket.to_proto(
                    resource.tcp_socket
                )
            )
        else:
            res.ClearField("tcp_socket")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersLivenessProbe(
            initial_delay_seconds=Primitive.from_proto(resource.initial_delay_seconds),
            timeout_seconds=Primitive.from_proto(resource.timeout_seconds),
            period_seconds=Primitive.from_proto(resource.period_seconds),
            success_threshold=Primitive.from_proto(resource.success_threshold),
            failure_threshold=Primitive.from_proto(resource.failure_threshold),
            exec=ServiceSpecTemplateSpecContainersLivenessProbeExec.from_proto(
                resource.exec
            ),
            http_get=ServiceSpecTemplateSpecContainersLivenessProbeHttpGet.from_proto(
                resource.http_get
            ),
            tcp_socket=ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket.from_proto(
                resource.tcp_socket
            ),
        )


class ServiceSpecTemplateSpecContainersLivenessProbeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersLivenessProbe.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersLivenessProbe.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersLivenessProbeExec(object):
    def __init__(self, command: str = None):
        self.command = command

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersLivenessProbeExec()
        if Primitive.to_proto(resource.command):
            res.command = Primitive.to_proto(resource.command)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersLivenessProbeExec(
            command=Primitive.from_proto(resource.command),
        )


class ServiceSpecTemplateSpecContainersLivenessProbeExecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersLivenessProbeExec.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersLivenessProbeExec.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersLivenessProbeHttpGet(object):
    def __init__(
        self,
        path: str = None,
        host: str = None,
        scheme: str = None,
        http_headers: list = None,
    ):
        self.path = path
        self.host = host
        self.scheme = scheme
        self.http_headers = http_headers

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersLivenessProbeHttpGet()
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        if Primitive.to_proto(resource.scheme):
            res.scheme = Primitive.to_proto(resource.scheme)
        if ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersArray.to_proto(
            resource.http_headers
        ):
            res.http_headers.extend(
                ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersArray.to_proto(
                    resource.http_headers
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersLivenessProbeHttpGet(
            path=Primitive.from_proto(resource.path),
            host=Primitive.from_proto(resource.host),
            scheme=Primitive.from_proto(resource.scheme),
            http_headers=ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersArray.from_proto(
                resource.http_headers
            ),
        )


class ServiceSpecTemplateSpecContainersLivenessProbeHttpGetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersLivenessProbeHttpGet.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersLivenessProbeHttpGet.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(object):
    def __init__(self, name: str = None, value: str = None):
        self.name = name
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_pb2.RunServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders(
            name=Primitive.from_proto(resource.name),
            value=Primitive.from_proto(resource.value),
        )


class ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeadersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders.from_proto(
                i
            )
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(object):
    def __init__(self, port: int = None, host: str = None):
        self.port = port
        self.host = host

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersLivenessProbeTcpSocket()
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket(
            port=Primitive.from_proto(resource.port),
            host=Primitive.from_proto(resource.host),
        )


class ServiceSpecTemplateSpecContainersLivenessProbeTcpSocketArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersLivenessProbeTcpSocket.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersReadinessProbe(object):
    def __init__(
        self,
        initial_delay_seconds: int = None,
        timeout_seconds: int = None,
        period_seconds: int = None,
        success_threshold: int = None,
        failure_threshold: int = None,
        exec: dict = None,
        http_get: dict = None,
        tcp_socket: dict = None,
    ):
        self.initial_delay_seconds = initial_delay_seconds
        self.timeout_seconds = timeout_seconds
        self.period_seconds = period_seconds
        self.success_threshold = success_threshold
        self.failure_threshold = failure_threshold
        self.exec = exec
        self.http_get = http_get
        self.tcp_socket = tcp_socket

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersReadinessProbe()
        if Primitive.to_proto(resource.initial_delay_seconds):
            res.initial_delay_seconds = Primitive.to_proto(
                resource.initial_delay_seconds
            )
        if Primitive.to_proto(resource.timeout_seconds):
            res.timeout_seconds = Primitive.to_proto(resource.timeout_seconds)
        if Primitive.to_proto(resource.period_seconds):
            res.period_seconds = Primitive.to_proto(resource.period_seconds)
        if Primitive.to_proto(resource.success_threshold):
            res.success_threshold = Primitive.to_proto(resource.success_threshold)
        if Primitive.to_proto(resource.failure_threshold):
            res.failure_threshold = Primitive.to_proto(resource.failure_threshold)
        if ServiceSpecTemplateSpecContainersReadinessProbeExec.to_proto(resource.exec):
            res.exec.CopyFrom(
                ServiceSpecTemplateSpecContainersReadinessProbeExec.to_proto(
                    resource.exec
                )
            )
        else:
            res.ClearField("exec")
        if ServiceSpecTemplateSpecContainersReadinessProbeHttpGet.to_proto(
            resource.http_get
        ):
            res.http_get.CopyFrom(
                ServiceSpecTemplateSpecContainersReadinessProbeHttpGet.to_proto(
                    resource.http_get
                )
            )
        else:
            res.ClearField("http_get")
        if ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket.to_proto(
            resource.tcp_socket
        ):
            res.tcp_socket.CopyFrom(
                ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket.to_proto(
                    resource.tcp_socket
                )
            )
        else:
            res.ClearField("tcp_socket")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersReadinessProbe(
            initial_delay_seconds=Primitive.from_proto(resource.initial_delay_seconds),
            timeout_seconds=Primitive.from_proto(resource.timeout_seconds),
            period_seconds=Primitive.from_proto(resource.period_seconds),
            success_threshold=Primitive.from_proto(resource.success_threshold),
            failure_threshold=Primitive.from_proto(resource.failure_threshold),
            exec=ServiceSpecTemplateSpecContainersReadinessProbeExec.from_proto(
                resource.exec
            ),
            http_get=ServiceSpecTemplateSpecContainersReadinessProbeHttpGet.from_proto(
                resource.http_get
            ),
            tcp_socket=ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket.from_proto(
                resource.tcp_socket
            ),
        )


class ServiceSpecTemplateSpecContainersReadinessProbeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersReadinessProbe.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersReadinessProbe.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersReadinessProbeExec(object):
    def __init__(self, command: str = None):
        self.command = command

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersReadinessProbeExec()
        if Primitive.to_proto(resource.command):
            res.command = Primitive.to_proto(resource.command)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersReadinessProbeExec(
            command=Primitive.from_proto(resource.command),
        )


class ServiceSpecTemplateSpecContainersReadinessProbeExecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersReadinessProbeExec.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersReadinessProbeExec.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersReadinessProbeHttpGet(object):
    def __init__(
        self,
        path: str = None,
        host: str = None,
        scheme: str = None,
        http_headers: list = None,
    ):
        self.path = path
        self.host = host
        self.scheme = scheme
        self.http_headers = http_headers

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersReadinessProbeHttpGet()
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        if Primitive.to_proto(resource.scheme):
            res.scheme = Primitive.to_proto(resource.scheme)
        if ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersArray.to_proto(
            resource.http_headers
        ):
            res.http_headers.extend(
                ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersArray.to_proto(
                    resource.http_headers
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersReadinessProbeHttpGet(
            path=Primitive.from_proto(resource.path),
            host=Primitive.from_proto(resource.host),
            scheme=Primitive.from_proto(resource.scheme),
            http_headers=ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersArray.from_proto(
                resource.http_headers
            ),
        )


class ServiceSpecTemplateSpecContainersReadinessProbeHttpGetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersReadinessProbeHttpGet.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersReadinessProbeHttpGet.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(object):
    def __init__(self, name: str = None, value: str = None):
        self.name = name
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_pb2.RunServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders(
            name=Primitive.from_proto(resource.name),
            value=Primitive.from_proto(resource.value),
        )


class ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeadersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders.from_proto(
                i
            )
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(object):
    def __init__(self, port: int = None, host: str = None):
        self.port = port
        self.host = host

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersReadinessProbeTcpSocket()
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket(
            port=Primitive.from_proto(resource.port),
            host=Primitive.from_proto(resource.host),
        )


class ServiceSpecTemplateSpecContainersReadinessProbeTcpSocketArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersReadinessProbeTcpSocket.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecContainersSecurityContext(object):
    def __init__(self, run_as_user: int = None):
        self.run_as_user = run_as_user

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecContainersSecurityContext()
        if Primitive.to_proto(resource.run_as_user):
            res.run_as_user = Primitive.to_proto(resource.run_as_user)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecContainersSecurityContext(
            run_as_user=Primitive.from_proto(resource.run_as_user),
        )


class ServiceSpecTemplateSpecContainersSecurityContextArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecContainersSecurityContext.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecContainersSecurityContext.from_proto(i)
            for i in resources
        ]


class ServiceSpecTemplateSpecVolumes(object):
    def __init__(self, name: str = None, secret: dict = None, config_map: dict = None):
        self.name = name
        self.secret = secret
        self.config_map = config_map

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecVolumes()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if ServiceSpecTemplateSpecVolumesSecret.to_proto(resource.secret):
            res.secret.CopyFrom(
                ServiceSpecTemplateSpecVolumesSecret.to_proto(resource.secret)
            )
        else:
            res.ClearField("secret")
        if ServiceSpecTemplateSpecVolumesConfigMap.to_proto(resource.config_map):
            res.config_map.CopyFrom(
                ServiceSpecTemplateSpecVolumesConfigMap.to_proto(resource.config_map)
            )
        else:
            res.ClearField("config_map")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecVolumes(
            name=Primitive.from_proto(resource.name),
            secret=ServiceSpecTemplateSpecVolumesSecret.from_proto(resource.secret),
            config_map=ServiceSpecTemplateSpecVolumesConfigMap.from_proto(
                resource.config_map
            ),
        )


class ServiceSpecTemplateSpecVolumesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceSpecTemplateSpecVolumes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceSpecTemplateSpecVolumes.from_proto(i) for i in resources]


class ServiceSpecTemplateSpecVolumesSecret(object):
    def __init__(
        self,
        secret_name: str = None,
        items: list = None,
        default_mode: int = None,
        optional: bool = None,
    ):
        self.secret_name = secret_name
        self.items = items
        self.default_mode = default_mode
        self.optional = optional

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecVolumesSecret()
        if Primitive.to_proto(resource.secret_name):
            res.secret_name = Primitive.to_proto(resource.secret_name)
        if ServiceSpecTemplateSpecVolumesSecretItemsArray.to_proto(resource.items):
            res.items.extend(
                ServiceSpecTemplateSpecVolumesSecretItemsArray.to_proto(resource.items)
            )
        if Primitive.to_proto(resource.default_mode):
            res.default_mode = Primitive.to_proto(resource.default_mode)
        if Primitive.to_proto(resource.optional):
            res.optional = Primitive.to_proto(resource.optional)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecVolumesSecret(
            secret_name=Primitive.from_proto(resource.secret_name),
            items=ServiceSpecTemplateSpecVolumesSecretItemsArray.from_proto(
                resource.items
            ),
            default_mode=Primitive.from_proto(resource.default_mode),
            optional=Primitive.from_proto(resource.optional),
        )


class ServiceSpecTemplateSpecVolumesSecretArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceSpecTemplateSpecVolumesSecret.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceSpecTemplateSpecVolumesSecret.from_proto(i) for i in resources]


class ServiceSpecTemplateSpecVolumesSecretItems(object):
    def __init__(self, key: str = None, path: str = None, mode: int = None):
        self.key = key
        self.path = path
        self.mode = mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecVolumesSecretItems()
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.mode):
            res.mode = Primitive.to_proto(resource.mode)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecVolumesSecretItems(
            key=Primitive.from_proto(resource.key),
            path=Primitive.from_proto(resource.path),
            mode=Primitive.from_proto(resource.mode),
        )


class ServiceSpecTemplateSpecVolumesSecretItemsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecVolumesSecretItems.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecVolumesSecretItems.from_proto(i) for i in resources
        ]


class ServiceSpecTemplateSpecVolumesConfigMap(object):
    def __init__(
        self,
        name: str = None,
        items: list = None,
        default_mode: int = None,
        optional: bool = None,
    ):
        self.name = name
        self.items = items
        self.default_mode = default_mode
        self.optional = optional

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecVolumesConfigMap()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if ServiceSpecTemplateSpecVolumesConfigMapItemsArray.to_proto(resource.items):
            res.items.extend(
                ServiceSpecTemplateSpecVolumesConfigMapItemsArray.to_proto(
                    resource.items
                )
            )
        if Primitive.to_proto(resource.default_mode):
            res.default_mode = Primitive.to_proto(resource.default_mode)
        if Primitive.to_proto(resource.optional):
            res.optional = Primitive.to_proto(resource.optional)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecVolumesConfigMap(
            name=Primitive.from_proto(resource.name),
            items=ServiceSpecTemplateSpecVolumesConfigMapItemsArray.from_proto(
                resource.items
            ),
            default_mode=Primitive.from_proto(resource.default_mode),
            optional=Primitive.from_proto(resource.optional),
        )


class ServiceSpecTemplateSpecVolumesConfigMapArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceSpecTemplateSpecVolumesConfigMap.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecVolumesConfigMap.from_proto(i) for i in resources
        ]


class ServiceSpecTemplateSpecVolumesConfigMapItems(object):
    def __init__(self, key: str = None, path: str = None, mode: int = None):
        self.key = key
        self.path = path
        self.mode = mode

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTemplateSpecVolumesConfigMapItems()
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.mode):
            res.mode = Primitive.to_proto(resource.mode)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTemplateSpecVolumesConfigMapItems(
            key=Primitive.from_proto(resource.key),
            path=Primitive.from_proto(resource.path),
            mode=Primitive.from_proto(resource.mode),
        )


class ServiceSpecTemplateSpecVolumesConfigMapItemsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceSpecTemplateSpecVolumesConfigMapItems.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceSpecTemplateSpecVolumesConfigMapItems.from_proto(i)
            for i in resources
        ]


class ServiceSpecTraffic(object):
    def __init__(
        self,
        configuration_name: str = None,
        revision_name: str = None,
        percent: int = None,
        tag: str = None,
        latest_revision: bool = None,
        url: str = None,
    ):
        self.configuration_name = configuration_name
        self.revision_name = revision_name
        self.percent = percent
        self.tag = tag
        self.latest_revision = latest_revision
        self.url = url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceSpecTraffic()
        if Primitive.to_proto(resource.configuration_name):
            res.configuration_name = Primitive.to_proto(resource.configuration_name)
        if Primitive.to_proto(resource.revision_name):
            res.revision_name = Primitive.to_proto(resource.revision_name)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.latest_revision):
            res.latest_revision = Primitive.to_proto(resource.latest_revision)
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceSpecTraffic(
            configuration_name=Primitive.from_proto(resource.configuration_name),
            revision_name=Primitive.from_proto(resource.revision_name),
            percent=Primitive.from_proto(resource.percent),
            tag=Primitive.from_proto(resource.tag),
            latest_revision=Primitive.from_proto(resource.latest_revision),
            url=Primitive.from_proto(resource.url),
        )


class ServiceSpecTrafficArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceSpecTraffic.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceSpecTraffic.from_proto(i) for i in resources]


class ServiceStatus(object):
    def __init__(
        self,
        observed_generation: int = None,
        conditions: list = None,
        latest_ready_revision_name: str = None,
        latest_created_revision_name: str = None,
        traffic: list = None,
        url: str = None,
        address: dict = None,
    ):
        self.observed_generation = observed_generation
        self.conditions = conditions
        self.latest_ready_revision_name = latest_ready_revision_name
        self.latest_created_revision_name = latest_created_revision_name
        self.traffic = traffic
        self.url = url
        self.address = address

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceStatus()
        if Primitive.to_proto(resource.observed_generation):
            res.observed_generation = Primitive.to_proto(resource.observed_generation)
        if ServiceStatusConditionsArray.to_proto(resource.conditions):
            res.conditions.extend(
                ServiceStatusConditionsArray.to_proto(resource.conditions)
            )
        if Primitive.to_proto(resource.latest_ready_revision_name):
            res.latest_ready_revision_name = Primitive.to_proto(
                resource.latest_ready_revision_name
            )
        if Primitive.to_proto(resource.latest_created_revision_name):
            res.latest_created_revision_name = Primitive.to_proto(
                resource.latest_created_revision_name
            )
        if ServiceStatusTrafficArray.to_proto(resource.traffic):
            res.traffic.extend(ServiceStatusTrafficArray.to_proto(resource.traffic))
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        if ServiceStatusAddress.to_proto(resource.address):
            res.address.CopyFrom(ServiceStatusAddress.to_proto(resource.address))
        else:
            res.ClearField("address")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceStatus(
            observed_generation=Primitive.from_proto(resource.observed_generation),
            conditions=ServiceStatusConditionsArray.from_proto(resource.conditions),
            latest_ready_revision_name=Primitive.from_proto(
                resource.latest_ready_revision_name
            ),
            latest_created_revision_name=Primitive.from_proto(
                resource.latest_created_revision_name
            ),
            traffic=ServiceStatusTrafficArray.from_proto(resource.traffic),
            url=Primitive.from_proto(resource.url),
            address=ServiceStatusAddress.from_proto(resource.address),
        )


class ServiceStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceStatus.from_proto(i) for i in resources]


class ServiceStatusConditions(object):
    def __init__(
        self,
        type: str = None,
        status: str = None,
        reason: str = None,
        message: str = None,
        last_transition_time: dict = None,
        severity: str = None,
    ):
        self.type = type
        self.status = status
        self.reason = reason
        self.message = message
        self.last_transition_time = last_transition_time
        self.severity = severity

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceStatusConditions()
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.status):
            res.status = Primitive.to_proto(resource.status)
        if Primitive.to_proto(resource.reason):
            res.reason = Primitive.to_proto(resource.reason)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if ServiceStatusConditionsLastTransitionTime.to_proto(
            resource.last_transition_time
        ):
            res.last_transition_time.CopyFrom(
                ServiceStatusConditionsLastTransitionTime.to_proto(
                    resource.last_transition_time
                )
            )
        else:
            res.ClearField("last_transition_time")
        if Primitive.to_proto(resource.severity):
            res.severity = Primitive.to_proto(resource.severity)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceStatusConditions(
            type=Primitive.from_proto(resource.type),
            status=Primitive.from_proto(resource.status),
            reason=Primitive.from_proto(resource.reason),
            message=Primitive.from_proto(resource.message),
            last_transition_time=ServiceStatusConditionsLastTransitionTime.from_proto(
                resource.last_transition_time
            ),
            severity=Primitive.from_proto(resource.severity),
        )


class ServiceStatusConditionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceStatusConditions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceStatusConditions.from_proto(i) for i in resources]


class ServiceStatusConditionsLastTransitionTime(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceStatusConditionsLastTransitionTime()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceStatusConditionsLastTransitionTime(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class ServiceStatusConditionsLastTransitionTimeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceStatusConditionsLastTransitionTime.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceStatusConditionsLastTransitionTime.from_proto(i) for i in resources
        ]


class ServiceStatusTraffic(object):
    def __init__(
        self,
        configuration_name: str = None,
        revision_name: str = None,
        percent: int = None,
        tag: str = None,
        latest_revision: bool = None,
        url: str = None,
    ):
        self.configuration_name = configuration_name
        self.revision_name = revision_name
        self.percent = percent
        self.tag = tag
        self.latest_revision = latest_revision
        self.url = url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceStatusTraffic()
        if Primitive.to_proto(resource.configuration_name):
            res.configuration_name = Primitive.to_proto(resource.configuration_name)
        if Primitive.to_proto(resource.revision_name):
            res.revision_name = Primitive.to_proto(resource.revision_name)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.latest_revision):
            res.latest_revision = Primitive.to_proto(resource.latest_revision)
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceStatusTraffic(
            configuration_name=Primitive.from_proto(resource.configuration_name),
            revision_name=Primitive.from_proto(resource.revision_name),
            percent=Primitive.from_proto(resource.percent),
            tag=Primitive.from_proto(resource.tag),
            latest_revision=Primitive.from_proto(resource.latest_revision),
            url=Primitive.from_proto(resource.url),
        )


class ServiceStatusTrafficArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceStatusTraffic.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceStatusTraffic.from_proto(i) for i in resources]


class ServiceStatusAddress(object):
    def __init__(self, url: str = None):
        self.url = url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.RunServiceStatusAddress()
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceStatusAddress(url=Primitive.from_proto(resource.url),)


class ServiceStatusAddressArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceStatusAddress.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceStatusAddress.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
