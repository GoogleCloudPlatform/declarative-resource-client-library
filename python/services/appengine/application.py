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
from google3.cloud.graphite.mmv2.services.google.appengine import application_pb2
from google3.cloud.graphite.mmv2.services.google.appengine import application_pb2_grpc

from typing import List


class Application(object):
    def __init__(
        self,
        auth_domain: str = None,
        blocked_addresses: list = None,
        code_bucket: str = None,
        consumer_project: dict = None,
        database_type: str = None,
        default_bucket: str = None,
        default_hostname: str = None,
        dispatch_rules: list = None,
        domains: list = None,
        feature_settings: dict = None,
        gcr_domain: str = None,
        iap: dict = None,
        parent: dict = None,
        name: str = None,
        location: str = None,
        serving_status: str = None,
        service_account: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.auth_domain = auth_domain
        self.blocked_addresses = blocked_addresses
        self.code_bucket = code_bucket
        self.consumer_project = consumer_project
        self.database_type = database_type
        self.default_bucket = default_bucket
        self.default_hostname = default_hostname
        self.dispatch_rules = dispatch_rules
        self.domains = domains
        self.feature_settings = feature_settings
        self.gcr_domain = gcr_domain
        self.iap = iap
        self.parent = parent
        self.name = name
        self.location = location
        self.serving_status = serving_status
        self.service_account = service_account
        self.service_account_file = service_account_file

    def apply(self):
        stub = application_pb2_grpc.AppengineApplicationServiceStub(channel.Channel())
        request = application_pb2.ApplyAppengineApplicationRequest()
        if Primitive.to_proto(self.auth_domain):
            request.resource.auth_domain = Primitive.to_proto(self.auth_domain)

        if Primitive.to_proto(self.blocked_addresses):
            request.resource.blocked_addresses.extend(
                Primitive.to_proto(self.blocked_addresses)
            )
        if Primitive.to_proto(self.code_bucket):
            request.resource.code_bucket = Primitive.to_proto(self.code_bucket)

        if ApplicationConsumerProject.to_proto(self.consumer_project):
            request.resource.consumer_project.CopyFrom(
                ApplicationConsumerProject.to_proto(self.consumer_project)
            )
        else:
            request.resource.ClearField("consumer_project")
        if ApplicationDatabaseTypeEnum.to_proto(self.database_type):
            request.resource.database_type = ApplicationDatabaseTypeEnum.to_proto(
                self.database_type
            )

        if Primitive.to_proto(self.default_bucket):
            request.resource.default_bucket = Primitive.to_proto(self.default_bucket)

        if Primitive.to_proto(self.default_hostname):
            request.resource.default_hostname = Primitive.to_proto(
                self.default_hostname
            )

        if ApplicationDispatchRulesArray.to_proto(self.dispatch_rules):
            request.resource.dispatch_rules.extend(
                ApplicationDispatchRulesArray.to_proto(self.dispatch_rules)
            )
        if Primitive.to_proto(self.domains):
            request.resource.domains.extend(Primitive.to_proto(self.domains))
        if ApplicationFeatureSettings.to_proto(self.feature_settings):
            request.resource.feature_settings.CopyFrom(
                ApplicationFeatureSettings.to_proto(self.feature_settings)
            )
        else:
            request.resource.ClearField("feature_settings")
        if Primitive.to_proto(self.gcr_domain):
            request.resource.gcr_domain = Primitive.to_proto(self.gcr_domain)

        if ApplicationIap.to_proto(self.iap):
            request.resource.iap.CopyFrom(ApplicationIap.to_proto(self.iap))
        else:
            request.resource.ClearField("iap")
        if ApplicationParent.to_proto(self.parent):
            request.resource.parent.CopyFrom(ApplicationParent.to_proto(self.parent))
        else:
            request.resource.ClearField("parent")
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        if ApplicationServingStatusEnum.to_proto(self.serving_status):
            request.resource.serving_status = ApplicationServingStatusEnum.to_proto(
                self.serving_status
            )

        if Primitive.to_proto(self.service_account):
            request.resource.service_account = Primitive.to_proto(self.service_account)

        request.service_account_file = self.service_account_file

        response = stub.ApplyAppengineApplication(request)
        self.auth_domain = Primitive.from_proto(response.auth_domain)
        self.blocked_addresses = Primitive.from_proto(response.blocked_addresses)
        self.code_bucket = Primitive.from_proto(response.code_bucket)
        self.consumer_project = ApplicationConsumerProject.from_proto(
            response.consumer_project
        )
        self.database_type = ApplicationDatabaseTypeEnum.from_proto(
            response.database_type
        )
        self.default_bucket = Primitive.from_proto(response.default_bucket)
        self.default_hostname = Primitive.from_proto(response.default_hostname)
        self.dispatch_rules = ApplicationDispatchRulesArray.from_proto(
            response.dispatch_rules
        )
        self.domains = Primitive.from_proto(response.domains)
        self.feature_settings = ApplicationFeatureSettings.from_proto(
            response.feature_settings
        )
        self.gcr_domain = Primitive.from_proto(response.gcr_domain)
        self.iap = ApplicationIap.from_proto(response.iap)
        self.parent = ApplicationParent.from_proto(response.parent)
        self.name = Primitive.from_proto(response.name)
        self.location = Primitive.from_proto(response.location)
        self.serving_status = ApplicationServingStatusEnum.from_proto(
            response.serving_status
        )
        self.service_account = Primitive.from_proto(response.service_account)

    @classmethod
    def delete(self, service_account_file=""):
        stub = application_pb2_grpc.AppengineApplicationServiceStub(channel.Channel())
        request = application_pb2.DeleteAppengineApplicationRequest()
        request.service_account_file = service_account_file
        response = stub.DeleteAppengineApplication(request)

    @classmethod
    def list(self, service_account_file=""):
        stub = application_pb2_grpc.AppengineApplicationServiceStub(channel.Channel())
        request = application_pb2.ListAppengineApplicationRequest()
        request.service_account_file = service_account_file

        return stub.ListAppengineApplication(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = application_pb2.AppengineApplication()
        any_proto.Unpack(res_proto)

        res = Application()
        res.auth_domain = Primitive.from_proto(res_proto.auth_domain)
        res.blocked_addresses = Primitive.from_proto(res_proto.blocked_addresses)
        res.code_bucket = Primitive.from_proto(res_proto.code_bucket)
        res.consumer_project = ApplicationConsumerProject.from_proto(
            res_proto.consumer_project
        )
        res.database_type = ApplicationDatabaseTypeEnum.from_proto(
            res_proto.database_type
        )
        res.default_bucket = Primitive.from_proto(res_proto.default_bucket)
        res.default_hostname = Primitive.from_proto(res_proto.default_hostname)
        res.dispatch_rules = ApplicationDispatchRulesArray.from_proto(
            res_proto.dispatch_rules
        )
        res.domains = Primitive.from_proto(res_proto.domains)
        res.feature_settings = ApplicationFeatureSettings.from_proto(
            res_proto.feature_settings
        )
        res.gcr_domain = Primitive.from_proto(res_proto.gcr_domain)
        res.iap = ApplicationIap.from_proto(res_proto.iap)
        res.parent = ApplicationParent.from_proto(res_proto.parent)
        res.name = Primitive.from_proto(res_proto.name)
        res.location = Primitive.from_proto(res_proto.location)
        res.serving_status = ApplicationServingStatusEnum.from_proto(
            res_proto.serving_status
        )
        res.service_account = Primitive.from_proto(res_proto.service_account)
        return res


class ApplicationConsumerProject(object):
    def __init__(
        self,
        service: str = None,
        name: str = None,
        project_id: str = None,
        route_hash: dict = None,
        region: str = None,
    ):
        self.service = service
        self.name = name
        self.project_id = project_id
        self.route_hash = route_hash
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = application_pb2.AppengineApplicationConsumerProject()
        if ApplicationConsumerProjectServiceEnum.to_proto(resource.service):
            res.service = ApplicationConsumerProjectServiceEnum.to_proto(
                resource.service
            )
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.project_id):
            res.project_id = Primitive.to_proto(resource.project_id)
        if ApplicationConsumerProjectRouteHash.to_proto(resource.route_hash):
            res.route_hash.CopyFrom(
                ApplicationConsumerProjectRouteHash.to_proto(resource.route_hash)
            )
        else:
            res.ClearField("route_hash")
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ApplicationConsumerProject(
            service=resource.service,
            name=resource.name,
            project_id=resource.project_id,
            route_hash=resource.route_hash,
            region=resource.region,
        )


class ApplicationConsumerProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ApplicationConsumerProject.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ApplicationConsumerProject.from_proto(i) for i in resources]


