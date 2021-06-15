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
from google3.cloud.graphite.mmv2.services.google.privateca import (
    certificate_authority_pb2,
)
from google3.cloud.graphite.mmv2.services.google.privateca import (
    certificate_authority_pb2_grpc,
)

from typing import List


class CertificateAuthority(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        config: dict = None,
        lifetime: str = None,
        key_spec: dict = None,
        subordinate_config: dict = None,
        tier: str = None,
        state: str = None,
        pem_ca_certificates: list = None,
        ca_certificate_descriptions: list = None,
        gcs_bucket: str = None,
        access_urls: dict = None,
        create_time: str = None,
        update_time: str = None,
        delete_time: str = None,
        labels: dict = None,
        certificate_policy: dict = None,
        issuing_options: dict = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.type = type
        self.config = config
        self.lifetime = lifetime
        self.key_spec = key_spec
        self.subordinate_config = subordinate_config
        self.tier = tier
        self.gcs_bucket = gcs_bucket
        self.labels = labels
        self.certificate_policy = certificate_policy
        self.issuing_options = issuing_options
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = certificate_authority_pb2_grpc.PrivatecaBetaCertificateAuthorityServiceStub(
            channel.Channel()
        )
        request = (
            certificate_authority_pb2.ApplyPrivatecaBetaCertificateAuthorityRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if CertificateAuthorityTypeEnum.to_proto(self.type):
            request.resource.type = CertificateAuthorityTypeEnum.to_proto(self.type)

        if CertificateAuthorityConfig.to_proto(self.config):
            request.resource.config.CopyFrom(
                CertificateAuthorityConfig.to_proto(self.config)
            )
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.lifetime):
            request.resource.lifetime = Primitive.to_proto(self.lifetime)

        if CertificateAuthorityKeySpec.to_proto(self.key_spec):
            request.resource.key_spec.CopyFrom(
                CertificateAuthorityKeySpec.to_proto(self.key_spec)
            )
        else:
            request.resource.ClearField("key_spec")
        if CertificateAuthoritySubordinateConfig.to_proto(self.subordinate_config):
            request.resource.subordinate_config.CopyFrom(
                CertificateAuthoritySubordinateConfig.to_proto(self.subordinate_config)
            )
        else:
            request.resource.ClearField("subordinate_config")
        if CertificateAuthorityTierEnum.to_proto(self.tier):
            request.resource.tier = CertificateAuthorityTierEnum.to_proto(self.tier)

        if Primitive.to_proto(self.gcs_bucket):
            request.resource.gcs_bucket = Primitive.to_proto(self.gcs_bucket)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if CertificateAuthorityCertificatePolicy.to_proto(self.certificate_policy):
            request.resource.certificate_policy.CopyFrom(
                CertificateAuthorityCertificatePolicy.to_proto(self.certificate_policy)
            )
        else:
            request.resource.ClearField("certificate_policy")
        if CertificateAuthorityIssuingOptions.to_proto(self.issuing_options):
            request.resource.issuing_options.CopyFrom(
                CertificateAuthorityIssuingOptions.to_proto(self.issuing_options)
            )
        else:
            request.resource.ClearField("issuing_options")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyPrivatecaBetaCertificateAuthority(request)
        self.name = Primitive.from_proto(response.name)
        self.type = CertificateAuthorityTypeEnum.from_proto(response.type)
        self.config = CertificateAuthorityConfig.from_proto(response.config)
        self.lifetime = Primitive.from_proto(response.lifetime)
        self.key_spec = CertificateAuthorityKeySpec.from_proto(response.key_spec)
        self.subordinate_config = CertificateAuthoritySubordinateConfig.from_proto(
            response.subordinate_config
        )
        self.tier = CertificateAuthorityTierEnum.from_proto(response.tier)
        self.state = CertificateAuthorityStateEnum.from_proto(response.state)
        self.pem_ca_certificates = Primitive.from_proto(response.pem_ca_certificates)
        self.ca_certificate_descriptions = CertificateAuthorityCaCertificateDescriptionsArray.from_proto(
            response.ca_certificate_descriptions
        )
        self.gcs_bucket = Primitive.from_proto(response.gcs_bucket)
        self.access_urls = CertificateAuthorityAccessUrls.from_proto(
            response.access_urls
        )
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.delete_time = Primitive.from_proto(response.delete_time)
        self.labels = Primitive.from_proto(response.labels)
        self.certificate_policy = CertificateAuthorityCertificatePolicy.from_proto(
            response.certificate_policy
        )
        self.issuing_options = CertificateAuthorityIssuingOptions.from_proto(
            response.issuing_options
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = certificate_authority_pb2_grpc.PrivatecaBetaCertificateAuthorityServiceStub(
            channel.Channel()
        )
        request = (
            certificate_authority_pb2.DeletePrivatecaBetaCertificateAuthorityRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if CertificateAuthorityTypeEnum.to_proto(self.type):
            request.resource.type = CertificateAuthorityTypeEnum.to_proto(self.type)

        if CertificateAuthorityConfig.to_proto(self.config):
            request.resource.config.CopyFrom(
                CertificateAuthorityConfig.to_proto(self.config)
            )
        else:
            request.resource.ClearField("config")
        if Primitive.to_proto(self.lifetime):
            request.resource.lifetime = Primitive.to_proto(self.lifetime)

        if CertificateAuthorityKeySpec.to_proto(self.key_spec):
            request.resource.key_spec.CopyFrom(
                CertificateAuthorityKeySpec.to_proto(self.key_spec)
            )
        else:
            request.resource.ClearField("key_spec")
        if CertificateAuthoritySubordinateConfig.to_proto(self.subordinate_config):
            request.resource.subordinate_config.CopyFrom(
                CertificateAuthoritySubordinateConfig.to_proto(self.subordinate_config)
            )
        else:
            request.resource.ClearField("subordinate_config")
        if CertificateAuthorityTierEnum.to_proto(self.tier):
            request.resource.tier = CertificateAuthorityTierEnum.to_proto(self.tier)

        if Primitive.to_proto(self.gcs_bucket):
            request.resource.gcs_bucket = Primitive.to_proto(self.gcs_bucket)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if CertificateAuthorityCertificatePolicy.to_proto(self.certificate_policy):
            request.resource.certificate_policy.CopyFrom(
                CertificateAuthorityCertificatePolicy.to_proto(self.certificate_policy)
            )
        else:
            request.resource.ClearField("certificate_policy")
        if CertificateAuthorityIssuingOptions.to_proto(self.issuing_options):
            request.resource.issuing_options.CopyFrom(
                CertificateAuthorityIssuingOptions.to_proto(self.issuing_options)
            )
        else:
            request.resource.ClearField("issuing_options")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeletePrivatecaBetaCertificateAuthority(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = certificate_authority_pb2_grpc.PrivatecaBetaCertificateAuthorityServiceStub(
            channel.Channel()
        )
        request = (
            certificate_authority_pb2.ListPrivatecaBetaCertificateAuthorityRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListPrivatecaBetaCertificateAuthority(request).items

    def to_proto(self):
        resource = certificate_authority_pb2.PrivatecaBetaCertificateAuthority()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if CertificateAuthorityTypeEnum.to_proto(self.type):
            resource.type = CertificateAuthorityTypeEnum.to_proto(self.type)
        if CertificateAuthorityConfig.to_proto(self.config):
            resource.config.CopyFrom(CertificateAuthorityConfig.to_proto(self.config))
        else:
            resource.ClearField("config")
        if Primitive.to_proto(self.lifetime):
            resource.lifetime = Primitive.to_proto(self.lifetime)
        if CertificateAuthorityKeySpec.to_proto(self.key_spec):
            resource.key_spec.CopyFrom(
                CertificateAuthorityKeySpec.to_proto(self.key_spec)
            )
        else:
            resource.ClearField("key_spec")
        if CertificateAuthoritySubordinateConfig.to_proto(self.subordinate_config):
            resource.subordinate_config.CopyFrom(
                CertificateAuthoritySubordinateConfig.to_proto(self.subordinate_config)
            )
        else:
            resource.ClearField("subordinate_config")
        if CertificateAuthorityTierEnum.to_proto(self.tier):
            resource.tier = CertificateAuthorityTierEnum.to_proto(self.tier)
        if Primitive.to_proto(self.gcs_bucket):
            resource.gcs_bucket = Primitive.to_proto(self.gcs_bucket)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if CertificateAuthorityCertificatePolicy.to_proto(self.certificate_policy):
            resource.certificate_policy.CopyFrom(
                CertificateAuthorityCertificatePolicy.to_proto(self.certificate_policy)
            )
        else:
            resource.ClearField("certificate_policy")
        if CertificateAuthorityIssuingOptions.to_proto(self.issuing_options):
            resource.issuing_options.CopyFrom(
                CertificateAuthorityIssuingOptions.to_proto(self.issuing_options)
            )
        else:
            resource.ClearField("issuing_options")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class CertificateAuthorityConfig(object):
    def __init__(
        self,
        subject_config: dict = None,
        public_key: dict = None,
        reusable_config: dict = None,
    ):
        self.subject_config = subject_config
        self.public_key = public_key
        self.reusable_config = reusable_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfig()
        if CertificateAuthorityConfigSubjectConfig.to_proto(resource.subject_config):
            res.subject_config.CopyFrom(
                CertificateAuthorityConfigSubjectConfig.to_proto(
                    resource.subject_config
                )
            )
        else:
            res.ClearField("subject_config")
        if CertificateAuthorityConfigPublicKey.to_proto(resource.public_key):
            res.public_key.CopyFrom(
                CertificateAuthorityConfigPublicKey.to_proto(resource.public_key)
            )
        else:
            res.ClearField("public_key")
        if CertificateAuthorityConfigReusableConfig.to_proto(resource.reusable_config):
            res.reusable_config.CopyFrom(
                CertificateAuthorityConfigReusableConfig.to_proto(
                    resource.reusable_config
                )
            )
        else:
            res.ClearField("reusable_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfig(
            subject_config=CertificateAuthorityConfigSubjectConfig.from_proto(
                resource.subject_config
            ),
            public_key=CertificateAuthorityConfigPublicKey.from_proto(
                resource.public_key
            ),
            reusable_config=CertificateAuthorityConfigReusableConfig.from_proto(
                resource.reusable_config
            ),
        )


class CertificateAuthorityConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthorityConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateAuthorityConfig.from_proto(i) for i in resources]


class CertificateAuthorityConfigSubjectConfig(object):
    def __init__(
        self,
        subject: dict = None,
        common_name: str = None,
        subject_alt_name: dict = None,
    ):
        self.subject = subject
        self.common_name = common_name
        self.subject_alt_name = subject_alt_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigSubjectConfig()
        )
        if CertificateAuthorityConfigSubjectConfigSubject.to_proto(resource.subject):
            res.subject.CopyFrom(
                CertificateAuthorityConfigSubjectConfigSubject.to_proto(
                    resource.subject
                )
            )
        else:
            res.ClearField("subject")
        if Primitive.to_proto(resource.common_name):
            res.common_name = Primitive.to_proto(resource.common_name)
        if CertificateAuthorityConfigSubjectConfigSubjectAltName.to_proto(
            resource.subject_alt_name
        ):
            res.subject_alt_name.CopyFrom(
                CertificateAuthorityConfigSubjectConfigSubjectAltName.to_proto(
                    resource.subject_alt_name
                )
            )
        else:
            res.ClearField("subject_alt_name")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigSubjectConfig(
            subject=CertificateAuthorityConfigSubjectConfigSubject.from_proto(
                resource.subject
            ),
            common_name=Primitive.from_proto(resource.common_name),
            subject_alt_name=CertificateAuthorityConfigSubjectConfigSubjectAltName.from_proto(
                resource.subject_alt_name
            ),
        )


class CertificateAuthorityConfigSubjectConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthorityConfigSubjectConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigSubjectConfig.from_proto(i) for i in resources
        ]


class CertificateAuthorityConfigSubjectConfigSubject(object):
    def __init__(
        self,
        country_code: str = None,
        organization: str = None,
        organizational_unit: str = None,
        locality: str = None,
        province: str = None,
        street_address: str = None,
        postal_code: str = None,
    ):
        self.country_code = country_code
        self.organization = organization
        self.organizational_unit = organizational_unit
        self.locality = locality
        self.province = province
        self.street_address = street_address
        self.postal_code = postal_code

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubject()
        )
        if Primitive.to_proto(resource.country_code):
            res.country_code = Primitive.to_proto(resource.country_code)
        if Primitive.to_proto(resource.organization):
            res.organization = Primitive.to_proto(resource.organization)
        if Primitive.to_proto(resource.organizational_unit):
            res.organizational_unit = Primitive.to_proto(resource.organizational_unit)
        if Primitive.to_proto(resource.locality):
            res.locality = Primitive.to_proto(resource.locality)
        if Primitive.to_proto(resource.province):
            res.province = Primitive.to_proto(resource.province)
        if Primitive.to_proto(resource.street_address):
            res.street_address = Primitive.to_proto(resource.street_address)
        if Primitive.to_proto(resource.postal_code):
            res.postal_code = Primitive.to_proto(resource.postal_code)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigSubjectConfigSubject(
            country_code=Primitive.from_proto(resource.country_code),
            organization=Primitive.from_proto(resource.organization),
            organizational_unit=Primitive.from_proto(resource.organizational_unit),
            locality=Primitive.from_proto(resource.locality),
            province=Primitive.from_proto(resource.province),
            street_address=Primitive.from_proto(resource.street_address),
            postal_code=Primitive.from_proto(resource.postal_code),
        )


class CertificateAuthorityConfigSubjectConfigSubjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigSubjectConfigSubject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigSubjectConfigSubject.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityConfigSubjectConfigSubjectAltName(object):
    def __init__(
        self,
        dns_names: list = None,
        uris: list = None,
        email_addresses: list = None,
        ip_addresses: list = None,
        custom_sans: list = None,
    ):
        self.dns_names = dns_names
        self.uris = uris
        self.email_addresses = email_addresses
        self.ip_addresses = ip_addresses
        self.custom_sans = custom_sans

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltName()
        )
        if Primitive.to_proto(resource.dns_names):
            res.dns_names.extend(Primitive.to_proto(resource.dns_names))
        if Primitive.to_proto(resource.uris):
            res.uris.extend(Primitive.to_proto(resource.uris))
        if Primitive.to_proto(resource.email_addresses):
            res.email_addresses.extend(Primitive.to_proto(resource.email_addresses))
        if Primitive.to_proto(resource.ip_addresses):
            res.ip_addresses.extend(Primitive.to_proto(resource.ip_addresses))
        if CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansArray.to_proto(
            resource.custom_sans
        ):
            res.custom_sans.extend(
                CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansArray.to_proto(
                    resource.custom_sans
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigSubjectConfigSubjectAltName(
            dns_names=Primitive.from_proto(resource.dns_names),
            uris=Primitive.from_proto(resource.uris),
            email_addresses=Primitive.from_proto(resource.email_addresses),
            ip_addresses=Primitive.from_proto(resource.ip_addresses),
            custom_sans=CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansArray.from_proto(
                resource.custom_sans
            ),
        )


class CertificateAuthorityConfigSubjectConfigSubjectAltNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigSubjectConfigSubjectAltName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigSubjectConfigSubjectAltName.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(object):
    def __init__(
        self, object_id: dict = None, critical: bool = None, value: str = None
    ):
        self.object_id = object_id
        self.critical = critical
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans()
        )
        if CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId.to_proto(
                    resource.object_id
                )
            )
        else:
            res.ClearField("object_id")
        if Primitive.to_proto(resource.critical):
            res.critical = Primitive.to_proto(resource.critical)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans(
            object_id=CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectIdArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityConfigPublicKey(object):
    def __init__(self, key: str = None, type: str = None):
        self.key = key
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigPublicKey()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if CertificateAuthorityConfigPublicKeyTypeEnum.to_proto(resource.type):
            res.type = CertificateAuthorityConfigPublicKeyTypeEnum.to_proto(
                resource.type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigPublicKey(
            key=Primitive.from_proto(resource.key),
            type=CertificateAuthorityConfigPublicKeyTypeEnum.from_proto(resource.type),
        )


class CertificateAuthorityConfigPublicKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthorityConfigPublicKey.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateAuthorityConfigPublicKey.from_proto(i) for i in resources]


class CertificateAuthorityConfigReusableConfig(object):
    def __init__(
        self, reusable_config: str = None, reusable_config_values: dict = None
    ):
        self.reusable_config = reusable_config
        self.reusable_config_values = reusable_config_values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigReusableConfig()
        )
        if Primitive.to_proto(resource.reusable_config):
            res.reusable_config = Primitive.to_proto(resource.reusable_config)
        if CertificateAuthorityConfigReusableConfigReusableConfigValues.to_proto(
            resource.reusable_config_values
        ):
            res.reusable_config_values.CopyFrom(
                CertificateAuthorityConfigReusableConfigReusableConfigValues.to_proto(
                    resource.reusable_config_values
                )
            )
        else:
            res.ClearField("reusable_config_values")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigReusableConfig(
            reusable_config=Primitive.from_proto(resource.reusable_config),
            reusable_config_values=CertificateAuthorityConfigReusableConfigReusableConfigValues.from_proto(
                resource.reusable_config_values
            ),
        )


class CertificateAuthorityConfigReusableConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthorityConfigReusableConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigReusableConfig.from_proto(i) for i in resources
        ]


class CertificateAuthorityConfigReusableConfigReusableConfigValues(object):
    def __init__(
        self,
        key_usage: dict = None,
        ca_options: dict = None,
        policy_ids: list = None,
        aia_ocsp_servers: list = None,
        additional_extensions: list = None,
    ):
        self.key_usage = key_usage
        self.ca_options = ca_options
        self.policy_ids = policy_ids
        self.aia_ocsp_servers = aia_ocsp_servers
        self.additional_extensions = additional_extensions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValues()
        )
        if CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage.to_proto(
            resource.key_usage
        ):
            res.key_usage.CopyFrom(
                CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage.to_proto(
                    resource.key_usage
                )
            )
        else:
            res.ClearField("key_usage")
        if CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions.to_proto(
            resource.ca_options
        ):
            res.ca_options.CopyFrom(
                CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions.to_proto(
                    resource.ca_options
                )
            )
        else:
            res.ClearField("ca_options")
        if CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsArray.to_proto(
            resource.policy_ids
        ):
            res.policy_ids.extend(
                CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsArray.to_proto(
                    resource.policy_ids
                )
            )
        if Primitive.to_proto(resource.aia_ocsp_servers):
            res.aia_ocsp_servers.extend(Primitive.to_proto(resource.aia_ocsp_servers))
        if CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsArray.to_proto(
            resource.additional_extensions
        ):
            res.additional_extensions.extend(
                CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsArray.to_proto(
                    resource.additional_extensions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigReusableConfigReusableConfigValues(
            key_usage=CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage.from_proto(
                resource.key_usage
            ),
            ca_options=CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions.from_proto(
                resource.ca_options
            ),
            policy_ids=CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsArray.from_proto(
                resource.policy_ids
            ),
            aia_ocsp_servers=Primitive.from_proto(resource.aia_ocsp_servers),
            additional_extensions=CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsArray.from_proto(
                resource.additional_extensions
            ),
        )


class CertificateAuthorityConfigReusableConfigReusableConfigValuesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValues.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValues.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(object):
    def __init__(
        self,
        base_key_usage: dict = None,
        extended_key_usage: dict = None,
        unknown_extended_key_usages: list = None,
    ):
        self.base_key_usage = base_key_usage
        self.extended_key_usage = extended_key_usage
        self.unknown_extended_key_usages = unknown_extended_key_usages

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage()
        )
        if CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage.to_proto(
            resource.base_key_usage
        ):
            res.base_key_usage.CopyFrom(
                CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage.to_proto(
                    resource.base_key_usage
                )
            )
        else:
            res.ClearField("base_key_usage")
        if CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage.to_proto(
            resource.extended_key_usage
        ):
            res.extended_key_usage.CopyFrom(
                CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage.to_proto(
                    resource.extended_key_usage
                )
            )
        else:
            res.ClearField("extended_key_usage")
        if CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
            resource.unknown_extended_key_usages
        ):
            res.unknown_extended_key_usages.extend(
                CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
                    resource.unknown_extended_key_usages
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage(
            base_key_usage=CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage.from_proto(
                resource.base_key_usage
            ),
            extended_key_usage=CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage.from_proto(
                resource.extended_key_usage
            ),
            unknown_extended_key_usages=CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesArray.from_proto(
                resource.unknown_extended_key_usages
            ),
        )


class CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(
    object
):
    def __init__(
        self,
        digital_signature: bool = None,
        content_commitment: bool = None,
        key_encipherment: bool = None,
        data_encipherment: bool = None,
        key_agreement: bool = None,
        cert_sign: bool = None,
        crl_sign: bool = None,
        encipher_only: bool = None,
        decipher_only: bool = None,
    ):
        self.digital_signature = digital_signature
        self.content_commitment = content_commitment
        self.key_encipherment = key_encipherment
        self.data_encipherment = data_encipherment
        self.key_agreement = key_agreement
        self.cert_sign = cert_sign
        self.crl_sign = crl_sign
        self.encipher_only = encipher_only
        self.decipher_only = decipher_only

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage()
        )
        if Primitive.to_proto(resource.digital_signature):
            res.digital_signature = Primitive.to_proto(resource.digital_signature)
        if Primitive.to_proto(resource.content_commitment):
            res.content_commitment = Primitive.to_proto(resource.content_commitment)
        if Primitive.to_proto(resource.key_encipherment):
            res.key_encipherment = Primitive.to_proto(resource.key_encipherment)
        if Primitive.to_proto(resource.data_encipherment):
            res.data_encipherment = Primitive.to_proto(resource.data_encipherment)
        if Primitive.to_proto(resource.key_agreement):
            res.key_agreement = Primitive.to_proto(resource.key_agreement)
        if Primitive.to_proto(resource.cert_sign):
            res.cert_sign = Primitive.to_proto(resource.cert_sign)
        if Primitive.to_proto(resource.crl_sign):
            res.crl_sign = Primitive.to_proto(resource.crl_sign)
        if Primitive.to_proto(resource.encipher_only):
            res.encipher_only = Primitive.to_proto(resource.encipher_only)
        if Primitive.to_proto(resource.decipher_only):
            res.decipher_only = Primitive.to_proto(resource.decipher_only)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage(
            digital_signature=Primitive.from_proto(resource.digital_signature),
            content_commitment=Primitive.from_proto(resource.content_commitment),
            key_encipherment=Primitive.from_proto(resource.key_encipherment),
            data_encipherment=Primitive.from_proto(resource.data_encipherment),
            key_agreement=Primitive.from_proto(resource.key_agreement),
            cert_sign=Primitive.from_proto(resource.cert_sign),
            crl_sign=Primitive.from_proto(resource.crl_sign),
            encipher_only=Primitive.from_proto(resource.encipher_only),
            decipher_only=Primitive.from_proto(resource.decipher_only),
        )


class CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsageArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(
    object
):
    def __init__(
        self,
        server_auth: bool = None,
        client_auth: bool = None,
        code_signing: bool = None,
        email_protection: bool = None,
        time_stamping: bool = None,
        ocsp_signing: bool = None,
    ):
        self.server_auth = server_auth
        self.client_auth = client_auth
        self.code_signing = code_signing
        self.email_protection = email_protection
        self.time_stamping = time_stamping
        self.ocsp_signing = ocsp_signing

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage()
        )
        if Primitive.to_proto(resource.server_auth):
            res.server_auth = Primitive.to_proto(resource.server_auth)
        if Primitive.to_proto(resource.client_auth):
            res.client_auth = Primitive.to_proto(resource.client_auth)
        if Primitive.to_proto(resource.code_signing):
            res.code_signing = Primitive.to_proto(resource.code_signing)
        if Primitive.to_proto(resource.email_protection):
            res.email_protection = Primitive.to_proto(resource.email_protection)
        if Primitive.to_proto(resource.time_stamping):
            res.time_stamping = Primitive.to_proto(resource.time_stamping)
        if Primitive.to_proto(resource.ocsp_signing):
            res.ocsp_signing = Primitive.to_proto(resource.ocsp_signing)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage(
            server_auth=Primitive.from_proto(resource.server_auth),
            client_auth=Primitive.from_proto(resource.client_auth),
            code_signing=Primitive.from_proto(resource.code_signing),
            email_protection=Primitive.from_proto(resource.email_protection),
            time_stamping=Primitive.from_proto(resource.time_stamping),
            ocsp_signing=Primitive.from_proto(resource.ocsp_signing),
        )


class CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsageArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(object):
    def __init__(self, is_ca: bool = None, max_issuer_path_length: int = None):
        self.is_ca = is_ca
        self.max_issuer_path_length = max_issuer_path_length

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions()
        )
        if Primitive.to_proto(resource.is_ca):
            res.is_ca = Primitive.to_proto(resource.is_ca)
        if Primitive.to_proto(resource.max_issuer_path_length):
            res.max_issuer_path_length = Primitive.to_proto(
                resource.max_issuer_path_length
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions(
            is_ca=Primitive.from_proto(resource.is_ca),
            max_issuer_path_length=Primitive.from_proto(
                resource.max_issuer_path_length
            ),
        )


class CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIdsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions(
    object
):
    def __init__(
        self, object_id: dict = None, critical: bool = None, value: str = None
    ):
        self.object_id = object_id
        self.critical = critical
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions()
        )
        if CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId.to_proto(
                    resource.object_id
                )
            )
        else:
            res.ClearField("object_id")
        if Primitive.to_proto(resource.critical):
            res.critical = Primitive.to_proto(resource.critical)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions(
            object_id=CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectIdArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityKeySpec(object):
    def __init__(self, cloud_kms_key_version: str = None, algorithm: str = None):
        self.cloud_kms_key_version = cloud_kms_key_version
        self.algorithm = algorithm

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_authority_pb2.PrivatecaBetaCertificateAuthorityKeySpec()
        if Primitive.to_proto(resource.cloud_kms_key_version):
            res.cloud_kms_key_version = Primitive.to_proto(
                resource.cloud_kms_key_version
            )
        if CertificateAuthorityKeySpecAlgorithmEnum.to_proto(resource.algorithm):
            res.algorithm = CertificateAuthorityKeySpecAlgorithmEnum.to_proto(
                resource.algorithm
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityKeySpec(
            cloud_kms_key_version=Primitive.from_proto(resource.cloud_kms_key_version),
            algorithm=CertificateAuthorityKeySpecAlgorithmEnum.from_proto(
                resource.algorithm
            ),
        )


class CertificateAuthorityKeySpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthorityKeySpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateAuthorityKeySpec.from_proto(i) for i in resources]


class CertificateAuthoritySubordinateConfig(object):
    def __init__(
        self, certificate_authority: str = None, pem_issuer_chain: dict = None
    ):
        self.certificate_authority = certificate_authority
        self.pem_issuer_chain = pem_issuer_chain

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthoritySubordinateConfig()
        )
        if Primitive.to_proto(resource.certificate_authority):
            res.certificate_authority = Primitive.to_proto(
                resource.certificate_authority
            )
        if CertificateAuthoritySubordinateConfigPemIssuerChain.to_proto(
            resource.pem_issuer_chain
        ):
            res.pem_issuer_chain.CopyFrom(
                CertificateAuthoritySubordinateConfigPemIssuerChain.to_proto(
                    resource.pem_issuer_chain
                )
            )
        else:
            res.ClearField("pem_issuer_chain")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthoritySubordinateConfig(
            certificate_authority=Primitive.from_proto(resource.certificate_authority),
            pem_issuer_chain=CertificateAuthoritySubordinateConfigPemIssuerChain.from_proto(
                resource.pem_issuer_chain
            ),
        )


class CertificateAuthoritySubordinateConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthoritySubordinateConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateAuthoritySubordinateConfig.from_proto(i) for i in resources]


class CertificateAuthoritySubordinateConfigPemIssuerChain(object):
    def __init__(self, pem_certificates: list = None):
        self.pem_certificates = pem_certificates

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthoritySubordinateConfigPemIssuerChain()
        )
        if Primitive.to_proto(resource.pem_certificates):
            res.pem_certificates.extend(Primitive.to_proto(resource.pem_certificates))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthoritySubordinateConfigPemIssuerChain(
            pem_certificates=Primitive.from_proto(resource.pem_certificates),
        )


