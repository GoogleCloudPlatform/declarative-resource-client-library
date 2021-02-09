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
from google3.cloud.graphite.mmv2.services.google.appengine import version_pb2
from google3.cloud.graphite.mmv2.services.google.appengine import version_pb2_grpc

from typing import List


class Version(object):
    def __init__(
        self,
        consumer_name: str = None,
        name: str = None,
        automatic_scaling: dict = None,
        basic_scaling: dict = None,
        manual_scaling: dict = None,
        job_scaling: dict = None,
        pool_scaling: dict = None,
        inbound_services: list = None,
        instance_class: str = None,
        network: dict = None,
        zones: list = None,
        resources: dict = None,
        runtime: str = None,
        runtime_channel: str = None,
        threadsafe: bool = None,
        vm: bool = None,
        beta_settings: dict = None,
        env: str = None,
        serving_status: str = None,
        created_by: str = None,
        create_time: str = None,
        disk_usage_bytes: int = None,
        runtime_api_version: str = None,
        runtime_main_executable_path: str = None,
        handlers: list = None,
        error_handlers: list = None,
        libraries: list = None,
        api_config: dict = None,
        env_variables: dict = None,
        build_env_variables: dict = None,
        default_expiration: str = None,
        deployment: dict = None,
        health_check: dict = None,
        readiness_check: dict = None,
        liveness_check: dict = None,
        nobuild_files_regex: str = None,
        version_url: str = None,
        service_auth_spec: dict = None,
        service_cors_spec: dict = None,
        route_hash: str = None,
        entrypoint: dict = None,
        vpc_access_connector: dict = None,
        network_settings: dict = None,
        instance_spec: dict = None,
        app: str = None,
        service: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.consumer_name = consumer_name
        self.name = name
        self.automatic_scaling = automatic_scaling
        self.basic_scaling = basic_scaling
        self.manual_scaling = manual_scaling
        self.job_scaling = job_scaling
        self.pool_scaling = pool_scaling
        self.inbound_services = inbound_services
        self.instance_class = instance_class
        self.network = network
        self.zones = zones
        self.resources = resources
        self.runtime = runtime
        self.runtime_channel = runtime_channel
        self.threadsafe = threadsafe
        self.vm = vm
        self.beta_settings = beta_settings
        self.env = env
        self.serving_status = serving_status
        self.created_by = created_by
        self.create_time = create_time
        self.disk_usage_bytes = disk_usage_bytes
        self.runtime_api_version = runtime_api_version
        self.runtime_main_executable_path = runtime_main_executable_path
        self.handlers = handlers
        self.error_handlers = error_handlers
        self.libraries = libraries
        self.api_config = api_config
        self.env_variables = env_variables
        self.build_env_variables = build_env_variables
        self.default_expiration = default_expiration
        self.deployment = deployment
        self.health_check = health_check
        self.readiness_check = readiness_check
        self.liveness_check = liveness_check
        self.nobuild_files_regex = nobuild_files_regex
        self.version_url = version_url
        self.service_auth_spec = service_auth_spec
        self.service_cors_spec = service_cors_spec
        self.route_hash = route_hash
        self.entrypoint = entrypoint
        self.vpc_access_connector = vpc_access_connector
        self.network_settings = network_settings
        self.instance_spec = instance_spec
        self.app = app
        self.service = service
        self.service_account_file = service_account_file

    def apply(self):
        stub = version_pb2_grpc.AppengineVersionServiceStub(channel.Channel())
        request = version_pb2.ApplyAppengineVersionRequest()
        if Primitive.to_proto(self.consumer_name):
            request.resource.consumer_name = Primitive.to_proto(self.consumer_name)

        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if VersionAutomaticScaling.to_proto(self.automatic_scaling):
            request.resource.automatic_scaling.CopyFrom(
                VersionAutomaticScaling.to_proto(self.automatic_scaling)
            )
        else:
            request.resource.ClearField("automatic_scaling")
        if VersionBasicScaling.to_proto(self.basic_scaling):
            request.resource.basic_scaling.CopyFrom(
                VersionBasicScaling.to_proto(self.basic_scaling)
            )
        else:
            request.resource.ClearField("basic_scaling")
        if VersionManualScaling.to_proto(self.manual_scaling):
            request.resource.manual_scaling.CopyFrom(
                VersionManualScaling.to_proto(self.manual_scaling)
            )
        else:
            request.resource.ClearField("manual_scaling")
        if VersionJobScaling.to_proto(self.job_scaling):
            request.resource.job_scaling.CopyFrom(
                VersionJobScaling.to_proto(self.job_scaling)
            )
        else:
            request.resource.ClearField("job_scaling")
        if VersionPoolScaling.to_proto(self.pool_scaling):
            request.resource.pool_scaling.CopyFrom(
                VersionPoolScaling.to_proto(self.pool_scaling)
            )
        else:
            request.resource.ClearField("pool_scaling")
        if VersionInboundServicesEnumArray.to_proto(self.inbound_services):
            request.resource.inbound_services.extend(
                VersionInboundServicesEnumArray.to_proto(self.inbound_services)
            )
        if Primitive.to_proto(self.instance_class):
            request.resource.instance_class = Primitive.to_proto(self.instance_class)

        if VersionNetwork.to_proto(self.network):
            request.resource.network.CopyFrom(VersionNetwork.to_proto(self.network))
        else:
            request.resource.ClearField("network")
        if Primitive.to_proto(self.zones):
            request.resource.zones.extend(Primitive.to_proto(self.zones))
        if VersionResources.to_proto(self.resources):
            request.resource.resources.CopyFrom(
                VersionResources.to_proto(self.resources)
            )
        else:
            request.resource.ClearField("resources")
        if Primitive.to_proto(self.runtime):
            request.resource.runtime = Primitive.to_proto(self.runtime)

        if Primitive.to_proto(self.runtime_channel):
            request.resource.runtime_channel = Primitive.to_proto(self.runtime_channel)

        if Primitive.to_proto(self.threadsafe):
            request.resource.threadsafe = Primitive.to_proto(self.threadsafe)

        if Primitive.to_proto(self.vm):
            request.resource.vm = Primitive.to_proto(self.vm)

        if Primitive.to_proto(self.beta_settings):
            request.resource.beta_settings = Primitive.to_proto(self.beta_settings)

        if Primitive.to_proto(self.env):
            request.resource.env = Primitive.to_proto(self.env)

        if VersionServingStatusEnum.to_proto(self.serving_status):
            request.resource.serving_status = VersionServingStatusEnum.to_proto(
                self.serving_status
            )

        if Primitive.to_proto(self.created_by):
            request.resource.created_by = Primitive.to_proto(self.created_by)

        if Primitive.to_proto(self.create_time):
            request.resource.create_time = Primitive.to_proto(self.create_time)

        if Primitive.to_proto(self.disk_usage_bytes):
            request.resource.disk_usage_bytes = Primitive.to_proto(
                self.disk_usage_bytes
            )

        if Primitive.to_proto(self.runtime_api_version):
            request.resource.runtime_api_version = Primitive.to_proto(
                self.runtime_api_version
            )

        if Primitive.to_proto(self.runtime_main_executable_path):
            request.resource.runtime_main_executable_path = Primitive.to_proto(
                self.runtime_main_executable_path
            )

        if VersionHandlersArray.to_proto(self.handlers):
            request.resource.handlers.extend(
                VersionHandlersArray.to_proto(self.handlers)
            )
        if VersionErrorHandlersArray.to_proto(self.error_handlers):
            request.resource.error_handlers.extend(
                VersionErrorHandlersArray.to_proto(self.error_handlers)
            )
        if VersionLibrariesArray.to_proto(self.libraries):
            request.resource.libraries.extend(
                VersionLibrariesArray.to_proto(self.libraries)
            )
        if VersionApiConfig.to_proto(self.api_config):
            request.resource.api_config.CopyFrom(
                VersionApiConfig.to_proto(self.api_config)
            )
        else:
            request.resource.ClearField("api_config")
        if Primitive.to_proto(self.env_variables):
            request.resource.env_variables = Primitive.to_proto(self.env_variables)

        if Primitive.to_proto(self.build_env_variables):
            request.resource.build_env_variables = Primitive.to_proto(
                self.build_env_variables
            )

        if Primitive.to_proto(self.default_expiration):
            request.resource.default_expiration = Primitive.to_proto(
                self.default_expiration
            )

        if VersionDeployment.to_proto(self.deployment):
            request.resource.deployment.CopyFrom(
                VersionDeployment.to_proto(self.deployment)
            )
        else:
            request.resource.ClearField("deployment")
        if VersionHealthCheck.to_proto(self.health_check):
            request.resource.health_check.CopyFrom(
                VersionHealthCheck.to_proto(self.health_check)
            )
        else:
            request.resource.ClearField("health_check")
        if VersionReadinessCheck.to_proto(self.readiness_check):
            request.resource.readiness_check.CopyFrom(
                VersionReadinessCheck.to_proto(self.readiness_check)
            )
        else:
            request.resource.ClearField("readiness_check")
        if VersionLivenessCheck.to_proto(self.liveness_check):
            request.resource.liveness_check.CopyFrom(
                VersionLivenessCheck.to_proto(self.liveness_check)
            )
        else:
            request.resource.ClearField("liveness_check")
        if Primitive.to_proto(self.nobuild_files_regex):
            request.resource.nobuild_files_regex = Primitive.to_proto(
                self.nobuild_files_regex
            )

        if Primitive.to_proto(self.version_url):
            request.resource.version_url = Primitive.to_proto(self.version_url)

        if VersionServiceAuthSpec.to_proto(self.service_auth_spec):
            request.resource.service_auth_spec.CopyFrom(
                VersionServiceAuthSpec.to_proto(self.service_auth_spec)
            )
        else:
            request.resource.ClearField("service_auth_spec")
        if VersionServiceCorsSpec.to_proto(self.service_cors_spec):
            request.resource.service_cors_spec.CopyFrom(
                VersionServiceCorsSpec.to_proto(self.service_cors_spec)
            )
        else:
            request.resource.ClearField("service_cors_spec")
        if Primitive.to_proto(self.route_hash):
            request.resource.route_hash = Primitive.to_proto(self.route_hash)

        if VersionEntrypoint.to_proto(self.entrypoint):
            request.resource.entrypoint.CopyFrom(
                VersionEntrypoint.to_proto(self.entrypoint)
            )
        else:
            request.resource.ClearField("entrypoint")
        if VersionVPCAccessConnector.to_proto(self.vpc_access_connector):
            request.resource.vpc_access_connector.CopyFrom(
                VersionVPCAccessConnector.to_proto(self.vpc_access_connector)
            )
        else:
            request.resource.ClearField("vpc_access_connector")
        if VersionNetworkSettings.to_proto(self.network_settings):
            request.resource.network_settings.CopyFrom(
                VersionNetworkSettings.to_proto(self.network_settings)
            )
        else:
            request.resource.ClearField("network_settings")
        if VersionInstanceSpec.to_proto(self.instance_spec):
            request.resource.instance_spec.CopyFrom(
                VersionInstanceSpec.to_proto(self.instance_spec)
            )
        else:
            request.resource.ClearField("instance_spec")
        if Primitive.to_proto(self.app):
            request.resource.app = Primitive.to_proto(self.app)

        if Primitive.to_proto(self.service):
            request.resource.service = Primitive.to_proto(self.service)

        request.service_account_file = self.service_account_file

        response = stub.ApplyAppengineVersion(request)
        self.consumer_name = Primitive.from_proto(response.consumer_name)
        self.name = Primitive.from_proto(response.name)
        self.automatic_scaling = VersionAutomaticScaling.from_proto(
            response.automatic_scaling
        )
        self.basic_scaling = VersionBasicScaling.from_proto(response.basic_scaling)
        self.manual_scaling = VersionManualScaling.from_proto(response.manual_scaling)
        self.job_scaling = VersionJobScaling.from_proto(response.job_scaling)
        self.pool_scaling = VersionPoolScaling.from_proto(response.pool_scaling)
        self.inbound_services = VersionInboundServicesEnumArray.from_proto(
            response.inbound_services
        )
        self.instance_class = Primitive.from_proto(response.instance_class)
        self.network = VersionNetwork.from_proto(response.network)
        self.zones = Primitive.from_proto(response.zones)
        self.resources = VersionResources.from_proto(response.resources)
        self.runtime = Primitive.from_proto(response.runtime)
        self.runtime_channel = Primitive.from_proto(response.runtime_channel)
        self.threadsafe = Primitive.from_proto(response.threadsafe)
        self.vm = Primitive.from_proto(response.vm)
        self.beta_settings = Primitive.from_proto(response.beta_settings)
        self.env = Primitive.from_proto(response.env)
        self.serving_status = VersionServingStatusEnum.from_proto(
            response.serving_status
        )
        self.created_by = Primitive.from_proto(response.created_by)
        self.create_time = Primitive.from_proto(response.create_time)
        self.disk_usage_bytes = Primitive.from_proto(response.disk_usage_bytes)
        self.runtime_api_version = Primitive.from_proto(response.runtime_api_version)
        self.runtime_main_executable_path = Primitive.from_proto(
            response.runtime_main_executable_path
        )
        self.handlers = VersionHandlersArray.from_proto(response.handlers)
        self.error_handlers = VersionErrorHandlersArray.from_proto(
            response.error_handlers
        )
        self.libraries = VersionLibrariesArray.from_proto(response.libraries)
        self.api_config = VersionApiConfig.from_proto(response.api_config)
        self.env_variables = Primitive.from_proto(response.env_variables)
        self.build_env_variables = Primitive.from_proto(response.build_env_variables)
        self.default_expiration = Primitive.from_proto(response.default_expiration)
        self.deployment = VersionDeployment.from_proto(response.deployment)
        self.health_check = VersionHealthCheck.from_proto(response.health_check)
        self.readiness_check = VersionReadinessCheck.from_proto(
            response.readiness_check
        )
        self.liveness_check = VersionLivenessCheck.from_proto(response.liveness_check)
        self.nobuild_files_regex = Primitive.from_proto(response.nobuild_files_regex)
        self.version_url = Primitive.from_proto(response.version_url)
        self.service_auth_spec = VersionServiceAuthSpec.from_proto(
            response.service_auth_spec
        )
        self.service_cors_spec = VersionServiceCorsSpec.from_proto(
            response.service_cors_spec
        )
        self.route_hash = Primitive.from_proto(response.route_hash)
        self.entrypoint = VersionEntrypoint.from_proto(response.entrypoint)
        self.vpc_access_connector = VersionVPCAccessConnector.from_proto(
            response.vpc_access_connector
        )
        self.network_settings = VersionNetworkSettings.from_proto(
            response.network_settings
        )
        self.instance_spec = VersionInstanceSpec.from_proto(response.instance_spec)
        self.app = Primitive.from_proto(response.app)
        self.service = Primitive.from_proto(response.service)

    @classmethod
    def delete(self, app, service, name, service_account_file=""):
        stub = version_pb2_grpc.AppengineVersionServiceStub(channel.Channel())
        request = version_pb2.DeleteAppengineVersionRequest()
        request.service_account_file = service_account_file
        request.App = app

        request.Service = service

        request.Name = name

        response = stub.DeleteAppengineVersion(request)

    @classmethod
    def list(self, app, service, service_account_file=""):
        stub = version_pb2_grpc.AppengineVersionServiceStub(channel.Channel())
        request = version_pb2.ListAppengineVersionRequest()
        request.service_account_file = service_account_file
        request.App = app

        request.Service = service

        return stub.ListAppengineVersion(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = version_pb2.AppengineVersion()
        any_proto.Unpack(res_proto)

        res = Version()
        res.consumer_name = Primitive.from_proto(res_proto.consumer_name)
        res.name = Primitive.from_proto(res_proto.name)
        res.automatic_scaling = VersionAutomaticScaling.from_proto(
            res_proto.automatic_scaling
        )
        res.basic_scaling = VersionBasicScaling.from_proto(res_proto.basic_scaling)
        res.manual_scaling = VersionManualScaling.from_proto(res_proto.manual_scaling)
        res.job_scaling = VersionJobScaling.from_proto(res_proto.job_scaling)
        res.pool_scaling = VersionPoolScaling.from_proto(res_proto.pool_scaling)
        res.inbound_services = VersionInboundServicesEnumArray.from_proto(
            res_proto.inbound_services
        )
        res.instance_class = Primitive.from_proto(res_proto.instance_class)
        res.network = VersionNetwork.from_proto(res_proto.network)
        res.zones = Primitive.from_proto(res_proto.zones)
        res.resources = VersionResources.from_proto(res_proto.resources)
        res.runtime = Primitive.from_proto(res_proto.runtime)
        res.runtime_channel = Primitive.from_proto(res_proto.runtime_channel)
        res.threadsafe = Primitive.from_proto(res_proto.threadsafe)
        res.vm = Primitive.from_proto(res_proto.vm)
        res.beta_settings = Primitive.from_proto(res_proto.beta_settings)
        res.env = Primitive.from_proto(res_proto.env)
        res.serving_status = VersionServingStatusEnum.from_proto(
            res_proto.serving_status
        )
        res.created_by = Primitive.from_proto(res_proto.created_by)
        res.create_time = Primitive.from_proto(res_proto.create_time)
        res.disk_usage_bytes = Primitive.from_proto(res_proto.disk_usage_bytes)
        res.runtime_api_version = Primitive.from_proto(res_proto.runtime_api_version)
        res.runtime_main_executable_path = Primitive.from_proto(
            res_proto.runtime_main_executable_path
        )
        res.handlers = VersionHandlersArray.from_proto(res_proto.handlers)
        res.error_handlers = VersionErrorHandlersArray.from_proto(
            res_proto.error_handlers
        )
        res.libraries = VersionLibrariesArray.from_proto(res_proto.libraries)
        res.api_config = VersionApiConfig.from_proto(res_proto.api_config)
        res.env_variables = Primitive.from_proto(res_proto.env_variables)
        res.build_env_variables = Primitive.from_proto(res_proto.build_env_variables)
        res.default_expiration = Primitive.from_proto(res_proto.default_expiration)
        res.deployment = VersionDeployment.from_proto(res_proto.deployment)
        res.health_check = VersionHealthCheck.from_proto(res_proto.health_check)
        res.readiness_check = VersionReadinessCheck.from_proto(
            res_proto.readiness_check
        )
        res.liveness_check = VersionLivenessCheck.from_proto(res_proto.liveness_check)
        res.nobuild_files_regex = Primitive.from_proto(res_proto.nobuild_files_regex)
        res.version_url = Primitive.from_proto(res_proto.version_url)
        res.service_auth_spec = VersionServiceAuthSpec.from_proto(
            res_proto.service_auth_spec
        )
        res.service_cors_spec = VersionServiceCorsSpec.from_proto(
            res_proto.service_cors_spec
        )
        res.route_hash = Primitive.from_proto(res_proto.route_hash)
        res.entrypoint = VersionEntrypoint.from_proto(res_proto.entrypoint)
        res.vpc_access_connector = VersionVPCAccessConnector.from_proto(
            res_proto.vpc_access_connector
        )
        res.network_settings = VersionNetworkSettings.from_proto(
            res_proto.network_settings
        )
        res.instance_spec = VersionInstanceSpec.from_proto(res_proto.instance_spec)
        res.app = Primitive.from_proto(res_proto.app)
        res.service = Primitive.from_proto(res_proto.service)
        return res


class VersionAutomaticScaling(object):
    def __init__(
        self,
        cool_down_period: str = None,
        cpu_utilization: dict = None,
        max_concurrent_requests: int = None,
        max_idle_instances: int = None,
        max_total_instances: int = None,
        max_pending_latency: str = None,
        min_idle_instances: int = None,
        min_total_instances: int = None,
        min_pending_latency: str = None,
        request_utilization: dict = None,
        disk_utilization: dict = None,
        network_utilization: dict = None,
        standard_scheduler_settings: dict = None,
    ):
        self.cool_down_period = cool_down_period
        self.cpu_utilization = cpu_utilization
        self.max_concurrent_requests = max_concurrent_requests
        self.max_idle_instances = max_idle_instances
        self.max_total_instances = max_total_instances
        self.max_pending_latency = max_pending_latency
        self.min_idle_instances = min_idle_instances
        self.min_total_instances = min_total_instances
        self.min_pending_latency = min_pending_latency
        self.request_utilization = request_utilization
        self.disk_utilization = disk_utilization
        self.network_utilization = network_utilization
        self.standard_scheduler_settings = standard_scheduler_settings

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionAutomaticScaling()
        if Primitive.to_proto(resource.cool_down_period):
            res.cool_down_period = Primitive.to_proto(resource.cool_down_period)
        if VersionAutomaticScalingCpuUtilization.to_proto(resource.cpu_utilization):
            res.cpu_utilization.CopyFrom(
                VersionAutomaticScalingCpuUtilization.to_proto(resource.cpu_utilization)
            )
        else:
            res.ClearField("cpu_utilization")
        if Primitive.to_proto(resource.max_concurrent_requests):
            res.max_concurrent_requests = Primitive.to_proto(
                resource.max_concurrent_requests
            )
        if Primitive.to_proto(resource.max_idle_instances):
            res.max_idle_instances = Primitive.to_proto(resource.max_idle_instances)
        if Primitive.to_proto(resource.max_total_instances):
            res.max_total_instances = Primitive.to_proto(resource.max_total_instances)
        if Primitive.to_proto(resource.max_pending_latency):
            res.max_pending_latency = Primitive.to_proto(resource.max_pending_latency)
        if Primitive.to_proto(resource.min_idle_instances):
            res.min_idle_instances = Primitive.to_proto(resource.min_idle_instances)
        if Primitive.to_proto(resource.min_total_instances):
            res.min_total_instances = Primitive.to_proto(resource.min_total_instances)
        if Primitive.to_proto(resource.min_pending_latency):
            res.min_pending_latency = Primitive.to_proto(resource.min_pending_latency)
        if VersionAutomaticScalingRequestUtilization.to_proto(
            resource.request_utilization
        ):
            res.request_utilization.CopyFrom(
                VersionAutomaticScalingRequestUtilization.to_proto(
                    resource.request_utilization
                )
            )
        else:
            res.ClearField("request_utilization")
        if VersionAutomaticScalingDiskUtilization.to_proto(resource.disk_utilization):
            res.disk_utilization.CopyFrom(
                VersionAutomaticScalingDiskUtilization.to_proto(
                    resource.disk_utilization
                )
            )
        else:
            res.ClearField("disk_utilization")
        if VersionAutomaticScalingNetworkUtilization.to_proto(
            resource.network_utilization
        ):
            res.network_utilization.CopyFrom(
                VersionAutomaticScalingNetworkUtilization.to_proto(
                    resource.network_utilization
                )
            )
        else:
            res.ClearField("network_utilization")
        if VersionAutomaticScalingStandardSchedulerSettings.to_proto(
            resource.standard_scheduler_settings
        ):
            res.standard_scheduler_settings.CopyFrom(
                VersionAutomaticScalingStandardSchedulerSettings.to_proto(
                    resource.standard_scheduler_settings
                )
            )
        else:
            res.ClearField("standard_scheduler_settings")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionAutomaticScaling(
            cool_down_period=resource.cool_down_period,
            cpu_utilization=resource.cpu_utilization,
            max_concurrent_requests=resource.max_concurrent_requests,
            max_idle_instances=resource.max_idle_instances,
            max_total_instances=resource.max_total_instances,
            max_pending_latency=resource.max_pending_latency,
            min_idle_instances=resource.min_idle_instances,
            min_total_instances=resource.min_total_instances,
            min_pending_latency=resource.min_pending_latency,
            request_utilization=resource.request_utilization,
            disk_utilization=resource.disk_utilization,
            network_utilization=resource.network_utilization,
            standard_scheduler_settings=resource.standard_scheduler_settings,
        )


class VersionAutomaticScalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionAutomaticScaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionAutomaticScaling.from_proto(i) for i in resources]


class VersionAutomaticScalingCpuUtilization(object):
    def __init__(
        self, aggregation_window_length: str = None, target_utilization: float = None
    ):
        self.aggregation_window_length = aggregation_window_length
        self.target_utilization = target_utilization

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionAutomaticScalingCpuUtilization()
        if Primitive.to_proto(resource.aggregation_window_length):
            res.aggregation_window_length = Primitive.to_proto(
                resource.aggregation_window_length
            )
        if Primitive.to_proto(resource.target_utilization):
            res.target_utilization = Primitive.to_proto(resource.target_utilization)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionAutomaticScalingCpuUtilization(
            aggregation_window_length=resource.aggregation_window_length,
            target_utilization=resource.target_utilization,
        )


class VersionAutomaticScalingCpuUtilizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionAutomaticScalingCpuUtilization.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionAutomaticScalingCpuUtilization.from_proto(i) for i in resources]


