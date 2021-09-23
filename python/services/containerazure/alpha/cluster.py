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
from google3.cloud.graphite.mmv2.services.google.container_azure import cluster_pb2
from google3.cloud.graphite.mmv2.services.google.container_azure import cluster_pb2_grpc

from typing import List


class Cluster(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        azure_region: str = None,
        resource_group_id: str = None,
        client: str = None,
        networking: dict = None,
        control_plane: dict = None,
        authorization: dict = None,
        state: str = None,
        endpoint: str = None,
        uid: str = None,
        reconciling: bool = None,
        create_time: str = None,
        update_time: str = None,
        etag: str = None,
        annotations: dict = None,
        workload_identity_config: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.azure_region = azure_region
        self.resource_group_id = resource_group_id
        self.client = client
        self.networking = networking
        self.control_plane = control_plane
        self.authorization = authorization
        self.annotations = annotations
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = cluster_pb2_grpc.ContainerazureAlphaClusterServiceStub(channel.Channel())
        request = cluster_pb2.ApplyContainerazureAlphaClusterRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.azure_region):
            request.resource.azure_region = Primitive.to_proto(self.azure_region)

        if Primitive.to_proto(self.resource_group_id):
            request.resource.resource_group_id = Primitive.to_proto(
                self.resource_group_id
            )

        if Primitive.to_proto(self.client):
            request.resource.client = Primitive.to_proto(self.client)

        if ClusterNetworking.to_proto(self.networking):
            request.resource.networking.CopyFrom(
                ClusterNetworking.to_proto(self.networking)
            )
        else:
            request.resource.ClearField("networking")
        if ClusterControlPlane.to_proto(self.control_plane):
            request.resource.control_plane.CopyFrom(
                ClusterControlPlane.to_proto(self.control_plane)
            )
        else:
            request.resource.ClearField("control_plane")
        if ClusterAuthorization.to_proto(self.authorization):
            request.resource.authorization.CopyFrom(
                ClusterAuthorization.to_proto(self.authorization)
            )
        else:
            request.resource.ClearField("authorization")
        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyContainerazureAlphaCluster(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.azure_region = Primitive.from_proto(response.azure_region)
        self.resource_group_id = Primitive.from_proto(response.resource_group_id)
        self.client = Primitive.from_proto(response.client)
        self.networking = ClusterNetworking.from_proto(response.networking)
        self.control_plane = ClusterControlPlane.from_proto(response.control_plane)
        self.authorization = ClusterAuthorization.from_proto(response.authorization)
        self.state = ClusterStateEnum.from_proto(response.state)
        self.endpoint = Primitive.from_proto(response.endpoint)
        self.uid = Primitive.from_proto(response.uid)
        self.reconciling = Primitive.from_proto(response.reconciling)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.etag = Primitive.from_proto(response.etag)
        self.annotations = Primitive.from_proto(response.annotations)
        self.workload_identity_config = ClusterWorkloadIdentityConfig.from_proto(
            response.workload_identity_config
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = cluster_pb2_grpc.ContainerazureAlphaClusterServiceStub(channel.Channel())
        request = cluster_pb2.DeleteContainerazureAlphaClusterRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.azure_region):
            request.resource.azure_region = Primitive.to_proto(self.azure_region)

        if Primitive.to_proto(self.resource_group_id):
            request.resource.resource_group_id = Primitive.to_proto(
                self.resource_group_id
            )

        if Primitive.to_proto(self.client):
            request.resource.client = Primitive.to_proto(self.client)

        if ClusterNetworking.to_proto(self.networking):
            request.resource.networking.CopyFrom(
                ClusterNetworking.to_proto(self.networking)
            )
        else:
            request.resource.ClearField("networking")
        if ClusterControlPlane.to_proto(self.control_plane):
            request.resource.control_plane.CopyFrom(
                ClusterControlPlane.to_proto(self.control_plane)
            )
        else:
            request.resource.ClearField("control_plane")
        if ClusterAuthorization.to_proto(self.authorization):
            request.resource.authorization.CopyFrom(
                ClusterAuthorization.to_proto(self.authorization)
            )
        else:
            request.resource.ClearField("authorization")
        if Primitive.to_proto(self.annotations):
            request.resource.annotations = Primitive.to_proto(self.annotations)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteContainerazureAlphaCluster(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = cluster_pb2_grpc.ContainerazureAlphaClusterServiceStub(channel.Channel())
        request = cluster_pb2.ListContainerazureAlphaClusterRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListContainerazureAlphaCluster(request).items

    def to_proto(self):
        resource = cluster_pb2.ContainerazureAlphaCluster()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.azure_region):
            resource.azure_region = Primitive.to_proto(self.azure_region)
        if Primitive.to_proto(self.resource_group_id):
            resource.resource_group_id = Primitive.to_proto(self.resource_group_id)
        if Primitive.to_proto(self.client):
            resource.client = Primitive.to_proto(self.client)
        if ClusterNetworking.to_proto(self.networking):
            resource.networking.CopyFrom(ClusterNetworking.to_proto(self.networking))
        else:
            resource.ClearField("networking")
        if ClusterControlPlane.to_proto(self.control_plane):
            resource.control_plane.CopyFrom(
                ClusterControlPlane.to_proto(self.control_plane)
            )
        else:
            resource.ClearField("control_plane")
        if ClusterAuthorization.to_proto(self.authorization):
            resource.authorization.CopyFrom(
                ClusterAuthorization.to_proto(self.authorization)
            )
        else:
            resource.ClearField("authorization")
        if Primitive.to_proto(self.annotations):
            resource.annotations = Primitive.to_proto(self.annotations)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class ClusterNetworking(object):
    def __init__(
        self,
        virtual_network_id: str = None,
        pod_address_cidr_blocks: list = None,
        service_address_cidr_blocks: list = None,
    ):
        self.virtual_network_id = virtual_network_id
        self.pod_address_cidr_blocks = pod_address_cidr_blocks
        self.service_address_cidr_blocks = service_address_cidr_blocks

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerazureAlphaClusterNetworking()
        if Primitive.to_proto(resource.virtual_network_id):
            res.virtual_network_id = Primitive.to_proto(resource.virtual_network_id)
        if Primitive.to_proto(resource.pod_address_cidr_blocks):
            res.pod_address_cidr_blocks.extend(
                Primitive.to_proto(resource.pod_address_cidr_blocks)
            )
        if Primitive.to_proto(resource.service_address_cidr_blocks):
            res.service_address_cidr_blocks.extend(
                Primitive.to_proto(resource.service_address_cidr_blocks)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterNetworking(
            virtual_network_id=Primitive.from_proto(resource.virtual_network_id),
            pod_address_cidr_blocks=Primitive.from_proto(
                resource.pod_address_cidr_blocks
            ),
            service_address_cidr_blocks=Primitive.from_proto(
                resource.service_address_cidr_blocks
            ),
        )


class ClusterNetworkingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterNetworking.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterNetworking.from_proto(i) for i in resources]


class ClusterControlPlane(object):
    def __init__(
        self,
        version: str = None,
        subnet_id: str = None,
        vm_size: str = None,
        ssh_config: dict = None,
        root_volume: dict = None,
        main_volume: dict = None,
        database_encryption: dict = None,
        tags: dict = None,
    ):
        self.version = version
        self.subnet_id = subnet_id
        self.vm_size = vm_size
        self.ssh_config = ssh_config
        self.root_volume = root_volume
        self.main_volume = main_volume
        self.database_encryption = database_encryption
        self.tags = tags

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerazureAlphaClusterControlPlane()
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        if Primitive.to_proto(resource.subnet_id):
            res.subnet_id = Primitive.to_proto(resource.subnet_id)
        if Primitive.to_proto(resource.vm_size):
            res.vm_size = Primitive.to_proto(resource.vm_size)
        if ClusterControlPlaneSshConfig.to_proto(resource.ssh_config):
            res.ssh_config.CopyFrom(
                ClusterControlPlaneSshConfig.to_proto(resource.ssh_config)
            )
        else:
            res.ClearField("ssh_config")
        if ClusterControlPlaneRootVolume.to_proto(resource.root_volume):
            res.root_volume.CopyFrom(
                ClusterControlPlaneRootVolume.to_proto(resource.root_volume)
            )
        else:
            res.ClearField("root_volume")
        if ClusterControlPlaneMainVolume.to_proto(resource.main_volume):
            res.main_volume.CopyFrom(
                ClusterControlPlaneMainVolume.to_proto(resource.main_volume)
            )
        else:
            res.ClearField("main_volume")
        if ClusterControlPlaneDatabaseEncryption.to_proto(resource.database_encryption):
            res.database_encryption.CopyFrom(
                ClusterControlPlaneDatabaseEncryption.to_proto(
                    resource.database_encryption
                )
            )
        else:
            res.ClearField("database_encryption")
        if Primitive.to_proto(resource.tags):
            res.tags = Primitive.to_proto(resource.tags)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterControlPlane(
            version=Primitive.from_proto(resource.version),
            subnet_id=Primitive.from_proto(resource.subnet_id),
            vm_size=Primitive.from_proto(resource.vm_size),
            ssh_config=ClusterControlPlaneSshConfig.from_proto(resource.ssh_config),
            root_volume=ClusterControlPlaneRootVolume.from_proto(resource.root_volume),
            main_volume=ClusterControlPlaneMainVolume.from_proto(resource.main_volume),
            database_encryption=ClusterControlPlaneDatabaseEncryption.from_proto(
                resource.database_encryption
            ),
            tags=Primitive.from_proto(resource.tags),
        )


class ClusterControlPlaneArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterControlPlane.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterControlPlane.from_proto(i) for i in resources]


