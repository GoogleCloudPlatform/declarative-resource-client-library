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
from google3.cloud.graphite.mmv2.services.google.monitoring import service_pb2
from google3.cloud.graphite.mmv2.services.google.monitoring import service_pb2_grpc

from typing import List


class Service(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        custom: dict = None,
        app_engine: dict = None,
        cloud_endpoints: dict = None,
        cluster_istio: dict = None,
        mesh_istio: dict = None,
        istio_canonical_service: dict = None,
        cloud_run: dict = None,
        gke_namespace: dict = None,
        gke_workload: dict = None,
        gke_service: dict = None,
        telemetry: dict = None,
        user_labels: dict = None,
        deleted: bool = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.custom = custom
        self.app_engine = app_engine
        self.cloud_endpoints = cloud_endpoints
        self.cluster_istio = cluster_istio
        self.mesh_istio = mesh_istio
        self.istio_canonical_service = istio_canonical_service
        self.cloud_run = cloud_run
        self.gke_namespace = gke_namespace
        self.gke_workload = gke_workload
        self.gke_service = gke_service
        self.telemetry = telemetry
        self.user_labels = user_labels
        self.deleted = deleted
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = service_pb2_grpc.MonitoringAlphaServiceServiceStub(channel.Channel())
        request = service_pb2.ApplyMonitoringAlphaServiceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if ServiceCustom.to_proto(self.custom):
            request.resource.custom.CopyFrom(ServiceCustom.to_proto(self.custom))
        else:
            request.resource.ClearField("custom")
        if ServiceAppEngine.to_proto(self.app_engine):
            request.resource.app_engine.CopyFrom(
                ServiceAppEngine.to_proto(self.app_engine)
            )
        else:
            request.resource.ClearField("app_engine")
        if ServiceCloudEndpoints.to_proto(self.cloud_endpoints):
            request.resource.cloud_endpoints.CopyFrom(
                ServiceCloudEndpoints.to_proto(self.cloud_endpoints)
            )
        else:
            request.resource.ClearField("cloud_endpoints")
        if ServiceClusterIstio.to_proto(self.cluster_istio):
            request.resource.cluster_istio.CopyFrom(
                ServiceClusterIstio.to_proto(self.cluster_istio)
            )
        else:
            request.resource.ClearField("cluster_istio")
        if ServiceMeshIstio.to_proto(self.mesh_istio):
            request.resource.mesh_istio.CopyFrom(
                ServiceMeshIstio.to_proto(self.mesh_istio)
            )
        else:
            request.resource.ClearField("mesh_istio")
        if ServiceIstioCanonicalService.to_proto(self.istio_canonical_service):
            request.resource.istio_canonical_service.CopyFrom(
                ServiceIstioCanonicalService.to_proto(self.istio_canonical_service)
            )
        else:
            request.resource.ClearField("istio_canonical_service")
        if ServiceCloudRun.to_proto(self.cloud_run):
            request.resource.cloud_run.CopyFrom(
                ServiceCloudRun.to_proto(self.cloud_run)
            )
        else:
            request.resource.ClearField("cloud_run")
        if ServiceGkeNamespace.to_proto(self.gke_namespace):
            request.resource.gke_namespace.CopyFrom(
                ServiceGkeNamespace.to_proto(self.gke_namespace)
            )
        else:
            request.resource.ClearField("gke_namespace")
        if ServiceGkeWorkload.to_proto(self.gke_workload):
            request.resource.gke_workload.CopyFrom(
                ServiceGkeWorkload.to_proto(self.gke_workload)
            )
        else:
            request.resource.ClearField("gke_workload")
        if ServiceGkeService.to_proto(self.gke_service):
            request.resource.gke_service.CopyFrom(
                ServiceGkeService.to_proto(self.gke_service)
            )
        else:
            request.resource.ClearField("gke_service")
        if ServiceTelemetry.to_proto(self.telemetry):
            request.resource.telemetry.CopyFrom(
                ServiceTelemetry.to_proto(self.telemetry)
            )
        else:
            request.resource.ClearField("telemetry")
        if Primitive.to_proto(self.user_labels):
            request.resource.user_labels = Primitive.to_proto(self.user_labels)

        if Primitive.to_proto(self.deleted):
            request.resource.deleted = Primitive.to_proto(self.deleted)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyMonitoringAlphaService(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.custom = ServiceCustom.from_proto(response.custom)
        self.app_engine = ServiceAppEngine.from_proto(response.app_engine)
        self.cloud_endpoints = ServiceCloudEndpoints.from_proto(
            response.cloud_endpoints
        )
        self.cluster_istio = ServiceClusterIstio.from_proto(response.cluster_istio)
        self.mesh_istio = ServiceMeshIstio.from_proto(response.mesh_istio)
        self.istio_canonical_service = ServiceIstioCanonicalService.from_proto(
            response.istio_canonical_service
        )
        self.cloud_run = ServiceCloudRun.from_proto(response.cloud_run)
        self.gke_namespace = ServiceGkeNamespace.from_proto(response.gke_namespace)
        self.gke_workload = ServiceGkeWorkload.from_proto(response.gke_workload)
        self.gke_service = ServiceGkeService.from_proto(response.gke_service)
        self.telemetry = ServiceTelemetry.from_proto(response.telemetry)
        self.user_labels = Primitive.from_proto(response.user_labels)
        self.deleted = Primitive.from_proto(response.deleted)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = service_pb2_grpc.MonitoringAlphaServiceServiceStub(channel.Channel())
        request = service_pb2.DeleteMonitoringAlphaServiceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if ServiceCustom.to_proto(self.custom):
            request.resource.custom.CopyFrom(ServiceCustom.to_proto(self.custom))
        else:
            request.resource.ClearField("custom")
        if ServiceAppEngine.to_proto(self.app_engine):
            request.resource.app_engine.CopyFrom(
                ServiceAppEngine.to_proto(self.app_engine)
            )
        else:
            request.resource.ClearField("app_engine")
        if ServiceCloudEndpoints.to_proto(self.cloud_endpoints):
            request.resource.cloud_endpoints.CopyFrom(
                ServiceCloudEndpoints.to_proto(self.cloud_endpoints)
            )
        else:
            request.resource.ClearField("cloud_endpoints")
        if ServiceClusterIstio.to_proto(self.cluster_istio):
            request.resource.cluster_istio.CopyFrom(
                ServiceClusterIstio.to_proto(self.cluster_istio)
            )
        else:
            request.resource.ClearField("cluster_istio")
        if ServiceMeshIstio.to_proto(self.mesh_istio):
            request.resource.mesh_istio.CopyFrom(
                ServiceMeshIstio.to_proto(self.mesh_istio)
            )
        else:
            request.resource.ClearField("mesh_istio")
        if ServiceIstioCanonicalService.to_proto(self.istio_canonical_service):
            request.resource.istio_canonical_service.CopyFrom(
                ServiceIstioCanonicalService.to_proto(self.istio_canonical_service)
            )
        else:
            request.resource.ClearField("istio_canonical_service")
        if ServiceCloudRun.to_proto(self.cloud_run):
            request.resource.cloud_run.CopyFrom(
                ServiceCloudRun.to_proto(self.cloud_run)
            )
        else:
            request.resource.ClearField("cloud_run")
        if ServiceGkeNamespace.to_proto(self.gke_namespace):
            request.resource.gke_namespace.CopyFrom(
                ServiceGkeNamespace.to_proto(self.gke_namespace)
            )
        else:
            request.resource.ClearField("gke_namespace")
        if ServiceGkeWorkload.to_proto(self.gke_workload):
            request.resource.gke_workload.CopyFrom(
                ServiceGkeWorkload.to_proto(self.gke_workload)
            )
        else:
            request.resource.ClearField("gke_workload")
        if ServiceGkeService.to_proto(self.gke_service):
            request.resource.gke_service.CopyFrom(
                ServiceGkeService.to_proto(self.gke_service)
            )
        else:
            request.resource.ClearField("gke_service")
        if ServiceTelemetry.to_proto(self.telemetry):
            request.resource.telemetry.CopyFrom(
                ServiceTelemetry.to_proto(self.telemetry)
            )
        else:
            request.resource.ClearField("telemetry")
        if Primitive.to_proto(self.user_labels):
            request.resource.user_labels = Primitive.to_proto(self.user_labels)

        if Primitive.to_proto(self.deleted):
            request.resource.deleted = Primitive.to_proto(self.deleted)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteMonitoringAlphaService(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = service_pb2_grpc.MonitoringAlphaServiceServiceStub(channel.Channel())
        request = service_pb2.ListMonitoringAlphaServiceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListMonitoringAlphaService(request).items

    def to_proto(self):
        resource = service_pb2.MonitoringAlphaService()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if ServiceCustom.to_proto(self.custom):
            resource.custom.CopyFrom(ServiceCustom.to_proto(self.custom))
        else:
            resource.ClearField("custom")
        if ServiceAppEngine.to_proto(self.app_engine):
            resource.app_engine.CopyFrom(ServiceAppEngine.to_proto(self.app_engine))
        else:
            resource.ClearField("app_engine")
        if ServiceCloudEndpoints.to_proto(self.cloud_endpoints):
            resource.cloud_endpoints.CopyFrom(
                ServiceCloudEndpoints.to_proto(self.cloud_endpoints)
            )
        else:
            resource.ClearField("cloud_endpoints")
        if ServiceClusterIstio.to_proto(self.cluster_istio):
            resource.cluster_istio.CopyFrom(
                ServiceClusterIstio.to_proto(self.cluster_istio)
            )
        else:
            resource.ClearField("cluster_istio")
        if ServiceMeshIstio.to_proto(self.mesh_istio):
            resource.mesh_istio.CopyFrom(ServiceMeshIstio.to_proto(self.mesh_istio))
        else:
            resource.ClearField("mesh_istio")
        if ServiceIstioCanonicalService.to_proto(self.istio_canonical_service):
            resource.istio_canonical_service.CopyFrom(
                ServiceIstioCanonicalService.to_proto(self.istio_canonical_service)
            )
        else:
            resource.ClearField("istio_canonical_service")
        if ServiceCloudRun.to_proto(self.cloud_run):
            resource.cloud_run.CopyFrom(ServiceCloudRun.to_proto(self.cloud_run))
        else:
            resource.ClearField("cloud_run")
        if ServiceGkeNamespace.to_proto(self.gke_namespace):
            resource.gke_namespace.CopyFrom(
                ServiceGkeNamespace.to_proto(self.gke_namespace)
            )
        else:
            resource.ClearField("gke_namespace")
        if ServiceGkeWorkload.to_proto(self.gke_workload):
            resource.gke_workload.CopyFrom(
                ServiceGkeWorkload.to_proto(self.gke_workload)
            )
        else:
            resource.ClearField("gke_workload")
        if ServiceGkeService.to_proto(self.gke_service):
            resource.gke_service.CopyFrom(ServiceGkeService.to_proto(self.gke_service))
        else:
            resource.ClearField("gke_service")
        if ServiceTelemetry.to_proto(self.telemetry):
            resource.telemetry.CopyFrom(ServiceTelemetry.to_proto(self.telemetry))
        else:
            resource.ClearField("telemetry")
        if Primitive.to_proto(self.user_labels):
            resource.user_labels = Primitive.to_proto(self.user_labels)
        if Primitive.to_proto(self.deleted):
            resource.deleted = Primitive.to_proto(self.deleted)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class ServiceCustom(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.MonitoringAlphaServiceCustom()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceCustom()


class ServiceCustomArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceCustom.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceCustom.from_proto(i) for i in resources]


class ServiceAppEngine(object):
    def __init__(self, module_id: str = None):
        self.module_id = module_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.MonitoringAlphaServiceAppEngine()
        if Primitive.to_proto(resource.module_id):
            res.module_id = Primitive.to_proto(resource.module_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceAppEngine(module_id=Primitive.from_proto(resource.module_id),)


class ServiceAppEngineArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceAppEngine.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceAppEngine.from_proto(i) for i in resources]


class ServiceCloudEndpoints(object):
    def __init__(self, service: str = None):
        self.service = service

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.MonitoringAlphaServiceCloudEndpoints()
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceCloudEndpoints(service=Primitive.from_proto(resource.service),)


class ServiceCloudEndpointsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceCloudEndpoints.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceCloudEndpoints.from_proto(i) for i in resources]


class ServiceClusterIstio(object):
    def __init__(
        self,
        location: str = None,
        cluster_name: str = None,
        service_namespace: str = None,
        service_name: str = None,
    ):
        self.location = location
        self.cluster_name = cluster_name
        self.service_namespace = service_namespace
        self.service_name = service_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.MonitoringAlphaServiceClusterIstio()
        if Primitive.to_proto(resource.location):
            res.location = Primitive.to_proto(resource.location)
        if Primitive.to_proto(resource.cluster_name):
            res.cluster_name = Primitive.to_proto(resource.cluster_name)
        if Primitive.to_proto(resource.service_namespace):
            res.service_namespace = Primitive.to_proto(resource.service_namespace)
        if Primitive.to_proto(resource.service_name):
            res.service_name = Primitive.to_proto(resource.service_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceClusterIstio(
            location=Primitive.from_proto(resource.location),
            cluster_name=Primitive.from_proto(resource.cluster_name),
            service_namespace=Primitive.from_proto(resource.service_namespace),
            service_name=Primitive.from_proto(resource.service_name),
        )


class ServiceClusterIstioArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceClusterIstio.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceClusterIstio.from_proto(i) for i in resources]


class ServiceMeshIstio(object):
    def __init__(
        self,
        mesh_uid: str = None,
        service_namespace: str = None,
        service_name: str = None,
    ):
        self.mesh_uid = mesh_uid
        self.service_namespace = service_namespace
        self.service_name = service_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.MonitoringAlphaServiceMeshIstio()
        if Primitive.to_proto(resource.mesh_uid):
            res.mesh_uid = Primitive.to_proto(resource.mesh_uid)
        if Primitive.to_proto(resource.service_namespace):
            res.service_namespace = Primitive.to_proto(resource.service_namespace)
        if Primitive.to_proto(resource.service_name):
            res.service_name = Primitive.to_proto(resource.service_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceMeshIstio(
            mesh_uid=Primitive.from_proto(resource.mesh_uid),
            service_namespace=Primitive.from_proto(resource.service_namespace),
            service_name=Primitive.from_proto(resource.service_name),
        )


class ServiceMeshIstioArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceMeshIstio.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceMeshIstio.from_proto(i) for i in resources]


class ServiceIstioCanonicalService(object):
    def __init__(
        self,
        mesh_uid: str = None,
        canonical_service_namespace: str = None,
        canonical_service: str = None,
    ):
        self.mesh_uid = mesh_uid
        self.canonical_service_namespace = canonical_service_namespace
        self.canonical_service = canonical_service

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.MonitoringAlphaServiceIstioCanonicalService()
        if Primitive.to_proto(resource.mesh_uid):
            res.mesh_uid = Primitive.to_proto(resource.mesh_uid)
        if Primitive.to_proto(resource.canonical_service_namespace):
            res.canonical_service_namespace = Primitive.to_proto(
                resource.canonical_service_namespace
            )
        if Primitive.to_proto(resource.canonical_service):
            res.canonical_service = Primitive.to_proto(resource.canonical_service)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceIstioCanonicalService(
            mesh_uid=Primitive.from_proto(resource.mesh_uid),
            canonical_service_namespace=Primitive.from_proto(
                resource.canonical_service_namespace
            ),
            canonical_service=Primitive.from_proto(resource.canonical_service),
        )


class ServiceIstioCanonicalServiceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceIstioCanonicalService.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceIstioCanonicalService.from_proto(i) for i in resources]


class ServiceCloudRun(object):
    def __init__(self, service_name: str = None, location: str = None):
        self.service_name = service_name
        self.location = location

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.MonitoringAlphaServiceCloudRun()
        if Primitive.to_proto(resource.service_name):
            res.service_name = Primitive.to_proto(resource.service_name)
        if Primitive.to_proto(resource.location):
            res.location = Primitive.to_proto(resource.location)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceCloudRun(
            service_name=Primitive.from_proto(resource.service_name),
            location=Primitive.from_proto(resource.location),
        )


class ServiceCloudRunArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceCloudRun.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceCloudRun.from_proto(i) for i in resources]


