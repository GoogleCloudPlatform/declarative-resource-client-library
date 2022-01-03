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
from google3.cloud.graphite.mmv2.services.google.gke_hub import feature_membership_pb2
from google3.cloud.graphite.mmv2.services.google.gke_hub import (
    feature_membership_pb2_grpc,
)

from typing import List


class FeatureMembership(object):
    def __init__(
        self,
        configmanagement: dict = None,
        project: str = None,
        location: str = None,
        feature: str = None,
        membership: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.configmanagement = configmanagement
        self.project = project
        self.location = location
        self.feature = feature
        self.membership = membership
        self.service_account_file = service_account_file

    def apply(self):
        stub = feature_membership_pb2_grpc.GkehubBetaFeatureMembershipServiceStub(
            channel.Channel()
        )
        request = feature_membership_pb2.ApplyGkehubBetaFeatureMembershipRequest()
        if FeatureMembershipConfigmanagement.to_proto(self.configmanagement):
            request.resource.configmanagement.CopyFrom(
                FeatureMembershipConfigmanagement.to_proto(self.configmanagement)
            )
        else:
            request.resource.ClearField("configmanagement")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.feature):
            request.resource.feature = Primitive.to_proto(self.feature)

        if Primitive.to_proto(self.membership):
            request.resource.membership = Primitive.to_proto(self.membership)

        request.service_account_file = self.service_account_file

        response = stub.ApplyGkehubBetaFeatureMembership(request)
        self.configmanagement = FeatureMembershipConfigmanagement.from_proto(
            response.configmanagement
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)
        self.feature = Primitive.from_proto(response.feature)
        self.membership = Primitive.from_proto(response.membership)

    def delete(self):
        stub = feature_membership_pb2_grpc.GkehubBetaFeatureMembershipServiceStub(
            channel.Channel()
        )
        request = feature_membership_pb2.DeleteGkehubBetaFeatureMembershipRequest()
        request.service_account_file = self.service_account_file
        if FeatureMembershipConfigmanagement.to_proto(self.configmanagement):
            request.resource.configmanagement.CopyFrom(
                FeatureMembershipConfigmanagement.to_proto(self.configmanagement)
            )
        else:
            request.resource.ClearField("configmanagement")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if Primitive.to_proto(self.feature):
            request.resource.feature = Primitive.to_proto(self.feature)

        if Primitive.to_proto(self.membership):
            request.resource.membership = Primitive.to_proto(self.membership)

        response = stub.DeleteGkehubBetaFeatureMembership(request)

    @classmethod
    def list(self, project, location, feature, service_account_file=""):
        stub = feature_membership_pb2_grpc.GkehubBetaFeatureMembershipServiceStub(
            channel.Channel()
        )
        request = feature_membership_pb2.ListGkehubBetaFeatureMembershipRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        request.Feature = feature

        return stub.ListGkehubBetaFeatureMembership(request).items

    def to_proto(self):
        resource = feature_membership_pb2.GkehubBetaFeatureMembership()
        if FeatureMembershipConfigmanagement.to_proto(self.configmanagement):
            resource.configmanagement.CopyFrom(
                FeatureMembershipConfigmanagement.to_proto(self.configmanagement)
            )
        else:
            resource.ClearField("configmanagement")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        if Primitive.to_proto(self.feature):
            resource.feature = Primitive.to_proto(self.feature)
        if Primitive.to_proto(self.membership):
            resource.membership = Primitive.to_proto(self.membership)
        return resource


class FeatureMembershipConfigmanagement(object):
    def __init__(
        self,
        config_sync: dict = None,
        policy_controller: dict = None,
        binauthz: dict = None,
        hierarchy_controller: dict = None,
        version: str = None,
    ):
        self.config_sync = config_sync
        self.policy_controller = policy_controller
        self.binauthz = binauthz
        self.hierarchy_controller = hierarchy_controller
        self.version = version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = feature_membership_pb2.GkehubBetaFeatureMembershipConfigmanagement()
        if FeatureMembershipConfigmanagementConfigSync.to_proto(resource.config_sync):
            res.config_sync.CopyFrom(
                FeatureMembershipConfigmanagementConfigSync.to_proto(
                    resource.config_sync
                )
            )
        else:
            res.ClearField("config_sync")
        if FeatureMembershipConfigmanagementPolicyController.to_proto(
            resource.policy_controller
        ):
            res.policy_controller.CopyFrom(
                FeatureMembershipConfigmanagementPolicyController.to_proto(
                    resource.policy_controller
                )
            )
        else:
            res.ClearField("policy_controller")
        if FeatureMembershipConfigmanagementBinauthz.to_proto(resource.binauthz):
            res.binauthz.CopyFrom(
                FeatureMembershipConfigmanagementBinauthz.to_proto(resource.binauthz)
            )
        else:
            res.ClearField("binauthz")
        if FeatureMembershipConfigmanagementHierarchyController.to_proto(
            resource.hierarchy_controller
        ):
            res.hierarchy_controller.CopyFrom(
                FeatureMembershipConfigmanagementHierarchyController.to_proto(
                    resource.hierarchy_controller
                )
            )
        else:
            res.ClearField("hierarchy_controller")
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureMembershipConfigmanagement(
            config_sync=FeatureMembershipConfigmanagementConfigSync.from_proto(
                resource.config_sync
            ),
            policy_controller=FeatureMembershipConfigmanagementPolicyController.from_proto(
                resource.policy_controller
            ),
            binauthz=FeatureMembershipConfigmanagementBinauthz.from_proto(
                resource.binauthz
            ),
            hierarchy_controller=FeatureMembershipConfigmanagementHierarchyController.from_proto(
                resource.hierarchy_controller
            ),
            version=Primitive.from_proto(resource.version),
        )


class FeatureMembershipConfigmanagementArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [FeatureMembershipConfigmanagement.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [FeatureMembershipConfigmanagement.from_proto(i) for i in resources]


class FeatureMembershipConfigmanagementConfigSync(object):
    def __init__(self, git: dict = None, source_format: str = None):
        self.git = git
        self.source_format = source_format

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            feature_membership_pb2.GkehubBetaFeatureMembershipConfigmanagementConfigSync()
        )
        if FeatureMembershipConfigmanagementConfigSyncGit.to_proto(resource.git):
            res.git.CopyFrom(
                FeatureMembershipConfigmanagementConfigSyncGit.to_proto(resource.git)
            )
        else:
            res.ClearField("git")
        if Primitive.to_proto(resource.source_format):
            res.source_format = Primitive.to_proto(resource.source_format)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureMembershipConfigmanagementConfigSync(
            git=FeatureMembershipConfigmanagementConfigSyncGit.from_proto(resource.git),
            source_format=Primitive.from_proto(resource.source_format),
        )


class FeatureMembershipConfigmanagementConfigSyncArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            FeatureMembershipConfigmanagementConfigSync.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            FeatureMembershipConfigmanagementConfigSync.from_proto(i) for i in resources
        ]


class FeatureMembershipConfigmanagementConfigSyncGit(object):
    def __init__(
        self,
        sync_repo: str = None,
        sync_branch: str = None,
        policy_dir: str = None,
        sync_wait_secs: str = None,
        sync_rev: str = None,
        secret_type: str = None,
        https_proxy: str = None,
        gcp_service_account_email: str = None,
    ):
        self.sync_repo = sync_repo
        self.sync_branch = sync_branch
        self.policy_dir = policy_dir
        self.sync_wait_secs = sync_wait_secs
        self.sync_rev = sync_rev
        self.secret_type = secret_type
        self.https_proxy = https_proxy
        self.gcp_service_account_email = gcp_service_account_email

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            feature_membership_pb2.GkehubBetaFeatureMembershipConfigmanagementConfigSyncGit()
        )
        if Primitive.to_proto(resource.sync_repo):
            res.sync_repo = Primitive.to_proto(resource.sync_repo)
        if Primitive.to_proto(resource.sync_branch):
            res.sync_branch = Primitive.to_proto(resource.sync_branch)
        if Primitive.to_proto(resource.policy_dir):
            res.policy_dir = Primitive.to_proto(resource.policy_dir)
        if Primitive.to_proto(resource.sync_wait_secs):
            res.sync_wait_secs = Primitive.to_proto(resource.sync_wait_secs)
        if Primitive.to_proto(resource.sync_rev):
            res.sync_rev = Primitive.to_proto(resource.sync_rev)
        if Primitive.to_proto(resource.secret_type):
            res.secret_type = Primitive.to_proto(resource.secret_type)
        if Primitive.to_proto(resource.https_proxy):
            res.https_proxy = Primitive.to_proto(resource.https_proxy)
        if Primitive.to_proto(resource.gcp_service_account_email):
            res.gcp_service_account_email = Primitive.to_proto(
                resource.gcp_service_account_email
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureMembershipConfigmanagementConfigSyncGit(
            sync_repo=Primitive.from_proto(resource.sync_repo),
            sync_branch=Primitive.from_proto(resource.sync_branch),
            policy_dir=Primitive.from_proto(resource.policy_dir),
            sync_wait_secs=Primitive.from_proto(resource.sync_wait_secs),
            sync_rev=Primitive.from_proto(resource.sync_rev),
            secret_type=Primitive.from_proto(resource.secret_type),
            https_proxy=Primitive.from_proto(resource.https_proxy),
            gcp_service_account_email=Primitive.from_proto(
                resource.gcp_service_account_email
            ),
        )


class FeatureMembershipConfigmanagementConfigSyncGitArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            FeatureMembershipConfigmanagementConfigSyncGit.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            FeatureMembershipConfigmanagementConfigSyncGit.from_proto(i)
            for i in resources
        ]


