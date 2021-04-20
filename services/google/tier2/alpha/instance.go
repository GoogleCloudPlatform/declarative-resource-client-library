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
	Name                       *string                           `json:"name"`
	DisplayName                *string                           `json:"displayName"`
	Labels                     map[string]string                 `json:"labels"`
	Zone                       *string                           `json:"zone"`
	Sku                        *InstanceSku                      `json:"sku"`
	AuthorizedNetworkId        *string                           `json:"authorizedNetworkId"`
	ReservedIPRange            *string                           `json:"reservedIPRange"`
	HostName                   *string                           `json:"hostName"`
	PortNumber                 *int64                            `json:"portNumber"`
	CurrentZone                *string                           `json:"currentZone"`
	CreationTime               *string                           `json:"creationTime"`
	State                      *InstanceStateEnum                `json:"state"`
	StatusMessage              *string                           `json:"statusMessage"`
	ExtraField                 *string                           `json:"extraField"`
	PreprocessCreateRecipe     *InstancePreprocessCreateRecipe   `json:"preprocessCreateRecipe"`
	InitiateCreateRecipe       *InstanceInitiateCreateRecipe     `json:"initiateCreateRecipe"`
	CreateRecipe               *InstanceCreateRecipe             `json:"createRecipe"`
	DeleteRecipe               *InstanceDeleteRecipe             `json:"deleteRecipe"`
	UpdateRecipe               *InstanceUpdateRecipe             `json:"updateRecipe"`
	PreprocessResetRecipe      *InstancePreprocessResetRecipe    `json:"preprocessResetRecipe"`
	InitiateResetRecipe        *InstanceInitiateResetRecipe      `json:"initiateResetRecipe"`
	ResetRecipe                *InstanceResetRecipe              `json:"resetRecipe"`
	PreprocessRepairRecipe     *InstancePreprocessRepairRecipe   `json:"preprocessRepairRecipe"`
	InitiateRepairRecipe       *InstanceInitiateRepairRecipe     `json:"initiateRepairRecipe"`
	RepairRecipe               *InstanceRepairRecipe             `json:"repairRecipe"`
	PreprocessDeleteRecipe     *InstancePreprocessDeleteRecipe   `json:"preprocessDeleteRecipe"`
	InitiateDeleteRecipe       *InstanceInitiateDeleteRecipe     `json:"initiateDeleteRecipe"`
	PreprocessUpdateRecipe     *InstancePreprocessUpdateRecipe   `json:"preprocessUpdateRecipe"`
	InitiateUpdateRecipe       *InstanceInitiateUpdateRecipe     `json:"initiateUpdateRecipe"`
	PreprocessFreezeRecipe     *InstancePreprocessFreezeRecipe   `json:"preprocessFreezeRecipe"`
	FreezeRecipe               *InstanceFreezeRecipe             `json:"freezeRecipe"`
	PreprocessUnfreezeRecipe   *InstancePreprocessUnfreezeRecipe `json:"preprocessUnfreezeRecipe"`
	UnfreezeRecipe             *InstanceUnfreezeRecipe           `json:"unfreezeRecipe"`
	ReadonlyRecipe             *InstanceReadonlyRecipe           `json:"readonlyRecipe"`
	EnableCallHistory          *bool                             `json:"enableCallHistory"`
	History                    []InstanceHistory                 `json:"history"`
	PublicResourceViewOverride *string                           `json:"publicResourceViewOverride"`
	Project                    *string                           `json:"project"`
	Location                   *string                           `json:"location"`
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

// The enum InstanceInitiateCreateRecipeStepsActionEnum.
type InstanceInitiateCreateRecipeStepsActionEnum string

// InstanceInitiateCreateRecipeStepsActionEnumRef returns a *InstanceInitiateCreateRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceInitiateCreateRecipeStepsActionEnumRef(s string) *InstanceInitiateCreateRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceInitiateCreateRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceInitiateCreateRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceInitiateCreateRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum",
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

// The enum InstanceInitiateResetRecipeStepsActionEnum.
type InstanceInitiateResetRecipeStepsActionEnum string