class VersionAutomaticScalingRequestUtilization(object):
    def __init__(
        self,
        target_request_count_per_second: int = None,
        target_concurrent_requests: int = None,
    ):
        self.target_request_count_per_second = target_request_count_per_second
        self.target_concurrent_requests = target_concurrent_requests

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionAutomaticScalingRequestUtilization()
        if Primitive.to_proto(resource.target_request_count_per_second):
            res.target_request_count_per_second = Primitive.to_proto(
                resource.target_request_count_per_second
            )
        if Primitive.to_proto(resource.target_concurrent_requests):
            res.target_concurrent_requests = Primitive.to_proto(
                resource.target_concurrent_requests
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionAutomaticScalingRequestUtilization(
            target_request_count_per_second=resource.target_request_count_per_second,
            target_concurrent_requests=resource.target_concurrent_requests,
        )


class VersionAutomaticScalingRequestUtilizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            VersionAutomaticScalingRequestUtilization.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            VersionAutomaticScalingRequestUtilization.from_proto(i) for i in resources
        ]


class VersionAutomaticScalingDiskUtilization(object):
    def __init__(
        self,
        target_write_bytes_per_second: int = None,
        target_write_ops_per_second: int = None,
        target_read_bytes_per_second: int = None,
        target_read_ops_per_second: int = None,
    ):
        self.target_write_bytes_per_second = target_write_bytes_per_second
        self.target_write_ops_per_second = target_write_ops_per_second
        self.target_read_bytes_per_second = target_read_bytes_per_second
        self.target_read_ops_per_second = target_read_ops_per_second

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionAutomaticScalingDiskUtilization()
        if Primitive.to_proto(resource.target_write_bytes_per_second):
            res.target_write_bytes_per_second = Primitive.to_proto(
                resource.target_write_bytes_per_second
            )
        if Primitive.to_proto(resource.target_write_ops_per_second):
            res.target_write_ops_per_second = Primitive.to_proto(
                resource.target_write_ops_per_second
            )
        if Primitive.to_proto(resource.target_read_bytes_per_second):
            res.target_read_bytes_per_second = Primitive.to_proto(
                resource.target_read_bytes_per_second
            )
        if Primitive.to_proto(resource.target_read_ops_per_second):
            res.target_read_ops_per_second = Primitive.to_proto(
                resource.target_read_ops_per_second
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionAutomaticScalingDiskUtilization(
            target_write_bytes_per_second=resource.target_write_bytes_per_second,
            target_write_ops_per_second=resource.target_write_ops_per_second,
            target_read_bytes_per_second=resource.target_read_bytes_per_second,
            target_read_ops_per_second=resource.target_read_ops_per_second,
        )


class VersionAutomaticScalingDiskUtilizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionAutomaticScalingDiskUtilization.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionAutomaticScalingDiskUtilization.from_proto(i) for i in resources]


