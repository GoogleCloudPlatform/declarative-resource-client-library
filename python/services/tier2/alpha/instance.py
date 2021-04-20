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
from google3.cloud.graphite.mmv2.services.google.tier2 import instance_pb2
from google3.cloud.graphite.mmv2.services.google.tier2 import instance_pb2_grpc

from typing import List


class Instance(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        labels: dict = None,
        zone: str = None,
        sku: dict = None,
        authorized_network_id: str = None,
        reserved_ip_range: str = None,
        host_name: str = None,
        port_number: int = None,
        current_zone: str = None,
        creation_time: str = None,
        state: str = None,
        status_message: str = None,
        extra_field: str = None,
        preprocess_create_recipe: dict = None,
        initiate_create_recipe: dict = None,
        create_recipe: dict = None,
        delete_recipe: dict = None,
        update_recipe: dict = None,
        preprocess_reset_recipe: dict = None,
        initiate_reset_recipe: dict = None,
        reset_recipe: dict = None,
        preprocess_repair_recipe: dict = None,
        initiate_repair_recipe: dict = None,
        repair_recipe: dict = None,
        preprocess_delete_recipe: dict = None,
        initiate_delete_recipe: dict = None,
        preprocess_update_recipe: dict = None,
        initiate_update_recipe: dict = None,
        preprocess_freeze_recipe: dict = None,
        freeze_recipe: dict = None,
        preprocess_unfreeze_recipe: dict = None,
        unfreeze_recipe: dict = None,
        readonly_recipe: dict = None,
        enable_call_history: bool = None,
        history: list = None,
        public_resource_view_override: str = None,
        project: str = None,
        location: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.labels = labels
        self.zone = zone
        self.sku = sku
        self.authorized_network_id = authorized_network_id
        self.reserved_ip_range = reserved_ip_range
        self.host_name = host_name
        self.port_number = port_number
        self.current_zone = current_zone
        self.creation_time = creation_time
        self.state = state
        self.status_message = status_message
        self.extra_field = extra_field
        self.preprocess_create_recipe = preprocess_create_recipe
        self.initiate_create_recipe = initiate_create_recipe
        self.create_recipe = create_recipe
        self.delete_recipe = delete_recipe
        self.update_recipe = update_recipe
        self.preprocess_reset_recipe = preprocess_reset_recipe
        self.initiate_reset_recipe = initiate_reset_recipe
        self.reset_recipe = reset_recipe
        self.preprocess_repair_recipe = preprocess_repair_recipe
        self.initiate_repair_recipe = initiate_repair_recipe
        self.repair_recipe = repair_recipe
        self.preprocess_delete_recipe = preprocess_delete_recipe
        self.initiate_delete_recipe = initiate_delete_recipe
        self.preprocess_update_recipe = preprocess_update_recipe
        self.initiate_update_recipe = initiate_update_recipe
        self.preprocess_freeze_recipe = preprocess_freeze_recipe
        self.freeze_recipe = freeze_recipe
        self.preprocess_unfreeze_recipe = preprocess_unfreeze_recipe
        self.unfreeze_recipe = unfreeze_recipe
        self.readonly_recipe = readonly_recipe
        self.enable_call_history = enable_call_history
        self.history = history
        self.public_resource_view_override = public_resource_view_override
        self.project = project
        self.location = location
        self.service_account_file = service_account_file

    def apply(self):
        stub = instance_pb2_grpc.Tier2AlphaInstanceServiceStub(channel.Channel())
        request = instance_pb2.ApplyTier2AlphaInstanceRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        if InstanceSku.to_proto(self.sku):
            request.resource.sku.CopyFrom(InstanceSku.to_proto(self.sku))
        else:
            request.resource.ClearField("sku")
        if Primitive.to_proto(self.authorized_network_id):
            request.resource.authorized_network_id = Primitive.to_proto(
                self.authorized_network_id
            )

        if Primitive.to_proto(self.reserved_ip_range):
            request.resource.reserved_ip_range = Primitive.to_proto(
                self.reserved_ip_range
            )

        if Primitive.to_proto(self.host_name):
            request.resource.host_name = Primitive.to_proto(self.host_name)

        if Primitive.to_proto(self.port_number):
            request.resource.port_number = Primitive.to_proto(self.port_number)

        if Primitive.to_proto(self.current_zone):
            request.resource.current_zone = Primitive.to_proto(self.current_zone)

        if Primitive.to_proto(self.creation_time):
            request.resource.creation_time = Primitive.to_proto(self.creation_time)

        if InstanceStateEnum.to_proto(self.state):
            request.resource.state = InstanceStateEnum.to_proto(self.state)

        if Primitive.to_proto(self.status_message):
            request.resource.status_message = Primitive.to_proto(self.status_message)

        if Primitive.to_proto(self.extra_field):
            request.resource.extra_field = Primitive.to_proto(self.extra_field)

        if InstancePreprocessCreateRecipe.to_proto(self.preprocess_create_recipe):
            request.resource.preprocess_create_recipe.CopyFrom(
                InstancePreprocessCreateRecipe.to_proto(self.preprocess_create_recipe)
            )
        else:
            request.resource.ClearField("preprocess_create_recipe")
        if InstanceInitiateCreateRecipe.to_proto(self.initiate_create_recipe):
            request.resource.initiate_create_recipe.CopyFrom(
                InstanceInitiateCreateRecipe.to_proto(self.initiate_create_recipe)
            )
        else:
            request.resource.ClearField("initiate_create_recipe")
        if InstanceCreateRecipe.to_proto(self.create_recipe):
            request.resource.create_recipe.CopyFrom(
                InstanceCreateRecipe.to_proto(self.create_recipe)
            )
        else:
            request.resource.ClearField("create_recipe")
        if InstanceDeleteRecipe.to_proto(self.delete_recipe):
            request.resource.delete_recipe.CopyFrom(
                InstanceDeleteRecipe.to_proto(self.delete_recipe)
            )
        else:
            request.resource.ClearField("delete_recipe")
        if InstanceUpdateRecipe.to_proto(self.update_recipe):
            request.resource.update_recipe.CopyFrom(
                InstanceUpdateRecipe.to_proto(self.update_recipe)
            )
        else:
            request.resource.ClearField("update_recipe")
        if InstancePreprocessResetRecipe.to_proto(self.preprocess_reset_recipe):
            request.resource.preprocess_reset_recipe.CopyFrom(
                InstancePreprocessResetRecipe.to_proto(self.preprocess_reset_recipe)
            )
        else:
            request.resource.ClearField("preprocess_reset_recipe")
        if InstanceInitiateResetRecipe.to_proto(self.initiate_reset_recipe):
            request.resource.initiate_reset_recipe.CopyFrom(
                InstanceInitiateResetRecipe.to_proto(self.initiate_reset_recipe)
            )
        else:
            request.resource.ClearField("initiate_reset_recipe")
        if InstanceResetRecipe.to_proto(self.reset_recipe):
            request.resource.reset_recipe.CopyFrom(
                InstanceResetRecipe.to_proto(self.reset_recipe)
            )
        else:
            request.resource.ClearField("reset_recipe")
        if InstancePreprocessRepairRecipe.to_proto(self.preprocess_repair_recipe):
            request.resource.preprocess_repair_recipe.CopyFrom(
                InstancePreprocessRepairRecipe.to_proto(self.preprocess_repair_recipe)
            )
        else:
            request.resource.ClearField("preprocess_repair_recipe")
        if InstanceInitiateRepairRecipe.to_proto(self.initiate_repair_recipe):
            request.resource.initiate_repair_recipe.CopyFrom(
                InstanceInitiateRepairRecipe.to_proto(self.initiate_repair_recipe)
            )
        else:
            request.resource.ClearField("initiate_repair_recipe")
        if InstanceRepairRecipe.to_proto(self.repair_recipe):
            request.resource.repair_recipe.CopyFrom(
                InstanceRepairRecipe.to_proto(self.repair_recipe)
            )
        else:
            request.resource.ClearField("repair_recipe")
        if InstancePreprocessDeleteRecipe.to_proto(self.preprocess_delete_recipe):
            request.resource.preprocess_delete_recipe.CopyFrom(
                InstancePreprocessDeleteRecipe.to_proto(self.preprocess_delete_recipe)
            )
        else:
            request.resource.ClearField("preprocess_delete_recipe")
        if InstanceInitiateDeleteRecipe.to_proto(self.initiate_delete_recipe):
            request.resource.initiate_delete_recipe.CopyFrom(
                InstanceInitiateDeleteRecipe.to_proto(self.initiate_delete_recipe)
            )
        else:
            request.resource.ClearField("initiate_delete_recipe")
        if InstancePreprocessUpdateRecipe.to_proto(self.preprocess_update_recipe):
            request.resource.preprocess_update_recipe.CopyFrom(
                InstancePreprocessUpdateRecipe.to_proto(self.preprocess_update_recipe)
            )
        else:
            request.resource.ClearField("preprocess_update_recipe")
        if InstanceInitiateUpdateRecipe.to_proto(self.initiate_update_recipe):
            request.resource.initiate_update_recipe.CopyFrom(
                InstanceInitiateUpdateRecipe.to_proto(self.initiate_update_recipe)
            )
        else:
            request.resource.ClearField("initiate_update_recipe")
        if InstancePreprocessFreezeRecipe.to_proto(self.preprocess_freeze_recipe):
            request.resource.preprocess_freeze_recipe.CopyFrom(
                InstancePreprocessFreezeRecipe.to_proto(self.preprocess_freeze_recipe)
            )
        else:
            request.resource.ClearField("preprocess_freeze_recipe")
        if InstanceFreezeRecipe.to_proto(self.freeze_recipe):
            request.resource.freeze_recipe.CopyFrom(
                InstanceFreezeRecipe.to_proto(self.freeze_recipe)
            )
        else:
            request.resource.ClearField("freeze_recipe")
        if InstancePreprocessUnfreezeRecipe.to_proto(self.preprocess_unfreeze_recipe):
            request.resource.preprocess_unfreeze_recipe.CopyFrom(
                InstancePreprocessUnfreezeRecipe.to_proto(
                    self.preprocess_unfreeze_recipe
                )
            )
        else:
            request.resource.ClearField("preprocess_unfreeze_recipe")
        if InstanceUnfreezeRecipe.to_proto(self.unfreeze_recipe):
            request.resource.unfreeze_recipe.CopyFrom(
                InstanceUnfreezeRecipe.to_proto(self.unfreeze_recipe)
            )
        else:
            request.resource.ClearField("unfreeze_recipe")
        if InstanceReadonlyRecipe.to_proto(self.readonly_recipe):
            request.resource.readonly_recipe.CopyFrom(
                InstanceReadonlyRecipe.to_proto(self.readonly_recipe)
            )
        else:
            request.resource.ClearField("readonly_recipe")
        if Primitive.to_proto(self.enable_call_history):
            request.resource.enable_call_history = Primitive.to_proto(
                self.enable_call_history
            )

        if InstanceHistoryArray.to_proto(self.history):
            request.resource.history.extend(InstanceHistoryArray.to_proto(self.history))
        if Primitive.to_proto(self.public_resource_view_override):
            request.resource.public_resource_view_override = Primitive.to_proto(
                self.public_resource_view_override
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        request.service_account_file = self.service_account_file

        response = stub.ApplyTier2AlphaInstance(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.labels = Primitive.from_proto(response.labels)
        self.zone = Primitive.from_proto(response.zone)
        self.sku = InstanceSku.from_proto(response.sku)
        self.authorized_network_id = Primitive.from_proto(
            response.authorized_network_id
        )
        self.reserved_ip_range = Primitive.from_proto(response.reserved_ip_range)
        self.host_name = Primitive.from_proto(response.host_name)
        self.port_number = Primitive.from_proto(response.port_number)
        self.current_zone = Primitive.from_proto(response.current_zone)
        self.creation_time = Primitive.from_proto(response.creation_time)
        self.state = InstanceStateEnum.from_proto(response.state)
        self.status_message = Primitive.from_proto(response.status_message)
        self.extra_field = Primitive.from_proto(response.extra_field)
        self.preprocess_create_recipe = InstancePreprocessCreateRecipe.from_proto(
            response.preprocess_create_recipe
        )
        self.initiate_create_recipe = InstanceInitiateCreateRecipe.from_proto(
            response.initiate_create_recipe
        )
        self.create_recipe = InstanceCreateRecipe.from_proto(response.create_recipe)
        self.delete_recipe = InstanceDeleteRecipe.from_proto(response.delete_recipe)
        self.update_recipe = InstanceUpdateRecipe.from_proto(response.update_recipe)
        self.preprocess_reset_recipe = InstancePreprocessResetRecipe.from_proto(
            response.preprocess_reset_recipe
        )
        self.initiate_reset_recipe = InstanceInitiateResetRecipe.from_proto(
            response.initiate_reset_recipe
        )
        self.reset_recipe = InstanceResetRecipe.from_proto(response.reset_recipe)
        self.preprocess_repair_recipe = InstancePreprocessRepairRecipe.from_proto(
            response.preprocess_repair_recipe
        )
        self.initiate_repair_recipe = InstanceInitiateRepairRecipe.from_proto(
            response.initiate_repair_recipe
        )
        self.repair_recipe = InstanceRepairRecipe.from_proto(response.repair_recipe)
        self.preprocess_delete_recipe = InstancePreprocessDeleteRecipe.from_proto(
            response.preprocess_delete_recipe
        )
        self.initiate_delete_recipe = InstanceInitiateDeleteRecipe.from_proto(
            response.initiate_delete_recipe
        )
        self.preprocess_update_recipe = InstancePreprocessUpdateRecipe.from_proto(
            response.preprocess_update_recipe
        )
        self.initiate_update_recipe = InstanceInitiateUpdateRecipe.from_proto(
            response.initiate_update_recipe
        )
        self.preprocess_freeze_recipe = InstancePreprocessFreezeRecipe.from_proto(
            response.preprocess_freeze_recipe
        )
        self.freeze_recipe = InstanceFreezeRecipe.from_proto(response.freeze_recipe)
        self.preprocess_unfreeze_recipe = InstancePreprocessUnfreezeRecipe.from_proto(
            response.preprocess_unfreeze_recipe
        )
        self.unfreeze_recipe = InstanceUnfreezeRecipe.from_proto(
            response.unfreeze_recipe
        )
        self.readonly_recipe = InstanceReadonlyRecipe.from_proto(
            response.readonly_recipe
        )
        self.enable_call_history = Primitive.from_proto(response.enable_call_history)
        self.history = InstanceHistoryArray.from_proto(response.history)
        self.public_resource_view_override = Primitive.from_proto(
            response.public_resource_view_override
        )
        self.project = Primitive.from_proto(response.project)
        self.location = Primitive.from_proto(response.location)

    def delete(self):
        stub = instance_pb2_grpc.Tier2AlphaInstanceServiceStub(channel.Channel())
        request = instance_pb2.DeleteTier2AlphaInstanceRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if Primitive.to_proto(self.labels):
            request.resource.labels = Primitive.to_proto(self.labels)

        if Primitive.to_proto(self.zone):
            request.resource.zone = Primitive.to_proto(self.zone)

        if InstanceSku.to_proto(self.sku):
            request.resource.sku.CopyFrom(InstanceSku.to_proto(self.sku))
        else:
            request.resource.ClearField("sku")
        if Primitive.to_proto(self.authorized_network_id):
            request.resource.authorized_network_id = Primitive.to_proto(
                self.authorized_network_id
            )

        if Primitive.to_proto(self.reserved_ip_range):
            request.resource.reserved_ip_range = Primitive.to_proto(
                self.reserved_ip_range
            )

        if Primitive.to_proto(self.host_name):
            request.resource.host_name = Primitive.to_proto(self.host_name)

        if Primitive.to_proto(self.port_number):
            request.resource.port_number = Primitive.to_proto(self.port_number)

        if Primitive.to_proto(self.current_zone):
            request.resource.current_zone = Primitive.to_proto(self.current_zone)

        if Primitive.to_proto(self.creation_time):
            request.resource.creation_time = Primitive.to_proto(self.creation_time)

        if InstanceStateEnum.to_proto(self.state):
            request.resource.state = InstanceStateEnum.to_proto(self.state)

        if Primitive.to_proto(self.status_message):
            request.resource.status_message = Primitive.to_proto(self.status_message)

        if Primitive.to_proto(self.extra_field):
            request.resource.extra_field = Primitive.to_proto(self.extra_field)

        if InstancePreprocessCreateRecipe.to_proto(self.preprocess_create_recipe):
            request.resource.preprocess_create_recipe.CopyFrom(
                InstancePreprocessCreateRecipe.to_proto(self.preprocess_create_recipe)
            )
        else:
            request.resource.ClearField("preprocess_create_recipe")
        if InstanceInitiateCreateRecipe.to_proto(self.initiate_create_recipe):
            request.resource.initiate_create_recipe.CopyFrom(
                InstanceInitiateCreateRecipe.to_proto(self.initiate_create_recipe)
            )
        else:
            request.resource.ClearField("initiate_create_recipe")
        if InstanceCreateRecipe.to_proto(self.create_recipe):
            request.resource.create_recipe.CopyFrom(
                InstanceCreateRecipe.to_proto(self.create_recipe)
            )
        else:
            request.resource.ClearField("create_recipe")
        if InstanceDeleteRecipe.to_proto(self.delete_recipe):
            request.resource.delete_recipe.CopyFrom(
                InstanceDeleteRecipe.to_proto(self.delete_recipe)
            )
        else:
            request.resource.ClearField("delete_recipe")
        if InstanceUpdateRecipe.to_proto(self.update_recipe):
            request.resource.update_recipe.CopyFrom(
                InstanceUpdateRecipe.to_proto(self.update_recipe)
            )
        else:
            request.resource.ClearField("update_recipe")
        if InstancePreprocessResetRecipe.to_proto(self.preprocess_reset_recipe):
            request.resource.preprocess_reset_recipe.CopyFrom(
                InstancePreprocessResetRecipe.to_proto(self.preprocess_reset_recipe)
            )
        else:
            request.resource.ClearField("preprocess_reset_recipe")
        if InstanceInitiateResetRecipe.to_proto(self.initiate_reset_recipe):
            request.resource.initiate_reset_recipe.CopyFrom(
                InstanceInitiateResetRecipe.to_proto(self.initiate_reset_recipe)
            )
        else:
            request.resource.ClearField("initiate_reset_recipe")
        if InstanceResetRecipe.to_proto(self.reset_recipe):
            request.resource.reset_recipe.CopyFrom(
                InstanceResetRecipe.to_proto(self.reset_recipe)
            )
        else:
            request.resource.ClearField("reset_recipe")
        if InstancePreprocessRepairRecipe.to_proto(self.preprocess_repair_recipe):
            request.resource.preprocess_repair_recipe.CopyFrom(
                InstancePreprocessRepairRecipe.to_proto(self.preprocess_repair_recipe)
            )
        else:
            request.resource.ClearField("preprocess_repair_recipe")
        if InstanceInitiateRepairRecipe.to_proto(self.initiate_repair_recipe):
            request.resource.initiate_repair_recipe.CopyFrom(
                InstanceInitiateRepairRecipe.to_proto(self.initiate_repair_recipe)
            )
        else:
            request.resource.ClearField("initiate_repair_recipe")
        if InstanceRepairRecipe.to_proto(self.repair_recipe):
            request.resource.repair_recipe.CopyFrom(
                InstanceRepairRecipe.to_proto(self.repair_recipe)
            )
        else:
            request.resource.ClearField("repair_recipe")
        if InstancePreprocessDeleteRecipe.to_proto(self.preprocess_delete_recipe):
            request.resource.preprocess_delete_recipe.CopyFrom(
                InstancePreprocessDeleteRecipe.to_proto(self.preprocess_delete_recipe)
            )
        else:
            request.resource.ClearField("preprocess_delete_recipe")
        if InstanceInitiateDeleteRecipe.to_proto(self.initiate_delete_recipe):
            request.resource.initiate_delete_recipe.CopyFrom(
                InstanceInitiateDeleteRecipe.to_proto(self.initiate_delete_recipe)
            )
        else:
            request.resource.ClearField("initiate_delete_recipe")
        if InstancePreprocessUpdateRecipe.to_proto(self.preprocess_update_recipe):
            request.resource.preprocess_update_recipe.CopyFrom(
                InstancePreprocessUpdateRecipe.to_proto(self.preprocess_update_recipe)
            )
        else:
            request.resource.ClearField("preprocess_update_recipe")
        if InstanceInitiateUpdateRecipe.to_proto(self.initiate_update_recipe):
            request.resource.initiate_update_recipe.CopyFrom(
                InstanceInitiateUpdateRecipe.to_proto(self.initiate_update_recipe)
            )
        else:
            request.resource.ClearField("initiate_update_recipe")
        if InstancePreprocessFreezeRecipe.to_proto(self.preprocess_freeze_recipe):
            request.resource.preprocess_freeze_recipe.CopyFrom(
                InstancePreprocessFreezeRecipe.to_proto(self.preprocess_freeze_recipe)
            )
        else:
            request.resource.ClearField("preprocess_freeze_recipe")
        if InstanceFreezeRecipe.to_proto(self.freeze_recipe):
            request.resource.freeze_recipe.CopyFrom(
                InstanceFreezeRecipe.to_proto(self.freeze_recipe)
            )
        else:
            request.resource.ClearField("freeze_recipe")
        if InstancePreprocessUnfreezeRecipe.to_proto(self.preprocess_unfreeze_recipe):
            request.resource.preprocess_unfreeze_recipe.CopyFrom(
                InstancePreprocessUnfreezeRecipe.to_proto(
                    self.preprocess_unfreeze_recipe
                )
            )
        else:
            request.resource.ClearField("preprocess_unfreeze_recipe")
        if InstanceUnfreezeRecipe.to_proto(self.unfreeze_recipe):
            request.resource.unfreeze_recipe.CopyFrom(
                InstanceUnfreezeRecipe.to_proto(self.unfreeze_recipe)
            )
        else:
            request.resource.ClearField("unfreeze_recipe")
        if InstanceReadonlyRecipe.to_proto(self.readonly_recipe):
            request.resource.readonly_recipe.CopyFrom(
                InstanceReadonlyRecipe.to_proto(self.readonly_recipe)
            )
        else:
            request.resource.ClearField("readonly_recipe")
        if Primitive.to_proto(self.enable_call_history):
            request.resource.enable_call_history = Primitive.to_proto(
                self.enable_call_history
            )

        if InstanceHistoryArray.to_proto(self.history):
            request.resource.history.extend(InstanceHistoryArray.to_proto(self.history))
        if Primitive.to_proto(self.public_resource_view_override):
            request.resource.public_resource_view_override = Primitive.to_proto(
                self.public_resource_view_override
            )

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.location):
            request.resource.location = Primitive.to_proto(self.location)

        response = stub.DeleteTier2AlphaInstance(request)

    @classmethod
    def list(self, project, location, service_account_file=""):
        stub = instance_pb2_grpc.Tier2AlphaInstanceServiceStub(channel.Channel())
        request = instance_pb2.ListTier2AlphaInstanceRequest()
        request.service_account_file = service_account_file
        request.Project = project

        request.Location = location

        return stub.ListTier2AlphaInstance(request).items

    @classmethod
    def from_any(self, any_proto):
        # Marshal any proto to regular proto.
        res_proto = instance_pb2.Tier2AlphaInstance()
        any_proto.Unpack(res_proto)

        res = Instance()
        res.name = Primitive.from_proto(res_proto.name)
        res.display_name = Primitive.from_proto(res_proto.display_name)
        res.labels = Primitive.from_proto(res_proto.labels)
        res.zone = Primitive.from_proto(res_proto.zone)
        res.sku = InstanceSku.from_proto(res_proto.sku)
        res.authorized_network_id = Primitive.from_proto(
            res_proto.authorized_network_id
        )
        res.reserved_ip_range = Primitive.from_proto(res_proto.reserved_ip_range)
        res.host_name = Primitive.from_proto(res_proto.host_name)
        res.port_number = Primitive.from_proto(res_proto.port_number)
        res.current_zone = Primitive.from_proto(res_proto.current_zone)
        res.creation_time = Primitive.from_proto(res_proto.creation_time)
        res.state = InstanceStateEnum.from_proto(res_proto.state)
        res.status_message = Primitive.from_proto(res_proto.status_message)
        res.extra_field = Primitive.from_proto(res_proto.extra_field)
        res.preprocess_create_recipe = InstancePreprocessCreateRecipe.from_proto(
            res_proto.preprocess_create_recipe
        )
        res.initiate_create_recipe = InstanceInitiateCreateRecipe.from_proto(
            res_proto.initiate_create_recipe
        )
        res.create_recipe = InstanceCreateRecipe.from_proto(res_proto.create_recipe)
        res.delete_recipe = InstanceDeleteRecipe.from_proto(res_proto.delete_recipe)
        res.update_recipe = InstanceUpdateRecipe.from_proto(res_proto.update_recipe)
        res.preprocess_reset_recipe = InstancePreprocessResetRecipe.from_proto(
            res_proto.preprocess_reset_recipe
        )
        res.initiate_reset_recipe = InstanceInitiateResetRecipe.from_proto(
            res_proto.initiate_reset_recipe
        )
        res.reset_recipe = InstanceResetRecipe.from_proto(res_proto.reset_recipe)
        res.preprocess_repair_recipe = InstancePreprocessRepairRecipe.from_proto(
            res_proto.preprocess_repair_recipe
        )
        res.initiate_repair_recipe = InstanceInitiateRepairRecipe.from_proto(
            res_proto.initiate_repair_recipe
        )
        res.repair_recipe = InstanceRepairRecipe.from_proto(res_proto.repair_recipe)
        res.preprocess_delete_recipe = InstancePreprocessDeleteRecipe.from_proto(
            res_proto.preprocess_delete_recipe
        )
        res.initiate_delete_recipe = InstanceInitiateDeleteRecipe.from_proto(
            res_proto.initiate_delete_recipe
        )
        res.preprocess_update_recipe = InstancePreprocessUpdateRecipe.from_proto(
            res_proto.preprocess_update_recipe
        )
        res.initiate_update_recipe = InstanceInitiateUpdateRecipe.from_proto(
            res_proto.initiate_update_recipe
        )
        res.preprocess_freeze_recipe = InstancePreprocessFreezeRecipe.from_proto(
            res_proto.preprocess_freeze_recipe
        )
        res.freeze_recipe = InstanceFreezeRecipe.from_proto(res_proto.freeze_recipe)
        res.preprocess_unfreeze_recipe = InstancePreprocessUnfreezeRecipe.from_proto(
            res_proto.preprocess_unfreeze_recipe
        )
        res.unfreeze_recipe = InstanceUnfreezeRecipe.from_proto(
            res_proto.unfreeze_recipe
        )
        res.readonly_recipe = InstanceReadonlyRecipe.from_proto(
            res_proto.readonly_recipe
        )
        res.enable_call_history = Primitive.from_proto(res_proto.enable_call_history)
        res.history = InstanceHistoryArray.from_proto(res_proto.history)
        res.public_resource_view_override = Primitive.from_proto(
            res_proto.public_resource_view_override
        )
        res.project = Primitive.from_proto(res_proto.project)
        res.location = Primitive.from_proto(res_proto.location)
        return res

    def to_proto(self):
        resource = instance_pb2.Tier2AlphaInstance()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if Primitive.to_proto(self.labels):
            resource.labels = Primitive.to_proto(self.labels)
        if Primitive.to_proto(self.zone):
            resource.zone = Primitive.to_proto(self.zone)
        if InstanceSku.to_proto(self.sku):
            resource.sku.CopyFrom(InstanceSku.to_proto(self.sku))
        else:
            resource.ClearField("sku")
        if Primitive.to_proto(self.authorized_network_id):
            resource.authorized_network_id = Primitive.to_proto(
                self.authorized_network_id
            )
        if Primitive.to_proto(self.reserved_ip_range):
            resource.reserved_ip_range = Primitive.to_proto(self.reserved_ip_range)
        if Primitive.to_proto(self.host_name):
            resource.host_name = Primitive.to_proto(self.host_name)
        if Primitive.to_proto(self.port_number):
            resource.port_number = Primitive.to_proto(self.port_number)
        if Primitive.to_proto(self.current_zone):
            resource.current_zone = Primitive.to_proto(self.current_zone)
        if Primitive.to_proto(self.creation_time):
            resource.creation_time = Primitive.to_proto(self.creation_time)
        if InstanceStateEnum.to_proto(self.state):
            resource.state = InstanceStateEnum.to_proto(self.state)
        if Primitive.to_proto(self.status_message):
            resource.status_message = Primitive.to_proto(self.status_message)
        if Primitive.to_proto(self.extra_field):
            resource.extra_field = Primitive.to_proto(self.extra_field)
        if InstancePreprocessCreateRecipe.to_proto(self.preprocess_create_recipe):
            resource.preprocess_create_recipe.CopyFrom(
                InstancePreprocessCreateRecipe.to_proto(self.preprocess_create_recipe)
            )
        else:
            resource.ClearField("preprocess_create_recipe")
        if InstanceInitiateCreateRecipe.to_proto(self.initiate_create_recipe):
            resource.initiate_create_recipe.CopyFrom(
                InstanceInitiateCreateRecipe.to_proto(self.initiate_create_recipe)
            )
        else:
            resource.ClearField("initiate_create_recipe")
        if InstanceCreateRecipe.to_proto(self.create_recipe):
            resource.create_recipe.CopyFrom(
                InstanceCreateRecipe.to_proto(self.create_recipe)
            )
        else:
            resource.ClearField("create_recipe")
        if InstanceDeleteRecipe.to_proto(self.delete_recipe):
            resource.delete_recipe.CopyFrom(
                InstanceDeleteRecipe.to_proto(self.delete_recipe)
            )
        else:
            resource.ClearField("delete_recipe")
        if InstanceUpdateRecipe.to_proto(self.update_recipe):
            resource.update_recipe.CopyFrom(
                InstanceUpdateRecipe.to_proto(self.update_recipe)
            )
        else:
            resource.ClearField("update_recipe")
        if InstancePreprocessResetRecipe.to_proto(self.preprocess_reset_recipe):
            resource.preprocess_reset_recipe.CopyFrom(
                InstancePreprocessResetRecipe.to_proto(self.preprocess_reset_recipe)
            )
        else:
            resource.ClearField("preprocess_reset_recipe")
        if InstanceInitiateResetRecipe.to_proto(self.initiate_reset_recipe):
            resource.initiate_reset_recipe.CopyFrom(
                InstanceInitiateResetRecipe.to_proto(self.initiate_reset_recipe)
            )
        else:
            resource.ClearField("initiate_reset_recipe")
        if InstanceResetRecipe.to_proto(self.reset_recipe):
            resource.reset_recipe.CopyFrom(
                InstanceResetRecipe.to_proto(self.reset_recipe)
            )
        else:
            resource.ClearField("reset_recipe")
        if InstancePreprocessRepairRecipe.to_proto(self.preprocess_repair_recipe):
            resource.preprocess_repair_recipe.CopyFrom(
                InstancePreprocessRepairRecipe.to_proto(self.preprocess_repair_recipe)
            )
        else:
            resource.ClearField("preprocess_repair_recipe")
        if InstanceInitiateRepairRecipe.to_proto(self.initiate_repair_recipe):
            resource.initiate_repair_recipe.CopyFrom(
                InstanceInitiateRepairRecipe.to_proto(self.initiate_repair_recipe)
            )
        else:
            resource.ClearField("initiate_repair_recipe")
        if InstanceRepairRecipe.to_proto(self.repair_recipe):
            resource.repair_recipe.CopyFrom(
                InstanceRepairRecipe.to_proto(self.repair_recipe)
            )
        else:
            resource.ClearField("repair_recipe")
        if InstancePreprocessDeleteRecipe.to_proto(self.preprocess_delete_recipe):
            resource.preprocess_delete_recipe.CopyFrom(
                InstancePreprocessDeleteRecipe.to_proto(self.preprocess_delete_recipe)
            )
        else:
            resource.ClearField("preprocess_delete_recipe")
        if InstanceInitiateDeleteRecipe.to_proto(self.initiate_delete_recipe):
            resource.initiate_delete_recipe.CopyFrom(
                InstanceInitiateDeleteRecipe.to_proto(self.initiate_delete_recipe)
            )
        else:
            resource.ClearField("initiate_delete_recipe")
        if InstancePreprocessUpdateRecipe.to_proto(self.preprocess_update_recipe):
            resource.preprocess_update_recipe.CopyFrom(
                InstancePreprocessUpdateRecipe.to_proto(self.preprocess_update_recipe)
            )
        else:
            resource.ClearField("preprocess_update_recipe")
        if InstanceInitiateUpdateRecipe.to_proto(self.initiate_update_recipe):
            resource.initiate_update_recipe.CopyFrom(
                InstanceInitiateUpdateRecipe.to_proto(self.initiate_update_recipe)
            )
        else:
            resource.ClearField("initiate_update_recipe")
        if InstancePreprocessFreezeRecipe.to_proto(self.preprocess_freeze_recipe):
            resource.preprocess_freeze_recipe.CopyFrom(
                InstancePreprocessFreezeRecipe.to_proto(self.preprocess_freeze_recipe)
            )
        else:
            resource.ClearField("preprocess_freeze_recipe")
        if InstanceFreezeRecipe.to_proto(self.freeze_recipe):
            resource.freeze_recipe.CopyFrom(
                InstanceFreezeRecipe.to_proto(self.freeze_recipe)
            )
        else:
            resource.ClearField("freeze_recipe")
        if InstancePreprocessUnfreezeRecipe.to_proto(self.preprocess_unfreeze_recipe):
            resource.preprocess_unfreeze_recipe.CopyFrom(
                InstancePreprocessUnfreezeRecipe.to_proto(
                    self.preprocess_unfreeze_recipe
                )
            )
        else:
            resource.ClearField("preprocess_unfreeze_recipe")
        if InstanceUnfreezeRecipe.to_proto(self.unfreeze_recipe):
            resource.unfreeze_recipe.CopyFrom(
                InstanceUnfreezeRecipe.to_proto(self.unfreeze_recipe)
            )
        else:
            resource.ClearField("unfreeze_recipe")
        if InstanceReadonlyRecipe.to_proto(self.readonly_recipe):
            resource.readonly_recipe.CopyFrom(
                InstanceReadonlyRecipe.to_proto(self.readonly_recipe)
            )
        else:
            resource.ClearField("readonly_recipe")
        if Primitive.to_proto(self.enable_call_history):
            resource.enable_call_history = Primitive.to_proto(self.enable_call_history)
        if InstanceHistoryArray.to_proto(self.history):
            resource.history.extend(InstanceHistoryArray.to_proto(self.history))
        if Primitive.to_proto(self.public_resource_view_override):
            resource.public_resource_view_override = Primitive.to_proto(
                self.public_resource_view_override
            )
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.location):
            resource.location = Primitive.to_proto(self.location)
        return resource


class InstanceSku(object):
    def __init__(self, tier: str = None, size: str = None):
        self.tier = tier
        self.size = size

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceSku()
        if InstanceSkuTierEnum.to_proto(resource.tier):
            res.tier = InstanceSkuTierEnum.to_proto(resource.tier)
        if InstanceSkuSizeEnum.to_proto(resource.size):
            res.size = InstanceSkuSizeEnum.to_proto(resource.size)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceSku(tier=resource.tier, size=resource.size,)


class InstanceSkuArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceSku.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceSku.from_proto(i) for i in resources]


class InstancePreprocessCreateRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessCreateRecipe()
        if InstancePreprocessCreateRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessCreateRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstancePreprocessCreateRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessCreateRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessCreateRecipe.from_proto(i) for i in resources]


class InstancePreprocessCreateRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessCreateRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessCreateRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessCreateRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessCreateRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessCreateRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessCreateRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessCreateRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessCreateRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessCreateRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessCreateRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessCreateRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessCreateRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstancePreprocessCreateRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessCreateRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessCreateRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessCreateRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessCreateRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessCreateRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstancePreprocessCreateRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstancePreprocessCreateRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstancePreprocessCreateRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstancePreprocessCreateRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstancePreprocessCreateRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstancePreprocessCreateRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstancePreprocessCreateRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs()


class InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateCreateRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateCreateRecipe()
        if InstanceInitiateCreateRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstanceInitiateCreateRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstanceInitiateCreateRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateCreateRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceInitiateCreateRecipe.from_proto(i) for i in resources]


class InstanceInitiateCreateRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceInitiateCreateRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceInitiateCreateRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstanceInitiateCreateRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceInitiateCreateRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceInitiateCreateRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceInitiateCreateRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceInitiateCreateRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceInitiateCreateRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceInitiateCreateRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceInitiateCreateRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceInitiateCreateRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceInitiateCreateRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstanceInitiateCreateRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateCreateRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceInitiateCreateRecipeSteps.from_proto(i) for i in resources]


class InstanceInitiateCreateRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceInitiateCreateRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstanceInitiateCreateRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstanceInitiateCreateRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateCreateRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateCreateRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstanceInitiateCreateRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstanceInitiateCreateRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateCreateRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateCreateRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstanceInitiateCreateRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstanceInitiateCreateRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateCreateRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateCreateRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstanceInitiateCreateRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstanceInitiateCreateRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateCreateRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateCreateRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstanceInitiateCreateRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstanceInitiateCreateRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateCreateRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateCreateRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceInitiateCreateRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfo()
        if InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstanceInitiateCreateRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateCreateRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateCreateRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs()


class InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdate()
        )
        if InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstanceCreateRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipe()
        if InstanceCreateRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceCreateRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstanceCreateRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceCreateRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceCreateRecipe.from_proto(i) for i in resources]


class InstanceCreateRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceCreateRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceCreateRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceCreateRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceCreateRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceCreateRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceCreateRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceCreateRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceCreateRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceCreateRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceCreateRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceCreateRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceCreateRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceCreateRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceCreateRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstanceCreateRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceCreateRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceCreateRecipeSteps.from_proto(i) for i in resources]


class InstanceCreateRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceCreateRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceCreateRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstanceCreateRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceCreateRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceCreateRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceCreateRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstanceCreateRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceCreateRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceCreateRecipeStepsStatusDetails.from_proto(i) for i in resources]


class InstanceCreateRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstanceCreateRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsQuotaRequestDeltas.from_proto(i) for i in resources
        ]


class InstanceCreateRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstanceCreateRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceCreateRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstanceCreateRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceCreateRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfo()
        if InstanceCreateRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceCreateRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceCreateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceCreateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceCreateRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstanceCreateRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstanceCreateRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceCreateRecipeStepsPermissionsInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceCreateRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstanceCreateRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceCreateRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstanceCreateRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceCreateRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoApiAttrs()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsPermissionsInfoApiAttrs()


class InstanceCreateRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstanceCreateRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdate()
        if InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstanceCreateRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstanceDeleteRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipe()
        if InstanceDeleteRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceDeleteRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstanceDeleteRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDeleteRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceDeleteRecipe.from_proto(i) for i in resources]


class InstanceDeleteRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceDeleteRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceDeleteRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceDeleteRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceDeleteRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceDeleteRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceDeleteRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceDeleteRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceDeleteRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceDeleteRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceDeleteRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceDeleteRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceDeleteRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceDeleteRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceDeleteRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstanceDeleteRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDeleteRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceDeleteRecipeSteps.from_proto(i) for i in resources]


class InstanceDeleteRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceDeleteRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceDeleteRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstanceDeleteRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDeleteRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceDeleteRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceDeleteRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstanceDeleteRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDeleteRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceDeleteRecipeStepsStatusDetails.from_proto(i) for i in resources]


class InstanceDeleteRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstanceDeleteRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsQuotaRequestDeltas.from_proto(i) for i in resources
        ]


class InstanceDeleteRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstanceDeleteRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceDeleteRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstanceDeleteRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceDeleteRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfo()
        if InstanceDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceDeleteRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceDeleteRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceDeleteRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstanceDeleteRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstanceDeleteRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceDeleteRecipeStepsPermissionsInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceDeleteRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstanceDeleteRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceDeleteRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstanceDeleteRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceDeleteRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoApiAttrs()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsPermissionsInfoApiAttrs()


class InstanceDeleteRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstanceDeleteRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdate()
        if InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstanceDeleteRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstanceUpdateRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipe()
        if InstanceUpdateRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceUpdateRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstanceUpdateRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUpdateRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceUpdateRecipe.from_proto(i) for i in resources]


class InstanceUpdateRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceUpdateRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceUpdateRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceUpdateRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceUpdateRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceUpdateRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceUpdateRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceUpdateRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceUpdateRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceUpdateRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceUpdateRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceUpdateRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceUpdateRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceUpdateRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceUpdateRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstanceUpdateRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUpdateRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceUpdateRecipeSteps.from_proto(i) for i in resources]


class InstanceUpdateRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceUpdateRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceUpdateRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstanceUpdateRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUpdateRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceUpdateRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceUpdateRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstanceUpdateRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUpdateRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceUpdateRecipeStepsStatusDetails.from_proto(i) for i in resources]


class InstanceUpdateRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstanceUpdateRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsQuotaRequestDeltas.from_proto(i) for i in resources
        ]


class InstanceUpdateRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstanceUpdateRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceUpdateRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstanceUpdateRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceUpdateRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfo()
        if InstanceUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceUpdateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceUpdateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceUpdateRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstanceUpdateRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstanceUpdateRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUpdateRecipeStepsPermissionsInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceUpdateRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstanceUpdateRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceUpdateRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstanceUpdateRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceUpdateRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoApiAttrs()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsPermissionsInfoApiAttrs()


class InstanceUpdateRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstanceUpdateRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdate()
        if InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstanceUpdateRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessResetRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessResetRecipe()
        if InstancePreprocessResetRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessResetRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstancePreprocessResetRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessResetRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessResetRecipe.from_proto(i) for i in resources]


class InstancePreprocessResetRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessResetRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessResetRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessResetRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessResetRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessResetRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessResetRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessResetRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessResetRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessResetRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessResetRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessResetRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessResetRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessResetRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessResetRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessResetRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstancePreprocessResetRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessResetRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessResetRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessResetRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessResetRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessResetRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstancePreprocessResetRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessResetRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstancePreprocessResetRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstancePreprocessResetRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstancePreprocessResetRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstancePreprocessResetRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstancePreprocessResetRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfo()
        if InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstancePreprocessResetRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs()


class InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstancePreprocessResetRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateResetRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateResetRecipe()
        if InstanceInitiateResetRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstanceInitiateResetRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstanceInitiateResetRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateResetRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceInitiateResetRecipe.from_proto(i) for i in resources]


class InstanceInitiateResetRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateResetRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceInitiateResetRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceInitiateResetRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstanceInitiateResetRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceInitiateResetRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceInitiateResetRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceInitiateResetRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceInitiateResetRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceInitiateResetRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceInitiateResetRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceInitiateResetRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceInitiateResetRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceInitiateResetRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceInitiateResetRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceInitiateResetRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstanceInitiateResetRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateResetRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceInitiateResetRecipeSteps.from_proto(i) for i in resources]


class InstanceInitiateResetRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceInitiateResetRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstanceInitiateResetRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstanceInitiateResetRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateResetRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceInitiateResetRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceInitiateResetRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstanceInitiateResetRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateResetRecipeStepsStatusDetails.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateResetRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstanceInitiateResetRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstanceInitiateResetRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateResetRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateResetRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstanceInitiateResetRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstanceInitiateResetRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateResetRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateResetRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstanceInitiateResetRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstanceInitiateResetRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateResetRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateResetRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceInitiateResetRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfo()
        if InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstanceInitiateResetRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateResetRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateResetRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstanceInitiateResetRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs()


class InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstanceInitiateResetRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdate()
        )
        if InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstanceInitiateResetRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateResetRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateResetRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstanceResetRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipe()
        if InstanceResetRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceResetRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstanceResetRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceResetRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceResetRecipe.from_proto(i) for i in resources]


class InstanceResetRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceResetRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceResetRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceResetRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceResetRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceResetRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceResetRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceResetRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceResetRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceResetRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceResetRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceResetRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceResetRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceResetRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceResetRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstanceResetRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceResetRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceResetRecipeSteps.from_proto(i) for i in resources]


class InstanceResetRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceResetRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceResetRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstanceResetRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceResetRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceResetRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceResetRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstanceResetRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceResetRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceResetRecipeStepsStatusDetails.from_proto(i) for i in resources]


class InstanceResetRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstanceResetRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsQuotaRequestDeltas.from_proto(i) for i in resources
        ]


class InstanceResetRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstanceResetRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceResetRecipeStepsPreprocessUpdate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceResetRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceResetRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceResetRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstanceResetRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceResetRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsPermissionsInfo()
        if InstanceResetRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceResetRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceResetRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceResetRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceResetRecipeStepsPermissionsInfoApiAttrs.to_proto(resource.api_attrs):
            res.api_attrs.CopyFrom(
                InstanceResetRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstanceResetRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceResetRecipeStepsPermissionsInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceResetRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyName()
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstanceResetRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceResetRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstanceResetRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceResetRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoApiAttrs()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsPermissionsInfoApiAttrs()


class InstanceResetRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstanceResetRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdate()
        if InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstanceResetRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessRepairRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessRepairRecipe()
        if InstancePreprocessRepairRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessRepairRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstancePreprocessRepairRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessRepairRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessRepairRecipe.from_proto(i) for i in resources]


class InstancePreprocessRepairRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessRepairRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessRepairRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessRepairRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessRepairRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessRepairRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessRepairRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessRepairRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessRepairRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessRepairRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessRepairRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessRepairRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessRepairRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstancePreprocessRepairRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessRepairRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessRepairRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessRepairRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessRepairRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessRepairRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstancePreprocessRepairRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstancePreprocessRepairRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstancePreprocessRepairRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstancePreprocessRepairRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstancePreprocessRepairRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstancePreprocessRepairRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstancePreprocessRepairRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs()


class InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateRepairRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateRepairRecipe()
        if InstanceInitiateRepairRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstanceInitiateRepairRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstanceInitiateRepairRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateRepairRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceInitiateRepairRecipe.from_proto(i) for i in resources]


class InstanceInitiateRepairRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceInitiateRepairRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceInitiateRepairRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstanceInitiateRepairRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceInitiateRepairRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceInitiateRepairRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceInitiateRepairRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceInitiateRepairRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceInitiateRepairRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceInitiateRepairRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceInitiateRepairRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceInitiateRepairRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceInitiateRepairRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstanceInitiateRepairRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateRepairRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceInitiateRepairRecipeSteps.from_proto(i) for i in resources]


class InstanceInitiateRepairRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceInitiateRepairRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstanceInitiateRepairRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstanceInitiateRepairRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateRepairRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateRepairRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstanceInitiateRepairRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstanceInitiateRepairRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateRepairRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateRepairRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstanceInitiateRepairRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstanceInitiateRepairRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateRepairRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateRepairRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstanceInitiateRepairRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstanceInitiateRepairRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateRepairRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateRepairRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstanceInitiateRepairRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstanceInitiateRepairRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateRepairRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateRepairRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceInitiateRepairRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfo()
        if InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstanceInitiateRepairRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateRepairRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateRepairRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs()


class InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdate()
        )
        if InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstanceRepairRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipe()
        if InstanceRepairRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceRepairRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstanceRepairRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceRepairRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceRepairRecipe.from_proto(i) for i in resources]


class InstanceRepairRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceRepairRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceRepairRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceRepairRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceRepairRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceRepairRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceRepairRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceRepairRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceRepairRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceRepairRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceRepairRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceRepairRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceRepairRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceRepairRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceRepairRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstanceRepairRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceRepairRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceRepairRecipeSteps.from_proto(i) for i in resources]


class InstanceRepairRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceRepairRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceRepairRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstanceRepairRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceRepairRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceRepairRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceRepairRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstanceRepairRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceRepairRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceRepairRecipeStepsStatusDetails.from_proto(i) for i in resources]


class InstanceRepairRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstanceRepairRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsQuotaRequestDeltas.from_proto(i) for i in resources
        ]


class InstanceRepairRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstanceRepairRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceRepairRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstanceRepairRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceRepairRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfo()
        if InstanceRepairRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceRepairRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceRepairRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceRepairRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceRepairRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstanceRepairRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstanceRepairRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceRepairRecipeStepsPermissionsInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceRepairRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstanceRepairRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceRepairRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstanceRepairRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceRepairRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoApiAttrs()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsPermissionsInfoApiAttrs()


class InstanceRepairRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstanceRepairRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdate()
        if InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstanceRepairRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessDeleteRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipe()
        if InstancePreprocessDeleteRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessDeleteRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstancePreprocessDeleteRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessDeleteRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessDeleteRecipe.from_proto(i) for i in resources]


class InstancePreprocessDeleteRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessDeleteRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessDeleteRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessDeleteRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessDeleteRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessDeleteRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessDeleteRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessDeleteRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessDeleteRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessDeleteRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessDeleteRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessDeleteRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessDeleteRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstancePreprocessDeleteRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessDeleteRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessDeleteRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessDeleteRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessDeleteRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessDeleteRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstancePreprocessDeleteRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstancePreprocessDeleteRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstancePreprocessDeleteRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstancePreprocessDeleteRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstancePreprocessDeleteRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstancePreprocessDeleteRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs()


class InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateDeleteRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipe()
        if InstanceInitiateDeleteRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstanceInitiateDeleteRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstanceInitiateDeleteRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateDeleteRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceInitiateDeleteRecipe.from_proto(i) for i in resources]


class InstanceInitiateDeleteRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceInitiateDeleteRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceInitiateDeleteRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstanceInitiateDeleteRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceInitiateDeleteRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceInitiateDeleteRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceInitiateDeleteRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceInitiateDeleteRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceInitiateDeleteRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceInitiateDeleteRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceInitiateDeleteRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceInitiateDeleteRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceInitiateDeleteRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstanceInitiateDeleteRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateDeleteRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceInitiateDeleteRecipeSteps.from_proto(i) for i in resources]


class InstanceInitiateDeleteRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceInitiateDeleteRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstanceInitiateDeleteRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstanceInitiateDeleteRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateDeleteRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateDeleteRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstanceInitiateDeleteRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstanceInitiateDeleteRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateDeleteRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateDeleteRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstanceInitiateDeleteRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstanceInitiateDeleteRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstanceInitiateDeleteRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateDeleteRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateDeleteRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstanceInitiateDeleteRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstanceInitiateDeleteRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateDeleteRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateDeleteRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceInitiateDeleteRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfo()
        if InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstanceInitiateDeleteRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateDeleteRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateDeleteRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs()


class InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate()
        )
        if InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessUpdateRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipe()
        if InstancePreprocessUpdateRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessUpdateRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstancePreprocessUpdateRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessUpdateRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessUpdateRecipe.from_proto(i) for i in resources]


class InstancePreprocessUpdateRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessUpdateRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessUpdateRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessUpdateRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessUpdateRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessUpdateRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessUpdateRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessUpdateRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessUpdateRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessUpdateRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessUpdateRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessUpdateRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessUpdateRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstancePreprocessUpdateRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessUpdateRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessUpdateRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessUpdateRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessUpdateRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessUpdateRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstancePreprocessUpdateRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstancePreprocessUpdateRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstancePreprocessUpdateRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstancePreprocessUpdateRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstancePreprocessUpdateRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstancePreprocessUpdateRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs()


class InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateUpdateRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipe()
        if InstanceInitiateUpdateRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstanceInitiateUpdateRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstanceInitiateUpdateRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateUpdateRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceInitiateUpdateRecipe.from_proto(i) for i in resources]


class InstanceInitiateUpdateRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceInitiateUpdateRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceInitiateUpdateRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstanceInitiateUpdateRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceInitiateUpdateRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceInitiateUpdateRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceInitiateUpdateRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceInitiateUpdateRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceInitiateUpdateRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceInitiateUpdateRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceInitiateUpdateRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceInitiateUpdateRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceInitiateUpdateRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstanceInitiateUpdateRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateUpdateRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceInitiateUpdateRecipeSteps.from_proto(i) for i in resources]


class InstanceInitiateUpdateRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceInitiateUpdateRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstanceInitiateUpdateRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstanceInitiateUpdateRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceInitiateUpdateRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateUpdateRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstanceInitiateUpdateRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstanceInitiateUpdateRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateUpdateRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateUpdateRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstanceInitiateUpdateRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstanceInitiateUpdateRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstanceInitiateUpdateRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateUpdateRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateUpdateRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstanceInitiateUpdateRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstanceInitiateUpdateRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateUpdateRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateUpdateRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceInitiateUpdateRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfo()
        if InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstanceInitiateUpdateRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateUpdateRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateUpdateRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs()


class InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate()
        )
        if InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessFreezeRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipe()
        if InstancePreprocessFreezeRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessFreezeRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstancePreprocessFreezeRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessFreezeRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessFreezeRecipe.from_proto(i) for i in resources]


class InstancePreprocessFreezeRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessFreezeRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessFreezeRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessFreezeRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessFreezeRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessFreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessFreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessFreezeRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessFreezeRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessFreezeRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessFreezeRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessFreezeRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessFreezeRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstancePreprocessFreezeRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessFreezeRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessFreezeRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessFreezeRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessFreezeRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessFreezeRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstancePreprocessFreezeRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstancePreprocessFreezeRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstancePreprocessFreezeRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstancePreprocessFreezeRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstancePreprocessFreezeRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstancePreprocessFreezeRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs()


class InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstanceFreezeRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipe()
        if InstanceFreezeRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceFreezeRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstanceFreezeRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceFreezeRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceFreezeRecipe.from_proto(i) for i in resources]


class InstanceFreezeRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceFreezeRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceFreezeRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceFreezeRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceFreezeRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceFreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceFreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceFreezeRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceFreezeRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceFreezeRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceFreezeRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceFreezeRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceFreezeRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceFreezeRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceFreezeRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstanceFreezeRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceFreezeRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceFreezeRecipeSteps.from_proto(i) for i in resources]


class InstanceFreezeRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceFreezeRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceFreezeRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstanceFreezeRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceFreezeRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceFreezeRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceFreezeRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstanceFreezeRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceFreezeRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceFreezeRecipeStepsStatusDetails.from_proto(i) for i in resources]


class InstanceFreezeRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstanceFreezeRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsQuotaRequestDeltas.from_proto(i) for i in resources
        ]


class InstanceFreezeRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstanceFreezeRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceFreezeRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstanceFreezeRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceFreezeRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfo()
        if InstanceFreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceFreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceFreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceFreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceFreezeRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstanceFreezeRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstanceFreezeRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceFreezeRecipeStepsPermissionsInfo.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceFreezeRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstanceFreezeRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceFreezeRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstanceFreezeRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceFreezeRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoApiAttrs()
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsPermissionsInfoApiAttrs()


class InstanceFreezeRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstanceFreezeRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdate()
        if InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstanceFreezeRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipe()
        if InstancePreprocessUnfreezeRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(
                InstancePreprocessUnfreezeRecipeStepsArray.to_proto(resource.steps)
            )
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstancePreprocessUnfreezeRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessUnfreezeRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessUnfreezeRecipe.from_proto(i) for i in resources]


class InstancePreprocessUnfreezeRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstancePreprocessUnfreezeRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstancePreprocessUnfreezeRecipeStepsActionEnum.to_proto(
                resource.action
            )
        if InstancePreprocessUnfreezeRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstancePreprocessUnfreezeRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstancePreprocessUnfreezeRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstancePreprocessUnfreezeRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstancePreprocessUnfreezeRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstancePreprocessUnfreezeRecipeSteps.from_proto(i) for i in resources]


class InstancePreprocessUnfreezeRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstancePreprocessUnfreezeRecipeStepsStatusDetailsArray.to_proto(
            resource.details
        ):
            res.details.extend(
                InstancePreprocessUnfreezeRecipeStepsStatusDetailsArray.to_proto(
                    resource.details
                )
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstancePreprocessUnfreezeRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsStatus.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsStatus.from_proto(i) for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetails()
        )
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstancePreprocessUnfreezeRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsStatusDetails.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsStatusDetails.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas()
        )
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate()
        )
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstancePreprocessUnfreezeRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject()
        )
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfo()
        )
        if InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfo.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfo.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs()


class InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate()
        )
        if InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
    object
):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstanceUnfreezeRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipe()
        if InstanceUnfreezeRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceUnfreezeRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstanceUnfreezeRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUnfreezeRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceUnfreezeRecipe.from_proto(i) for i in resources]


class InstanceUnfreezeRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceUnfreezeRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceUnfreezeRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceUnfreezeRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceUnfreezeRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceUnfreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceUnfreezeRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceUnfreezeRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceUnfreezeRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceUnfreezeRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceUnfreezeRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceUnfreezeRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceUnfreezeRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceUnfreezeRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceUnfreezeRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstanceUnfreezeRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUnfreezeRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceUnfreezeRecipeSteps.from_proto(i) for i in resources]


class InstanceUnfreezeRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceUnfreezeRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceUnfreezeRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstanceUnfreezeRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUnfreezeRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceUnfreezeRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceUnfreezeRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstanceUnfreezeRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceUnfreezeRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsStatusDetails.from_proto(i) for i in resources
        ]


class InstanceUnfreezeRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstanceUnfreezeRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstanceUnfreezeRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceUnfreezeRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstanceUnfreezeRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfo()
        if InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstanceUnfreezeRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfo.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs()


class InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdate()
        if InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstanceUnfreezeRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstanceReadonlyRecipe(object):
    def __init__(
        self,
        steps: list = None,
        honor_cancel_request: bool = None,
        ignore_recipe_after: int = None,
        verify_deadline_seconds_below: float = None,
        populate_operation_result: bool = None,
        readonly_recipe_start_time: str = None,
        resource_names_stored_in_clh_with_delay: list = None,
        delay_to_store_resources_in_clh_db_nanos: int = None,
    ):
        self.steps = steps
        self.honor_cancel_request = honor_cancel_request
        self.ignore_recipe_after = ignore_recipe_after
        self.verify_deadline_seconds_below = verify_deadline_seconds_below
        self.populate_operation_result = populate_operation_result
        self.readonly_recipe_start_time = readonly_recipe_start_time
        self.resource_names_stored_in_clh_with_delay = (
            resource_names_stored_in_clh_with_delay
        )
        self.delay_to_store_resources_in_clh_db_nanos = (
            delay_to_store_resources_in_clh_db_nanos
        )

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipe()
        if InstanceReadonlyRecipeStepsArray.to_proto(resource.steps):
            res.steps.extend(InstanceReadonlyRecipeStepsArray.to_proto(resource.steps))
        if Primitive.to_proto(resource.honor_cancel_request):
            res.honor_cancel_request = Primitive.to_proto(resource.honor_cancel_request)
        if Primitive.to_proto(resource.ignore_recipe_after):
            res.ignore_recipe_after = Primitive.to_proto(resource.ignore_recipe_after)
        if Primitive.to_proto(resource.verify_deadline_seconds_below):
            res.verify_deadline_seconds_below = Primitive.to_proto(
                resource.verify_deadline_seconds_below
            )
        if Primitive.to_proto(resource.populate_operation_result):
            res.populate_operation_result = Primitive.to_proto(
                resource.populate_operation_result
            )
        if Primitive.to_proto(resource.readonly_recipe_start_time):
            res.readonly_recipe_start_time = Primitive.to_proto(
                resource.readonly_recipe_start_time
            )
        if Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay):
            res.resource_names_stored_in_clh_with_delay.extend(
                Primitive.to_proto(resource.resource_names_stored_in_clh_with_delay)
            )
        if Primitive.to_proto(resource.delay_to_store_resources_in_clh_db_nanos):
            res.delay_to_store_resources_in_clh_db_nanos = Primitive.to_proto(
                resource.delay_to_store_resources_in_clh_db_nanos
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipe(
            steps=resource.steps,
            honor_cancel_request=resource.honor_cancel_request,
            ignore_recipe_after=resource.ignore_recipe_after,
            verify_deadline_seconds_below=resource.verify_deadline_seconds_below,
            populate_operation_result=resource.populate_operation_result,
            readonly_recipe_start_time=resource.readonly_recipe_start_time,
            resource_names_stored_in_clh_with_delay=resource.resource_names_stored_in_clh_with_delay,
            delay_to_store_resources_in_clh_db_nanos=resource.delay_to_store_resources_in_clh_db_nanos,
        )


class InstanceReadonlyRecipeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReadonlyRecipe.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceReadonlyRecipe.from_proto(i) for i in resources]


class InstanceReadonlyRecipeSteps(object):
    def __init__(
        self,
        relative_time: int = None,
        sleep_duration: int = None,
        action: str = None,
        status: dict = None,
        error_space: str = None,
        p4_service_account: str = None,
        resource_metadata_size: int = None,
        description: str = None,
        updated_repeat_operation_delay_sec: float = None,
        quota_request_deltas: list = None,
        preprocess_update: dict = None,
        public_operation_metadata: str = None,
        requested_tenant_project: dict = None,
        permissions_info: list = None,
        key_notifications_update: dict = None,
        clh_data_update_time: str = None,
    ):
        self.relative_time = relative_time
        self.sleep_duration = sleep_duration
        self.action = action
        self.status = status
        self.error_space = error_space
        self.p4_service_account = p4_service_account
        self.resource_metadata_size = resource_metadata_size
        self.description = description
        self.updated_repeat_operation_delay_sec = updated_repeat_operation_delay_sec
        self.quota_request_deltas = quota_request_deltas
        self.preprocess_update = preprocess_update
        self.public_operation_metadata = public_operation_metadata
        self.requested_tenant_project = requested_tenant_project
        self.permissions_info = permissions_info
        self.key_notifications_update = key_notifications_update
        self.clh_data_update_time = clh_data_update_time

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeSteps()
        if Primitive.to_proto(resource.relative_time):
            res.relative_time = Primitive.to_proto(resource.relative_time)
        if Primitive.to_proto(resource.sleep_duration):
            res.sleep_duration = Primitive.to_proto(resource.sleep_duration)
        if InstanceReadonlyRecipeStepsActionEnum.to_proto(resource.action):
            res.action = InstanceReadonlyRecipeStepsActionEnum.to_proto(resource.action)
        if InstanceReadonlyRecipeStepsStatus.to_proto(resource.status):
            res.status.CopyFrom(
                InstanceReadonlyRecipeStepsStatus.to_proto(resource.status)
            )
        else:
            res.ClearField("status")
        if Primitive.to_proto(resource.error_space):
            res.error_space = Primitive.to_proto(resource.error_space)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        if Primitive.to_proto(resource.resource_metadata_size):
            res.resource_metadata_size = Primitive.to_proto(
                resource.resource_metadata_size
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.updated_repeat_operation_delay_sec):
            res.updated_repeat_operation_delay_sec = Primitive.to_proto(
                resource.updated_repeat_operation_delay_sec
            )
        if InstanceReadonlyRecipeStepsQuotaRequestDeltasArray.to_proto(
            resource.quota_request_deltas
        ):
            res.quota_request_deltas.extend(
                InstanceReadonlyRecipeStepsQuotaRequestDeltasArray.to_proto(
                    resource.quota_request_deltas
                )
            )
        if InstanceReadonlyRecipeStepsPreprocessUpdate.to_proto(
            resource.preprocess_update
        ):
            res.preprocess_update.CopyFrom(
                InstanceReadonlyRecipeStepsPreprocessUpdate.to_proto(
                    resource.preprocess_update
                )
            )
        else:
            res.ClearField("preprocess_update")
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        if InstanceReadonlyRecipeStepsRequestedTenantProject.to_proto(
            resource.requested_tenant_project
        ):
            res.requested_tenant_project.CopyFrom(
                InstanceReadonlyRecipeStepsRequestedTenantProject.to_proto(
                    resource.requested_tenant_project
                )
            )
        else:
            res.ClearField("requested_tenant_project")
        if InstanceReadonlyRecipeStepsPermissionsInfoArray.to_proto(
            resource.permissions_info
        ):
            res.permissions_info.extend(
                InstanceReadonlyRecipeStepsPermissionsInfoArray.to_proto(
                    resource.permissions_info
                )
            )
        if InstanceReadonlyRecipeStepsKeyNotificationsUpdate.to_proto(
            resource.key_notifications_update
        ):
            res.key_notifications_update.CopyFrom(
                InstanceReadonlyRecipeStepsKeyNotificationsUpdate.to_proto(
                    resource.key_notifications_update
                )
            )
        else:
            res.ClearField("key_notifications_update")
        if Primitive.to_proto(resource.clh_data_update_time):
            res.clh_data_update_time = Primitive.to_proto(resource.clh_data_update_time)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeSteps(
            relative_time=resource.relative_time,
            sleep_duration=resource.sleep_duration,
            action=resource.action,
            status=resource.status,
            error_space=resource.error_space,
            p4_service_account=resource.p4_service_account,
            resource_metadata_size=resource.resource_metadata_size,
            description=resource.description,
            updated_repeat_operation_delay_sec=resource.updated_repeat_operation_delay_sec,
            quota_request_deltas=resource.quota_request_deltas,
            preprocess_update=resource.preprocess_update,
            public_operation_metadata=resource.public_operation_metadata,
            requested_tenant_project=resource.requested_tenant_project,
            permissions_info=resource.permissions_info,
            key_notifications_update=resource.key_notifications_update,
            clh_data_update_time=resource.clh_data_update_time,
        )


class InstanceReadonlyRecipeStepsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReadonlyRecipeSteps.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceReadonlyRecipeSteps.from_proto(i) for i in resources]


class InstanceReadonlyRecipeStepsStatus(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsStatus()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if InstanceReadonlyRecipeStepsStatusDetailsArray.to_proto(resource.details):
            res.details.extend(
                InstanceReadonlyRecipeStepsStatusDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsStatus(
            code=resource.code, message=resource.message, details=resource.details,
        )


class InstanceReadonlyRecipeStepsStatusArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReadonlyRecipeStepsStatus.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceReadonlyRecipeStepsStatus.from_proto(i) for i in resources]


class InstanceReadonlyRecipeStepsStatusDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsStatusDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsStatusDetails(
            type_url=resource.type_url, value=resource.value,
        )


class InstanceReadonlyRecipeStepsStatusDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceReadonlyRecipeStepsStatusDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsStatusDetails.from_proto(i) for i in resources
        ]


class InstanceReadonlyRecipeStepsQuotaRequestDeltas(object):
    def __init__(self, metric_name: str = None, amount: int = None):
        self.metric_name = metric_name
        self.amount = amount

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltas()
        if Primitive.to_proto(resource.metric_name):
            res.metric_name = Primitive.to_proto(resource.metric_name)
        if Primitive.to_proto(resource.amount):
            res.amount = Primitive.to_proto(resource.amount)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsQuotaRequestDeltas(
            metric_name=resource.metric_name, amount=resource.amount,
        )


class InstanceReadonlyRecipeStepsQuotaRequestDeltasArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsQuotaRequestDeltas.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsQuotaRequestDeltas.from_proto(i)
            for i in resources
        ]


class InstanceReadonlyRecipeStepsPreprocessUpdate(object):
    def __init__(
        self, latency_slo_bucket_name: str = None, public_operation_metadata: str = None
    ):
        self.latency_slo_bucket_name = latency_slo_bucket_name
        self.public_operation_metadata = public_operation_metadata

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdate()
        if Primitive.to_proto(resource.latency_slo_bucket_name):
            res.latency_slo_bucket_name = Primitive.to_proto(
                resource.latency_slo_bucket_name
            )
        if Primitive.to_proto(resource.public_operation_metadata):
            res.public_operation_metadata = Primitive.to_proto(
                resource.public_operation_metadata
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsPreprocessUpdate(
            latency_slo_bucket_name=resource.latency_slo_bucket_name,
            public_operation_metadata=resource.public_operation_metadata,
        )


class InstanceReadonlyRecipeStepsPreprocessUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsPreprocessUpdate.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsPreprocessUpdate.from_proto(i) for i in resources
        ]


class InstanceReadonlyRecipeStepsRequestedTenantProject(object):
    def __init__(self, tag: str = None, folder: str = None, scope: str = None):
        self.tag = tag
        self.folder = folder
        self.scope = scope

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProject()
        if Primitive.to_proto(resource.tag):
            res.tag = Primitive.to_proto(resource.tag)
        if Primitive.to_proto(resource.folder):
            res.folder = Primitive.to_proto(resource.folder)
        if InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
            resource.scope
        ):
            res.scope = InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum.to_proto(
                resource.scope
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsRequestedTenantProject(
            tag=resource.tag, folder=resource.folder, scope=resource.scope,
        )


class InstanceReadonlyRecipeStepsRequestedTenantProjectArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsRequestedTenantProject.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsRequestedTenantProject.from_proto(i)
            for i in resources
        ]


