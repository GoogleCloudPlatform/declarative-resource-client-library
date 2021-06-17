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
from google3.cloud.graphite.mmv2.services.google.os_config import (
    os_policy_assignment_pb2,
)
from google3.cloud.graphite.mmv2.services.google.os_config import (
    os_policy_assignment_pb2_grpc,
)

from typing import List


class OsPolicyAssignment(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        os_policies: list = None,
        instance_filter: dict = None,
        rollout: dict = None,
        revision_id: str = None,
        revision_create_time: str = None,
        rollout_state: str = None,
        baseline: bool = None,
        deleted: bool = None,
        reconciling: bool = None,
        uid: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.os_policies = os_policies
        self.instance_filter = instance_filter
        self.rollout = rollout
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = os_policy_assignment_pb2_grpc.OsconfigAlphaOsPolicyAssignmentServiceStub(
            channel.Channel()
        )
        request = os_policy_assignment_pb2.ApplyOsconfigAlphaOsPolicyAssignmentRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if OsPolicyAssignmentOsPoliciesArray.to_proto(self.os_policies):
            request.resource.os_policies.extend(
                OsPolicyAssignmentOsPoliciesArray.to_proto(self.os_policies)
            )
        if OsPolicyAssignmentInstanceFilter.to_proto(self.instance_filter):
            request.resource.instance_filter.CopyFrom(
                OsPolicyAssignmentInstanceFilter.to_proto(self.instance_filter)
            )
        else:
            request.resource.ClearField("instance_filter")
        if OsPolicyAssignmentRollout.to_proto(self.rollout):
            request.resource.rollout.CopyFrom(
                OsPolicyAssignmentRollout.to_proto(self.rollout)
            )
        else:
            request.resource.ClearField("rollout")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyOsconfigAlphaOsPolicyAssignment(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.os_policies = OsPolicyAssignmentOsPoliciesArray.from_proto(
            response.os_policies
        )
        self.instance_filter = OsPolicyAssignmentInstanceFilter.from_proto(
            response.instance_filter
        )
        self.rollout = OsPolicyAssignmentRollout.from_proto(response.rollout)
        self.revision_id = Primitive.from_proto(response.revision_id)
        self.revision_create_time = Primitive.from_proto(response.revision_create_time)
        self.rollout_state = OsPolicyAssignmentRolloutStateEnum.from_proto(
            response.rollout_state
        )
        self.baseline = Primitive.from_proto(response.baseline)
        self.deleted = Primitive.from_proto(response.deleted)
        self.reconciling = Primitive.from_proto(response.reconciling)
        self.uid = Primitive.from_proto(response.uid)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = os_policy_assignment_pb2_grpc.OsconfigAlphaOsPolicyAssignmentServiceStub(
            channel.Channel()
        )
        request = (
            os_policy_assignment_pb2.DeleteOsconfigAlphaOsPolicyAssignmentRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if OsPolicyAssignmentOsPoliciesArray.to_proto(self.os_policies):
            request.resource.os_policies.extend(
                OsPolicyAssignmentOsPoliciesArray.to_proto(self.os_policies)
            )
        if OsPolicyAssignmentInstanceFilter.to_proto(self.instance_filter):
            request.resource.instance_filter.CopyFrom(
                OsPolicyAssignmentInstanceFilter.to_proto(self.instance_filter)
            )
        else:
            request.resource.ClearField("instance_filter")
        if OsPolicyAssignmentRollout.to_proto(self.rollout):
            request.resource.rollout.CopyFrom(
                OsPolicyAssignmentRollout.to_proto(self.rollout)
            )
        else:
            request.resource.ClearField("rollout")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteOsconfigAlphaOsPolicyAssignment(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = os_policy_assignment_pb2_grpc.OsconfigAlphaOsPolicyAssignmentServiceStub(
            channel.Channel()
        )
        request = os_policy_assignment_pb2.ListOsconfigAlphaOsPolicyAssignmentRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListOsconfigAlphaOsPolicyAssignment(request).items

    def to_proto(self):
        resource = os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignment()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if OsPolicyAssignmentOsPoliciesArray.to_proto(self.os_policies):
            resource.os_policies.extend(
                OsPolicyAssignmentOsPoliciesArray.to_proto(self.os_policies)
            )
        if OsPolicyAssignmentInstanceFilter.to_proto(self.instance_filter):
            resource.instance_filter.CopyFrom(
                OsPolicyAssignmentInstanceFilter.to_proto(self.instance_filter)
            )
        else:
            resource.ClearField("instance_filter")
        if OsPolicyAssignmentRollout.to_proto(self.rollout):
            resource.rollout.CopyFrom(OsPolicyAssignmentRollout.to_proto(self.rollout))
        else:
            resource.ClearField("rollout")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class OsPolicyAssignmentOsPolicies(object):
    def __init__(
        self,
        id: str = None,
        description: str = None,
        mode: str = None,
        resource_groups: list = None,
        allow_no_resource_group_match: bool = None,
    ):
        self.id = id
        self.description = description
        self.mode = mode
        self.resource_groups = resource_groups
        self.allow_no_resource_group_match = allow_no_resource_group_match

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPolicies()
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if OsPolicyAssignmentOsPoliciesModeEnum.to_proto(resource.mode):
            res.mode = OsPolicyAssignmentOsPoliciesModeEnum.to_proto(resource.mode)
        if OsPolicyAssignmentOsPoliciesResourceGroupsArray.to_proto(
            resource.resource_groups
        ):
            res.resource_groups.extend(
                OsPolicyAssignmentOsPoliciesResourceGroupsArray.to_proto(
                    resource.resource_groups
                )
            )
        if Primitive.to_proto(resource.allow_no_resource_group_match):
            res.allow_no_resource_group_match = Primitive.to_proto(
                resource.allow_no_resource_group_match
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPolicies(
            id=Primitive.from_proto(resource.id),
            description=Primitive.from_proto(resource.description),
            mode=OsPolicyAssignmentOsPoliciesModeEnum.from_proto(resource.mode),
            resource_groups=OsPolicyAssignmentOsPoliciesResourceGroupsArray.from_proto(
                resource.resource_groups
            ),
            allow_no_resource_group_match=Primitive.from_proto(
                resource.allow_no_resource_group_match
            ),
        )


class OsPolicyAssignmentOsPoliciesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OsPolicyAssignmentOsPolicies.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OsPolicyAssignmentOsPolicies.from_proto(i) for i in resources]


class OsPolicyAssignmentOsPoliciesResourceGroups(object):
    def __init__(self, os_filter: dict = None, resources: list = None):
        self.os_filter = os_filter
        self.resources = resources

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroups()
        )
        if OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter.to_proto(
            resource.os_filter
        ):
            res.os_filter.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter.to_proto(
                    resource.os_filter
                )
            )
        else:
            res.ClearField("os_filter")
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesArray.to_proto(
            resource.resources
        ):
            res.resources.extend(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesArray.to_proto(
                    resource.resources
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroups(
            os_filter=OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter.from_proto(
                resource.os_filter
            ),
            resources=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesArray.from_proto(
                resource.resources
            ),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroups.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroups.from_proto(i) for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(object):
    def __init__(self, os_short_name: str = None, os_version: str = None):
        self.os_short_name = os_short_name
        self.os_version = os_version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsOsFilter()
        )
        if Primitive.to_proto(resource.os_short_name):
            res.os_short_name = Primitive.to_proto(resource.os_short_name)
        if Primitive.to_proto(resource.os_version):
            res.os_version = Primitive.to_proto(resource.os_version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter(
            os_short_name=Primitive.from_proto(resource.os_short_name),
            os_version=Primitive.from_proto(resource.os_version),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsOsFilterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsOsFilter.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResources(object):
    def __init__(
        self,
        id: str = None,
        pkg: dict = None,
        repository: dict = None,
        exec: dict = None,
        file: dict = None,
    ):
        self.id = id
        self.pkg = pkg
        self.repository = repository
        self.exec = exec
        self.file = file

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResources()
        )
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg.to_proto(
            resource.pkg
        ):
            res.pkg.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg.to_proto(
                    resource.pkg
                )
            )
        else:
            res.ClearField("pkg")
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository.to_proto(
            resource.repository
        ):
            res.repository.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository.to_proto(
                    resource.repository
                )
            )
        else:
            res.ClearField("repository")
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec.to_proto(
            resource.exec
        ):
            res.exec.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec.to_proto(
                    resource.exec
                )
            )
        else:
            res.ClearField("exec")
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile.to_proto(
            resource.file
        ):
            res.file.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile.to_proto(
                    resource.file
                )
            )
        else:
            res.ClearField("file")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResources(
            id=Primitive.from_proto(resource.id),
            pkg=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg.from_proto(
                resource.pkg
            ),
            repository=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository.from_proto(
                resource.repository
            ),
            exec=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec.from_proto(
                resource.exec
            ),
            file=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile.from_proto(
                resource.file
            ),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResources.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResources.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(object):
    def __init__(
        self,
        desired_state: str = None,
        apt: dict = None,
        deb: dict = None,
        yum: dict = None,
        zypper: dict = None,
        rpm: dict = None,
        googet: dict = None,
        msi: dict = None,
    ):
        self.desired_state = desired_state
        self.apt = apt
        self.deb = deb
        self.yum = yum
        self.zypper = zypper
        self.rpm = rpm
        self.googet = googet
        self.msi = msi

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg()
        )
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum.to_proto(
            resource.desired_state
        ):
            res.desired_state = OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum.to_proto(
                resource.desired_state
            )
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt.to_proto(
            resource.apt
        ):
            res.apt.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt.to_proto(
                    resource.apt
                )
            )
        else:
            res.ClearField("apt")
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb.to_proto(
            resource.deb
        ):
            res.deb.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb.to_proto(
                    resource.deb
                )
            )
        else:
            res.ClearField("deb")
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum.to_proto(
            resource.yum
        ):
            res.yum.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum.to_proto(
                    resource.yum
                )
            )
        else:
            res.ClearField("yum")
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper.to_proto(
            resource.zypper
        ):
            res.zypper.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper.to_proto(
                    resource.zypper
                )
            )
        else:
            res.ClearField("zypper")
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm.to_proto(
            resource.rpm
        ):
            res.rpm.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm.to_proto(
                    resource.rpm
                )
            )
        else:
            res.ClearField("rpm")
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget.to_proto(
            resource.googet
        ):
            res.googet.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget.to_proto(
                    resource.googet
                )
            )
        else:
            res.ClearField("googet")
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi.to_proto(
            resource.msi
        ):
            res.msi.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi.to_proto(
                    resource.msi
                )
            )
        else:
            res.ClearField("msi")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg(
            desired_state=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum.from_proto(
                resource.desired_state
            ),
            apt=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt.from_proto(
                resource.apt
            ),
            deb=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb.from_proto(
                resource.deb
            ),
            yum=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum.from_proto(
                resource.yum
            ),
            zypper=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper.from_proto(
                resource.zypper
            ),
            rpm=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm.from_proto(
                resource.rpm
            ),
            googet=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget.from_proto(
                resource.googet
            ),
            msi=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi.from_proto(
                resource.msi
            ),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkg.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt(
            name=Primitive.from_proto(resource.name),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgAptArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgApt.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(object):
    def __init__(self, source: dict = None, pull_deps: bool = None):
        self.source = source
        self.pull_deps = pull_deps

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb()
        )
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource.to_proto(
            resource.source
        ):
            res.source.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource.to_proto(
                    resource.source
                )
            )
        else:
            res.ClearField("source")
        if Primitive.to_proto(resource.pull_deps):
            res.pull_deps = Primitive.to_proto(resource.pull_deps)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb(
            source=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource.from_proto(
                resource.source
            ),
            pull_deps=Primitive.from_proto(resource.pull_deps),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(object):
    def __init__(
        self,
        remote: dict = None,
        gcs: dict = None,
        local_path: str = None,
        allow_insecure: bool = None,
    ):
        self.remote = remote
        self.gcs = gcs
        self.local_path = local_path
        self.allow_insecure = allow_insecure

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource()
        )
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote.to_proto(
            resource.remote
        ):
            res.remote.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote.to_proto(
                    resource.remote
                )
            )
        else:
            res.ClearField("remote")
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs.to_proto(
            resource.gcs
        ):
            res.gcs.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs.to_proto(
                    resource.gcs
                )
            )
        else:
            res.ClearField("gcs")
        if Primitive.to_proto(resource.local_path):
            res.local_path = Primitive.to_proto(resource.local_path)
        if Primitive.to_proto(resource.allow_insecure):
            res.allow_insecure = Primitive.to_proto(resource.allow_insecure)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource(
            remote=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote.from_proto(
                resource.remote
            ),
            gcs=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs.from_proto(
                resource.gcs
            ),
            local_path=Primitive.from_proto(resource.local_path),
            allow_insecure=Primitive.from_proto(resource.allow_insecure),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource.from_proto(
                i
            )
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(object):
    def __init__(self, uri: str = None, sha256_checksum: str = None):
        self.uri = uri
        self.sha256_checksum = sha256_checksum

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote()
        )
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        if Primitive.to_proto(resource.sha256_checksum):
            res.sha256_checksum = Primitive.to_proto(resource.sha256_checksum)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote(
            uri=Primitive.from_proto(resource.uri),
            sha256_checksum=Primitive.from_proto(resource.sha256_checksum),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote.from_proto(
                i
            )
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(object):
    def __init__(self, bucket: str = None, object: str = None, generation: int = None):
        self.bucket = bucket
        self.object = object
        self.generation = generation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs()
        )
        if Primitive.to_proto(resource.bucket):
            res.bucket = Primitive.to_proto(resource.bucket)
        if Primitive.to_proto(resource.object):
            res.object = Primitive.to_proto(resource.object)
        if Primitive.to_proto(resource.generation):
            res.generation = Primitive.to_proto(resource.generation)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation=Primitive.from_proto(resource.generation),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceGcs.from_proto(
                i
            )
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum(
            name=Primitive.from_proto(resource.name),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYumArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgYum.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper(
            name=Primitive.from_proto(resource.name),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypperArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgZypper.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(object):
    def __init__(self, source: dict = None, pull_deps: bool = None):
        self.source = source
        self.pull_deps = pull_deps

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm()
        )
        if OsPolicyAssignmentFile.to_proto(resource.source):
            res.source.CopyFrom(OsPolicyAssignmentFile.to_proto(resource.source))
        else:
            res.ClearField("source")
        if Primitive.to_proto(resource.pull_deps):
            res.pull_deps = Primitive.to_proto(resource.pull_deps)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm(
            source=OsPolicyAssignmentFile.from_proto(resource.source),
            pull_deps=Primitive.from_proto(resource.pull_deps),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpm.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentFile(object):
    def __init__(
        self,
        remote: dict = None,
        gcs: dict = None,
        local_path: str = None,
        allow_insecure: bool = None,
    ):
        self.remote = remote
        self.gcs = gcs
        self.local_path = local_path
        self.allow_insecure = allow_insecure

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentFile()
        if OsPolicyAssignmentFileRemote.to_proto(resource.remote):
            res.remote.CopyFrom(OsPolicyAssignmentFileRemote.to_proto(resource.remote))
        else:
            res.ClearField("remote")
        if OsPolicyAssignmentFileGcs.to_proto(resource.gcs):
            res.gcs.CopyFrom(OsPolicyAssignmentFileGcs.to_proto(resource.gcs))
        else:
            res.ClearField("gcs")
        if Primitive.to_proto(resource.local_path):
            res.local_path = Primitive.to_proto(resource.local_path)
        if Primitive.to_proto(resource.allow_insecure):
            res.allow_insecure = Primitive.to_proto(resource.allow_insecure)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentFile(
            remote=OsPolicyAssignmentFileRemote.from_proto(resource.remote),
            gcs=OsPolicyAssignmentFileGcs.from_proto(resource.gcs),
            local_path=Primitive.from_proto(resource.local_path),
            allow_insecure=Primitive.from_proto(resource.allow_insecure),
        )


class OsPolicyAssignmentFileArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OsPolicyAssignmentFile.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OsPolicyAssignmentFile.from_proto(i) for i in resources]


class OsPolicyAssignmentFileRemote(object):
    def __init__(self, uri: str = None, sha256_checksum: str = None):
        self.uri = uri
        self.sha256_checksum = sha256_checksum

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentFileRemote()
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        if Primitive.to_proto(resource.sha256_checksum):
            res.sha256_checksum = Primitive.to_proto(resource.sha256_checksum)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentFileRemote(
            uri=Primitive.from_proto(resource.uri),
            sha256_checksum=Primitive.from_proto(resource.sha256_checksum),
        )


class OsPolicyAssignmentFileRemoteArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OsPolicyAssignmentFileRemote.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OsPolicyAssignmentFileRemote.from_proto(i) for i in resources]