class ServiceGkeNamespace(object):
    def __init__(
        self,
        project_id: str = None,
        location: str = None,
        cluster_name: str = None,
        namespace_name: str = None,
    ):
        self.project_id = project_id
        self.location = location
        self.cluster_name = cluster_name
        self.namespace_name = namespace_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.MonitoringAlphaServiceGkeNamespace()
        if Primitive.to_proto(resource.project_id):
            res.project_id = Primitive.to_proto(resource.project_id)
        if Primitive.to_proto(resource.location):
            res.location = Primitive.to_proto(resource.location)
        if Primitive.to_proto(resource.cluster_name):
            res.cluster_name = Primitive.to_proto(resource.cluster_name)
        if Primitive.to_proto(resource.namespace_name):
            res.namespace_name = Primitive.to_proto(resource.namespace_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceGkeNamespace(
            project_id=Primitive.from_proto(resource.project_id),
            location=Primitive.from_proto(resource.location),
            cluster_name=Primitive.from_proto(resource.cluster_name),
            namespace_name=Primitive.from_proto(resource.namespace_name),
        )


class ServiceGkeNamespaceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceGkeNamespace.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceGkeNamespace.from_proto(i) for i in resources]


class ServiceGkeWorkload(object):
    def __init__(
        self,
        project_id: str = None,
        location: str = None,
        cluster_name: str = None,
        namespace_name: str = None,
        top_level_controller_type: str = None,
        top_level_controller_name: str = None,
    ):
        self.project_id = project_id
        self.location = location
        self.cluster_name = cluster_name
        self.namespace_name = namespace_name
        self.top_level_controller_type = top_level_controller_type
        self.top_level_controller_name = top_level_controller_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.MonitoringAlphaServiceGkeWorkload()
        if Primitive.to_proto(resource.project_id):
            res.project_id = Primitive.to_proto(resource.project_id)
        if Primitive.to_proto(resource.location):
            res.location = Primitive.to_proto(resource.location)
        if Primitive.to_proto(resource.cluster_name):
            res.cluster_name = Primitive.to_proto(resource.cluster_name)
        if Primitive.to_proto(resource.namespace_name):
            res.namespace_name = Primitive.to_proto(resource.namespace_name)
        if Primitive.to_proto(resource.top_level_controller_type):
            res.top_level_controller_type = Primitive.to_proto(
                resource.top_level_controller_type
            )
        if Primitive.to_proto(resource.top_level_controller_name):
            res.top_level_controller_name = Primitive.to_proto(
                resource.top_level_controller_name
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceGkeWorkload(
            project_id=Primitive.from_proto(resource.project_id),
            location=Primitive.from_proto(resource.location),
            cluster_name=Primitive.from_proto(resource.cluster_name),
            namespace_name=Primitive.from_proto(resource.namespace_name),
            top_level_controller_type=Primitive.from_proto(
                resource.top_level_controller_type
            ),
            top_level_controller_name=Primitive.from_proto(
                resource.top_level_controller_name
            ),
        )


class ServiceGkeWorkloadArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceGkeWorkload.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceGkeWorkload.from_proto(i) for i in resources]


