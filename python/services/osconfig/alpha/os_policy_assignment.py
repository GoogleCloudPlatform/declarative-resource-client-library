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


class OSPolicyAssignment(object):
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
        stub = os_policy_assignment_pb2_grpc.OsconfigAlphaOSPolicyAssignmentServiceStub(
            channel.Channel()
        )
        request = os_policy_assignment_pb2.ApplyOsconfigAlphaOSPolicyAssignmentRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if OSPolicyAssignmentOSPoliciesArray.to_proto(self.os_policies):
            request.resource.os_policies.extend(
                OSPolicyAssignmentOSPoliciesArray.to_proto(self.os_policies)
            )
        if OSPolicyAssignmentInstanceFilter.to_proto(self.instance_filter):
            request.resource.instance_filter.CopyFrom(
                OSPolicyAssignmentInstanceFilter.to_proto(self.instance_filter)
            )
        else:
            request.resource.ClearField("instance_filter")
        if OSPolicyAssignmentRollout.to_proto(self.rollout):
            request.resource.rollout.CopyFrom(
                OSPolicyAssignmentRollout.to_proto(self.rollout)
            )
        else:
            request.resource.ClearField("rollout")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyOsconfigAlphaOSPolicyAssignment(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.os_policies = OSPolicyAssignmentOSPoliciesArray.from_proto(
            response.os_policies
        )
        self.instance_filter = OSPolicyAssignmentInstanceFilter.from_proto(
            response.instance_filter
        )
        self.rollout = OSPolicyAssignmentRollout.from_proto(response.rollout)
        self.revision_id = Primitive.from_proto(response.revision_id)
        self.revision_create_time = Primitive.from_proto(response.revision_create_time)
        self.rollout_state = OSPolicyAssignmentRolloutStateEnum.from_proto(
            response.rollout_state
        )
        self.baseline = Primitive.from_proto(response.baseline)
        self.deleted = Primitive.from_proto(response.deleted)
        self.reconciling = Primitive.from_proto(response.reconciling)
        self.uid = Primitive.from_proto(response.uid)
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = os_policy_assignment_pb2_grpc.OsconfigAlphaOSPolicyAssignmentServiceStub(
            channel.Channel()
        )
        request = (
            os_policy_assignment_pb2.DeleteOsconfigAlphaOSPolicyAssignmentRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if OSPolicyAssignmentOSPoliciesArray.to_proto(self.os_policies):
            request.resource.os_policies.extend(
                OSPolicyAssignmentOSPoliciesArray.to_proto(self.os_policies)
            )
        if OSPolicyAssignmentInstanceFilter.to_proto(self.instance_filter):
            request.resource.instance_filter.CopyFrom(
                OSPolicyAssignmentInstanceFilter.to_proto(self.instance_filter)
            )
        else:
            request.resource.ClearField("instance_filter")
        if OSPolicyAssignmentRollout.to_proto(self.rollout):
            request.resource.rollout.CopyFrom(
                OSPolicyAssignmentRollout.to_proto(self.rollout)
            )
        else:
            request.resource.ClearField("rollout")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteOsconfigAlphaOSPolicyAssignment(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = os_policy_assignment_pb2_grpc.OsconfigAlphaOSPolicyAssignmentServiceStub(
            channel.Channel()
        )
        request = os_policy_assignment_pb2.ListOsconfigAlphaOSPolicyAssignmentRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListOsconfigAlphaOSPolicyAssignment(request).items

    def to_proto(self):
        resource = os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignment()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if OSPolicyAssignmentOSPoliciesArray.to_proto(self.os_policies):
            resource.os_policies.extend(
                OSPolicyAssignmentOSPoliciesArray.to_proto(self.os_policies)
            )
        if OSPolicyAssignmentInstanceFilter.to_proto(self.instance_filter):
            resource.instance_filter.CopyFrom(
                OSPolicyAssignmentInstanceFilter.to_proto(self.instance_filter)
            )
        else:
            resource.ClearField("instance_filter")
        if OSPolicyAssignmentRollout.to_proto(self.rollout):
            resource.rollout.CopyFrom(OSPolicyAssignmentRollout.to_proto(self.rollout))
        else:
            resource.ClearField("rollout")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class OSPolicyAssignmentOSPolicies(object):
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

        res = os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPolicies()
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if OSPolicyAssignmentOSPoliciesModeEnum.to_proto(resource.mode):
            res.mode = OSPolicyAssignmentOSPoliciesModeEnum.to_proto(resource.mode)
        if OSPolicyAssignmentOSPoliciesResourceGroupsArray.to_proto(
            resource.resource_groups
        ):
            res.resource_groups.extend(
                OSPolicyAssignmentOSPoliciesResourceGroupsArray.to_proto(
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

        return OSPolicyAssignmentOSPolicies(
            id=Primitive.from_proto(resource.id),
            description=Primitive.from_proto(resource.description),
            mode=OSPolicyAssignmentOSPoliciesModeEnum.from_proto(resource.mode),
            resource_groups=OSPolicyAssignmentOSPoliciesResourceGroupsArray.from_proto(
                resource.resource_groups
            ),
            allow_no_resource_group_match=Primitive.from_proto(
                resource.allow_no_resource_group_match
            ),
        )


class OSPolicyAssignmentOSPoliciesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OSPolicyAssignmentOSPolicies.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OSPolicyAssignmentOSPolicies.from_proto(i) for i in resources]


class OSPolicyAssignmentOSPoliciesResourceGroups(object):
    def __init__(self, os_filter: dict = None, resources: list = None):
        self.os_filter = os_filter
        self.resources = resources

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroups()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsOSFilter.to_proto(
            resource.os_filter
        ):
            res.os_filter.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsOSFilter.to_proto(
                    resource.os_filter
                )
            )
        else:
            res.ClearField("os_filter")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesArray.to_proto(
            resource.resources
        ):
            res.resources.extend(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesArray.to_proto(
                    resource.resources
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroups(
            os_filter=OSPolicyAssignmentOSPoliciesResourceGroupsOSFilter.from_proto(
                resource.os_filter
            ),
            resources=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesArray.from_proto(
                resource.resources
            ),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroups.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroups.from_proto(i) for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsOSFilter(object):
    def __init__(self, os_short_name: str = None, os_version: str = None):
        self.os_short_name = os_short_name
        self.os_version = os_version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsOSFilter()
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

        return OSPolicyAssignmentOSPoliciesResourceGroupsOSFilter(
            os_short_name=Primitive.from_proto(resource.os_short_name),
            os_version=Primitive.from_proto(resource.os_version),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsOSFilterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsOSFilter.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsOSFilter.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResources(object):
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
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResources()
        )
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg.to_proto(
            resource.pkg
        ):
            res.pkg.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg.to_proto(
                    resource.pkg
                )
            )
        else:
            res.ClearField("pkg")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository.to_proto(
            resource.repository
        ):
            res.repository.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository.to_proto(
                    resource.repository
                )
            )
        else:
            res.ClearField("repository")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec.to_proto(
            resource.exec
        ):
            res.exec.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec.to_proto(
                    resource.exec
                )
            )
        else:
            res.ClearField("exec")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile.to_proto(
            resource.file
        ):
            res.file.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile.to_proto(
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

        return OSPolicyAssignmentOSPoliciesResourceGroupsResources(
            id=Primitive.from_proto(resource.id),
            pkg=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg.from_proto(
                resource.pkg
            ),
            repository=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository.from_proto(
                resource.repository
            ),
            exec=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec.from_proto(
                resource.exec
            ),
            file=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile.from_proto(
                resource.file
            ),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResources.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResources.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg(object):
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
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum.to_proto(
            resource.desired_state
        ):
            res.desired_state = OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum.to_proto(
                resource.desired_state
            )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt.to_proto(
            resource.apt
        ):
            res.apt.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt.to_proto(
                    resource.apt
                )
            )
        else:
            res.ClearField("apt")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb.to_proto(
            resource.deb
        ):
            res.deb.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb.to_proto(
                    resource.deb
                )
            )
        else:
            res.ClearField("deb")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum.to_proto(
            resource.yum
        ):
            res.yum.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum.to_proto(
                    resource.yum
                )
            )
        else:
            res.ClearField("yum")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper.to_proto(
            resource.zypper
        ):
            res.zypper.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper.to_proto(
                    resource.zypper
                )
            )
        else:
            res.ClearField("zypper")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm.to_proto(
            resource.rpm
        ):
            res.rpm.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm.to_proto(
                    resource.rpm
                )
            )
        else:
            res.ClearField("rpm")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget.to_proto(
            resource.googet
        ):
            res.googet.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget.to_proto(
                    resource.googet
                )
            )
        else:
            res.ClearField("googet")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi.to_proto(
            resource.msi
        ):
            res.msi.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi.to_proto(
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

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg(
            desired_state=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum.from_proto(
                resource.desired_state
            ),
            apt=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt.from_proto(
                resource.apt
            ),
            deb=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb.from_proto(
                resource.deb
            ),
            yum=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum.from_proto(
                resource.yum
            ),
            zypper=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper.from_proto(
                resource.zypper
            ),
            rpm=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm.from_proto(
                resource.rpm
            ),
            googet=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget.from_proto(
                resource.googet
            ),
            msi=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi.from_proto(
                resource.msi
            ),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt(
            name=Primitive.from_proto(resource.name),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgAptArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb(object):
    def __init__(self, source: dict = None, pull_deps: bool = None):
        self.source = source
        self.pull_deps = pull_deps

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource.to_proto(
            resource.source
        ):
            res.source.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource.to_proto(
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

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb(
            source=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource.from_proto(
                resource.source
            ),
            pull_deps=Primitive.from_proto(resource.pull_deps),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource(object):
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
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote.to_proto(
            resource.remote
        ):
            res.remote.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote.to_proto(
                    resource.remote
                )
            )
        else:
            res.ClearField("remote")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs.to_proto(
            resource.gcs
        ):
            res.gcs.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs.to_proto(
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

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource(
            remote=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote.from_proto(
                resource.remote
            ),
            gcs=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs.from_proto(
                resource.gcs
            ),
            local_path=Primitive.from_proto(resource.local_path),
            allow_insecure=Primitive.from_proto(resource.allow_insecure),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote(object):
    def __init__(self, uri: str = None, sha256_checksum: str = None):
        self.uri = uri
        self.sha256_checksum = sha256_checksum

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote()
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

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote(
            uri=Primitive.from_proto(resource.uri),
            sha256_checksum=Primitive.from_proto(resource.sha256_checksum),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemoteArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs(object):
    def __init__(self, bucket: str = None, object: str = None, generation: int = None):
        self.bucket = bucket
        self.object = object
        self.generation = generation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs()
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

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation=Primitive.from_proto(resource.generation),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum(
            name=Primitive.from_proto(resource.name),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYumArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper(
            name=Primitive.from_proto(resource.name),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypperArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm(object):
    def __init__(self, source: dict = None, pull_deps: bool = None):
        self.source = source
        self.pull_deps = pull_deps

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm()
        )
        if OSPolicyAssignmentFile.to_proto(resource.source):
            res.source.CopyFrom(OSPolicyAssignmentFile.to_proto(resource.source))
        else:
            res.ClearField("source")
        if Primitive.to_proto(resource.pull_deps):
            res.pull_deps = Primitive.to_proto(resource.pull_deps)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm(
            source=OSPolicyAssignmentFile.from_proto(resource.source),
            pull_deps=Primitive.from_proto(resource.pull_deps),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentFile(object):
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

        res = os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentFile()
        if OSPolicyAssignmentFileRemote.to_proto(resource.remote):
            res.remote.CopyFrom(OSPolicyAssignmentFileRemote.to_proto(resource.remote))
        else:
            res.ClearField("remote")
        if OSPolicyAssignmentFileGcs.to_proto(resource.gcs):
            res.gcs.CopyFrom(OSPolicyAssignmentFileGcs.to_proto(resource.gcs))
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

        return OSPolicyAssignmentFile(
            remote=OSPolicyAssignmentFileRemote.from_proto(resource.remote),
            gcs=OSPolicyAssignmentFileGcs.from_proto(resource.gcs),
            local_path=Primitive.from_proto(resource.local_path),
            allow_insecure=Primitive.from_proto(resource.allow_insecure),
        )


class OSPolicyAssignmentFileArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OSPolicyAssignmentFile.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OSPolicyAssignmentFile.from_proto(i) for i in resources]


class OSPolicyAssignmentFileRemote(object):
    def __init__(self, uri: str = None, sha256_checksum: str = None):
        self.uri = uri
        self.sha256_checksum = sha256_checksum

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentFileRemote()
        if Primitive.to_proto(resource.uri):
            res.uri = Primitive.to_proto(resource.uri)
        if Primitive.to_proto(resource.sha256_checksum):
            res.sha256_checksum = Primitive.to_proto(resource.sha256_checksum)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentFileRemote(
            uri=Primitive.from_proto(resource.uri),
            sha256_checksum=Primitive.from_proto(resource.sha256_checksum),
        )


class OSPolicyAssignmentFileRemoteArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OSPolicyAssignmentFileRemote.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OSPolicyAssignmentFileRemote.from_proto(i) for i in resources]


class OSPolicyAssignmentFileGcs(object):
    def __init__(self, bucket: str = None, object: str = None, generation: int = None):
        self.bucket = bucket
        self.object = object
        self.generation = generation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentFileGcs()
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

        return OSPolicyAssignmentFileGcs(
            bucket=Primitive.from_proto(resource.bucket),
            object=Primitive.from_proto(resource.object),
            generation=Primitive.from_proto(resource.generation),
        )


class OSPolicyAssignmentFileGcsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OSPolicyAssignmentFileGcs.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OSPolicyAssignmentFileGcs.from_proto(i) for i in resources]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget(object):
    def __init__(self, name: str = None):
        self.name = name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget()
        )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget(
            name=Primitive.from_proto(resource.name),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGoogetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi(object):
    def __init__(self, source: dict = None, properties: list = None):
        self.source = source
        self.properties = properties

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi()
        )
        if OSPolicyAssignmentFile.to_proto(resource.source):
            res.source.CopyFrom(OSPolicyAssignmentFile.to_proto(resource.source))
        else:
            res.ClearField("source")
        if Primitive.to_proto(resource.properties):
            res.properties.extend(Primitive.to_proto(resource.properties))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi(
            source=OSPolicyAssignmentFile.from_proto(resource.source),
            properties=Primitive.from_proto(resource.properties),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository(object):
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
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt.to_proto(
            resource.apt
        ):
            res.apt.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt.to_proto(
                    resource.apt
                )
            )
        else:
            res.ClearField("apt")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum.to_proto(
            resource.yum
        ):
            res.yum.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum.to_proto(
                    resource.yum
                )
            )
        else:
            res.ClearField("yum")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper.to_proto(
            resource.zypper
        ):
            res.zypper.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper.to_proto(
                    resource.zypper
                )
            )
        else:
            res.ClearField("zypper")
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo.to_proto(
            resource.goo
        ):
            res.goo.CopyFrom(
                OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo.to_proto(
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

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository(
            apt=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt.from_proto(
                resource.apt
            ),
            yum=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum.from_proto(
                resource.yum
            ),
            zypper=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper.from_proto(
                resource.zypper
            ),
            goo=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo.from_proto(
                resource.goo
            ),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt(object):
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
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt()
        )
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.to_proto(
            resource.archive_type
        ):
            res.archive_type = OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.to_proto(
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

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt(
            archive_type=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.from_proto(
                resource.archive_type
            ),
            uri=Primitive.from_proto(resource.uri),
            distribution=Primitive.from_proto(resource.distribution),
            components=Primitive.from_proto(resource.components),
            gpg_key=Primitive.from_proto(resource.gpg_key),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum(object):
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
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum()
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

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum(
            id=Primitive.from_proto(resource.id),
            display_name=Primitive.from_proto(resource.display_name),
            base_url=Primitive.from_proto(resource.base_url),
            gpg_keys=Primitive.from_proto(resource.gpg_keys),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYumArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper(object):
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
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper()
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

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper(
            id=Primitive.from_proto(resource.id),
            display_name=Primitive.from_proto(resource.display_name),
            base_url=Primitive.from_proto(resource.base_url),
            gpg_keys=Primitive.from_proto(resource.gpg_keys),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypperArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo(object):
    def __init__(self, name: str = None, url: str = None):
        self.name = name
        self.url = url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo()
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

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo(
            name=Primitive.from_proto(resource.name),
            url=Primitive.from_proto(resource.url),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGooArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo.from_proto(
                i
            )
            for i in resources
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec(object):
    def __init__(self, validate: dict = None, enforce: dict = None):
        self.validate = validate
        self.enforce = enforce

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec()
        )
        if OSPolicyAssignmentExec.to_proto(resource.validate):
            res.validate.CopyFrom(OSPolicyAssignmentExec.to_proto(resource.validate))
        else:
            res.ClearField("validate")
        if OSPolicyAssignmentExec.to_proto(resource.enforce):
            res.enforce.CopyFrom(OSPolicyAssignmentExec.to_proto(resource.enforce))
        else:
            res.ClearField("enforce")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec(
            validate=OSPolicyAssignmentExec.from_proto(resource.validate),
            enforce=OSPolicyAssignmentExec.from_proto(resource.enforce),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentExec(object):
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

        res = os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentExec()
        if OSPolicyAssignmentFile.to_proto(resource.file):
            res.file.CopyFrom(OSPolicyAssignmentFile.to_proto(resource.file))
        else:
            res.ClearField("file")
        if Primitive.to_proto(resource.script):
            res.script = Primitive.to_proto(resource.script)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if OSPolicyAssignmentExecInterpreterEnum.to_proto(resource.interpreter):
            res.interpreter = OSPolicyAssignmentExecInterpreterEnum.to_proto(
                resource.interpreter
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentExec(
            file=OSPolicyAssignmentFile.from_proto(resource.file),
            script=Primitive.from_proto(resource.script),
            args=Primitive.from_proto(resource.args),
            interpreter=OSPolicyAssignmentExecInterpreterEnum.from_proto(
                resource.interpreter
            ),
        )


class OSPolicyAssignmentExecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OSPolicyAssignmentExec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OSPolicyAssignmentExec.from_proto(i) for i in resources]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile(object):
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
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile()
        )
        if OSPolicyAssignmentFile.to_proto(resource.file):
            res.file.CopyFrom(OSPolicyAssignmentFile.to_proto(resource.file))
        else:
            res.ClearField("file")
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum.to_proto(
            resource.state
        ):
            res.state = OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum.to_proto(
                resource.state
            )
        if Primitive.to_proto(resource.permissions):
            res.permissions = Primitive.to_proto(resource.permissions)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile(
            file=OSPolicyAssignmentFile.from_proto(resource.file),
            content=Primitive.from_proto(resource.content),
            path=Primitive.from_proto(resource.path),
            state=OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum.from_proto(
                resource.state
            ),
            permissions=Primitive.from_proto(resource.permissions),
        )


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentInstanceFilter(object):
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

        res = os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentInstanceFilter()
        if Primitive.to_proto(resource.all):
            res.all = Primitive.to_proto(resource.all)
        if Primitive.to_proto(resource.os_short_names):
            res.os_short_names.extend(Primitive.to_proto(resource.os_short_names))
        if OSPolicyAssignmentInstanceFilterInclusionLabelsArray.to_proto(
            resource.inclusion_labels
        ):
            res.inclusion_labels.extend(
                OSPolicyAssignmentInstanceFilterInclusionLabelsArray.to_proto(
                    resource.inclusion_labels
                )
            )
        if OSPolicyAssignmentInstanceFilterExclusionLabelsArray.to_proto(
            resource.exclusion_labels
        ):
            res.exclusion_labels.extend(
                OSPolicyAssignmentInstanceFilterExclusionLabelsArray.to_proto(
                    resource.exclusion_labels
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentInstanceFilter(
            all=Primitive.from_proto(resource.all),
            os_short_names=Primitive.from_proto(resource.os_short_names),
            inclusion_labels=OSPolicyAssignmentInstanceFilterInclusionLabelsArray.from_proto(
                resource.inclusion_labels
            ),
            exclusion_labels=OSPolicyAssignmentInstanceFilterExclusionLabelsArray.from_proto(
                resource.exclusion_labels
            ),
        )


class OSPolicyAssignmentInstanceFilterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OSPolicyAssignmentInstanceFilter.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OSPolicyAssignmentInstanceFilter.from_proto(i) for i in resources]


class OSPolicyAssignmentInstanceFilterInclusionLabels(object):
    def __init__(self, labels: dict = None):
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabels()
        )
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentInstanceFilterInclusionLabels(
            labels=Primitive.from_proto(resource.labels),
        )


class OSPolicyAssignmentInstanceFilterInclusionLabelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentInstanceFilterInclusionLabels.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentInstanceFilterInclusionLabels.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentInstanceFilterExclusionLabels(object):
    def __init__(self, labels: dict = None):
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabels()
        )
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OSPolicyAssignmentInstanceFilterExclusionLabels(
            labels=Primitive.from_proto(resource.labels),
        )


class OSPolicyAssignmentInstanceFilterExclusionLabelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentInstanceFilterExclusionLabels.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentInstanceFilterExclusionLabels.from_proto(i)
            for i in resources
        ]


class OSPolicyAssignmentRollout(object):
    def __init__(self, disruption_budget: dict = None, min_wait_duration: str = None):
        self.disruption_budget = disruption_budget
        self.min_wait_duration = min_wait_duration

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentRollout()
        if OSPolicyAssignmentRolloutDisruptionBudget.to_proto(
            resource.disruption_budget
        ):
            res.disruption_budget.CopyFrom(
                OSPolicyAssignmentRolloutDisruptionBudget.to_proto(
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

        return OSPolicyAssignmentRollout(
            disruption_budget=OSPolicyAssignmentRolloutDisruptionBudget.from_proto(
                resource.disruption_budget
            ),
            min_wait_duration=Primitive.from_proto(resource.min_wait_duration),
        )


class OSPolicyAssignmentRolloutArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OSPolicyAssignmentRollout.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OSPolicyAssignmentRollout.from_proto(i) for i in resources]


class OSPolicyAssignmentRolloutDisruptionBudget(object):
    def __init__(self, fixed: int = None, percent: int = None):
        self.fixed = fixed
        self.percent = percent

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudget()
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

        return OSPolicyAssignmentRolloutDisruptionBudget(
            fixed=Primitive.from_proto(resource.fixed),
            percent=Primitive.from_proto(resource.percent),
        )


class OSPolicyAssignmentRolloutDisruptionBudgetArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            OSPolicyAssignmentRolloutDisruptionBudget.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            OSPolicyAssignmentRolloutDisruptionBudget.from_proto(i) for i in resources
        ]


class OSPolicyAssignmentOSPoliciesModeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum.Value(
            "OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum.Name(
            resource
        )[
            len("OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum") :
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum.Value(
            "OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum.Name(
            resource
        )[
            len(
                "OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum"
            ) :
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.Value(
            "OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum.Name(
            resource
        )[
            len(
                "OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum"
            ) :
        ]


class OSPolicyAssignmentExecInterpreterEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentExecInterpreterEnum.Value(
            "OsconfigAlphaOSPolicyAssignmentExecInterpreterEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentExecInterpreterEnum.Name(
            resource
        )[
            len("OsconfigAlphaOSPolicyAssignmentExecInterpreterEnum") :
        ]


class OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum.Value(
            "OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum.Name(
            resource
        )[
            len(
                "OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum"
            ) :
        ]


class OSPolicyAssignmentRolloutStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum.Value(
            "OsconfigAlphaOSPolicyAssignmentRolloutStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return os_policy_assignment_pb2.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum.Name(
            resource
        )[
            len("OsconfigAlphaOSPolicyAssignmentRolloutStateEnum") :
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
