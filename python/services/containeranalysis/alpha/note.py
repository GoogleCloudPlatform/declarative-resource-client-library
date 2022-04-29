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
from google3.cloud.graphite.mmv2.services.google.container_analysis import note_pb2
from google3.cloud.graphite.mmv2.services.google.container_analysis import note_pb2_grpc

from typing import List


class Note(object):
    def __init__(
        self,
        name: str = None,
        short_description: str = None,
        long_description: str = None,
        related_url: list = None,
        expiration_time: str = None,
        create_time: str = None,
        update_time: str = None,
        image: dict = None,
        package: dict = None,
        discovery: dict = None,
        deployment: dict = None,
        attestation: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.short_description = short_description
        self.long_description = long_description
        self.related_url = related_url
        self.expiration_time = expiration_time
        self.image = image
        self.package = package
        self.discovery = discovery
        self.deployment = deployment
        self.attestation = attestation
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = note_pb2_grpc.ContaineranalysisAlphaNoteServiceStub(channel.Channel())
        request = note_pb2.ApplyContaineranalysisAlphaNoteRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.short_description):
            request.resource.short_description = Primitive.to_proto(
                self.short_description
            )

        if Primitive.to_proto(self.long_description):
            request.resource.long_description = Primitive.to_proto(
                self.long_description
            )

        if NoteRelatedUrlArray.to_proto(self.related_url):
            request.resource.related_url.extend(
                NoteRelatedUrlArray.to_proto(self.related_url)
            )
        if Primitive.to_proto(self.expiration_time):
            request.resource.expiration_time = Primitive.to_proto(self.expiration_time)

        if NoteImage.to_proto(self.image):
            request.resource.image.CopyFrom(NoteImage.to_proto(self.image))
        else:
            request.resource.ClearField("image")
        if NotePackage.to_proto(self.package):
            request.resource.package.CopyFrom(NotePackage.to_proto(self.package))
        else:
            request.resource.ClearField("package")
        if NoteDiscovery.to_proto(self.discovery):
            request.resource.discovery.CopyFrom(NoteDiscovery.to_proto(self.discovery))
        else:
            request.resource.ClearField("discovery")
        if NoteDeployment.to_proto(self.deployment):
            request.resource.deployment.CopyFrom(
                NoteDeployment.to_proto(self.deployment)
            )
        else:
            request.resource.ClearField("deployment")
        if NoteAttestation.to_proto(self.attestation):
            request.resource.attestation.CopyFrom(
                NoteAttestation.to_proto(self.attestation)
            )
        else:
            request.resource.ClearField("attestation")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyContaineranalysisAlphaNote(request)
        self.name = Primitive.from_proto(response.name)
        self.short_description = Primitive.from_proto(response.short_description)
        self.long_description = Primitive.from_proto(response.long_description)
        self.related_url = NoteRelatedUrlArray.from_proto(response.related_url)
        self.expiration_time = Primitive.from_proto(response.expiration_time)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.image = NoteImage.from_proto(response.image)
        self.package = NotePackage.from_proto(response.package)
        self.discovery = NoteDiscovery.from_proto(response.discovery)
        self.deployment = NoteDeployment.from_proto(response.deployment)
        self.attestation = NoteAttestation.from_proto(response.attestation)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = note_pb2_grpc.ContaineranalysisAlphaNoteServiceStub(channel.Channel())
        request = note_pb2.DeleteContaineranalysisAlphaNoteRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.short_description):
            request.resource.short_description = Primitive.to_proto(
                self.short_description
            )

        if Primitive.to_proto(self.long_description):
            request.resource.long_description = Primitive.to_proto(
                self.long_description
            )

        if NoteRelatedUrlArray.to_proto(self.related_url):
            request.resource.related_url.extend(
                NoteRelatedUrlArray.to_proto(self.related_url)
            )
        if Primitive.to_proto(self.expiration_time):
            request.resource.expiration_time = Primitive.to_proto(self.expiration_time)

        if NoteImage.to_proto(self.image):
            request.resource.image.CopyFrom(NoteImage.to_proto(self.image))
        else:
            request.resource.ClearField("image")
        if NotePackage.to_proto(self.package):
            request.resource.package.CopyFrom(NotePackage.to_proto(self.package))
        else:
            request.resource.ClearField("package")
        if NoteDiscovery.to_proto(self.discovery):
            request.resource.discovery.CopyFrom(NoteDiscovery.to_proto(self.discovery))
        else:
            request.resource.ClearField("discovery")
        if NoteDeployment.to_proto(self.deployment):
            request.resource.deployment.CopyFrom(
                NoteDeployment.to_proto(self.deployment)
            )
        else:
            request.resource.ClearField("deployment")
        if NoteAttestation.to_proto(self.attestation):
            request.resource.attestation.CopyFrom(
                NoteAttestation.to_proto(self.attestation)
            )
        else:
            request.resource.ClearField("attestation")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteContaineranalysisAlphaNote(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = note_pb2_grpc.ContaineranalysisAlphaNoteServiceStub(channel.Channel())
        request = note_pb2.ListContaineranalysisAlphaNoteRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListContaineranalysisAlphaNote(request).items

    def to_proto(self):
        resource = note_pb2.ContaineranalysisAlphaNote()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.short_description):
            resource.short_description = Primitive.to_proto(self.short_description)
        if Primitive.to_proto(self.long_description):
            resource.long_description = Primitive.to_proto(self.long_description)
        if NoteRelatedUrlArray.to_proto(self.related_url):
            resource.related_url.extend(NoteRelatedUrlArray.to_proto(self.related_url))
        if Primitive.to_proto(self.expiration_time):
            resource.expiration_time = Primitive.to_proto(self.expiration_time)
        if NoteImage.to_proto(self.image):
            resource.image.CopyFrom(NoteImage.to_proto(self.image))
        else:
            resource.ClearField("image")
        if NotePackage.to_proto(self.package):
            resource.package.CopyFrom(NotePackage.to_proto(self.package))
        else:
            resource.ClearField("package")
        if NoteDiscovery.to_proto(self.discovery):
            resource.discovery.CopyFrom(NoteDiscovery.to_proto(self.discovery))
        else:
            resource.ClearField("discovery")
        if NoteDeployment.to_proto(self.deployment):
            resource.deployment.CopyFrom(NoteDeployment.to_proto(self.deployment))
        else:
            resource.ClearField("deployment")
        if NoteAttestation.to_proto(self.attestation):
            resource.attestation.CopyFrom(NoteAttestation.to_proto(self.attestation))
        else:
            resource.ClearField("attestation")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class NoteRelatedUrl(object):
    def __init__(self, url: str = None, label: str = None):
        self.url = url
        self.label = label

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisAlphaNoteRelatedUrl()
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        if Primitive.to_proto(resource.label):
            res.label = Primitive.to_proto(resource.label)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteRelatedUrl(
            url=Primitive.from_proto(resource.url),
            label=Primitive.from_proto(resource.label),
        )


class NoteRelatedUrlArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteRelatedUrl.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteRelatedUrl.from_proto(i) for i in resources]


class NoteImage(object):
    def __init__(self, resource_url: str = None, fingerprint: dict = None):
        self.resource_url = resource_url
        self.fingerprint = fingerprint

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisAlphaNoteImage()
        if Primitive.to_proto(resource.resource_url):
            res.resource_url = Primitive.to_proto(resource.resource_url)
        if NoteImageFingerprint.to_proto(resource.fingerprint):
            res.fingerprint.CopyFrom(
                NoteImageFingerprint.to_proto(resource.fingerprint)
            )
        else:
            res.ClearField("fingerprint")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteImage(
            resource_url=Primitive.from_proto(resource.resource_url),
            fingerprint=NoteImageFingerprint.from_proto(resource.fingerprint),
        )


class NoteImageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteImage.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteImage.from_proto(i) for i in resources]


class NoteImageFingerprint(object):
    def __init__(self, v1_name: str = None, v2_blob: list = None, v2_name: str = None):
        self.v1_name = v1_name
        self.v2_blob = v2_blob
        self.v2_name = v2_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisAlphaNoteImageFingerprint()
        if Primitive.to_proto(resource.v1_name):
            res.v1_name = Primitive.to_proto(resource.v1_name)
        if Primitive.to_proto(resource.v2_blob):
            res.v2_blob.extend(Primitive.to_proto(resource.v2_blob))
        if Primitive.to_proto(resource.v2_name):
            res.v2_name = Primitive.to_proto(resource.v2_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteImageFingerprint(
            v1_name=Primitive.from_proto(resource.v1_name),
            v2_blob=Primitive.from_proto(resource.v2_blob),
            v2_name=Primitive.from_proto(resource.v2_name),
        )


class NoteImageFingerprintArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteImageFingerprint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteImageFingerprint.from_proto(i) for i in resources]


