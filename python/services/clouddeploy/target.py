# Copyright 2025 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.clouddeploy import target_pb2
from google3.cloud.graphite.mmv2.services.google.clouddeploy import target_pb2_grpc

from typing import List


class Target(object):
    def __init__(
        self,
        name: str = None,
        target_id: str = None,
        uid: str = None,
        description: str = None,
        annotations: dict = None,
        labels: dict = None,
        require_approval: bool = None,
        create_time: str = None,
        update_time: str = None,
        gke: dict = None,
        anthos_cluster: dict = None,
        etag: str = None,
        execution_configs: list = None,
        project: str = None,
        location: str = None,
        run: dict = None,
        multi_target: dict = None,
        deploy_parameters: dict = None,
        custom_target: dict = None,
        associated_entities: dict = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.annotations = annotations
        self.labels = labels
        self.require_approval = require_approval
        self.gke = gke
        self.anthos_cluster = anthos_cluster
        self.execution_configs = execution_configs
        self.project = project
        self.location = location
        self.run = run
        self.multi_target = multi_target
        self.deploy_parameters = deploy_parameters
        self.custom_target = custom_target
        self.associated_entities = associated_entities
        self.service_account_file = service_account_file

    def apply(self):
        stub = target_pb2_grpc.ClouddeployTargetServiceStub(channel.Channel())
        request = target_pb2.ApplyClouddeployTargetRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.require_approval):
            request.resource.require_approval = Primitive.to_proto(
                self.require_approval
            )

        if TargetGke.to_proto(self.gke):
            request.resource.gke.CopyFrom(TargetGke.to_proto(self.gke))
        else:
            request.resource.ClearField("gke")
        if TargetAnthosCluster.to_proto(self.anthos_cluster):
            request.resource.anthos_cluster.CopyFrom(
                TargetAnthosCluster.to_proto(self.anthos_cluster)
            )
        else:
            request.resource.ClearField("anthos_cluster")
        if TargetExecutionConfigsArray.to_proto(self.execution_configs):
            request.resource.execution_configs.extend(
                TargetExecutionConfigsArray.to_proto(self.execution_configs)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if TargetRun.to_proto(self.run):
            request.resource.run.CopyFrom(TargetRun.to_proto(self.run))
        else:
            request.resource.ClearField("run")
        if TargetMultiTarget.to_proto(self.multi_target):
            request.resource.multi_target.CopyFrom(
                TargetMultiTarget.to_proto(self.multi_target)
            )
        else:
            request.resource.ClearField("multi_target")
        if Primitive.to_proto(self.deploy_parameters):
            request.resource.deploy_parameters = Primitive.to_proto(
                self.deploy_parameters
            )

        if TargetCustomTarget.to_proto(self.custom_target):
            request.resource.custom_target.CopyFrom(
                TargetCustomTarget.to_proto(self.custom_target)
            )
        else:
            request.resource.ClearField("custom_target")
        if Primitive.to_proto(self.associated_entities):
            request.resource.associated_entities = Primitive.to_proto(
                self.associated_entities
            )

        request.service_account_file = self.service_account_file

        response = stub.ApplyClouddeployTarget(request)
        self.name = Primitive.from_proto(response.name)
        self.target_id = Primitive.from_proto(response.target_id)
        self.uid = Primitive.from_proto(response.uid)
        self.description = Primitive.from_proto(response.description)
        self.annotations = Primitive.from_proto(response.annotations)
        self.labels = Primitive.from_proto(response.labels)
        self.require_approval = Primitive.from_proto(response.require_approval)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.gke = TargetGke.from_proto(response.gke)
        self.anthos_cluster = TargetAnthosCluster.from_proto(response.anthos_cluster)
        self.etag = Primitive.from_proto(response.etag)
        self.execution_configs = TargetExecutionConfigsArray.from_proto(
            response.execution_configs
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.run = TargetRun.from_proto(response.run)
        self.multi_target = TargetMultiTarget.from_proto(response.multi_target)
        self.deploy_parameters = Primitive.from_proto(response.deploy_parameters)
        self.custom_target = TargetCustomTarget.from_proto(response.custom_target)
        self.associated_entities = Primitive.from_proto(response.associated_entities)

    def delete(self):
        stub = target_pb2_grpc.ClouddeployTargetServiceStub(channel.Channel())
        request = target_pb2.DeleteClouddeployTargetRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.require_approval):
            request.resource.require_approval = Primitive.to_proto(
                self.require_approval
            )

        if TargetGke.to_proto(self.gke):
            request.resource.gke.CopyFrom(TargetGke.to_proto(self.gke))
        else:
            request.resource.ClearField("gke")
        if TargetAnthosCluster.to_proto(self.anthos_cluster):
            request.resource.anthos_cluster.CopyFrom(
                TargetAnthosCluster.to_proto(self.anthos_cluster)
            )
        else:
            request.resource.ClearField("anthos_cluster")
        if TargetExecutionConfigsArray.to_proto(self.execution_configs):
            request.resource.execution_configs.extend(
                TargetExecutionConfigsArray.to_proto(self.execution_configs)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if TargetRun.to_proto(self.run):
            request.resource.run.CopyFrom(TargetRun.to_proto(self.run))
        else:
            request.resource.ClearField("run")
        if TargetMultiTarget.to_proto(self.multi_target):
            request.resource.multi_target.CopyFrom(
                TargetMultiTarget.to_proto(self.multi_target)
            )
        else:
            request.resource.ClearField("multi_target")
        if Primitive.to_proto(self.deploy_parameters):
            request.resource.deploy_parameters = Primitive.to_proto(
                self.deploy_parameters
            )

        if TargetCustomTarget.to_proto(self.custom_target):
            request.resource.custom_target.CopyFrom(
                TargetCustomTarget.to_proto(self.custom_target)
            )
        else:
            request.resource.ClearField("custom_target")
        if Primitive.to_proto(self.associated_entities):
            request.resource.associated_entities = Primitive.to_proto(
                self.associated_entities
            )

        response = stub.DeleteClouddeployTarget(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = target_pb2_grpc.ClouddeployTargetServiceStub(channel.Channel())
        request = target_pb2.ListClouddeployTargetRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListClouddeployTarget(request).items

    def to_proto(self):
        resource = target_pb2.ClouddeployTarget()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.require_approval):
            resource.require_approval = Primitive.to_proto(self.require_approval)
        if TargetGke.to_proto(self.gke):
            resource.gke.CopyFrom(TargetGke.to_proto(self.gke))
        else:
            resource.ClearField("gke")
        if TargetAnthosCluster.to_proto(self.anthos_cluster):
            resource.anthos_cluster.CopyFrom(
                TargetAnthosCluster.to_proto(self.anthos_cluster)
            )
        else:
            resource.ClearField("anthos_cluster")
        if TargetExecutionConfigsArray.to_proto(self.execution_configs):
            resource.execution_configs.extend(
                TargetExecutionConfigsArray.to_proto(self.execution_configs)
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if TargetRun.to_proto(self.run):
            resource.run.CopyFrom(TargetRun.to_proto(self.run))
        else:
            resource.ClearField("run")
        if TargetMultiTarget.to_proto(self.multi_target):
            resource.multi_target.CopyFrom(
                TargetMultiTarget.to_proto(self.multi_target)
            )
        else:
            resource.ClearField("multi_target")
        if Primitive.to_proto(self.deploy_parameters):
            resource.deploy_parameters = Primitive.to_proto(self.deploy_parameters)
        if TargetCustomTarget.to_proto(self.custom_target):
            resource.custom_target.CopyFrom(
                TargetCustomTarget.to_proto(self.custom_target)
            )
        else:
            resource.ClearField("custom_target")
        if Primitive.to_proto(self.associated_entities):
            resource.associated_entities = Primitive.to_proto(self.associated_entities)
        return resource


class TargetGke(object):

    def __init__(
        self,
        cluster: str = None,
        internal_ip: bool = None,
        proxy_url: str = None,
        dns_endpoint: bool = None,
    ):
        self.cluster = cluster
        self.internal_ip = internal_ip
        self.proxy_url = proxy_url
        self.dns_endpoint = dns_endpoint

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = target_pb2.ClouddeployTargetGke()
        if Primitive.to_proto(resource.cluster):
            res.cluster = Primitive.to_proto(resource.cluster)
        if Primitive.to_proto(resource.internal_ip):
            res.internal_ip = Primitive.to_proto(resource.internal_ip)
        if Primitive.to_proto(resource.proxy_url):
            res.proxy_url = Primitive.to_proto(resource.proxy_url)
        if Primitive.to_proto(resource.dns_endpoint):
            res.dns_endpoint = Primitive.to_proto(resource.dns_endpoint)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TargetGke(
            cluster=Primitive.from_proto(resource.cluster),
            internal_ip=Primitive.from_proto(resource.internal_ip),
            proxy_url=Primitive.from_proto(resource.proxy_url),
            dns_endpoint=Primitive.from_proto(resource.dns_endpoint),
        )


class TargetGkeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TargetGke.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TargetGke.from_proto(i) for i in resources]


class TargetAnthosCluster(object):

    def __init__(self, membership: str = None):
        self.membership = membership

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = target_pb2.ClouddeployTargetAnthosCluster()
        if Primitive.to_proto(resource.membership):
            res.membership = Primitive.to_proto(resource.membership)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TargetAnthosCluster(
            membership=Primitive.from_proto(resource.membership),
        )


class TargetAnthosClusterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TargetAnthosCluster.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TargetAnthosCluster.from_proto(i) for i in resources]


class TargetExecutionConfigs(object):

    def __init__(
        self,
        usages: list = None,
        worker_pool: str = None,
        service_account: str = None,
        artifact_storage: str = None,
        execution_timeout: str = None,
        verbose: bool = None,
    ):
        self.usages = usages
        self.worker_pool = worker_pool
        self.service_account = service_account
        self.artifact_storage = artifact_storage
        self.execution_timeout = execution_timeout
        self.verbose = verbose

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = target_pb2.ClouddeployTargetExecutionConfigs()
        if TargetExecutionConfigsUsagesEnumArray.to_proto(resource.usages):
            res.usages.extend(
                TargetExecutionConfigsUsagesEnumArray.to_proto(resource.usages)
            )
        if Primitive.to_proto(resource.worker_pool):
            res.worker_pool = Primitive.to_proto(resource.worker_pool)
        if Primitive.to_proto(resource.service_account):
            res.service_account = Primitive.to_proto(resource.service_account)
        if Primitive.to_proto(resource.artifact_storage):
            res.artifact_storage = Primitive.to_proto(resource.artifact_storage)
        if Primitive.to_proto(resource.execution_timeout):
            res.execution_timeout = Primitive.to_proto(resource.execution_timeout)
        if Primitive.to_proto(resource.verbose):
            res.verbose = Primitive.to_proto(resource.verbose)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TargetExecutionConfigs(
            usages=TargetExecutionConfigsUsagesEnumArray.from_proto(resource.usages),
            worker_pool=Primitive.from_proto(resource.worker_pool),
            service_account=Primitive.from_proto(resource.service_account),
            artifact_storage=Primitive.from_proto(resource.artifact_storage),
            execution_timeout=Primitive.from_proto(resource.execution_timeout),
            verbose=Primitive.from_proto(resource.verbose),
        )


class TargetExecutionConfigsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TargetExecutionConfigs.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TargetExecutionConfigs.from_proto(i) for i in resources]