class ApplicationConsumerProjectRouteHash(object):
    def __init__(self, project_hash: str = None):
        self.project_hash = project_hash

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = application_pb2.AppengineApplicationConsumerProjectRouteHash()
        if Primitive.to_proto(resource.project_hash):
            res.project_hash = Primitive.to_proto(resource.project_hash)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ApplicationConsumerProjectRouteHash(project_hash=resource.project_hash,)


class ApplicationConsumerProjectRouteHashArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ApplicationConsumerProjectRouteHash.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ApplicationConsumerProjectRouteHash.from_proto(i) for i in resources]


class ApplicationDispatchRules(object):
    def __init__(self, domain: str = None, path: str = None, service: str = None):
        self.domain = domain
        self.path = path
        self.service = service

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = application_pb2.AppengineApplicationDispatchRules()
        if Primitive.to_proto(resource.domain):
            res.domain = Primitive.to_proto(resource.domain)
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.service):
            res.service = Primitive.to_proto(resource.service)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ApplicationDispatchRules(
            domain=resource.domain, path=resource.path, service=resource.service,
        )


class ApplicationDispatchRulesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ApplicationDispatchRules.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ApplicationDispatchRules.from_proto(i) for i in resources]


class ApplicationFeatureSettings(object):
    def __init__(
        self, split_health_checks: bool = None, use_container_optimized_os: bool = None
    ):
        self.split_health_checks = split_health_checks
        self.use_container_optimized_os = use_container_optimized_os

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = application_pb2.AppengineApplicationFeatureSettings()
        if Primitive.to_proto(resource.split_health_checks):
            res.split_health_checks = Primitive.to_proto(resource.split_health_checks)
        if Primitive.to_proto(resource.use_container_optimized_os):
            res.use_container_optimized_os = Primitive.to_proto(
                resource.use_container_optimized_os
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ApplicationFeatureSettings(
            split_health_checks=resource.split_health_checks,
            use_container_optimized_os=resource.use_container_optimized_os,
        )


class ApplicationFeatureSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ApplicationFeatureSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ApplicationFeatureSettings.from_proto(i) for i in resources]