class VersionAutomaticScalingNetworkUtilization(object):
    def __init__(
        self,
        target_sent_bytes_per_second: int = None,
        target_sent_packets_per_second: int = None,
        target_received_bytes_per_second: int = None,
        target_received_packets_per_second: int = None,
    ):
        self.target_sent_bytes_per_second = target_sent_bytes_per_second
        self.target_sent_packets_per_second = target_sent_packets_per_second
        self.target_received_bytes_per_second = target_received_bytes_per_second
        self.target_received_packets_per_second = target_received_packets_per_second

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionAutomaticScalingNetworkUtilization()
        if Primitive.to_proto(resource.target_sent_bytes_per_second):
            res.target_sent_bytes_per_second = Primitive.to_proto(
                resource.target_sent_bytes_per_second
            )
        if Primitive.to_proto(resource.target_sent_packets_per_second):
            res.target_sent_packets_per_second = Primitive.to_proto(
                resource.target_sent_packets_per_second
            )
        if Primitive.to_proto(resource.target_received_bytes_per_second):
            res.target_received_bytes_per_second = Primitive.to_proto(
                resource.target_received_bytes_per_second
            )
        if Primitive.to_proto(resource.target_received_packets_per_second):
            res.target_received_packets_per_second = Primitive.to_proto(
                resource.target_received_packets_per_second
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionAutomaticScalingNetworkUtilization(
            target_sent_bytes_per_second=resource.target_sent_bytes_per_second,
            target_sent_packets_per_second=resource.target_sent_packets_per_second,
            target_received_bytes_per_second=resource.target_received_bytes_per_second,
            target_received_packets_per_second=resource.target_received_packets_per_second,
        )


class VersionAutomaticScalingNetworkUtilizationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            VersionAutomaticScalingNetworkUtilization.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            VersionAutomaticScalingNetworkUtilization.from_proto(i) for i in resources
        ]