class TargetRun(object):

    def __init__(self, location: str = None):
        self.location = location

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = target_pb2.ClouddeployTargetRun()
        if Primitive.to_proto(resource.location):
            res.location = Primitive.to_proto(resource.location)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TargetRun(
            location=Primitive.from_proto(resource.location),
        )


class TargetRunArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TargetRun.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TargetRun.from_proto(i) for i in resources]


class TargetMultiTarget(object):

    def __init__(self, target_ids: list = None):
        self.target_ids = target_ids

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = target_pb2.ClouddeployTargetMultiTarget()
        if Primitive.to_proto(resource.target_ids):
            res.target_ids.extend(Primitive.to_proto(resource.target_ids))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TargetMultiTarget(
            target_ids=Primitive.from_proto(resource.target_ids),
        )


class TargetMultiTargetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TargetMultiTarget.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TargetMultiTarget.from_proto(i) for i in resources]


class TargetCustomTarget(object):

    def __init__(self, custom_target_type: str = None):
        self.custom_target_type = custom_target_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = target_pb2.ClouddeployTargetCustomTarget()
        if Primitive.to_proto(resource.custom_target_type):
            res.custom_target_type = Primitive.to_proto(resource.custom_target_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TargetCustomTarget(
            custom_target_type=Primitive.from_proto(resource.custom_target_type),
        )


class TargetCustomTargetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TargetCustomTarget.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TargetCustomTarget.from_proto(i) for i in resources]


