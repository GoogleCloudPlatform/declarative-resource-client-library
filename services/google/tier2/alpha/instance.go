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
package alpha

import (
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Instance struct {
	Name                                 *string                                       `json:"name"`
	DisplayName                          *string                                       `json:"displayName"`
	Labels                               map[string]string                             `json:"labels"`
	Zone                                 *string                                       `json:"zone"`
	AlternativeZone                      *string                                       `json:"alternativeZone"`
	Sku                                  *InstanceSku                                  `json:"sku"`
	AuthorizedNetworkId                  *string                                       `json:"authorizedNetworkId"`
	ReservedIPRange                      *string                                       `json:"reservedIPRange"`
	Host                                 *string                                       `json:"host"`
	Port                                 *int64                                        `json:"port"`
	CurrentZone                          *string                                       `json:"currentZone"`
	CreateTime                           *string                                       `json:"createTime"`
	State                                *InstanceStateEnum                            `json:"state"`
	StatusMessage                        *string                                       `json:"statusMessage"`
	UpdateTime                           *string                                       `json:"updateTime"`
	MutateUserId                         *int64                                        `json:"mutateUserId"`
	ReadUserId                           *int64                                        `json:"readUserId"`
	References                           []InstanceReferences                          `json:"references"`
	PreprocessCreateRecipe               *InstancePreprocessCreateRecipe               `json:"preprocessCreateRecipe"`
	CreateRecipe                         *InstanceCreateRecipe                         `json:"createRecipe"`
	DeleteRecipe                         *InstanceDeleteRecipe                         `json:"deleteRecipe"`
	UpdateRecipe                         *InstanceUpdateRecipe                         `json:"updateRecipe"`
	PreprocessResetRecipe                *InstancePreprocessResetRecipe                `json:"preprocessResetRecipe"`
	ResetRecipe                          *InstanceResetRecipe                          `json:"resetRecipe"`
	PreprocessRepairRecipe               *InstancePreprocessRepairRecipe               `json:"preprocessRepairRecipe"`
	RepairRecipe                         *InstanceRepairRecipe                         `json:"repairRecipe"`
	PreprocessDeleteRecipe               *InstancePreprocessDeleteRecipe               `json:"preprocessDeleteRecipe"`
	PreprocessUpdateRecipe               *InstancePreprocessUpdateRecipe               `json:"preprocessUpdateRecipe"`
	PreprocessFreezeRecipe               *InstancePreprocessFreezeRecipe               `json:"preprocessFreezeRecipe"`
	FreezeRecipe                         *InstanceFreezeRecipe                         `json:"freezeRecipe"`
	PreprocessUnfreezeRecipe             *InstancePreprocessUnfreezeRecipe             `json:"preprocessUnfreezeRecipe"`
	UnfreezeRecipe                       *InstanceUnfreezeRecipe                       `json:"unfreezeRecipe"`
	PreprocessReportInstanceHealthRecipe *InstancePreprocessReportInstanceHealthRecipe `json:"preprocessReportInstanceHealthRecipe"`
	ReportInstanceHealthRecipe           *InstanceReportInstanceHealthRecipe           `json:"reportInstanceHealthRecipe"`
	PreprocessGetRecipe                  *InstancePreprocessGetRecipe                  `json:"preprocessGetRecipe"`
	NotifyKeyAvailableRecipe             *InstanceNotifyKeyAvailableRecipe             `json:"notifyKeyAvailableRecipe"`
	NotifyKeyUnavailableRecipe           *InstanceNotifyKeyUnavailableRecipe           `json:"notifyKeyUnavailableRecipe"`
	ReadonlyRecipe                       *InstanceReadonlyRecipe                       `json:"readonlyRecipe"`
	ReconcileRecipe                      *InstanceReconcileRecipe                      `json:"reconcileRecipe"`
	PreprocessPassthroughRecipe          *InstancePreprocessPassthroughRecipe          `json:"preprocessPassthroughRecipe"`
	PreprocessReconcileRecipe            *InstancePreprocessReconcileRecipe            `json:"preprocessReconcileRecipe"`
	EnableCallHistory                    *bool                                         `json:"enableCallHistory"`
	History                              []InstanceHistory                             `json:"history"`
	PublicResourceViewOverride           *string                                       `json:"publicResourceViewOverride"`
	ExtraInfo                            *string                                       `json:"extraInfo"`
	Uid                                  *string                                       `json:"uid"`
	Etag                                 *string                                       `json:"etag"`
	Project                              *string                                       `json:"project"`
	Location                             *string                                       `json:"location"`
}

func (r *Instance) String() string {
	return dcl.SprintResource(r)
}

// The enum InstanceSkuTierEnum.
type InstanceSkuTierEnum string

// InstanceSkuTierEnumRef returns a *InstanceSkuTierEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceSkuTierEnumRef(s string) *InstanceSkuTierEnum {
	if s == "" {
		return nil
	}

	v := InstanceSkuTierEnum(s)
	return &v
}