class NotePackage(object):
    def __init__(self, name: str = None, distribution: list = None):
        self.name = name
        self.distribution = distribution

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisAlphaNotePackage()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if NotePackageDistributionArray.to_proto(resource.distribution):
            res.distribution.extend(
                NotePackageDistributionArray.to_proto(resource.distribution)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NotePackage(
            name=Primitive.from_proto(resource.name),
            distribution=NotePackageDistributionArray.from_proto(resource.distribution),
        )


class NotePackageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NotePackage.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NotePackage.from_proto(i) for i in resources]


class NotePackageDistribution(object):
    def __init__(
        self,
        cpe_uri: str = None,
        architecture: str = None,
        latest_version: dict = None,
        maintainer: str = None,
        url: str = None,
        description: str = None,
    ):
        self.cpe_uri = cpe_uri
        self.architecture = architecture
        self.latest_version = latest_version
        self.maintainer = maintainer
        self.url = url
        self.description = description

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisAlphaNotePackageDistribution()
        if Primitive.to_proto(resource.cpe_uri):
            res.cpe_uri = Primitive.to_proto(resource.cpe_uri)
        if NotePackageDistributionArchitectureEnum.to_proto(resource.architecture):
            res.architecture = NotePackageDistributionArchitectureEnum.to_proto(
                resource.architecture
            )
        if NotePackageDistributionLatestVersion.to_proto(resource.latest_version):
            res.latest_version.CopyFrom(
                NotePackageDistributionLatestVersion.to_proto(resource.latest_version)
            )
        else:
            res.ClearField("latest_version")
        if Primitive.to_proto(resource.maintainer):
            res.maintainer = Primitive.to_proto(resource.maintainer)
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NotePackageDistribution(
            cpe_uri=Primitive.from_proto(resource.cpe_uri),
            architecture=NotePackageDistributionArchitectureEnum.from_proto(
                resource.architecture
            ),
            latest_version=NotePackageDistributionLatestVersion.from_proto(
                resource.latest_version
            ),
            maintainer=Primitive.from_proto(resource.maintainer),
            url=Primitive.from_proto(resource.url),
            description=Primitive.from_proto(resource.description),
        )


class NotePackageDistributionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NotePackageDistribution.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NotePackageDistribution.from_proto(i) for i in resources]


