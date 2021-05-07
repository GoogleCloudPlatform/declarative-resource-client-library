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
from google3.cloud.graphite.mmv2.services.google.dataproc import workflow_template_pb2
from google3.cloud.graphite.mmv2.services.google.dataproc import (
    workflow_template_pb2_grpc,
)

from typing import List


class WorkflowTemplate(object):
    def __init__(
        self,
        name: str = None,
        version: int = None,
        create_time: str = None,
        update_time: str = None,
        labels: dict = None,
        placement: dict = None,
        jobs: list = None,
        parameters: list = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.version = version
        self.labels = labels
        self.placement = placement
        self.jobs = jobs
        self.parameters = parameters
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = workflow_template_pb2_grpc.DataprocWorkflowTemplateServiceStub(
            channel.Channel()
        )
        request = workflow_template_pb2.ApplyDataprocWorkflowTemplateRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if WorkflowTemplatePlacement.to_proto(self.placement):
            request.resource.placement.CopyFrom(
                WorkflowTemplatePlacement.to_proto(self.placement)
            )
        else:
            request.resource.ClearField("placement")
        if WorkflowTemplateJobsArray.to_proto(self.jobs):
            request.resource.jobs.extend(WorkflowTemplateJobsArray.to_proto(self.jobs))
        if WorkflowTemplateParametersArray.to_proto(self.parameters):
            request.resource.parameters.extend(
                WorkflowTemplateParametersArray.to_proto(self.parameters)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyDataprocWorkflowTemplate(request)
        self.name = Primitive.from_proto(response.name)
        self.version = Primitive.from_proto(response.version)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.labels = Primitive.from_proto(response.labels)
        self.placement = WorkflowTemplatePlacement.from_proto(response.placement)
        self.jobs = WorkflowTemplateJobsArray.from_proto(response.jobs)
        self.parameters = WorkflowTemplateParametersArray.from_proto(
            response.parameters
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = workflow_template_pb2_grpc.DataprocWorkflowTemplateServiceStub(
            channel.Channel()
        )
        request = workflow_template_pb2.DeleteDataprocWorkflowTemplateRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.version):
            request.resource.version = Primitive.to_proto(self.version)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if WorkflowTemplatePlacement.to_proto(self.placement):
            request.resource.placement.CopyFrom(
                WorkflowTemplatePlacement.to_proto(self.placement)
            )
        else:
            request.resource.ClearField("placement")
        if WorkflowTemplateJobsArray.to_proto(self.jobs):
            request.resource.jobs.extend(WorkflowTemplateJobsArray.to_proto(self.jobs))
        if WorkflowTemplateParametersArray.to_proto(self.parameters):
            request.resource.parameters.extend(
                WorkflowTemplateParametersArray.to_proto(self.parameters)
            )
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteDataprocWorkflowTemplate(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = workflow_template_pb2_grpc.DataprocWorkflowTemplateServiceStub(
            channel.Channel()
        )
        request = workflow_template_pb2.ListDataprocWorkflowTemplateRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListDataprocWorkflowTemplate(request).items

    def to_proto(self):
        resource = workflow_template_pb2.DataprocWorkflowTemplate()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.version):
            resource.version = Primitive.to_proto(self.version)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if WorkflowTemplatePlacement.to_proto(self.placement):
            resource.placement.CopyFrom(
                WorkflowTemplatePlacement.to_proto(self.placement)
            )
        else:
            resource.ClearField("placement")
        if WorkflowTemplateJobsArray.to_proto(self.jobs):
            resource.jobs.extend(WorkflowTemplateJobsArray.to_proto(self.jobs))
        if WorkflowTemplateParametersArray.to_proto(self.parameters):
            resource.parameters.extend(
                WorkflowTemplateParametersArray.to_proto(self.parameters)
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class WorkflowTemplatePlacement(object):
    def __init__(self, managed_cluster: dict = None, cluster_selector: dict = None):
        self.managed_cluster = managed_cluster
        self.cluster_selector = cluster_selector

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplatePlacement()
        if WorkflowTemplatePlacementManagedCluster.to_proto(resource.managed_cluster):
            res.managed_cluster.CopyFrom(
                WorkflowTemplatePlacementManagedCluster.to_proto(
                    resource.managed_cluster
                )
            )
        else:
            res.ClearField("managed_cluster")
        if WorkflowTemplatePlacementClusterSelector.to_proto(resource.cluster_selector):
            res.cluster_selector.CopyFrom(
                WorkflowTemplatePlacementClusterSelector.to_proto(
                    resource.cluster_selector
                )
            )
        else:
            res.ClearField("cluster_selector")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacement(
            managed_cluster=resource.managed_cluster,
            cluster_selector=resource.cluster_selector,
        )


class WorkflowTemplatePlacementArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplatePlacement.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplatePlacement.from_proto(i) for i in resources]


class WorkflowTemplatePlacementManagedCluster(object):
    def __init__(
        self, cluster_name: str = None, config: dict = None, labels: dict = None
    ):
        self.cluster_name = cluster_name
        self.config = config
        self.labels = labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplatePlacementManagedCluster()
        if Primitive.to_proto(resource.cluster_name):
            res.cluster_name = Primitive.to_proto(resource.cluster_name)
        if ClusterClusterConfig.to_proto(resource.config):
            res.config.CopyFrom(ClusterClusterConfig.to_proto(resource.config))
        else:
            res.ClearField("config")
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementManagedCluster(
            cluster_name=resource.cluster_name,
            config=resource.config,
            labels=resource.labels,
        )


class WorkflowTemplatePlacementManagedClusterArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplatePlacementManagedCluster.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementManagedCluster.from_proto(i) for i in resources
        ]


class WorkflowTemplatePlacementClusterSelector(object):
    def __init__(self, zone: str = None, cluster_labels: dict = None):
        self.zone = zone
        self.cluster_labels = cluster_labels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplatePlacementClusterSelector()
        if Primitive.to_proto(resource.zone):
            res.zone = Primitive.to_proto(resource.zone)
        if Primitive.to_proto(resource.cluster_labels):
            res.cluster_labels = Primitive.to_proto(resource.cluster_labels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplatePlacementClusterSelector(
            zone=resource.zone, cluster_labels=resource.cluster_labels,
        )


class WorkflowTemplatePlacementClusterSelectorArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplatePlacementClusterSelector.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplatePlacementClusterSelector.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobs(object):
    def __init__(
        self,
        step_id: str = None,
        hadoop_job: dict = None,
        spark_job: dict = None,
        pyspark_job: dict = None,
        hive_job: dict = None,
        pig_job: dict = None,
        spark_r_job: dict = None,
        spark_sql_job: dict = None,
        presto_job: dict = None,
        labels: dict = None,
        scheduling: dict = None,
        prerequisite_step_ids: list = None,
    ):
        self.step_id = step_id
        self.hadoop_job = hadoop_job
        self.spark_job = spark_job
        self.pyspark_job = pyspark_job
        self.hive_job = hive_job
        self.pig_job = pig_job
        self.spark_r_job = spark_r_job
        self.spark_sql_job = spark_sql_job
        self.presto_job = presto_job
        self.labels = labels
        self.scheduling = scheduling
        self.prerequisite_step_ids = prerequisite_step_ids

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobs()
        if Primitive.to_proto(resource.step_id):
            res.step_id = Primitive.to_proto(resource.step_id)
        if WorkflowTemplateJobsHadoopJob.to_proto(resource.hadoop_job):
            res.hadoop_job.CopyFrom(
                WorkflowTemplateJobsHadoopJob.to_proto(resource.hadoop_job)
            )
        else:
            res.ClearField("hadoop_job")
        if WorkflowTemplateJobsSparkJob.to_proto(resource.spark_job):
            res.spark_job.CopyFrom(
                WorkflowTemplateJobsSparkJob.to_proto(resource.spark_job)
            )
        else:
            res.ClearField("spark_job")
        if WorkflowTemplateJobsPysparkJob.to_proto(resource.pyspark_job):
            res.pyspark_job.CopyFrom(
                WorkflowTemplateJobsPysparkJob.to_proto(resource.pyspark_job)
            )
        else:
            res.ClearField("pyspark_job")
        if WorkflowTemplateJobsHiveJob.to_proto(resource.hive_job):
            res.hive_job.CopyFrom(
                WorkflowTemplateJobsHiveJob.to_proto(resource.hive_job)
            )
        else:
            res.ClearField("hive_job")
        if WorkflowTemplateJobsPigJob.to_proto(resource.pig_job):
            res.pig_job.CopyFrom(WorkflowTemplateJobsPigJob.to_proto(resource.pig_job))
        else:
            res.ClearField("pig_job")
        if WorkflowTemplateJobsSparkRJob.to_proto(resource.spark_r_job):
            res.spark_r_job.CopyFrom(
                WorkflowTemplateJobsSparkRJob.to_proto(resource.spark_r_job)
            )
        else:
            res.ClearField("spark_r_job")
        if WorkflowTemplateJobsSparkSqlJob.to_proto(resource.spark_sql_job):
            res.spark_sql_job.CopyFrom(
                WorkflowTemplateJobsSparkSqlJob.to_proto(resource.spark_sql_job)
            )
        else:
            res.ClearField("spark_sql_job")
        if WorkflowTemplateJobsPrestoJob.to_proto(resource.presto_job):
            res.presto_job.CopyFrom(
                WorkflowTemplateJobsPrestoJob.to_proto(resource.presto_job)
            )
        else:
            res.ClearField("presto_job")
        if Primitive.to_proto(resource.labels):
            res.labels = Primitive.to_proto(resource.labels)
        if WorkflowTemplateJobsScheduling.to_proto(resource.scheduling):
            res.scheduling.CopyFrom(
                WorkflowTemplateJobsScheduling.to_proto(resource.scheduling)
            )
        else:
            res.ClearField("scheduling")
        if Primitive.to_proto(resource.prerequisite_step_ids):
            res.prerequisite_step_ids.extend(
                Primitive.to_proto(resource.prerequisite_step_ids)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobs(
            step_id=resource.step_id,
            hadoop_job=resource.hadoop_job,
            spark_job=resource.spark_job,
            pyspark_job=resource.pyspark_job,
            hive_job=resource.hive_job,
            pig_job=resource.pig_job,
            spark_r_job=resource.spark_r_job,
            spark_sql_job=resource.spark_sql_job,
            presto_job=resource.presto_job,
            labels=resource.labels,
            scheduling=resource.scheduling,
            prerequisite_step_ids=resource.prerequisite_step_ids,
        )


class WorkflowTemplateJobsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobs.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobs.from_proto(i) for i in resources]


class WorkflowTemplateJobsHadoopJob(object):
    def __init__(
        self,
        main_jar_file_uri: str = None,
        main_class: str = None,
        args: list = None,
        jar_file_uris: list = None,
        file_uris: list = None,
        archive_uris: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.main_jar_file_uri = main_jar_file_uri
        self.main_class = main_class
        self.args = args
        self.jar_file_uris = jar_file_uris
        self.file_uris = file_uris
        self.archive_uris = archive_uris
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsHadoopJob()
        if Primitive.to_proto(resource.main_jar_file_uri):
            res.main_jar_file_uri = Primitive.to_proto(resource.main_jar_file_uri)
        if Primitive.to_proto(resource.main_class):
            res.main_class = Primitive.to_proto(resource.main_class)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if Primitive.to_proto(resource.file_uris):
            res.file_uris.extend(Primitive.to_proto(resource.file_uris))
        if Primitive.to_proto(resource.archive_uris):
            res.archive_uris.extend(Primitive.to_proto(resource.archive_uris))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if WorkflowTemplateJobsHadoopJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                WorkflowTemplateJobsHadoopJobLoggingConfig.to_proto(
                    resource.logging_config
                )
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsHadoopJob(
            main_jar_file_uri=resource.main_jar_file_uri,
            main_class=resource.main_class,
            args=resource.args,
            jar_file_uris=resource.jar_file_uris,
            file_uris=resource.file_uris,
            archive_uris=resource.archive_uris,
            properties=resource.properties,
            logging_config=resource.logging_config,
        )


class WorkflowTemplateJobsHadoopJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsHadoopJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsHadoopJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsHadoopJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsHadoopJobLoggingConfig()
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsHadoopJobLoggingConfig(
            driver_log_levels=resource.driver_log_levels,
        )


class WorkflowTemplateJobsHadoopJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateJobsHadoopJobLoggingConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsHadoopJobLoggingConfig.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobsSparkJob(object):
    def __init__(
        self,
        main_jar_file_uri: str = None,
        main_class: str = None,
        args: list = None,
        jar_file_uris: list = None,
        file_uris: list = None,
        archive_uris: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.main_jar_file_uri = main_jar_file_uri
        self.main_class = main_class
        self.args = args
        self.jar_file_uris = jar_file_uris
        self.file_uris = file_uris
        self.archive_uris = archive_uris
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsSparkJob()
        if Primitive.to_proto(resource.main_jar_file_uri):
            res.main_jar_file_uri = Primitive.to_proto(resource.main_jar_file_uri)
        if Primitive.to_proto(resource.main_class):
            res.main_class = Primitive.to_proto(resource.main_class)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if Primitive.to_proto(resource.file_uris):
            res.file_uris.extend(Primitive.to_proto(resource.file_uris))
        if Primitive.to_proto(resource.archive_uris):
            res.archive_uris.extend(Primitive.to_proto(resource.archive_uris))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if WorkflowTemplateJobsSparkJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                WorkflowTemplateJobsSparkJobLoggingConfig.to_proto(
                    resource.logging_config
                )
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsSparkJob(
            main_jar_file_uri=resource.main_jar_file_uri,
            main_class=resource.main_class,
            args=resource.args,
            jar_file_uris=resource.jar_file_uris,
            file_uris=resource.file_uris,
            archive_uris=resource.archive_uris,
            properties=resource.properties,
            logging_config=resource.logging_config,
        )


class WorkflowTemplateJobsSparkJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsSparkJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsSparkJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsSparkJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsSparkJobLoggingConfig()
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsSparkJobLoggingConfig(
            driver_log_levels=resource.driver_log_levels,
        )


class WorkflowTemplateJobsSparkJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateJobsSparkJobLoggingConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsSparkJobLoggingConfig.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobsPysparkJob(object):
    def __init__(
        self,
        main_python_file_uri: str = None,
        args: list = None,
        python_file_uris: list = None,
        jar_file_uris: list = None,
        file_uris: list = None,
        archive_uris: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.main_python_file_uri = main_python_file_uri
        self.args = args
        self.python_file_uris = python_file_uris
        self.jar_file_uris = jar_file_uris
        self.file_uris = file_uris
        self.archive_uris = archive_uris
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsPysparkJob()
        if Primitive.to_proto(resource.main_python_file_uri):
            res.main_python_file_uri = Primitive.to_proto(resource.main_python_file_uri)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if Primitive.to_proto(resource.python_file_uris):
            res.python_file_uris.extend(Primitive.to_proto(resource.python_file_uris))
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if Primitive.to_proto(resource.file_uris):
            res.file_uris.extend(Primitive.to_proto(resource.file_uris))
        if Primitive.to_proto(resource.archive_uris):
            res.archive_uris.extend(Primitive.to_proto(resource.archive_uris))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if WorkflowTemplateJobsPysparkJobLoggingConfig.to_proto(
            resource.logging_config
        ):
            res.logging_config.CopyFrom(
                WorkflowTemplateJobsPysparkJobLoggingConfig.to_proto(
                    resource.logging_config
                )
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPysparkJob(
            main_python_file_uri=resource.main_python_file_uri,
            args=resource.args,
            python_file_uris=resource.python_file_uris,
            jar_file_uris=resource.jar_file_uris,
            file_uris=resource.file_uris,
            archive_uris=resource.archive_uris,
            properties=resource.properties,
            logging_config=resource.logging_config,
        )


class WorkflowTemplateJobsPysparkJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsPysparkJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsPysparkJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsPysparkJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocWorkflowTemplateJobsPysparkJobLoggingConfig()
        )
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPysparkJobLoggingConfig(
            driver_log_levels=resource.driver_log_levels,
        )


class WorkflowTemplateJobsPysparkJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateJobsPysparkJobLoggingConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsPysparkJobLoggingConfig.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobsHiveJob(object):
    def __init__(
        self,
        query_file_uri: str = None,
        query_list: dict = None,
        continue_on_failure: bool = None,
        script_variables: dict = None,
        properties: dict = None,
        jar_file_uris: list = None,
    ):
        self.query_file_uri = query_file_uri
        self.query_list = query_list
        self.continue_on_failure = continue_on_failure
        self.script_variables = script_variables
        self.properties = properties
        self.jar_file_uris = jar_file_uris

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsHiveJob()
        if Primitive.to_proto(resource.query_file_uri):
            res.query_file_uri = Primitive.to_proto(resource.query_file_uri)
        if WorkflowTemplateJobsHiveJobQueryList.to_proto(resource.query_list):
            res.query_list.CopyFrom(
                WorkflowTemplateJobsHiveJobQueryList.to_proto(resource.query_list)
            )
        else:
            res.ClearField("query_list")
        if Primitive.to_proto(resource.continue_on_failure):
            res.continue_on_failure = Primitive.to_proto(resource.continue_on_failure)
        if Primitive.to_proto(resource.script_variables):
            res.script_variables = Primitive.to_proto(resource.script_variables)
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsHiveJob(
            query_file_uri=resource.query_file_uri,
            query_list=resource.query_list,
            continue_on_failure=resource.continue_on_failure,
            script_variables=resource.script_variables,
            properties=resource.properties,
            jar_file_uris=resource.jar_file_uris,
        )


class WorkflowTemplateJobsHiveJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsHiveJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsHiveJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsHiveJobQueryList(object):
    def __init__(self, queries: list = None):
        self.queries = queries

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsHiveJobQueryList()
        if Primitive.to_proto(resource.queries):
            res.queries.extend(Primitive.to_proto(resource.queries))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsHiveJobQueryList(queries=resource.queries,)


class WorkflowTemplateJobsHiveJobQueryListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsHiveJobQueryList.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsHiveJobQueryList.from_proto(i) for i in resources]


class WorkflowTemplateJobsPigJob(object):
    def __init__(
        self,
        query_file_uri: str = None,
        query_list: dict = None,
        continue_on_failure: bool = None,
        script_variables: dict = None,
        properties: dict = None,
        jar_file_uris: list = None,
        logging_config: dict = None,
    ):
        self.query_file_uri = query_file_uri
        self.query_list = query_list
        self.continue_on_failure = continue_on_failure
        self.script_variables = script_variables
        self.properties = properties
        self.jar_file_uris = jar_file_uris
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsPigJob()
        if Primitive.to_proto(resource.query_file_uri):
            res.query_file_uri = Primitive.to_proto(resource.query_file_uri)
        if WorkflowTemplateJobsPigJobQueryList.to_proto(resource.query_list):
            res.query_list.CopyFrom(
                WorkflowTemplateJobsPigJobQueryList.to_proto(resource.query_list)
            )
        else:
            res.ClearField("query_list")
        if Primitive.to_proto(resource.continue_on_failure):
            res.continue_on_failure = Primitive.to_proto(resource.continue_on_failure)
        if Primitive.to_proto(resource.script_variables):
            res.script_variables = Primitive.to_proto(resource.script_variables)
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if WorkflowTemplateJobsPigJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                WorkflowTemplateJobsPigJobLoggingConfig.to_proto(
                    resource.logging_config
                )
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPigJob(
            query_file_uri=resource.query_file_uri,
            query_list=resource.query_list,
            continue_on_failure=resource.continue_on_failure,
            script_variables=resource.script_variables,
            properties=resource.properties,
            jar_file_uris=resource.jar_file_uris,
            logging_config=resource.logging_config,
        )


class WorkflowTemplateJobsPigJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsPigJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsPigJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsPigJobQueryList(object):
    def __init__(self, queries: list = None):
        self.queries = queries

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsPigJobQueryList()
        if Primitive.to_proto(resource.queries):
            res.queries.extend(Primitive.to_proto(resource.queries))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPigJobQueryList(queries=resource.queries,)


class WorkflowTemplateJobsPigJobQueryListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsPigJobQueryList.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsPigJobQueryList.from_proto(i) for i in resources]


class WorkflowTemplateJobsPigJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsPigJobLoggingConfig()
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPigJobLoggingConfig(
            driver_log_levels=resource.driver_log_levels,
        )


class WorkflowTemplateJobsPigJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsPigJobLoggingConfig.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsPigJobLoggingConfig.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobsSparkRJob(object):
    def __init__(
        self,
        main_r_file_uri: str = None,
        args: list = None,
        file_uris: list = None,
        archive_uris: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.main_r_file_uri = main_r_file_uri
        self.args = args
        self.file_uris = file_uris
        self.archive_uris = archive_uris
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsSparkRJob()
        if Primitive.to_proto(resource.main_r_file_uri):
            res.main_r_file_uri = Primitive.to_proto(resource.main_r_file_uri)
        if Primitive.to_proto(resource.args):
            res.args.extend(Primitive.to_proto(resource.args))
        if Primitive.to_proto(resource.file_uris):
            res.file_uris.extend(Primitive.to_proto(resource.file_uris))
        if Primitive.to_proto(resource.archive_uris):
            res.archive_uris.extend(Primitive.to_proto(resource.archive_uris))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if WorkflowTemplateJobsSparkRJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                WorkflowTemplateJobsSparkRJobLoggingConfig.to_proto(
                    resource.logging_config
                )
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsSparkRJob(
            main_r_file_uri=resource.main_r_file_uri,
            args=resource.args,
            file_uris=resource.file_uris,
            archive_uris=resource.archive_uris,
            properties=resource.properties,
            logging_config=resource.logging_config,
        )


class WorkflowTemplateJobsSparkRJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsSparkRJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsSparkRJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsSparkRJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsSparkRJobLoggingConfig()
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsSparkRJobLoggingConfig(
            driver_log_levels=resource.driver_log_levels,
        )


class WorkflowTemplateJobsSparkRJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateJobsSparkRJobLoggingConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsSparkRJobLoggingConfig.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobsSparkSqlJob(object):
    def __init__(
        self,
        query_file_uri: str = None,
        query_list: dict = None,
        script_variables: dict = None,
        properties: dict = None,
        jar_file_uris: list = None,
        logging_config: dict = None,
    ):
        self.query_file_uri = query_file_uri
        self.query_list = query_list
        self.script_variables = script_variables
        self.properties = properties
        self.jar_file_uris = jar_file_uris
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsSparkSqlJob()
        if Primitive.to_proto(resource.query_file_uri):
            res.query_file_uri = Primitive.to_proto(resource.query_file_uri)
        if WorkflowTemplateJobsSparkSqlJobQueryList.to_proto(resource.query_list):
            res.query_list.CopyFrom(
                WorkflowTemplateJobsSparkSqlJobQueryList.to_proto(resource.query_list)
            )
        else:
            res.ClearField("query_list")
        if Primitive.to_proto(resource.script_variables):
            res.script_variables = Primitive.to_proto(resource.script_variables)
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if Primitive.to_proto(resource.jar_file_uris):
            res.jar_file_uris.extend(Primitive.to_proto(resource.jar_file_uris))
        if WorkflowTemplateJobsSparkSqlJobLoggingConfig.to_proto(
            resource.logging_config
        ):
            res.logging_config.CopyFrom(
                WorkflowTemplateJobsSparkSqlJobLoggingConfig.to_proto(
                    resource.logging_config
                )
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsSparkSqlJob(
            query_file_uri=resource.query_file_uri,
            query_list=resource.query_list,
            script_variables=resource.script_variables,
            properties=resource.properties,
            jar_file_uris=resource.jar_file_uris,
            logging_config=resource.logging_config,
        )


class WorkflowTemplateJobsSparkSqlJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsSparkSqlJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsSparkSqlJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsSparkSqlJobQueryList(object):
    def __init__(self, queries: list = None):
        self.queries = queries

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsSparkSqlJobQueryList()
        if Primitive.to_proto(resource.queries):
            res.queries.extend(Primitive.to_proto(resource.queries))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsSparkSqlJobQueryList(queries=resource.queries,)


class WorkflowTemplateJobsSparkSqlJobQueryListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsSparkSqlJobQueryList.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsSparkSqlJobQueryList.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobsSparkSqlJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            workflow_template_pb2.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig()
        )
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsSparkSqlJobLoggingConfig(
            driver_log_levels=resource.driver_log_levels,
        )


class WorkflowTemplateJobsSparkSqlJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateJobsSparkSqlJobLoggingConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsSparkSqlJobLoggingConfig.from_proto(i)
            for i in resources
        ]


class WorkflowTemplateJobsPrestoJob(object):
    def __init__(
        self,
        query_file_uri: str = None,
        query_list: dict = None,
        continue_on_failure: bool = None,
        output_format: str = None,
        client_tags: list = None,
        properties: dict = None,
        logging_config: dict = None,
    ):
        self.query_file_uri = query_file_uri
        self.query_list = query_list
        self.continue_on_failure = continue_on_failure
        self.output_format = output_format
        self.client_tags = client_tags
        self.properties = properties
        self.logging_config = logging_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsPrestoJob()
        if Primitive.to_proto(resource.query_file_uri):
            res.query_file_uri = Primitive.to_proto(resource.query_file_uri)
        if WorkflowTemplateJobsPrestoJobQueryList.to_proto(resource.query_list):
            res.query_list.CopyFrom(
                WorkflowTemplateJobsPrestoJobQueryList.to_proto(resource.query_list)
            )
        else:
            res.ClearField("query_list")
        if Primitive.to_proto(resource.continue_on_failure):
            res.continue_on_failure = Primitive.to_proto(resource.continue_on_failure)
        if Primitive.to_proto(resource.output_format):
            res.output_format = Primitive.to_proto(resource.output_format)
        if Primitive.to_proto(resource.client_tags):
            res.client_tags.extend(Primitive.to_proto(resource.client_tags))
        if Primitive.to_proto(resource.properties):
            res.properties = Primitive.to_proto(resource.properties)
        if WorkflowTemplateJobsPrestoJobLoggingConfig.to_proto(resource.logging_config):
            res.logging_config.CopyFrom(
                WorkflowTemplateJobsPrestoJobLoggingConfig.to_proto(
                    resource.logging_config
                )
            )
        else:
            res.ClearField("logging_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPrestoJob(
            query_file_uri=resource.query_file_uri,
            query_list=resource.query_list,
            continue_on_failure=resource.continue_on_failure,
            output_format=resource.output_format,
            client_tags=resource.client_tags,
            properties=resource.properties,
            logging_config=resource.logging_config,
        )


class WorkflowTemplateJobsPrestoJobArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsPrestoJob.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsPrestoJob.from_proto(i) for i in resources]


class WorkflowTemplateJobsPrestoJobQueryList(object):
    def __init__(self, queries: list = None):
        self.queries = queries

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsPrestoJobQueryList()
        if Primitive.to_proto(resource.queries):
            res.queries.extend(Primitive.to_proto(resource.queries))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPrestoJobQueryList(queries=resource.queries,)


class WorkflowTemplateJobsPrestoJobQueryListArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsPrestoJobQueryList.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsPrestoJobQueryList.from_proto(i) for i in resources]


class WorkflowTemplateJobsPrestoJobLoggingConfig(object):
    def __init__(self, driver_log_levels: dict = None):
        self.driver_log_levels = driver_log_levels

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsPrestoJobLoggingConfig()
        if Primitive.to_proto(resource.driver_log_levels):
            res.driver_log_levels = Primitive.to_proto(resource.driver_log_levels)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsPrestoJobLoggingConfig(
            driver_log_levels=resource.driver_log_levels,
        )


class WorkflowTemplateJobsPrestoJobLoggingConfigArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateJobsPrestoJobLoggingConfig.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateJobsPrestoJobLoggingConfig.from_proto(i) for i in resources
        ]