// InstanceInitiateResetRecipeStepsActionEnumRef returns a *InstanceInitiateResetRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceInitiateResetRecipeStepsActionEnumRef(s string) *InstanceInitiateResetRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceInitiateResetRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceInitiateResetRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceInitiateResetRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum",
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

// The enum InstanceInitiateRepairRecipeStepsActionEnum.
type InstanceInitiateRepairRecipeStepsActionEnum string

// InstanceInitiateRepairRecipeStepsActionEnumRef returns a *InstanceInitiateRepairRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceInitiateRepairRecipeStepsActionEnumRef(s string) *InstanceInitiateRepairRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceInitiateRepairRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceInitiateRepairRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceInitiateRepairRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum",
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

// The enum InstanceInitiateDeleteRecipeStepsActionEnum.
type InstanceInitiateDeleteRecipeStepsActionEnum string

// InstanceInitiateDeleteRecipeStepsActionEnumRef returns a *InstanceInitiateDeleteRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceInitiateDeleteRecipeStepsActionEnumRef(s string) *InstanceInitiateDeleteRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceInitiateDeleteRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceInitiateDeleteRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceInitiateDeleteRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum",
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

// The enum InstanceInitiateUpdateRecipeStepsActionEnum.
type InstanceInitiateUpdateRecipeStepsActionEnum string

// InstanceInitiateUpdateRecipeStepsActionEnumRef returns a *InstanceInitiateUpdateRecipeStepsActionEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceInitiateUpdateRecipeStepsActionEnumRef(s string) *InstanceInitiateUpdateRecipeStepsActionEnum {
	if s == "" {
		return nil
	}

	v := InstanceInitiateUpdateRecipeStepsActionEnum(s)
	return &v
}