class ClusterControlPlaneSshConfig(object):
    def __init__(self, authorized_key: str = None):
        self.authorized_key = authorized_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerazureAlphaClusterControlPlaneSshConfig()
        if Primitive.to_proto(resource.authorized_key):
            res.authorized_key = Primitive.to_proto(resource.authorized_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterControlPlaneSshConfig(
            authorized_key=Primitive.from_proto(resource.authorized_key),
        )


class ClusterControlPlaneSshConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterControlPlaneSshConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterControlPlaneSshConfig.from_proto(i) for i in resources]


class ClusterControlPlaneRootVolume(object):
    def __init__(self, size_gib: int = None):
        self.size_gib = size_gib

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerazureAlphaClusterControlPlaneRootVolume()
        if Primitive.to_proto(resource.size_gib):
            res.size_gib = Primitive.to_proto(resource.size_gib)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterControlPlaneRootVolume(
            size_gib=Primitive.from_proto(resource.size_gib),
        )


class ClusterControlPlaneRootVolumeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterControlPlaneRootVolume.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterControlPlaneRootVolume.from_proto(i) for i in resources]


class ClusterControlPlaneMainVolume(object):
    def __init__(self, size_gib: int = None):
        self.size_gib = size_gib

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerazureAlphaClusterControlPlaneMainVolume()
        if Primitive.to_proto(resource.size_gib):
            res.size_gib = Primitive.to_proto(resource.size_gib)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterControlPlaneMainVolume(
            size_gib=Primitive.from_proto(resource.size_gib),
        )


class ClusterControlPlaneMainVolumeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterControlPlaneMainVolume.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterControlPlaneMainVolume.from_proto(i) for i in resources]


class ClusterControlPlaneDatabaseEncryption(object):
    def __init__(self, resource_group_id: str = None, kms_key_identifier: str = None):
        self.resource_group_id = resource_group_id
        self.kms_key_identifier = kms_key_identifier

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerazureAlphaClusterControlPlaneDatabaseEncryption()
        if Primitive.to_proto(resource.resource_group_id):
            res.resource_group_id = Primitive.to_proto(resource.resource_group_id)
        if Primitive.to_proto(resource.kms_key_identifier):
            res.kms_key_identifier = Primitive.to_proto(resource.kms_key_identifier)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterControlPlaneDatabaseEncryption(
            resource_group_id=Primitive.from_proto(resource.resource_group_id),
            kms_key_identifier=Primitive.from_proto(resource.kms_key_identifier),
        )


class ClusterControlPlaneDatabaseEncryptionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterControlPlaneDatabaseEncryption.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterControlPlaneDatabaseEncryption.from_proto(i) for i in resources]


class ClusterAuthorization(object):
    def __init__(self, admin_users: list = None):
        self.admin_users = admin_users

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerazureAlphaClusterAuthorization()
        if ClusterAuthorizationAdminUsersArray.to_proto(resource.admin_users):
            res.admin_users.extend(
                ClusterAuthorizationAdminUsersArray.to_proto(resource.admin_users)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAuthorization(
            admin_users=ClusterAuthorizationAdminUsersArray.from_proto(
                resource.admin_users
            ),
        )


class ClusterAuthorizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAuthorization.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAuthorization.from_proto(i) for i in resources]


class ClusterAuthorizationAdminUsers(object):
    def __init__(self, username: str = None):
        self.username = username

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerazureAlphaClusterAuthorizationAdminUsers()
        if Primitive.to_proto(resource.username):
            res.username = Primitive.to_proto(resource.username)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterAuthorizationAdminUsers(
            username=Primitive.from_proto(resource.username),
        )


class ClusterAuthorizationAdminUsersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterAuthorizationAdminUsers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterAuthorizationAdminUsers.from_proto(i) for i in resources]


class ClusterWorkloadIdentityConfig(object):
    def __init__(
        self,
        issuer_uri: str = None,
        workload_pool: str = None,
        identity_provider: str = None,
    ):
        self.issuer_uri = issuer_uri
        self.workload_pool = workload_pool
        self.identity_provider = identity_provider

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = cluster_pb2.ContainerazureAlphaClusterWorkloadIdentityConfig()
        if Primitive.to_proto(resource.issuer_uri):
            res.issuer_uri = Primitive.to_proto(resource.issuer_uri)
        if Primitive.to_proto(resource.workload_pool):
            res.workload_pool = Primitive.to_proto(resource.workload_pool)
        if Primitive.to_proto(resource.identity_provider):
            res.identity_provider = Primitive.to_proto(resource.identity_provider)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ClusterWorkloadIdentityConfig(
            issuer_uri=Primitive.from_proto(resource.issuer_uri),
            workload_pool=Primitive.from_proto(resource.workload_pool),
            identity_provider=Primitive.from_proto(resource.identity_provider),
        )


class ClusterWorkloadIdentityConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ClusterWorkloadIdentityConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ClusterWorkloadIdentityConfig.from_proto(i) for i in resources]


class ClusterStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerazureAlphaClusterStateEnum.Value(
            "ContainerazureAlphaClusterStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return cluster_pb2.ContainerazureAlphaClusterStateEnum.Name(resource)[
            len("ContainerazureAlphaClusterStateEnum") :
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