class OsPolicyAssignmentFileGcs(object):
    def __init__(self, bucket: str = None, object: str = None, generation: int = None):
        self.bucket = bucket
        self.object = object
        self.generation = generation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentFileGcs()
        if Primitive.to_proto(resource.bucket):
            res.bucket = Primitive.to_proto(resource.bucket)
        if Primitive.to_proto(resource.object):
            res.object = Primitive.to_proto(resource.object)
        if Primitive.to_proto(resource.generation):
            res.generation = Primitive.to_proto(resource.generation)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentFileGcs(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation=Primitive.from_proto(resource.generation),
        )


class OsPolicyAssignmentFileGcsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OsPolicyAssignmentFileGcs.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OsPolicyAssignmentFileGcs.from_proto(i) for i in resources]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget(
            name=Primitive.from_proto(resource.name),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGoogetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgGooget.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(object):
    def __init__(self, source: dict = None, properties: list = None):
        self.source = source
        self.properties = properties

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi()
        )
        if OsPolicyAssignmentFile.to_proto(resource.source):
            res.source.CopyFrom(OsPolicyAssignmentFile.to_proto(resource.source))
        else:
            res.ClearField("source")
        if Primitive.to_proto(resource.properties):
            res.properties.extend(Primitive.to_proto(resource.properties))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi(
            source=OsPolicyAssignmentFile.from_proto(resource.source),
            properties=Primitive.from_proto(resource.properties),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsiArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgMsi.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(object):
    def __init__(
        self, apt: dict = None, yum: dict = None, zypper: dict = None, goo: dict = None
    ):
        self.apt = apt
        self.yum = yum
        self.zypper = zypper
        self.goo = goo

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository()
        )
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt.to_proto(
            resource.apt
        ):
            res.apt.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt.to_proto(
                    resource.apt
                )
            )
        else:
            res.ClearField("apt")
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum.to_proto(
            resource.yum
        ):
            res.yum.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum.to_proto(
                    resource.yum
                )
            )
        else:
            res.ClearField("yum")
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper.to_proto(
            resource.zypper
        ):
            res.zypper.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper.to_proto(
                    resource.zypper
                )
            )
        else:
            res.ClearField("zypper")
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo.to_proto(
            resource.goo
        ):
            res.goo.CopyFrom(
                OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo.to_proto(
                    resource.goo
                )
            )
        else:
            res.ClearField("goo")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository(
            apt=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt.from_proto(
                resource.apt
            ),
            yum=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum.from_proto(
                resource.yum
            ),
            zypper=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper.from_proto(
                resource.zypper
            ),
            goo=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo.from_proto(
                resource.goo
            ),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(object):
    def __init__(
        self,
        archive_type: str = None,
        uri: str = None,
        distribution: str = None,
        components: list = None,
        gpg_key: str = None,
    ):
        self.archive_type = archive_type
        self.uri = uri
        self.distribution = distribution
        self.components = components
        self.gpg_key = gpg_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt()
        )
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.to_proto(
            resource.archive_type
        ):
            res.archive_type = OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.to_proto(
                resource.archive_type
            )
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        if Primitive.to_proto(resource.distribution):
            res.distribution = Primitive.to_proto(resource.distribution)
        if Primitive.to_proto(resource.components):
            res.components.extend(Primitive.to_proto(resource.components))
        if Primitive.to_proto(resource.gpg_key):
            res.gpg_key = Primitive.to_proto(resource.gpg_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt(
            archive_type=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.from_proto(
                resource.archive_type
            ),
            uri=Primitive.from_proto(resource.uri),
            distribution=Primitive.from_proto(resource.distribution),
            components=Primitive.from_proto(resource.components),
            gpg_key=Primitive.from_proto(resource.gpg_key),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt.from_proto(
                i
            )
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(object):
    def __init__(
        self,
        id: str = None,
        display_name: str = None,
        base_url: str = None,
        gpg_keys: list = None,
    ):
        self.id = id
        self.display_name = display_name
        self.base_url = base_url
        self.gpg_keys = gpg_keys

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum()
        )
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.display_name):
            res.display_name = Primitive.to_proto(resource.display_name)
        if Primitive.to_proto(resource.base_url):
            res.base_url = Primitive.to_proto(resource.base_url)
        if Primitive.to_proto(resource.gpg_keys):
            res.gpg_keys.extend(Primitive.to_proto(resource.gpg_keys))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum(
            id=Primitive.from_proto(resource.id),
            display_name=Primitive.from_proto(resource.display_name),
            base_url=Primitive.from_proto(resource.base_url),
            gpg_keys=Primitive.from_proto(resource.gpg_keys),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum.from_proto(
                i
            )
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(object):
    def __init__(
        self,
        id: str = None,
        display_name: str = None,
        base_url: str = None,
        gpg_keys: list = None,
    ):
        self.id = id
        self.display_name = display_name
        self.base_url = base_url
        self.gpg_keys = gpg_keys

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper()
        )
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.display_name):
            res.display_name = Primitive.to_proto(resource.display_name)
        if Primitive.to_proto(resource.base_url):
            res.base_url = Primitive.to_proto(resource.base_url)
        if Primitive.to_proto(resource.gpg_keys):
            res.gpg_keys.extend(Primitive.to_proto(resource.gpg_keys))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper(
            id=Primitive.from_proto(resource.id),
            display_name=Primitive.from_proto(resource.display_name),
            base_url=Primitive.from_proto(resource.base_url),
            gpg_keys=Primitive.from_proto(resource.gpg_keys),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper.from_proto(
                i
            )
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(object):
    def __init__(self, name: str = None, url: str = None):
        self.name = name
        self.url = url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo(
            name=Primitive.from_proto(resource.name),
            url=Primitive.from_proto(resource.url),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo.from_proto(
                i
            )
            for i in resources
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(object):
    def __init__(self, validate: dict = None, enforce: dict = None):
        self.validate = validate
        self.enforce = enforce

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec()
        )
        if OsPolicyAssignmentExec.to_proto(resource.validate):
            res.validate.CopyFrom(OsPolicyAssignmentExec.to_proto(resource.validate))
        else:
            res.ClearField("validate")
        if OsPolicyAssignmentExec.to_proto(resource.enforce):
            res.enforce.CopyFrom(OsPolicyAssignmentExec.to_proto(resource.enforce))
        else:
            res.ClearField("enforce")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec(
            validate=OsPolicyAssignmentExec.from_proto(resource.validate),
            enforce=OsPolicyAssignmentExec.from_proto(resource.enforce),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesExec.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentExec(object):
    def __init__(
        self,
        file: dict = None,
        script: str = None,
        args: list = None,
        interpreter: str = None,
    ):
        self.file = file
        self.script = script
        self.args = args
        self.interpreter = interpreter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentExec()
        if OsPolicyAssignmentFile.to_proto(resource.file):
            res.file.CopyFrom(OsPolicyAssignmentFile.to_proto(resource.file))
        else:
            res.ClearField("file")
        if Primitive.to_proto(resource.script):
            res.script = Primitive.to_proto(resource.script)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if OsPolicyAssignmentExecInterpreterEnum.to_proto(resource.interpreter):
            res.interpreter = OsPolicyAssignmentExecInterpreterEnum.to_proto(
                resource.interpreter
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentExec(
            file=OsPolicyAssignmentFile.from_proto(resource.file),
            script=Primitive.from_proto(resource.script),
            args=Primitive.from_proto(resource.args),
            interpreter=OsPolicyAssignmentExecInterpreterEnum.from_proto(
                resource.interpreter
            ),
        )


class OsPolicyAssignmentExecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OsPolicyAssignmentExec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OsPolicyAssignmentExec.from_proto(i) for i in resources]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(object):
    def __init__(
        self,
        file: dict = None,
        content: str = None,
        path: str = None,
        state: str = None,
        permissions: str = None,
    ):
        self.file = file
        self.content = content
        self.path = path
        self.state = state
        self.permissions = permissions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile()
        )
        if OsPolicyAssignmentFile.to_proto(resource.file):
            res.file.CopyFrom(OsPolicyAssignmentFile.to_proto(resource.file))
        else:
            res.ClearField("file")
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum.to_proto(
            resource.state
        ):
            res.state = OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum.to_proto(
                resource.state
            )
        if Primitive.to_proto(resource.permissions):
            res.permissions = Primitive.to_proto(resource.permissions)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile(
            file=OsPolicyAssignmentFile.from_proto(resource.file),
            content=Primitive.from_proto(resource.content),
            path=Primitive.from_proto(resource.path),
            state=OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum.from_proto(
                resource.state
            ),
            permissions=Primitive.from_proto(resource.permissions),
        )


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFile.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentInstanceFilter(object):
    def __init__(
        self,
        all: bool = None,
        os_short_names: list = None,
        inclusion_labels: list = None,
        exclusion_labels: list = None,
    ):
        self.all = all
        self.os_short_names = os_short_names
        self.inclusion_labels = inclusion_labels
        self.exclusion_labels = exclusion_labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentInstanceFilter()
        if Primitive.to_proto(resource.all):
            res.all = Primitive.to_proto(resource.all)
        if Primitive.to_proto(resource.os_short_names):
            res.os_short_names.extend(Primitive.to_proto(resource.os_short_names))
        if OsPolicyAssignmentInstanceFilterInclusionLabelsArray.to_proto(
            resource.inclusion_labels
        ):
            res.inclusion_labels.extend(
                OsPolicyAssignmentInstanceFilterInclusionLabelsArray.to_proto(
                    resource.inclusion_labels
                )
            )
        if OsPolicyAssignmentInstanceFilterExclusionLabelsArray.to_proto(
            resource.exclusion_labels
        ):
            res.exclusion_labels.extend(
                OsPolicyAssignmentInstanceFilterExclusionLabelsArray.to_proto(
                    resource.exclusion_labels
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentInstanceFilter(
            all=Primitive.from_proto(resource.all),
            os_short_names=Primitive.from_proto(resource.os_short_names),
            inclusion_labels=OsPolicyAssignmentInstanceFilterInclusionLabelsArray.from_proto(
                resource.inclusion_labels
            ),
            exclusion_labels=OsPolicyAssignmentInstanceFilterExclusionLabelsArray.from_proto(
                resource.exclusion_labels
            ),
        )


class OsPolicyAssignmentInstanceFilterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OsPolicyAssignmentInstanceFilter.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OsPolicyAssignmentInstanceFilter.from_proto(i) for i in resources]


class OsPolicyAssignmentInstanceFilterInclusionLabels(object):
    def __init__(self, labels: dict = None):
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentInstanceFilterInclusionLabels()
        )
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentInstanceFilterInclusionLabels(
            labels=Primitive.from_proto(resource.labels),
        )


class OsPolicyAssignmentInstanceFilterInclusionLabelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentInstanceFilterInclusionLabels.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentInstanceFilterInclusionLabels.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentInstanceFilterExclusionLabels(object):
    def __init__(self, labels: dict = None):
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentInstanceFilterExclusionLabels()
        )
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentInstanceFilterExclusionLabels(
            labels=Primitive.from_proto(resource.labels),
        )