class InstanceReadonlyRecipeStepsPermissionsInfo(object):
    def __init__(
        self,
        policy_name: dict = None,
        iam_permissions: list = None,
        resource_path: str = None,
        api_attrs: dict = None,
    ):
        self.policy_name = policy_name
        self.iam_permissions = iam_permissions
        self.resource_path = resource_path
        self.api_attrs = api_attrs

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfo()
        if InstanceReadonlyRecipeStepsPermissionsInfoPolicyName.to_proto(
            resource.policy_name
        ):
            res.policy_name.CopyFrom(
                InstanceReadonlyRecipeStepsPermissionsInfoPolicyName.to_proto(
                    resource.policy_name
                )
            )
        else:
            res.ClearField("policy_name")
        if InstanceReadonlyRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
            resource.iam_permissions
        ):
            res.iam_permissions.extend(
                InstanceReadonlyRecipeStepsPermissionsInfoIamPermissionsArray.to_proto(
                    resource.iam_permissions
                )
            )
        if Primitive.to_proto(resource.resource_path):
            res.resource_path = Primitive.to_proto(resource.resource_path)
        if InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs.to_proto(
            resource.api_attrs
        ):
            res.api_attrs.CopyFrom(
                InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs.to_proto(
                    resource.api_attrs
                )
            )
        else:
            res.ClearField("api_attrs")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsPermissionsInfo(
            policy_name=resource.policy_name,
            iam_permissions=resource.iam_permissions,
            resource_path=resource.resource_path,
            api_attrs=resource.api_attrs,
        )


class InstanceReadonlyRecipeStepsPermissionsInfoArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsPermissionsInfo.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsPermissionsInfo.from_proto(i) for i in resources
        ]


class InstanceReadonlyRecipeStepsPermissionsInfoPolicyName(object):
    def __init__(self, type: str = None, id: str = None, region: str = None):
        self.type = type
        self.id = id
        self.region = region

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyName()
        )
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if Primitive.to_proto(resource.id):
            res.id = Primitive.to_proto(resource.id)
        if Primitive.to_proto(resource.region):
            res.region = Primitive.to_proto(resource.region)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsPermissionsInfoPolicyName(
            type=resource.type, id=resource.id, region=resource.region,
        )


class InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsPermissionsInfoPolicyName.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsPermissionsInfoPolicyName.from_proto(i)
            for i in resources
        ]


class InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions(object):
    def __init__(self, permission: str = None):
        self.permission = permission

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions()
        )
        if Primitive.to_proto(resource.permission):
            res.permission = Primitive.to_proto(resource.permission)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions(
            permission=resource.permission,
        )


class InstanceReadonlyRecipeStepsPermissionsInfoIamPermissionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions.from_proto(i)
            for i in resources
        ]


class InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoApiAttrs()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs()


class InstanceReadonlyRecipeStepsPermissionsInfoApiAttrsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs.from_proto(i)
            for i in resources
        ]


class InstanceReadonlyRecipeStepsKeyNotificationsUpdate(object):
    def __init__(self, key_notifications_info: dict = None):
        self.key_notifications_info = key_notifications_info

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdate()
        if InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
            resource.key_notifications_info
        ):
            res.key_notifications_info.CopyFrom(
                InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                    resource.key_notifications_info
                )
            )
        else:
            res.ClearField("key_notifications_info")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsKeyNotificationsUpdate(
            key_notifications_info=resource.key_notifications_info,
        )


class InstanceReadonlyRecipeStepsKeyNotificationsUpdateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsKeyNotificationsUpdate.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsKeyNotificationsUpdate.from_proto(i)
            for i in resources
        ]


class InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(object):
    def __init__(
        self, key_configs: dict = None, data_version: int = None, delegate: str = None
    ):
        self.key_configs = key_configs
        self.data_version = data_version
        self.delegate = delegate

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo()
        )
        if InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
            resource.key_configs
        ):
            res.key_configs.CopyFrom(
                InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                    resource.key_configs
                )
            )
        else:
            res.ClearField("key_configs")
        if Primitive.to_proto(resource.data_version):
            res.data_version = Primitive.to_proto(resource.data_version)
        if Primitive.to_proto(resource.delegate):
            res.delegate = Primitive.to_proto(resource.delegate)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(
            key_configs=resource.key_configs,
            data_version=resource.data_version,
            delegate=resource.delegate,
        )


class InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo.from_proto(
                i
            )
            for i in resources
        ]


class InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
    object
):
    def __init__(self, key_config: dict = None):
        self.key_config = key_config

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs()
        )
        if InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
            resource.key_config
        ):
            res.key_config.CopyFrom(
                InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                    resource.key_config
                )
            )
        else:
            res.ClearField("key_config")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(
            key_config=resource.key_config,
        )


class InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs.from_proto(
                i
            )
            for i in resources
        ]


class InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
    object
):
    def __init__(self, key_or_version_name: str = None):
        self.key_or_version_name = key_or_version_name

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig()
        )
        if Primitive.to_proto(resource.key_or_version_name):
            res.key_or_version_name = Primitive.to_proto(resource.key_or_version_name)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(
            key_or_version_name=resource.key_or_version_name,
        )


class InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig.from_proto(
                i
            )
            for i in resources
        ]


class InstanceHistory(object):
    def __init__(
        self,
        timestamp: str = None,
        operation_handle: str = None,
        description: str = None,
        step_index: int = None,
        tenant_project_number: int = None,
        tenant_project_id: str = None,
        p4_service_account: str = None,
    ):
        self.timestamp = timestamp
        self.operation_handle = operation_handle
        self.description = description
        self.step_index = step_index
        self.tenant_project_number = tenant_project_number
        self.tenant_project_id = tenant_project_id
        self.p4_service_account = p4_service_account

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = instance_pb2.Tier2AlphaInstanceHistory()
        if Primitive.to_proto(resource.timestamp):
            res.timestamp = Primitive.to_proto(resource.timestamp)
        if Primitive.to_proto(resource.operation_handle):
            res.operation_handle = Primitive.to_proto(resource.operation_handle)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.step_index):
            res.step_index = Primitive.to_proto(resource.step_index)
        if Primitive.to_proto(resource.tenant_project_number):
            res.tenant_project_number = Primitive.to_proto(
                resource.tenant_project_number
            )
        if Primitive.to_proto(resource.tenant_project_id):
            res.tenant_project_id = Primitive.to_proto(resource.tenant_project_id)
        if Primitive.to_proto(resource.p4_service_account):
            res.p4_service_account = Primitive.to_proto(resource.p4_service_account)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return InstanceHistory(
            timestamp=resource.timestamp,
            operation_handle=resource.operation_handle,
            description=resource.description,
            step_index=resource.step_index,
            tenant_project_number=resource.tenant_project_number,
            tenant_project_id=resource.tenant_project_id,
            p4_service_account=resource.p4_service_account,
        )


class InstanceHistoryArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [InstanceHistory.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [InstanceHistory.from_proto(i) for i in resources]


class InstanceSkuTierEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceSkuTierEnum.Value(
            "Tier2AlphaInstanceSkuTierEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceSkuTierEnum.Name(resource)[
            len("Tier2AlphaInstanceSkuTierEnum") :
        ]


class InstanceSkuSizeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceSkuSizeEnum.Value(
            "Tier2AlphaInstanceSkuSizeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceSkuSizeEnum.Name(resource)[
            len("Tier2AlphaInstanceSkuSizeEnum") :
        ]


class InstanceStateEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceStateEnum.Value(
            "Tier2AlphaInstanceStateEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceStateEnum.Name(resource)[
            len("Tier2AlphaInstanceStateEnum") :
        ]


class InstancePreprocessCreateRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum") :
        ]


class InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceInitiateCreateRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceInitiateCreateRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceInitiateCreateRecipeStepsActionEnum") :]


class InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceCreateRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceCreateRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceCreateRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceCreateRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceCreateRecipeStepsActionEnum") :]


class InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum") :
        ]


class InstanceDeleteRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceDeleteRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceDeleteRecipeStepsActionEnum") :]


class InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum") :
        ]


class InstanceUpdateRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceUpdateRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceUpdateRecipeStepsActionEnum") :]


class InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum") :
        ]


class InstancePreprocessResetRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum") :]


class InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceInitiateResetRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceInitiateResetRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceInitiateResetRecipeStepsActionEnum") :]


class InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceResetRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceResetRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceResetRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceResetRecipeStepsActionEnum.Name(resource)[
            len("Tier2AlphaInstanceResetRecipeStepsActionEnum") :
        ]


class InstanceResetRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum") :
        ]


class InstancePreprocessRepairRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum") :
        ]


class InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceInitiateRepairRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceInitiateRepairRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceInitiateRepairRecipeStepsActionEnum") :]


class InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceRepairRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceRepairRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceRepairRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceRepairRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceRepairRecipeStepsActionEnum") :]


class InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum") :
        ]


class InstancePreprocessDeleteRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum") :
        ]


class InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceInitiateDeleteRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceInitiateDeleteRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceInitiateDeleteRecipeStepsActionEnum") :]


class InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstancePreprocessUpdateRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum") :
        ]


class InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceInitiateUpdateRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceInitiateUpdateRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceInitiateUpdateRecipeStepsActionEnum") :]


class InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstancePreprocessFreezeRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum") :
        ]


class InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceFreezeRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceFreezeRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceFreezeRecipeStepsActionEnum") :]


class InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len("Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum") :
        ]


class InstancePreprocessUnfreezeRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum.Value(
            "Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum.Name(
            resource
        )[
            len("Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum") :
        ]


class InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceUnfreezeRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum") :]


class InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum"
            ) :
        ]


class InstanceReadonlyRecipeStepsActionEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum.Value(
            "Tier2AlphaInstanceReadonlyRecipeStepsActionEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum.Name(
            resource
        )[len("Tier2AlphaInstanceReadonlyRecipeStepsActionEnum") :]


class InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum.Value(
            "Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return instance_pb2.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum.Name(
            resource
        )[
            len(
                "Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum"
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