class VersionAutomaticScalingStandardSchedulerSettings(object):
    def __init__(
        self,
        target_cpu_utilization: float = None,
        target_throughput_utilization: float = None,
        min_instances: int = None,
        max_instances: int = None,
    ):
        self.target_cpu_utilization = target_cpu_utilization
        self.target_throughput_utilization = target_throughput_utilization
        self.min_instances = min_instances
        self.max_instances = max_instances

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionAutomaticScalingStandardSchedulerSettings()
        if Primitive.to_proto(resource.target_cpu_utilization):
            res.target_cpu_utilization = Primitive.to_proto(
                resource.target_cpu_utilization
            )
        if Primitive.to_proto(resource.target_throughput_utilization):
            res.target_throughput_utilization = Primitive.to_proto(
                resource.target_throughput_utilization
            )
        if Primitive.to_proto(resource.min_instances):
            res.min_instances = Primitive.to_proto(resource.min_instances)
        if Primitive.to_proto(resource.max_instances):
            res.max_instances = Primitive.to_proto(resource.max_instances)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionAutomaticScalingStandardSchedulerSettings(
            target_cpu_utilization=resource.target_cpu_utilization,
            target_throughput_utilization=resource.target_throughput_utilization,
            min_instances=resource.min_instances,
            max_instances=resource.max_instances,
        )


class VersionAutomaticScalingStandardSchedulerSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            VersionAutomaticScalingStandardSchedulerSettings.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            VersionAutomaticScalingStandardSchedulerSettings.from_proto(i)
            for i in resources
        ]


class VersionBasicScaling(object):
    def __init__(self, idle_timeout: str = None, max_instances: int = None):
        self.idle_timeout = idle_timeout
        self.max_instances = max_instances

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionBasicScaling()
        if Primitive.to_proto(resource.idle_timeout):
            res.idle_timeout = Primitive.to_proto(resource.idle_timeout)
        if Primitive.to_proto(resource.max_instances):
            res.max_instances = Primitive.to_proto(resource.max_instances)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionBasicScaling(
            idle_timeout=resource.idle_timeout, max_instances=resource.max_instances,
        )


class VersionBasicScalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionBasicScaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionBasicScaling.from_proto(i) for i in resources]


class VersionManualScaling(object):
    def __init__(self, instances: int = None):
        self.instances = instances

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionManualScaling()
        if Primitive.to_proto(resource.instances):
            res.instances = Primitive.to_proto(resource.instances)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionManualScaling(instances=resource.instances,)


class VersionManualScalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionManualScaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionManualScaling.from_proto(i) for i in resources]


class VersionJobScaling(object):
    def __init__(
        self,
        completions: int = None,
        parallelism: int = None,
        job_deadline: str = None,
        instance_retries: int = None,
        instance_deadline: str = None,
        instance_termination_window: str = None,
    ):
        self.completions = completions
        self.parallelism = parallelism
        self.job_deadline = job_deadline
        self.instance_retries = instance_retries
        self.instance_deadline = instance_deadline
        self.instance_termination_window = instance_termination_window

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionJobScaling()
        if Primitive.to_proto(resource.completions):
            res.completions = Primitive.to_proto(resource.completions)
        if Primitive.to_proto(resource.parallelism):
            res.parallelism = Primitive.to_proto(resource.parallelism)
        if Primitive.to_proto(resource.job_deadline):
            res.job_deadline = Primitive.to_proto(resource.job_deadline)
        if Primitive.to_proto(resource.instance_retries):
            res.instance_retries = Primitive.to_proto(resource.instance_retries)
        if Primitive.to_proto(resource.instance_deadline):
            res.instance_deadline = Primitive.to_proto(resource.instance_deadline)
        if Primitive.to_proto(resource.instance_termination_window):
            res.instance_termination_window = Primitive.to_proto(
                resource.instance_termination_window
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionJobScaling(
            completions=resource.completions,
            parallelism=resource.parallelism,
            job_deadline=resource.job_deadline,
            instance_retries=resource.instance_retries,
            instance_deadline=resource.instance_deadline,
            instance_termination_window=resource.instance_termination_window,
        )


class VersionJobScalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionJobScaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionJobScaling.from_proto(i) for i in resources]


