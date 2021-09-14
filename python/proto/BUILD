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

load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "empty_proto",
    srcs = ["empty.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "empty_go_proto",
    importpath = "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto",
    proto = ":empty_proto",
    visibility = ["//visibility:public"],
)

load("//:connector_rules.bzl", "python_protos")

python_protos(packages = ["ud/graphite/mmv2/connector","assuredworkloads","assuredworkloads/alpha","assuredworkloads/beta","bigqueryreservation","bigqueryreservation/beta","binaryauthorization","binaryauthorization/beta","cloudbuild/beta","cloudbuild/alpha","cloudresourcemanager","cloudresourcemanager/beta","cloudscheduler","cloudscheduler/beta","compute","compute/beta","compute/alpha","containeranalysis","containeranalysis/beta","containeraws/alpha","containerazure/alpha","datafusion/beta","dataproc","dataproc/alpha","dataproc/beta","eventarc","eventarc/alpha","eventarc/beta","filestore","filestore/beta","gameservices","gameservices/beta","gkehub/beta","gkehub/alpha","iam","iam/beta","iap","iap/beta","identitytoolkit","identitytoolkit/beta","logging","logging/beta","monitoring","monitoring/beta","networksecurity/alpha","networksecurity/beta","networkservices/alpha","networkservices/beta","orgpolicy","orgpolicy/beta","osconfig/alpha","osconfig/beta","pubsub","pubsub/beta","storage","storage/beta","privateca","privateca/alpha","privateca/beta"])