func (v InstanceInitiateUpdateRecipeStepsActionEnum) Validate() error {
	for _, s := range []string{"NO_ACTION", "ALLOW", "ALLOW_WITH_LOG", "DENY", "DENY_WITH_LOG", "LOG"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceInitiateUpdateRecipeStepsActionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum.
type InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum string

// InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnumRef returns a *InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnumRef(s string) *InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum {
	if s == "" {
		return nil
	}

	v := InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum(s)
	return &v
}

func (v InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum) Validate() error {
	for _, s := range []string{"UNKNOWN_SCOPE", "PROJECT", "RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum",
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
	ApiAttrs       *InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs *InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs = &InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                                     `json:"-"`
	KeyConfigs  *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                                   `json:"dataVersion"`
	Delegate    *string                                                                                  `json:"delegate"`
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

type InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                              `json:"-"`
	KeyConfig *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipe struct {
	empty                             bool                                `json:"-"`
	Steps                             []InstanceInitiateCreateRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                               `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                              `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                            `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                               `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                             `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                            `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                              `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipe *InstanceInitiateCreateRecipe = &InstanceInitiateCreateRecipe{empty: true}

func (r *InstanceInitiateCreateRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipeSteps struct {
	empty                          bool                                                     `json:"-"`
	RelativeTime                   *int64                                                   `json:"relativeTime"`
	SleepDuration                  *int64                                                   `json:"sleepDuration"`
	Action                         *InstanceInitiateCreateRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceInitiateCreateRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                  `json:"errorSpace"`
	P4ServiceAccount               *string                                                  `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                   `json:"resourceMetadataSize"`
	Description                    *string                                                  `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                 `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceInitiateCreateRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceInitiateCreateRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                  `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceInitiateCreateRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceInitiateCreateRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                  `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipeSteps *InstanceInitiateCreateRecipeSteps = &InstanceInitiateCreateRecipeSteps{empty: true}

func (r *InstanceInitiateCreateRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipeStepsStatus struct {
	empty   bool                                             `json:"-"`
	Code    *int64                                           `json:"code"`
	Message *string                                          `json:"message"`
	Details []InstanceInitiateCreateRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipeStepsStatus *InstanceInitiateCreateRecipeStepsStatus = &InstanceInitiateCreateRecipeStepsStatus{empty: true}

func (r *InstanceInitiateCreateRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipeStepsStatusDetails *InstanceInitiateCreateRecipeStepsStatusDetails = &InstanceInitiateCreateRecipeStepsStatusDetails{empty: true}

func (r *InstanceInitiateCreateRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipeStepsQuotaRequestDeltas *InstanceInitiateCreateRecipeStepsQuotaRequestDeltas = &InstanceInitiateCreateRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceInitiateCreateRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipeStepsPreprocessUpdate *InstanceInitiateCreateRecipeStepsPreprocessUpdate = &InstanceInitiateCreateRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceInitiateCreateRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipeStepsRequestedTenantProject struct {
	empty  bool                                                              `json:"-"`
	Tag    *string                                                           `json:"tag"`
	Folder *string                                                           `json:"folder"`
	Scope  *InstanceInitiateCreateRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipeStepsRequestedTenantProject *InstanceInitiateCreateRecipeStepsRequestedTenantProject = &InstanceInitiateCreateRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceInitiateCreateRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipeStepsPermissionsInfo struct {
	empty          bool                                                             `json:"-"`
	PolicyName     *InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                          `json:"resourcePath"`
	ApiAttrs       *InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipeStepsPermissionsInfo *InstanceInitiateCreateRecipeStepsPermissionsInfo = &InstanceInitiateCreateRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceInitiateCreateRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName *InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName = &InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions *InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions = &InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs *InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs = &InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                         `json:"-"`
	KeyNotificationsInfo *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipeStepsKeyNotificationsUpdate *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate = &InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty       bool                                                                                   `json:"-"`
	KeyConfigs  *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                                 `json:"dataVersion"`
	Delegate    *string                                                                                `json:"delegate"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                            `json:"-"`
	KeyConfig *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
	ApiAttrs       *InstanceCreateRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstanceCreateRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeStepsPermissionsInfoApiAttrs *InstanceCreateRecipeStepsPermissionsInfoApiAttrs = &InstanceCreateRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstanceCreateRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                           `json:"-"`
	KeyConfigs  *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                         `json:"dataVersion"`
	Delegate    *string                                                                        `json:"delegate"`
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

type InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                    `json:"-"`
	KeyConfig *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceCreateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
	ApiAttrs       *InstanceDeleteRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstanceDeleteRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeStepsPermissionsInfoApiAttrs *InstanceDeleteRecipeStepsPermissionsInfoApiAttrs = &InstanceDeleteRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstanceDeleteRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                           `json:"-"`
	KeyConfigs  *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                         `json:"dataVersion"`
	Delegate    *string                                                                        `json:"delegate"`
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

type InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                    `json:"-"`
	KeyConfig *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
	ApiAttrs       *InstanceUpdateRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstanceUpdateRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeStepsPermissionsInfoApiAttrs *InstanceUpdateRecipeStepsPermissionsInfoApiAttrs = &InstanceUpdateRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstanceUpdateRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                           `json:"-"`
	KeyConfigs  *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                         `json:"dataVersion"`
	Delegate    *string                                                                        `json:"delegate"`
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

type InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                    `json:"-"`
	KeyConfig *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
	ApiAttrs       *InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs *InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs = &InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                                    `json:"-"`
	KeyConfigs  *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                                  `json:"dataVersion"`
	Delegate    *string                                                                                 `json:"delegate"`
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

type InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                             `json:"-"`
	KeyConfig *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipe struct {
	empty                             bool                               `json:"-"`
	Steps                             []InstanceInitiateResetRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                              `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                             `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                           `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                              `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                            `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                           `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                             `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipe *InstanceInitiateResetRecipe = &InstanceInitiateResetRecipe{empty: true}

func (r *InstanceInitiateResetRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipeSteps struct {
	empty                          bool                                                    `json:"-"`
	RelativeTime                   *int64                                                  `json:"relativeTime"`
	SleepDuration                  *int64                                                  `json:"sleepDuration"`
	Action                         *InstanceInitiateResetRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceInitiateResetRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                 `json:"errorSpace"`
	P4ServiceAccount               *string                                                 `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                  `json:"resourceMetadataSize"`
	Description                    *string                                                 `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceInitiateResetRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceInitiateResetRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                 `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceInitiateResetRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceInitiateResetRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceInitiateResetRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                 `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipeSteps *InstanceInitiateResetRecipeSteps = &InstanceInitiateResetRecipeSteps{empty: true}

func (r *InstanceInitiateResetRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipeStepsStatus struct {
	empty   bool                                            `json:"-"`
	Code    *int64                                          `json:"code"`
	Message *string                                         `json:"message"`
	Details []InstanceInitiateResetRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipeStepsStatus *InstanceInitiateResetRecipeStepsStatus = &InstanceInitiateResetRecipeStepsStatus{empty: true}

func (r *InstanceInitiateResetRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipeStepsStatusDetails *InstanceInitiateResetRecipeStepsStatusDetails = &InstanceInitiateResetRecipeStepsStatusDetails{empty: true}

func (r *InstanceInitiateResetRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipeStepsQuotaRequestDeltas *InstanceInitiateResetRecipeStepsQuotaRequestDeltas = &InstanceInitiateResetRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceInitiateResetRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipeStepsPreprocessUpdate *InstanceInitiateResetRecipeStepsPreprocessUpdate = &InstanceInitiateResetRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceInitiateResetRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipeStepsRequestedTenantProject struct {
	empty  bool                                                             `json:"-"`
	Tag    *string                                                          `json:"tag"`
	Folder *string                                                          `json:"folder"`
	Scope  *InstanceInitiateResetRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipeStepsRequestedTenantProject *InstanceInitiateResetRecipeStepsRequestedTenantProject = &InstanceInitiateResetRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceInitiateResetRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipeStepsPermissionsInfo struct {
	empty          bool                                                            `json:"-"`
	PolicyName     *InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                         `json:"resourcePath"`
	ApiAttrs       *InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipeStepsPermissionsInfo *InstanceInitiateResetRecipeStepsPermissionsInfo = &InstanceInitiateResetRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceInitiateResetRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipeStepsPermissionsInfoPolicyName *InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName = &InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions *InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions = &InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs *InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs = &InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                        `json:"-"`
	KeyNotificationsInfo *InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipeStepsKeyNotificationsUpdate *InstanceInitiateResetRecipeStepsKeyNotificationsUpdate = &InstanceInitiateResetRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceInitiateResetRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty       bool                                                                                  `json:"-"`
	KeyConfigs  *InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                                `json:"dataVersion"`
	Delegate    *string                                                                               `json:"delegate"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                           `json:"-"`
	KeyConfig *InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
	ApiAttrs       *InstanceResetRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstanceResetRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstanceResetRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeStepsPermissionsInfoApiAttrs *InstanceResetRecipeStepsPermissionsInfoApiAttrs = &InstanceResetRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstanceResetRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                          `json:"-"`
	KeyConfigs  *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                        `json:"dataVersion"`
	Delegate    *string                                                                       `json:"delegate"`
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

type InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                   `json:"-"`
	KeyConfig *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceResetRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
	ApiAttrs       *InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs *InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs = &InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                                     `json:"-"`
	KeyConfigs  *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                                   `json:"dataVersion"`
	Delegate    *string                                                                                  `json:"delegate"`
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

type InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                              `json:"-"`
	KeyConfig *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipe struct {
	empty                             bool                                `json:"-"`
	Steps                             []InstanceInitiateRepairRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                               `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                              `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                            `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                               `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                             `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                            `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                              `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipe *InstanceInitiateRepairRecipe = &InstanceInitiateRepairRecipe{empty: true}

func (r *InstanceInitiateRepairRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipeSteps struct {
	empty                          bool                                                     `json:"-"`
	RelativeTime                   *int64                                                   `json:"relativeTime"`
	SleepDuration                  *int64                                                   `json:"sleepDuration"`
	Action                         *InstanceInitiateRepairRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceInitiateRepairRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                  `json:"errorSpace"`
	P4ServiceAccount               *string                                                  `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                   `json:"resourceMetadataSize"`
	Description                    *string                                                  `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                 `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceInitiateRepairRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceInitiateRepairRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                  `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceInitiateRepairRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceInitiateRepairRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                  `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipeSteps *InstanceInitiateRepairRecipeSteps = &InstanceInitiateRepairRecipeSteps{empty: true}

func (r *InstanceInitiateRepairRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipeStepsStatus struct {
	empty   bool                                             `json:"-"`
	Code    *int64                                           `json:"code"`
	Message *string                                          `json:"message"`
	Details []InstanceInitiateRepairRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipeStepsStatus *InstanceInitiateRepairRecipeStepsStatus = &InstanceInitiateRepairRecipeStepsStatus{empty: true}

func (r *InstanceInitiateRepairRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipeStepsStatusDetails *InstanceInitiateRepairRecipeStepsStatusDetails = &InstanceInitiateRepairRecipeStepsStatusDetails{empty: true}

func (r *InstanceInitiateRepairRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipeStepsQuotaRequestDeltas *InstanceInitiateRepairRecipeStepsQuotaRequestDeltas = &InstanceInitiateRepairRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceInitiateRepairRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipeStepsPreprocessUpdate *InstanceInitiateRepairRecipeStepsPreprocessUpdate = &InstanceInitiateRepairRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceInitiateRepairRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipeStepsRequestedTenantProject struct {
	empty  bool                                                              `json:"-"`
	Tag    *string                                                           `json:"tag"`
	Folder *string                                                           `json:"folder"`
	Scope  *InstanceInitiateRepairRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipeStepsRequestedTenantProject *InstanceInitiateRepairRecipeStepsRequestedTenantProject = &InstanceInitiateRepairRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceInitiateRepairRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipeStepsPermissionsInfo struct {
	empty          bool                                                             `json:"-"`
	PolicyName     *InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                          `json:"resourcePath"`
	ApiAttrs       *InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipeStepsPermissionsInfo *InstanceInitiateRepairRecipeStepsPermissionsInfo = &InstanceInitiateRepairRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceInitiateRepairRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName *InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName = &InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions *InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions = &InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs *InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs = &InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                         `json:"-"`
	KeyNotificationsInfo *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipeStepsKeyNotificationsUpdate *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate = &InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty       bool                                                                                   `json:"-"`
	KeyConfigs  *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                                 `json:"dataVersion"`
	Delegate    *string                                                                                `json:"delegate"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                            `json:"-"`
	KeyConfig *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
	ApiAttrs       *InstanceRepairRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstanceRepairRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeStepsPermissionsInfoApiAttrs *InstanceRepairRecipeStepsPermissionsInfoApiAttrs = &InstanceRepairRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstanceRepairRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                           `json:"-"`
	KeyConfigs  *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                         `json:"dataVersion"`
	Delegate    *string                                                                        `json:"delegate"`
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

type InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                    `json:"-"`
	KeyConfig *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceRepairRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
	ApiAttrs       *InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs *InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs = &InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                                     `json:"-"`
	KeyConfigs  *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                                   `json:"dataVersion"`
	Delegate    *string                                                                                  `json:"delegate"`
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

type InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                              `json:"-"`
	KeyConfig *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipe struct {
	empty                             bool                                `json:"-"`
	Steps                             []InstanceInitiateDeleteRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                               `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                              `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                            `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                               `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                             `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                            `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                              `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipe *InstanceInitiateDeleteRecipe = &InstanceInitiateDeleteRecipe{empty: true}

func (r *InstanceInitiateDeleteRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipeSteps struct {
	empty                          bool                                                     `json:"-"`
	RelativeTime                   *int64                                                   `json:"relativeTime"`
	SleepDuration                  *int64                                                   `json:"sleepDuration"`
	Action                         *InstanceInitiateDeleteRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceInitiateDeleteRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                  `json:"errorSpace"`
	P4ServiceAccount               *string                                                  `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                   `json:"resourceMetadataSize"`
	Description                    *string                                                  `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                 `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceInitiateDeleteRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                  `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceInitiateDeleteRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceInitiateDeleteRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                  `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipeSteps *InstanceInitiateDeleteRecipeSteps = &InstanceInitiateDeleteRecipeSteps{empty: true}

func (r *InstanceInitiateDeleteRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipeStepsStatus struct {
	empty   bool                                             `json:"-"`
	Code    *int64                                           `json:"code"`
	Message *string                                          `json:"message"`
	Details []InstanceInitiateDeleteRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipeStepsStatus *InstanceInitiateDeleteRecipeStepsStatus = &InstanceInitiateDeleteRecipeStepsStatus{empty: true}

func (r *InstanceInitiateDeleteRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipeStepsStatusDetails *InstanceInitiateDeleteRecipeStepsStatusDetails = &InstanceInitiateDeleteRecipeStepsStatusDetails{empty: true}

func (r *InstanceInitiateDeleteRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipeStepsQuotaRequestDeltas *InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas = &InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipeStepsPreprocessUpdate *InstanceInitiateDeleteRecipeStepsPreprocessUpdate = &InstanceInitiateDeleteRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceInitiateDeleteRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipeStepsRequestedTenantProject struct {
	empty  bool                                                              `json:"-"`
	Tag    *string                                                           `json:"tag"`
	Folder *string                                                           `json:"folder"`
	Scope  *InstanceInitiateDeleteRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipeStepsRequestedTenantProject *InstanceInitiateDeleteRecipeStepsRequestedTenantProject = &InstanceInitiateDeleteRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceInitiateDeleteRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipeStepsPermissionsInfo struct {
	empty          bool                                                             `json:"-"`
	PolicyName     *InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                          `json:"resourcePath"`
	ApiAttrs       *InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipeStepsPermissionsInfo *InstanceInitiateDeleteRecipeStepsPermissionsInfo = &InstanceInitiateDeleteRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceInitiateDeleteRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName *InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName = &InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions *InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions = &InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs *InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs = &InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                         `json:"-"`
	KeyNotificationsInfo *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate = &InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty       bool                                                                                   `json:"-"`
	KeyConfigs  *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                                 `json:"dataVersion"`
	Delegate    *string                                                                                `json:"delegate"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                            `json:"-"`
	KeyConfig *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateDeleteRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
	ApiAttrs       *InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs *InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs = &InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                                     `json:"-"`
	KeyConfigs  *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                                   `json:"dataVersion"`
	Delegate    *string                                                                                  `json:"delegate"`
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

type InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                              `json:"-"`
	KeyConfig *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipe struct {
	empty                             bool                                `json:"-"`
	Steps                             []InstanceInitiateUpdateRecipeSteps `json:"steps"`
	HonorCancelRequest                *bool                               `json:"honorCancelRequest"`
	IgnoreRecipeAfter                 *int64                              `json:"ignoreRecipeAfter"`
	VerifyDeadlineSecondsBelow        *float64                            `json:"verifyDeadlineSecondsBelow"`
	PopulateOperationResult           *bool                               `json:"populateOperationResult"`
	ReadonlyRecipeStartTime           *string                             `json:"readonlyRecipeStartTime"`
	ResourceNamesStoredInClhWithDelay []string                            `json:"resourceNamesStoredInClhWithDelay"`
	DelayToStoreResourcesInClhDbNanos *int64                              `json:"delayToStoreResourcesInClhDbNanos"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipe is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipe *InstanceInitiateUpdateRecipe = &InstanceInitiateUpdateRecipe{empty: true}

func (r *InstanceInitiateUpdateRecipe) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipe) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipeSteps struct {
	empty                          bool                                                     `json:"-"`
	RelativeTime                   *int64                                                   `json:"relativeTime"`
	SleepDuration                  *int64                                                   `json:"sleepDuration"`
	Action                         *InstanceInitiateUpdateRecipeStepsActionEnum             `json:"action"`
	Status                         *InstanceInitiateUpdateRecipeStepsStatus                 `json:"status"`
	ErrorSpace                     *string                                                  `json:"errorSpace"`
	P4ServiceAccount               *string                                                  `json:"p4ServiceAccount"`
	ResourceMetadataSize           *int64                                                   `json:"resourceMetadataSize"`
	Description                    *string                                                  `json:"description"`
	UpdatedRepeatOperationDelaySec *float64                                                 `json:"updatedRepeatOperationDelaySec"`
	QuotaRequestDeltas             []InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas    `json:"quotaRequestDeltas"`
	PreprocessUpdate               *InstanceInitiateUpdateRecipeStepsPreprocessUpdate       `json:"preprocessUpdate"`
	PublicOperationMetadata        *string                                                  `json:"publicOperationMetadata"`
	RequestedTenantProject         *InstanceInitiateUpdateRecipeStepsRequestedTenantProject `json:"requestedTenantProject"`
	PermissionsInfo                []InstanceInitiateUpdateRecipeStepsPermissionsInfo       `json:"permissionsInfo"`
	KeyNotificationsUpdate         *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate `json:"keyNotificationsUpdate"`
	ClhDataUpdateTime              *string                                                  `json:"clhDataUpdateTime"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipeSteps is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipeSteps *InstanceInitiateUpdateRecipeSteps = &InstanceInitiateUpdateRecipeSteps{empty: true}

func (r *InstanceInitiateUpdateRecipeSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipeSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipeStepsStatus struct {
	empty   bool                                             `json:"-"`
	Code    *int64                                           `json:"code"`
	Message *string                                          `json:"message"`
	Details []InstanceInitiateUpdateRecipeStepsStatusDetails `json:"details"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipeStepsStatus is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipeStepsStatus *InstanceInitiateUpdateRecipeStepsStatus = &InstanceInitiateUpdateRecipeStepsStatus{empty: true}

func (r *InstanceInitiateUpdateRecipeStepsStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipeStepsStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipeStepsStatusDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipeStepsStatusDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipeStepsStatusDetails *InstanceInitiateUpdateRecipeStepsStatusDetails = &InstanceInitiateUpdateRecipeStepsStatusDetails{empty: true}

func (r *InstanceInitiateUpdateRecipeStepsStatusDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipeStepsStatusDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas struct {
	empty      bool    `json:"-"`
	MetricName *string `json:"metricName"`
	Amount     *int64  `json:"amount"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipeStepsQuotaRequestDeltas *InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas = &InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas{empty: true}

func (r *InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipeStepsQuotaRequestDeltas) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipeStepsPreprocessUpdate struct {
	empty                   bool    `json:"-"`
	LatencySloBucketName    *string `json:"latencySloBucketName"`
	PublicOperationMetadata *string `json:"publicOperationMetadata"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipeStepsPreprocessUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipeStepsPreprocessUpdate *InstanceInitiateUpdateRecipeStepsPreprocessUpdate = &InstanceInitiateUpdateRecipeStepsPreprocessUpdate{empty: true}

func (r *InstanceInitiateUpdateRecipeStepsPreprocessUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipeStepsPreprocessUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipeStepsRequestedTenantProject struct {
	empty  bool                                                              `json:"-"`
	Tag    *string                                                           `json:"tag"`
	Folder *string                                                           `json:"folder"`
	Scope  *InstanceInitiateUpdateRecipeStepsRequestedTenantProjectScopeEnum `json:"scope"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipeStepsRequestedTenantProject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipeStepsRequestedTenantProject *InstanceInitiateUpdateRecipeStepsRequestedTenantProject = &InstanceInitiateUpdateRecipeStepsRequestedTenantProject{empty: true}

func (r *InstanceInitiateUpdateRecipeStepsRequestedTenantProject) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipeStepsRequestedTenantProject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipeStepsPermissionsInfo struct {
	empty          bool                                                             `json:"-"`
	PolicyName     *InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName      `json:"policyName"`
	IamPermissions []InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions `json:"iamPermissions"`
	ResourcePath   *string                                                          `json:"resourcePath"`
	ApiAttrs       *InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipeStepsPermissionsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipeStepsPermissionsInfo *InstanceInitiateUpdateRecipeStepsPermissionsInfo = &InstanceInitiateUpdateRecipeStepsPermissionsInfo{empty: true}

func (r *InstanceInitiateUpdateRecipeStepsPermissionsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipeStepsPermissionsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName struct {
	empty  bool    `json:"-"`
	Type   *string `json:"type"`
	Id     *string `json:"id"`
	Region *string `json:"region"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName *InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName = &InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName{empty: true}

func (r *InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipeStepsPermissionsInfoPolicyName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions struct {
	empty      bool    `json:"-"`
	Permission *string `json:"permission"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions *InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions = &InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions{empty: true}

func (r *InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipeStepsPermissionsInfoIamPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs *InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs = &InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate struct {
	empty                bool                                                                         `json:"-"`
	KeyNotificationsInfo *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo `json:"keyNotificationsInfo"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate = &InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate{empty: true}

func (r *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo struct {
	empty       bool                                                                                   `json:"-"`
	KeyConfigs  *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                                 `json:"dataVersion"`
	Delegate    *string                                                                                `json:"delegate"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo = &InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo{empty: true}

func (r *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                            `json:"-"`
	KeyConfig *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceInitiateUpdateRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
	ApiAttrs       *InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs *InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs = &InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                                     `json:"-"`
	KeyConfigs  *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                                   `json:"dataVersion"`
	Delegate    *string                                                                                  `json:"delegate"`
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

type InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                              `json:"-"`
	KeyConfig *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
	ApiAttrs       *InstanceFreezeRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstanceFreezeRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeStepsPermissionsInfoApiAttrs *InstanceFreezeRecipeStepsPermissionsInfoApiAttrs = &InstanceFreezeRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstanceFreezeRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                           `json:"-"`
	KeyConfigs  *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                         `json:"dataVersion"`
	Delegate    *string                                                                        `json:"delegate"`
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

type InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                    `json:"-"`
	KeyConfig *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceFreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
	ApiAttrs       *InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs *InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs = &InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                                       `json:"-"`
	KeyConfigs  *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                                     `json:"dataVersion"`
	Delegate    *string                                                                                    `json:"delegate"`
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

type InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                                `json:"-"`
	KeyConfig *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstancePreprocessUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
	ApiAttrs       *InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs *InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs = &InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                             `json:"-"`
	KeyConfigs  *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                           `json:"dataVersion"`
	Delegate    *string                                                                          `json:"delegate"`
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

type InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                      `json:"-"`
	KeyConfig *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceUnfreezeRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
	ApiAttrs       *InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs        `json:"apiAttrs"`
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

type InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs struct {
	empty bool `json:"-"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeStepsPermissionsInfoApiAttrs *InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs = &InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs{empty: true}

func (r *InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeStepsPermissionsInfoApiAttrs) HashCode() string {
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
	empty       bool                                                                             `json:"-"`
	KeyConfigs  *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs `json:"keyConfigs"`
	DataVersion *int64                                                                           `json:"dataVersion"`
	Delegate    *string                                                                          `json:"delegate"`
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

type InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs struct {
	empty     bool                                                                                      `json:"-"`
	KeyConfig *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig `json:"keyConfig"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs = &InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs{empty: true}

func (r *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig struct {
	empty            bool    `json:"-"`
	KeyOrVersionName *string `json:"keyOrVersionName"`
}

// This object is used to assert a desired state where this InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyInstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig = &InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig{empty: true}

func (r *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *InstanceReadonlyRecipeStepsKeyNotificationsUpdateKeyNotificationsInfoKeyConfigsKeyConfig) HashCode() string {
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
