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
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/tier2/alpha/tier2_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/tier2/alpha"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceSkuTierEnum converts a InstanceSkuTierEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceSkuTierEnum(e alphapb.Tier2AlphaInstanceSkuTierEnum) *alpha.InstanceSkuTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceSkuTierEnum_name[int32(e)]; ok {
		e := alpha.InstanceSkuTierEnum(n[len("InstanceSkuTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSkuSizeEnum converts a InstanceSkuSizeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceSkuSizeEnum(e alphapb.Tier2AlphaInstanceSkuSizeEnum) *alpha.InstanceSkuSizeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceSkuSizeEnum_name[int32(e)]; ok {
		e := alpha.InstanceSkuSizeEnum(n[len("InstanceSkuSizeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceStateEnum(e alphapb.Tier2AlphaInstanceStateEnum) *alpha.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceStateEnum_name[int32(e)]; ok {
		e := alpha.InstanceStateEnum(n[len("InstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessCreateRecipeStepsActionEnum converts a InstancePreprocessCreateRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum) *alpha.InstancePreprocessCreateRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessCreateRecipeStepsActionEnum(n[len("InstancePreprocessCreateRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceInitiateCreateRecipeStepsActionEnum converts a InstanceInitiateCreateRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsActionEnum) *alpha.InstanceInitiateCreateRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceInitiateCreateRecipeStepsActionEnum(n[len("InstanceInitiateCreateRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceCreateRecipeStepsActionEnum converts a InstanceCreateRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceCreateRecipeStepsActionEnum) *alpha.InstanceCreateRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceCreateRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceCreateRecipeStepsActionEnum(n[len("InstanceCreateRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDeleteRecipeStepsActionEnum converts a InstanceDeleteRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceDeleteRecipeStepsActionEnum) *alpha.InstanceDeleteRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceDeleteRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceDeleteRecipeStepsActionEnum(n[len("InstanceDeleteRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceUpdateRecipeStepsActionEnum converts a InstanceUpdateRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceUpdateRecipeStepsActionEnum) *alpha.InstanceUpdateRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceUpdateRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceUpdateRecipeStepsActionEnum(n[len("InstanceUpdateRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessResetRecipeStepsActionEnum converts a InstancePreprocessResetRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum) *alpha.InstancePreprocessResetRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessResetRecipeStepsActionEnum(n[len("InstancePreprocessResetRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceInitiateResetRecipeStepsActionEnum converts a InstanceInitiateResetRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsActionEnum) *alpha.InstanceInitiateResetRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceInitiateResetRecipeStepsActionEnum(n[len("InstanceInitiateResetRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceResetRecipeStepsActionEnum converts a InstanceResetRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceResetRecipeStepsActionEnum) *alpha.InstanceResetRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceResetRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceResetRecipeStepsActionEnum(n[len("InstanceResetRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceResetRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceResetRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceResetRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceResetRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstanceResetRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessRepairRecipeStepsActionEnum converts a InstancePreprocessRepairRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum) *alpha.InstancePreprocessRepairRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessRepairRecipeStepsActionEnum(n[len("InstancePreprocessRepairRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceInitiateRepairRecipeStepsActionEnum converts a InstanceInitiateRepairRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsActionEnum) *alpha.InstanceInitiateRepairRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceInitiateRepairRecipeStepsActionEnum(n[len("InstanceInitiateRepairRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceRepairRecipeStepsActionEnum converts a InstanceRepairRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceRepairRecipeStepsActionEnum) *alpha.InstanceRepairRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceRepairRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceRepairRecipeStepsActionEnum(n[len("InstanceRepairRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessDeleteRecipeStepsActionEnum converts a InstancePreprocessDeleteRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum) *alpha.InstancePreprocessDeleteRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessDeleteRecipeStepsActionEnum(n[len("InstancePreprocessDeleteRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceInitiateDeleteRecipeStepsActionEnum converts a InstanceInitiateDeleteRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsActionEnum) *alpha.InstanceInitiateDeleteRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceInitiateDeleteRecipeStepsActionEnum(n[len("InstanceInitiateDeleteRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessUpdateRecipeStepsActionEnum converts a InstancePreprocessUpdateRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum) *alpha.InstancePreprocessUpdateRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessUpdateRecipeStepsActionEnum(n[len("InstancePreprocessUpdateRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceInitiateUpdateRecipeStepsActionEnum converts a InstanceInitiateUpdateRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsActionEnum) *alpha.InstanceInitiateUpdateRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceInitiateUpdateRecipeStepsActionEnum(n[len("InstanceInitiateUpdateRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessFreezeRecipeStepsActionEnum converts a InstancePreprocessFreezeRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum) *alpha.InstancePreprocessFreezeRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessFreezeRecipeStepsActionEnum(n[len("InstancePreprocessFreezeRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFreezeRecipeStepsActionEnum converts a InstanceFreezeRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceFreezeRecipeStepsActionEnum) *alpha.InstanceFreezeRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceFreezeRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceFreezeRecipeStepsActionEnum(n[len("InstanceFreezeRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsActionEnum converts a InstancePreprocessUnfreezeRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum(e alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum) *alpha.InstancePreprocessUnfreezeRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessUnfreezeRecipeStepsActionEnum(n[len("InstancePreprocessUnfreezeRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum converts a InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceUnfreezeRecipeStepsActionEnum converts a InstanceUnfreezeRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum) *alpha.InstanceUnfreezeRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceUnfreezeRecipeStepsActionEnum(n[len("InstanceUnfreezeRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceReadonlyRecipeStepsActionEnum converts a InstanceReadonlyRecipeStepsActionEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsActionEnum(e alphapb.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum) *alpha.InstanceReadonlyRecipeStepsActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceReadonlyRecipeStepsActionEnum(n[len("InstanceReadonlyRecipeStepsActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum converts a InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum enum from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(e alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum) *alpha.InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum_name[int32(e)]; ok {
		e := alpha.InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(n[len("InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceSku converts a InstanceSku resource from its proto representation.
func ProtoToTier2AlphaInstanceSku(p *alphapb.Tier2AlphaInstanceSku) *alpha.InstanceSku {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceSku{
		Tier: ProtoToTier2AlphaInstanceSkuTierEnum(p.GetTier()),
		Size: ProtoToTier2AlphaInstanceSkuSizeEnum(p.GetSize()),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipe converts a InstancePreprocessCreateRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipe(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipe) *alpha.InstancePreprocessCreateRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessCreateRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeSteps converts a InstancePreprocessCreateRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeSteps) *alpha.InstancePreprocessCreateRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsStatus converts a InstancePreprocessCreateRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsStatus) *alpha.InstancePreprocessCreateRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsStatusDetails converts a InstancePreprocessCreateRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsStatusDetails) *alpha.InstancePreprocessCreateRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsQuotaRequestDeltas converts a InstancePreprocessCreateRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessCreateRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsPreprocessUpdate converts a InstancePreprocessCreateRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessCreateRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsRequestedTenantProject converts a InstancePreprocessCreateRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessCreateRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsPermissionsInfo converts a InstancePreprocessCreateRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfo) *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs converts a InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs) *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstanceInitiateCreateRecipe converts a InstanceInitiateCreateRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipe(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipe) *alpha.InstanceInitiateCreateRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceInitiateCreateRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceInitiateCreateRecipeSteps converts a InstanceInitiateCreateRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeSteps(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipeSteps) *alpha.InstanceInitiateCreateRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceInitiateCreateRecipeStepsStatus converts a InstanceInitiateCreateRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsStatus) *alpha.InstanceInitiateCreateRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceInitiateCreateRecipeStepsStatusDetails converts a InstanceInitiateCreateRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsStatusDetails) *alpha.InstanceInitiateCreateRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceInitiateCreateRecipeStepsQuotaRequestDeltas converts a InstanceInitiateCreateRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsQuotaRequestDeltas) *alpha.InstanceInitiateCreateRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstanceInitiateCreateRecipeStepsPreprocessUpdate converts a InstanceInitiateCreateRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPreprocessUpdate) *alpha.InstanceInitiateCreateRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceInitiateCreateRecipeStepsRequestedTenantProject converts a InstanceInitiateCreateRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProject) *alpha.InstanceInitiateCreateRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceInitiateCreateRecipeStepsPermissionsInfo converts a InstanceInitiateCreateRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfo) *alpha.InstanceInitiateCreateRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName converts a InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions converts a InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs converts a InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs) *alpha.InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstanceInitiateCreateRecipeStepsKeyNotificationsUpdate converts a InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdate) *alpha.InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstanceCreateRecipe converts a InstanceCreateRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipe(p *alphapb.Tier2AlphaInstanceCreateRecipe) *alpha.InstanceCreateRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceCreateRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceCreateRecipeSteps converts a InstanceCreateRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeSteps(p *alphapb.Tier2AlphaInstanceCreateRecipeSteps) *alpha.InstanceCreateRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceCreateRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceCreateRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceCreateRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceCreateRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsStatus converts a InstanceCreateRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsStatus) *alpha.InstanceCreateRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceCreateRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsStatusDetails converts a InstanceCreateRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsStatusDetails) *alpha.InstanceCreateRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsQuotaRequestDeltas converts a InstanceCreateRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltas) *alpha.InstanceCreateRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsPreprocessUpdate converts a InstanceCreateRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsPreprocessUpdate) *alpha.InstanceCreateRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsRequestedTenantProject converts a InstanceCreateRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProject) *alpha.InstanceCreateRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsPermissionsInfo converts a InstanceCreateRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfo) *alpha.InstanceCreateRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsPermissionsInfoPolicyName converts a InstanceCreateRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceCreateRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsPermissionsInfoIamPermissions converts a InstanceCreateRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceCreateRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsPermissionsInfoApiAttrs converts a InstanceCreateRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoApiAttrs) *alpha.InstanceCreateRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstanceCreateRecipeStepsKeyNotificationsUpdate converts a InstanceCreateRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdate) *alpha.InstanceCreateRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstanceDeleteRecipe converts a InstanceDeleteRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipe(p *alphapb.Tier2AlphaInstanceDeleteRecipe) *alpha.InstanceDeleteRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceDeleteRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceDeleteRecipeSteps converts a InstanceDeleteRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeSteps(p *alphapb.Tier2AlphaInstanceDeleteRecipeSteps) *alpha.InstanceDeleteRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceDeleteRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceDeleteRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceDeleteRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceDeleteRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsStatus converts a InstanceDeleteRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsStatus) *alpha.InstanceDeleteRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceDeleteRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsStatusDetails converts a InstanceDeleteRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsStatusDetails) *alpha.InstanceDeleteRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsQuotaRequestDeltas converts a InstanceDeleteRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltas) *alpha.InstanceDeleteRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsPreprocessUpdate converts a InstanceDeleteRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPreprocessUpdate) *alpha.InstanceDeleteRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsRequestedTenantProject converts a InstanceDeleteRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProject) *alpha.InstanceDeleteRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsPermissionsInfo converts a InstanceDeleteRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfo) *alpha.InstanceDeleteRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsPermissionsInfoPolicyName converts a InstanceDeleteRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceDeleteRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsPermissionsInfoIamPermissions converts a InstanceDeleteRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceDeleteRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsPermissionsInfoApiAttrs converts a InstanceDeleteRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoApiAttrs) *alpha.InstanceDeleteRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsKeyNotificationsUpdate converts a InstanceDeleteRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdate) *alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstanceUpdateRecipe converts a InstanceUpdateRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipe(p *alphapb.Tier2AlphaInstanceUpdateRecipe) *alpha.InstanceUpdateRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceUpdateRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceUpdateRecipeSteps converts a InstanceUpdateRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeSteps(p *alphapb.Tier2AlphaInstanceUpdateRecipeSteps) *alpha.InstanceUpdateRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceUpdateRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceUpdateRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceUpdateRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceUpdateRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsStatus converts a InstanceUpdateRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsStatus) *alpha.InstanceUpdateRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceUpdateRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsStatusDetails converts a InstanceUpdateRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsStatusDetails) *alpha.InstanceUpdateRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsQuotaRequestDeltas converts a InstanceUpdateRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltas) *alpha.InstanceUpdateRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsPreprocessUpdate converts a InstanceUpdateRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPreprocessUpdate) *alpha.InstanceUpdateRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsRequestedTenantProject converts a InstanceUpdateRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProject) *alpha.InstanceUpdateRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsPermissionsInfo converts a InstanceUpdateRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfo) *alpha.InstanceUpdateRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsPermissionsInfoPolicyName converts a InstanceUpdateRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceUpdateRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsPermissionsInfoIamPermissions converts a InstanceUpdateRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceUpdateRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsPermissionsInfoApiAttrs converts a InstanceUpdateRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoApiAttrs) *alpha.InstanceUpdateRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsKeyNotificationsUpdate converts a InstanceUpdateRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdate) *alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipe converts a InstancePreprocessResetRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipe(p *alphapb.Tier2AlphaInstancePreprocessResetRecipe) *alpha.InstancePreprocessResetRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessResetRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeSteps converts a InstancePreprocessResetRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeSteps) *alpha.InstancePreprocessResetRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessResetRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessResetRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsStatus converts a InstancePreprocessResetRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsStatus) *alpha.InstancePreprocessResetRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessResetRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsStatusDetails converts a InstancePreprocessResetRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsStatusDetails) *alpha.InstancePreprocessResetRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsQuotaRequestDeltas converts a InstancePreprocessResetRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessResetRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsPreprocessUpdate converts a InstancePreprocessResetRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessResetRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsRequestedTenantProject converts a InstancePreprocessResetRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessResetRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsPermissionsInfo converts a InstancePreprocessResetRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfo) *alpha.InstancePreprocessResetRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs converts a InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs) *alpha.InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessResetRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstanceInitiateResetRecipe converts a InstanceInitiateResetRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipe(p *alphapb.Tier2AlphaInstanceInitiateResetRecipe) *alpha.InstanceInitiateResetRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceInitiateResetRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceInitiateResetRecipeSteps converts a InstanceInitiateResetRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeSteps(p *alphapb.Tier2AlphaInstanceInitiateResetRecipeSteps) *alpha.InstanceInitiateResetRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceInitiateResetRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceInitiateResetRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceInitiateResetRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceInitiateResetRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceInitiateResetRecipeStepsStatus converts a InstanceInitiateResetRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsStatus) *alpha.InstanceInitiateResetRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceInitiateResetRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceInitiateResetRecipeStepsStatusDetails converts a InstanceInitiateResetRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsStatusDetails) *alpha.InstanceInitiateResetRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceInitiateResetRecipeStepsQuotaRequestDeltas converts a InstanceInitiateResetRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsQuotaRequestDeltas) *alpha.InstanceInitiateResetRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstanceInitiateResetRecipeStepsPreprocessUpdate converts a InstanceInitiateResetRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPreprocessUpdate) *alpha.InstanceInitiateResetRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceInitiateResetRecipeStepsRequestedTenantProject converts a InstanceInitiateResetRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProject) *alpha.InstanceInitiateResetRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceInitiateResetRecipeStepsPermissionsInfo converts a InstanceInitiateResetRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfo) *alpha.InstanceInitiateResetRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceInitiateResetRecipeStepsPermissionsInfoPolicyName converts a InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions converts a InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs converts a InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs) *alpha.InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstanceInitiateResetRecipeStepsKeyNotificationsUpdate converts a InstanceInitiateResetRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdate) *alpha.InstanceInitiateResetRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstanceResetRecipe converts a InstanceResetRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipe(p *alphapb.Tier2AlphaInstanceResetRecipe) *alpha.InstanceResetRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceResetRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceResetRecipeSteps converts a InstanceResetRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeSteps(p *alphapb.Tier2AlphaInstanceResetRecipeSteps) *alpha.InstanceResetRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceResetRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceResetRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceResetRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceResetRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceResetRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsStatus converts a InstanceResetRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceResetRecipeStepsStatus) *alpha.InstanceResetRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceResetRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsStatusDetails converts a InstanceResetRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceResetRecipeStepsStatusDetails) *alpha.InstanceResetRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsQuotaRequestDeltas converts a InstanceResetRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceResetRecipeStepsQuotaRequestDeltas) *alpha.InstanceResetRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsPreprocessUpdate converts a InstanceResetRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceResetRecipeStepsPreprocessUpdate) *alpha.InstanceResetRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsRequestedTenantProject converts a InstanceResetRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProject) *alpha.InstanceResetRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsPermissionsInfo converts a InstanceResetRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfo) *alpha.InstanceResetRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsPermissionsInfoPolicyName converts a InstanceResetRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceResetRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsPermissionsInfoIamPermissions converts a InstanceResetRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceResetRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsPermissionsInfoApiAttrs converts a InstanceResetRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoApiAttrs) *alpha.InstanceResetRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstanceResetRecipeStepsKeyNotificationsUpdate converts a InstanceResetRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdate) *alpha.InstanceResetRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipe converts a InstancePreprocessRepairRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipe(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipe) *alpha.InstancePreprocessRepairRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessRepairRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeSteps converts a InstancePreprocessRepairRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeSteps) *alpha.InstancePreprocessRepairRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsStatus converts a InstancePreprocessRepairRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsStatus) *alpha.InstancePreprocessRepairRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsStatusDetails converts a InstancePreprocessRepairRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsStatusDetails) *alpha.InstancePreprocessRepairRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsQuotaRequestDeltas converts a InstancePreprocessRepairRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessRepairRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsPreprocessUpdate converts a InstancePreprocessRepairRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessRepairRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsRequestedTenantProject converts a InstancePreprocessRepairRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessRepairRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsPermissionsInfo converts a InstancePreprocessRepairRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfo) *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs converts a InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs) *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstanceInitiateRepairRecipe converts a InstanceInitiateRepairRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipe(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipe) *alpha.InstanceInitiateRepairRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceInitiateRepairRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceInitiateRepairRecipeSteps converts a InstanceInitiateRepairRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeSteps(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipeSteps) *alpha.InstanceInitiateRepairRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceInitiateRepairRecipeStepsStatus converts a InstanceInitiateRepairRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsStatus) *alpha.InstanceInitiateRepairRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceInitiateRepairRecipeStepsStatusDetails converts a InstanceInitiateRepairRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsStatusDetails) *alpha.InstanceInitiateRepairRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceInitiateRepairRecipeStepsQuotaRequestDeltas converts a InstanceInitiateRepairRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsQuotaRequestDeltas) *alpha.InstanceInitiateRepairRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstanceInitiateRepairRecipeStepsPreprocessUpdate converts a InstanceInitiateRepairRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPreprocessUpdate) *alpha.InstanceInitiateRepairRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceInitiateRepairRecipeStepsRequestedTenantProject converts a InstanceInitiateRepairRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProject) *alpha.InstanceInitiateRepairRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceInitiateRepairRecipeStepsPermissionsInfo converts a InstanceInitiateRepairRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfo) *alpha.InstanceInitiateRepairRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName converts a InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions converts a InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs converts a InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs) *alpha.InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstanceInitiateRepairRecipeStepsKeyNotificationsUpdate converts a InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdate) *alpha.InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstanceRepairRecipe converts a InstanceRepairRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipe(p *alphapb.Tier2AlphaInstanceRepairRecipe) *alpha.InstanceRepairRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceRepairRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceRepairRecipeSteps converts a InstanceRepairRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeSteps(p *alphapb.Tier2AlphaInstanceRepairRecipeSteps) *alpha.InstanceRepairRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceRepairRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceRepairRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceRepairRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceRepairRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsStatus converts a InstanceRepairRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsStatus) *alpha.InstanceRepairRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceRepairRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsStatusDetails converts a InstanceRepairRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsStatusDetails) *alpha.InstanceRepairRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsQuotaRequestDeltas converts a InstanceRepairRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltas) *alpha.InstanceRepairRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsPreprocessUpdate converts a InstanceRepairRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsPreprocessUpdate) *alpha.InstanceRepairRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsRequestedTenantProject converts a InstanceRepairRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProject) *alpha.InstanceRepairRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsPermissionsInfo converts a InstanceRepairRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfo) *alpha.InstanceRepairRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsPermissionsInfoPolicyName converts a InstanceRepairRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceRepairRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsPermissionsInfoIamPermissions converts a InstanceRepairRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceRepairRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsPermissionsInfoApiAttrs converts a InstanceRepairRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoApiAttrs) *alpha.InstanceRepairRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstanceRepairRecipeStepsKeyNotificationsUpdate converts a InstanceRepairRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdate) *alpha.InstanceRepairRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipe converts a InstancePreprocessDeleteRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipe(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipe) *alpha.InstancePreprocessDeleteRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessDeleteRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeSteps converts a InstancePreprocessDeleteRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeSteps) *alpha.InstancePreprocessDeleteRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsStatus converts a InstancePreprocessDeleteRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatus) *alpha.InstancePreprocessDeleteRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsStatusDetails converts a InstancePreprocessDeleteRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetails) *alpha.InstancePreprocessDeleteRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas converts a InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsPreprocessUpdate converts a InstancePreprocessDeleteRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessDeleteRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsRequestedTenantProject converts a InstancePreprocessDeleteRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessDeleteRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsPermissionsInfo converts a InstancePreprocessDeleteRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfo) *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs converts a InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs) *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipe converts a InstanceInitiateDeleteRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipe(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipe) *alpha.InstanceInitiateDeleteRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceInitiateDeleteRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipeSteps converts a InstanceInitiateDeleteRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeSteps(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeSteps) *alpha.InstanceInitiateDeleteRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipeStepsStatus converts a InstanceInitiateDeleteRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsStatus) *alpha.InstanceInitiateDeleteRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipeStepsStatusDetails converts a InstanceInitiateDeleteRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsStatusDetails) *alpha.InstanceInitiateDeleteRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipeStepsQuotaRequestDeltas converts a InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsQuotaRequestDeltas) *alpha.InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipeStepsPreprocessUpdate converts a InstanceInitiateDeleteRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPreprocessUpdate) *alpha.InstanceInitiateDeleteRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipeStepsRequestedTenantProject converts a InstanceInitiateDeleteRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProject) *alpha.InstanceInitiateDeleteRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipeStepsPermissionsInfo converts a InstanceInitiateDeleteRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfo) *alpha.InstanceInitiateDeleteRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName converts a InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions converts a InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs converts a InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs) *alpha.InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate converts a InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate) *alpha.InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipe converts a InstancePreprocessUpdateRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipe(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipe) *alpha.InstancePreprocessUpdateRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessUpdateRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeSteps converts a InstancePreprocessUpdateRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeSteps) *alpha.InstancePreprocessUpdateRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsStatus converts a InstancePreprocessUpdateRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatus) *alpha.InstancePreprocessUpdateRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsStatusDetails converts a InstancePreprocessUpdateRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetails) *alpha.InstancePreprocessUpdateRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas converts a InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsPreprocessUpdate converts a InstancePreprocessUpdateRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessUpdateRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsRequestedTenantProject converts a InstancePreprocessUpdateRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessUpdateRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsPermissionsInfo converts a InstancePreprocessUpdateRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfo) *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs converts a InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs) *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipe converts a InstanceInitiateUpdateRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipe(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipe) *alpha.InstanceInitiateUpdateRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceInitiateUpdateRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipeSteps converts a InstanceInitiateUpdateRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeSteps(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeSteps) *alpha.InstanceInitiateUpdateRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipeStepsStatus converts a InstanceInitiateUpdateRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsStatus) *alpha.InstanceInitiateUpdateRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipeStepsStatusDetails converts a InstanceInitiateUpdateRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsStatusDetails) *alpha.InstanceInitiateUpdateRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipeStepsQuotaRequestDeltas converts a InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsQuotaRequestDeltas) *alpha.InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipeStepsPreprocessUpdate converts a InstanceInitiateUpdateRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPreprocessUpdate) *alpha.InstanceInitiateUpdateRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipeStepsRequestedTenantProject converts a InstanceInitiateUpdateRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProject) *alpha.InstanceInitiateUpdateRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipeStepsPermissionsInfo converts a InstanceInitiateUpdateRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfo) *alpha.InstanceInitiateUpdateRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName converts a InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions converts a InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs converts a InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs) *alpha.InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate converts a InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate) *alpha.InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipe converts a InstancePreprocessFreezeRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipe(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipe) *alpha.InstancePreprocessFreezeRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessFreezeRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeSteps converts a InstancePreprocessFreezeRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeSteps) *alpha.InstancePreprocessFreezeRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsStatus converts a InstancePreprocessFreezeRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatus) *alpha.InstancePreprocessFreezeRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsStatusDetails converts a InstancePreprocessFreezeRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetails) *alpha.InstancePreprocessFreezeRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas converts a InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsPreprocessUpdate converts a InstancePreprocessFreezeRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessFreezeRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsRequestedTenantProject converts a InstancePreprocessFreezeRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessFreezeRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsPermissionsInfo converts a InstancePreprocessFreezeRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfo) *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs converts a InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs) *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstanceFreezeRecipe converts a InstanceFreezeRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipe(p *alphapb.Tier2AlphaInstanceFreezeRecipe) *alpha.InstanceFreezeRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceFreezeRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceFreezeRecipeSteps converts a InstanceFreezeRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeSteps(p *alphapb.Tier2AlphaInstanceFreezeRecipeSteps) *alpha.InstanceFreezeRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceFreezeRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceFreezeRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceFreezeRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceFreezeRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsStatus converts a InstanceFreezeRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsStatus) *alpha.InstanceFreezeRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceFreezeRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsStatusDetails converts a InstanceFreezeRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsStatusDetails) *alpha.InstanceFreezeRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsQuotaRequestDeltas converts a InstanceFreezeRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltas) *alpha.InstanceFreezeRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsPreprocessUpdate converts a InstanceFreezeRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPreprocessUpdate) *alpha.InstanceFreezeRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsRequestedTenantProject converts a InstanceFreezeRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProject) *alpha.InstanceFreezeRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsPermissionsInfo converts a InstanceFreezeRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfo) *alpha.InstanceFreezeRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsPermissionsInfoPolicyName converts a InstanceFreezeRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceFreezeRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsPermissionsInfoIamPermissions converts a InstanceFreezeRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceFreezeRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsPermissionsInfoApiAttrs converts a InstanceFreezeRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoApiAttrs) *alpha.InstanceFreezeRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsKeyNotificationsUpdate converts a InstanceFreezeRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdate) *alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipe converts a InstancePreprocessUnfreezeRecipe resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipe(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipe) *alpha.InstancePreprocessUnfreezeRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeSteps converts a InstancePreprocessUnfreezeRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeSteps(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeSteps) *alpha.InstancePreprocessUnfreezeRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsStatus converts a InstancePreprocessUnfreezeRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsStatus(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatus) *alpha.InstancePreprocessUnfreezeRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsStatusDetails converts a InstancePreprocessUnfreezeRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetails) *alpha.InstancePreprocessUnfreezeRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas converts a InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas) *alpha.InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate converts a InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate) *alpha.InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject converts a InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject) *alpha.InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsPermissionsInfo converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfo) *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName) *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions) *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs) *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate converts a InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate) *alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipe converts a InstanceUnfreezeRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipe(p *alphapb.Tier2AlphaInstanceUnfreezeRecipe) *alpha.InstanceUnfreezeRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceUnfreezeRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeSteps converts a InstanceUnfreezeRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeSteps(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeSteps) *alpha.InstanceUnfreezeRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceUnfreezeRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceUnfreezeRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsStatus converts a InstanceUnfreezeRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsStatus) *alpha.InstanceUnfreezeRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceUnfreezeRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsStatusDetails converts a InstanceUnfreezeRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsStatusDetails) *alpha.InstanceUnfreezeRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsQuotaRequestDeltas converts a InstanceUnfreezeRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltas) *alpha.InstanceUnfreezeRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsPreprocessUpdate converts a InstanceUnfreezeRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdate) *alpha.InstanceUnfreezeRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsRequestedTenantProject converts a InstanceUnfreezeRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProject) *alpha.InstanceUnfreezeRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsPermissionsInfo converts a InstanceUnfreezeRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfo) *alpha.InstanceUnfreezeRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName converts a InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions converts a InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs converts a InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs) *alpha.InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsKeyNotificationsUpdate converts a InstanceUnfreezeRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdate) *alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipe converts a InstanceReadonlyRecipe resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipe(p *alphapb.Tier2AlphaInstanceReadonlyRecipe) *alpha.InstanceReadonlyRecipe {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipe{
		HonorCancelRequest:                dcl.Bool(p.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.Int64OrNil(p.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.Float64OrNil(p.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.Bool(p.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.StringOrNil(p.GetReadonlyRecipeStartTime()),
		DelayToStoreResourcesInClhDbNanos: dcl.Int64OrNil(p.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range p.GetSteps() {
		obj.Steps = append(obj.Steps, *ProtoToTier2AlphaInstanceReadonlyRecipeSteps(r))
	}
	for _, r := range p.GetResourceNamesStoredInClhWithDelay() {
		obj.ResourceNamesStoredInClhWithDelay = append(obj.ResourceNamesStoredInClhWithDelay, r)
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeSteps converts a InstanceReadonlyRecipeSteps resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeSteps(p *alphapb.Tier2AlphaInstanceReadonlyRecipeSteps) *alpha.InstanceReadonlyRecipeSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeSteps{
		RelativeTime:                   dcl.Int64OrNil(p.RelativeTime),
		SleepDuration:                  dcl.Int64OrNil(p.SleepDuration),
		Action:                         ProtoToTier2AlphaInstanceReadonlyRecipeStepsActionEnum(p.GetAction()),
		Status:                         ProtoToTier2AlphaInstanceReadonlyRecipeStepsStatus(p.GetStatus()),
		ErrorSpace:                     dcl.StringOrNil(p.ErrorSpace),
		P4ServiceAccount:               dcl.StringOrNil(p.P4ServiceAccount),
		ResourceMetadataSize:           dcl.Int64OrNil(p.ResourceMetadataSize),
		Description:                    dcl.StringOrNil(p.Description),
		UpdatedRepeatOperationDelaySec: dcl.Float64OrNil(p.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               ProtoToTier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdate(p.GetPreprocessUpdate()),
		PublicOperationMetadata:        dcl.StringOrNil(p.PublicOperationMetadata),
		RequestedTenantProject:         ProtoToTier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProject(p.GetRequestedTenantProject()),
		KeyNotificationsUpdate:         ProtoToTier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdate(p.GetKeyNotificationsUpdate()),
		ClhDataUpdateTime:              dcl.StringOrNil(p.GetClhDataUpdateTime()),
	}
	for _, r := range p.GetQuotaRequestDeltas() {
		obj.QuotaRequestDeltas = append(obj.QuotaRequestDeltas, *ProtoToTier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltas(r))
	}
	for _, r := range p.GetPermissionsInfo() {
		obj.PermissionsInfo = append(obj.PermissionsInfo, *ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfo(r))
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsStatus converts a InstanceReadonlyRecipeStepsStatus resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsStatus(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsStatus) *alpha.InstanceReadonlyRecipeStepsStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsStatus{
		Code:    dcl.Int64OrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetDetails() {
		obj.Details = append(obj.Details, *ProtoToTier2AlphaInstanceReadonlyRecipeStepsStatusDetails(r))
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsStatusDetails converts a InstanceReadonlyRecipeStepsStatusDetails resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsStatusDetails(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsStatusDetails) *alpha.InstanceReadonlyRecipeStepsStatusDetails {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsStatusDetails{
		TypeUrl: dcl.StringOrNil(p.TypeUrl),
		Value:   dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsQuotaRequestDeltas converts a InstanceReadonlyRecipeStepsQuotaRequestDeltas resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltas(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltas) *alpha.InstanceReadonlyRecipeStepsQuotaRequestDeltas {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.StringOrNil(p.MetricName),
		Amount:     dcl.Int64OrNil(p.Amount),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsPreprocessUpdate converts a InstanceReadonlyRecipeStepsPreprocessUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdate(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdate) *alpha.InstanceReadonlyRecipeStepsPreprocessUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.StringOrNil(p.LatencySloBucketName),
		PublicOperationMetadata: dcl.StringOrNil(p.PublicOperationMetadata),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsRequestedTenantProject converts a InstanceReadonlyRecipeStepsRequestedTenantProject resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProject(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProject) *alpha.InstanceReadonlyRecipeStepsRequestedTenantProject {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsRequestedTenantProject{
		Tag:    dcl.StringOrNil(p.Tag),
		Folder: dcl.StringOrNil(p.Folder),
		Scope:  ProtoToTier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(p.GetScope()),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsPermissionsInfo converts a InstanceReadonlyRecipeStepsPermissionsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfo(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfo) *alpha.InstanceReadonlyRecipeStepsPermissionsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsPermissionsInfo{
		PolicyName:   ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyName(p.GetPolicyName()),
		ResourcePath: dcl.StringOrNil(p.ResourcePath),
		ApiAttrs:     ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoApiAttrs(p.GetApiAttrs()),
	}
	for _, r := range p.GetIamPermissions() {
		obj.IamPermissions = append(obj.IamPermissions, *ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions(r))
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsPermissionsInfoPolicyName converts a InstanceReadonlyRecipeStepsPermissionsInfoPolicyName resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyName(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyName) *alpha.InstanceReadonlyRecipeStepsPermissionsInfoPolicyName {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.StringOrNil(p.Type),
		Id:     dcl.StringOrNil(p.Id),
		Region: dcl.StringOrNil(p.Region),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions converts a InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions) *alpha.InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.StringOrNil(p.Permission),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsPermissionsInfoApiAttrs converts a InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoApiAttrs(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoApiAttrs) *alpha.InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs{}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsKeyNotificationsUpdate converts a InstanceReadonlyRecipeStepsKeyNotificationsUpdate resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdate(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdate) *alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdate {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: ProtoToTier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p.GetKeyNotificationsInfo()),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo converts a InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  ProtoToTier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p.GetKeyConfigs()),
		DataVersion: dcl.Int64OrNil(p.DataVersion),
		Delegate:    dcl.StringOrNil(p.Delegate),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs converts a InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: ProtoToTier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p.GetKeyConfig()),
	}
	return obj
}

// ProtoToInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig converts a InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource from its proto representation.
func ProtoToTier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig(p *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.StringOrNil(p.KeyOrVersionName),
	}
	return obj
}

// ProtoToInstanceHistory converts a InstanceHistory resource from its proto representation.
func ProtoToTier2AlphaInstanceHistory(p *alphapb.Tier2AlphaInstanceHistory) *alpha.InstanceHistory {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceHistory{
		Timestamp:           dcl.StringOrNil(p.GetTimestamp()),
		OperationHandle:     dcl.StringOrNil(p.OperationHandle),
		Description:         dcl.StringOrNil(p.Description),
		StepIndex:           dcl.Int64OrNil(p.StepIndex),
		TenantProjectNumber: dcl.Int64OrNil(p.TenantProjectNumber),
		TenantProjectId:     dcl.StringOrNil(p.TenantProjectId),
		P4ServiceAccount:    dcl.StringOrNil(p.P4ServiceAccount),
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *alphapb.Tier2AlphaInstance) *alpha.Instance {
	obj := &alpha.Instance{
		Name:                       dcl.StringOrNil(p.Name),
		DisplayName:                dcl.StringOrNil(p.DisplayName),
		Zone:                       dcl.StringOrNil(p.Zone),
		Sku:                        ProtoToTier2AlphaInstanceSku(p.GetSku()),
		AuthorizedNetworkId:        dcl.StringOrNil(p.AuthorizedNetworkId),
		ReservedIPRange:            dcl.StringOrNil(p.ReservedIpRange),
		HostName:                   dcl.StringOrNil(p.HostName),
		PortNumber:                 dcl.Int64OrNil(p.PortNumber),
		CurrentZone:                dcl.StringOrNil(p.CurrentZone),
		CreationTime:               dcl.StringOrNil(p.GetCreationTime()),
		State:                      ProtoToTier2AlphaInstanceStateEnum(p.GetState()),
		StatusMessage:              dcl.StringOrNil(p.StatusMessage),
		ExtraField:                 dcl.StringOrNil(p.ExtraField),
		PreprocessCreateRecipe:     ProtoToTier2AlphaInstancePreprocessCreateRecipe(p.GetPreprocessCreateRecipe()),
		InitiateCreateRecipe:       ProtoToTier2AlphaInstanceInitiateCreateRecipe(p.GetInitiateCreateRecipe()),
		CreateRecipe:               ProtoToTier2AlphaInstanceCreateRecipe(p.GetCreateRecipe()),
		DeleteRecipe:               ProtoToTier2AlphaInstanceDeleteRecipe(p.GetDeleteRecipe()),
		UpdateRecipe:               ProtoToTier2AlphaInstanceUpdateRecipe(p.GetUpdateRecipe()),
		PreprocessResetRecipe:      ProtoToTier2AlphaInstancePreprocessResetRecipe(p.GetPreprocessResetRecipe()),
		InitiateResetRecipe:        ProtoToTier2AlphaInstanceInitiateResetRecipe(p.GetInitiateResetRecipe()),
		ResetRecipe:                ProtoToTier2AlphaInstanceResetRecipe(p.GetResetRecipe()),
		PreprocessRepairRecipe:     ProtoToTier2AlphaInstancePreprocessRepairRecipe(p.GetPreprocessRepairRecipe()),
		InitiateRepairRecipe:       ProtoToTier2AlphaInstanceInitiateRepairRecipe(p.GetInitiateRepairRecipe()),
		RepairRecipe:               ProtoToTier2AlphaInstanceRepairRecipe(p.GetRepairRecipe()),
		PreprocessDeleteRecipe:     ProtoToTier2AlphaInstancePreprocessDeleteRecipe(p.GetPreprocessDeleteRecipe()),
		InitiateDeleteRecipe:       ProtoToTier2AlphaInstanceInitiateDeleteRecipe(p.GetInitiateDeleteRecipe()),
		PreprocessUpdateRecipe:     ProtoToTier2AlphaInstancePreprocessUpdateRecipe(p.GetPreprocessUpdateRecipe()),
		InitiateUpdateRecipe:       ProtoToTier2AlphaInstanceInitiateUpdateRecipe(p.GetInitiateUpdateRecipe()),
		PreprocessFreezeRecipe:     ProtoToTier2AlphaInstancePreprocessFreezeRecipe(p.GetPreprocessFreezeRecipe()),
		FreezeRecipe:               ProtoToTier2AlphaInstanceFreezeRecipe(p.GetFreezeRecipe()),
		PreprocessUnfreezeRecipe:   ProtoToTier2AlphaInstancePreprocessUnfreezeRecipe(p.GetPreprocessUnfreezeRecipe()),
		UnfreezeRecipe:             ProtoToTier2AlphaInstanceUnfreezeRecipe(p.GetUnfreezeRecipe()),
		ReadonlyRecipe:             ProtoToTier2AlphaInstanceReadonlyRecipe(p.GetReadonlyRecipe()),
		EnableCallHistory:          dcl.Bool(p.EnableCallHistory),
		PublicResourceViewOverride: dcl.StringOrNil(p.PublicResourceViewOverride),
		Project:                    dcl.StringOrNil(p.Project),
		Location:                   dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetHistory() {
		obj.History = append(obj.History, *ProtoToTier2AlphaInstanceHistory(r))
	}
	return obj
}

// InstanceSkuTierEnumToProto converts a InstanceSkuTierEnum enum to its proto representation.
func Tier2AlphaInstanceSkuTierEnumToProto(e *alpha.InstanceSkuTierEnum) alphapb.Tier2AlphaInstanceSkuTierEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceSkuTierEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceSkuTierEnum_value["InstanceSkuTierEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceSkuTierEnum(v)
	}
	return alphapb.Tier2AlphaInstanceSkuTierEnum(0)
}

// InstanceSkuSizeEnumToProto converts a InstanceSkuSizeEnum enum to its proto representation.
func Tier2AlphaInstanceSkuSizeEnumToProto(e *alpha.InstanceSkuSizeEnum) alphapb.Tier2AlphaInstanceSkuSizeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceSkuSizeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceSkuSizeEnum_value["InstanceSkuSizeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceSkuSizeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceSkuSizeEnum(0)
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func Tier2AlphaInstanceStateEnumToProto(e *alpha.InstanceStateEnum) alphapb.Tier2AlphaInstanceStateEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceStateEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceStateEnum(v)
	}
	return alphapb.Tier2AlphaInstanceStateEnum(0)
}

// InstancePreprocessCreateRecipeStepsActionEnumToProto converts a InstancePreprocessCreateRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessCreateRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum_value["InstancePreprocessCreateRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnum(0)
}

// InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceInitiateCreateRecipeStepsActionEnumToProto converts a InstanceInitiateCreateRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsActionEnumToProto(e *alpha.InstanceInitiateCreateRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsActionEnum_value["InstanceInitiateCreateRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsActionEnum(0)
}

// InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceCreateRecipeStepsActionEnumToProto converts a InstanceCreateRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsActionEnumToProto(e *alpha.InstanceCreateRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceCreateRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceCreateRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceCreateRecipeStepsActionEnum_value["InstanceCreateRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceCreateRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceCreateRecipeStepsActionEnum(0)
}

// InstanceCreateRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceDeleteRecipeStepsActionEnumToProto converts a InstanceDeleteRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsActionEnumToProto(e *alpha.InstanceDeleteRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceDeleteRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceDeleteRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceDeleteRecipeStepsActionEnum_value["InstanceDeleteRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceDeleteRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceDeleteRecipeStepsActionEnum(0)
}

// InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceUpdateRecipeStepsActionEnumToProto converts a InstanceUpdateRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsActionEnumToProto(e *alpha.InstanceUpdateRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceUpdateRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceUpdateRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceUpdateRecipeStepsActionEnum_value["InstanceUpdateRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceUpdateRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceUpdateRecipeStepsActionEnum(0)
}

// InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessResetRecipeStepsActionEnumToProto converts a InstancePreprocessResetRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessResetRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum_value["InstancePreprocessResetRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsActionEnum(0)
}

// InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceInitiateResetRecipeStepsActionEnumToProto converts a InstanceInitiateResetRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsActionEnumToProto(e *alpha.InstanceInitiateResetRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsActionEnum_value["InstanceInitiateResetRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsActionEnum(0)
}

// InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceResetRecipeStepsActionEnumToProto converts a InstanceResetRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsActionEnumToProto(e *alpha.InstanceResetRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceResetRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceResetRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceResetRecipeStepsActionEnum_value["InstanceResetRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceResetRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceResetRecipeStepsActionEnum(0)
}

// InstanceResetRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceResetRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceResetRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceResetRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessRepairRecipeStepsActionEnumToProto converts a InstancePreprocessRepairRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessRepairRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum_value["InstancePreprocessRepairRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnum(0)
}

// InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceInitiateRepairRecipeStepsActionEnumToProto converts a InstanceInitiateRepairRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsActionEnumToProto(e *alpha.InstanceInitiateRepairRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsActionEnum_value["InstanceInitiateRepairRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsActionEnum(0)
}

// InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceRepairRecipeStepsActionEnumToProto converts a InstanceRepairRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsActionEnumToProto(e *alpha.InstanceRepairRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceRepairRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceRepairRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceRepairRecipeStepsActionEnum_value["InstanceRepairRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceRepairRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceRepairRecipeStepsActionEnum(0)
}

// InstanceRepairRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessDeleteRecipeStepsActionEnumToProto converts a InstancePreprocessDeleteRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessDeleteRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum_value["InstancePreprocessDeleteRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnum(0)
}

// InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceInitiateDeleteRecipeStepsActionEnumToProto converts a InstanceInitiateDeleteRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsActionEnumToProto(e *alpha.InstanceInitiateDeleteRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsActionEnum_value["InstanceInitiateDeleteRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsActionEnum(0)
}

// InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessUpdateRecipeStepsActionEnumToProto converts a InstancePreprocessUpdateRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessUpdateRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum_value["InstancePreprocessUpdateRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnum(0)
}

// InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceInitiateUpdateRecipeStepsActionEnumToProto converts a InstanceInitiateUpdateRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsActionEnumToProto(e *alpha.InstanceInitiateUpdateRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsActionEnum_value["InstanceInitiateUpdateRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsActionEnum(0)
}

// InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessFreezeRecipeStepsActionEnumToProto converts a InstancePreprocessFreezeRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessFreezeRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum_value["InstancePreprocessFreezeRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnum(0)
}

// InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceFreezeRecipeStepsActionEnumToProto converts a InstanceFreezeRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsActionEnumToProto(e *alpha.InstanceFreezeRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceFreezeRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceFreezeRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceFreezeRecipeStepsActionEnum_value["InstanceFreezeRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceFreezeRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceFreezeRecipeStepsActionEnum(0)
}

// InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstancePreprocessUnfreezeRecipeStepsActionEnumToProto converts a InstancePreprocessUnfreezeRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnumToProto(e *alpha.InstancePreprocessUnfreezeRecipeStepsActionEnum) alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum_value["InstancePreprocessUnfreezeRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnum(0)
}

// InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum_value["InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceUnfreezeRecipeStepsActionEnumToProto converts a InstanceUnfreezeRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsActionEnumToProto(e *alpha.InstanceUnfreezeRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum_value["InstanceUnfreezeRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsActionEnum(0)
}

// InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceReadonlyRecipeStepsActionEnumToProto converts a InstanceReadonlyRecipeStepsActionEnum enum to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsActionEnumToProto(e *alpha.InstanceReadonlyRecipeStepsActionEnum) alphapb.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum_value["InstanceReadonlyRecipeStepsActionEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum(v)
	}
	return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsActionEnum(0)
}

// InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnumToProto converts a InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum enum to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnumToProto(e *alpha.InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum) alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum {
	if e == nil {
		return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(0)
	}
	if v, ok := alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum_value["InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum"+string(*e)]; ok {
		return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(v)
	}
	return alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(0)
}

// InstanceSkuToProto converts a InstanceSku resource to its proto representation.
func Tier2AlphaInstanceSkuToProto(o *alpha.InstanceSku) *alphapb.Tier2AlphaInstanceSku {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceSku{
		Tier: Tier2AlphaInstanceSkuTierEnumToProto(o.Tier),
		Size: Tier2AlphaInstanceSkuSizeEnumToProto(o.Size),
	}
	return p
}

