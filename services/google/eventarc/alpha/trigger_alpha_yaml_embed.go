// Copyright 2021 Google LLC. All Rights Reserved.
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
// gen_go_data -package alpha -var YAML_trigger blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/eventarc/alpha/trigger.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/eventarc/alpha/trigger.yaml
var YAML_trigger = []byte("info:\n  title: Eventarc/Trigger\n  description: DCL Specification for the Eventarc Trigger resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Trigger\n    parameters:\n    - name: Trigger\n      required: true\n      description: A full instance of a Trigger\n  apply:\n    description: The function used to apply information about a Trigger\n    parameters:\n    - name: Trigger\n      required: true\n      description: A full instance of a Trigger\n  delete:\n    description: The function used to delete a Trigger\n    parameters:\n    - name: Trigger\n      required: true\n      description: A full instance of a Trigger\n  deleteAll:\n    description: The function used to delete all Trigger\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Trigger\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Trigger:\n      title: Trigger\n      x-dcl-id: projects/{{project}}/locations/{{location}}/triggers/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      type: object\n      required:\n      - name\n      - matchingCriteria\n      - destination\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The creation time.\n          x-kubernetes-immutable: true\n        destination:\n          type: object\n          x-dcl-go-name: Destination\n          x-dcl-go-type: TriggerDestination\n          description: Required. Destination specifies where the events should be\n            sent to.\n          properties:\n            cloudFunction:\n              type: string\n              x-dcl-go-name: CloudFunction\n              description: 'The Cloud Function resource name. Only Cloud Functions\n                V2 is supported. Format: projects/{project}/locations/{location}/functions/{function}'\n              x-dcl-references:\n              - resource: Cloudfunctions/CloudFunction\n                field: selfLink\n            cloudRunService:\n              type: object\n              x-dcl-go-name: CloudRunService\n              x-dcl-go-type: TriggerDestinationCloudRunService\n              description: Cloud Run fully-managed service that receives the events.\n                The service should be running in the same project of the trigger.\n              required:\n              - service\n              - region\n              properties:\n                path:\n                  type: string\n                  x-dcl-go-name: Path\n                  description: 'Optional. The relative path on the Cloud Run service\n                    the events should be sent to. The value must conform to the definition\n                    of URI path segment (section 3.3 of RFC2396). Examples: \"/route\",\n                    \"route\", \"route/subroute\".'\n                region:\n                  type: string\n                  x-dcl-go-name: Region\n                  description: Required. The region the Cloud Run service is deployed\n                    in.\n                service:\n                  type: string\n                  x-dcl-go-name: Service\n                  description: Required. The name of the Cloud Run service being addressed.\n                    See https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services.\n                    Only services located in the same project of the trigger object\n                    can be addressed.\n                  x-dcl-references:\n                  - resource: Run/Service\n                    field: selfLink\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Output only. This checksum is computed by the server based\n            on the value of other fields, and may be sent only on create requests\n            to ensure the client has an up-to-date value before proceeding.\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. User labels attached to the triggers that can be\n            used to group resources.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        matchingCriteria:\n          type: array\n          x-dcl-go-name: MatchingCriteria\n          description: Required. null The list of filters that applies to event attributes.\n            Only events that match all the provided filters will be sent to the destination.\n          x-dcl-send-empty: true\n          x-dcl-list-type: set\n          items:\n            type: object\n            x-dcl-go-type: TriggerMatchingCriteria\n            required:\n            - attribute\n            - value\n            properties:\n              attribute:\n                type: string\n                x-dcl-go-name: Attribute\n                description: Required. The name of a CloudEvents attribute. Currently,\n                  only a subset of attributes are supported for filtering. All triggers\n                  MUST provide a filter for the 'type' attribute.\n              value:\n                type: string\n                x-dcl-go-name: Value\n                description: Required. The value for the attribute.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. The resource name of the trigger. Must be unique\n            within the location on the project and must be in \\\\\\`projects/{project}/locations/{location}/triggers/{trigger}\\\\\\`\n            format.\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        serviceAccount:\n          type: string\n          x-dcl-go-name: ServiceAccount\n          description: Optional. The IAM service account email associated with the\n            trigger. The service account represents the identity of the trigger. The\n            principal who calls this API must have \\\\\\`iam.serviceAccounts.actAs\\\\\\`\n            permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa\\\\\\_common\n            for more information. For Cloud Run destinations, this service account\n            is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account\n            for information on how to invoke authenticated Cloud Run services. In\n            order to create Audit Log triggers, the service account should also have\n            \\\\\\`roles/eventarc.eventReceiver\\\\\\` IAM role.\n          x-dcl-references:\n          - resource: Iam/ServiceAccount\n            field: email\n        transport:\n          type: object\n          x-dcl-go-name: Transport\n          x-dcl-go-type: TriggerTransport\n          description: Optional. In order to deliver messages, Eventarc may use other\n            GCP products as transport intermediary. This field contains a reference\n            to that transport intermediary. This information can be used for debugging\n            purposes.\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n          properties:\n            pubsub:\n              type: object\n              x-dcl-go-name: Pubsub\n              x-dcl-go-type: TriggerTransportPubsub\n              description: The Pub/Sub topic and subscription used by Eventarc as\n                delivery intermediary.\n              x-kubernetes-immutable: true\n              properties:\n                subscription:\n                  type: string\n                  x-dcl-go-name: Subscription\n                  readOnly: true\n                  description: 'Output only. The name of the Pub/Sub subscription\n                    created and managed by Eventarc system as a transport for the\n                    event delivery. Format: \\\\\\\\\\\\\\\\\\\\`projects/{PROJECT\\\\\\\\\\\\\\\\\\\\_ID}/subscriptions/{SUBSCRIPTION\\\\\\\\\\\\\\\\\\\\_NAME}\\\\\\\\\\\\\\\\\\\\`.'\n                  x-kubernetes-immutable: true\n                topic:\n                  type: string\n                  x-dcl-go-name: Topic\n                  description: 'Optional. The name of the Pub/Sub topic created and\n                    managed by Eventarc system as a transport for the event delivery.\n                    Format: \\\\\\\\\\\\\\\\\\\\`projects/{PROJECT\\\\\\\\\\\\\\\\\\\\_ID}/topics/{TOPIC\\\\\\\\\\\\\\\\\\\\_NAME}\\\\\\\\\\\\\\\\\\\\`.\n                    You may set an existing topic for triggers of the type \\\\\\\\\\\\\\\\\\\\`google.cloud.pubsub.topic.v1.messagePublished\\\\\\\\\\\\\\\\\\\\`\n                    only. The topic you provide here will not be deleted by Eventarc\n                    at trigger deletion.'\n                  x-kubernetes-immutable: true\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. Server assigned unique identifier for the trigger.\n            The value is a UUID4 string and guaranteed to remain unchanged until the\n            resource is deleted.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The last-modified time.\n          x-kubernetes-immutable: true\n")

// 9934 bytes
// MD5: 6382480e475d3770777aca4bb73d9806
