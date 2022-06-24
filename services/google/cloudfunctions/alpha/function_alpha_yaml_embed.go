// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// GENERATED BY gen_go_data.go
// gen_go_data -package alpha -var YAML_function blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudfunctions/alpha/function.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudfunctions/alpha/function.yaml
var YAML_function = []byte("info:\n  title: CloudFunctions/Function\n  description: The CloudFunctions Function resource\n  x-dcl-struct-name: Function\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a Function\n    parameters:\n    - name: Function\n      required: true\n      description: A full instance of a Function\n  apply:\n    description: The function used to apply information about a Function\n    parameters:\n    - name: Function\n      required: true\n      description: A full instance of a Function\n  delete:\n    description: The function used to delete a Function\n    parameters:\n    - name: Function\n      required: true\n      description: A full instance of a Function\n  deleteAll:\n    description: The function used to delete all Function\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: region\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Function\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: region\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Function:\n      title: Function\n      x-dcl-id: projects/{{project}}/locations/{{region}}/functions/{{name}}\n      x-dcl-locations:\n      - region\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: true\n      type: object\n      required:\n      - name\n      - runtime\n      - region\n      - project\n      properties:\n        availableMemoryMb:\n          type: integer\n          format: int64\n          x-dcl-go-name: AvailableMemoryMb\n          description: 'Memory (in MB), available to the function. Default value is\n            256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.'\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: User-provided description of a function.\n        entryPoint:\n          type: string\n          x-dcl-go-name: EntryPoint\n          description: |-\n            The name of the function (as defined in source code) that will be\n            executed. Defaults to the resource name suffix, if not specified. For\n            backward compatibility, if function with given name is not found, then the\n            system will try to use function named \"function\".\n            For Node.js this is name of a function exported by the module specified\n            in `source_location`.\n          x-kubernetes-immutable: true\n        environmentVariables:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: EnvironmentVariables\n          description: Environment variables that shall be available during function\n            execution.\n        eventTrigger:\n          type: object\n          x-dcl-go-name: EventTrigger\n          x-dcl-go-type: FunctionEventTrigger\n          description: A source that fires events in response to a condition in another\n            service.\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - httpsTrigger\n          required:\n          - eventType\n          - resource\n          properties:\n            eventType:\n              type: string\n              x-dcl-go-name: EventType\n              description: |-\n                Required. The type of event to observe. For example:\n                `providers/cloud.storage/eventTypes/object.change` and\n                `providers/cloud.pubsub/eventTypes/topic.publish`.\n\n                Event types match pattern `providers/*/eventTypes/*.*`.\n                The pattern contains:\n\n                1. namespace: For example, `cloud.storage` and\n                   `google.firebase.analytics`.\n                2. resource type: The type of resource on which event occurs. For\n                   example, the Google Cloud Storage API includes the type `object`.\n                3. action: The action that generates the event. For example, action for\n                   a Google Cloud Storage Object is 'change'.\n                These parts are lower case.\n              x-kubernetes-immutable: true\n            failurePolicy:\n              type: boolean\n              x-dcl-go-name: FailurePolicy\n              description: Specifies policy for failed executions.\n              x-kubernetes-immutable: true\n            resource:\n              type: string\n              x-dcl-go-name: Resource\n              description: |-\n                Required. The resource(s) from which to observe events, for example,\n                `projects/_/buckets/myBucket`.\n\n                Not all syntactically correct values are accepted by all services. For\n                example:\n\n                1. The authorization model must support it. Google Cloud Functions\n                   only allows EventTriggers to be deployed that observe resources in the\n                   same project as the `Function`.\n                2. The resource type must match the pattern expected for an\n                   `event_type`. For example, an `EventTrigger` that has an\n                   `event_type` of \"google.pubsub.topic.publish\" should have a resource\n                   that matches Google Cloud Pub/Sub topics.\n\n                Additionally, some services may support short names when creating an\n                `EventTrigger`. These will always be returned in the normalized \"long\"\n                format.\n\n                See each *service's* documentation for supported formats.\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Storage/Bucket\n                field: name\n                format: projects/{{project}}/buckets/{{name}}\n              - resource: Pubsub/Topic\n                field: name\n            service:\n              type: string\n              x-dcl-go-name: Service\n              description: |\n                The hostname of the service that should be observed.\n\n                If no string is provided, the default service implementing the API will\n                be used. For example, `storage.googleapis.com` is the default for all\n                event types in the `google.storage` namespace.\n              x-kubernetes-immutable: true\n        httpsTrigger:\n          type: object\n          x-dcl-go-name: HttpsTrigger\n          x-dcl-go-type: FunctionHttpsTrigger\n          description: An HTTPS endpoint type of source that can be triggered via\n            URL.\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - eventTrigger\n          properties:\n            securityLevel:\n              type: string\n              x-dcl-go-name: SecurityLevel\n              x-dcl-go-type: FunctionHttpsTriggerSecurityLevelEnum\n              description: 'Both HTTP and HTTPS requests with URLs that match the\n                handler succeed without redirects. The application can examine the\n                request to determine which protocol was used and respond accordingly.\n                Possible values: SECURITY_LEVEL_UNSPECIFIED, SECURE_ALWAYS, SECURE_OPTIONAL'\n              x-kubernetes-immutable: true\n              enum:\n              - SECURITY_LEVEL_UNSPECIFIED\n              - SECURE_ALWAYS\n              - SECURE_OPTIONAL\n            url:\n              type: string\n              x-dcl-go-name: Url\n              readOnly: true\n              description: Output only. The deployed url for the function.\n              x-kubernetes-immutable: true\n        ingressSettings:\n          type: string\n          x-dcl-go-name: IngressSettings\n          x-dcl-go-type: FunctionIngressSettingsEnum\n          description: |-\n            The ingress settings for the function, controlling what traffic can reach\n            it. Possible values: INGRESS_SETTINGS_UNSPECIFIED, ALLOW_ALL, ALLOW_INTERNAL_ONLY, ALLOW_INTERNAL_AND_GCLB\n          enum:\n          - INGRESS_SETTINGS_UNSPECIFIED\n          - ALLOW_ALL\n          - ALLOW_INTERNAL_ONLY\n          - ALLOW_INTERNAL_AND_GCLB\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Labels associated with this Cloud Function.\n        maxInstances:\n          type: integer\n          format: int64\n          x-dcl-go-name: MaxInstances\n          description: |-\n            The limit on the maximum number of function instances that may coexist at a\n            given time.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: A user-defined name of the function. Function names must be\n            unique globally.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project id of the function.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        region:\n          type: string\n          x-dcl-go-name: Region\n          description: The name of the Cloud Functions region of the function.\n          x-kubernetes-immutable: true\n        runtime:\n          type: string\n          x-dcl-go-name: Runtime\n          description: |\n            The runtime in which to run the function. Required when deploying a new\n            function, optional when updating an existing function. For a complete\n            list of possible choices, see the\n            [`gcloud` command\n            reference](/sdk/gcloud/reference/functions/deploy#--runtime).\n        serviceAccountEmail:\n          type: string\n          x-dcl-go-name: ServiceAccountEmail\n          description: |-\n            The email of the function's service account. If empty, defaults to\n            `{project_id}@appspot.gserviceaccount.com`.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Iam/ServiceAccount\n            field: email\n        sourceArchiveUrl:\n          type: string\n          x-dcl-go-name: SourceArchiveUrl\n          description: The Google Cloud Storage URL, starting with gs://, pointing\n            to the zip archive which contains the function.\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - sourceRepository\n        sourceRepository:\n          type: object\n          x-dcl-go-name: SourceRepository\n          x-dcl-go-type: FunctionSourceRepository\n          description: Represents parameters related to source repository where a\n            function is hosted.\n          x-kubernetes-immutable: true\n          x-dcl-conflicts:\n          - sourceArchiveUrl\n          required:\n          - url\n          properties:\n            deployedUrl:\n              type: string\n              x-dcl-go-name: DeployedUrl\n              readOnly: true\n              description: |-\n                Output only. The URL pointing to the hosted repository where the function\n                were defined at the time of deployment. It always points to a specific\n                commit in the format described above.\n              x-kubernetes-immutable: true\n            url:\n              type: string\n              x-dcl-go-name: Url\n              description: |-\n                The URL pointing to the hosted repository where the function is defined.\n                There are supported Cloud Source Repository URLs in the following\n                formats:\n\n                To refer to a specific commit:\n                `https://source.developers.google.com/projects/*/repos/*/revisions/*/paths/*`\n                To refer to a moveable alias (branch):\n                `https://source.developers.google.com/projects/*/repos/*/moveable-aliases/*/paths/*`\n                In particular, to refer to HEAD use `master` moveable alias.\n                To refer to a specific fixed alias (tag):\n                `https://source.developers.google.com/projects/*/repos/*/fixed-aliases/*/paths/*`\n\n                You may omit `paths/*` if you want to use the main directory.\n              x-kubernetes-immutable: true\n        status:\n          type: string\n          x-dcl-go-name: Status\n          x-dcl-go-type: FunctionStatusEnum\n          readOnly: true\n          description: 'Output only. Status of the function deployment. Possible values:\n            CLOUD_FUNCTION_STATUS_UNSPECIFIED, ACTIVE, OFFLINE, DEPLOY_IN_PROGRESS,\n            DELETE_IN_PROGRESS, UNKNOWN'\n          x-kubernetes-immutable: true\n          enum:\n          - CLOUD_FUNCTION_STATUS_UNSPECIFIED\n          - ACTIVE\n          - OFFLINE\n          - DEPLOY_IN_PROGRESS\n          - DELETE_IN_PROGRESS\n          - UNKNOWN\n        timeout:\n          type: string\n          x-dcl-go-name: Timeout\n          description: |-\n            The function execution timeout. Execution is considered failed and\n            can be terminated if the function is not completed at the end of the\n            timeout period. Defaults to 60 seconds.\n        updateTime:\n          type: string\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The last update timestamp of a Cloud Function\n            in RFC3339 UTC 'Zulu' format, with nanosecond resolution and up to nine\n            fractional digits.\n          x-kubernetes-immutable: true\n        versionId:\n          type: integer\n          format: int64\n          x-dcl-go-name: VersionId\n          readOnly: true\n          description: |-\n            Output only. The version identifier of the Cloud Function. Each deployment attempt\n            results in a new version of a function being created.\n          x-kubernetes-immutable: true\n        vpcConnector:\n          type: string\n          x-dcl-go-name: VPCConnector\n          description: |-\n            The VPC Network Connector that this cloud function can connect to. It can\n            be either the fully-qualified URI, or the short name of the network\n            connector resource. The format of this field is\n            `projects/*/locations/*/connectors/*`\n          x-dcl-server-default: true\n          x-dcl-references:\n          - resource: Vpcaccess/Connector\n            field: name\n        vpcConnectorEgressSettings:\n          type: string\n          x-dcl-go-name: VPCConnectorEgressSettings\n          x-dcl-go-type: FunctionVPCConnectorEgressSettingsEnum\n          description: |-\n            The egress settings for the connector, controlling what traffic is diverted\n            through it. Possible values: VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED, PRIVATE_RANGES_ONLY, ALL_TRAFFIC\n          enum:\n          - VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED\n          - PRIVATE_RANGES_ONLY\n          - ALL_TRAFFIC\n")

// 14783 bytes
// MD5: 330da9c12587c3371ee1d702aceddfdd