class CertificateAuthoritySubordinateConfigPemIssuerChainArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthoritySubordinateConfigPemIssuerChain.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthoritySubordinateConfigPemIssuerChain.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptions(object):
    def __init__(
        self,
        subject_description: dict = None,
        public_key: dict = None,
        subject_key_id: dict = None,
        authority_key_id: dict = None,
        crl_distribution_points: list = None,
        aia_issuing_certificate_urls: list = None,
        cert_fingerprint: dict = None,
        config_values: dict = None,
    ):
        self.subject_description = subject_description
        self.public_key = public_key
        self.subject_key_id = subject_key_id
        self.authority_key_id = authority_key_id
        self.crl_distribution_points = crl_distribution_points
        self.aia_issuing_certificate_urls = aia_issuing_certificate_urls
        self.cert_fingerprint = cert_fingerprint
        self.config_values = config_values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptions()
        )
        if CertificateAuthorityCaCertificateDescriptionsSubjectDescription.to_proto(
            resource.subject_description
        ):
            res.subject_description.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsSubjectDescription.to_proto(
                    resource.subject_description
                )
            )
        else:
            res.ClearField("subject_description")
        if CertificateAuthorityCaCertificateDescriptionsPublicKey.to_proto(
            resource.public_key
        ):
            res.public_key.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsPublicKey.to_proto(
                    resource.public_key
                )
            )
        else:
            res.ClearField("public_key")
        if CertificateAuthorityCaCertificateDescriptionsSubjectKeyId.to_proto(
            resource.subject_key_id
        ):
            res.subject_key_id.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsSubjectKeyId.to_proto(
                    resource.subject_key_id
                )
            )
        else:
            res.ClearField("subject_key_id")
        if CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId.to_proto(
            resource.authority_key_id
        ):
            res.authority_key_id.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId.to_proto(
                    resource.authority_key_id
                )
            )
        else:
            res.ClearField("authority_key_id")
        if Primitive.to_proto(resource.crl_distribution_points):
            res.crl_distribution_points.extend(
                Primitive.to_proto(resource.crl_distribution_points)
            )
        if Primitive.to_proto(resource.aia_issuing_certificate_urls):
            res.aia_issuing_certificate_urls.extend(
                Primitive.to_proto(resource.aia_issuing_certificate_urls)
            )
        if CertificateAuthorityCaCertificateDescriptionsCertFingerprint.to_proto(
            resource.cert_fingerprint
        ):
            res.cert_fingerprint.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsCertFingerprint.to_proto(
                    resource.cert_fingerprint
                )
            )
        else:
            res.ClearField("cert_fingerprint")
        if CertificateAuthorityCaCertificateDescriptionsConfigValues.to_proto(
            resource.config_values
        ):
            res.config_values.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsConfigValues.to_proto(
                    resource.config_values
                )
            )
        else:
            res.ClearField("config_values")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptions(
            subject_description=CertificateAuthorityCaCertificateDescriptionsSubjectDescription.from_proto(
                resource.subject_description
            ),
            public_key=CertificateAuthorityCaCertificateDescriptionsPublicKey.from_proto(
                resource.public_key
            ),
            subject_key_id=CertificateAuthorityCaCertificateDescriptionsSubjectKeyId.from_proto(
                resource.subject_key_id
            ),
            authority_key_id=CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId.from_proto(
                resource.authority_key_id
            ),
            crl_distribution_points=Primitive.from_proto(
                resource.crl_distribution_points
            ),
            aia_issuing_certificate_urls=Primitive.from_proto(
                resource.aia_issuing_certificate_urls
            ),
            cert_fingerprint=CertificateAuthorityCaCertificateDescriptionsCertFingerprint.from_proto(
                resource.cert_fingerprint
            ),
            config_values=CertificateAuthorityCaCertificateDescriptionsConfigValues.from_proto(
                resource.config_values
            ),
        )


class CertificateAuthorityCaCertificateDescriptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptions.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptions.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsSubjectDescription(object):
    def __init__(
        self,
        subject: dict = None,
        subject_alt_name: dict = None,
        hex_serial_number: str = None,
        lifetime: str = None,
        not_before_time: str = None,
        not_after_time: str = None,
        common_name: str = None,
    ):
        self.subject = subject
        self.subject_alt_name = subject_alt_name
        self.hex_serial_number = hex_serial_number
        self.lifetime = lifetime
        self.not_before_time = not_before_time
        self.not_after_time = not_after_time
        self.common_name = common_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescription()
        )
        if CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject.to_proto(
            resource.subject
        ):
            res.subject.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject.to_proto(
                    resource.subject
                )
            )
        else:
            res.ClearField("subject")
        if CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName.to_proto(
            resource.subject_alt_name
        ):
            res.subject_alt_name.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName.to_proto(
                    resource.subject_alt_name
                )
            )
        else:
            res.ClearField("subject_alt_name")
        if Primitive.to_proto(resource.hex_serial_number):
            res.hex_serial_number = Primitive.to_proto(resource.hex_serial_number)
        if Primitive.to_proto(resource.lifetime):
            res.lifetime = Primitive.to_proto(resource.lifetime)
        if Primitive.to_proto(resource.not_before_time):
            res.not_before_time = Primitive.to_proto(resource.not_before_time)
        if Primitive.to_proto(resource.not_after_time):
            res.not_after_time = Primitive.to_proto(resource.not_after_time)
        if Primitive.to_proto(resource.common_name):
            res.common_name = Primitive.to_proto(resource.common_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsSubjectDescription(
            subject=CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject.from_proto(
                resource.subject
            ),
            subject_alt_name=CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName.from_proto(
                resource.subject_alt_name
            ),
            hex_serial_number=Primitive.from_proto(resource.hex_serial_number),
            lifetime=Primitive.from_proto(resource.lifetime),
            not_before_time=Primitive.from_proto(resource.not_before_time),
            not_after_time=Primitive.from_proto(resource.not_after_time),
            common_name=Primitive.from_proto(resource.common_name),
        )


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescription.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescription.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(object):
    def __init__(
        self,
        common_name: str = None,
        country_code: str = None,
        organization: str = None,
        organizational_unit: str = None,
        locality: str = None,
        province: str = None,
        street_address: str = None,
        postal_code: str = None,
    ):
        self.common_name = common_name
        self.country_code = country_code
        self.organization = organization
        self.organizational_unit = organizational_unit
        self.locality = locality
        self.province = province
        self.street_address = street_address
        self.postal_code = postal_code

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject()
        )
        if Primitive.to_proto(resource.common_name):
            res.common_name = Primitive.to_proto(resource.common_name)
        if Primitive.to_proto(resource.country_code):
            res.country_code = Primitive.to_proto(resource.country_code)
        if Primitive.to_proto(resource.organization):
            res.organization = Primitive.to_proto(resource.organization)
        if Primitive.to_proto(resource.organizational_unit):
            res.organizational_unit = Primitive.to_proto(resource.organizational_unit)
        if Primitive.to_proto(resource.locality):
            res.locality = Primitive.to_proto(resource.locality)
        if Primitive.to_proto(resource.province):
            res.province = Primitive.to_proto(resource.province)
        if Primitive.to_proto(resource.street_address):
            res.street_address = Primitive.to_proto(resource.street_address)
        if Primitive.to_proto(resource.postal_code):
            res.postal_code = Primitive.to_proto(resource.postal_code)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject(
            common_name=Primitive.from_proto(resource.common_name),
            country_code=Primitive.from_proto(resource.country_code),
            organization=Primitive.from_proto(resource.organization),
            organizational_unit=Primitive.from_proto(resource.organizational_unit),
            locality=Primitive.from_proto(resource.locality),
            province=Primitive.from_proto(resource.province),
            street_address=Primitive.from_proto(resource.street_address),
            postal_code=Primitive.from_proto(resource.postal_code),
        )


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(
    object
):
    def __init__(
        self,
        dns_names: list = None,
        uris: list = None,
        email_addresses: list = None,
        ip_addresses: list = None,
        custom_sans: list = None,
    ):
        self.dns_names = dns_names
        self.uris = uris
        self.email_addresses = email_addresses
        self.ip_addresses = ip_addresses
        self.custom_sans = custom_sans

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName()
        )
        if Primitive.to_proto(resource.dns_names):
            res.dns_names.extend(Primitive.to_proto(resource.dns_names))
        if Primitive.to_proto(resource.uris):
            res.uris.extend(Primitive.to_proto(resource.uris))
        if Primitive.to_proto(resource.email_addresses):
            res.email_addresses.extend(Primitive.to_proto(resource.email_addresses))
        if Primitive.to_proto(resource.ip_addresses):
            res.ip_addresses.extend(Primitive.to_proto(resource.ip_addresses))
        if CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansArray.to_proto(
            resource.custom_sans
        ):
            res.custom_sans.extend(
                CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansArray.to_proto(
                    resource.custom_sans
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName(
            dns_names=Primitive.from_proto(resource.dns_names),
            uris=Primitive.from_proto(resource.uris),
            email_addresses=Primitive.from_proto(resource.email_addresses),
            ip_addresses=Primitive.from_proto(resource.ip_addresses),
            custom_sans=CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansArray.from_proto(
                resource.custom_sans
            ),
        )


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(
    object
):
    def __init__(
        self, object_id: dict = None, critical: bool = None, value: str = None
    ):
        self.object_id = object_id
        self.critical = critical
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans()
        )
        if CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId.to_proto(
                    resource.object_id
                )
            )
        else:
            res.ClearField("object_id")
        if Primitive.to_proto(resource.critical):
            res.critical = Primitive.to_proto(resource.critical)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans(
            object_id=CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectIdArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsPublicKey(object):
    def __init__(self, key: str = None, format: str = None, type: str = None):
        self.key = key
        self.format = format
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKey()
        )
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum.to_proto(
            resource.format
        ):
            res.format = CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum.to_proto(
                resource.format
            )
        if CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum.to_proto(
            resource.type
        ):
            res.type = CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum.to_proto(
                resource.type
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsPublicKey(
            key=Primitive.from_proto(resource.key),
            format=CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum.from_proto(
                resource.format
            ),
            type=CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum.from_proto(
                resource.type
            ),
        )


class CertificateAuthorityCaCertificateDescriptionsPublicKeyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsPublicKey.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsPublicKey.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsSubjectKeyId(object):
    def __init__(self, key_id: str = None):
        self.key_id = key_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsSubjectKeyId()
        )
        if Primitive.to_proto(resource.key_id):
            res.key_id = Primitive.to_proto(resource.key_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsSubjectKeyId(
            key_id=Primitive.from_proto(resource.key_id),
        )


class CertificateAuthorityCaCertificateDescriptionsSubjectKeyIdArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectKeyId.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsSubjectKeyId.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(object):
    def __init__(self, key_id: str = None):
        self.key_id = key_id

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId()
        )
        if Primitive.to_proto(resource.key_id):
            res.key_id = Primitive.to_proto(resource.key_id)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId(
            key_id=Primitive.from_proto(resource.key_id),
        )


class CertificateAuthorityCaCertificateDescriptionsAuthorityKeyIdArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsCertFingerprint(object):
    def __init__(self, sha256_hash: str = None):
        self.sha256_hash = sha256_hash

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsCertFingerprint()
        )
        if Primitive.to_proto(resource.sha256_hash):
            res.sha256_hash = Primitive.to_proto(resource.sha256_hash)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsCertFingerprint(
            sha256_hash=Primitive.from_proto(resource.sha256_hash),
        )


