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
from google3.cloud.graphite.mmv2.services.google.compute import ssl_certificate_pb2
from google3.cloud.graphite.mmv2.services.google.compute import ssl_certificate_pb2_grpc

from typing import List


class SslCertificate(object):
    def __init__(
        self,
        id: int = None,
        name: str = None,
        description: str = None,
        self_link: str = None,
        self_managed: dict = None,
        type: str = None,
        subject_alternative_names: list = None,
        expire_time: str = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.description = description
        self.self_managed = self_managed
        self.type = type
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = ssl_certificate_pb2_grpc.ComputeSslCertificateServiceStub(
            channel.Channel()
        )
        request = ssl_certificate_pb2.ApplyComputeSslCertificateRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if SslCertificateSelfManaged.to_proto(self.self_managed):
            request.resource.self_managed.CopyFrom(
                SslCertificateSelfManaged.to_proto(self.self_managed)
            )
        else:
            request.resource.ClearField("self_managed")
        if SslCertificateTypeEnum.to_proto(self.type):
            request.resource.type = SslCertificateTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyComputeSslCertificate(request)
        self.id = Primitive.from_proto(response.id)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.self_link = Primitive.from_proto(response.self_link)
        self.self_managed = SslCertificateSelfManaged.from_proto(response.self_managed)
        self.type = SslCertificateTypeEnum.from_proto(response.type)
        self.subject_alternative_names = Primitive.from_proto(
            response.subject_alternative_names
        )
        self.expire_time = Primitive.from_proto(response.expire_time)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = ssl_certificate_pb2_grpc.ComputeSslCertificateServiceStub(
            channel.Channel()
        )
        request = ssl_certificate_pb2.DeleteComputeSslCertificateRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if SslCertificateSelfManaged.to_proto(self.self_managed):
            request.resource.self_managed.CopyFrom(
                SslCertificateSelfManaged.to_proto(self.self_managed)
            )
        else:
            request.resource.ClearField("self_managed")
        if SslCertificateTypeEnum.to_proto(self.type):
            request.resource.type = SslCertificateTypeEnum.to_proto(self.type)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteComputeSslCertificate(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = ssl_certificate_pb2_grpc.ComputeSslCertificateServiceStub(
            channel.Channel()
        )
        request = ssl_certificate_pb2.ListComputeSslCertificateRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListComputeSslCertificate(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = ssl_certificate_pb2.ComputeSslCertificate()
        any_proto.Unpack(res_proto)

        res = SslCertificate()
        res.id = Primitive.from_proto(res_proto.id)
        res.name = Primitive.from_proto(res_proto.name)
        res.description = Primitive.from_proto(res_proto.description)
        res.self_link = Primitive.from_proto(res_proto.self_link)
        res.self_managed = SslCertificateSelfManaged.from_proto(res_proto.self_managed)
        res.type = SslCertificateTypeEnum.from_proto(res_proto.type)
        res.subject_alternative_names = Primitive.from_proto(
            res_proto.subject_alternative_names
        )
        res.expire_time = Primitive.from_proto(res_proto.expire_time)
        res.project = Primitive.from_proto(res_proto.project)
        return res

    def to_proto(self):
        resource = ssl_certificate_pb2.ComputeSslCertificate()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if SslCertificateSelfManaged.to_proto(self.self_managed):
            resource.self_managed.CopyFrom(
                SslCertificateSelfManaged.to_proto(self.self_managed)
            )
        else:
            resource.ClearField("self_managed")
        if SslCertificateTypeEnum.to_proto(self.type):
            resource.type = SslCertificateTypeEnum.to_proto(self.type)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class SslCertificateSelfManaged(object):
    def __init__(self, certificate: str = None, private_key: str = None):
        self.certificate = certificate
        self.private_key = private_key

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = ssl_certificate_pb2.ComputeSslCertificateSelfManaged()
        if Primitive.to_proto(resource.certificate):
            res.certificate = Primitive.to_proto(resource.certificate)
        if Primitive.to_proto(resource.private_key):
            res.private_key = Primitive.to_proto(resource.private_key)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return SslCertificateSelfManaged(
            certificate=resource.certificate, private_key=resource.private_key,
        )


class SslCertificateSelfManagedArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [SslCertificateSelfManaged.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [SslCertificateSelfManaged.from_proto(i) for i in resources]


class SslCertificateTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return ssl_certificate_pb2.ComputeSslCertificateTypeEnum.Value(
            "ComputeSslCertificateTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return ssl_certificate_pb2.ComputeSslCertificateTypeEnum.Name(resource)[
            len("ComputeSslCertificateTypeEnum") :
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