class ApplicationIap(object):
    def __init__(
        self,
        enabled: bool = None,
        oauth2_client_id: str = None,
        oauth2_client_secret: str = None,
        oauth2_client_secret_sha256: str = None,
    ):
        self.enabled = enabled
        self.oauth2_client_id = oauth2_client_id
        self.oauth2_client_secret = oauth2_client_secret
        self.oauth2_client_secret_sha256 = oauth2_client_secret_sha256

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = application_pb2.AppengineApplicationIap()
        if Primitive.to_proto(resource.enabled):
            res.enabled = Primitive.to_proto(resource.enabled)
        if Primitive.to_proto(resource.oauth2_client_id):
            res.oauth2_client_id = Primitive.to_proto(resource.oauth2_client_id)
        if Primitive.to_proto(resource.oauth2_client_secret):
            res.oauth2_client_secret = Primitive.to_proto(resource.oauth2_client_secret)
        if Primitive.to_proto(resource.oauth2_client_secret_sha256):
            res.oauth2_client_secret_sha256 = Primitive.to_proto(
                resource.oauth2_client_secret_sha256
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ApplicationIap(
            enabled=resource.enabled,
            oauth2_client_id=resource.oauth2_client_id,
            oauth2_client_secret=resource.oauth2_client_secret,
            oauth2_client_secret_sha256=resource.oauth2_client_secret_sha256,
        )


class ApplicationIapArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ApplicationIap.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ApplicationIap.from_proto(i) for i in resources]


class ApplicationParent(object):
    def __init__(self, type: str = None, id: str = None):
        self.type = type
        self.id = id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = application_pb2.AppengineApplicationParent()
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ApplicationParent(type=resource.type, id=resource.id,)


class ApplicationParentArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [ApplicationParent.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [ApplicationParent.from_proto(i) for i in resources]


class ApplicationConsumerProjectServiceEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return application_pb2.AppengineApplicationConsumerProjectServiceEnum.Value(
            "AppengineApplicationConsumerProjectServiceEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return application_pb2.AppengineApplicationConsumerProjectServiceEnum.Name(
            resource
        )[len("AppengineApplicationConsumerProjectServiceEnum") :]


class ApplicationDatabaseTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return application_pb2.AppengineApplicationDatabaseTypeEnum.Value(
            "AppengineApplicationDatabaseTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return application_pb2.AppengineApplicationDatabaseTypeEnum.Name(resource)[
            len("AppengineApplicationDatabaseTypeEnum") :
        ]


class ApplicationServingStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return application_pb2.AppengineApplicationServingStatusEnum.Value(
            "AppengineApplicationServingStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return application_pb2.AppengineApplicationServingStatusEnum.Name(resource)[
            len("AppengineApplicationServingStatusEnum") :
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