class CertificateAuthorityCaCertificateDescriptionsCertFingerprintArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsCertFingerprint.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsCertFingerprint.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsConfigValues(object):
    def __init__(
        self,
        key_usage: dict = None,
        ca_options: dict = None,
        policy_ids: list = None,
        aia_ocsp_servers: list = None,
        additional_extensions: list = None,
    ):
        self.key_usage = key_usage
        self.ca_options = ca_options
        self.policy_ids = policy_ids
        self.aia_ocsp_servers = aia_ocsp_servers
        self.additional_extensions = additional_extensions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValues()
        )
        if CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage.to_proto(
            resource.key_usage
        ):
            res.key_usage.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage.to_proto(
                    resource.key_usage
                )
            )
        else:
            res.ClearField("key_usage")
        if CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions.to_proto(
            resource.ca_options
        ):
            res.ca_options.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions.to_proto(
                    resource.ca_options
                )
            )
        else:
            res.ClearField("ca_options")
        if CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsArray.to_proto(
            resource.policy_ids
        ):
            res.policy_ids.extend(
                CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsArray.to_proto(
                    resource.policy_ids
                )
            )
        if Primitive.to_proto(resource.aia_ocsp_servers):
            res.aia_ocsp_servers.extend(Primitive.to_proto(resource.aia_ocsp_servers))
        if CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsArray.to_proto(
            resource.additional_extensions
        ):
            res.additional_extensions.extend(
                CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsArray.to_proto(
                    resource.additional_extensions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsConfigValues(
            key_usage=CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage.from_proto(
                resource.key_usage
            ),
            ca_options=CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions.from_proto(
                resource.ca_options
            ),
            policy_ids=CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsArray.from_proto(
                resource.policy_ids
            ),
            aia_ocsp_servers=Primitive.from_proto(resource.aia_ocsp_servers),
            additional_extensions=CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsArray.from_proto(
                resource.additional_extensions
            ),
        )


class CertificateAuthorityCaCertificateDescriptionsConfigValuesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValues.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValues.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(object):
    def __init__(
        self,
        base_key_usage: dict = None,
        extended_key_usage: dict = None,
        unknown_extended_key_usages: list = None,
    ):
        self.base_key_usage = base_key_usage
        self.extended_key_usage = extended_key_usage
        self.unknown_extended_key_usages = unknown_extended_key_usages

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage()
        )
        if CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage.to_proto(
            resource.base_key_usage
        ):
            res.base_key_usage.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage.to_proto(
                    resource.base_key_usage
                )
            )
        else:
            res.ClearField("base_key_usage")
        if CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage.to_proto(
            resource.extended_key_usage
        ):
            res.extended_key_usage.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage.to_proto(
                    resource.extended_key_usage
                )
            )
        else:
            res.ClearField("extended_key_usage")
        if CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
            resource.unknown_extended_key_usages
        ):
            res.unknown_extended_key_usages.extend(
                CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
                    resource.unknown_extended_key_usages
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage(
            base_key_usage=CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage.from_proto(
                resource.base_key_usage
            ),
            extended_key_usage=CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage.from_proto(
                resource.extended_key_usage
            ),
            unknown_extended_key_usages=CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesArray.from_proto(
                resource.unknown_extended_key_usages
            ),
        )


class CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(
    object
):
    def __init__(
        self,
        digital_signature: bool = None,
        content_commitment: bool = None,
        key_encipherment: bool = None,
        data_encipherment: bool = None,
        key_agreement: bool = None,
        cert_sign: bool = None,
        crl_sign: bool = None,
        encipher_only: bool = None,
        decipher_only: bool = None,
    ):
        self.digital_signature = digital_signature
        self.content_commitment = content_commitment
        self.key_encipherment = key_encipherment
        self.data_encipherment = data_encipherment
        self.key_agreement = key_agreement
        self.cert_sign = cert_sign
        self.crl_sign = crl_sign
        self.encipher_only = encipher_only
        self.decipher_only = decipher_only

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage()
        )
        if Primitive.to_proto(resource.digital_signature):
            res.digital_signature = Primitive.to_proto(resource.digital_signature)
        if Primitive.to_proto(resource.content_commitment):
            res.content_commitment = Primitive.to_proto(resource.content_commitment)
        if Primitive.to_proto(resource.key_encipherment):
            res.key_encipherment = Primitive.to_proto(resource.key_encipherment)
        if Primitive.to_proto(resource.data_encipherment):
            res.data_encipherment = Primitive.to_proto(resource.data_encipherment)
        if Primitive.to_proto(resource.key_agreement):
            res.key_agreement = Primitive.to_proto(resource.key_agreement)
        if Primitive.to_proto(resource.cert_sign):
            res.cert_sign = Primitive.to_proto(resource.cert_sign)
        if Primitive.to_proto(resource.crl_sign):
            res.crl_sign = Primitive.to_proto(resource.crl_sign)
        if Primitive.to_proto(resource.encipher_only):
            res.encipher_only = Primitive.to_proto(resource.encipher_only)
        if Primitive.to_proto(resource.decipher_only):
            res.decipher_only = Primitive.to_proto(resource.decipher_only)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage(
            digital_signature=Primitive.from_proto(resource.digital_signature),
            content_commitment=Primitive.from_proto(resource.content_commitment),
            key_encipherment=Primitive.from_proto(resource.key_encipherment),
            data_encipherment=Primitive.from_proto(resource.data_encipherment),
            key_agreement=Primitive.from_proto(resource.key_agreement),
            cert_sign=Primitive.from_proto(resource.cert_sign),
            crl_sign=Primitive.from_proto(resource.crl_sign),
            encipher_only=Primitive.from_proto(resource.encipher_only),
            decipher_only=Primitive.from_proto(resource.decipher_only),
        )


class CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsageArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(
    object
):
    def __init__(
        self,
        server_auth: bool = None,
        client_auth: bool = None,
        code_signing: bool = None,
        email_protection: bool = None,
        time_stamping: bool = None,
        ocsp_signing: bool = None,
    ):
        self.server_auth = server_auth
        self.client_auth = client_auth
        self.code_signing = code_signing
        self.email_protection = email_protection
        self.time_stamping = time_stamping
        self.ocsp_signing = ocsp_signing

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage()
        )
        if Primitive.to_proto(resource.server_auth):
            res.server_auth = Primitive.to_proto(resource.server_auth)
        if Primitive.to_proto(resource.client_auth):
            res.client_auth = Primitive.to_proto(resource.client_auth)
        if Primitive.to_proto(resource.code_signing):
            res.code_signing = Primitive.to_proto(resource.code_signing)
        if Primitive.to_proto(resource.email_protection):
            res.email_protection = Primitive.to_proto(resource.email_protection)
        if Primitive.to_proto(resource.time_stamping):
            res.time_stamping = Primitive.to_proto(resource.time_stamping)
        if Primitive.to_proto(resource.ocsp_signing):
            res.ocsp_signing = Primitive.to_proto(resource.ocsp_signing)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage(
            server_auth=Primitive.from_proto(resource.server_auth),
            client_auth=Primitive.from_proto(resource.client_auth),
            code_signing=Primitive.from_proto(resource.code_signing),
            email_protection=Primitive.from_proto(resource.email_protection),
            time_stamping=Primitive.from_proto(resource.time_stamping),
            ocsp_signing=Primitive.from_proto(resource.ocsp_signing),
        )


class CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsageArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsagesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(object):
    def __init__(self, is_ca: bool = None, max_issuer_path_length: int = None):
        self.is_ca = is_ca
        self.max_issuer_path_length = max_issuer_path_length

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions()
        )
        if Primitive.to_proto(resource.is_ca):
            res.is_ca = Primitive.to_proto(resource.is_ca)
        if Primitive.to_proto(resource.max_issuer_path_length):
            res.max_issuer_path_length = Primitive.to_proto(
                resource.max_issuer_path_length
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions(
            is_ca=Primitive.from_proto(resource.is_ca),
            max_issuer_path_length=Primitive.from_proto(
                resource.max_issuer_path_length
            ),
        )


class CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds(object):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIdsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions(
    object
):
    def __init__(
        self, object_id: dict = None, critical: bool = None, value: str = None
    ):
        self.object_id = object_id
        self.critical = critical
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions()
        )
        if CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId.to_proto(
                    resource.object_id
                )
            )
        else:
            res.ClearField("object_id")
        if Primitive.to_proto(resource.critical):
            res.critical = Primitive.to_proto(resource.critical)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions(
            object_id=CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectIdArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityAccessUrls(object):
    def __init__(
        self,
        ca_certificate_access_url: str = None,
        crl_access_urls: list = None,
        crl_access_url: str = None,
    ):
        self.ca_certificate_access_url = ca_certificate_access_url
        self.crl_access_urls = crl_access_urls
        self.crl_access_url = crl_access_url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = certificate_authority_pb2.PrivatecaBetaCertificateAuthorityAccessUrls()
        if Primitive.to_proto(resource.ca_certificate_access_url):
            res.ca_certificate_access_url = Primitive.to_proto(
                resource.ca_certificate_access_url
            )
        if Primitive.to_proto(resource.crl_access_urls):
            res.crl_access_urls.extend(Primitive.to_proto(resource.crl_access_urls))
        if Primitive.to_proto(resource.crl_access_url):
            res.crl_access_url = Primitive.to_proto(resource.crl_access_url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityAccessUrls(
            ca_certificate_access_url=Primitive.from_proto(
                resource.ca_certificate_access_url
            ),
            crl_access_urls=Primitive.from_proto(resource.crl_access_urls),
            crl_access_url=Primitive.from_proto(resource.crl_access_url),
        )


class CertificateAuthorityAccessUrlsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthorityAccessUrls.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateAuthorityAccessUrls.from_proto(i) for i in resources]


class CertificateAuthorityCertificatePolicy(object):
    def __init__(
        self,
        allowed_config_list: dict = None,
        overwrite_config_values: dict = None,
        allowed_locations_and_organizations: list = None,
        allowed_common_names: list = None,
        allowed_sans: dict = None,
        maximum_lifetime: str = None,
        allowed_issuance_modes: dict = None,
    ):
        self.allowed_config_list = allowed_config_list
        self.overwrite_config_values = overwrite_config_values
        self.allowed_locations_and_organizations = allowed_locations_and_organizations
        self.allowed_common_names = allowed_common_names
        self.allowed_sans = allowed_sans
        self.maximum_lifetime = maximum_lifetime
        self.allowed_issuance_modes = allowed_issuance_modes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicy()
        )
        if CertificateAuthorityCertificatePolicyAllowedConfigList.to_proto(
            resource.allowed_config_list
        ):
            res.allowed_config_list.CopyFrom(
                CertificateAuthorityCertificatePolicyAllowedConfigList.to_proto(
                    resource.allowed_config_list
                )
            )
        else:
            res.ClearField("allowed_config_list")
        if CertificateAuthorityCertificatePolicyOverwriteConfigValues.to_proto(
            resource.overwrite_config_values
        ):
            res.overwrite_config_values.CopyFrom(
                CertificateAuthorityCertificatePolicyOverwriteConfigValues.to_proto(
                    resource.overwrite_config_values
                )
            )
        else:
            res.ClearField("overwrite_config_values")
        if CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsArray.to_proto(
            resource.allowed_locations_and_organizations
        ):
            res.allowed_locations_and_organizations.extend(
                CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsArray.to_proto(
                    resource.allowed_locations_and_organizations
                )
            )
        if Primitive.to_proto(resource.allowed_common_names):
            res.allowed_common_names.extend(
                Primitive.to_proto(resource.allowed_common_names)
            )
        if CertificateAuthorityCertificatePolicyAllowedSans.to_proto(
            resource.allowed_sans
        ):
            res.allowed_sans.CopyFrom(
                CertificateAuthorityCertificatePolicyAllowedSans.to_proto(
                    resource.allowed_sans
                )
            )
        else:
            res.ClearField("allowed_sans")
        if Primitive.to_proto(resource.maximum_lifetime):
            res.maximum_lifetime = Primitive.to_proto(resource.maximum_lifetime)
        if CertificateAuthorityCertificatePolicyAllowedIssuanceModes.to_proto(
            resource.allowed_issuance_modes
        ):
            res.allowed_issuance_modes.CopyFrom(
                CertificateAuthorityCertificatePolicyAllowedIssuanceModes.to_proto(
                    resource.allowed_issuance_modes
                )
            )
        else:
            res.ClearField("allowed_issuance_modes")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicy(
            allowed_config_list=CertificateAuthorityCertificatePolicyAllowedConfigList.from_proto(
                resource.allowed_config_list
            ),
            overwrite_config_values=CertificateAuthorityCertificatePolicyOverwriteConfigValues.from_proto(
                resource.overwrite_config_values
            ),
            allowed_locations_and_organizations=CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsArray.from_proto(
                resource.allowed_locations_and_organizations
            ),
            allowed_common_names=Primitive.from_proto(resource.allowed_common_names),
            allowed_sans=CertificateAuthorityCertificatePolicyAllowedSans.from_proto(
                resource.allowed_sans
            ),
            maximum_lifetime=Primitive.from_proto(resource.maximum_lifetime),
            allowed_issuance_modes=CertificateAuthorityCertificatePolicyAllowedIssuanceModes.from_proto(
                resource.allowed_issuance_modes
            ),
        )


class CertificateAuthorityCertificatePolicyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthorityCertificatePolicy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateAuthorityCertificatePolicy.from_proto(i) for i in resources]


class CertificateAuthorityCertificatePolicyAllowedConfigList(object):
    def __init__(self, allowed_config_values: list = None):
        self.allowed_config_values = allowed_config_values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigList()
        )
        if CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesArray.to_proto(
            resource.allowed_config_values
        ):
            res.allowed_config_values.extend(
                CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesArray.to_proto(
                    resource.allowed_config_values
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyAllowedConfigList(
            allowed_config_values=CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesArray.from_proto(
                resource.allowed_config_values
            ),
        )


class CertificateAuthorityCertificatePolicyAllowedConfigListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigList.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigList.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues(object):
    def __init__(
        self, reusable_config: str = None, reusable_config_values: dict = None
    ):
        self.reusable_config = reusable_config
        self.reusable_config_values = reusable_config_values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues()
        )
        if Primitive.to_proto(resource.reusable_config):
            res.reusable_config = Primitive.to_proto(resource.reusable_config)
        if CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues.to_proto(
            resource.reusable_config_values
        ):
            res.reusable_config_values.CopyFrom(
                CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues.to_proto(
                    resource.reusable_config_values
                )
            )
        else:
            res.ClearField("reusable_config_values")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues(
            reusable_config=Primitive.from_proto(resource.reusable_config),
            reusable_config_values=CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues.from_proto(
                resource.reusable_config_values
            ),
        )


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(
    object
):
    def __init__(
        self,
        key_usage: dict = None,
        ca_options: dict = None,
        policy_ids: list = None,
        aia_ocsp_servers: list = None,
        additional_extensions: list = None,
    ):
        self.key_usage = key_usage
        self.ca_options = ca_options
        self.policy_ids = policy_ids
        self.aia_ocsp_servers = aia_ocsp_servers
        self.additional_extensions = additional_extensions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues()
        )
        if CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage.to_proto(
            resource.key_usage
        ):
            res.key_usage.CopyFrom(
                CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage.to_proto(
                    resource.key_usage
                )
            )
        else:
            res.ClearField("key_usage")
        if CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions.to_proto(
            resource.ca_options
        ):
            res.ca_options.CopyFrom(
                CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions.to_proto(
                    resource.ca_options
                )
            )
        else:
            res.ClearField("ca_options")
        if CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsArray.to_proto(
            resource.policy_ids
        ):
            res.policy_ids.extend(
                CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsArray.to_proto(
                    resource.policy_ids
                )
            )
        if Primitive.to_proto(resource.aia_ocsp_servers):
            res.aia_ocsp_servers.extend(Primitive.to_proto(resource.aia_ocsp_servers))
        if CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsArray.to_proto(
            resource.additional_extensions
        ):
            res.additional_extensions.extend(
                CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsArray.to_proto(
                    resource.additional_extensions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues(
            key_usage=CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage.from_proto(
                resource.key_usage
            ),
            ca_options=CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions.from_proto(
                resource.ca_options
            ),
            policy_ids=CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsArray.from_proto(
                resource.policy_ids
            ),
            aia_ocsp_servers=Primitive.from_proto(resource.aia_ocsp_servers),
            additional_extensions=CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsArray.from_proto(
                resource.additional_extensions
            ),
        )


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(
    object
):
    def __init__(
        self,
        base_key_usage: dict = None,
        extended_key_usage: dict = None,
        unknown_extended_key_usages: list = None,
    ):
        self.base_key_usage = base_key_usage
        self.extended_key_usage = extended_key_usage
        self.unknown_extended_key_usages = unknown_extended_key_usages

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage()
        )
        if CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage.to_proto(
            resource.base_key_usage
        ):
            res.base_key_usage.CopyFrom(
                CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage.to_proto(
                    resource.base_key_usage
                )
            )
        else:
            res.ClearField("base_key_usage")
        if CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage.to_proto(
            resource.extended_key_usage
        ):
            res.extended_key_usage.CopyFrom(
                CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage.to_proto(
                    resource.extended_key_usage
                )
            )
        else:
            res.ClearField("extended_key_usage")
        if CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
            resource.unknown_extended_key_usages
        ):
            res.unknown_extended_key_usages.extend(
                CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
                    resource.unknown_extended_key_usages
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage(
            base_key_usage=CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage.from_proto(
                resource.base_key_usage
            ),
            extended_key_usage=CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage.from_proto(
                resource.extended_key_usage
            ),
            unknown_extended_key_usages=CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesArray.from_proto(
                resource.unknown_extended_key_usages
            ),
        )


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(
    object
):
    def __init__(
        self,
        digital_signature: bool = None,
        content_commitment: bool = None,
        key_encipherment: bool = None,
        data_encipherment: bool = None,
        key_agreement: bool = None,
        cert_sign: bool = None,
        crl_sign: bool = None,
        encipher_only: bool = None,
        decipher_only: bool = None,
    ):
        self.digital_signature = digital_signature
        self.content_commitment = content_commitment
        self.key_encipherment = key_encipherment
        self.data_encipherment = data_encipherment
        self.key_agreement = key_agreement
        self.cert_sign = cert_sign
        self.crl_sign = crl_sign
        self.encipher_only = encipher_only
        self.decipher_only = decipher_only

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage()
        )
        if Primitive.to_proto(resource.digital_signature):
            res.digital_signature = Primitive.to_proto(resource.digital_signature)
        if Primitive.to_proto(resource.content_commitment):
            res.content_commitment = Primitive.to_proto(resource.content_commitment)
        if Primitive.to_proto(resource.key_encipherment):
            res.key_encipherment = Primitive.to_proto(resource.key_encipherment)
        if Primitive.to_proto(resource.data_encipherment):
            res.data_encipherment = Primitive.to_proto(resource.data_encipherment)
        if Primitive.to_proto(resource.key_agreement):
            res.key_agreement = Primitive.to_proto(resource.key_agreement)
        if Primitive.to_proto(resource.cert_sign):
            res.cert_sign = Primitive.to_proto(resource.cert_sign)
        if Primitive.to_proto(resource.crl_sign):
            res.crl_sign = Primitive.to_proto(resource.crl_sign)
        if Primitive.to_proto(resource.encipher_only):
            res.encipher_only = Primitive.to_proto(resource.encipher_only)
        if Primitive.to_proto(resource.decipher_only):
            res.decipher_only = Primitive.to_proto(resource.decipher_only)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(
            digital_signature=Primitive.from_proto(resource.digital_signature),
            content_commitment=Primitive.from_proto(resource.content_commitment),
            key_encipherment=Primitive.from_proto(resource.key_encipherment),
            data_encipherment=Primitive.from_proto(resource.data_encipherment),
            key_agreement=Primitive.from_proto(resource.key_agreement),
            cert_sign=Primitive.from_proto(resource.cert_sign),
            crl_sign=Primitive.from_proto(resource.crl_sign),
            encipher_only=Primitive.from_proto(resource.encipher_only),
            decipher_only=Primitive.from_proto(resource.decipher_only),
        )


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(
    object
):
    def __init__(
        self,
        server_auth: bool = None,
        client_auth: bool = None,
        code_signing: bool = None,
        email_protection: bool = None,
        time_stamping: bool = None,
        ocsp_signing: bool = None,
    ):
        self.server_auth = server_auth
        self.client_auth = client_auth
        self.code_signing = code_signing
        self.email_protection = email_protection
        self.time_stamping = time_stamping
        self.ocsp_signing = ocsp_signing

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage()
        )
        if Primitive.to_proto(resource.server_auth):
            res.server_auth = Primitive.to_proto(resource.server_auth)
        if Primitive.to_proto(resource.client_auth):
            res.client_auth = Primitive.to_proto(resource.client_auth)
        if Primitive.to_proto(resource.code_signing):
            res.code_signing = Primitive.to_proto(resource.code_signing)
        if Primitive.to_proto(resource.email_protection):
            res.email_protection = Primitive.to_proto(resource.email_protection)
        if Primitive.to_proto(resource.time_stamping):
            res.time_stamping = Primitive.to_proto(resource.time_stamping)
        if Primitive.to_proto(resource.ocsp_signing):
            res.ocsp_signing = Primitive.to_proto(resource.ocsp_signing)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(
            server_auth=Primitive.from_proto(resource.server_auth),
            client_auth=Primitive.from_proto(resource.client_auth),
            code_signing=Primitive.from_proto(resource.code_signing),
            email_protection=Primitive.from_proto(resource.email_protection),
            time_stamping=Primitive.from_proto(resource.time_stamping),
            ocsp_signing=Primitive.from_proto(resource.ocsp_signing),
        )


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(
    object
):
    def __init__(self, is_ca: bool = None, max_issuer_path_length: int = None):
        self.is_ca = is_ca
        self.max_issuer_path_length = max_issuer_path_length

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions()
        )
        if Primitive.to_proto(resource.is_ca):
            res.is_ca = Primitive.to_proto(resource.is_ca)
        if Primitive.to_proto(resource.max_issuer_path_length):
            res.max_issuer_path_length = Primitive.to_proto(
                resource.max_issuer_path_length
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions(
            is_ca=Primitive.from_proto(resource.is_ca),
            max_issuer_path_length=Primitive.from_proto(
                resource.max_issuer_path_length
            ),
        )


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIdsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions(
    object
):
    def __init__(
        self, object_id: dict = None, critical: bool = None, value: str = None
    ):
        self.object_id = object_id
        self.critical = critical
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions()
        )
        if CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId.to_proto(
                    resource.object_id
                )
            )
        else:
            res.ClearField("object_id")
        if Primitive.to_proto(resource.critical):
            res.critical = Primitive.to_proto(resource.critical)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions(
            object_id=CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyOverwriteConfigValues(object):
    def __init__(
        self, reusable_config: str = None, reusable_config_values: dict = None
    ):
        self.reusable_config = reusable_config
        self.reusable_config_values = reusable_config_values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValues()
        )
        if Primitive.to_proto(resource.reusable_config):
            res.reusable_config = Primitive.to_proto(resource.reusable_config)
        if CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues.to_proto(
            resource.reusable_config_values
        ):
            res.reusable_config_values.CopyFrom(
                CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues.to_proto(
                    resource.reusable_config_values
                )
            )
        else:
            res.ClearField("reusable_config_values")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyOverwriteConfigValues(
            reusable_config=Primitive.from_proto(resource.reusable_config),
            reusable_config_values=CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues.from_proto(
                resource.reusable_config_values
            ),
        )


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValues.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValues.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(
    object
):
    def __init__(
        self,
        key_usage: dict = None,
        ca_options: dict = None,
        policy_ids: list = None,
        aia_ocsp_servers: list = None,
        additional_extensions: list = None,
    ):
        self.key_usage = key_usage
        self.ca_options = ca_options
        self.policy_ids = policy_ids
        self.aia_ocsp_servers = aia_ocsp_servers
        self.additional_extensions = additional_extensions

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues()
        )
        if CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage.to_proto(
            resource.key_usage
        ):
            res.key_usage.CopyFrom(
                CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage.to_proto(
                    resource.key_usage
                )
            )
        else:
            res.ClearField("key_usage")
        if CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions.to_proto(
            resource.ca_options
        ):
            res.ca_options.CopyFrom(
                CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions.to_proto(
                    resource.ca_options
                )
            )
        else:
            res.ClearField("ca_options")
        if CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsArray.to_proto(
            resource.policy_ids
        ):
            res.policy_ids.extend(
                CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsArray.to_proto(
                    resource.policy_ids
                )
            )
        if Primitive.to_proto(resource.aia_ocsp_servers):
            res.aia_ocsp_servers.extend(Primitive.to_proto(resource.aia_ocsp_servers))
        if CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsArray.to_proto(
            resource.additional_extensions
        ):
            res.additional_extensions.extend(
                CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsArray.to_proto(
                    resource.additional_extensions
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues(
            key_usage=CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage.from_proto(
                resource.key_usage
            ),
            ca_options=CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions.from_proto(
                resource.ca_options
            ),
            policy_ids=CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsArray.from_proto(
                resource.policy_ids
            ),
            aia_ocsp_servers=Primitive.from_proto(resource.aia_ocsp_servers),
            additional_extensions=CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsArray.from_proto(
                resource.additional_extensions
            ),
        )


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(
    object
):
    def __init__(
        self,
        base_key_usage: dict = None,
        extended_key_usage: dict = None,
        unknown_extended_key_usages: list = None,
    ):
        self.base_key_usage = base_key_usage
        self.extended_key_usage = extended_key_usage
        self.unknown_extended_key_usages = unknown_extended_key_usages

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage()
        )
        if CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage.to_proto(
            resource.base_key_usage
        ):
            res.base_key_usage.CopyFrom(
                CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage.to_proto(
                    resource.base_key_usage
                )
            )
        else:
            res.ClearField("base_key_usage")
        if CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage.to_proto(
            resource.extended_key_usage
        ):
            res.extended_key_usage.CopyFrom(
                CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage.to_proto(
                    resource.extended_key_usage
                )
            )
        else:
            res.ClearField("extended_key_usage")
        if CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
            resource.unknown_extended_key_usages
        ):
            res.unknown_extended_key_usages.extend(
                CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesArray.to_proto(
                    resource.unknown_extended_key_usages
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage(
            base_key_usage=CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage.from_proto(
                resource.base_key_usage
            ),
            extended_key_usage=CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage.from_proto(
                resource.extended_key_usage
            ),
            unknown_extended_key_usages=CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesArray.from_proto(
                resource.unknown_extended_key_usages
            ),
        )


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(
    object
):
    def __init__(
        self,
        digital_signature: bool = None,
        content_commitment: bool = None,
        key_encipherment: bool = None,
        data_encipherment: bool = None,
        key_agreement: bool = None,
        cert_sign: bool = None,
        crl_sign: bool = None,
        encipher_only: bool = None,
        decipher_only: bool = None,
    ):
        self.digital_signature = digital_signature
        self.content_commitment = content_commitment
        self.key_encipherment = key_encipherment
        self.data_encipherment = data_encipherment
        self.key_agreement = key_agreement
        self.cert_sign = cert_sign
        self.crl_sign = crl_sign
        self.encipher_only = encipher_only
        self.decipher_only = decipher_only

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage()
        )
        if Primitive.to_proto(resource.digital_signature):
            res.digital_signature = Primitive.to_proto(resource.digital_signature)
        if Primitive.to_proto(resource.content_commitment):
            res.content_commitment = Primitive.to_proto(resource.content_commitment)
        if Primitive.to_proto(resource.key_encipherment):
            res.key_encipherment = Primitive.to_proto(resource.key_encipherment)
        if Primitive.to_proto(resource.data_encipherment):
            res.data_encipherment = Primitive.to_proto(resource.data_encipherment)
        if Primitive.to_proto(resource.key_agreement):
            res.key_agreement = Primitive.to_proto(resource.key_agreement)
        if Primitive.to_proto(resource.cert_sign):
            res.cert_sign = Primitive.to_proto(resource.cert_sign)
        if Primitive.to_proto(resource.crl_sign):
            res.crl_sign = Primitive.to_proto(resource.crl_sign)
        if Primitive.to_proto(resource.encipher_only):
            res.encipher_only = Primitive.to_proto(resource.encipher_only)
        if Primitive.to_proto(resource.decipher_only):
            res.decipher_only = Primitive.to_proto(resource.decipher_only)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage(
            digital_signature=Primitive.from_proto(resource.digital_signature),
            content_commitment=Primitive.from_proto(resource.content_commitment),
            key_encipherment=Primitive.from_proto(resource.key_encipherment),
            data_encipherment=Primitive.from_proto(resource.data_encipherment),
            key_agreement=Primitive.from_proto(resource.key_agreement),
            cert_sign=Primitive.from_proto(resource.cert_sign),
            crl_sign=Primitive.from_proto(resource.crl_sign),
            encipher_only=Primitive.from_proto(resource.encipher_only),
            decipher_only=Primitive.from_proto(resource.decipher_only),
        )


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsageArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(
    object
):
    def __init__(
        self,
        server_auth: bool = None,
        client_auth: bool = None,
        code_signing: bool = None,
        email_protection: bool = None,
        time_stamping: bool = None,
        ocsp_signing: bool = None,
    ):
        self.server_auth = server_auth
        self.client_auth = client_auth
        self.code_signing = code_signing
        self.email_protection = email_protection
        self.time_stamping = time_stamping
        self.ocsp_signing = ocsp_signing

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage()
        )
        if Primitive.to_proto(resource.server_auth):
            res.server_auth = Primitive.to_proto(resource.server_auth)
        if Primitive.to_proto(resource.client_auth):
            res.client_auth = Primitive.to_proto(resource.client_auth)
        if Primitive.to_proto(resource.code_signing):
            res.code_signing = Primitive.to_proto(resource.code_signing)
        if Primitive.to_proto(resource.email_protection):
            res.email_protection = Primitive.to_proto(resource.email_protection)
        if Primitive.to_proto(resource.time_stamping):
            res.time_stamping = Primitive.to_proto(resource.time_stamping)
        if Primitive.to_proto(resource.ocsp_signing):
            res.ocsp_signing = Primitive.to_proto(resource.ocsp_signing)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage(
            server_auth=Primitive.from_proto(resource.server_auth),
            client_auth=Primitive.from_proto(resource.client_auth),
            code_signing=Primitive.from_proto(resource.code_signing),
            email_protection=Primitive.from_proto(resource.email_protection),
            time_stamping=Primitive.from_proto(resource.time_stamping),
            ocsp_signing=Primitive.from_proto(resource.ocsp_signing),
        )


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsageArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsagesArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(
    object
):
    def __init__(self, is_ca: bool = None, max_issuer_path_length: int = None):
        self.is_ca = is_ca
        self.max_issuer_path_length = max_issuer_path_length

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions()
        )
        if Primitive.to_proto(resource.is_ca):
            res.is_ca = Primitive.to_proto(resource.is_ca)
        if Primitive.to_proto(resource.max_issuer_path_length):
            res.max_issuer_path_length = Primitive.to_proto(
                resource.max_issuer_path_length
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions(
            is_ca=Primitive.from_proto(resource.is_ca),
            max_issuer_path_length=Primitive.from_proto(
                resource.max_issuer_path_length
            ),
        )


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIdsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions(
    object
):
    def __init__(
        self, object_id: dict = None, critical: bool = None, value: str = None
    ):
        self.object_id = object_id
        self.critical = critical
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions()
        )
        if CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId.to_proto(
            resource.object_id
        ):
            res.object_id.CopyFrom(
                CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId.to_proto(
                    resource.object_id
                )
            )
        else:
            res.ClearField("object_id")
        if Primitive.to_proto(resource.critical):
            res.critical = Primitive.to_proto(resource.critical)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions(
            object_id=CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId.from_proto(
                resource.object_id
            ),
            critical=Primitive.from_proto(resource.critical),
            value=Primitive.from_proto(resource.value),
        )


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(
    object
):
    def __init__(self, object_id_path: list = None):
        self.object_id_path = object_id_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId()
        )
        if int64Array.to_proto(resource.object_id_path):
            res.object_id_path.extend(int64Array.to_proto(resource.object_id_path))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId(
            object_id_path=int64Array.from_proto(resource.object_id_path),
        )


class CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectIdArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations(object):
    def __init__(
        self,
        country_code: str = None,
        organization: str = None,
        organizational_unit: str = None,
        locality: str = None,
        province: str = None,
        street_address: str = None,
        postal_code: str = None,
    ):
        self.country_code = country_code
        self.organization = organization
        self.organizational_unit = organizational_unit
        self.locality = locality
        self.province = province
        self.street_address = street_address
        self.postal_code = postal_code

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations()
        )
        if Primitive.to_proto(resource.country_code):
            res.country_code = Primitive.to_proto(resource.country_code)
        if Primitive.to_proto(resource.organization):
            res.organization = Primitive.to_proto(resource.organization)
        if Primitive.to_proto(resource.organizational_unit):
            res.organizational_unit = Primitive.to_proto(resource.organizational_unit)
        if Primitive.to_proto(resource.locality):
            res.locality = Primitive.to_proto(resource.locality)
        if Primitive.to_proto(resource.province):
            res.province = Primitive.to_proto(resource.province)
        if Primitive.to_proto(resource.street_address):
            res.street_address = Primitive.to_proto(resource.street_address)
        if Primitive.to_proto(resource.postal_code):
            res.postal_code = Primitive.to_proto(resource.postal_code)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations(
            country_code=Primitive.from_proto(resource.country_code),
            organization=Primitive.from_proto(resource.organization),
            organizational_unit=Primitive.from_proto(resource.organizational_unit),
            locality=Primitive.from_proto(resource.locality),
            province=Primitive.from_proto(resource.province),
            street_address=Primitive.from_proto(resource.street_address),
            postal_code=Primitive.from_proto(resource.postal_code),
        )


class CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizationsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations.from_proto(
                i
            )
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyAllowedSans(object):
    def __init__(
        self,
        allowed_dns_names: list = None,
        allowed_uris: list = None,
        allowed_email_addresses: list = None,
        allowed_ips: list = None,
        allow_globbing_dns_wildcards: bool = None,
        allow_custom_sans: bool = None,
    ):
        self.allowed_dns_names = allowed_dns_names
        self.allowed_uris = allowed_uris
        self.allowed_email_addresses = allowed_email_addresses
        self.allowed_ips = allowed_ips
        self.allow_globbing_dns_wildcards = allow_globbing_dns_wildcards
        self.allow_custom_sans = allow_custom_sans

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedSans()
        )
        if Primitive.to_proto(resource.allowed_dns_names):
            res.allowed_dns_names.extend(Primitive.to_proto(resource.allowed_dns_names))
        if Primitive.to_proto(resource.allowed_uris):
            res.allowed_uris.extend(Primitive.to_proto(resource.allowed_uris))
        if Primitive.to_proto(resource.allowed_email_addresses):
            res.allowed_email_addresses.extend(
                Primitive.to_proto(resource.allowed_email_addresses)
            )
        if Primitive.to_proto(resource.allowed_ips):
            res.allowed_ips.extend(Primitive.to_proto(resource.allowed_ips))
        if Primitive.to_proto(resource.allow_globbing_dns_wildcards):
            res.allow_globbing_dns_wildcards = Primitive.to_proto(
                resource.allow_globbing_dns_wildcards
            )
        if Primitive.to_proto(resource.allow_custom_sans):
            res.allow_custom_sans = Primitive.to_proto(resource.allow_custom_sans)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyAllowedSans(
            allowed_dns_names=Primitive.from_proto(resource.allowed_dns_names),
            allowed_uris=Primitive.from_proto(resource.allowed_uris),
            allowed_email_addresses=Primitive.from_proto(
                resource.allowed_email_addresses
            ),
            allowed_ips=Primitive.from_proto(resource.allowed_ips),
            allow_globbing_dns_wildcards=Primitive.from_proto(
                resource.allow_globbing_dns_wildcards
            ),
            allow_custom_sans=Primitive.from_proto(resource.allow_custom_sans),
        )


class CertificateAuthorityCertificatePolicyAllowedSansArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyAllowedSans.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyAllowedSans.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityCertificatePolicyAllowedIssuanceModes(object):
    def __init__(
        self,
        allow_csr_based_issuance: bool = None,
        allow_config_based_issuance: bool = None,
    ):
        self.allow_csr_based_issuance = allow_csr_based_issuance
        self.allow_config_based_issuance = allow_config_based_issuance

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCertificatePolicyAllowedIssuanceModes()
        )
        if Primitive.to_proto(resource.allow_csr_based_issuance):
            res.allow_csr_based_issuance = Primitive.to_proto(
                resource.allow_csr_based_issuance
            )
        if Primitive.to_proto(resource.allow_config_based_issuance):
            res.allow_config_based_issuance = Primitive.to_proto(
                resource.allow_config_based_issuance
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityCertificatePolicyAllowedIssuanceModes(
            allow_csr_based_issuance=Primitive.from_proto(
                resource.allow_csr_based_issuance
            ),
            allow_config_based_issuance=Primitive.from_proto(
                resource.allow_config_based_issuance
            ),
        )


class CertificateAuthorityCertificatePolicyAllowedIssuanceModesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            CertificateAuthorityCertificatePolicyAllowedIssuanceModes.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            CertificateAuthorityCertificatePolicyAllowedIssuanceModes.from_proto(i)
            for i in resources
        ]


class CertificateAuthorityIssuingOptions(object):
    def __init__(
        self, include_ca_cert_url: bool = None, include_crl_access_url: bool = None
    ):
        self.include_ca_cert_url = include_ca_cert_url
        self.include_crl_access_url = include_crl_access_url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            certificate_authority_pb2.PrivatecaBetaCertificateAuthorityIssuingOptions()
        )
        if Primitive.to_proto(resource.include_ca_cert_url):
            res.include_ca_cert_url = Primitive.to_proto(resource.include_ca_cert_url)
        if Primitive.to_proto(resource.include_crl_access_url):
            res.include_crl_access_url = Primitive.to_proto(
                resource.include_crl_access_url
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return CertificateAuthorityIssuingOptions(
            include_ca_cert_url=Primitive.from_proto(resource.include_ca_cert_url),
            include_crl_access_url=Primitive.from_proto(
                resource.include_crl_access_url
            ),
        )


class CertificateAuthorityIssuingOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [CertificateAuthorityIssuingOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [CertificateAuthorityIssuingOptions.from_proto(i) for i in resources]


class CertificateAuthorityTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaBetaCertificateAuthorityTypeEnum.Value(
            "PrivatecaBetaCertificateAuthorityTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaBetaCertificateAuthorityTypeEnum.Name(
            resource
        )[len("PrivatecaBetaCertificateAuthorityTypeEnum") :]


class CertificateAuthorityConfigPublicKeyTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnum.Value(
            "PrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnum.Name(
            resource
        )[
            len("PrivatecaBetaCertificateAuthorityConfigPublicKeyTypeEnum") :
        ]


class CertificateAuthorityKeySpecAlgorithmEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum.Value(
            "PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum.Name(
            resource
        )[
            len("PrivatecaBetaCertificateAuthorityKeySpecAlgorithmEnum") :
        ]


class CertificateAuthorityTierEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaBetaCertificateAuthorityTierEnum.Value(
            "PrivatecaBetaCertificateAuthorityTierEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaBetaCertificateAuthorityTierEnum.Name(
            resource
        )[len("PrivatecaBetaCertificateAuthorityTierEnum") :]


class CertificateAuthorityStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaBetaCertificateAuthorityStateEnum.Value(
            "PrivatecaBetaCertificateAuthorityStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaBetaCertificateAuthorityStateEnum.Name(
            resource
        )[
            len("PrivatecaBetaCertificateAuthorityStateEnum") :
        ]


class CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum.Value(
            "PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum.Name(
            resource
        )[
            len(
                "PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum"
            ) :
        ]


class CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum.Value(
            "PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return certificate_authority_pb2.PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum.Name(
            resource
        )[
            len(
                "PrivatecaBetaCertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum"
            ) :
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
