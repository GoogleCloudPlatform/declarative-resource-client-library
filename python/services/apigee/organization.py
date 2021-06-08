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
from google3.cloud.graphite.mmv2.services.google.apigee import organization_pb2
from google3.cloud.graphite.mmv2.services.google.apigee import organization_pb2_grpc

from typing import List


class Organization(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        description: str = None,
        created_at: int = None,
        last_modified_at: int = None,
        expires_at: int = None,
        environments: list = None,
        properties: dict = None,
        analytics_region: str = None,
        authorized_network: str = None,
        runtime_type: str = None,
        subscription_type: str = None,
        billing_type: str = None,
        ca_certificate: str = None,
        runtime_database_encryption_key_name: str = None,
        project_id: str = None,
        state: str = None,
        parent: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.description = description
        self.properties = properties
        self.analytics_region = analytics_region
        self.authorized_network = authorized_network
        self.runtime_type = runtime_type
        self.runtime_database_encryption_key_name = runtime_database_encryption_key_name
        self.parent = parent
        self.service_account_file = service_account_file

    def apply(self):
        stub = organization_pb2_grpc.ApigeeOrganizationServiceStub(channel.Channel())
        request = organization_pb2.ApplyApigeeOrganizationRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if OrganizationProperties.to_proto(self.properties):
            request.resource.properties.CopyFrom(
                OrganizationProperties.to_proto(self.properties)
            )
        else:
            request.resource.ClearField("properties")
        if Primitive.to_proto(self.analytics_region):
            request.resource.analytics_region = Primitive.to_proto(
                self.analytics_region
            )

        if Primitive.to_proto(self.authorized_network):
            request.resource.authorized_network = Primitive.to_proto(
                self.authorized_network
            )

        if OrganizationRuntimeTypeEnum.to_proto(self.runtime_type):
            request.resource.runtime_type = OrganizationRuntimeTypeEnum.to_proto(
                self.runtime_type
            )

        if Primitive.to_proto(self.runtime_database_encryption_key_name):
            request.resource.runtime_database_encryption_key_name = Primitive.to_proto(
                self.runtime_database_encryption_key_name
            )

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        request.service_account_file = self.service_account_file

        response = stub.ApplyApigeeOrganization(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.description = Primitive.from_proto(response.description)
        self.created_at = Primitive.from_proto(response.created_at)
        self.last_modified_at = Primitive.from_proto(response.last_modified_at)
        self.expires_at = Primitive.from_proto(response.expires_at)
        self.environments = Primitive.from_proto(response.environments)
        self.properties = OrganizationProperties.from_proto(response.properties)
        self.analytics_region = Primitive.from_proto(response.analytics_region)
        self.authorized_network = Primitive.from_proto(response.authorized_network)
        self.runtime_type = OrganizationRuntimeTypeEnum.from_proto(
            response.runtime_type
        )
        self.subscription_type = OrganizationSubscriptionTypeEnum.from_proto(
            response.subscription_type
        )
        self.billing_type = OrganizationBillingTypeEnum.from_proto(
            response.billing_type
        )
        self.ca_certificate = Primitive.from_proto(response.ca_certificate)
        self.runtime_database_encryption_key_name = Primitive.from_proto(
            response.runtime_database_encryption_key_name
        )
        self.project_id = Primitive.from_proto(response.project_id)
        self.state = OrganizationStateEnum.from_proto(response.state)
        self.parent = Primitive.from_proto(response.parent)

    def delete(self):
        stub = organization_pb2_grpc.ApigeeOrganizationServiceStub(channel.Channel())
        request = organization_pb2.DeleteApigeeOrganizationRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if OrganizationProperties.to_proto(self.properties):
            request.resource.properties.CopyFrom(
                OrganizationProperties.to_proto(self.properties)
            )
        else:
            request.resource.ClearField("properties")
        if Primitive.to_proto(self.analytics_region):
            request.resource.analytics_region = Primitive.to_proto(
                self.analytics_region
            )

        if Primitive.to_proto(self.authorized_network):
            request.resource.authorized_network = Primitive.to_proto(
                self.authorized_network
            )

        if OrganizationRuntimeTypeEnum.to_proto(self.runtime_type):
            request.resource.runtime_type = OrganizationRuntimeTypeEnum.to_proto(
                self.runtime_type
            )

        if Primitive.to_proto(self.runtime_database_encryption_key_name):
            request.resource.runtime_database_encryption_key_name = Primitive.to_proto(
                self.runtime_database_encryption_key_name
            )

        if Primitive.to_proto(self.parent):
            request.resource.parent = Primitive.to_proto(self.parent)

        response = stub.DeleteApigeeOrganization(request)

    @classmethod
    def list(self, service_account_file=""):
        stub = organization_pb2_grpc.ApigeeOrganizationServiceStub(channel.Channel())
        request = organization_pb2.ListApigeeOrganizationRequest()
        request.service_account_file = service_account_file

        return stub.ListApigeeOrganization(request).items

    def to_proto(self):
        resource = organization_pb2.ApigeeOrganization()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if OrganizationProperties.to_proto(self.properties):
            resource.properties.CopyFrom(
                OrganizationProperties.to_proto(self.properties)
            )
        else:
            resource.ClearField("properties")
        if Primitive.to_proto(self.analytics_region):
            resource.analytics_region = Primitive.to_proto(self.analytics_region)
        if Primitive.to_proto(self.authorized_network):
            resource.authorized_network = Primitive.to_proto(self.authorized_network)
        if OrganizationRuntimeTypeEnum.to_proto(self.runtime_type):
            resource.runtime_type = OrganizationRuntimeTypeEnum.to_proto(
                self.runtime_type
            )
        if Primitive.to_proto(self.runtime_database_encryption_key_name):
            resource.runtime_database_encryption_key_name = Primitive.to_proto(
                self.runtime_database_encryption_key_name
            )
        if Primitive.to_proto(self.parent):
            resource.parent = Primitive.to_proto(self.parent)
        return resource


class OrganizationProperties(object):
    def __init__(self, property: list = None):
        self.property = property

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = organization_pb2.ApigeeOrganizationProperties()
        if OrganizationPropertiesPropertyArray.to_proto(resource.property):
            res.property.extend(
                OrganizationPropertiesPropertyArray.to_proto(resource.property)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OrganizationProperties(
            property=OrganizationPropertiesPropertyArray.from_proto(resource.property),
        )


class OrganizationPropertiesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OrganizationProperties.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OrganizationProperties.from_proto(i) for i in resources]


class OrganizationPropertiesProperty(object):
    def __init__(self, name: str = None, value: str = None):
        self.name = name
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = organization_pb2.ApigeeOrganizationPropertiesProperty()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return OrganizationPropertiesProperty(
            name=Primitive.from_proto(resource.name),
            value=Primitive.from_proto(resource.value),
        )


class OrganizationPropertiesPropertyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [OrganizationPropertiesProperty.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [OrganizationPropertiesProperty.from_proto(i) for i in resources]


class OrganizationRuntimeTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeOrganizationRuntimeTypeEnum.Value(
            "ApigeeOrganizationRuntimeTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeOrganizationRuntimeTypeEnum.Name(resource)[
            len("ApigeeOrganizationRuntimeTypeEnum") :
        ]


class OrganizationSubscriptionTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeOrganizationSubscriptionTypeEnum.Value(
            "ApigeeOrganizationSubscriptionTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeOrganizationSubscriptionTypeEnum.Name(resource)[
            len("ApigeeOrganizationSubscriptionTypeEnum") :
        ]


class OrganizationBillingTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeOrganizationBillingTypeEnum.Value(
            "ApigeeOrganizationBillingTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeOrganizationBillingTypeEnum.Name(resource)[
            len("ApigeeOrganizationBillingTypeEnum") :
        ]


class OrganizationStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeOrganizationStateEnum.Value(
            "ApigeeOrganizationStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return organization_pb2.ApigeeOrganizationStateEnum.Name(resource)[
            len("ApigeeOrganizationStateEnum") :
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