class OsPolicyAssignmentInstanceFilterExclusionLabelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentInstanceFilterExclusionLabels.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentInstanceFilterExclusionLabels.from_proto(i)
            for i in resources
        ]


class OsPolicyAssignmentRollout(object):
    def __init__(self, disruption_budget: dict = None, min_wait_duration: str = None):
        self.disruption_budget = disruption_budget
        self.min_wait_duration = min_wait_duration

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentRollout()
        if OsPolicyAssignmentRolloutDisruptionBudget.to_proto(
            resource.disruption_budget
        ):
            res.disruption_budget.CopyFrom(
                OsPolicyAssignmentRolloutDisruptionBudget.to_proto(
                    resource.disruption_budget
                )
            )
        else:
            res.ClearField("disruption_budget")
        if Primitive.to_proto(resource.min_wait_duration):
            res.min_wait_duration = Primitive.to_proto(resource.min_wait_duration)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentRollout(
            disruption_budget=OsPolicyAssignmentRolloutDisruptionBudget.from_proto(
                resource.disruption_budget
            ),
            min_wait_duration=Primitive.from_proto(resource.min_wait_duration),
        )


class OsPolicyAssignmentRolloutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OsPolicyAssignmentRollout.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OsPolicyAssignmentRollout.from_proto(i) for i in resources]