class NotePackageDistributionLatestVersion(object):
    def __init__(
        self,
        epoch: int = None,
        name: str = None,
        revision: str = None,
        kind: str = None,
        full_name: str = None,
    ):
        self.epoch = epoch
        self.name = name
        self.revision = revision
        self.kind = kind
        self.full_name = full_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisAlphaNotePackageDistributionLatestVersion()
        if Primitive.to_proto(resource.epoch):
            res.epoch = Primitive.to_proto(resource.epoch)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.revision):
            res.revision = Primitive.to_proto(resource.revision)
        if NotePackageDistributionLatestVersionKindEnum.to_proto(resource.kind):
            res.kind = NotePackageDistributionLatestVersionKindEnum.to_proto(
                resource.kind
            )
        if Primitive.to_proto(resource.full_name):
            res.full_name = Primitive.to_proto(resource.full_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NotePackageDistributionLatestVersion(
            epoch=Primitive.from_proto(resource.epoch),
            name=Primitive.from_proto(resource.name),
            revision=Primitive.from_proto(resource.revision),
            kind=NotePackageDistributionLatestVersionKindEnum.from_proto(resource.kind),
            full_name=Primitive.from_proto(resource.full_name),
        )


class NotePackageDistributionLatestVersionArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NotePackageDistributionLatestVersion.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NotePackageDistributionLatestVersion.from_proto(i) for i in resources]


class NoteDiscovery(object):
    def __init__(self, analysis_kind: str = None):
        self.analysis_kind = analysis_kind

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisAlphaNoteDiscovery()
        if NoteDiscoveryAnalysisKindEnum.to_proto(resource.analysis_kind):
            res.analysis_kind = NoteDiscoveryAnalysisKindEnum.to_proto(
                resource.analysis_kind
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteDiscovery(
            analysis_kind=NoteDiscoveryAnalysisKindEnum.from_proto(
                resource.analysis_kind
            ),
        )


class NoteDiscoveryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteDiscovery.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteDiscovery.from_proto(i) for i in resources]


class NoteDeployment(object):
    def __init__(self, resource_uri: list = None):
        self.resource_uri = resource_uri

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisAlphaNoteDeployment()
        if Primitive.to_proto(resource.resource_uri):
            res.resource_uri.extend(Primitive.to_proto(resource.resource_uri))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteDeployment(
            resource_uri=Primitive.from_proto(resource.resource_uri),
        )


class NoteDeploymentArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteDeployment.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteDeployment.from_proto(i) for i in resources]


class NoteAttestation(object):
    def __init__(self, hint: dict = None):
        self.hint = hint

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisAlphaNoteAttestation()
        if NoteAttestationHint.to_proto(resource.hint):
            res.hint.CopyFrom(NoteAttestationHint.to_proto(resource.hint))
        else:
            res.ClearField("hint")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteAttestation(
            hint=NoteAttestationHint.from_proto(resource.hint),
        )


class NoteAttestationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteAttestation.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteAttestation.from_proto(i) for i in resources]


class NoteAttestationHint(object):
    def __init__(self, human_readable_name: str = None):
        self.human_readable_name = human_readable_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = note_pb2.ContaineranalysisAlphaNoteAttestationHint()
        if Primitive.to_proto(resource.human_readable_name):
            res.human_readable_name = Primitive.to_proto(resource.human_readable_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return NoteAttestationHint(
            human_readable_name=Primitive.from_proto(resource.human_readable_name),
        )


class NoteAttestationHintArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [NoteAttestationHint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [NoteAttestationHint.from_proto(i) for i in resources]


class NotePackageDistributionArchitectureEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum.Value(
            "ContaineranalysisAlphaNotePackageDistributionArchitectureEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            note_pb2.ContaineranalysisAlphaNotePackageDistributionArchitectureEnum.Name(
                resource
            )[len("ContaineranalysisAlphaNotePackageDistributionArchitectureEnum") :]
        )


class NotePackageDistributionLatestVersionKindEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum.Value(
            "ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum.Name(
            resource
        )[
            len("ContaineranalysisAlphaNotePackageDistributionLatestVersionKindEnum") :
        ]


class NoteDiscoveryAnalysisKindEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum.Value(
            "ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return note_pb2.ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum.Name(
            resource
        )[len("ContaineranalysisAlphaNoteDiscoveryAnalysisKindEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