// InstancePreprocessCreateRecipeToProto converts a InstancePreprocessCreateRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeToProto(o *alpha.InstancePreprocessCreateRecipe) *alphapb.Tier2AlphaInstancePreprocessCreateRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessCreateRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessCreateRecipeStepsToProto converts a InstancePreprocessCreateRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsToProto(o *alpha.InstancePreprocessCreateRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessCreateRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessCreateRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessCreateRecipeStepsStatusToProto converts a InstancePreprocessCreateRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsStatusToProto(o *alpha.InstancePreprocessCreateRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessCreateRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessCreateRecipeStepsStatusDetailsToProto converts a InstancePreprocessCreateRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessCreateRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessCreateRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessCreateRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessCreateRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessCreateRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessCreateRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessCreateRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsPermissionsInfoToProto converts a InstancePreprocessCreateRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrsToProto converts a InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstanceInitiateCreateRecipeToProto converts a InstanceInitiateCreateRecipe resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeToProto(o *alpha.InstanceInitiateCreateRecipe) *alphapb.Tier2AlphaInstanceInitiateCreateRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceInitiateCreateRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceInitiateCreateRecipeStepsToProto converts a InstanceInitiateCreateRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsToProto(o *alpha.InstanceInitiateCreateRecipeSteps) *alphapb.Tier2AlphaInstanceInitiateCreateRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceInitiateCreateRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceInitiateCreateRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceInitiateCreateRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceInitiateCreateRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceInitiateCreateRecipeStepsStatusToProto converts a InstanceInitiateCreateRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsStatusToProto(o *alpha.InstanceInitiateCreateRecipeStepsStatus) *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceInitiateCreateRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceInitiateCreateRecipeStepsStatusDetailsToProto converts a InstanceInitiateCreateRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsStatusDetailsToProto(o *alpha.InstanceInitiateCreateRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceInitiateCreateRecipeStepsQuotaRequestDeltasToProto converts a InstanceInitiateCreateRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceInitiateCreateRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstanceInitiateCreateRecipeStepsPreprocessUpdateToProto converts a InstanceInitiateCreateRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceInitiateCreateRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceInitiateCreateRecipeStepsRequestedTenantProjectToProto converts a InstanceInitiateCreateRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceInitiateCreateRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceInitiateCreateRecipeStepsPermissionsInfoToProto converts a InstanceInitiateCreateRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoToProto(o *alpha.InstanceInitiateCreateRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrsToProto converts a InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateToProto converts a InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstanceCreateRecipeToProto converts a InstanceCreateRecipe resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeToProto(o *alpha.InstanceCreateRecipe) *alphapb.Tier2AlphaInstanceCreateRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceCreateRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceCreateRecipeStepsToProto converts a InstanceCreateRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsToProto(o *alpha.InstanceCreateRecipeSteps) *alphapb.Tier2AlphaInstanceCreateRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceCreateRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceCreateRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceCreateRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceCreateRecipeStepsStatusToProto converts a InstanceCreateRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsStatusToProto(o *alpha.InstanceCreateRecipeStepsStatus) *alphapb.Tier2AlphaInstanceCreateRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceCreateRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceCreateRecipeStepsStatusDetailsToProto converts a InstanceCreateRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsStatusDetailsToProto(o *alpha.InstanceCreateRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceCreateRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceCreateRecipeStepsQuotaRequestDeltasToProto converts a InstanceCreateRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceCreateRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstanceCreateRecipeStepsPreprocessUpdateToProto converts a InstanceCreateRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceCreateRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceCreateRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceCreateRecipeStepsRequestedTenantProjectToProto converts a InstanceCreateRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceCreateRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceCreateRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceCreateRecipeStepsPermissionsInfoToProto converts a InstanceCreateRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoToProto(o *alpha.InstanceCreateRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceCreateRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceCreateRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceCreateRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceCreateRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceCreateRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceCreateRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceCreateRecipeStepsPermissionsInfoApiAttrsToProto converts a InstanceCreateRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstanceCreateRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstanceCreateRecipeStepsKeyNotificationsUpdateToProto converts a InstanceCreateRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceCreateRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstanceDeleteRecipeToProto converts a InstanceDeleteRecipe resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeToProto(o *alpha.InstanceDeleteRecipe) *alphapb.Tier2AlphaInstanceDeleteRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceDeleteRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceDeleteRecipeStepsToProto converts a InstanceDeleteRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsToProto(o *alpha.InstanceDeleteRecipeSteps) *alphapb.Tier2AlphaInstanceDeleteRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceDeleteRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceDeleteRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceDeleteRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceDeleteRecipeStepsStatusToProto converts a InstanceDeleteRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsStatusToProto(o *alpha.InstanceDeleteRecipeStepsStatus) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceDeleteRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceDeleteRecipeStepsStatusDetailsToProto converts a InstanceDeleteRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsStatusDetailsToProto(o *alpha.InstanceDeleteRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceDeleteRecipeStepsQuotaRequestDeltasToProto converts a InstanceDeleteRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceDeleteRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstanceDeleteRecipeStepsPreprocessUpdateToProto converts a InstanceDeleteRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceDeleteRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceDeleteRecipeStepsRequestedTenantProjectToProto converts a InstanceDeleteRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceDeleteRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceDeleteRecipeStepsPermissionsInfoToProto converts a InstanceDeleteRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoToProto(o *alpha.InstanceDeleteRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceDeleteRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceDeleteRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceDeleteRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceDeleteRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceDeleteRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceDeleteRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceDeleteRecipeStepsPermissionsInfoApiAttrsToProto converts a InstanceDeleteRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstanceDeleteRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstanceDeleteRecipeStepsKeyNotificationsUpdateToProto converts a InstanceDeleteRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstanceUpdateRecipeToProto converts a InstanceUpdateRecipe resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeToProto(o *alpha.InstanceUpdateRecipe) *alphapb.Tier2AlphaInstanceUpdateRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceUpdateRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceUpdateRecipeStepsToProto converts a InstanceUpdateRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsToProto(o *alpha.InstanceUpdateRecipeSteps) *alphapb.Tier2AlphaInstanceUpdateRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceUpdateRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceUpdateRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceUpdateRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceUpdateRecipeStepsStatusToProto converts a InstanceUpdateRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsStatusToProto(o *alpha.InstanceUpdateRecipeStepsStatus) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceUpdateRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceUpdateRecipeStepsStatusDetailsToProto converts a InstanceUpdateRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsStatusDetailsToProto(o *alpha.InstanceUpdateRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceUpdateRecipeStepsQuotaRequestDeltasToProto converts a InstanceUpdateRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceUpdateRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstanceUpdateRecipeStepsPreprocessUpdateToProto converts a InstanceUpdateRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceUpdateRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceUpdateRecipeStepsRequestedTenantProjectToProto converts a InstanceUpdateRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceUpdateRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceUpdateRecipeStepsPermissionsInfoToProto converts a InstanceUpdateRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoToProto(o *alpha.InstanceUpdateRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceUpdateRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceUpdateRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceUpdateRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceUpdateRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceUpdateRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceUpdateRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceUpdateRecipeStepsPermissionsInfoApiAttrsToProto converts a InstanceUpdateRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstanceUpdateRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstanceUpdateRecipeStepsKeyNotificationsUpdateToProto converts a InstanceUpdateRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstancePreprocessResetRecipeToProto converts a InstancePreprocessResetRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeToProto(o *alpha.InstancePreprocessResetRecipe) *alphapb.Tier2AlphaInstancePreprocessResetRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessResetRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessResetRecipeStepsToProto converts a InstancePreprocessResetRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsToProto(o *alpha.InstancePreprocessResetRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessResetRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessResetRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessResetRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessResetRecipeStepsStatusToProto converts a InstancePreprocessResetRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsStatusToProto(o *alpha.InstancePreprocessResetRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessResetRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessResetRecipeStepsStatusDetailsToProto converts a InstancePreprocessResetRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessResetRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessResetRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessResetRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessResetRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstancePreprocessResetRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessResetRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessResetRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessResetRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessResetRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessResetRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessResetRecipeStepsPermissionsInfoToProto converts a InstancePreprocessResetRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessResetRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrsToProto converts a InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstancePreprocessResetRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessResetRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstanceInitiateResetRecipeToProto converts a InstanceInitiateResetRecipe resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeToProto(o *alpha.InstanceInitiateResetRecipe) *alphapb.Tier2AlphaInstanceInitiateResetRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceInitiateResetRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceInitiateResetRecipeStepsToProto converts a InstanceInitiateResetRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsToProto(o *alpha.InstanceInitiateResetRecipeSteps) *alphapb.Tier2AlphaInstanceInitiateResetRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceInitiateResetRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceInitiateResetRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceInitiateResetRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceInitiateResetRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceInitiateResetRecipeStepsStatusToProto converts a InstanceInitiateResetRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsStatusToProto(o *alpha.InstanceInitiateResetRecipeStepsStatus) *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceInitiateResetRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceInitiateResetRecipeStepsStatusDetailsToProto converts a InstanceInitiateResetRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsStatusDetailsToProto(o *alpha.InstanceInitiateResetRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceInitiateResetRecipeStepsQuotaRequestDeltasToProto converts a InstanceInitiateResetRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceInitiateResetRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstanceInitiateResetRecipeStepsPreprocessUpdateToProto converts a InstanceInitiateResetRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceInitiateResetRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceInitiateResetRecipeStepsRequestedTenantProjectToProto converts a InstanceInitiateResetRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceInitiateResetRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceInitiateResetRecipeStepsPermissionsInfoToProto converts a InstanceInitiateResetRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoToProto(o *alpha.InstanceInitiateResetRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceInitiateResetRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrsToProto converts a InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstanceInitiateResetRecipeStepsKeyNotificationsUpdateToProto converts a InstanceInitiateResetRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceInitiateResetRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstanceResetRecipeToProto converts a InstanceResetRecipe resource to its proto representation.
func Tier2AlphaInstanceResetRecipeToProto(o *alpha.InstanceResetRecipe) *alphapb.Tier2AlphaInstanceResetRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceResetRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceResetRecipeStepsToProto converts a InstanceResetRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsToProto(o *alpha.InstanceResetRecipeSteps) *alphapb.Tier2AlphaInstanceResetRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceResetRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceResetRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceResetRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceResetRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceResetRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceResetRecipeStepsStatusToProto converts a InstanceResetRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsStatusToProto(o *alpha.InstanceResetRecipeStepsStatus) *alphapb.Tier2AlphaInstanceResetRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceResetRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceResetRecipeStepsStatusDetailsToProto converts a InstanceResetRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsStatusDetailsToProto(o *alpha.InstanceResetRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceResetRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceResetRecipeStepsQuotaRequestDeltasToProto converts a InstanceResetRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceResetRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceResetRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstanceResetRecipeStepsPreprocessUpdateToProto converts a InstanceResetRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceResetRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceResetRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceResetRecipeStepsRequestedTenantProjectToProto converts a InstanceResetRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceResetRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceResetRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceResetRecipeStepsPermissionsInfoToProto converts a InstanceResetRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsPermissionsInfoToProto(o *alpha.InstanceResetRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstanceResetRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceResetRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceResetRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceResetRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceResetRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceResetRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceResetRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceResetRecipeStepsPermissionsInfoApiAttrsToProto converts a InstanceResetRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstanceResetRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstanceResetRecipeStepsKeyNotificationsUpdateToProto converts a InstanceResetRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceResetRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstancePreprocessRepairRecipeToProto converts a InstancePreprocessRepairRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeToProto(o *alpha.InstancePreprocessRepairRecipe) *alphapb.Tier2AlphaInstancePreprocessRepairRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessRepairRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessRepairRecipeStepsToProto converts a InstancePreprocessRepairRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsToProto(o *alpha.InstancePreprocessRepairRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessRepairRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessRepairRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessRepairRecipeStepsStatusToProto converts a InstancePreprocessRepairRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsStatusToProto(o *alpha.InstancePreprocessRepairRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessRepairRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessRepairRecipeStepsStatusDetailsToProto converts a InstancePreprocessRepairRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessRepairRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessRepairRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessRepairRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessRepairRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessRepairRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessRepairRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessRepairRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsPermissionsInfoToProto converts a InstancePreprocessRepairRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrsToProto converts a InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstanceInitiateRepairRecipeToProto converts a InstanceInitiateRepairRecipe resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeToProto(o *alpha.InstanceInitiateRepairRecipe) *alphapb.Tier2AlphaInstanceInitiateRepairRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceInitiateRepairRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceInitiateRepairRecipeStepsToProto converts a InstanceInitiateRepairRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsToProto(o *alpha.InstanceInitiateRepairRecipeSteps) *alphapb.Tier2AlphaInstanceInitiateRepairRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceInitiateRepairRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceInitiateRepairRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceInitiateRepairRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceInitiateRepairRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceInitiateRepairRecipeStepsStatusToProto converts a InstanceInitiateRepairRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsStatusToProto(o *alpha.InstanceInitiateRepairRecipeStepsStatus) *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceInitiateRepairRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceInitiateRepairRecipeStepsStatusDetailsToProto converts a InstanceInitiateRepairRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsStatusDetailsToProto(o *alpha.InstanceInitiateRepairRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceInitiateRepairRecipeStepsQuotaRequestDeltasToProto converts a InstanceInitiateRepairRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceInitiateRepairRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstanceInitiateRepairRecipeStepsPreprocessUpdateToProto converts a InstanceInitiateRepairRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceInitiateRepairRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceInitiateRepairRecipeStepsRequestedTenantProjectToProto converts a InstanceInitiateRepairRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceInitiateRepairRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceInitiateRepairRecipeStepsPermissionsInfoToProto converts a InstanceInitiateRepairRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoToProto(o *alpha.InstanceInitiateRepairRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrsToProto converts a InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateToProto converts a InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstanceRepairRecipeToProto converts a InstanceRepairRecipe resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeToProto(o *alpha.InstanceRepairRecipe) *alphapb.Tier2AlphaInstanceRepairRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceRepairRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceRepairRecipeStepsToProto converts a InstanceRepairRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsToProto(o *alpha.InstanceRepairRecipeSteps) *alphapb.Tier2AlphaInstanceRepairRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceRepairRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceRepairRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceRepairRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceRepairRecipeStepsStatusToProto converts a InstanceRepairRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsStatusToProto(o *alpha.InstanceRepairRecipeStepsStatus) *alphapb.Tier2AlphaInstanceRepairRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceRepairRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceRepairRecipeStepsStatusDetailsToProto converts a InstanceRepairRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsStatusDetailsToProto(o *alpha.InstanceRepairRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceRepairRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceRepairRecipeStepsQuotaRequestDeltasToProto converts a InstanceRepairRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceRepairRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstanceRepairRecipeStepsPreprocessUpdateToProto converts a InstanceRepairRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceRepairRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceRepairRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceRepairRecipeStepsRequestedTenantProjectToProto converts a InstanceRepairRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceRepairRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceRepairRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceRepairRecipeStepsPermissionsInfoToProto converts a InstanceRepairRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoToProto(o *alpha.InstanceRepairRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceRepairRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceRepairRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceRepairRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceRepairRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceRepairRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceRepairRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceRepairRecipeStepsPermissionsInfoApiAttrsToProto converts a InstanceRepairRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstanceRepairRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstanceRepairRecipeStepsKeyNotificationsUpdateToProto converts a InstanceRepairRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceRepairRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstancePreprocessDeleteRecipeToProto converts a InstancePreprocessDeleteRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeToProto(o *alpha.InstancePreprocessDeleteRecipe) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessDeleteRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsToProto converts a InstancePreprocessDeleteRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsToProto(o *alpha.InstancePreprocessDeleteRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessDeleteRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsStatusToProto converts a InstancePreprocessDeleteRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusToProto(o *alpha.InstancePreprocessDeleteRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsStatusDetailsToProto converts a InstancePreprocessDeleteRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessDeleteRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessDeleteRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessDeleteRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessDeleteRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessDeleteRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsPermissionsInfoToProto converts a InstancePreprocessDeleteRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrsToProto converts a InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstanceInitiateDeleteRecipeToProto converts a InstanceInitiateDeleteRecipe resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeToProto(o *alpha.InstanceInitiateDeleteRecipe) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceInitiateDeleteRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceInitiateDeleteRecipeStepsToProto converts a InstanceInitiateDeleteRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsToProto(o *alpha.InstanceInitiateDeleteRecipeSteps) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceInitiateDeleteRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceInitiateDeleteRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceInitiateDeleteRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceInitiateDeleteRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceInitiateDeleteRecipeStepsStatusToProto converts a InstanceInitiateDeleteRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsStatusToProto(o *alpha.InstanceInitiateDeleteRecipeStepsStatus) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceInitiateDeleteRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceInitiateDeleteRecipeStepsStatusDetailsToProto converts a InstanceInitiateDeleteRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsStatusDetailsToProto(o *alpha.InstanceInitiateDeleteRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceInitiateDeleteRecipeStepsQuotaRequestDeltasToProto converts a InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstanceInitiateDeleteRecipeStepsPreprocessUpdateToProto converts a InstanceInitiateDeleteRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceInitiateDeleteRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceInitiateDeleteRecipeStepsRequestedTenantProjectToProto converts a InstanceInitiateDeleteRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceInitiateDeleteRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceInitiateDeleteRecipeStepsPermissionsInfoToProto converts a InstanceInitiateDeleteRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoToProto(o *alpha.InstanceInitiateDeleteRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrsToProto converts a InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateToProto converts a InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstancePreprocessUpdateRecipeToProto converts a InstancePreprocessUpdateRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeToProto(o *alpha.InstancePreprocessUpdateRecipe) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessUpdateRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsToProto converts a InstancePreprocessUpdateRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsToProto(o *alpha.InstancePreprocessUpdateRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessUpdateRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsStatusToProto converts a InstancePreprocessUpdateRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusToProto(o *alpha.InstancePreprocessUpdateRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsStatusDetailsToProto converts a InstancePreprocessUpdateRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessUpdateRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessUpdateRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessUpdateRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessUpdateRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessUpdateRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsPermissionsInfoToProto converts a InstancePreprocessUpdateRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrsToProto converts a InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstanceInitiateUpdateRecipeToProto converts a InstanceInitiateUpdateRecipe resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeToProto(o *alpha.InstanceInitiateUpdateRecipe) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceInitiateUpdateRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceInitiateUpdateRecipeStepsToProto converts a InstanceInitiateUpdateRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsToProto(o *alpha.InstanceInitiateUpdateRecipeSteps) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceInitiateUpdateRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceInitiateUpdateRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceInitiateUpdateRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceInitiateUpdateRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceInitiateUpdateRecipeStepsStatusToProto converts a InstanceInitiateUpdateRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsStatusToProto(o *alpha.InstanceInitiateUpdateRecipeStepsStatus) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceInitiateUpdateRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceInitiateUpdateRecipeStepsStatusDetailsToProto converts a InstanceInitiateUpdateRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsStatusDetailsToProto(o *alpha.InstanceInitiateUpdateRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceInitiateUpdateRecipeStepsQuotaRequestDeltasToProto converts a InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstanceInitiateUpdateRecipeStepsPreprocessUpdateToProto converts a InstanceInitiateUpdateRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceInitiateUpdateRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceInitiateUpdateRecipeStepsRequestedTenantProjectToProto converts a InstanceInitiateUpdateRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceInitiateUpdateRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceInitiateUpdateRecipeStepsPermissionsInfoToProto converts a InstanceInitiateUpdateRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoToProto(o *alpha.InstanceInitiateUpdateRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrsToProto converts a InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateToProto converts a InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstancePreprocessFreezeRecipeToProto converts a InstancePreprocessFreezeRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeToProto(o *alpha.InstancePreprocessFreezeRecipe) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessFreezeRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsToProto converts a InstancePreprocessFreezeRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsToProto(o *alpha.InstancePreprocessFreezeRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessFreezeRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsStatusToProto converts a InstancePreprocessFreezeRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusToProto(o *alpha.InstancePreprocessFreezeRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsStatusDetailsToProto converts a InstancePreprocessFreezeRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessFreezeRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessFreezeRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessFreezeRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessFreezeRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessFreezeRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsPermissionsInfoToProto converts a InstancePreprocessFreezeRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrsToProto converts a InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstanceFreezeRecipeToProto converts a InstanceFreezeRecipe resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeToProto(o *alpha.InstanceFreezeRecipe) *alphapb.Tier2AlphaInstanceFreezeRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceFreezeRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceFreezeRecipeStepsToProto converts a InstanceFreezeRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsToProto(o *alpha.InstanceFreezeRecipeSteps) *alphapb.Tier2AlphaInstanceFreezeRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceFreezeRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceFreezeRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceFreezeRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceFreezeRecipeStepsStatusToProto converts a InstanceFreezeRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsStatusToProto(o *alpha.InstanceFreezeRecipeStepsStatus) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceFreezeRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceFreezeRecipeStepsStatusDetailsToProto converts a InstanceFreezeRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsStatusDetailsToProto(o *alpha.InstanceFreezeRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceFreezeRecipeStepsQuotaRequestDeltasToProto converts a InstanceFreezeRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceFreezeRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstanceFreezeRecipeStepsPreprocessUpdateToProto converts a InstanceFreezeRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceFreezeRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceFreezeRecipeStepsRequestedTenantProjectToProto converts a InstanceFreezeRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceFreezeRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceFreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceFreezeRecipeStepsPermissionsInfoToProto converts a InstanceFreezeRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoToProto(o *alpha.InstanceFreezeRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceFreezeRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceFreezeRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceFreezeRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceFreezeRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceFreezeRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceFreezeRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceFreezeRecipeStepsPermissionsInfoApiAttrsToProto converts a InstanceFreezeRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstanceFreezeRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstanceFreezeRecipeStepsKeyNotificationsUpdateToProto converts a InstanceFreezeRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeToProto converts a InstancePreprocessUnfreezeRecipe resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeToProto(o *alpha.InstancePreprocessUnfreezeRecipe) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstancePreprocessUnfreezeRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsToProto converts a InstancePreprocessUnfreezeRecipeSteps resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsToProto(o *alpha.InstancePreprocessUnfreezeRecipeSteps) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstancePreprocessUnfreezeRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsStatusToProto converts a InstancePreprocessUnfreezeRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsStatus) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsStatusDetailsToProto converts a InstancePreprocessUnfreezeRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetailsToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltasToProto converts a InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsPreprocessUpdateToProto converts a InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdateToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectToProto converts a InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsPermissionsInfoToProto converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameToProto converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrsToProto converts a InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateToProto converts a InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstanceUnfreezeRecipeToProto converts a InstanceUnfreezeRecipe resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeToProto(o *alpha.InstanceUnfreezeRecipe) *alphapb.Tier2AlphaInstanceUnfreezeRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceUnfreezeRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceUnfreezeRecipeStepsToProto converts a InstanceUnfreezeRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsToProto(o *alpha.InstanceUnfreezeRecipeSteps) *alphapb.Tier2AlphaInstanceUnfreezeRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceUnfreezeRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceUnfreezeRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceUnfreezeRecipeStepsStatusToProto converts a InstanceUnfreezeRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsStatusToProto(o *alpha.InstanceUnfreezeRecipeStepsStatus) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceUnfreezeRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceUnfreezeRecipeStepsStatusDetailsToProto converts a InstanceUnfreezeRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsStatusDetailsToProto(o *alpha.InstanceUnfreezeRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceUnfreezeRecipeStepsQuotaRequestDeltasToProto converts a InstanceUnfreezeRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceUnfreezeRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstanceUnfreezeRecipeStepsPreprocessUpdateToProto converts a InstanceUnfreezeRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceUnfreezeRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceUnfreezeRecipeStepsRequestedTenantProjectToProto converts a InstanceUnfreezeRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceUnfreezeRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceUnfreezeRecipeStepsPermissionsInfoToProto converts a InstanceUnfreezeRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoToProto(o *alpha.InstanceUnfreezeRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrsToProto converts a InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstanceUnfreezeRecipeStepsKeyNotificationsUpdateToProto converts a InstanceUnfreezeRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstanceReadonlyRecipeToProto converts a InstanceReadonlyRecipe resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeToProto(o *alpha.InstanceReadonlyRecipe) *alphapb.Tier2AlphaInstanceReadonlyRecipe {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipe{
		HonorCancelRequest:                dcl.ValueOrEmptyBool(o.HonorCancelRequest),
		IgnoreRecipeAfter:                 dcl.ValueOrEmptyInt64(o.IgnoreRecipeAfter),
		VerifyDeadlineSecondsBelow:        dcl.ValueOrEmptyDouble(o.VerifyDeadlineSecondsBelow),
		PopulateOperationResult:           dcl.ValueOrEmptyBool(o.PopulateOperationResult),
		ReadonlyRecipeStartTime:           dcl.ValueOrEmptyString(o.ReadonlyRecipeStartTime),
		DelayToStoreResourcesInClhDbNanos: dcl.ValueOrEmptyInt64(o.DelayToStoreResourcesInClhDbNanos),
	}
	for _, r := range o.Steps {
		p.Steps = append(p.Steps, Tier2AlphaInstanceReadonlyRecipeStepsToProto(&r))
	}
	for _, r := range o.ResourceNamesStoredInClhWithDelay {
		p.ResourceNamesStoredInClhWithDelay = append(p.ResourceNamesStoredInClhWithDelay, r)
	}
	return p
}

// InstanceReadonlyRecipeStepsToProto converts a InstanceReadonlyRecipeSteps resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsToProto(o *alpha.InstanceReadonlyRecipeSteps) *alphapb.Tier2AlphaInstanceReadonlyRecipeSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeSteps{
		RelativeTime:                   dcl.ValueOrEmptyInt64(o.RelativeTime),
		SleepDuration:                  dcl.ValueOrEmptyInt64(o.SleepDuration),
		Action:                         Tier2AlphaInstanceReadonlyRecipeStepsActionEnumToProto(o.Action),
		Status:                         Tier2AlphaInstanceReadonlyRecipeStepsStatusToProto(o.Status),
		ErrorSpace:                     dcl.ValueOrEmptyString(o.ErrorSpace),
		P4ServiceAccount:               dcl.ValueOrEmptyString(o.P4ServiceAccount),
		ResourceMetadataSize:           dcl.ValueOrEmptyInt64(o.ResourceMetadataSize),
		Description:                    dcl.ValueOrEmptyString(o.Description),
		UpdatedRepeatOperationDelaySec: dcl.ValueOrEmptyDouble(o.UpdatedRepeatOperationDelaySec),
		PreprocessUpdate:               Tier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdateToProto(o.PreprocessUpdate),
		PublicOperationMetadata:        dcl.ValueOrEmptyString(o.PublicOperationMetadata),
		RequestedTenantProject:         Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectToProto(o.RequestedTenantProject),
		KeyNotificationsUpdate:         Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateToProto(o.KeyNotificationsUpdate),
		ClhDataUpdateTime:              dcl.ValueOrEmptyString(o.ClhDataUpdateTime),
	}
	for _, r := range o.QuotaRequestDeltas {
		p.QuotaRequestDeltas = append(p.QuotaRequestDeltas, Tier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltasToProto(&r))
	}
	for _, r := range o.PermissionsInfo {
		p.PermissionsInfo = append(p.PermissionsInfo, Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoToProto(&r))
	}
	return p
}

// InstanceReadonlyRecipeStepsStatusToProto converts a InstanceReadonlyRecipeStepsStatus resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsStatusToProto(o *alpha.InstanceReadonlyRecipeStepsStatus) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsStatus{
		Code:    dcl.ValueOrEmptyInt64(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Details {
		p.Details = append(p.Details, Tier2AlphaInstanceReadonlyRecipeStepsStatusDetailsToProto(&r))
	}
	return p
}

// InstanceReadonlyRecipeStepsStatusDetailsToProto converts a InstanceReadonlyRecipeStepsStatusDetails resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsStatusDetailsToProto(o *alpha.InstanceReadonlyRecipeStepsStatusDetails) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsStatusDetails {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsStatusDetails{
		TypeUrl: dcl.ValueOrEmptyString(o.TypeUrl),
		Value:   dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// InstanceReadonlyRecipeStepsQuotaRequestDeltasToProto converts a InstanceReadonlyRecipeStepsQuotaRequestDeltas resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltasToProto(o *alpha.InstanceReadonlyRecipeStepsQuotaRequestDeltas) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltas {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsQuotaRequestDeltas{
		MetricName: dcl.ValueOrEmptyString(o.MetricName),
		Amount:     dcl.ValueOrEmptyInt64(o.Amount),
	}
	return p
}

// InstanceReadonlyRecipeStepsPreprocessUpdateToProto converts a InstanceReadonlyRecipeStepsPreprocessUpdate resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdateToProto(o *alpha.InstanceReadonlyRecipeStepsPreprocessUpdate) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPreprocessUpdate{
		LatencySloBucketName:    dcl.ValueOrEmptyString(o.LatencySloBucketName),
		PublicOperationMetadata: dcl.ValueOrEmptyString(o.PublicOperationMetadata),
	}
	return p
}

// InstanceReadonlyRecipeStepsRequestedTenantProjectToProto converts a InstanceReadonlyRecipeStepsRequestedTenantProject resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectToProto(o *alpha.InstanceReadonlyRecipeStepsRequestedTenantProject) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProject {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProject{
		Tag:    dcl.ValueOrEmptyString(o.Tag),
		Folder: dcl.ValueOrEmptyString(o.Folder),
		Scope:  Tier2AlphaInstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnumToProto(o.Scope),
	}
	return p
}

// InstanceReadonlyRecipeStepsPermissionsInfoToProto converts a InstanceReadonlyRecipeStepsPermissionsInfo resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoToProto(o *alpha.InstanceReadonlyRecipeStepsPermissionsInfo) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfo{
		PolicyName:   Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameToProto(o.PolicyName),
		ResourcePath: dcl.ValueOrEmptyString(o.ResourcePath),
		ApiAttrs:     Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoApiAttrsToProto(o.ApiAttrs),
	}
	for _, r := range o.IamPermissions {
		p.IamPermissions = append(p.IamPermissions, Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissionsToProto(&r))
	}
	return p
}

// InstanceReadonlyRecipeStepsPermissionsInfoPolicyNameToProto converts a InstanceReadonlyRecipeStepsPermissionsInfoPolicyName resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyNameToProto(o *alpha.InstanceReadonlyRecipeStepsPermissionsInfoPolicyName) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyName {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoPolicyName{
		Type:   dcl.ValueOrEmptyString(o.Type),
		Id:     dcl.ValueOrEmptyString(o.Id),
		Region: dcl.ValueOrEmptyString(o.Region),
	}
	return p
}

// InstanceReadonlyRecipeStepsPermissionsInfoIamPermissionsToProto converts a InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissionsToProto(o *alpha.InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions{
		Permission: dcl.ValueOrEmptyString(o.Permission),
	}
	return p
}

// InstanceReadonlyRecipeStepsPermissionsInfoApiAttrsToProto converts a InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoApiAttrsToProto(o *alpha.InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoApiAttrs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsPermissionsInfoApiAttrs{}
	return p
}

// InstanceReadonlyRecipeStepsKeyNotificationsUpdateToProto converts a InstanceReadonlyRecipeStepsKeyNotificationsUpdate resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateToProto(o *alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdate) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdate {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdate{
		KeyNotificationsInfo: Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o.KeyNotificationsInfo),
	}
	return p
}

// InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto converts a InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoToProto(o *alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{
		KeyConfigs:  Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o.KeyConfigs),
		DataVersion: dcl.ValueOrEmptyInt64(o.DataVersion),
		Delegate:    dcl.ValueOrEmptyString(o.Delegate),
	}
	return p
}

// InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto converts a InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsToProto(o *alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{
		KeyConfig: Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o.KeyConfig),
	}
	return p
}

// InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto converts a InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig resource to its proto representation.
func Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfigToProto(o *alpha.InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) *alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{
		KeyOrVersionName: dcl.ValueOrEmptyString(o.KeyOrVersionName),
	}
	return p
}

// InstanceHistoryToProto converts a InstanceHistory resource to its proto representation.
func Tier2AlphaInstanceHistoryToProto(o *alpha.InstanceHistory) *alphapb.Tier2AlphaInstanceHistory {
	if o == nil {
		return nil
	}
	p := &alphapb.Tier2AlphaInstanceHistory{
		Timestamp:           dcl.ValueOrEmptyString(o.Timestamp),
		OperationHandle:     dcl.ValueOrEmptyString(o.OperationHandle),
		Description:         dcl.ValueOrEmptyString(o.Description),
		StepIndex:           dcl.ValueOrEmptyInt64(o.StepIndex),
		TenantProjectNumber: dcl.ValueOrEmptyInt64(o.TenantProjectNumber),
		TenantProjectId:     dcl.ValueOrEmptyString(o.TenantProjectId),
		P4ServiceAccount:    dcl.ValueOrEmptyString(o.P4ServiceAccount),
	}
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *alpha.Instance) *alphapb.Tier2AlphaInstance {
	p := &alphapb.Tier2AlphaInstance{
		Name:                       dcl.ValueOrEmptyString(resource.Name),
		DisplayName:                dcl.ValueOrEmptyString(resource.DisplayName),
		Zone:                       dcl.ValueOrEmptyString(resource.Zone),
		Sku:                        Tier2AlphaInstanceSkuToProto(resource.Sku),
		AuthorizedNetworkId:        dcl.ValueOrEmptyString(resource.AuthorizedNetworkId),
		ReservedIpRange:            dcl.ValueOrEmptyString(resource.ReservedIPRange),
		HostName:                   dcl.ValueOrEmptyString(resource.HostName),
		PortNumber:                 dcl.ValueOrEmptyInt64(resource.PortNumber),
		CurrentZone:                dcl.ValueOrEmptyString(resource.CurrentZone),
		CreationTime:               dcl.ValueOrEmptyString(resource.CreationTime),
		State:                      Tier2AlphaInstanceStateEnumToProto(resource.State),
		StatusMessage:              dcl.ValueOrEmptyString(resource.StatusMessage),
		ExtraField:                 dcl.ValueOrEmptyString(resource.ExtraField),
		PreprocessCreateRecipe:     Tier2AlphaInstancePreprocessCreateRecipeToProto(resource.PreprocessCreateRecipe),
		InitiateCreateRecipe:       Tier2AlphaInstanceInitiateCreateRecipeToProto(resource.InitiateCreateRecipe),
		CreateRecipe:               Tier2AlphaInstanceCreateRecipeToProto(resource.CreateRecipe),
		DeleteRecipe:               Tier2AlphaInstanceDeleteRecipeToProto(resource.DeleteRecipe),
		UpdateRecipe:               Tier2AlphaInstanceUpdateRecipeToProto(resource.UpdateRecipe),
		PreprocessResetRecipe:      Tier2AlphaInstancePreprocessResetRecipeToProto(resource.PreprocessResetRecipe),
		InitiateResetRecipe:        Tier2AlphaInstanceInitiateResetRecipeToProto(resource.InitiateResetRecipe),
		ResetRecipe:                Tier2AlphaInstanceResetRecipeToProto(resource.ResetRecipe),
		PreprocessRepairRecipe:     Tier2AlphaInstancePreprocessRepairRecipeToProto(resource.PreprocessRepairRecipe),
		InitiateRepairRecipe:       Tier2AlphaInstanceInitiateRepairRecipeToProto(resource.InitiateRepairRecipe),
		RepairRecipe:               Tier2AlphaInstanceRepairRecipeToProto(resource.RepairRecipe),
		PreprocessDeleteRecipe:     Tier2AlphaInstancePreprocessDeleteRecipeToProto(resource.PreprocessDeleteRecipe),
		InitiateDeleteRecipe:       Tier2AlphaInstanceInitiateDeleteRecipeToProto(resource.InitiateDeleteRecipe),
		PreprocessUpdateRecipe:     Tier2AlphaInstancePreprocessUpdateRecipeToProto(resource.PreprocessUpdateRecipe),
		InitiateUpdateRecipe:       Tier2AlphaInstanceInitiateUpdateRecipeToProto(resource.InitiateUpdateRecipe),
		PreprocessFreezeRecipe:     Tier2AlphaInstancePreprocessFreezeRecipeToProto(resource.PreprocessFreezeRecipe),
		FreezeRecipe:               Tier2AlphaInstanceFreezeRecipeToProto(resource.FreezeRecipe),
		PreprocessUnfreezeRecipe:   Tier2AlphaInstancePreprocessUnfreezeRecipeToProto(resource.PreprocessUnfreezeRecipe),
		UnfreezeRecipe:             Tier2AlphaInstanceUnfreezeRecipeToProto(resource.UnfreezeRecipe),
		ReadonlyRecipe:             Tier2AlphaInstanceReadonlyRecipeToProto(resource.ReadonlyRecipe),
		EnableCallHistory:          dcl.ValueOrEmptyBool(resource.EnableCallHistory),
		PublicResourceViewOverride: dcl.ValueOrEmptyString(resource.PublicResourceViewOverride),
		Project:                    dcl.ValueOrEmptyString(resource.Project),
		Location:                   dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.History {
		p.History = append(p.History, Tier2AlphaInstanceHistoryToProto(&r))
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *alpha.Client, request *alphapb.ApplyTier2AlphaInstanceRequest) (*alphapb.Tier2AlphaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyTier2AlphaInstance(ctx context.Context, request *alphapb.ApplyTier2AlphaInstanceRequest) (*alphapb.Tier2AlphaInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteTier2AlphaInstance(ctx context.Context, request *alphapb.DeleteTier2AlphaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListTier2AlphaInstance(ctx context.Context, request *alphapb.ListTier2AlphaInstanceRequest) (*alphapb.ListTier2AlphaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.Tier2AlphaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListTier2AlphaInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