class TargetAssociatedEntities(object):

    def __init__(self, gke_clusters: list = None, anthos_clusters: list = None):
        self.gke_clusters = gke_clusters
        self.anthos_clusters = anthos_clusters

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = target_pb2.ClouddeployTargetAssociatedEntities()
        if TargetAssociatedEntitiesGkeClustersArray.to_proto(resource.gke_clusters):
            res.gke_clusters.extend(
                TargetAssociatedEntitiesGkeClustersArray.to_proto(resource.gke_clusters)
            )
        if TargetAssociatedEntitiesAnthosClustersArray.to_proto(
            resource.anthos_clusters
        ):
            res.anthos_clusters.extend(
                TargetAssociatedEntitiesAnthosClustersArray.to_proto(
                    resource.anthos_clusters
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TargetAssociatedEntities(
            gke_clusters=TargetAssociatedEntitiesGkeClustersArray.from_proto(
                resource.gke_clusters
            ),
            anthos_clusters=TargetAssociatedEntitiesAnthosClustersArray.from_proto(
                resource.anthos_clusters
            ),
        )


class TargetAssociatedEntitiesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TargetAssociatedEntities.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TargetAssociatedEntities.from_proto(i) for i in resources]


class TargetAssociatedEntitiesGkeClusters(object):

    def __init__(
        self, cluster: str = None, internal_ip: bool = None, proxy_url: str = None
    ):
        self.cluster = cluster
        self.internal_ip = internal_ip
        self.proxy_url = proxy_url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = target_pb2.ClouddeployTargetAssociatedEntitiesGkeClusters()
        if Primitive.to_proto(resource.cluster):
            res.cluster = Primitive.to_proto(resource.cluster)
        if Primitive.to_proto(resource.internal_ip):
            res.internal_ip = Primitive.to_proto(resource.internal_ip)
        if Primitive.to_proto(resource.proxy_url):
            res.proxy_url = Primitive.to_proto(resource.proxy_url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TargetAssociatedEntitiesGkeClusters(
            cluster=Primitive.from_proto(resource.cluster),
            internal_ip=Primitive.from_proto(resource.internal_ip),
            proxy_url=Primitive.from_proto(resource.proxy_url),
        )


class TargetAssociatedEntitiesGkeClustersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TargetAssociatedEntitiesGkeClusters.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TargetAssociatedEntitiesGkeClusters.from_proto(i) for i in resources]


class TargetAssociatedEntitiesAnthosClusters(object):

    def __init__(self, membership: str = None):
        self.membership = membership

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = target_pb2.ClouddeployTargetAssociatedEntitiesAnthosClusters()
        if Primitive.to_proto(resource.membership):
            res.membership = Primitive.to_proto(resource.membership)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return TargetAssociatedEntitiesAnthosClusters(
            membership=Primitive.from_proto(resource.membership),
        )


class TargetAssociatedEntitiesAnthosClustersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [TargetAssociatedEntitiesAnthosClusters.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [TargetAssociatedEntitiesAnthosClusters.from_proto(i) for i in resources]


class TargetExecutionConfigsUsagesEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return target_pb2.ClouddeployTargetExecutionConfigsUsagesEnum.Value(
            "ClouddeployTargetExecutionConfigsUsagesEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return target_pb2.ClouddeployTargetExecutionConfigsUsagesEnum.Name(resource)[
            len("ClouddeployTargetExecutionConfigsUsagesEnum") :
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