class OsPolicyAssignmentRolloutDisruptionBudget(object):
    def __init__(self, fixed: int = None, percent: int = None):
        self.fixed = fixed
        self.percent = percent

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentRolloutDisruptionBudget()
        )
        if Primitive.to_proto(resource.fixed):
            res.fixed = Primitive.to_proto(resource.fixed)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OsPolicyAssignmentRolloutDisruptionBudget(
            fixed=Primitive.from_proto(resource.fixed),
            percent=Primitive.from_proto(resource.percent),
        )


class OsPolicyAssignmentRolloutDisruptionBudgetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OsPolicyAssignmentRolloutDisruptionBudget.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OsPolicyAssignmentRolloutDisruptionBudget.from_proto(i) for i in resources
        ]


class OsPolicyAssignmentOsPoliciesModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnum.Value(
            "OsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnum.Name(
            resource
        )[
            len("OsconfigAlphaOsPolicyAssignmentOsPoliciesModeEnum") :
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum.Value(
            "OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum.Name(
            resource
        )[
            len(
                "OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDesiredStateEnum"
            ) :
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.Value(
            "OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.Name(
            resource
        )[
            len(
                "OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum"
            ) :
        ]


class OsPolicyAssignmentExecInterpreterEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentExecInterpreterEnum.Value(
            "OsconfigAlphaOsPolicyAssignmentExecInterpreterEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentExecInterpreterEnum.Name(
            resource
        )[
            len("OsconfigAlphaOsPolicyAssignmentExecInterpreterEnum") :
        ]


class OsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum.Value(
            "OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum.Name(
            resource
        )[
            len(
                "OsconfigAlphaOsPolicyAssignmentOsPoliciesResourceGroupsResourcesFileStateEnum"
            ) :
        ]


class OsPolicyAssignmentRolloutStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentRolloutStateEnum.Value(
            "OsconfigAlphaOsPolicyAssignmentRolloutStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOsPolicyAssignmentRolloutStateEnum.Name(
            resource
        )[
            len("OsconfigAlphaOsPolicyAssignmentRolloutStateEnum") :
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