class VersionPoolScaling(object):
    def __init__(
        self, replicas: int = None, max_unavailable: int = None, max_surge: int = None
    ):
        self.replicas = replicas
        self.max_unavailable = max_unavailable
        self.max_surge = max_surge

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionPoolScaling()
        if Primitive.to_proto(resource.replicas):
            res.replicas = Primitive.to_proto(resource.replicas)
        if Primitive.to_proto(resource.max_unavailable):
            res.max_unavailable = Primitive.to_proto(resource.max_unavailable)
        if Primitive.to_proto(resource.max_surge):
            res.max_surge = Primitive.to_proto(resource.max_surge)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionPoolScaling(
            replicas=resource.replicas,
            max_unavailable=resource.max_unavailable,
            max_surge=resource.max_surge,
        )


class VersionPoolScalingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionPoolScaling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionPoolScaling.from_proto(i) for i in resources]


class VersionNetwork(object):
    def __init__(
        self,
        forwarded_ports: list = None,
        instance_tag: str = None,
        name: str = None,
        subnetwork_name: str = None,
        session_affinity: bool = None,
    ):
        self.forwarded_ports = forwarded_ports
        self.instance_tag = instance_tag
        self.name = name
        self.subnetwork_name = subnetwork_name
        self.session_affinity = session_affinity

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionNetwork()
        if Primitive.to_proto(resource.forwarded_ports):
            res.forwarded_ports.extend(Primitive.to_proto(resource.forwarded_ports))
        if Primitive.to_proto(resource.instance_tag):
            res.instance_tag = Primitive.to_proto(resource.instance_tag)
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.subnetwork_name):
            res.subnetwork_name = Primitive.to_proto(resource.subnetwork_name)
        if Primitive.to_proto(resource.session_affinity):
            res.session_affinity = Primitive.to_proto(resource.session_affinity)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionNetwork(
            forwarded_ports=resource.forwarded_ports,
            instance_tag=resource.instance_tag,
            name=resource.name,
            subnetwork_name=resource.subnetwork_name,
            session_affinity=resource.session_affinity,
        )


class VersionNetworkArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionNetwork.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionNetwork.from_proto(i) for i in resources]


class VersionResources(object):
    def __init__(
        self,
        cpu: float = None,
        disk_gb: float = None,
        memory_gb: float = None,
        volumes: list = None,
        kms_key_reference: str = None,
    ):
        self.cpu = cpu
        self.disk_gb = disk_gb
        self.memory_gb = memory_gb
        self.volumes = volumes
        self.kms_key_reference = kms_key_reference

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionResources()
        if Primitive.to_proto(resource.cpu):
            res.cpu = Primitive.to_proto(resource.cpu)
        if Primitive.to_proto(resource.disk_gb):
            res.disk_gb = Primitive.to_proto(resource.disk_gb)
        if Primitive.to_proto(resource.memory_gb):
            res.memory_gb = Primitive.to_proto(resource.memory_gb)
        if VersionResourcesVolumesArray.to_proto(resource.volumes):
            res.volumes.extend(VersionResourcesVolumesArray.to_proto(resource.volumes))
        if Primitive.to_proto(resource.kms_key_reference):
            res.kms_key_reference = Primitive.to_proto(resource.kms_key_reference)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionResources(
            cpu=resource.cpu,
            disk_gb=resource.disk_gb,
            memory_gb=resource.memory_gb,
            volumes=resource.volumes,
            kms_key_reference=resource.kms_key_reference,
        )


class VersionResourcesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionResources.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionResources.from_proto(i) for i in resources]


class VersionResourcesVolumes(object):
    def __init__(
        self, name: str = None, volume_type: str = None, size_gb: float = None
    ):
        self.name = name
        self.volume_type = volume_type
        self.size_gb = size_gb

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionResourcesVolumes()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.volume_type):
            res.volume_type = Primitive.to_proto(resource.volume_type)
        if Primitive.to_proto(resource.size_gb):
            res.size_gb = Primitive.to_proto(resource.size_gb)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionResourcesVolumes(
            name=resource.name,
            volume_type=resource.volume_type,
            size_gb=resource.size_gb,
        )


class VersionResourcesVolumesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionResourcesVolumes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionResourcesVolumes.from_proto(i) for i in resources]


class VersionHandlers(object):
    def __init__(
        self,
        url_regex: str = None,
        static_files: dict = None,
        script: dict = None,
        api_endpoint: dict = None,
        security_level: str = None,
        login: str = None,
        auth_fail_action: str = None,
        redirect_http_response_code: str = None,
    ):
        self.url_regex = url_regex
        self.static_files = static_files
        self.script = script
        self.api_endpoint = api_endpoint
        self.security_level = security_level
        self.login = login
        self.auth_fail_action = auth_fail_action
        self.redirect_http_response_code = redirect_http_response_code

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionHandlers()
        if Primitive.to_proto(resource.url_regex):
            res.url_regex = Primitive.to_proto(resource.url_regex)
        if VersionHandlersStaticFiles.to_proto(resource.static_files):
            res.static_files.CopyFrom(
                VersionHandlersStaticFiles.to_proto(resource.static_files)
            )
        else:
            res.ClearField("static_files")
        if VersionHandlersScript.to_proto(resource.script):
            res.script.CopyFrom(VersionHandlersScript.to_proto(resource.script))
        else:
            res.ClearField("script")
        if VersionHandlersApiEndpoint.to_proto(resource.api_endpoint):
            res.api_endpoint.CopyFrom(
                VersionHandlersApiEndpoint.to_proto(resource.api_endpoint)
            )
        else:
            res.ClearField("api_endpoint")
        if VersionHandlersSecurityLevelEnum.to_proto(resource.security_level):
            res.security_level = VersionHandlersSecurityLevelEnum.to_proto(
                resource.security_level
            )
        if VersionHandlersLoginEnum.to_proto(resource.login):
            res.login = VersionHandlersLoginEnum.to_proto(resource.login)
        if VersionHandlersAuthFailActionEnum.to_proto(resource.auth_fail_action):
            res.auth_fail_action = VersionHandlersAuthFailActionEnum.to_proto(
                resource.auth_fail_action
            )
        if VersionHandlersRedirectHttpResponseCodeEnum.to_proto(
            resource.redirect_http_response_code
        ):
            res.redirect_http_response_code = VersionHandlersRedirectHttpResponseCodeEnum.to_proto(
                resource.redirect_http_response_code
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionHandlers(
            url_regex=resource.url_regex,
            static_files=resource.static_files,
            script=resource.script,
            api_endpoint=resource.api_endpoint,
            security_level=resource.security_level,
            login=resource.login,
            auth_fail_action=resource.auth_fail_action,
            redirect_http_response_code=resource.redirect_http_response_code,
        )


class VersionHandlersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionHandlers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionHandlers.from_proto(i) for i in resources]


class VersionHandlersStaticFiles(object):
    def __init__(
        self,
        path: str = None,
        upload_path_regex: str = None,
        http_headers: dict = None,
        mime_type: str = None,
        expiration: str = None,
        require_matching_file: bool = None,
        application_readable: bool = None,
    ):
        self.path = path
        self.upload_path_regex = upload_path_regex
        self.http_headers = http_headers
        self.mime_type = mime_type
        self.expiration = expiration
        self.require_matching_file = require_matching_file
        self.application_readable = application_readable

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionHandlersStaticFiles()
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.upload_path_regex):
            res.upload_path_regex = Primitive.to_proto(resource.upload_path_regex)
        if Primitive.to_proto(resource.http_headers):
            res.http_headers = Primitive.to_proto(resource.http_headers)
        if Primitive.to_proto(resource.mime_type):
            res.mime_type = Primitive.to_proto(resource.mime_type)
        if Primitive.to_proto(resource.expiration):
            res.expiration = Primitive.to_proto(resource.expiration)
        if Primitive.to_proto(resource.require_matching_file):
            res.require_matching_file = Primitive.to_proto(
                resource.require_matching_file
            )
        if Primitive.to_proto(resource.application_readable):
            res.application_readable = Primitive.to_proto(resource.application_readable)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionHandlersStaticFiles(
            path=resource.path,
            upload_path_regex=resource.upload_path_regex,
            http_headers=resource.http_headers,
            mime_type=resource.mime_type,
            expiration=resource.expiration,
            require_matching_file=resource.require_matching_file,
            application_readable=resource.application_readable,
        )


class VersionHandlersStaticFilesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionHandlersStaticFiles.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionHandlersStaticFiles.from_proto(i) for i in resources]


class VersionHandlersScript(object):
    def __init__(self, script_path: str = None):
        self.script_path = script_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionHandlersScript()
        if Primitive.to_proto(resource.script_path):
            res.script_path = Primitive.to_proto(resource.script_path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionHandlersScript(script_path=resource.script_path,)


class VersionHandlersScriptArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionHandlersScript.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionHandlersScript.from_proto(i) for i in resources]


class VersionHandlersApiEndpoint(object):
    def __init__(self, script_path: str = None):
        self.script_path = script_path

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionHandlersApiEndpoint()
        if Primitive.to_proto(resource.script_path):
            res.script_path = Primitive.to_proto(resource.script_path)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionHandlersApiEndpoint(script_path=resource.script_path,)


class VersionHandlersApiEndpointArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionHandlersApiEndpoint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionHandlersApiEndpoint.from_proto(i) for i in resources]


class VersionErrorHandlers(object):
    def __init__(
        self, error_code: str = None, static_file: str = None, mime_type: str = None
    ):
        self.error_code = error_code
        self.static_file = static_file
        self.mime_type = mime_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionErrorHandlers()
        if VersionErrorHandlersErrorCodeEnum.to_proto(resource.error_code):
            res.error_code = VersionErrorHandlersErrorCodeEnum.to_proto(
                resource.error_code
            )
        if Primitive.to_proto(resource.static_file):
            res.static_file = Primitive.to_proto(resource.static_file)
        if Primitive.to_proto(resource.mime_type):
            res.mime_type = Primitive.to_proto(resource.mime_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionErrorHandlers(
            error_code=resource.error_code,
            static_file=resource.static_file,
            mime_type=resource.mime_type,
        )


class VersionErrorHandlersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionErrorHandlers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionErrorHandlers.from_proto(i) for i in resources]


class VersionLibraries(object):
    def __init__(self, name: str = None, version: str = None):
        self.name = name
        self.version = version

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionLibraries()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.version):
            res.version = Primitive.to_proto(resource.version)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionLibraries(name=resource.name, version=resource.version,)


class VersionLibrariesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionLibraries.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionLibraries.from_proto(i) for i in resources]


class VersionApiConfig(object):
    def __init__(
        self,
        auth_fail_action: str = None,
        login: str = None,
        script: str = None,
        security_level: str = None,
        url: str = None,
    ):
        self.auth_fail_action = auth_fail_action
        self.login = login
        self.script = script
        self.security_level = security_level
        self.url = url

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionApiConfig()
        if VersionApiConfigAuthFailActionEnum.to_proto(resource.auth_fail_action):
            res.auth_fail_action = VersionApiConfigAuthFailActionEnum.to_proto(
                resource.auth_fail_action
            )
        if VersionApiConfigLoginEnum.to_proto(resource.login):
            res.login = VersionApiConfigLoginEnum.to_proto(resource.login)
        if Primitive.to_proto(resource.script):
            res.script = Primitive.to_proto(resource.script)
        if VersionApiConfigSecurityLevelEnum.to_proto(resource.security_level):
            res.security_level = VersionApiConfigSecurityLevelEnum.to_proto(
                resource.security_level
            )
        if Primitive.to_proto(resource.url):
            res.url = Primitive.to_proto(resource.url)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionApiConfig(
            auth_fail_action=resource.auth_fail_action,
            login=resource.login,
            script=resource.script,
            security_level=resource.security_level,
            url=resource.url,
        )


class VersionApiConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionApiConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionApiConfig.from_proto(i) for i in resources]


class VersionDeployment(object):
    def __init__(
        self,
        files: dict = None,
        container: dict = None,
        zip: dict = None,
        cloud_build_options: dict = None,
    ):
        self.files = files
        self.container = container
        self.zip = zip
        self.cloud_build_options = cloud_build_options

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionDeployment()
        if Primitive.to_proto(resource.files):
            res.files = Primitive.to_proto(resource.files)
        if VersionDeploymentContainer.to_proto(resource.container):
            res.container.CopyFrom(
                VersionDeploymentContainer.to_proto(resource.container)
            )
        else:
            res.ClearField("container")
        if VersionDeploymentZip.to_proto(resource.zip):
            res.zip.CopyFrom(VersionDeploymentZip.to_proto(resource.zip))
        else:
            res.ClearField("zip")
        if VersionDeploymentCloudBuildOptions.to_proto(resource.cloud_build_options):
            res.cloud_build_options.CopyFrom(
                VersionDeploymentCloudBuildOptions.to_proto(
                    resource.cloud_build_options
                )
            )
        else:
            res.ClearField("cloud_build_options")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionDeployment(
            files=resource.files,
            container=resource.container,
            zip=resource.zip,
            cloud_build_options=resource.cloud_build_options,
        )


class VersionDeploymentArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionDeployment.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionDeployment.from_proto(i) for i in resources]


class VersionDeploymentFiles(object):
    def __init__(
        self, source_url: str = None, sha1_sum: str = None, mime_type: str = None
    ):
        self.source_url = source_url
        self.sha1_sum = sha1_sum
        self.mime_type = mime_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionDeploymentFiles()
        if Primitive.to_proto(resource.source_url):
            res.source_url = Primitive.to_proto(resource.source_url)
        if Primitive.to_proto(resource.sha1_sum):
            res.sha1_sum = Primitive.to_proto(resource.sha1_sum)
        if Primitive.to_proto(resource.mime_type):
            res.mime_type = Primitive.to_proto(resource.mime_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionDeploymentFiles(
            source_url=resource.source_url,
            sha1_sum=resource.sha1_sum,
            mime_type=resource.mime_type,
        )


class VersionDeploymentFilesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionDeploymentFiles.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionDeploymentFiles.from_proto(i) for i in resources]


class VersionDeploymentContainer(object):
    def __init__(self, image: str = None):
        self.image = image

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionDeploymentContainer()
        if Primitive.to_proto(resource.image):
            res.image = Primitive.to_proto(resource.image)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionDeploymentContainer(image=resource.image,)


class VersionDeploymentContainerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionDeploymentContainer.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionDeploymentContainer.from_proto(i) for i in resources]


class VersionDeploymentZip(object):
    def __init__(self, source_url: str = None, files_count: int = None):
        self.source_url = source_url
        self.files_count = files_count

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionDeploymentZip()
        if Primitive.to_proto(resource.source_url):
            res.source_url = Primitive.to_proto(resource.source_url)
        if Primitive.to_proto(resource.files_count):
            res.files_count = Primitive.to_proto(resource.files_count)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionDeploymentZip(
            source_url=resource.source_url, files_count=resource.files_count,
        )


class VersionDeploymentZipArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionDeploymentZip.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionDeploymentZip.from_proto(i) for i in resources]


class VersionDeploymentCloudBuildOptions(object):
    def __init__(self, app_yaml_path: str = None, cloud_build_timeout: str = None):
        self.app_yaml_path = app_yaml_path
        self.cloud_build_timeout = cloud_build_timeout

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionDeploymentCloudBuildOptions()
        if Primitive.to_proto(resource.app_yaml_path):
            res.app_yaml_path = Primitive.to_proto(resource.app_yaml_path)
        if Primitive.to_proto(resource.cloud_build_timeout):
            res.cloud_build_timeout = Primitive.to_proto(resource.cloud_build_timeout)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionDeploymentCloudBuildOptions(
            app_yaml_path=resource.app_yaml_path,
            cloud_build_timeout=resource.cloud_build_timeout,
        )


class VersionDeploymentCloudBuildOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionDeploymentCloudBuildOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionDeploymentCloudBuildOptions.from_proto(i) for i in resources]