class ServiceGkeService(object):
    def __init__(
        self,
        project_id: str = None,
        location: str = None,
        cluster_name: str = None,
        namespace_name: str = None,
        service_name: str = None,
    ):
        self.project_id = project_id
        self.location = location
        self.cluster_name = cluster_name
        self.namespace_name = namespace_name
        self.service_name = service_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.MonitoringAlphaServiceGkeService()
        if Primitive.to_proto(resource.project_id):
            res.project_id = Primitive.to_proto(resource.project_id)
        if Primitive.to_proto(resource.location):
            res.location = Primitive.to_proto(resource.location)
        if Primitive.to_proto(resource.cluster_name):
            res.cluster_name = Primitive.to_proto(resource.cluster_name)
        if Primitive.to_proto(resource.namespace_name):
            res.namespace_name = Primitive.to_proto(resource.namespace_name)
        if Primitive.to_proto(resource.service_name):
            res.service_name = Primitive.to_proto(resource.service_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceGkeService(
            project_id=Primitive.from_proto(resource.project_id),
            location=Primitive.from_proto(resource.location),
            cluster_name=Primitive.from_proto(resource.cluster_name),
            namespace_name=Primitive.from_proto(resource.namespace_name),
            service_name=Primitive.from_proto(resource.service_name),
        )


class ServiceGkeServiceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceGkeService.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceGkeService.from_proto(i) for i in resources]


class ServiceTelemetry(object):
    def __init__(self, resource_name: str = None):
        self.resource_name = resource_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = service_pb2.MonitoringAlphaServiceTelemetry()
        if Primitive.to_proto(resource.resource_name):
            res.resource_name = Primitive.to_proto(resource.resource_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceTelemetry(
            resource_name=Primitive.from_proto(resource.resource_name),
        )


class ServiceTelemetryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ServiceTelemetry.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ServiceTelemetry.from_proto(i) for i in resources]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