func (v InstanceSkuTierEnum) Validate() error {
	for _, s := range []string{"TIER_UNSPECIFIED", "STANDALONE", "REPLICATED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceSkuTierEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceSkuSizeEnum.
type InstanceSkuSizeEnum string

// InstanceSkuSizeEnumRef returns a *InstanceSkuSizeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceSkuSizeEnumRef(s string) *InstanceSkuSizeEnum {
	if s == "" {
		return nil
	}

	v := InstanceSkuSizeEnum(s)
	return &v
}

func (v InstanceSkuSizeEnum) Validate() error {
	for _, s := range []string{"SIZE_UNSPECIFIED", "C1", "C2"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceSkuSizeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceStateEnum.
type InstanceStateEnum string

// InstanceStateEnumRef returns a *InstanceStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceStateEnumRef(s string) *InstanceStateEnum {
	if s == "" {
		return nil
	}

	v := InstanceStateEnum(s)
	return &v
}

func (v InstanceStateEnum) Validate() error {
	for _, s := range []string{"STATUS_UNSPECIFIED", "CREATING", "READY", "UPDATING", "DELETING", "REPAIRING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessCreateRecipeStepsActionEnum.
type InstancePreprocessCreateRecipeStepsActionEnum string

// InstancePreprocessCreateRecipeStepsActionEnumRef returns a *InstancePreprocessCreateRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessCreateRecipeStepsActionEnumRef(s string) *InstancePreprocessCreateRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessCreateRecipeStepsActionEnum(s)
	return &v
}

func (v InstancePreprocessCreateRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessCreateRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum.
type InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum string

// InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceCreateRecipeStepsActionEnum.
type InstanceCreateRecipeStepsActionEnum string

// InstanceCreateRecipeStepsActionEnumRef returns a *InstanceCreateRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceCreateRecipeStepsActionEnumRef(s string) *InstanceCreateRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceCreateRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceCreateRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceCreateRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceCreateRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceCreateRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceDeleteRecipeStepsActionEnum.
type InstanceDeleteRecipeStepsActionEnum string

// InstanceDeleteRecipeStepsActionEnumRef returns a *InstanceDeleteRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceDeleteRecipeStepsActionEnumRef(s string) *InstanceDeleteRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceDeleteRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceDeleteRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceDeleteRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceUpdateRecipeStepsActionEnum.
type InstanceUpdateRecipeStepsActionEnum string

// InstanceUpdateRecipeStepsActionEnumRef returns a *InstanceUpdateRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceUpdateRecipeStepsActionEnumRef(s string) *InstanceUpdateRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceUpdateRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceUpdateRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceUpdateRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessResetRecipeStepsActionEnum.
type InstancePreprocessResetRecipeStepsActionEnum string

// InstancePreprocessResetRecipeStepsActionEnumRef returns a *InstancePreprocessResetRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessResetRecipeStepsActionEnumRef(s string) *InstancePreprocessResetRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessResetRecipeStepsActionEnum(s)
	return &v
}

func (v InstancePreprocessResetRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessResetRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum.
type InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum string

// InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceResetRecipeStepsActionEnum.
type InstanceResetRecipeStepsActionEnum string

// InstanceResetRecipeStepsActionEnumRef returns a *InstanceResetRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceResetRecipeStepsActionEnumRef(s string) *InstanceResetRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceResetRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceResetRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceResetRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceResetRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceResetRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceResetRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceResetRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceResetRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceResetRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceResetRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceResetRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceResetRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessRepairRecipeStepsActionEnum.
type InstancePreprocessRepairRecipeStepsActionEnum string

// InstancePreprocessRepairRecipeStepsActionEnumRef returns a *InstancePreprocessRepairRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessRepairRecipeStepsActionEnumRef(s string) *InstancePreprocessRepairRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessRepairRecipeStepsActionEnum(s)
	return &v
}

func (v InstancePreprocessRepairRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessRepairRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum.
type InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum string

// InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceRepairRecipeStepsActionEnum.
type InstanceRepairRecipeStepsActionEnum string

// InstanceRepairRecipeStepsActionEnumRef returns a *InstanceRepairRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceRepairRecipeStepsActionEnumRef(s string) *InstanceRepairRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceRepairRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceRepairRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceRepairRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceRepairRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceRepairRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessDeleteRecipeStepsActionEnum.
type InstancePreprocessDeleteRecipeStepsActionEnum string

// InstancePreprocessDeleteRecipeStepsActionEnumRef returns a *InstancePreprocessDeleteRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessDeleteRecipeStepsActionEnumRef(s string) *InstancePreprocessDeleteRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessDeleteRecipeStepsActionEnum(s)
	return &v
}

func (v InstancePreprocessDeleteRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessDeleteRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum.
type InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum string

// InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessUpdateRecipeStepsActionEnum.
type InstancePreprocessUpdateRecipeStepsActionEnum string

// InstancePreprocessUpdateRecipeStepsActionEnumRef returns a *InstancePreprocessUpdateRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessUpdateRecipeStepsActionEnumRef(s string) *InstancePreprocessUpdateRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessUpdateRecipeStepsActionEnum(s)
	return &v
}

func (v InstancePreprocessUpdateRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessUpdateRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum.
type InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum string

// InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessFreezeRecipeStepsActionEnum.
type InstancePreprocessFreezeRecipeStepsActionEnum string

// InstancePreprocessFreezeRecipeStepsActionEnumRef returns a *InstancePreprocessFreezeRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessFreezeRecipeStepsActionEnumRef(s string) *InstancePreprocessFreezeRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessFreezeRecipeStepsActionEnum(s)
	return &v
}

func (v InstancePreprocessFreezeRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessFreezeRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum.
type InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum string

// InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceFreezeRecipeStepsActionEnum.
type InstanceFreezeRecipeStepsActionEnum string

// InstanceFreezeRecipeStepsActionEnumRef returns a *InstanceFreezeRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceFreezeRecipeStepsActionEnumRef(s string) *InstanceFreezeRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceFreezeRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceFreezeRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceFreezeRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessUnfreezeRecipeStepsActionEnum.
type InstancePreprocessUnfreezeRecipeStepsActionEnum string

// InstancePreprocessUnfreezeRecipeStepsActionEnumRef returns a *InstancePreprocessUnfreezeRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessUnfreezeRecipeStepsActionEnumRef(s string) *InstancePreprocessUnfreezeRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessUnfreezeRecipeStepsActionEnum(s)
	return &v
}

func (v InstancePreprocessUnfreezeRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessUnfreezeRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.
type InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum string

// InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceUnfreezeRecipeStepsActionEnum.
type InstanceUnfreezeRecipeStepsActionEnum string

// InstanceUnfreezeRecipeStepsActionEnumRef returns a *InstanceUnfreezeRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceUnfreezeRecipeStepsActionEnumRef(s string) *InstanceUnfreezeRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceUnfreezeRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceUnfreezeRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceUnfreezeRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessReportInstanceHealthRecipeStepsActionEnum.
type InstancePreprocessReportInstanceHealthRecipeStepsActionEnum string

// InstancePreprocessReportInstanceHealthRecipeStepsActionEnumRef returns a *InstancePreprocessReportInstanceHealthRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessReportInstanceHealthRecipeStepsActionEnumRef(s string) *InstancePreprocessReportInstanceHealthRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessReportInstanceHealthRecipeStepsActionEnum(s)
	return &v
}

func (v InstancePreprocessReportInstanceHealthRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessReportInstanceHealthRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum.
type InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum string

// InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceReportInstanceHealthRecipeStepsActionEnum.
type InstanceReportInstanceHealthRecipeStepsActionEnum string

// InstanceReportInstanceHealthRecipeStepsActionEnumRef returns a *InstanceReportInstanceHealthRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceReportInstanceHealthRecipeStepsActionEnumRef(s string) *InstanceReportInstanceHealthRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceReportInstanceHealthRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceReportInstanceHealthRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceReportInstanceHealthRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessGetRecipeStepsActionEnum.
type InstancePreprocessGetRecipeStepsActionEnum string

// InstancePreprocessGetRecipeStepsActionEnumRef returns a *InstancePreprocessGetRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessGetRecipeStepsActionEnumRef(s string) *InstancePreprocessGetRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessGetRecipeStepsActionEnum(s)
	return &v
}

func (v InstancePreprocessGetRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessGetRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum.
type InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum string

// InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceNotifyKeyAvailableRecipeStepsActionEnum.
type InstanceNotifyKeyAvailableRecipeStepsActionEnum string

// InstanceNotifyKeyAvailableRecipeStepsActionEnumRef returns a *InstanceNotifyKeyAvailableRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceNotifyKeyAvailableRecipeStepsActionEnumRef(s string) *InstanceNotifyKeyAvailableRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceNotifyKeyAvailableRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceNotifyKeyAvailableRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceNotifyKeyAvailableRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceNotifyKeyUnavailableRecipeStepsActionEnum.
type InstanceNotifyKeyUnavailableRecipeStepsActionEnum string

// InstanceNotifyKeyUnavailableRecipeStepsActionEnumRef returns a *InstanceNotifyKeyUnavailableRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceNotifyKeyUnavailableRecipeStepsActionEnumRef(s string) *InstanceNotifyKeyUnavailableRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceNotifyKeyUnavailableRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceNotifyKeyUnavailableRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceNotifyKeyUnavailableRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceReadonlyRecipeStepsActionEnum.
type InstanceReadonlyRecipeStepsActionEnum string

// InstanceReadonlyRecipeStepsActionEnumRef returns a *InstanceReadonlyRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceReadonlyRecipeStepsActionEnumRef(s string) *InstanceReadonlyRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceReadonlyRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceReadonlyRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceReadonlyRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceReconcileRecipeStepsActionEnum.
type InstanceReconcileRecipeStepsActionEnum string

// InstanceReconcileRecipeStepsActionEnumRef returns a *InstanceReconcileRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceReconcileRecipeStepsActionEnumRef(s string) *InstanceReconcileRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceReconcileRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceReconcileRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceReconcileRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessPassthroughRecipeStepsActionEnum.
type InstancePreprocessPassthroughRecipeStepsActionEnum string

// InstancePreprocessPassthroughRecipeStepsActionEnumRef returns a *InstancePreprocessPassthroughRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessPassthroughRecipeStepsActionEnumRef(s string) *InstancePreprocessPassthroughRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessPassthroughRecipeStepsActionEnum(s)
	return &v
}

func (v InstancePreprocessPassthroughRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessPassthroughRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum.
type InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum string

// InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessReconcileRecipeStepsActionEnum.
type InstancePreprocessReconcileRecipeStepsActionEnum string

// InstancePreprocessReconcileRecipeStepsActionEnumRef returns a *InstancePreprocessReconcileRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessReconcileRecipeStepsActionEnumRef(s string) *InstancePreprocessReconcileRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessReconcileRecipeStepsActionEnum(s)
	return &v
}

func (v InstancePreprocessReconcileRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessReconcileRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum.
type InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum string

// InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type InstanceSku struct {
	empty bool                 `json:"-"`
	Tier  *InstanceSkuTierEnum `json:"tier"`
	Size  *InstanceSkuSizeEnum `json:"size"`
}

// This object is used to assert a desired state where this InstanceSku is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceSku *InstanceSku = &InstanceSku{empty: true}

func (r *InstanceSku) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceSku) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReferences struct {
	empty          bool                        `json:"-"`
	Name           *string                     `json:"name"`
	Type           *string                     `json:"type"`
	SourceResource *string                     `json:"sourceResource"`
	Details        []InstanceReferencesDetails `json:"details"`
	CreateTime     *string                     `json:"createTime"`
}

// This object is used to assert a desired state where this InstanceReferences is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReferences *InstanceReferences = &InstanceReferences{empty: true}

func (r *InstanceReferences) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReferences) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReferencesDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceReferencesDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReferencesDetails *InstanceReferencesDetails = &InstanceReferencesDetails{empty: true}

func (r *InstanceReferencesDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReferencesDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessCreateRecipe struct {
	empty                             bool                                  `json:"-"`
	Steps                             []InstancePreprocessCreateRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                                 `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                                `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                              `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                                 `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                               `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                              `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                                `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipe *InstancePreprocessCreateRecipe = &InstancePreprocessCreateRecipe{empty: true}

func (r *InstancePreprocessCreateRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessCreateRecipeSteps struct {
	empty                          bool                                                       `json:"-"`
	RelativeTime                   *int64                                                     `json:"relativeTime"`
	SleepDuration                  *int64                                                     `json:"sleepDuration"`
	Action                         *InstancePreprocessCreateRecipeStepsActionEnum             `json:"action"`
	Status                         *InstancePreprocessCreateRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                    `json:"errorSpace"`
	P4ServiceAccount               *string                                                    `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                     `json:"resourceMetadataSize"`
	Description                    *string                                                    `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                   `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstancePreprocessCreateRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstancePreprocessCreateRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                    `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstancePreprocessCreateRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstancePreprocessCreateRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                    `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeSteps *InstancePreprocessCreateRecipeSteps = &InstancePreprocessCreateRecipeSteps{empty: true}

func (r *InstancePreprocessCreateRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessCreateRecipeStepsStatus struct {
	empty   bool                                               `json:"-"`
	Code    *int64                                             `json:"code"`
	Message *string                                            `json:"message"`
	Details []InstancePreprocessCreateRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeStepsStatus *InstancePreprocessCreateRecipeStepsStatus = &InstancePreprocessCreateRecipeStepsStatus{empty: true}

func (r *InstancePreprocessCreateRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessCreateRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeStepsStatusDetails *InstancePreprocessCreateRecipeStepsStatusDetails = &InstancePreprocessCreateRecipeStepsStatusDetails{empty: true}

func (r *InstancePreprocessCreateRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessCreateRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeStepsQuotaRequestDeltas *InstancePreprocessCreateRecipeStepsQuotaRequestDeltas = &InstancePreprocessCreateRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstancePreprocessCreateRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessCreateRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeStepsPreprocessUpdate *InstancePreprocessCreateRecipeStepsPreprocessUpdate = &InstancePreprocessCreateRecipeStepsPreprocessUpdate{empty: true}

func (r *InstancePreprocessCreateRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessCreateRecipeStepsRequestedTenantProject struct {
	empty  bool                                                                `json:"-"`
	Tag    *string                                                             `json:"tag"`
	Folder *string                                                             `json:"folder"`
	Scope  *InstancePreprocessCreateRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeStepsRequestedTenantProject *InstancePreprocessCreateRecipeStepsRequestedTenantProject = &InstancePreprocessCreateRecipeStepsRequestedTenantProject{empty: true}

func (r *InstancePreprocessCreateRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessCreateRecipeStepsPermissionsInfo struct {
	empty          bool                                                               `json:"-"`
	PolicyName     *InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                            `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                                      `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeStepsPermissionsInfo *InstancePreprocessCreateRecipeStepsPermissionsInfo = &InstancePreprocessCreateRecipeStepsPermissionsInfo{empty: true}

func (r *InstancePreprocessCreateRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName *InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName = &InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions *InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions = &InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceGoogleprotobufstruct struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstanceGoogleprotobufstruct is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceGoogleprotobufstruct *InstanceGoogleprotobufstruct = &InstanceGoogleprotobufstruct{empty: true}

func (r *InstanceGoogleprotobufstruct) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceGoogleprotobufstruct) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                           `json:"-"`
	KeyNotificationsInfo *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeStepsKeyNotificationsUpdate *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate = &InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                                  `json:"-"`
	DataVersion            *int64                                                                                                `json:"dataVersion"`
	Delegate               *string                                                                                               `json:"delegate"`
	KeyNotificationConfigs []InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCreateRecipe struct {
	empty                             bool                        `json:"-"`
	Steps                             []InstanceCreateRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                       `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                      `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                    `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                       `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                     `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                    `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                      `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceCreateRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipe *InstanceCreateRecipe = &InstanceCreateRecipe{empty: true}

func (r *InstanceCreateRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCreateRecipeSteps struct {
	empty                          bool                                             `json:"-"`
	RelativeTime                   *int64                                           `json:"relativeTime"`
	SleepDuration                  *int64                                           `json:"sleepDuration"`
	Action                         *InstanceCreateRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceCreateRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                          `json:"errorSpace"`
	P4ServiceAccount               *string                                          `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                           `json:"resourceMetadataSize"`
	Description                    *string                                          `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                         `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceCreateRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceCreateRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                          `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceCreateRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceCreateRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceCreateRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                          `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeSteps *InstanceCreateRecipeSteps = &InstanceCreateRecipeSteps{empty: true}

func (r *InstanceCreateRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCreateRecipeStepsStatus struct {
	empty   bool                                     `json:"-"`
	Code    *int64                                   `json:"code"`
	Message *string                                  `json:"message"`
	Details []InstanceCreateRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeStepsStatus *InstanceCreateRecipeStepsStatus = &InstanceCreateRecipeStepsStatus{empty: true}

func (r *InstanceCreateRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCreateRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeStepsStatusDetails *InstanceCreateRecipeStepsStatusDetails = &InstanceCreateRecipeStepsStatusDetails{empty: true}

func (r *InstanceCreateRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCreateRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeStepsQuotaRequestDeltas *InstanceCreateRecipeStepsQuotaRequestDeltas = &InstanceCreateRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceCreateRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCreateRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeStepsPreprocessUpdate *InstanceCreateRecipeStepsPreprocessUpdate = &InstanceCreateRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceCreateRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCreateRecipeStepsRequestedTenantProject struct {
	empty  bool                                                      `json:"-"`
	Tag    *string                                                   `json:"tag"`
	Folder *string                                                   `json:"folder"`
	Scope  *InstanceCreateRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeStepsRequestedTenantProject *InstanceCreateRecipeStepsRequestedTenantProject = &InstanceCreateRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceCreateRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCreateRecipeStepsPermissionsInfo struct {
	empty          bool                                                     `json:"-"`
	PolicyName     *InstanceCreateRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceCreateRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                  `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                            `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeStepsPermissionsInfo *InstanceCreateRecipeStepsPermissionsInfo = &InstanceCreateRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceCreateRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCreateRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeStepsPermissionsInfoPolicyName *InstanceCreateRecipeStepsPermissionsInfoPolicyName = &InstanceCreateRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceCreateRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCreateRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeStepsPermissionsInfoIamPermissions *InstanceCreateRecipeStepsPermissionsInfoIamPermissions = &InstanceCreateRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceCreateRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCreateRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                 `json:"-"`
	KeyNotificationsInfo *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeStepsKeyNotificationsUpdate *InstanceCreateRecipeStepsKeyNotificationsUpdate = &InstanceCreateRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceCreateRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                        `json:"-"`
	DataVersion            *int64                                                                                      `json:"dataVersion"`
	Delegate               *string                                                                                     `json:"delegate"`
	KeyNotificationConfigs []InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDeleteRecipe struct {
	empty                             bool                        `json:"-"`
	Steps                             []InstanceDeleteRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                       `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                      `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                    `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                       `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                     `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                    `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                      `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipe *InstanceDeleteRecipe = &InstanceDeleteRecipe{empty: true}

func (r *InstanceDeleteRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDeleteRecipeSteps struct {
	empty                          bool                                             `json:"-"`
	RelativeTime                   *int64                                           `json:"relativeTime"`
	SleepDuration                  *int64                                           `json:"sleepDuration"`
	Action                         *InstanceDeleteRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceDeleteRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                          `json:"errorSpace"`
	P4ServiceAccount               *string                                          `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                           `json:"resourceMetadataSize"`
	Description                    *string                                          `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                         `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceDeleteRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceDeleteRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                          `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceDeleteRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceDeleteRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceDeleteRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                          `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeSteps *InstanceDeleteRecipeSteps = &InstanceDeleteRecipeSteps{empty: true}

func (r *InstanceDeleteRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDeleteRecipeStepsStatus struct {
	empty   bool                                     `json:"-"`
	Code    *int64                                   `json:"code"`
	Message *string                                  `json:"message"`
	Details []InstanceDeleteRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeStepsStatus *InstanceDeleteRecipeStepsStatus = &InstanceDeleteRecipeStepsStatus{empty: true}

func (r *InstanceDeleteRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDeleteRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeStepsStatusDetails *InstanceDeleteRecipeStepsStatusDetails = &InstanceDeleteRecipeStepsStatusDetails{empty: true}

func (r *InstanceDeleteRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDeleteRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeStepsQuotaRequestDeltas *InstanceDeleteRecipeStepsQuotaRequestDeltas = &InstanceDeleteRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceDeleteRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDeleteRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeStepsPreprocessUpdate *InstanceDeleteRecipeStepsPreprocessUpdate = &InstanceDeleteRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceDeleteRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDeleteRecipeStepsRequestedTenantProject struct {
	empty  bool                                                      `json:"-"`
	Tag    *string                                                   `json:"tag"`
	Folder *string                                                   `json:"folder"`
	Scope  *InstanceDeleteRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeStepsRequestedTenantProject *InstanceDeleteRecipeStepsRequestedTenantProject = &InstanceDeleteRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceDeleteRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDeleteRecipeStepsPermissionsInfo struct {
	empty          bool                                                     `json:"-"`
	PolicyName     *InstanceDeleteRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceDeleteRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                  `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                            `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeStepsPermissionsInfo *InstanceDeleteRecipeStepsPermissionsInfo = &InstanceDeleteRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceDeleteRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDeleteRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeStepsPermissionsInfoPolicyName *InstanceDeleteRecipeStepsPermissionsInfoPolicyName = &InstanceDeleteRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceDeleteRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDeleteRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeStepsPermissionsInfoIamPermissions *InstanceDeleteRecipeStepsPermissionsInfoIamPermissions = &InstanceDeleteRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceDeleteRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDeleteRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                 `json:"-"`
	KeyNotificationsInfo *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeStepsKeyNotificationsUpdate *InstanceDeleteRecipeStepsKeyNotificationsUpdate = &InstanceDeleteRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceDeleteRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                        `json:"-"`
	DataVersion            *int64                                                                                      `json:"dataVersion"`
	Delegate               *string                                                                                     `json:"delegate"`
	KeyNotificationConfigs []InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUpdateRecipe struct {
	empty                             bool                        `json:"-"`
	Steps                             []InstanceUpdateRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                       `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                      `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                    `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                       `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                     `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                    `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                      `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipe *InstanceUpdateRecipe = &InstanceUpdateRecipe{empty: true}

func (r *InstanceUpdateRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUpdateRecipeSteps struct {
	empty                          bool                                             `json:"-"`
	RelativeTime                   *int64                                           `json:"relativeTime"`
	SleepDuration                  *int64                                           `json:"sleepDuration"`
	Action                         *InstanceUpdateRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceUpdateRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                          `json:"errorSpace"`
	P4ServiceAccount               *string                                          `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                           `json:"resourceMetadataSize"`
	Description                    *string                                          `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                         `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceUpdateRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceUpdateRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                          `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceUpdateRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceUpdateRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceUpdateRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                          `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeSteps *InstanceUpdateRecipeSteps = &InstanceUpdateRecipeSteps{empty: true}

func (r *InstanceUpdateRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUpdateRecipeStepsStatus struct {
	empty   bool                                     `json:"-"`
	Code    *int64                                   `json:"code"`
	Message *string                                  `json:"message"`
	Details []InstanceUpdateRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeStepsStatus *InstanceUpdateRecipeStepsStatus = &InstanceUpdateRecipeStepsStatus{empty: true}

func (r *InstanceUpdateRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUpdateRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeStepsStatusDetails *InstanceUpdateRecipeStepsStatusDetails = &InstanceUpdateRecipeStepsStatusDetails{empty: true}

func (r *InstanceUpdateRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUpdateRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeStepsQuotaRequestDeltas *InstanceUpdateRecipeStepsQuotaRequestDeltas = &InstanceUpdateRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceUpdateRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUpdateRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeStepsPreprocessUpdate *InstanceUpdateRecipeStepsPreprocessUpdate = &InstanceUpdateRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceUpdateRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUpdateRecipeStepsRequestedTenantProject struct {
	empty  bool                                                      `json:"-"`
	Tag    *string                                                   `json:"tag"`
	Folder *string                                                   `json:"folder"`
	Scope  *InstanceUpdateRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeStepsRequestedTenantProject *InstanceUpdateRecipeStepsRequestedTenantProject = &InstanceUpdateRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceUpdateRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUpdateRecipeStepsPermissionsInfo struct {
	empty          bool                                                     `json:"-"`
	PolicyName     *InstanceUpdateRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceUpdateRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                  `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                            `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeStepsPermissionsInfo *InstanceUpdateRecipeStepsPermissionsInfo = &InstanceUpdateRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceUpdateRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUpdateRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeStepsPermissionsInfoPolicyName *InstanceUpdateRecipeStepsPermissionsInfoPolicyName = &InstanceUpdateRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceUpdateRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUpdateRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeStepsPermissionsInfoIamPermissions *InstanceUpdateRecipeStepsPermissionsInfoIamPermissions = &InstanceUpdateRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceUpdateRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUpdateRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                 `json:"-"`
	KeyNotificationsInfo *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeStepsKeyNotificationsUpdate *InstanceUpdateRecipeStepsKeyNotificationsUpdate = &InstanceUpdateRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceUpdateRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                        `json:"-"`
	DataVersion            *int64                                                                                      `json:"dataVersion"`
	Delegate               *string                                                                                     `json:"delegate"`
	KeyNotificationConfigs []InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessResetRecipe struct {
	empty                             bool                                 `json:"-"`
	Steps                             []InstancePreprocessResetRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                                `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                               `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                             `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                                `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                              `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                             `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                               `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipe *InstancePreprocessResetRecipe = &InstancePreprocessResetRecipe{empty: true}

func (r *InstancePreprocessResetRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessResetRecipeSteps struct {
	empty                          bool                                                      `json:"-"`
	RelativeTime                   *int64                                                    `json:"relativeTime"`
	SleepDuration                  *int64                                                    `json:"sleepDuration"`
	Action                         *InstancePreprocessResetRecipeStepsActionEnum             `json:"action"`
	Status                         *InstancePreprocessResetRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                   `json:"errorSpace"`
	P4ServiceAccount               *string                                                   `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                    `json:"resourceMetadataSize"`
	Description                    *string                                                   `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                  `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstancePreprocessResetRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstancePreprocessResetRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                   `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstancePreprocessResetRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstancePreprocessResetRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstancePreprocessResetRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                   `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeSteps *InstancePreprocessResetRecipeSteps = &InstancePreprocessResetRecipeSteps{empty: true}

func (r *InstancePreprocessResetRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessResetRecipeStepsStatus struct {
	empty   bool                                              `json:"-"`
	Code    *int64                                            `json:"code"`
	Message *string                                           `json:"message"`
	Details []InstancePreprocessResetRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeStepsStatus *InstancePreprocessResetRecipeStepsStatus = &InstancePreprocessResetRecipeStepsStatus{empty: true}

func (r *InstancePreprocessResetRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessResetRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeStepsStatusDetails *InstancePreprocessResetRecipeStepsStatusDetails = &InstancePreprocessResetRecipeStepsStatusDetails{empty: true}

func (r *InstancePreprocessResetRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessResetRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeStepsQuotaRequestDeltas *InstancePreprocessResetRecipeStepsQuotaRequestDeltas = &InstancePreprocessResetRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstancePreprocessResetRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessResetRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeStepsPreprocessUpdate *InstancePreprocessResetRecipeStepsPreprocessUpdate = &InstancePreprocessResetRecipeStepsPreprocessUpdate{empty: true}

func (r *InstancePreprocessResetRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessResetRecipeStepsRequestedTenantProject struct {
	empty  bool                                                               `json:"-"`
	Tag    *string                                                            `json:"tag"`
	Folder *string                                                            `json:"folder"`
	Scope  *InstancePreprocessResetRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeStepsRequestedTenantProject *InstancePreprocessResetRecipeStepsRequestedTenantProject = &InstancePreprocessResetRecipeStepsRequestedTenantProject{empty: true}

func (r *InstancePreprocessResetRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessResetRecipeStepsPermissionsInfo struct {
	empty          bool                                                              `json:"-"`
	PolicyName     *InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                           `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                                     `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeStepsPermissionsInfo *InstancePreprocessResetRecipeStepsPermissionsInfo = &InstancePreprocessResetRecipeStepsPermissionsInfo{empty: true}

func (r *InstancePreprocessResetRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeStepsPermissionsInfoPolicyName *InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName = &InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions *InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions = &InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessResetRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                          `json:"-"`
	KeyNotificationsInfo *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeStepsKeyNotificationsUpdate *InstancePreprocessResetRecipeStepsKeyNotificationsUpdate = &InstancePreprocessResetRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstancePreprocessResetRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                                 `json:"-"`
	DataVersion            *int64                                                                                               `json:"dataVersion"`
	Delegate               *string                                                                                              `json:"delegate"`
	KeyNotificationConfigs []InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceResetRecipe struct {
	empty                             bool                       `json:"-"`
	Steps                             []InstanceResetRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                      `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                     `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                   `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                      `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                    `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                   `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                     `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceResetRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipe *InstanceResetRecipe = &InstanceResetRecipe{empty: true}

func (r *InstanceResetRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceResetRecipeSteps struct {
	empty                          bool                                            `json:"-"`
	RelativeTime                   *int64                                          `json:"relativeTime"`
	SleepDuration                  *int64                                          `json:"sleepDuration"`
	Action                         *InstanceResetRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceResetRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                         `json:"errorSpace"`
	P4ServiceAccount               *string                                         `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                          `json:"resourceMetadataSize"`
	Description                    *string                                         `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                        `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceResetRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceResetRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                         `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceResetRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceResetRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceResetRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                         `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceResetRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeSteps *InstanceResetRecipeSteps = &InstanceResetRecipeSteps{empty: true}

func (r *InstanceResetRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceResetRecipeStepsStatus struct {
	empty   bool                                    `json:"-"`
	Code    *int64                                  `json:"code"`
	Message *string                                 `json:"message"`
	Details []InstanceResetRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceResetRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeStepsStatus *InstanceResetRecipeStepsStatus = &InstanceResetRecipeStepsStatus{empty: true}

func (r *InstanceResetRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceResetRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceResetRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeStepsStatusDetails *InstanceResetRecipeStepsStatusDetails = &InstanceResetRecipeStepsStatusDetails{empty: true}

func (r *InstanceResetRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceResetRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceResetRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeStepsQuotaRequestDeltas *InstanceResetRecipeStepsQuotaRequestDeltas = &InstanceResetRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceResetRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceResetRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceResetRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeStepsPreprocessUpdate *InstanceResetRecipeStepsPreprocessUpdate = &InstanceResetRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceResetRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceResetRecipeStepsRequestedTenantProject struct {
	empty  bool                                                     `json:"-"`
	Tag    *string                                                  `json:"tag"`
	Folder *string                                                  `json:"folder"`
	Scope  *InstanceResetRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceResetRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeStepsRequestedTenantProject *InstanceResetRecipeStepsRequestedTenantProject = &InstanceResetRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceResetRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceResetRecipeStepsPermissionsInfo struct {
	empty          bool                                                    `json:"-"`
	PolicyName     *InstanceResetRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceResetRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                 `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                           `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceResetRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeStepsPermissionsInfo *InstanceResetRecipeStepsPermissionsInfo = &InstanceResetRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceResetRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceResetRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceResetRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeStepsPermissionsInfoPolicyName *InstanceResetRecipeStepsPermissionsInfoPolicyName = &InstanceResetRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceResetRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceResetRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceResetRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeStepsPermissionsInfoIamPermissions *InstanceResetRecipeStepsPermissionsInfoIamPermissions = &InstanceResetRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceResetRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceResetRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                `json:"-"`
	KeyNotificationsInfo *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceResetRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeStepsKeyNotificationsUpdate *InstanceResetRecipeStepsKeyNotificationsUpdate = &InstanceResetRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceResetRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                       `json:"-"`
	DataVersion            *int64                                                                                     `json:"dataVersion"`
	Delegate               *string                                                                                    `json:"delegate"`
	KeyNotificationConfigs []InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessRepairRecipe struct {
	empty                             bool                                  `json:"-"`
	Steps                             []InstancePreprocessRepairRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                                 `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                                `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                              `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                                 `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                               `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                              `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                                `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipe *InstancePreprocessRepairRecipe = &InstancePreprocessRepairRecipe{empty: true}

func (r *InstancePreprocessRepairRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessRepairRecipeSteps struct {
	empty                          bool                                                       `json:"-"`
	RelativeTime                   *int64                                                     `json:"relativeTime"`
	SleepDuration                  *int64                                                     `json:"sleepDuration"`
	Action                         *InstancePreprocessRepairRecipeStepsActionEnum             `json:"action"`
	Status                         *InstancePreprocessRepairRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                    `json:"errorSpace"`
	P4ServiceAccount               *string                                                    `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                     `json:"resourceMetadataSize"`
	Description                    *string                                                    `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                   `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstancePreprocessRepairRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstancePreprocessRepairRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                    `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstancePreprocessRepairRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstancePreprocessRepairRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                    `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeSteps *InstancePreprocessRepairRecipeSteps = &InstancePreprocessRepairRecipeSteps{empty: true}

func (r *InstancePreprocessRepairRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessRepairRecipeStepsStatus struct {
	empty   bool                                               `json:"-"`
	Code    *int64                                             `json:"code"`
	Message *string                                            `json:"message"`
	Details []InstancePreprocessRepairRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeStepsStatus *InstancePreprocessRepairRecipeStepsStatus = &InstancePreprocessRepairRecipeStepsStatus{empty: true}

func (r *InstancePreprocessRepairRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessRepairRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeStepsStatusDetails *InstancePreprocessRepairRecipeStepsStatusDetails = &InstancePreprocessRepairRecipeStepsStatusDetails{empty: true}

func (r *InstancePreprocessRepairRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessRepairRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeStepsQuotaRequestDeltas *InstancePreprocessRepairRecipeStepsQuotaRequestDeltas = &InstancePreprocessRepairRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstancePreprocessRepairRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessRepairRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeStepsPreprocessUpdate *InstancePreprocessRepairRecipeStepsPreprocessUpdate = &InstancePreprocessRepairRecipeStepsPreprocessUpdate{empty: true}

func (r *InstancePreprocessRepairRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessRepairRecipeStepsRequestedTenantProject struct {
	empty  bool                                                                `json:"-"`
	Tag    *string                                                             `json:"tag"`
	Folder *string                                                             `json:"folder"`
	Scope  *InstancePreprocessRepairRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeStepsRequestedTenantProject *InstancePreprocessRepairRecipeStepsRequestedTenantProject = &InstancePreprocessRepairRecipeStepsRequestedTenantProject{empty: true}

func (r *InstancePreprocessRepairRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessRepairRecipeStepsPermissionsInfo struct {
	empty          bool                                                               `json:"-"`
	PolicyName     *InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                            `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                                      `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeStepsPermissionsInfo *InstancePreprocessRepairRecipeStepsPermissionsInfo = &InstancePreprocessRepairRecipeStepsPermissionsInfo{empty: true}

func (r *InstancePreprocessRepairRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName *InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName = &InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions *InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions = &InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                           `json:"-"`
	KeyNotificationsInfo *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeStepsKeyNotificationsUpdate *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate = &InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                                  `json:"-"`
	DataVersion            *int64                                                                                                `json:"dataVersion"`
	Delegate               *string                                                                                               `json:"delegate"`
	KeyNotificationConfigs []InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRepairRecipe struct {
	empty                             bool                        `json:"-"`
	Steps                             []InstanceRepairRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                       `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                      `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                    `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                       `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                     `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                    `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                      `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceRepairRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipe *InstanceRepairRecipe = &InstanceRepairRecipe{empty: true}

func (r *InstanceRepairRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRepairRecipeSteps struct {
	empty                          bool                                             `json:"-"`
	RelativeTime                   *int64                                           `json:"relativeTime"`
	SleepDuration                  *int64                                           `json:"sleepDuration"`
	Action                         *InstanceRepairRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceRepairRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                          `json:"errorSpace"`
	P4ServiceAccount               *string                                          `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                           `json:"resourceMetadataSize"`
	Description                    *string                                          `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                         `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceRepairRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceRepairRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                          `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceRepairRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceRepairRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceRepairRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                          `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeSteps *InstanceRepairRecipeSteps = &InstanceRepairRecipeSteps{empty: true}

func (r *InstanceRepairRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRepairRecipeStepsStatus struct {
	empty   bool                                     `json:"-"`
	Code    *int64                                   `json:"code"`
	Message *string                                  `json:"message"`
	Details []InstanceRepairRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeStepsStatus *InstanceRepairRecipeStepsStatus = &InstanceRepairRecipeStepsStatus{empty: true}

func (r *InstanceRepairRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRepairRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeStepsStatusDetails *InstanceRepairRecipeStepsStatusDetails = &InstanceRepairRecipeStepsStatusDetails{empty: true}

func (r *InstanceRepairRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRepairRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeStepsQuotaRequestDeltas *InstanceRepairRecipeStepsQuotaRequestDeltas = &InstanceRepairRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceRepairRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRepairRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeStepsPreprocessUpdate *InstanceRepairRecipeStepsPreprocessUpdate = &InstanceRepairRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceRepairRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRepairRecipeStepsRequestedTenantProject struct {
	empty  bool                                                      `json:"-"`
	Tag    *string                                                   `json:"tag"`
	Folder *string                                                   `json:"folder"`
	Scope  *InstanceRepairRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeStepsRequestedTenantProject *InstanceRepairRecipeStepsRequestedTenantProject = &InstanceRepairRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceRepairRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRepairRecipeStepsPermissionsInfo struct {
	empty          bool                                                     `json:"-"`
	PolicyName     *InstanceRepairRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceRepairRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                  `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                            `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeStepsPermissionsInfo *InstanceRepairRecipeStepsPermissionsInfo = &InstanceRepairRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceRepairRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRepairRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeStepsPermissionsInfoPolicyName *InstanceRepairRecipeStepsPermissionsInfoPolicyName = &InstanceRepairRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceRepairRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRepairRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeStepsPermissionsInfoIamPermissions *InstanceRepairRecipeStepsPermissionsInfoIamPermissions = &InstanceRepairRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceRepairRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRepairRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                 `json:"-"`
	KeyNotificationsInfo *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeStepsKeyNotificationsUpdate *InstanceRepairRecipeStepsKeyNotificationsUpdate = &InstanceRepairRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceRepairRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                        `json:"-"`
	DataVersion            *int64                                                                                      `json:"dataVersion"`
	Delegate               *string                                                                                     `json:"delegate"`
	KeyNotificationConfigs []InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessDeleteRecipe struct {
	empty                             bool                                  `json:"-"`
	Steps                             []InstancePreprocessDeleteRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                                 `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                                `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                              `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                                 `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                               `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                              `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                                `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipe *InstancePreprocessDeleteRecipe = &InstancePreprocessDeleteRecipe{empty: true}

func (r *InstancePreprocessDeleteRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessDeleteRecipeSteps struct {
	empty                          bool                                                       `json:"-"`
	RelativeTime                   *int64                                                     `json:"relativeTime"`
	SleepDuration                  *int64                                                     `json:"sleepDuration"`
	Action                         *InstancePreprocessDeleteRecipeStepsActionEnum             `json:"action"`
	Status                         *InstancePreprocessDeleteRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                    `json:"errorSpace"`
	P4ServiceAccount               *string                                                    `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                     `json:"resourceMetadataSize"`
	Description                    *string                                                    `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                   `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstancePreprocessDeleteRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                    `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstancePreprocessDeleteRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstancePreprocessDeleteRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                    `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeSteps *InstancePreprocessDeleteRecipeSteps = &InstancePreprocessDeleteRecipeSteps{empty: true}

func (r *InstancePreprocessDeleteRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessDeleteRecipeStepsStatus struct {
	empty   bool                                               `json:"-"`
	Code    *int64                                             `json:"code"`
	Message *string                                            `json:"message"`
	Details []InstancePreprocessDeleteRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeStepsStatus *InstancePreprocessDeleteRecipeStepsStatus = &InstancePreprocessDeleteRecipeStepsStatus{empty: true}

func (r *InstancePreprocessDeleteRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessDeleteRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeStepsStatusDetails *InstancePreprocessDeleteRecipeStepsStatusDetails = &InstancePreprocessDeleteRecipeStepsStatusDetails{empty: true}

func (r *InstancePreprocessDeleteRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeStepsQuotaRequestDeltas *InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas = &InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessDeleteRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeStepsPreprocessUpdate *InstancePreprocessDeleteRecipeStepsPreprocessUpdate = &InstancePreprocessDeleteRecipeStepsPreprocessUpdate{empty: true}

func (r *InstancePreprocessDeleteRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessDeleteRecipeStepsRequestedTenantProject struct {
	empty  bool                                                                `json:"-"`
	Tag    *string                                                             `json:"tag"`
	Folder *string                                                             `json:"folder"`
	Scope  *InstancePreprocessDeleteRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeStepsRequestedTenantProject *InstancePreprocessDeleteRecipeStepsRequestedTenantProject = &InstancePreprocessDeleteRecipeStepsRequestedTenantProject{empty: true}

func (r *InstancePreprocessDeleteRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessDeleteRecipeStepsPermissionsInfo struct {
	empty          bool                                                               `json:"-"`
	PolicyName     *InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                            `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                                      `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeStepsPermissionsInfo *InstancePreprocessDeleteRecipeStepsPermissionsInfo = &InstancePreprocessDeleteRecipeStepsPermissionsInfo{empty: true}

func (r *InstancePreprocessDeleteRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName *InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName = &InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions *InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions = &InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                           `json:"-"`
	KeyNotificationsInfo *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate = &InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                                  `json:"-"`
	DataVersion            *int64                                                                                                `json:"dataVersion"`
	Delegate               *string                                                                                               `json:"delegate"`
	KeyNotificationConfigs []InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUpdateRecipe struct {
	empty                             bool                                  `json:"-"`
	Steps                             []InstancePreprocessUpdateRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                                 `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                                `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                              `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                                 `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                               `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                              `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                                `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipe *InstancePreprocessUpdateRecipe = &InstancePreprocessUpdateRecipe{empty: true}

func (r *InstancePreprocessUpdateRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUpdateRecipeSteps struct {
	empty                          bool                                                       `json:"-"`
	RelativeTime                   *int64                                                     `json:"relativeTime"`
	SleepDuration                  *int64                                                     `json:"sleepDuration"`
	Action                         *InstancePreprocessUpdateRecipeStepsActionEnum             `json:"action"`
	Status                         *InstancePreprocessUpdateRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                    `json:"errorSpace"`
	P4ServiceAccount               *string                                                    `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                     `json:"resourceMetadataSize"`
	Description                    *string                                                    `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                   `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstancePreprocessUpdateRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                    `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstancePreprocessUpdateRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstancePreprocessUpdateRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                    `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeSteps *InstancePreprocessUpdateRecipeSteps = &InstancePreprocessUpdateRecipeSteps{empty: true}

func (r *InstancePreprocessUpdateRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUpdateRecipeStepsStatus struct {
	empty   bool                                               `json:"-"`
	Code    *int64                                             `json:"code"`
	Message *string                                            `json:"message"`
	Details []InstancePreprocessUpdateRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeStepsStatus *InstancePreprocessUpdateRecipeStepsStatus = &InstancePreprocessUpdateRecipeStepsStatus{empty: true}

func (r *InstancePreprocessUpdateRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUpdateRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeStepsStatusDetails *InstancePreprocessUpdateRecipeStepsStatusDetails = &InstancePreprocessUpdateRecipeStepsStatusDetails{empty: true}

func (r *InstancePreprocessUpdateRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeStepsQuotaRequestDeltas *InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas = &InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUpdateRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeStepsPreprocessUpdate *InstancePreprocessUpdateRecipeStepsPreprocessUpdate = &InstancePreprocessUpdateRecipeStepsPreprocessUpdate{empty: true}

func (r *InstancePreprocessUpdateRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUpdateRecipeStepsRequestedTenantProject struct {
	empty  bool                                                                `json:"-"`
	Tag    *string                                                             `json:"tag"`
	Folder *string                                                             `json:"folder"`
	Scope  *InstancePreprocessUpdateRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeStepsRequestedTenantProject *InstancePreprocessUpdateRecipeStepsRequestedTenantProject = &InstancePreprocessUpdateRecipeStepsRequestedTenantProject{empty: true}

func (r *InstancePreprocessUpdateRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUpdateRecipeStepsPermissionsInfo struct {
	empty          bool                                                               `json:"-"`
	PolicyName     *InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                            `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                                      `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeStepsPermissionsInfo *InstancePreprocessUpdateRecipeStepsPermissionsInfo = &InstancePreprocessUpdateRecipeStepsPermissionsInfo{empty: true}

func (r *InstancePreprocessUpdateRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName *InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName = &InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions *InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions = &InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                           `json:"-"`
	KeyNotificationsInfo *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate = &InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                                  `json:"-"`
	DataVersion            *int64                                                                                                `json:"dataVersion"`
	Delegate               *string                                                                                               `json:"delegate"`
	KeyNotificationConfigs []InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessFreezeRecipe struct {
	empty                             bool                                  `json:"-"`
	Steps                             []InstancePreprocessFreezeRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                                 `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                                `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                              `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                                 `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                               `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                              `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                                `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipe *InstancePreprocessFreezeRecipe = &InstancePreprocessFreezeRecipe{empty: true}

func (r *InstancePreprocessFreezeRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessFreezeRecipeSteps struct {
	empty                          bool                                                       `json:"-"`
	RelativeTime                   *int64                                                     `json:"relativeTime"`
	SleepDuration                  *int64                                                     `json:"sleepDuration"`
	Action                         *InstancePreprocessFreezeRecipeStepsActionEnum             `json:"action"`
	Status                         *InstancePreprocessFreezeRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                    `json:"errorSpace"`
	P4ServiceAccount               *string                                                    `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                     `json:"resourceMetadataSize"`
	Description                    *string                                                    `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                   `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstancePreprocessFreezeRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                    `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstancePreprocessFreezeRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstancePreprocessFreezeRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                    `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeSteps *InstancePreprocessFreezeRecipeSteps = &InstancePreprocessFreezeRecipeSteps{empty: true}

func (r *InstancePreprocessFreezeRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessFreezeRecipeStepsStatus struct {
	empty   bool                                               `json:"-"`
	Code    *int64                                             `json:"code"`
	Message *string                                            `json:"message"`
	Details []InstancePreprocessFreezeRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeStepsStatus *InstancePreprocessFreezeRecipeStepsStatus = &InstancePreprocessFreezeRecipeStepsStatus{empty: true}

func (r *InstancePreprocessFreezeRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessFreezeRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeStepsStatusDetails *InstancePreprocessFreezeRecipeStepsStatusDetails = &InstancePreprocessFreezeRecipeStepsStatusDetails{empty: true}

func (r *InstancePreprocessFreezeRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeStepsQuotaRequestDeltas *InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas = &InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessFreezeRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeStepsPreprocessUpdate *InstancePreprocessFreezeRecipeStepsPreprocessUpdate = &InstancePreprocessFreezeRecipeStepsPreprocessUpdate{empty: true}

func (r *InstancePreprocessFreezeRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessFreezeRecipeStepsRequestedTenantProject struct {
	empty  bool                                                                `json:"-"`
	Tag    *string                                                             `json:"tag"`
	Folder *string                                                             `json:"folder"`
	Scope  *InstancePreprocessFreezeRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeStepsRequestedTenantProject *InstancePreprocessFreezeRecipeStepsRequestedTenantProject = &InstancePreprocessFreezeRecipeStepsRequestedTenantProject{empty: true}

func (r *InstancePreprocessFreezeRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessFreezeRecipeStepsPermissionsInfo struct {
	empty          bool                                                               `json:"-"`
	PolicyName     *InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                            `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                                      `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeStepsPermissionsInfo *InstancePreprocessFreezeRecipeStepsPermissionsInfo = &InstancePreprocessFreezeRecipeStepsPermissionsInfo{empty: true}

func (r *InstancePreprocessFreezeRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName *InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName = &InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions *InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions = &InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                           `json:"-"`
	KeyNotificationsInfo *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate = &InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                                  `json:"-"`
	DataVersion            *int64                                                                                                `json:"dataVersion"`
	Delegate               *string                                                                                               `json:"delegate"`
	KeyNotificationConfigs []InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFreezeRecipe struct {
	empty                             bool                        `json:"-"`
	Steps                             []InstanceFreezeRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                       `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                      `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                    `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                       `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                     `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                    `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                      `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipe *InstanceFreezeRecipe = &InstanceFreezeRecipe{empty: true}

func (r *InstanceFreezeRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFreezeRecipeSteps struct {
	empty                          bool                                             `json:"-"`
	RelativeTime                   *int64                                           `json:"relativeTime"`
	SleepDuration                  *int64                                           `json:"sleepDuration"`
	Action                         *InstanceFreezeRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceFreezeRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                          `json:"errorSpace"`
	P4ServiceAccount               *string                                          `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                           `json:"resourceMetadataSize"`
	Description                    *string                                          `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                         `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceFreezeRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceFreezeRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                          `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceFreezeRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceFreezeRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceFreezeRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                          `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeSteps *InstanceFreezeRecipeSteps = &InstanceFreezeRecipeSteps{empty: true}

func (r *InstanceFreezeRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFreezeRecipeStepsStatus struct {
	empty   bool                                     `json:"-"`
	Code    *int64                                   `json:"code"`
	Message *string                                  `json:"message"`
	Details []InstanceFreezeRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeStepsStatus *InstanceFreezeRecipeStepsStatus = &InstanceFreezeRecipeStepsStatus{empty: true}

func (r *InstanceFreezeRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFreezeRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeStepsStatusDetails *InstanceFreezeRecipeStepsStatusDetails = &InstanceFreezeRecipeStepsStatusDetails{empty: true}

func (r *InstanceFreezeRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFreezeRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeStepsQuotaRequestDeltas *InstanceFreezeRecipeStepsQuotaRequestDeltas = &InstanceFreezeRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceFreezeRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFreezeRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeStepsPreprocessUpdate *InstanceFreezeRecipeStepsPreprocessUpdate = &InstanceFreezeRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceFreezeRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFreezeRecipeStepsRequestedTenantProject struct {
	empty  bool                                                      `json:"-"`
	Tag    *string                                                   `json:"tag"`
	Folder *string                                                   `json:"folder"`
	Scope  *InstanceFreezeRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeStepsRequestedTenantProject *InstanceFreezeRecipeStepsRequestedTenantProject = &InstanceFreezeRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceFreezeRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFreezeRecipeStepsPermissionsInfo struct {
	empty          bool                                                     `json:"-"`
	PolicyName     *InstanceFreezeRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceFreezeRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                  `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                            `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeStepsPermissionsInfo *InstanceFreezeRecipeStepsPermissionsInfo = &InstanceFreezeRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceFreezeRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFreezeRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeStepsPermissionsInfoPolicyName *InstanceFreezeRecipeStepsPermissionsInfoPolicyName = &InstanceFreezeRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceFreezeRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFreezeRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeStepsPermissionsInfoIamPermissions *InstanceFreezeRecipeStepsPermissionsInfoIamPermissions = &InstanceFreezeRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceFreezeRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFreezeRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                 `json:"-"`
	KeyNotificationsInfo *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeStepsKeyNotificationsUpdate *InstanceFreezeRecipeStepsKeyNotificationsUpdate = &InstanceFreezeRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceFreezeRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                        `json:"-"`
	DataVersion            *int64                                                                                      `json:"dataVersion"`
	Delegate               *string                                                                                     `json:"delegate"`
	KeyNotificationConfigs []InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUnfreezeRecipe struct {
	empty                             bool                                    `json:"-"`
	Steps                             []InstancePreprocessUnfreezeRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                                   `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                                  `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                                `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                                   `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                                 `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                                `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                                  `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipe *InstancePreprocessUnfreezeRecipe = &InstancePreprocessUnfreezeRecipe{empty: true}

func (r *InstancePreprocessUnfreezeRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUnfreezeRecipeSteps struct {
	empty                          bool                                                         `json:"-"`
	RelativeTime                   *int64                                                       `json:"relativeTime"`
	SleepDuration                  *int64                                                       `json:"sleepDuration"`
	Action                         *InstancePreprocessUnfreezeRecipeStepsActionEnum             `json:"action"`
	Status                         *InstancePreprocessUnfreezeRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                      `json:"errorSpace"`
	P4ServiceAccount               *string                                                      `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                       `json:"resourceMetadataSize"`
	Description                    *string                                                      `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                     `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                      `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstancePreprocessUnfreezeRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                      `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeSteps *InstancePreprocessUnfreezeRecipeSteps = &InstancePreprocessUnfreezeRecipeSteps{empty: true}

func (r *InstancePreprocessUnfreezeRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUnfreezeRecipeStepsStatus struct {
	empty   bool                                                 `json:"-"`
	Code    *int64                                               `json:"code"`
	Message *string                                              `json:"message"`
	Details []InstancePreprocessUnfreezeRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeStepsStatus *InstancePreprocessUnfreezeRecipeStepsStatus = &InstancePreprocessUnfreezeRecipeStepsStatus{empty: true}

func (r *InstancePreprocessUnfreezeRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUnfreezeRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeStepsStatusDetails *InstancePreprocessUnfreezeRecipeStepsStatusDetails = &InstancePreprocessUnfreezeRecipeStepsStatusDetails{empty: true}

func (r *InstancePreprocessUnfreezeRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas *InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas = &InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeStepsPreprocessUpdate *InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate = &InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate{empty: true}

func (r *InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject struct {
	empty  bool                                                                  `json:"-"`
	Tag    *string                                                               `json:"tag"`
	Folder *string                                                               `json:"folder"`
	Scope  *InstancePreprocessUnfreezeRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeStepsRequestedTenantProject *InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject = &InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject{empty: true}

func (r *InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUnfreezeRecipeStepsPermissionsInfo struct {
	empty          bool                                                                 `json:"-"`
	PolicyName     *InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                              `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                                        `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeStepsPermissionsInfo *InstancePreprocessUnfreezeRecipeStepsPermissionsInfo = &InstancePreprocessUnfreezeRecipeStepsPermissionsInfo{empty: true}

func (r *InstancePreprocessUnfreezeRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName *InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName = &InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions *InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions = &InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                             `json:"-"`
	KeyNotificationsInfo *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate = &InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                                    `json:"-"`
	DataVersion            *int64                                                                                                  `json:"dataVersion"`
	Delegate               *string                                                                                                 `json:"delegate"`
	KeyNotificationConfigs []InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUnfreezeRecipe struct {
	empty                             bool                          `json:"-"`
	Steps                             []InstanceUnfreezeRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                         `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                        `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                      `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                         `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                       `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                      `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                        `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipe *InstanceUnfreezeRecipe = &InstanceUnfreezeRecipe{empty: true}

func (r *InstanceUnfreezeRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUnfreezeRecipeSteps struct {
	empty                          bool                                               `json:"-"`
	RelativeTime                   *int64                                             `json:"relativeTime"`
	SleepDuration                  *int64                                             `json:"sleepDuration"`
	Action                         *InstanceUnfreezeRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceUnfreezeRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                            `json:"errorSpace"`
	P4ServiceAccount               *string                                            `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                             `json:"resourceMetadataSize"`
	Description                    *string                                            `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                           `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceUnfreezeRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceUnfreezeRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                            `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceUnfreezeRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceUnfreezeRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceUnfreezeRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                            `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeSteps *InstanceUnfreezeRecipeSteps = &InstanceUnfreezeRecipeSteps{empty: true}

func (r *InstanceUnfreezeRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUnfreezeRecipeStepsStatus struct {
	empty   bool                                       `json:"-"`
	Code    *int64                                     `json:"code"`
	Message *string                                    `json:"message"`
	Details []InstanceUnfreezeRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeStepsStatus *InstanceUnfreezeRecipeStepsStatus = &InstanceUnfreezeRecipeStepsStatus{empty: true}

func (r *InstanceUnfreezeRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUnfreezeRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeStepsStatusDetails *InstanceUnfreezeRecipeStepsStatusDetails = &InstanceUnfreezeRecipeStepsStatusDetails{empty: true}

func (r *InstanceUnfreezeRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUnfreezeRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeStepsQuotaRequestDeltas *InstanceUnfreezeRecipeStepsQuotaRequestDeltas = &InstanceUnfreezeRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceUnfreezeRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUnfreezeRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeStepsPreprocessUpdate *InstanceUnfreezeRecipeStepsPreprocessUpdate = &InstanceUnfreezeRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceUnfreezeRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUnfreezeRecipeStepsRequestedTenantProject struct {
	empty  bool                                                        `json:"-"`
	Tag    *string                                                     `json:"tag"`
	Folder *string                                                     `json:"folder"`
	Scope  *InstanceUnfreezeRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeStepsRequestedTenantProject *InstanceUnfreezeRecipeStepsRequestedTenantProject = &InstanceUnfreezeRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceUnfreezeRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUnfreezeRecipeStepsPermissionsInfo struct {
	empty          bool                                                       `json:"-"`
	PolicyName     *InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                    `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                              `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeStepsPermissionsInfo *InstanceUnfreezeRecipeStepsPermissionsInfo = &InstanceUnfreezeRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceUnfreezeRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeStepsPermissionsInfoPolicyName *InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName = &InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions *InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions = &InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUnfreezeRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                   `json:"-"`
	KeyNotificationsInfo *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeStepsKeyNotificationsUpdate *InstanceUnfreezeRecipeStepsKeyNotificationsUpdate = &InstanceUnfreezeRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceUnfreezeRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                          `json:"-"`
	DataVersion            *int64                                                                                        `json:"dataVersion"`
	Delegate               *string                                                                                       `json:"delegate"`
	KeyNotificationConfigs []InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReportInstanceHealthRecipe struct {
	empty                             bool                                                `json:"-"`
	Steps                             []InstancePreprocessReportInstanceHealthRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                                               `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                                              `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                                            `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                                               `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                                             `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                                            `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                                              `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstancePreprocessReportInstanceHealthRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReportInstanceHealthRecipe *InstancePreprocessReportInstanceHealthRecipe = &InstancePreprocessReportInstanceHealthRecipe{empty: true}

func (r *InstancePreprocessReportInstanceHealthRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReportInstanceHealthRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReportInstanceHealthRecipeSteps struct {
	empty                          bool                                                                     `json:"-"`
	RelativeTime                   *int64                                                                   `json:"relativeTime"`
	SleepDuration                  *int64                                                                   `json:"sleepDuration"`
	Action                         *InstancePreprocessReportInstanceHealthRecipeStepsActionEnum             `json:"action"`
	Status                         *InstancePreprocessReportInstanceHealthRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                                  `json:"errorSpace"`
	P4ServiceAccount               *string                                                                  `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                                   `json:"resourceMetadataSize"`
	Description                    *string                                                                  `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                                 `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                                  `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                                  `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstancePreprocessReportInstanceHealthRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReportInstanceHealthRecipeSteps *InstancePreprocessReportInstanceHealthRecipeSteps = &InstancePreprocessReportInstanceHealthRecipeSteps{empty: true}

func (r *InstancePreprocessReportInstanceHealthRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReportInstanceHealthRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReportInstanceHealthRecipeStepsStatus struct {
	empty   bool                                                             `json:"-"`
	Code    *int64                                                           `json:"code"`
	Message *string                                                          `json:"message"`
	Details []InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstancePreprocessReportInstanceHealthRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReportInstanceHealthRecipeStepsStatus *InstancePreprocessReportInstanceHealthRecipeStepsStatus = &InstancePreprocessReportInstanceHealthRecipeStepsStatus{empty: true}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReportInstanceHealthRecipeStepsStatusDetails *InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails = &InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails{empty: true}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas *InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas = &InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate *InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate = &InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate{empty: true}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject struct {
	empty  bool                                                                              `json:"-"`
	Tag    *string                                                                           `json:"tag"`
	Folder *string                                                                           `json:"folder"`
	Scope  *InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject *InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject = &InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject{empty: true}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo struct {
	empty          bool                                                                             `json:"-"`
	PolicyName     *InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                                          `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                                                    `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo *InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo = &InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo{empty: true}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName *InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName = &InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions *InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions = &InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                                         `json:"-"`
	KeyNotificationsInfo *InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate *InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate = &InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                                                `json:"-"`
	DataVersion            *int64                                                                                                              `json:"dataVersion"`
	Delegate               *string                                                                                                             `json:"delegate"`
	KeyNotificationConfigs []InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReportInstanceHealthRecipe struct {
	empty                             bool                                      `json:"-"`
	Steps                             []InstanceReportInstanceHealthRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                                     `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                                    `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                                  `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                                     `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                                   `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                                  `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                                    `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceReportInstanceHealthRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReportInstanceHealthRecipe *InstanceReportInstanceHealthRecipe = &InstanceReportInstanceHealthRecipe{empty: true}

func (r *InstanceReportInstanceHealthRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReportInstanceHealthRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReportInstanceHealthRecipeSteps struct {
	empty                          bool                                                           `json:"-"`
	RelativeTime                   *int64                                                         `json:"relativeTime"`
	SleepDuration                  *int64                                                         `json:"sleepDuration"`
	Action                         *InstanceReportInstanceHealthRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceReportInstanceHealthRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                        `json:"errorSpace"`
	P4ServiceAccount               *string                                                        `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                         `json:"resourceMetadataSize"`
	Description                    *string                                                        `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                       `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceReportInstanceHealthRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                        `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceReportInstanceHealthRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceReportInstanceHealthRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                        `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceReportInstanceHealthRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReportInstanceHealthRecipeSteps *InstanceReportInstanceHealthRecipeSteps = &InstanceReportInstanceHealthRecipeSteps{empty: true}

func (r *InstanceReportInstanceHealthRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReportInstanceHealthRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReportInstanceHealthRecipeStepsStatus struct {
	empty   bool                                                   `json:"-"`
	Code    *int64                                                 `json:"code"`
	Message *string                                                `json:"message"`
	Details []InstanceReportInstanceHealthRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceReportInstanceHealthRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReportInstanceHealthRecipeStepsStatus *InstanceReportInstanceHealthRecipeStepsStatus = &InstanceReportInstanceHealthRecipeStepsStatus{empty: true}

func (r *InstanceReportInstanceHealthRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReportInstanceHealthRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReportInstanceHealthRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceReportInstanceHealthRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReportInstanceHealthRecipeStepsStatusDetails *InstanceReportInstanceHealthRecipeStepsStatusDetails = &InstanceReportInstanceHealthRecipeStepsStatusDetails{empty: true}

func (r *InstanceReportInstanceHealthRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReportInstanceHealthRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas *InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas = &InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReportInstanceHealthRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReportInstanceHealthRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceReportInstanceHealthRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReportInstanceHealthRecipeStepsPreprocessUpdate *InstanceReportInstanceHealthRecipeStepsPreprocessUpdate = &InstanceReportInstanceHealthRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceReportInstanceHealthRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReportInstanceHealthRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReportInstanceHealthRecipeStepsRequestedTenantProject struct {
	empty  bool                                                                    `json:"-"`
	Tag    *string                                                                 `json:"tag"`
	Folder *string                                                                 `json:"folder"`
	Scope  *InstanceReportInstanceHealthRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceReportInstanceHealthRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReportInstanceHealthRecipeStepsRequestedTenantProject *InstanceReportInstanceHealthRecipeStepsRequestedTenantProject = &InstanceReportInstanceHealthRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceReportInstanceHealthRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReportInstanceHealthRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReportInstanceHealthRecipeStepsPermissionsInfo struct {
	empty          bool                                                                   `json:"-"`
	PolicyName     *InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                                `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                                          `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceReportInstanceHealthRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReportInstanceHealthRecipeStepsPermissionsInfo *InstanceReportInstanceHealthRecipeStepsPermissionsInfo = &InstanceReportInstanceHealthRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceReportInstanceHealthRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReportInstanceHealthRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName *InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName = &InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReportInstanceHealthRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions *InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions = &InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReportInstanceHealthRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                               `json:"-"`
	KeyNotificationsInfo *InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate *InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate = &InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                                      `json:"-"`
	DataVersion            *int64                                                                                                    `json:"dataVersion"`
	Delegate               *string                                                                                                   `json:"delegate"`
	KeyNotificationConfigs []InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReportInstanceHealthRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessGetRecipe struct {
	empty                             bool                               `json:"-"`
	Steps                             []InstancePreprocessGetRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                              `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                             `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                           `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                              `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                            `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                           `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                             `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstancePreprocessGetRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessGetRecipe *InstancePreprocessGetRecipe = &InstancePreprocessGetRecipe{empty: true}

func (r *InstancePreprocessGetRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessGetRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessGetRecipeSteps struct {
	empty                          bool                                                    `json:"-"`
	RelativeTime                   *int64                                                  `json:"relativeTime"`
	SleepDuration                  *int64                                                  `json:"sleepDuration"`
	Action                         *InstancePreprocessGetRecipeStepsActionEnum             `json:"action"`
	Status                         *InstancePreprocessGetRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                 `json:"errorSpace"`
	P4ServiceAccount               *string                                                 `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                  `json:"resourceMetadataSize"`
	Description                    *string                                                 `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstancePreprocessGetRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstancePreprocessGetRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                 `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstancePreprocessGetRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstancePreprocessGetRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstancePreprocessGetRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                 `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstancePreprocessGetRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessGetRecipeSteps *InstancePreprocessGetRecipeSteps = &InstancePreprocessGetRecipeSteps{empty: true}

func (r *InstancePreprocessGetRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessGetRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessGetRecipeStepsStatus struct {
	empty   bool                                            `json:"-"`
	Code    *int64                                          `json:"code"`
	Message *string                                         `json:"message"`
	Details []InstancePreprocessGetRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstancePreprocessGetRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessGetRecipeStepsStatus *InstancePreprocessGetRecipeStepsStatus = &InstancePreprocessGetRecipeStepsStatus{empty: true}

func (r *InstancePreprocessGetRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessGetRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessGetRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstancePreprocessGetRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessGetRecipeStepsStatusDetails *InstancePreprocessGetRecipeStepsStatusDetails = &InstancePreprocessGetRecipeStepsStatusDetails{empty: true}

func (r *InstancePreprocessGetRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessGetRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessGetRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstancePreprocessGetRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessGetRecipeStepsQuotaRequestDeltas *InstancePreprocessGetRecipeStepsQuotaRequestDeltas = &InstancePreprocessGetRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstancePreprocessGetRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessGetRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessGetRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstancePreprocessGetRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessGetRecipeStepsPreprocessUpdate *InstancePreprocessGetRecipeStepsPreprocessUpdate = &InstancePreprocessGetRecipeStepsPreprocessUpdate{empty: true}

func (r *InstancePreprocessGetRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessGetRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessGetRecipeStepsRequestedTenantProject struct {
	empty  bool                                                             `json:"-"`
	Tag    *string                                                          `json:"tag"`
	Folder *string                                                          `json:"folder"`
	Scope  *InstancePreprocessGetRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstancePreprocessGetRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessGetRecipeStepsRequestedTenantProject *InstancePreprocessGetRecipeStepsRequestedTenantProject = &InstancePreprocessGetRecipeStepsRequestedTenantProject{empty: true}

func (r *InstancePreprocessGetRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessGetRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessGetRecipeStepsPermissionsInfo struct {
	empty          bool                                                            `json:"-"`
	PolicyName     *InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                         `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                                   `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstancePreprocessGetRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessGetRecipeStepsPermissionsInfo *InstancePreprocessGetRecipeStepsPermissionsInfo = &InstancePreprocessGetRecipeStepsPermissionsInfo{empty: true}

func (r *InstancePreprocessGetRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessGetRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessGetRecipeStepsPermissionsInfoPolicyName *InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName = &InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessGetRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions *InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions = &InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessGetRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessGetRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                        `json:"-"`
	KeyNotificationsInfo *InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstancePreprocessGetRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessGetRecipeStepsKeyNotificationsUpdate *InstancePreprocessGetRecipeStepsKeyNotificationsUpdate = &InstancePreprocessGetRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstancePreprocessGetRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessGetRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                               `json:"-"`
	DataVersion            *int64                                                                                             `json:"dataVersion"`
	Delegate               *string                                                                                            `json:"delegate"`
	KeyNotificationConfigs []InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessGetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyAvailableRecipe struct {
	empty                             bool                                    `json:"-"`
	Steps                             []InstanceNotifyKeyAvailableRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                                   `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                                  `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                                `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                                   `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                                 `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                                `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                                  `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyAvailableRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyAvailableRecipe *InstanceNotifyKeyAvailableRecipe = &InstanceNotifyKeyAvailableRecipe{empty: true}

func (r *InstanceNotifyKeyAvailableRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyAvailableRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyAvailableRecipeSteps struct {
	empty                          bool                                                         `json:"-"`
	RelativeTime                   *int64                                                       `json:"relativeTime"`
	SleepDuration                  *int64                                                       `json:"sleepDuration"`
	Action                         *InstanceNotifyKeyAvailableRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceNotifyKeyAvailableRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                      `json:"errorSpace"`
	P4ServiceAccount               *string                                                      `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                       `json:"resourceMetadataSize"`
	Description                    *string                                                      `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                     `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                      `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                      `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyAvailableRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyAvailableRecipeSteps *InstanceNotifyKeyAvailableRecipeSteps = &InstanceNotifyKeyAvailableRecipeSteps{empty: true}

func (r *InstanceNotifyKeyAvailableRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyAvailableRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyAvailableRecipeStepsStatus struct {
	empty   bool                                                 `json:"-"`
	Code    *int64                                               `json:"code"`
	Message *string                                              `json:"message"`
	Details []InstanceNotifyKeyAvailableRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyAvailableRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyAvailableRecipeStepsStatus *InstanceNotifyKeyAvailableRecipeStepsStatus = &InstanceNotifyKeyAvailableRecipeStepsStatus{empty: true}

func (r *InstanceNotifyKeyAvailableRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyAvailableRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyAvailableRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyAvailableRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyAvailableRecipeStepsStatusDetails *InstanceNotifyKeyAvailableRecipeStepsStatusDetails = &InstanceNotifyKeyAvailableRecipeStepsStatusDetails{empty: true}

func (r *InstanceNotifyKeyAvailableRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyAvailableRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas *InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas = &InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyAvailableRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate *InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate = &InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyAvailableRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject struct {
	empty  bool                                                                  `json:"-"`
	Tag    *string                                                               `json:"tag"`
	Folder *string                                                               `json:"folder"`
	Scope  *InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject *InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject = &InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyAvailableRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo struct {
	empty          bool                                                                 `json:"-"`
	PolicyName     *InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                              `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                                        `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyAvailableRecipeStepsPermissionsInfo *InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo = &InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyAvailableRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName *InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName = &InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions *InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions = &InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyAvailableRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                             `json:"-"`
	KeyNotificationsInfo *InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate *InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate = &InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                                    `json:"-"`
	DataVersion            *int64                                                                                                  `json:"dataVersion"`
	Delegate               *string                                                                                                 `json:"delegate"`
	KeyNotificationConfigs []InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyAvailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyUnavailableRecipe struct {
	empty                             bool                                      `json:"-"`
	Steps                             []InstanceNotifyKeyUnavailableRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                                     `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                                    `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                                  `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                                     `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                                   `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                                  `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                                    `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyUnavailableRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyUnavailableRecipe *InstanceNotifyKeyUnavailableRecipe = &InstanceNotifyKeyUnavailableRecipe{empty: true}

func (r *InstanceNotifyKeyUnavailableRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyUnavailableRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyUnavailableRecipeSteps struct {
	empty                          bool                                                           `json:"-"`
	RelativeTime                   *int64                                                         `json:"relativeTime"`
	SleepDuration                  *int64                                                         `json:"sleepDuration"`
	Action                         *InstanceNotifyKeyUnavailableRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceNotifyKeyUnavailableRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                        `json:"errorSpace"`
	P4ServiceAccount               *string                                                        `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                         `json:"resourceMetadataSize"`
	Description                    *string                                                        `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                       `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                        `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                        `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyUnavailableRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyUnavailableRecipeSteps *InstanceNotifyKeyUnavailableRecipeSteps = &InstanceNotifyKeyUnavailableRecipeSteps{empty: true}

func (r *InstanceNotifyKeyUnavailableRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyUnavailableRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyUnavailableRecipeStepsStatus struct {
	empty   bool                                                   `json:"-"`
	Code    *int64                                                 `json:"code"`
	Message *string                                                `json:"message"`
	Details []InstanceNotifyKeyUnavailableRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyUnavailableRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyUnavailableRecipeStepsStatus *InstanceNotifyKeyUnavailableRecipeStepsStatus = &InstanceNotifyKeyUnavailableRecipeStepsStatus{empty: true}

func (r *InstanceNotifyKeyUnavailableRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyUnavailableRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyUnavailableRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyUnavailableRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyUnavailableRecipeStepsStatusDetails *InstanceNotifyKeyUnavailableRecipeStepsStatusDetails = &InstanceNotifyKeyUnavailableRecipeStepsStatusDetails{empty: true}

func (r *InstanceNotifyKeyUnavailableRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyUnavailableRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas *InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas = &InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyUnavailableRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate *InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate = &InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyUnavailableRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject struct {
	empty  bool                                                                    `json:"-"`
	Tag    *string                                                                 `json:"tag"`
	Folder *string                                                                 `json:"folder"`
	Scope  *InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject *InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject = &InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyUnavailableRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo struct {
	empty          bool                                                                   `json:"-"`
	PolicyName     *InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                                `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                                          `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo *InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo = &InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName *InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName = &InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions *InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions = &InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyUnavailableRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                               `json:"-"`
	KeyNotificationsInfo *InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate *InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate = &InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                                      `json:"-"`
	DataVersion            *int64                                                                                                    `json:"dataVersion"`
	Delegate               *string                                                                                                   `json:"delegate"`
	KeyNotificationConfigs []InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceNotifyKeyUnavailableRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReadonlyRecipe struct {
	empty                             bool                          `json:"-"`
	Steps                             []InstanceReadonlyRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                         `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                        `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                      `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                         `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                       `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                      `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                        `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipe *InstanceReadonlyRecipe = &InstanceReadonlyRecipe{empty: true}

func (r *InstanceReadonlyRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReadonlyRecipeSteps struct {
	empty                          bool                                               `json:"-"`
	RelativeTime                   *int64                                             `json:"relativeTime"`
	SleepDuration                  *int64                                             `json:"sleepDuration"`
	Action                         *InstanceReadonlyRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceReadonlyRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                            `json:"errorSpace"`
	P4ServiceAccount               *string                                            `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                             `json:"resourceMetadataSize"`
	Description                    *string                                            `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                           `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceReadonlyRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceReadonlyRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                            `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceReadonlyRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceReadonlyRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceReadonlyRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                            `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeSteps *InstanceReadonlyRecipeSteps = &InstanceReadonlyRecipeSteps{empty: true}

func (r *InstanceReadonlyRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReadonlyRecipeStepsStatus struct {
	empty   bool                                       `json:"-"`
	Code    *int64                                     `json:"code"`
	Message *string                                    `json:"message"`
	Details []InstanceReadonlyRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeStepsStatus *InstanceReadonlyRecipeStepsStatus = &InstanceReadonlyRecipeStepsStatus{empty: true}

func (r *InstanceReadonlyRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReadonlyRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeStepsStatusDetails *InstanceReadonlyRecipeStepsStatusDetails = &InstanceReadonlyRecipeStepsStatusDetails{empty: true}

func (r *InstanceReadonlyRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReadonlyRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeStepsQuotaRequestDeltas *InstanceReadonlyRecipeStepsQuotaRequestDeltas = &InstanceReadonlyRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceReadonlyRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReadonlyRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeStepsPreprocessUpdate *InstanceReadonlyRecipeStepsPreprocessUpdate = &InstanceReadonlyRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceReadonlyRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReadonlyRecipeStepsRequestedTenantProject struct {
	empty  bool                                                        `json:"-"`
	Tag    *string                                                     `json:"tag"`
	Folder *string                                                     `json:"folder"`
	Scope  *InstanceReadonlyRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeStepsRequestedTenantProject *InstanceReadonlyRecipeStepsRequestedTenantProject = &InstanceReadonlyRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceReadonlyRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReadonlyRecipeStepsPermissionsInfo struct {
	empty          bool                                                       `json:"-"`
	PolicyName     *InstanceReadonlyRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                    `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                              `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeStepsPermissionsInfo *InstanceReadonlyRecipeStepsPermissionsInfo = &InstanceReadonlyRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceReadonlyRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReadonlyRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeStepsPermissionsInfoPolicyName *InstanceReadonlyRecipeStepsPermissionsInfoPolicyName = &InstanceReadonlyRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceReadonlyRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeStepsPermissionsInfoIamPermissions *InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions = &InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReadonlyRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                   `json:"-"`
	KeyNotificationsInfo *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeStepsKeyNotificationsUpdate *InstanceReadonlyRecipeStepsKeyNotificationsUpdate = &InstanceReadonlyRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceReadonlyRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                          `json:"-"`
	DataVersion            *int64                                                                                        `json:"dataVersion"`
	Delegate               *string                                                                                       `json:"delegate"`
	KeyNotificationConfigs []InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReconcileRecipe struct {
	empty                             bool                           `json:"-"`
	Steps                             []InstanceReconcileRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                          `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                         `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                       `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                          `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                        `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                       `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                         `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceReconcileRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReconcileRecipe *InstanceReconcileRecipe = &InstanceReconcileRecipe{empty: true}

func (r *InstanceReconcileRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReconcileRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReconcileRecipeSteps struct {
	empty                          bool                                                `json:"-"`
	RelativeTime                   *int64                                              `json:"relativeTime"`
	SleepDuration                  *int64                                              `json:"sleepDuration"`
	Action                         *InstanceReconcileRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceReconcileRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                             `json:"errorSpace"`
	P4ServiceAccount               *string                                             `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                              `json:"resourceMetadataSize"`
	Description                    *string                                             `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                            `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceReconcileRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceReconcileRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                             `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceReconcileRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceReconcileRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceReconcileRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                             `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceReconcileRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReconcileRecipeSteps *InstanceReconcileRecipeSteps = &InstanceReconcileRecipeSteps{empty: true}

func (r *InstanceReconcileRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReconcileRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReconcileRecipeStepsStatus struct {
	empty   bool                                        `json:"-"`
	Code    *int64                                      `json:"code"`
	Message *string                                     `json:"message"`
	Details []InstanceReconcileRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceReconcileRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReconcileRecipeStepsStatus *InstanceReconcileRecipeStepsStatus = &InstanceReconcileRecipeStepsStatus{empty: true}

func (r *InstanceReconcileRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReconcileRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReconcileRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceReconcileRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReconcileRecipeStepsStatusDetails *InstanceReconcileRecipeStepsStatusDetails = &InstanceReconcileRecipeStepsStatusDetails{empty: true}

func (r *InstanceReconcileRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReconcileRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReconcileRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceReconcileRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReconcileRecipeStepsQuotaRequestDeltas *InstanceReconcileRecipeStepsQuotaRequestDeltas = &InstanceReconcileRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceReconcileRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReconcileRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReconcileRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceReconcileRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReconcileRecipeStepsPreprocessUpdate *InstanceReconcileRecipeStepsPreprocessUpdate = &InstanceReconcileRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceReconcileRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReconcileRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReconcileRecipeStepsRequestedTenantProject struct {
	empty  bool                                                         `json:"-"`
	Tag    *string                                                      `json:"tag"`
	Folder *string                                                      `json:"folder"`
	Scope  *InstanceReconcileRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceReconcileRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReconcileRecipeStepsRequestedTenantProject *InstanceReconcileRecipeStepsRequestedTenantProject = &InstanceReconcileRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceReconcileRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReconcileRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReconcileRecipeStepsPermissionsInfo struct {
	empty          bool                                                        `json:"-"`
	PolicyName     *InstanceReconcileRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceReconcileRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                     `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                               `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceReconcileRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReconcileRecipeStepsPermissionsInfo *InstanceReconcileRecipeStepsPermissionsInfo = &InstanceReconcileRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceReconcileRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReconcileRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReconcileRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceReconcileRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReconcileRecipeStepsPermissionsInfoPolicyName *InstanceReconcileRecipeStepsPermissionsInfoPolicyName = &InstanceReconcileRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceReconcileRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReconcileRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReconcileRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceReconcileRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReconcileRecipeStepsPermissionsInfoIamPermissions *InstanceReconcileRecipeStepsPermissionsInfoIamPermissions = &InstanceReconcileRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceReconcileRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReconcileRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReconcileRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                    `json:"-"`
	KeyNotificationsInfo *InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceReconcileRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReconcileRecipeStepsKeyNotificationsUpdate *InstanceReconcileRecipeStepsKeyNotificationsUpdate = &InstanceReconcileRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceReconcileRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReconcileRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                           `json:"-"`
	DataVersion            *int64                                                                                         `json:"dataVersion"`
	Delegate               *string                                                                                        `json:"delegate"`
	KeyNotificationConfigs []InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessPassthroughRecipe struct {
	empty                             bool                                       `json:"-"`
	Steps                             []InstancePreprocessPassthroughRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                                      `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                                     `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                                   `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                                      `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                                    `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                                   `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                                     `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstancePreprocessPassthroughRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessPassthroughRecipe *InstancePreprocessPassthroughRecipe = &InstancePreprocessPassthroughRecipe{empty: true}

func (r *InstancePreprocessPassthroughRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessPassthroughRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessPassthroughRecipeSteps struct {
	empty                          bool                                                            `json:"-"`
	RelativeTime                   *int64                                                          `json:"relativeTime"`
	SleepDuration                  *int64                                                          `json:"sleepDuration"`
	Action                         *InstancePreprocessPassthroughRecipeStepsActionEnum             `json:"action"`
	Status                         *InstancePreprocessPassthroughRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                         `json:"errorSpace"`
	P4ServiceAccount               *string                                                         `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                          `json:"resourceMetadataSize"`
	Description                    *string                                                         `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                        `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstancePreprocessPassthroughRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                         `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstancePreprocessPassthroughRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstancePreprocessPassthroughRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                         `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstancePreprocessPassthroughRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessPassthroughRecipeSteps *InstancePreprocessPassthroughRecipeSteps = &InstancePreprocessPassthroughRecipeSteps{empty: true}

func (r *InstancePreprocessPassthroughRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessPassthroughRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessPassthroughRecipeStepsStatus struct {
	empty   bool                                                    `json:"-"`
	Code    *int64                                                  `json:"code"`
	Message *string                                                 `json:"message"`
	Details []InstancePreprocessPassthroughRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstancePreprocessPassthroughRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessPassthroughRecipeStepsStatus *InstancePreprocessPassthroughRecipeStepsStatus = &InstancePreprocessPassthroughRecipeStepsStatus{empty: true}

func (r *InstancePreprocessPassthroughRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessPassthroughRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessPassthroughRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstancePreprocessPassthroughRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessPassthroughRecipeStepsStatusDetails *InstancePreprocessPassthroughRecipeStepsStatusDetails = &InstancePreprocessPassthroughRecipeStepsStatusDetails{empty: true}

func (r *InstancePreprocessPassthroughRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessPassthroughRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas *InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas = &InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessPassthroughRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessPassthroughRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstancePreprocessPassthroughRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessPassthroughRecipeStepsPreprocessUpdate *InstancePreprocessPassthroughRecipeStepsPreprocessUpdate = &InstancePreprocessPassthroughRecipeStepsPreprocessUpdate{empty: true}

func (r *InstancePreprocessPassthroughRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessPassthroughRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessPassthroughRecipeStepsRequestedTenantProject struct {
	empty  bool                                                                     `json:"-"`
	Tag    *string                                                                  `json:"tag"`
	Folder *string                                                                  `json:"folder"`
	Scope  *InstancePreprocessPassthroughRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstancePreprocessPassthroughRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessPassthroughRecipeStepsRequestedTenantProject *InstancePreprocessPassthroughRecipeStepsRequestedTenantProject = &InstancePreprocessPassthroughRecipeStepsRequestedTenantProject{empty: true}

func (r *InstancePreprocessPassthroughRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessPassthroughRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessPassthroughRecipeStepsPermissionsInfo struct {
	empty          bool                                                                    `json:"-"`
	PolicyName     *InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                                 `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                                           `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstancePreprocessPassthroughRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessPassthroughRecipeStepsPermissionsInfo *InstancePreprocessPassthroughRecipeStepsPermissionsInfo = &InstancePreprocessPassthroughRecipeStepsPermissionsInfo{empty: true}

func (r *InstancePreprocessPassthroughRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessPassthroughRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName *InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName = &InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessPassthroughRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions *InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions = &InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessPassthroughRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                                `json:"-"`
	KeyNotificationsInfo *InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate *InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate = &InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                                       `json:"-"`
	DataVersion            *int64                                                                                                     `json:"dataVersion"`
	Delegate               *string                                                                                                    `json:"delegate"`
	KeyNotificationConfigs []InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessPassthroughRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReconcileRecipe struct {
	empty                             bool                                     `json:"-"`
	Steps                             []InstancePreprocessReconcileRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                                    `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                                   `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                                 `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                                    `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                                  `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                                 `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                                   `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstancePreprocessReconcileRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReconcileRecipe *InstancePreprocessReconcileRecipe = &InstancePreprocessReconcileRecipe{empty: true}

func (r *InstancePreprocessReconcileRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReconcileRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReconcileRecipeSteps struct {
	empty                          bool                                                          `json:"-"`
	RelativeTime                   *int64                                                        `json:"relativeTime"`
	SleepDuration                  *int64                                                        `json:"sleepDuration"`
	Action                         *InstancePreprocessReconcileRecipeStepsActionEnum             `json:"action"`
	Status                         *InstancePreprocessReconcileRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                       `json:"errorSpace"`
	P4ServiceAccount               *string                                                       `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                        `json:"resourceMetadataSize"`
	Description                    *string                                                       `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                      `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstancePreprocessReconcileRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                       `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstancePreprocessReconcileRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstancePreprocessReconcileRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                       `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstancePreprocessReconcileRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReconcileRecipeSteps *InstancePreprocessReconcileRecipeSteps = &InstancePreprocessReconcileRecipeSteps{empty: true}

func (r *InstancePreprocessReconcileRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReconcileRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReconcileRecipeStepsStatus struct {
	empty   bool                                                  `json:"-"`
	Code    *int64                                                `json:"code"`
	Message *string                                               `json:"message"`
	Details []InstancePreprocessReconcileRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstancePreprocessReconcileRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReconcileRecipeStepsStatus *InstancePreprocessReconcileRecipeStepsStatus = &InstancePreprocessReconcileRecipeStepsStatus{empty: true}

func (r *InstancePreprocessReconcileRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReconcileRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReconcileRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstancePreprocessReconcileRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReconcileRecipeStepsStatusDetails *InstancePreprocessReconcileRecipeStepsStatusDetails = &InstancePreprocessReconcileRecipeStepsStatusDetails{empty: true}

func (r *InstancePreprocessReconcileRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReconcileRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReconcileRecipeStepsQuotaRequestDeltas *InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas = &InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReconcileRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReconcileRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstancePreprocessReconcileRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReconcileRecipeStepsPreprocessUpdate *InstancePreprocessReconcileRecipeStepsPreprocessUpdate = &InstancePreprocessReconcileRecipeStepsPreprocessUpdate{empty: true}

func (r *InstancePreprocessReconcileRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReconcileRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReconcileRecipeStepsRequestedTenantProject struct {
	empty  bool                                                                   `json:"-"`
	Tag    *string                                                                `json:"tag"`
	Folder *string                                                                `json:"folder"`
	Scope  *InstancePreprocessReconcileRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstancePreprocessReconcileRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReconcileRecipeStepsRequestedTenantProject *InstancePreprocessReconcileRecipeStepsRequestedTenantProject = &InstancePreprocessReconcileRecipeStepsRequestedTenantProject{empty: true}

func (r *InstancePreprocessReconcileRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReconcileRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReconcileRecipeStepsPermissionsInfo struct {
	empty          bool                                                                  `json:"-"`
	PolicyName     *InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                               `json:"resourcePath"`
	ApiAttrs       *InstanceGoogleprotobufstruct                                         `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstancePreprocessReconcileRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReconcileRecipeStepsPermissionsInfo *InstancePreprocessReconcileRecipeStepsPermissionsInfo = &InstancePreprocessReconcileRecipeStepsPermissionsInfo{empty: true}

func (r *InstancePreprocessReconcileRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReconcileRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName *InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName = &InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReconcileRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions *InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions = &InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReconcileRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                              `json:"-"`
	KeyNotificationsInfo *InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate *InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate = &InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty                  bool                                                                                                     `json:"-"`
	DataVersion            *int64                                                                                                   `json:"dataVersion"`
	Delegate               *string                                                                                                  `json:"delegate"`
	KeyNotificationConfigs []InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs `json:"keyNotificationConfigs"`
}

// This object is used to assert a desired state where this InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
	Grant            *string `json:"grant"`
	DelegatorGaiaId  *int64  `json:"delegatorGaiaId"`
}

// This object is used to assert a desired state where this InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs *InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs = &InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs{empty: true}

func (r *InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessReconcileRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyNotificationConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceHistory struct {
	empty               bool    `json:"-"`
	Timestamp           *string `json:"timestamp"`
	OperationHandle     *string `json:"operationHandle"`
	Description         *string `json:"description"`
	StepIndex           *int64  `json:"stepIndex"`
	TenantProjectNumber *int64  `json:"tenantProjectNumber"`
	TenantProjectId     *string `json:"tenantProjectId"`
	P4ServiceAccount    *string `json:"p4ServiceAccount"`
}

// This object is used to assert a desired state where this InstanceHistory is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceHistory *InstanceHistory = &InstanceHistory{empty: true}

func (r *InstanceHistory) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceHistory) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Instance) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "tier2",
		Type:    "Instance",
		Version: "alpha",
	}
}

const InstanceMaxPage = -1

type InstanceList struct {
	Items []*Instance

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *InstanceList) HasNext() bool {
	return l.nextToken != ""
}

func (l *InstanceList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listInstance(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListInstance(ctx context.Context, project, location string) (*InstanceList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListInstanceWithMaxResults(ctx, project, location, InstanceMaxPage)

}

func (c *Client) ListInstanceWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*InstanceList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listInstance(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &InstanceList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetInstance(ctx context.Context, r *Instance) (*Instance, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getInstanceRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalInstance(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeInstanceNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteInstance(ctx context.Context, r *Instance) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("Instance resource is nil")
	}
	c.Config.Logger.Info("Deleting Instance...")
	deleteOp := deleteInstanceOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllInstance deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllInstance(ctx context.Context, project, location string, filter func(*Instance) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListInstance(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllInstance(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllInstance(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyInstance(ctx context.Context, rawDesired *Instance, opts ...dcl.ApplyOption) (*Instance, error) {
	c.Config.Logger.Info("Beginning ApplyInstance...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.instanceDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	var recreate bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				if dcl.HasLifecycleParam(lp, dcl.BlockDestruction) || dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
					return nil, dcl.ApplyInfeasibleError{
						Message: fmt.Sprintf("Infeasible update: (%v) would require recreation.", d),
					}
				}
				c.Config.Logger.Infof("Diff requires recreate: %+v\n", d)
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []instanceApiOperation
	if create {
		ops = append(ops, &createInstanceOperation{})
	} else if recreate {

		ops = append(ops, &deleteInstanceOperation{})

		ops = append(ops, &createInstanceOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeInstanceDesiredState(rawDesired, nil)
		if err != nil {
			return nil, err
		}
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.Infof("Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.Infof("Performing operation %#v", op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.Infof("Failed operation %#v: %v", op, err)
			return nil, err
		}
		c.Config.Logger.Infof("Finished operation %#v", op)
	}

	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.Info("Retrieving raw new state...")
	rawNew, err := c.GetInstance(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createInstanceOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapInstance(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeInstanceNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeInstanceNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeInstanceDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffInstance(c, newDesired, newState)
	if err != nil {
		return nil, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.Info("No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.Infof("Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.Info("Done Apply.")
	return newState, nil
}