class FeatureMembershipConfigmanagementPolicyController(object):
    def __init__(
        self,
        enabled: bool = None,
        exemptable_namespaces: list = None,
        referential_rules_enabled: bool = None,
        log_denies_enabled: bool = None,
        template_library_installed: bool = None,
        audit_interval_seconds: str = None,
    ):
        self.enabled = enabled
        self.exemptable_namespaces = exemptable_namespaces
        self.referential_rules_enabled = referential_rules_enabled
        self.log_denies_enabled = log_denies_enabled
        self.template_library_installed = template_library_installed
        self.audit_interval_seconds = audit_interval_seconds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            feature_membership_pb2.GkehubBetaFeatureMembershipConfigmanagementPolicyController()
        )
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if Primitive.to_proto(resource.exemptable_namespaces):
            res.exemptable_namespaces.extend(
                Primitive.to_proto(resource.exemptable_namespaces)
            )
        if Primitive.to_proto(resource.referential_rules_enabled):
            res.referential_rules_enabled = Primitive.to_proto(
                resource.referential_rules_enabled
            )
        if Primitive.to_proto(resource.log_denies_enabled):
            res.log_denies_enabled = Primitive.to_proto(resource.log_denies_enabled)
        if Primitive.to_proto(resource.template_library_installed):
            res.template_library_installed = Primitive.to_proto(
                resource.template_library_installed
            )
        if Primitive.to_proto(resource.audit_interval_seconds):
            res.audit_interval_seconds = Primitive.to_proto(
                resource.audit_interval_seconds
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureMembershipConfigmanagementPolicyController(
            enabled=Primitive.from_proto(resource.enabled),
            exemptable_namespaces=Primitive.from_proto(resource.exemptable_namespaces),
            referential_rules_enabled=Primitive.from_proto(
                resource.referential_rules_enabled
            ),
            log_denies_enabled=Primitive.from_proto(resource.log_denies_enabled),
            template_library_installed=Primitive.from_proto(
                resource.template_library_installed
            ),
            audit_interval_seconds=Primitive.from_proto(
                resource.audit_interval_seconds
            ),
        )


class FeatureMembershipConfigmanagementPolicyControllerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            FeatureMembershipConfigmanagementPolicyController.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            FeatureMembershipConfigmanagementPolicyController.from_proto(i)
            for i in resources
        ]


class FeatureMembershipConfigmanagementBinauthz(object):
    def __init__(self, enabled: bool = None):
        self.enabled = enabled

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            feature_membership_pb2.GkehubBetaFeatureMembershipConfigmanagementBinauthz()
        )
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureMembershipConfigmanagementBinauthz(
            enabled=Primitive.from_proto(resource.enabled),
        )


class FeatureMembershipConfigmanagementBinauthzArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            FeatureMembershipConfigmanagementBinauthz.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            FeatureMembershipConfigmanagementBinauthz.from_proto(i) for i in resources
        ]


class FeatureMembershipConfigmanagementHierarchyController(object):
    def __init__(
        self,
        enabled: bool = None,
        enable_pod_tree_labels: bool = None,
        enable_hierarchical_resource_quota: bool = None,
    ):
        self.enabled = enabled
        self.enable_pod_tree_labels = enable_pod_tree_labels
        self.enable_hierarchical_resource_quota = enable_hierarchical_resource_quota

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            feature_membership_pb2.GkehubBetaFeatureMembershipConfigmanagementHierarchyController()
        )
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if Primitive.to_proto(resource.enable_pod_tree_labels):
            res.enable_pod_tree_labels = Primitive.to_proto(
                resource.enable_pod_tree_labels
            )
        if Primitive.to_proto(resource.enable_hierarchical_resource_quota):
            res.enable_hierarchical_resource_quota = Primitive.to_proto(
                resource.enable_hierarchical_resource_quota
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return FeatureMembershipConfigmanagementHierarchyController(
            enabled=Primitive.from_proto(resource.enabled),
            enable_pod_tree_labels=Primitive.from_proto(
                resource.enable_pod_tree_labels
            ),
            enable_hierarchical_resource_quota=Primitive.from_proto(
                resource.enable_hierarchical_resource_quota
            ),
        )


class FeatureMembershipConfigmanagementHierarchyControllerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            FeatureMembershipConfigmanagementHierarchyController.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            FeatureMembershipConfigmanagementHierarchyController.from_proto(i)
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
