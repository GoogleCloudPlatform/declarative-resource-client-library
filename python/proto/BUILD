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

python_protos(packages = ["ud/graphite/mmv2/connector","accesscontextmanager","appengine","bigqueryconnection","bigqueryreservation","binaryauthorization","cloudbilling","cloudbuild","cloudfunctions","cloudresourcemanager","cloudscheduler","compute","compute/beta","container","container/beta","containeranalysis","containeranalysis/beta","datafusion/beta","dataproc","dataproc/beta","datastore","dns","file","eventarc/beta","gameservices","gkehub/beta","logging","iam","iap","identitytoolkit","monitoring","osconfig/beta","pubsub","pubsublite","redis","redis/beta","run","runtimeconfig","servicenetworking","sourcerepo","serviceusage","spanner","sql/beta","storage","tpu","vpcaccess"])