class WorkflowTemplateJobsScheduling(object):
    def __init__(
        self, max_failures_per_hour: int = None, max_failures_total: int = None
    ):
        self.max_failures_per_hour = max_failures_per_hour
        self.max_failures_total = max_failures_total

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateJobsScheduling()
        if Primitive.to_proto(resource.max_failures_per_hour):
            res.max_failures_per_hour = Primitive.to_proto(
                resource.max_failures_per_hour
            )
        if Primitive.to_proto(resource.max_failures_total):
            res.max_failures_total = Primitive.to_proto(resource.max_failures_total)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateJobsScheduling(
            max_failures_per_hour=resource.max_failures_per_hour,
            max_failures_total=resource.max_failures_total,
        )


class WorkflowTemplateJobsSchedulingArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateJobsScheduling.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateJobsScheduling.from_proto(i) for i in resources]


class WorkflowTemplateParameters(object):
    def __init__(
        self,
        name: str = None,
        fields: list = None,
        description: str = None,
        validation: dict = None,
    ):
        self.name = name
        self.fields = fields
        self.description = description
        self.validation = validation

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateParameters()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.fields):
            res.fields.extend(Primitive.to_proto(resource.fields))
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if WorkflowTemplateParametersValidation.to_proto(resource.validation):
            res.validation.CopyFrom(
                WorkflowTemplateParametersValidation.to_proto(resource.validation)
            )
        else:
            res.ClearField("validation")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateParameters(
            name=resource.name,
            fields=resource.fields,
            description=resource.description,
            validation=resource.validation,
        )


class WorkflowTemplateParametersArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateParameters.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateParameters.from_proto(i) for i in resources]


class WorkflowTemplateParametersValidation(object):
    def __init__(self, regex: dict = None, values: dict = None):
        self.regex = regex
        self.values = values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateParametersValidation()
        if WorkflowTemplateParametersValidationRegex.to_proto(resource.regex):
            res.regex.CopyFrom(
                WorkflowTemplateParametersValidationRegex.to_proto(resource.regex)
            )
        else:
            res.ClearField("regex")
        if WorkflowTemplateParametersValidationValues.to_proto(resource.values):
            res.values.CopyFrom(
                WorkflowTemplateParametersValidationValues.to_proto(resource.values)
            )
        else:
            res.ClearField("values")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateParametersValidation(
            regex=resource.regex, values=resource.values,
        )


class WorkflowTemplateParametersValidationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [WorkflowTemplateParametersValidation.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [WorkflowTemplateParametersValidation.from_proto(i) for i in resources]


class WorkflowTemplateParametersValidationRegex(object):
    def __init__(self, regexes: list = None):
        self.regexes = regexes

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateParametersValidationRegex()
        if Primitive.to_proto(resource.regexes):
            res.regexes.extend(Primitive.to_proto(resource.regexes))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateParametersValidationRegex(regexes=resource.regexes,)


class WorkflowTemplateParametersValidationRegexArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateParametersValidationRegex.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateParametersValidationRegex.from_proto(i) for i in resources
        ]


class WorkflowTemplateParametersValidationValues(object):
    def __init__(self, values: list = None):
        self.values = values

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = workflow_template_pb2.DataprocWorkflowTemplateParametersValidationValues()
        if Primitive.to_proto(resource.values):
            res.values.extend(Primitive.to_proto(resource.values))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return WorkflowTemplateParametersValidationValues(values=resource.values,)


class WorkflowTemplateParametersValidationValuesArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            WorkflowTemplateParametersValidationValues.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            WorkflowTemplateParametersValidationValues.from_proto(i) for i in resources
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
