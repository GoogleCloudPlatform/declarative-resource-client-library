// Copyright 2023 Google LLC. All Rights Reserved.
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
// gen_go_data -package beta -var YAML_job blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudscheduler/beta/job.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudscheduler/beta/job.yaml
var YAML_job = []byte("info:\n  title: CloudScheduler/Job\n  description: The CloudScheduler Job resource\n  x-dcl-struct-name: Job\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Job\n    parameters:\n    - name: job\n      required: true\n      description: A full instance of a Job\n  apply:\n    description: The function used to apply information about a Job\n    parameters:\n    - name: job\n      required: true\n      description: A full instance of a Job\n  delete:\n    description: The function used to delete a Job\n    parameters:\n    - name: job\n      required: true\n      description: A full instance of a Job\n  deleteAll:\n    description: The function used to delete all Job\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Job\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Job:\n      title: Job\n      x-dcl-id: projects/{{project}}/locations/{{location}}/jobs/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        appEngineHttpTarget:\n          type: object\n          x-dcl-go-name: AppEngineHttpTarget\n          x-dcl-go-type: JobAppEngineHttpTarget\n          description: App Engine HTTP target.\n          x-dcl-conflicts:\n          - pubsubTarget\n          - httpTarget\n          properties:\n            appEngineRouting:\n              type: object\n              x-dcl-go-name: AppEngineRouting\n              x-dcl-go-type: JobAppEngineHttpTargetAppEngineRouting\n              description: App Engine Routing setting for the job.\n              properties:\n                host:\n                  type: string\n                  x-dcl-go-name: Host\n                  readOnly: true\n                  description: 'Output only. The host that the job is sent to. For\n                    more information about how App Engine requests are routed, see\n                    [here](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed).\n                    The host is constructed as: * `host = [application_domain_name]`\n                    `| [service] + ''.'' + [application_domain_name]` `| [version]\n                    + ''.'' + [application_domain_name]` `| [version_dot_service]+\n                    ''.'' + [application_domain_name]` `| [instance] + ''.'' + [application_domain_name]`\n                    `| [instance_dot_service] + ''.'' + [application_domain_name]`\n                    `| [instance_dot_version] + ''.'' + [application_domain_name]`\n                    `| [instance_dot_version_dot_service] + ''.'' + [application_domain_name]`\n                    * `application_domain_name` = The domain name of the app, for\n                    example .appspot.com, which is associated with the job''s project\n                    ID. * `service =` service * `version =` version * `version_dot_service\n                    =` version `+ ''.'' +` service * `instance =` instance * `instance_dot_service\n                    =` instance `+ ''.'' +` service * `instance_dot_version =` instance\n                    `+ ''.'' +` version * `instance_dot_version_dot_service =` instance\n                    `+ ''.'' +` version `+ ''.'' +` service If service is empty, then\n                    the job will be sent to the service which is the default service\n                    when the job is attempted. If version is empty, then the job will\n                    be sent to the version which is the default version when the job\n                    is attempted. If instance is empty, then the job will be sent\n                    to an instance which is available when the job is attempted. If\n                    service, version, or instance is invalid, then the job will be\n                    sent to the default version of the default service when the job\n                    is attempted.'\n                instance:\n                  type: string\n                  x-dcl-go-name: Instance\n                  description: App instance. By default, the job is sent to an instance\n                    which is available when the job is attempted. Requests can only\n                    be sent to a specific instance if [manual scaling is used in App\n                    Engine Standard](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine?hl=en_US#scaling_types_and_instance_classes).\n                    App Engine Flex does not support instances. For more information,\n                    see [App Engine Standard request routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)\n                    and [App Engine Flex request routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).\n                service:\n                  type: string\n                  x-dcl-go-name: Service\n                  description: App service. By default, the job is sent to the service\n                    which is the default service when the job is attempted.\n                version:\n                  type: string\n                  x-dcl-go-name: Version\n                  description: App version. By default, the job is sent to the version\n                    which is the default version when the job is attempted.\n            body:\n              type: string\n              x-dcl-go-name: Body\n              description: Body. HTTP request body. A request body is allowed only\n                if the HTTP method is POST or PUT. It will result in invalid argument\n                error to set a body on a job with an incompatible HttpMethod.\n            headers:\n              type: object\n              additionalProperties:\n                type: string\n              x-dcl-go-name: Headers\n              description: 'HTTP request headers. This map contains the header field\n                names and values. Headers can be set when the job is created. Cloud\n                Scheduler sets some headers to default values: * `User-Agent`: By\n                default, this header is `\"App Engine-Google; (+http://code.google.com/appengine)\"`.\n                This header can be modified, but Cloud Scheduler will append `\"App\n                Engine-Google; (+http://code.google.com/appengine)\"` to the modified\n                `User-Agent`. * `X-CloudScheduler`: This header will be set to true.\n                The headers below are output only. They cannot be set or overridden:\n                * `X-Google-*`: For Google internal use only. * `X-App Engine-*`:\n                For Google internal use only. In addition, some App Engine headers,\n                which contain job-specific information, are also be sent to the job\n                handler.'\n            httpMethod:\n              type: string\n              x-dcl-go-name: HttpMethod\n              x-dcl-go-type: JobAppEngineHttpTargetHttpMethodEnum\n              description: 'The HTTP method to use for the request. PATCH and OPTIONS\n                are not permitted. Possible values: HTTP_METHOD_UNSPECIFIED, POST,\n                GET, HEAD, PUT, DELETE, PATCH, OPTIONS'\n              enum:\n              - HTTP_METHOD_UNSPECIFIED\n              - POST\n              - GET\n              - HEAD\n              - PUT\n              - DELETE\n              - PATCH\n              - OPTIONS\n            relativeUri:\n              type: string\n              x-dcl-go-name: RelativeUri\n              description: The relative URI. The relative URL must begin with \"/\"\n                and must be a valid HTTP relative URL. It can contain a path, query\n                string arguments, and `#` fragments. If the relative URL is empty,\n                then the root path \"/\" will be used. No spaces are allowed, and the\n                maximum length allowed is 2083 characters.\n        attemptDeadline:\n          type: string\n          x-dcl-go-name: AttemptDeadline\n          description: 'The deadline for job attempts. If the request handler does\n            not respond by this deadline then the request is cancelled and the attempt\n            is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be\n            viewed in execution logs. Cloud Scheduler will retry the job according\n            to the RetryConfig. The allowed duration for this deadline is: * For HTTP\n            targets, between 15 seconds and 30 minutes. * For App Engine HTTP targets,\n            between 15 seconds and 24 hours.'\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optionally caller-specified in CreateJob or UpdateJob. A human-readable\n            description for the job. This string must not contain more than 500 characters.\n        httpTarget:\n          type: object\n          x-dcl-go-name: HttpTarget\n          x-dcl-go-type: JobHttpTarget\n          description: HTTP target.\n          x-dcl-conflicts:\n          - pubsubTarget\n          - appEngineHttpTarget\n          required:\n          - uri\n          properties:\n            body:\n              type: string\n              x-dcl-go-name: Body\n              description: HTTP request body. A request body is allowed only if the\n                HTTP method is POST, PUT, or PATCH. It is an error to set body on\n                a job with an incompatible HttpMethod.\n            headers:\n              type: object\n              additionalProperties:\n                type: string\n              x-dcl-go-name: Headers\n              description: 'The user can specify HTTP request headers to send with\n                the job''s HTTP request. This map contains the header field names\n                and values. Repeated headers are not supported, but a header value\n                can contain commas. These headers represent a subset of the headers\n                that will accompany the job''s HTTP request. Some HTTP request headers\n                will be ignored or replaced. A partial list of headers that will be\n                ignored or replaced is below: - Host: This will be computed by Cloud\n                Scheduler and derived from uri. * `Content-Length`: This will be computed\n                by Cloud Scheduler. * `User-Agent`: This will be set to `\"Google-Cloud-Scheduler\"`.\n                * `X-Google-*`: Google internal use only. * `X-appengine-*`: Google\n                internal use only. The total size of headers must be less than 80KB.'\n            httpMethod:\n              type: string\n              x-dcl-go-name: HttpMethod\n              x-dcl-go-type: JobHttpTargetHttpMethodEnum\n              description: 'Which HTTP method to use for the request. Possible values:\n                HTTP_METHOD_UNSPECIFIED, POST, GET, HEAD, PUT, DELETE, PATCH, OPTIONS'\n              enum:\n              - HTTP_METHOD_UNSPECIFIED\n              - POST\n              - GET\n              - HEAD\n              - PUT\n              - DELETE\n              - PATCH\n              - OPTIONS\n            oauthToken:\n              type: object\n              x-dcl-go-name: OAuthToken\n              x-dcl-go-type: JobHttpTargetOAuthToken\n              description: If specified, an [OAuth token](https://developers.google.com/identity/protocols/OAuth2)\n                will be generated and attached as an `Authorization` header in the\n                HTTP request. This type of authorization should generally only be\n                used when calling Google APIs hosted on *.googleapis.com.\n              properties:\n                scope:\n                  type: string\n                  x-dcl-go-name: Scope\n                  description: OAuth scope to be used for generating OAuth access\n                    token. If not specified, \"https://www.googleapis.com/auth/cloud-platform\"\n                    will be used.\n                serviceAccountEmail:\n                  type: string\n                  x-dcl-go-name: ServiceAccountEmail\n                  description: '[Service account email](https://cloud.google.com/iam/docs/service-accounts)\n                    to be used for generating OAuth token. The service account must\n                    be within the same project as the job. The caller must have iam.serviceAccounts.actAs\n                    permission for the service account.'\n                  x-dcl-references:\n                  - resource: Iam/ServiceAccount\n                    field: email\n            oidcToken:\n              type: object\n              x-dcl-go-name: OidcToken\n              x-dcl-go-type: JobHttpTargetOidcToken\n              description: If specified, an [OIDC](https://developers.google.com/identity/protocols/OpenIDConnect)\n                token will be generated and attached as an `Authorization` header\n                in the HTTP request. This type of authorization can be used for many\n                scenarios, including calling Cloud Run, or endpoints where you intend\n                to validate the token yourself.\n              properties:\n                audience:\n                  type: string\n                  x-dcl-go-name: Audience\n                  description: Audience to be used when generating OIDC token. If\n                    not specified, the URI specified in target will be used.\n                serviceAccountEmail:\n                  type: string\n                  x-dcl-go-name: ServiceAccountEmail\n                  description: '[Service account email](https://cloud.google.com/iam/docs/service-accounts)\n                    to be used for generating OIDC token. The service account must\n                    be within the same project as the job. The caller must have iam.serviceAccounts.actAs\n                    permission for the service account.'\n                  x-dcl-references:\n                  - resource: Iam/ServiceAccount\n                    field: email\n            uri:\n              type: string\n              x-dcl-go-name: Uri\n              description: 'Required. The full URI path that the request will be sent\n                to. This string must begin with either \"http://\" or \"https://\". Some\n                examples of valid values for uri are: `http://acme.com` and `https://acme.com/sales:8080`.\n                Cloud Scheduler will encode some characters for safety and compatibility.\n                The maximum allowed URL length is 2083 characters after encoding.'\n        lastAttemptTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: LastAttemptTime\n          readOnly: true\n          description: Output only. The time the last job attempt started.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Optionally caller-specified in CreateJob, after which it becomes\n            output only. The job name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`.\n            * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens\n            (-), colons (:), or periods (.). For more information, see [Identifying\n            projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)\n            * `LOCATION_ID` is the canonical ID for the job''s location. The list\n            of available locations can be obtained by calling ListLocations. For more\n            information, see https://cloud.google.com/about/locations/. * `JOB_ID`\n            can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or\n            underscores (_). The maximum length is 500 characters.'\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        pubsubTarget:\n          type: object\n          x-dcl-go-name: PubsubTarget\n          x-dcl-go-type: JobPubsubTarget\n          description: Pub/Sub target.\n          x-dcl-conflicts:\n          - appEngineHttpTarget\n          - httpTarget\n          required:\n          - topicName\n          properties:\n            attributes:\n              type: object\n              additionalProperties:\n                type: string\n              x-dcl-go-name: Attributes\n              description: Attributes for PubsubMessage. Pubsub message must contain\n                either non-empty data, or at least one attribute.\n            data:\n              type: string\n              x-dcl-go-name: Data\n              description: The message payload for PubsubMessage. Pubsub message must\n                contain either non-empty data, or at least one attribute.\n            topicName:\n              type: string\n              x-dcl-go-name: TopicName\n              description: Required. The name of the Cloud Pub/Sub topic to which\n                messages will be published when a job is delivered. The topic name\n                must be in the same format as required by Pub/Sub's [PublishRequest.name](https://cloud.google.com/pubsub/docs/reference/rpc/google.pubsub.v1#publishrequest),\n                for example `projects/PROJECT_ID/topics/TOPIC_ID`. The topic must\n                be in the same project as the Cloud Scheduler job.\n              x-dcl-references:\n              - resource: Pubsub/Topic\n                field: name\n        retryConfig:\n          type: object\n          x-dcl-go-name: RetryConfig\n          x-dcl-go-type: JobRetryConfig\n          description: Settings that determine the retry behavior.\n          properties:\n            maxBackoffDuration:\n              type: string\n              x-dcl-go-name: MaxBackoffDuration\n              description: The maximum amount of time to wait before retrying a job\n                after it fails. The default value of this field is 1 hour.\n            maxDoublings:\n              type: integer\n              format: int64\n              x-dcl-go-name: MaxDoublings\n              description: The time between retries will double `max_doublings` times.\n                A job's retry interval starts at min_backoff_duration, then doubles\n                `max_doublings` times, then increases linearly, and finally retries\n                at intervals of max_backoff_duration up to retry_count times. For\n                example, if min_backoff_duration is 10s, max_backoff_duration is 300s,\n                and `max_doublings` is 3, then the a job will first be retried in\n                10s. The retry interval will double three times, and then increase\n                linearly by 2^3 * 10s. Finally, the job will retry at intervals of\n                max_backoff_duration until the job has been attempted retry_count\n                times. Thus, the requests will retry at 10s, 20s, 40s, 80s, 160s,\n                240s, 300s, 300s, .... The default value of this field is 5.\n            maxRetryDuration:\n              type: string\n              x-dcl-go-name: MaxRetryDuration\n              description: The time limit for retrying a failed job, measured from\n                time when an execution was first attempted. If specified with retry_count,\n                the job will be retried until both limits are reached. The default\n                value for max_retry_duration is zero, which means retry duration is\n                unlimited.\n            minBackoffDuration:\n              type: string\n              x-dcl-go-name: MinBackoffDuration\n              description: The minimum amount of time to wait before retrying a job\n                after it fails. The default value of this field is 5 seconds.\n            retryCount:\n              type: integer\n              format: int64\n              x-dcl-go-name: RetryCount\n              description: The number of attempts that the system will make to run\n                a job using the exponential backoff procedure described by max_doublings.\n                The default value of retry_count is zero. If retry_count is zero,\n                a job attempt will *not* be retried if it fails. Instead the Cloud\n                Scheduler system will wait for the next scheduled execution time.\n                If retry_count is set to a non-zero number then Cloud Scheduler will\n                retry failed attempts, using exponential backoff, retry_count times,\n                or until the next scheduled execution time, whichever comes first.\n                Values greater than 5 and negative values are not allowed.\n        schedule:\n          type: string\n          x-dcl-go-name: Schedule\n          description: 'Required, except when used with UpdateJob. Describes the schedule\n            on which the job will be executed. The schedule can be either of the following\n            types: * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview) * English-like\n            [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules)\n            As a general rule, execution `n + 1` of a job will not begin until execution\n            `n` has finished. Cloud Scheduler will never allow two simultaneously\n            outstanding executions. For example, this implies that if the `n+1`th\n            execution is scheduled to run at 16:00 but the `n`th execution takes until\n            16:15, the `n+1`th execution will not start until `16:15`. A scheduled\n            start time will be delayed if the previous execution has not ended when\n            its scheduled time occurs. If retry_count > 0 and a job attempt fails,\n            the job will be tried a total of retry_count times, with exponential backoff,\n            until the next scheduled start time.'\n        scheduleTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: ScheduleTime\n          readOnly: true\n          description: Output only. The next time the job is scheduled. Note that\n            this may be a retry of a previously failed attempt or the next execution\n            time according to the schedule.\n          x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: JobStateEnum\n          readOnly: true\n          description: 'Output only. State of the job. Possible values: STATE_UNSPECIFIED,\n            ENABLED, PAUSED, DISABLED, UPDATE_FAILED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - ENABLED\n          - PAUSED\n          - DISABLED\n          - UPDATE_FAILED\n        status:\n          type: object\n          x-dcl-go-name: Status\n          x-dcl-go-type: JobStatus\n          readOnly: true\n          description: Output only. The response from the target for the last attempted\n            execution.\n          x-kubernetes-immutable: true\n          properties:\n            code:\n              type: integer\n              format: int64\n              x-dcl-go-name: Code\n              description: The status code, which should be an enum value of google.rpc.Code.\n              x-kubernetes-immutable: true\n            details:\n              type: array\n              x-dcl-go-name: Details\n              description: A list of messages that carry the error details. There\n                is a common set of message types for APIs to use.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: JobStatusDetails\n                properties:\n                  typeUrl:\n                    type: string\n                    x-dcl-go-name: TypeUrl\n                    description: 'A URL/resource name that uniquely identifies the\n                      type of the serialized protocol buffer message. This string\n                      must contain at least one \"/\" character. The last segment of\n                      the URL''s path must represent the fully qualified name of the\n                      type (as in `path/google.protobuf.Duration`). The name should\n                      be in a canonical form (e.g., leading \".\" is not accepted).\n                      In practice, teams usually precompile into the binary all types\n                      that they expect it to use in the context of Any. However, for\n                      URLs which use the scheme `http`, `https`, or no scheme, one\n                      can optionally set up a type server that maps type URLs to message\n                      definitions as follows: * If no scheme is provided, `https`\n                      is assumed. * An HTTP GET on the URL must yield a google.protobuf.Type\n                      value in binary format, or produce an error. * Applications\n                      are allowed to cache lookup results based on the URL, or have\n                      them precompiled into a binary to avoid any lookup. Therefore,\n                      binary compatibility needs to be preserved on changes to types.\n                      (Use versioned type names to manage breaking changes.) Note:\n                      this functionality is not currently available in the official\n                      protobuf release, and it is not used for type URLs beginning\n                      with type.googleapis.com. Schemes other than `http`, `https`\n                      (or the empty scheme) might be used with implementation specific\n                      semantics.'\n                    x-kubernetes-immutable: true\n                  value:\n                    type: string\n                    x-dcl-go-name: Value\n                    description: Must be a valid serialized protocol buffer of the\n                      above specified type.\n                    x-kubernetes-immutable: true\n            message:\n              type: string\n              x-dcl-go-name: Message\n              description: A developer-facing error message, which should be in English.\n                Any user-facing error message should be localized and sent in the\n                google.rpc.Status.details field, or localized by the client.\n              x-kubernetes-immutable: true\n        timeZone:\n          type: string\n          x-dcl-go-name: TimeZone\n          description: Specifies the time zone to be used in interpreting schedule.\n            The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database).\n            Note that some time zones include a provision for daylight savings time.\n            The rules for daylight saving time are determined by the chosen tz. For\n            UTC use the string \"utc\". If a time zone is not specified, the default\n            will be in UTC (also known as GMT).\n          x-dcl-server-default: true\n        userUpdateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UserUpdateTime\n          readOnly: true\n          description: Output only. The creation time of the job.\n          x-kubernetes-immutable: true\n")

// 27604 bytes
// MD5: de916dc2ac4c7dbfec76420434698ae8