class VersionHealthCheck(object):
    def __init__(
        self,
        disable_health_check: bool = None,
        host: str = None,
        healthy_threshold: int = None,
        unhealthy_threshold: int = None,
        restart_threshold: int = None,
        check_interval: str = None,
        timeout: str = None,
    ):
        self.disable_health_check = disable_health_check
        self.host = host
        self.healthy_threshold = healthy_threshold
        self.unhealthy_threshold = unhealthy_threshold
        self.restart_threshold = restart_threshold
        self.check_interval = check_interval
        self.timeout = timeout

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionHealthCheck()
        if Primitive.to_proto(resource.disable_health_check):
            res.disable_health_check = Primitive.to_proto(resource.disable_health_check)
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        if Primitive.to_proto(resource.healthy_threshold):
            res.healthy_threshold = Primitive.to_proto(resource.healthy_threshold)
        if Primitive.to_proto(resource.unhealthy_threshold):
            res.unhealthy_threshold = Primitive.to_proto(resource.unhealthy_threshold)
        if Primitive.to_proto(resource.restart_threshold):
            res.restart_threshold = Primitive.to_proto(resource.restart_threshold)
        if Primitive.to_proto(resource.check_interval):
            res.check_interval = Primitive.to_proto(resource.check_interval)
        if Primitive.to_proto(resource.timeout):
            res.timeout = Primitive.to_proto(resource.timeout)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionHealthCheck(
            disable_health_check=resource.disable_health_check,
            host=resource.host,
            healthy_threshold=resource.healthy_threshold,
            unhealthy_threshold=resource.unhealthy_threshold,
            restart_threshold=resource.restart_threshold,
            check_interval=resource.check_interval,
            timeout=resource.timeout,
        )


class VersionHealthCheckArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionHealthCheck.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionHealthCheck.from_proto(i) for i in resources]


class VersionReadinessCheck(object):
    def __init__(
        self,
        path: str = None,
        host: str = None,
        failure_threshold: int = None,
        success_threshold: int = None,
        check_interval: str = None,
        timeout: str = None,
        app_start_timeout: str = None,
    ):
        self.path = path
        self.host = host
        self.failure_threshold = failure_threshold
        self.success_threshold = success_threshold
        self.check_interval = check_interval
        self.timeout = timeout
        self.app_start_timeout = app_start_timeout

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionReadinessCheck()
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        if Primitive.to_proto(resource.failure_threshold):
            res.failure_threshold = Primitive.to_proto(resource.failure_threshold)
        if Primitive.to_proto(resource.success_threshold):
            res.success_threshold = Primitive.to_proto(resource.success_threshold)
        if Primitive.to_proto(resource.check_interval):
            res.check_interval = Primitive.to_proto(resource.check_interval)
        if Primitive.to_proto(resource.timeout):
            res.timeout = Primitive.to_proto(resource.timeout)
        if Primitive.to_proto(resource.app_start_timeout):
            res.app_start_timeout = Primitive.to_proto(resource.app_start_timeout)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionReadinessCheck(
            path=resource.path,
            host=resource.host,
            failure_threshold=resource.failure_threshold,
            success_threshold=resource.success_threshold,
            check_interval=resource.check_interval,
            timeout=resource.timeout,
            app_start_timeout=resource.app_start_timeout,
        )


class VersionReadinessCheckArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionReadinessCheck.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionReadinessCheck.from_proto(i) for i in resources]


class VersionLivenessCheck(object):
    def __init__(
        self,
        path: str = None,
        host: str = None,
        failure_threshold: int = None,
        success_threshold: int = None,
        check_interval: str = None,
        timeout: str = None,
        initial_delay: str = None,
    ):
        self.path = path
        self.host = host
        self.failure_threshold = failure_threshold
        self.success_threshold = success_threshold
        self.check_interval = check_interval
        self.timeout = timeout
        self.initial_delay = initial_delay

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionLivenessCheck()
        if Primitive.to_proto(resource.path):
            res.path = Primitive.to_proto(resource.path)
        if Primitive.to_proto(resource.host):
            res.host = Primitive.to_proto(resource.host)
        if Primitive.to_proto(resource.failure_threshold):
            res.failure_threshold = Primitive.to_proto(resource.failure_threshold)
        if Primitive.to_proto(resource.success_threshold):
            res.success_threshold = Primitive.to_proto(resource.success_threshold)
        if Primitive.to_proto(resource.check_interval):
            res.check_interval = Primitive.to_proto(resource.check_interval)
        if Primitive.to_proto(resource.timeout):
            res.timeout = Primitive.to_proto(resource.timeout)
        if Primitive.to_proto(resource.initial_delay):
            res.initial_delay = Primitive.to_proto(resource.initial_delay)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionLivenessCheck(
            path=resource.path,
            host=resource.host,
            failure_threshold=resource.failure_threshold,
            success_threshold=resource.success_threshold,
            check_interval=resource.check_interval,
            timeout=resource.timeout,
            initial_delay=resource.initial_delay,
        )


class VersionLivenessCheckArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionLivenessCheck.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionLivenessCheck.from_proto(i) for i in resources]


class VersionServiceAuthSpec(object):
    def __init__(
        self,
        audiences: list = None,
        iam_service_name: str = None,
        iam_resource_name: str = None,
        iam_policy_id: str = None,
        iam_policy_type: str = None,
        iam_permission: str = None,
        accept_gcloud_client_id: bool = None,
        clear: bool = None,
    ):
        self.audiences = audiences
        self.iam_service_name = iam_service_name
        self.iam_resource_name = iam_resource_name
        self.iam_policy_id = iam_policy_id
        self.iam_policy_type = iam_policy_type
        self.iam_permission = iam_permission
        self.accept_gcloud_client_id = accept_gcloud_client_id
        self.clear = clear

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionServiceAuthSpec()
        if Primitive.to_proto(resource.audiences):
            res.audiences.extend(Primitive.to_proto(resource.audiences))
        if Primitive.to_proto(resource.iam_service_name):
            res.iam_service_name = Primitive.to_proto(resource.iam_service_name)
        if Primitive.to_proto(resource.iam_resource_name):
            res.iam_resource_name = Primitive.to_proto(resource.iam_resource_name)
        if Primitive.to_proto(resource.iam_policy_id):
            res.iam_policy_id = Primitive.to_proto(resource.iam_policy_id)
        if Primitive.to_proto(resource.iam_policy_type):
            res.iam_policy_type = Primitive.to_proto(resource.iam_policy_type)
        if Primitive.to_proto(resource.iam_permission):
            res.iam_permission = Primitive.to_proto(resource.iam_permission)
        if Primitive.to_proto(resource.accept_gcloud_client_id):
            res.accept_gcloud_client_id = Primitive.to_proto(
                resource.accept_gcloud_client_id
            )
        if Primitive.to_proto(resource.clear):
            res.clear = Primitive.to_proto(resource.clear)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionServiceAuthSpec(
            audiences=resource.audiences,
            iam_service_name=resource.iam_service_name,
            iam_resource_name=resource.iam_resource_name,
            iam_policy_id=resource.iam_policy_id,
            iam_policy_type=resource.iam_policy_type,
            iam_permission=resource.iam_permission,
            accept_gcloud_client_id=resource.accept_gcloud_client_id,
            clear=resource.clear,
        )


class VersionServiceAuthSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionServiceAuthSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionServiceAuthSpec.from_proto(i) for i in resources]


class VersionServiceCorsSpec(object):
    def __init__(
        self,
        origin: list = None,
        method: list = None,
        header: list = None,
        exposed_header: list = None,
        allow_credential: bool = None,
        max_age_seconds: int = None,
    ):
        self.origin = origin
        self.method = method
        self.header = header
        self.exposed_header = exposed_header
        self.allow_credential = allow_credential
        self.max_age_seconds = max_age_seconds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionServiceCorsSpec()
        if Primitive.to_proto(resource.origin):
            res.origin.extend(Primitive.to_proto(resource.origin))
        if Primitive.to_proto(resource.method):
            res.method.extend(Primitive.to_proto(resource.method))
        if Primitive.to_proto(resource.header):
            res.header.extend(Primitive.to_proto(resource.header))
        if Primitive.to_proto(resource.exposed_header):
            res.exposed_header.extend(Primitive.to_proto(resource.exposed_header))
        if Primitive.to_proto(resource.allow_credential):
            res.allow_credential = Primitive.to_proto(resource.allow_credential)
        if Primitive.to_proto(resource.max_age_seconds):
            res.max_age_seconds = Primitive.to_proto(resource.max_age_seconds)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionServiceCorsSpec(
            origin=resource.origin,
            method=resource.method,
            header=resource.header,
            exposed_header=resource.exposed_header,
            allow_credential=resource.allow_credential,
            max_age_seconds=resource.max_age_seconds,
        )


class VersionServiceCorsSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionServiceCorsSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionServiceCorsSpec.from_proto(i) for i in resources]


class VersionEntrypoint(object):
    def __init__(self, shell: str = None):
        self.shell = shell

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionEntrypoint()
        if Primitive.to_proto(resource.shell):
            res.shell = Primitive.to_proto(resource.shell)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionEntrypoint(shell=resource.shell,)


class VersionEntrypointArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionEntrypoint.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionEntrypoint.from_proto(i) for i in resources]


class VersionVPCAccessConnector(object):
    def __init__(self, name: str = None, egress_setting: str = None):
        self.name = name
        self.egress_setting = egress_setting

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionVPCAccessConnector()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if VersionVPCAccessConnectorEgressSettingEnum.to_proto(resource.egress_setting):
            res.egress_setting = VersionVPCAccessConnectorEgressSettingEnum.to_proto(
                resource.egress_setting
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionVPCAccessConnector(
            name=resource.name, egress_setting=resource.egress_setting,
        )


class VersionVPCAccessConnectorArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionVPCAccessConnector.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionVPCAccessConnector.from_proto(i) for i in resources]


class VersionNetworkSettings(object):
    def __init__(self, ingress_traffic_allowed: str = None):
        self.ingress_traffic_allowed = ingress_traffic_allowed

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionNetworkSettings()
        if VersionNetworkSettingsIngressTrafficAllowedEnum.to_proto(
            resource.ingress_traffic_allowed
        ):
            res.ingress_traffic_allowed = VersionNetworkSettingsIngressTrafficAllowedEnum.to_proto(
                resource.ingress_traffic_allowed
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionNetworkSettings(
            ingress_traffic_allowed=resource.ingress_traffic_allowed,
        )


class VersionNetworkSettingsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionNetworkSettings.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionNetworkSettings.from_proto(i) for i in resources]


class VersionInstanceSpec(object):
    def __init__(self, sandboxes: list = None, ports: list = None):
        self.sandboxes = sandboxes
        self.ports = ports

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionInstanceSpec()
        if VersionInstanceSpecSandboxesArray.to_proto(resource.sandboxes):
            res.sandboxes.extend(
                VersionInstanceSpecSandboxesArray.to_proto(resource.sandboxes)
            )
        if VersionInstanceSpecPortsArray.to_proto(resource.ports):
            res.ports.extend(VersionInstanceSpecPortsArray.to_proto(resource.ports))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionInstanceSpec(sandboxes=resource.sandboxes, ports=resource.ports,)


class VersionInstanceSpecArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionInstanceSpec.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionInstanceSpec.from_proto(i) for i in resources]


class VersionInstanceSpecSandboxes(object):
    def __init__(self, name: str = None, containers: list = None):
        self.name = name
        self.containers = containers

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionInstanceSpecSandboxes()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if VersionInstanceSpecSandboxesContainersArray.to_proto(resource.containers):
            res.containers.extend(
                VersionInstanceSpecSandboxesContainersArray.to_proto(
                    resource.containers
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionInstanceSpecSandboxes(
            name=resource.name, containers=resource.containers,
        )


class VersionInstanceSpecSandboxesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionInstanceSpecSandboxes.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionInstanceSpecSandboxes.from_proto(i) for i in resources]


class VersionInstanceSpecSandboxesContainers(object):
    def __init__(self, ports: list = None):
        self.ports = ports

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionInstanceSpecSandboxesContainers()
        if int64Array.to_proto(resource.ports):
            res.ports.extend(int64Array.to_proto(resource.ports))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionInstanceSpecSandboxesContainers(ports=resource.ports,)


class VersionInstanceSpecSandboxesContainersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionInstanceSpecSandboxesContainers.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionInstanceSpecSandboxesContainers.from_proto(i) for i in resources]


class VersionInstanceSpecPorts(object):
    def __init__(
        self,
        name: str = None,
        sandbox: str = None,
        port: int = None,
        protocol: str = None,
        is_default: bool = None,
    ):
        self.name = name
        self.sandbox = sandbox
        self.port = port
        self.protocol = protocol
        self.is_default = is_default

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = version_pb2.AppengineVersionInstanceSpecPorts()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.sandbox):
            res.sandbox = Primitive.to_proto(resource.sandbox)
        if Primitive.to_proto(resource.port):
            res.port = Primitive.to_proto(resource.port)
        if VersionInstanceSpecPortsProtocolEnum.to_proto(resource.protocol):
            res.protocol = VersionInstanceSpecPortsProtocolEnum.to_proto(
                resource.protocol
            )
        if Primitive.to_proto(resource.is_default):
            res.is_default = Primitive.to_proto(resource.is_default)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return VersionInstanceSpecPorts(
            name=resource.name,
            sandbox=resource.sandbox,
            port=resource.port,
            protocol=resource.protocol,
            is_default=resource.is_default,
        )


class VersionInstanceSpecPortsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [VersionInstanceSpecPorts.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [VersionInstanceSpecPorts.from_proto(i) for i in resources]


class VersionInboundServicesEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionInboundServicesEnum.Value(
            "AppengineVersionInboundServicesEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionInboundServicesEnum.Name(resource)[
            len("AppengineVersionInboundServicesEnum") :
        ]


class VersionServingStatusEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionServingStatusEnum.Value(
            "AppengineVersionServingStatusEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionServingStatusEnum.Name(resource)[
            len("AppengineVersionServingStatusEnum") :
        ]


class VersionHandlersSecurityLevelEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersSecurityLevelEnum.Value(
            "AppengineVersionHandlersSecurityLevelEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersSecurityLevelEnum.Name(resource)[
            len("AppengineVersionHandlersSecurityLevelEnum") :
        ]


class VersionHandlersLoginEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersLoginEnum.Value(
            "AppengineVersionHandlersLoginEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersLoginEnum.Name(resource)[
            len("AppengineVersionHandlersLoginEnum") :
        ]


class VersionHandlersAuthFailActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersAuthFailActionEnum.Value(
            "AppengineVersionHandlersAuthFailActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersAuthFailActionEnum.Name(resource)[
            len("AppengineVersionHandlersAuthFailActionEnum") :
        ]


class VersionHandlersRedirectHttpResponseCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersRedirectHttpResponseCodeEnum.Value(
            "AppengineVersionHandlersRedirectHttpResponseCodeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionHandlersRedirectHttpResponseCodeEnum.Name(
            resource
        )[len("AppengineVersionHandlersRedirectHttpResponseCodeEnum") :]


class VersionErrorHandlersErrorCodeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionErrorHandlersErrorCodeEnum.Value(
            "AppengineVersionErrorHandlersErrorCodeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionErrorHandlersErrorCodeEnum.Name(resource)[
            len("AppengineVersionErrorHandlersErrorCodeEnum") :
        ]


class VersionApiConfigAuthFailActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionApiConfigAuthFailActionEnum.Value(
            "AppengineVersionApiConfigAuthFailActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionApiConfigAuthFailActionEnum.Name(resource)[
            len("AppengineVersionApiConfigAuthFailActionEnum") :
        ]


class VersionApiConfigLoginEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionApiConfigLoginEnum.Value(
            "AppengineVersionApiConfigLoginEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionApiConfigLoginEnum.Name(resource)[
            len("AppengineVersionApiConfigLoginEnum") :
        ]


class VersionApiConfigSecurityLevelEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionApiConfigSecurityLevelEnum.Value(
            "AppengineVersionApiConfigSecurityLevelEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionApiConfigSecurityLevelEnum.Name(resource)[
            len("AppengineVersionApiConfigSecurityLevelEnum") :
        ]


class VersionVPCAccessConnectorEgressSettingEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionVPCAccessConnectorEgressSettingEnum.Value(
            "AppengineVersionVPCAccessConnectorEgressSettingEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionVPCAccessConnectorEgressSettingEnum.Name(
            resource
        )[len("AppengineVersionVPCAccessConnectorEgressSettingEnum") :]


class VersionNetworkSettingsIngressTrafficAllowedEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionNetworkSettingsIngressTrafficAllowedEnum.Value(
            "AppengineVersionNetworkSettingsIngressTrafficAllowedEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionNetworkSettingsIngressTrafficAllowedEnum.Name(
            resource
        )[
            len("AppengineVersionNetworkSettingsIngressTrafficAllowedEnum") :
        ]


class VersionInstanceSpecPortsProtocolEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionInstanceSpecPortsProtocolEnum.Value(
            "AppengineVersionInstanceSpecPortsProtocolEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return version_pb2.AppengineVersionInstanceSpecPortsProtocolEnum.Name(resource)[
            len("AppengineVersionInstanceSpecPortsProtocolEnum") :
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
