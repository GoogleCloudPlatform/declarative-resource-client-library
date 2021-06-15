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
package beta

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type CertificateAuthority struct {
	Name                      *string                                         `json:"name"`
	Type                      *CertificateAuthorityTypeEnum                   `json:"type"`
	Config                    *CertificateAuthorityConfig                     `json:"config"`
	Lifetime                  *string                                         `json:"lifetime"`
	KeySpec                   *CertificateAuthorityKeySpec                    `json:"keySpec"`
	SubordinateConfig         *CertificateAuthoritySubordinateConfig          `json:"subordinateConfig"`
	Tier                      *CertificateAuthorityTierEnum                   `json:"tier"`
	State                     *CertificateAuthorityStateEnum                  `json:"state"`
	PemCaCertificates         []string                                        `json:"pemCaCertificates"`
	CaCertificateDescriptions []CertificateAuthorityCaCertificateDescriptions `json:"caCertificateDescriptions"`
	GcsBucket                 *string                                         `json:"gcsBucket"`
	AccessUrls                *CertificateAuthorityAccessUrls                 `json:"accessUrls"`
	CreateTime                *string                                         `json:"createTime"`
	UpdateTime                *string                                         `json:"updateTime"`
	DeleteTime                *string                                         `json:"deleteTime"`
	Labels                    map[string]string                               `json:"labels"`
	CertificatePolicy         *CertificateAuthorityCertificatePolicy          `json:"certificatePolicy"`
	IssuingOptions            *CertificateAuthorityIssuingOptions             `json:"issuingOptions"`
	Project                   *string                                         `json:"project"`
	Location                  *string                                         `json:"location"`
}

func (r *CertificateAuthority) String() string {
	return dcl.SprintResource(r)
}

// The enum CertificateAuthorityTypeEnum.
type CertificateAuthorityTypeEnum string

// CertificateAuthorityTypeEnumRef returns a *CertificateAuthorityTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func CertificateAuthorityTypeEnumRef(s string) *CertificateAuthorityTypeEnum {
	if s == "" {
		return nil
	}

	v := CertificateAuthorityTypeEnum(s)
	return &v
}

func (v CertificateAuthorityTypeEnum) Validate() error {
	for _, s := range []string{"TYPE_UNSPECIFIED", "SELF_SIGNED", "SUBORDINATE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CertificateAuthorityTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum CertificateAuthorityConfigPublicKeyTypeEnum.
type CertificateAuthorityConfigPublicKeyTypeEnum string

// CertificateAuthorityConfigPublicKeyTypeEnumRef returns a *CertificateAuthorityConfigPublicKeyTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func CertificateAuthorityConfigPublicKeyTypeEnumRef(s string) *CertificateAuthorityConfigPublicKeyTypeEnum {
	if s == "" {
		return nil
	}

	v := CertificateAuthorityConfigPublicKeyTypeEnum(s)
	return &v
}

func (v CertificateAuthorityConfigPublicKeyTypeEnum) Validate() error {
	for _, s := range []string{"KEY_TYPE_UNSPECIFIED", "PEM_RSA_KEY", "PEM_EC_KEY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CertificateAuthorityConfigPublicKeyTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum CertificateAuthorityKeySpecAlgorithmEnum.
type CertificateAuthorityKeySpecAlgorithmEnum string

// CertificateAuthorityKeySpecAlgorithmEnumRef returns a *CertificateAuthorityKeySpecAlgorithmEnum with the value of string s
// If the empty string is provided, nil is returned.
func CertificateAuthorityKeySpecAlgorithmEnumRef(s string) *CertificateAuthorityKeySpecAlgorithmEnum {
	if s == "" {
		return nil
	}

	v := CertificateAuthorityKeySpecAlgorithmEnum(s)
	return &v
}

func (v CertificateAuthorityKeySpecAlgorithmEnum) Validate() error {
	for _, s := range []string{"SIGN_HASH_ALGORITHM_UNSPECIFIED", "RSA_PSS_2048_SHA256", "RSA_PSS_3072_SHA256", "RSA_PSS_4096_SHA256", "RSA_PKCS1_2048_SHA256", "RSA_PKCS1_3072_SHA256", "RSA_PKCS1_4096_SHA256", "EC_P256_SHA256", "EC_P384_SHA384"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CertificateAuthorityKeySpecAlgorithmEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum CertificateAuthorityTierEnum.
type CertificateAuthorityTierEnum string

// CertificateAuthorityTierEnumRef returns a *CertificateAuthorityTierEnum with the value of string s
// If the empty string is provided, nil is returned.
func CertificateAuthorityTierEnumRef(s string) *CertificateAuthorityTierEnum {
	if s == "" {
		return nil
	}

	v := CertificateAuthorityTierEnum(s)
	return &v
}

func (v CertificateAuthorityTierEnum) Validate() error {
	for _, s := range []string{"TIER_UNSPECIFIED", "ENTERPRISE", "DEVOPS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CertificateAuthorityTierEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum CertificateAuthorityStateEnum.
type CertificateAuthorityStateEnum string

// CertificateAuthorityStateEnumRef returns a *CertificateAuthorityStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func CertificateAuthorityStateEnumRef(s string) *CertificateAuthorityStateEnum {
	if s == "" {
		return nil
	}

	v := CertificateAuthorityStateEnum(s)
	return &v
}

func (v CertificateAuthorityStateEnum) Validate() error {
	for _, s := range []string{"STATE_UNSPECIFIED", "ENABLED", "DISABLED", "STAGED", "AWAITING_USER_ACTIVATION", "DELETED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CertificateAuthorityStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum.
type CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum string

// CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumRef returns a *CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum with the value of string s
// If the empty string is provided, nil is returned.
func CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnumRef(s string) *CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum {
	if s == "" {
		return nil
	}

	v := CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum(s)
	return &v
}

func (v CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum) Validate() error {
	for _, s := range []string{"KEY_FORMAT_UNSPECIFIED", "PEM"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum.
type CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum string

// CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnumRef returns a *CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnumRef(s string) *CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum {
	if s == "" {
		return nil
	}

	v := CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum(s)
	return &v
}

func (v CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum) Validate() error {
	for _, s := range []string{"TYPE_UNSPECIFIED", "SELF_SIGNED", "SUBORDINATE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type CertificateAuthorityConfig struct {
	empty          bool                                      `json:"-"`
	SubjectConfig  *CertificateAuthorityConfigSubjectConfig  `json:"subjectConfig"`
	PublicKey      *CertificateAuthorityConfigPublicKey      `json:"publicKey"`
	ReusableConfig *CertificateAuthorityConfigReusableConfig `json:"reusableConfig"`
}

type jsonCertificateAuthorityConfig CertificateAuthorityConfig

func (r *CertificateAuthorityConfig) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfig
	} else {

		r.SubjectConfig = res.SubjectConfig

		r.PublicKey = res.PublicKey

		r.ReusableConfig = res.ReusableConfig

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfig *CertificateAuthorityConfig = &CertificateAuthorityConfig{empty: true}

func (r *CertificateAuthorityConfig) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigSubjectConfig struct {
	empty          bool                                                   `json:"-"`
	Subject        *CertificateAuthorityConfigSubjectConfigSubject        `json:"subject"`
	CommonName     *string                                                `json:"commonName"`
	SubjectAltName *CertificateAuthorityConfigSubjectConfigSubjectAltName `json:"subjectAltName"`
}

type jsonCertificateAuthorityConfigSubjectConfig CertificateAuthorityConfigSubjectConfig

func (r *CertificateAuthorityConfigSubjectConfig) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigSubjectConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigSubjectConfig
	} else {

		r.Subject = res.Subject

		r.CommonName = res.CommonName

		r.SubjectAltName = res.SubjectAltName

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigSubjectConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigSubjectConfig *CertificateAuthorityConfigSubjectConfig = &CertificateAuthorityConfigSubjectConfig{empty: true}

func (r *CertificateAuthorityConfigSubjectConfig) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigSubjectConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigSubjectConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigSubjectConfigSubject struct {
	empty              bool    `json:"-"`
	CountryCode        *string `json:"countryCode"`
	Organization       *string `json:"organization"`
	OrganizationalUnit *string `json:"organizationalUnit"`
	Locality           *string `json:"locality"`
	Province           *string `json:"province"`
	StreetAddress      *string `json:"streetAddress"`
	PostalCode         *string `json:"postalCode"`
}

type jsonCertificateAuthorityConfigSubjectConfigSubject CertificateAuthorityConfigSubjectConfigSubject

func (r *CertificateAuthorityConfigSubjectConfigSubject) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigSubjectConfigSubject
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigSubjectConfigSubject
	} else {

		r.CountryCode = res.CountryCode

		r.Organization = res.Organization

		r.OrganizationalUnit = res.OrganizationalUnit

		r.Locality = res.Locality

		r.Province = res.Province

		r.StreetAddress = res.StreetAddress

		r.PostalCode = res.PostalCode

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigSubjectConfigSubject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigSubjectConfigSubject *CertificateAuthorityConfigSubjectConfigSubject = &CertificateAuthorityConfigSubjectConfigSubject{empty: true}

func (r *CertificateAuthorityConfigSubjectConfigSubject) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigSubjectConfigSubject) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigSubjectConfigSubject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigSubjectConfigSubjectAltName struct {
	empty          bool                                                              `json:"-"`
	DnsNames       []string                                                          `json:"dnsNames"`
	Uris           []string                                                          `json:"uris"`
	EmailAddresses []string                                                          `json:"emailAddresses"`
	IPAddresses    []string                                                          `json:"ipAddresses"`
	CustomSans     []CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans `json:"customSans"`
}

type jsonCertificateAuthorityConfigSubjectConfigSubjectAltName CertificateAuthorityConfigSubjectConfigSubjectAltName

func (r *CertificateAuthorityConfigSubjectConfigSubjectAltName) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigSubjectConfigSubjectAltName
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigSubjectConfigSubjectAltName
	} else {

		r.DnsNames = res.DnsNames

		r.Uris = res.Uris

		r.EmailAddresses = res.EmailAddresses

		r.IPAddresses = res.IPAddresses

		r.CustomSans = res.CustomSans

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigSubjectConfigSubjectAltName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigSubjectConfigSubjectAltName *CertificateAuthorityConfigSubjectConfigSubjectAltName = &CertificateAuthorityConfigSubjectConfigSubjectAltName{empty: true}

func (r *CertificateAuthorityConfigSubjectConfigSubjectAltName) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigSubjectConfigSubjectAltName) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigSubjectConfigSubjectAltName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans struct {
	empty    bool                                                                     `json:"-"`
	ObjectId *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId `json:"objectId"`
	Critical *bool                                                                    `json:"critical"`
	Value    *string                                                                  `json:"value"`
}

type jsonCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans

func (r *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans
	} else {

		r.ObjectId = res.ObjectId

		r.Critical = res.Critical

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans = &CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans{empty: true}

func (r *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSans) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId struct {
	empty        bool    `json:"-"`
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type jsonCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId

func (r *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId
	} else {

		r.ObjectIdPath = res.ObjectIdPath

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId = &CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId{empty: true}

func (r *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigSubjectConfigSubjectAltNameCustomSansObjectId) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigPublicKey struct {
	empty bool                                         `json:"-"`
	Key   *string                                      `json:"key"`
	Type  *CertificateAuthorityConfigPublicKeyTypeEnum `json:"type"`
}

type jsonCertificateAuthorityConfigPublicKey CertificateAuthorityConfigPublicKey

func (r *CertificateAuthorityConfigPublicKey) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigPublicKey
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigPublicKey
	} else {

		r.Key = res.Key

		r.Type = res.Type

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigPublicKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigPublicKey *CertificateAuthorityConfigPublicKey = &CertificateAuthorityConfigPublicKey{empty: true}

func (r *CertificateAuthorityConfigPublicKey) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigPublicKey) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigPublicKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigReusableConfig struct {
	empty                bool                                                          `json:"-"`
	ReusableConfig       *string                                                       `json:"reusableConfig"`
	ReusableConfigValues *CertificateAuthorityConfigReusableConfigReusableConfigValues `json:"reusableConfigValues"`
}

type jsonCertificateAuthorityConfigReusableConfig CertificateAuthorityConfigReusableConfig

func (r *CertificateAuthorityConfigReusableConfig) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigReusableConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigReusableConfig
	} else {

		r.ReusableConfig = res.ReusableConfig

		r.ReusableConfigValues = res.ReusableConfigValues

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigReusableConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigReusableConfig *CertificateAuthorityConfigReusableConfig = &CertificateAuthorityConfigReusableConfig{empty: true}

func (r *CertificateAuthorityConfigReusableConfig) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigReusableConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigReusableConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigReusableConfigReusableConfigValues struct {
	empty                bool                                                                               `json:"-"`
	KeyUsage             *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage              `json:"keyUsage"`
	CaOptions            *CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions             `json:"caOptions"`
	PolicyIds            []CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds            `json:"policyIds"`
	AiaOcspServers       []string                                                                           `json:"aiaOcspServers"`
	AdditionalExtensions []CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions `json:"additionalExtensions"`
}

type jsonCertificateAuthorityConfigReusableConfigReusableConfigValues CertificateAuthorityConfigReusableConfigReusableConfigValues

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValues) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigReusableConfigReusableConfigValues
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigReusableConfigReusableConfigValues
	} else {

		r.KeyUsage = res.KeyUsage

		r.CaOptions = res.CaOptions

		r.PolicyIds = res.PolicyIds

		r.AiaOcspServers = res.AiaOcspServers

		r.AdditionalExtensions = res.AdditionalExtensions

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigReusableConfigReusableConfigValues is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigReusableConfigReusableConfigValues *CertificateAuthorityConfigReusableConfigReusableConfigValues = &CertificateAuthorityConfigReusableConfigReusableConfigValues{empty: true}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValues) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValues) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValues) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage struct {
	empty                    bool                                                                                           `json:"-"`
	BaseKeyUsage             *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage              `json:"baseKeyUsage"`
	ExtendedKeyUsage         *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage          `json:"extendedKeyUsage"`
	UnknownExtendedKeyUsages []CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages `json:"unknownExtendedKeyUsages"`
}

type jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage
	} else {

		r.BaseKeyUsage = res.BaseKeyUsage

		r.ExtendedKeyUsage = res.ExtendedKeyUsage

		r.UnknownExtendedKeyUsages = res.UnknownExtendedKeyUsages

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage = &CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage{empty: true}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage struct {
	empty             bool  `json:"-"`
	DigitalSignature  *bool `json:"digitalSignature"`
	ContentCommitment *bool `json:"contentCommitment"`
	KeyEncipherment   *bool `json:"keyEncipherment"`
	DataEncipherment  *bool `json:"dataEncipherment"`
	KeyAgreement      *bool `json:"keyAgreement"`
	CertSign          *bool `json:"certSign"`
	CrlSign           *bool `json:"crlSign"`
	EncipherOnly      *bool `json:"encipherOnly"`
	DecipherOnly      *bool `json:"decipherOnly"`
}

type jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage
	} else {

		r.DigitalSignature = res.DigitalSignature

		r.ContentCommitment = res.ContentCommitment

		r.KeyEncipherment = res.KeyEncipherment

		r.DataEncipherment = res.DataEncipherment

		r.KeyAgreement = res.KeyAgreement

		r.CertSign = res.CertSign

		r.CrlSign = res.CrlSign

		r.EncipherOnly = res.EncipherOnly

		r.DecipherOnly = res.DecipherOnly

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage = &CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage{empty: true}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageBaseKeyUsage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage struct {
	empty           bool  `json:"-"`
	ServerAuth      *bool `json:"serverAuth"`
	ClientAuth      *bool `json:"clientAuth"`
	CodeSigning     *bool `json:"codeSigning"`
	EmailProtection *bool `json:"emailProtection"`
	TimeStamping    *bool `json:"timeStamping"`
	OcspSigning     *bool `json:"ocspSigning"`
}

type jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage
	} else {

		r.ServerAuth = res.ServerAuth

		r.ClientAuth = res.ClientAuth

		r.CodeSigning = res.CodeSigning

		r.EmailProtection = res.EmailProtection

		r.TimeStamping = res.TimeStamping

		r.OcspSigning = res.OcspSigning

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage = &CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage{empty: true}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageExtendedKeyUsage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages struct {
	empty        bool    `json:"-"`
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	} else {

		r.ObjectIdPath = res.ObjectIdPath

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages = &CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{empty: true}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions struct {
	empty               bool   `json:"-"`
	IsCa                *bool  `json:"isCa"`
	MaxIssuerPathLength *int64 `json:"maxIssuerPathLength"`
}

type jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions
	} else {

		r.IsCa = res.IsCa

		r.MaxIssuerPathLength = res.MaxIssuerPathLength

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions *CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions = &CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions{empty: true}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesCaOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds struct {
	empty        bool    `json:"-"`
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds
	} else {

		r.ObjectIdPath = res.ObjectIdPath

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds *CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds = &CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds{empty: true}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesPolicyIds) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions struct {
	empty    bool                                                                                      `json:"-"`
	ObjectId *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId `json:"objectId"`
	Critical *bool                                                                                     `json:"critical"`
	Value    *string                                                                                   `json:"value"`
}

type jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions
	} else {

		r.ObjectId = res.ObjectId

		r.Critical = res.Critical

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions = &CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions{empty: true}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId struct {
	empty        bool    `json:"-"`
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId
	} else {

		r.ObjectIdPath = res.ObjectIdPath

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId = &CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId{empty: true}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityConfigReusableConfigReusableConfigValuesAdditionalExtensionsObjectId) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityKeySpec struct {
	empty              bool                                      `json:"-"`
	CloudKmsKeyVersion *string                                   `json:"cloudKmsKeyVersion"`
	Algorithm          *CertificateAuthorityKeySpecAlgorithmEnum `json:"algorithm"`
}

type jsonCertificateAuthorityKeySpec CertificateAuthorityKeySpec

func (r *CertificateAuthorityKeySpec) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityKeySpec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityKeySpec
	} else {

		r.CloudKmsKeyVersion = res.CloudKmsKeyVersion

		r.Algorithm = res.Algorithm

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityKeySpec is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityKeySpec *CertificateAuthorityKeySpec = &CertificateAuthorityKeySpec{empty: true}

func (r *CertificateAuthorityKeySpec) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityKeySpec) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityKeySpec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthoritySubordinateConfig struct {
	empty                bool                                                 `json:"-"`
	CertificateAuthority *string                                              `json:"certificateAuthority"`
	PemIssuerChain       *CertificateAuthoritySubordinateConfigPemIssuerChain `json:"pemIssuerChain"`
}

type jsonCertificateAuthoritySubordinateConfig CertificateAuthoritySubordinateConfig

func (r *CertificateAuthoritySubordinateConfig) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthoritySubordinateConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthoritySubordinateConfig
	} else {

		r.CertificateAuthority = res.CertificateAuthority

		r.PemIssuerChain = res.PemIssuerChain

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthoritySubordinateConfig is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthoritySubordinateConfig *CertificateAuthoritySubordinateConfig = &CertificateAuthoritySubordinateConfig{empty: true}

func (r *CertificateAuthoritySubordinateConfig) Empty() bool {
	return r.empty
}

func (r *CertificateAuthoritySubordinateConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthoritySubordinateConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthoritySubordinateConfigPemIssuerChain struct {
	empty           bool     `json:"-"`
	PemCertificates []string `json:"pemCertificates"`
}

type jsonCertificateAuthoritySubordinateConfigPemIssuerChain CertificateAuthoritySubordinateConfigPemIssuerChain

func (r *CertificateAuthoritySubordinateConfigPemIssuerChain) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthoritySubordinateConfigPemIssuerChain
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthoritySubordinateConfigPemIssuerChain
	} else {

		r.PemCertificates = res.PemCertificates

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthoritySubordinateConfigPemIssuerChain is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthoritySubordinateConfigPemIssuerChain *CertificateAuthoritySubordinateConfigPemIssuerChain = &CertificateAuthoritySubordinateConfigPemIssuerChain{empty: true}

func (r *CertificateAuthoritySubordinateConfigPemIssuerChain) Empty() bool {
	return r.empty
}

func (r *CertificateAuthoritySubordinateConfigPemIssuerChain) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthoritySubordinateConfigPemIssuerChain) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptions struct {
	empty                     bool                                                             `json:"-"`
	SubjectDescription        *CertificateAuthorityCaCertificateDescriptionsSubjectDescription `json:"subjectDescription"`
	PublicKey                 *CertificateAuthorityCaCertificateDescriptionsPublicKey          `json:"publicKey"`
	SubjectKeyId              *CertificateAuthorityCaCertificateDescriptionsSubjectKeyId       `json:"subjectKeyId"`
	AuthorityKeyId            *CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId     `json:"authorityKeyId"`
	CrlDistributionPoints     []string                                                         `json:"crlDistributionPoints"`
	AiaIssuingCertificateUrls []string                                                         `json:"aiaIssuingCertificateUrls"`
	CertFingerprint           *CertificateAuthorityCaCertificateDescriptionsCertFingerprint    `json:"certFingerprint"`
	ConfigValues              *CertificateAuthorityCaCertificateDescriptionsConfigValues       `json:"configValues"`
}

type jsonCertificateAuthorityCaCertificateDescriptions CertificateAuthorityCaCertificateDescriptions

func (r *CertificateAuthorityCaCertificateDescriptions) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptions
	} else {

		r.SubjectDescription = res.SubjectDescription

		r.PublicKey = res.PublicKey

		r.SubjectKeyId = res.SubjectKeyId

		r.AuthorityKeyId = res.AuthorityKeyId

		r.CrlDistributionPoints = res.CrlDistributionPoints

		r.AiaIssuingCertificateUrls = res.AiaIssuingCertificateUrls

		r.CertFingerprint = res.CertFingerprint

		r.ConfigValues = res.ConfigValues

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptions *CertificateAuthorityCaCertificateDescriptions = &CertificateAuthorityCaCertificateDescriptions{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptions) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptions) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsSubjectDescription struct {
	empty           bool                                                                           `json:"-"`
	Subject         *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject        `json:"subject"`
	SubjectAltName  *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName `json:"subjectAltName"`
	HexSerialNumber *string                                                                        `json:"hexSerialNumber"`
	Lifetime        *string                                                                        `json:"lifetime"`
	NotBeforeTime   *string                                                                        `json:"notBeforeTime"`
	NotAfterTime    *string                                                                        `json:"notAfterTime"`
	CommonName      *string                                                                        `json:"commonName"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsSubjectDescription CertificateAuthorityCaCertificateDescriptionsSubjectDescription

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescription) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsSubjectDescription
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescription
	} else {

		r.Subject = res.Subject

		r.SubjectAltName = res.SubjectAltName

		r.HexSerialNumber = res.HexSerialNumber

		r.Lifetime = res.Lifetime

		r.NotBeforeTime = res.NotBeforeTime

		r.NotAfterTime = res.NotAfterTime

		r.CommonName = res.CommonName

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsSubjectDescription is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescription *CertificateAuthorityCaCertificateDescriptionsSubjectDescription = &CertificateAuthorityCaCertificateDescriptionsSubjectDescription{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescription) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescription) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescription) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject struct {
	empty              bool    `json:"-"`
	CommonName         *string `json:"commonName"`
	CountryCode        *string `json:"countryCode"`
	Organization       *string `json:"organization"`
	OrganizationalUnit *string `json:"organizationalUnit"`
	Locality           *string `json:"locality"`
	Province           *string `json:"province"`
	StreetAddress      *string `json:"streetAddress"`
	PostalCode         *string `json:"postalCode"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject
	} else {

		r.CommonName = res.CommonName

		r.CountryCode = res.CountryCode

		r.Organization = res.Organization

		r.OrganizationalUnit = res.OrganizationalUnit

		r.Locality = res.Locality

		r.Province = res.Province

		r.StreetAddress = res.StreetAddress

		r.PostalCode = res.PostalCode

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject = &CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubject) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName struct {
	empty          bool                                                                                      `json:"-"`
	DnsNames       []string                                                                                  `json:"dnsNames"`
	Uris           []string                                                                                  `json:"uris"`
	EmailAddresses []string                                                                                  `json:"emailAddresses"`
	IPAddresses    []string                                                                                  `json:"ipAddresses"`
	CustomSans     []CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans `json:"customSans"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName
	} else {

		r.DnsNames = res.DnsNames

		r.Uris = res.Uris

		r.EmailAddresses = res.EmailAddresses

		r.IPAddresses = res.IPAddresses

		r.CustomSans = res.CustomSans

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName = &CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltName) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans struct {
	empty    bool                                                                                             `json:"-"`
	ObjectId *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId `json:"objectId"`
	Critical *bool                                                                                            `json:"critical"`
	Value    *string                                                                                          `json:"value"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans
	} else {

		r.ObjectId = res.ObjectId

		r.Critical = res.Critical

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans = &CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSans) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId struct {
	empty        bool    `json:"-"`
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId
	} else {

		r.ObjectIdPath = res.ObjectIdPath

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId = &CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectDescriptionSubjectAltNameCustomSansObjectId) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsPublicKey struct {
	empty  bool                                                              `json:"-"`
	Key    *string                                                           `json:"key"`
	Format *CertificateAuthorityCaCertificateDescriptionsPublicKeyFormatEnum `json:"format"`
	Type   *CertificateAuthorityCaCertificateDescriptionsPublicKeyTypeEnum   `json:"type"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsPublicKey CertificateAuthorityCaCertificateDescriptionsPublicKey

func (r *CertificateAuthorityCaCertificateDescriptionsPublicKey) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsPublicKey
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsPublicKey
	} else {

		r.Key = res.Key

		r.Format = res.Format

		r.Type = res.Type

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsPublicKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsPublicKey *CertificateAuthorityCaCertificateDescriptionsPublicKey = &CertificateAuthorityCaCertificateDescriptionsPublicKey{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsPublicKey) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsPublicKey) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsPublicKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsSubjectKeyId struct {
	empty bool    `json:"-"`
	KeyId *string `json:"keyId"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsSubjectKeyId CertificateAuthorityCaCertificateDescriptionsSubjectKeyId

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsSubjectKeyId
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsSubjectKeyId
	} else {

		r.KeyId = res.KeyId

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsSubjectKeyId is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsSubjectKeyId *CertificateAuthorityCaCertificateDescriptionsSubjectKeyId = &CertificateAuthorityCaCertificateDescriptionsSubjectKeyId{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsSubjectKeyId) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId struct {
	empty bool    `json:"-"`
	KeyId *string `json:"keyId"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId

func (r *CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId
	} else {

		r.KeyId = res.KeyId

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsAuthorityKeyId *CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId = &CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsAuthorityKeyId) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsCertFingerprint struct {
	empty      bool    `json:"-"`
	Sha256Hash *string `json:"sha256Hash"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsCertFingerprint CertificateAuthorityCaCertificateDescriptionsCertFingerprint

func (r *CertificateAuthorityCaCertificateDescriptionsCertFingerprint) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsCertFingerprint
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsCertFingerprint
	} else {

		r.Sha256Hash = res.Sha256Hash

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsCertFingerprint is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsCertFingerprint *CertificateAuthorityCaCertificateDescriptionsCertFingerprint = &CertificateAuthorityCaCertificateDescriptionsCertFingerprint{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsCertFingerprint) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsCertFingerprint) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsCertFingerprint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsConfigValues struct {
	empty                bool                                                                            `json:"-"`
	KeyUsage             *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage              `json:"keyUsage"`
	CaOptions            *CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions             `json:"caOptions"`
	PolicyIds            []CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds            `json:"policyIds"`
	AiaOcspServers       []string                                                                        `json:"aiaOcspServers"`
	AdditionalExtensions []CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions `json:"additionalExtensions"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsConfigValues CertificateAuthorityCaCertificateDescriptionsConfigValues

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValues) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsConfigValues
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsConfigValues
	} else {

		r.KeyUsage = res.KeyUsage

		r.CaOptions = res.CaOptions

		r.PolicyIds = res.PolicyIds

		r.AiaOcspServers = res.AiaOcspServers

		r.AdditionalExtensions = res.AdditionalExtensions

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsConfigValues is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsConfigValues *CertificateAuthorityCaCertificateDescriptionsConfigValues = &CertificateAuthorityCaCertificateDescriptionsConfigValues{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValues) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValues) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValues) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage struct {
	empty                    bool                                                                                        `json:"-"`
	BaseKeyUsage             *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage              `json:"baseKeyUsage"`
	ExtendedKeyUsage         *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage          `json:"extendedKeyUsage"`
	UnknownExtendedKeyUsages []CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages `json:"unknownExtendedKeyUsages"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage
	} else {

		r.BaseKeyUsage = res.BaseKeyUsage

		r.ExtendedKeyUsage = res.ExtendedKeyUsage

		r.UnknownExtendedKeyUsages = res.UnknownExtendedKeyUsages

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage = &CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage struct {
	empty             bool  `json:"-"`
	DigitalSignature  *bool `json:"digitalSignature"`
	ContentCommitment *bool `json:"contentCommitment"`
	KeyEncipherment   *bool `json:"keyEncipherment"`
	DataEncipherment  *bool `json:"dataEncipherment"`
	KeyAgreement      *bool `json:"keyAgreement"`
	CertSign          *bool `json:"certSign"`
	CrlSign           *bool `json:"crlSign"`
	EncipherOnly      *bool `json:"encipherOnly"`
	DecipherOnly      *bool `json:"decipherOnly"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage
	} else {

		r.DigitalSignature = res.DigitalSignature

		r.ContentCommitment = res.ContentCommitment

		r.KeyEncipherment = res.KeyEncipherment

		r.DataEncipherment = res.DataEncipherment

		r.KeyAgreement = res.KeyAgreement

		r.CertSign = res.CertSign

		r.CrlSign = res.CrlSign

		r.EncipherOnly = res.EncipherOnly

		r.DecipherOnly = res.DecipherOnly

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage = &CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageBaseKeyUsage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage struct {
	empty           bool  `json:"-"`
	ServerAuth      *bool `json:"serverAuth"`
	ClientAuth      *bool `json:"clientAuth"`
	CodeSigning     *bool `json:"codeSigning"`
	EmailProtection *bool `json:"emailProtection"`
	TimeStamping    *bool `json:"timeStamping"`
	OcspSigning     *bool `json:"ocspSigning"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage
	} else {

		r.ServerAuth = res.ServerAuth

		r.ClientAuth = res.ClientAuth

		r.CodeSigning = res.CodeSigning

		r.EmailProtection = res.EmailProtection

		r.TimeStamping = res.TimeStamping

		r.OcspSigning = res.OcspSigning

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage = &CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageExtendedKeyUsage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages struct {
	empty        bool    `json:"-"`
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages
	} else {

		r.ObjectIdPath = res.ObjectIdPath

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages = &CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesKeyUsageUnknownExtendedKeyUsages) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions struct {
	empty               bool   `json:"-"`
	IsCa                *bool  `json:"isCa"`
	MaxIssuerPathLength *int64 `json:"maxIssuerPathLength"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions
	} else {

		r.IsCa = res.IsCa

		r.MaxIssuerPathLength = res.MaxIssuerPathLength

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions *CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions = &CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesCaOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds struct {
	empty        bool    `json:"-"`
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds
	} else {

		r.ObjectIdPath = res.ObjectIdPath

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds *CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds = &CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesPolicyIds) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions struct {
	empty    bool                                                                                   `json:"-"`
	ObjectId *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId `json:"objectId"`
	Critical *bool                                                                                  `json:"critical"`
	Value    *string                                                                                `json:"value"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions
	} else {

		r.ObjectId = res.ObjectId

		r.Critical = res.Critical

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions = &CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId struct {
	empty        bool    `json:"-"`
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId
	} else {

		r.ObjectIdPath = res.ObjectIdPath

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId = &CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId{empty: true}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCaCertificateDescriptionsConfigValuesAdditionalExtensionsObjectId) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityAccessUrls struct {
	empty                  bool     `json:"-"`
	CaCertificateAccessUrl *string  `json:"caCertificateAccessUrl"`
	CrlAccessUrls          []string `json:"crlAccessUrls"`
	CrlAccessUrl           *string  `json:"crlAccessUrl"`
}

type jsonCertificateAuthorityAccessUrls CertificateAuthorityAccessUrls

func (r *CertificateAuthorityAccessUrls) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityAccessUrls
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityAccessUrls
	} else {

		r.CaCertificateAccessUrl = res.CaCertificateAccessUrl

		r.CrlAccessUrls = res.CrlAccessUrls

		r.CrlAccessUrl = res.CrlAccessUrl

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityAccessUrls is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityAccessUrls *CertificateAuthorityAccessUrls = &CertificateAuthorityAccessUrls{empty: true}

func (r *CertificateAuthorityAccessUrls) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityAccessUrls) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityAccessUrls) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicy struct {
	empty                            bool                                                                    `json:"-"`
	AllowedConfigList                *CertificateAuthorityCertificatePolicyAllowedConfigList                 `json:"allowedConfigList"`
	OverwriteConfigValues            *CertificateAuthorityCertificatePolicyOverwriteConfigValues             `json:"overwriteConfigValues"`
	AllowedLocationsAndOrganizations []CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations `json:"allowedLocationsAndOrganizations"`
	AllowedCommonNames               []string                                                                `json:"allowedCommonNames"`
	AllowedSans                      *CertificateAuthorityCertificatePolicyAllowedSans                       `json:"allowedSans"`
	MaximumLifetime                  *string                                                                 `json:"maximumLifetime"`
	AllowedIssuanceModes             *CertificateAuthorityCertificatePolicyAllowedIssuanceModes              `json:"allowedIssuanceModes"`
}

type jsonCertificateAuthorityCertificatePolicy CertificateAuthorityCertificatePolicy

func (r *CertificateAuthorityCertificatePolicy) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicy
	} else {

		r.AllowedConfigList = res.AllowedConfigList

		r.OverwriteConfigValues = res.OverwriteConfigValues

		r.AllowedLocationsAndOrganizations = res.AllowedLocationsAndOrganizations

		r.AllowedCommonNames = res.AllowedCommonNames

		r.AllowedSans = res.AllowedSans

		r.MaximumLifetime = res.MaximumLifetime

		r.AllowedIssuanceModes = res.AllowedIssuanceModes

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicy *CertificateAuthorityCertificatePolicy = &CertificateAuthorityCertificatePolicy{empty: true}

func (r *CertificateAuthorityCertificatePolicy) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyAllowedConfigList struct {
	empty               bool                                                                        `json:"-"`
	AllowedConfigValues []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues `json:"allowedConfigValues"`
}

type jsonCertificateAuthorityCertificatePolicyAllowedConfigList CertificateAuthorityCertificatePolicyAllowedConfigList

func (r *CertificateAuthorityCertificatePolicyAllowedConfigList) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyAllowedConfigList
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyAllowedConfigList
	} else {

		r.AllowedConfigValues = res.AllowedConfigValues

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyAllowedConfigList is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyAllowedConfigList *CertificateAuthorityCertificatePolicyAllowedConfigList = &CertificateAuthorityCertificatePolicyAllowedConfigList{empty: true}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigList) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigList) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigList) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues struct {
	empty                bool                                                                                           `json:"-"`
	ReusableConfig       *string                                                                                        `json:"reusableConfig"`
	ReusableConfigValues *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues `json:"reusableConfigValues"`
}

type jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues
	} else {

		r.ReusableConfig = res.ReusableConfig

		r.ReusableConfigValues = res.ReusableConfigValues

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues = &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues{empty: true}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValues) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues struct {
	empty                bool                                                                                                                `json:"-"`
	KeyUsage             *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage              `json:"keyUsage"`
	CaOptions            *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions             `json:"caOptions"`
	PolicyIds            []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds            `json:"policyIds"`
	AiaOcspServers       []string                                                                                                            `json:"aiaOcspServers"`
	AdditionalExtensions []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions `json:"additionalExtensions"`
}

type jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues
	} else {

		r.KeyUsage = res.KeyUsage

		r.CaOptions = res.CaOptions

		r.PolicyIds = res.PolicyIds

		r.AiaOcspServers = res.AiaOcspServers

		r.AdditionalExtensions = res.AdditionalExtensions

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues = &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues{empty: true}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValues) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage struct {
	empty                    bool                                                                                                                            `json:"-"`
	BaseKeyUsage             *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage              `json:"baseKeyUsage"`
	ExtendedKeyUsage         *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage          `json:"extendedKeyUsage"`
	UnknownExtendedKeyUsages []CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages `json:"unknownExtendedKeyUsages"`
}

type jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage
	} else {

		r.BaseKeyUsage = res.BaseKeyUsage

		r.ExtendedKeyUsage = res.ExtendedKeyUsage

		r.UnknownExtendedKeyUsages = res.UnknownExtendedKeyUsages

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage = &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage{empty: true}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage struct {
	empty             bool  `json:"-"`
	DigitalSignature  *bool `json:"digitalSignature"`
	ContentCommitment *bool `json:"contentCommitment"`
	KeyEncipherment   *bool `json:"keyEncipherment"`
	DataEncipherment  *bool `json:"dataEncipherment"`
	KeyAgreement      *bool `json:"keyAgreement"`
	CertSign          *bool `json:"certSign"`
	CrlSign           *bool `json:"crlSign"`
	EncipherOnly      *bool `json:"encipherOnly"`
	DecipherOnly      *bool `json:"decipherOnly"`
}

type jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage
	} else {

		r.DigitalSignature = res.DigitalSignature

		r.ContentCommitment = res.ContentCommitment

		r.KeyEncipherment = res.KeyEncipherment

		r.DataEncipherment = res.DataEncipherment

		r.KeyAgreement = res.KeyAgreement

		r.CertSign = res.CertSign

		r.CrlSign = res.CrlSign

		r.EncipherOnly = res.EncipherOnly

		r.DecipherOnly = res.DecipherOnly

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage = &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{empty: true}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage struct {
	empty           bool  `json:"-"`
	ServerAuth      *bool `json:"serverAuth"`
	ClientAuth      *bool `json:"clientAuth"`
	CodeSigning     *bool `json:"codeSigning"`
	EmailProtection *bool `json:"emailProtection"`
	TimeStamping    *bool `json:"timeStamping"`
	OcspSigning     *bool `json:"ocspSigning"`
}

type jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage
	} else {

		r.ServerAuth = res.ServerAuth

		r.ClientAuth = res.ClientAuth

		r.CodeSigning = res.CodeSigning

		r.EmailProtection = res.EmailProtection

		r.TimeStamping = res.TimeStamping

		r.OcspSigning = res.OcspSigning

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage = &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{empty: true}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages struct {
	empty        bool    `json:"-"`
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	} else {

		r.ObjectIdPath = res.ObjectIdPath

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages = &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{empty: true}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions struct {
	empty               bool   `json:"-"`
	IsCa                *bool  `json:"isCa"`
	MaxIssuerPathLength *int64 `json:"maxIssuerPathLength"`
}

type jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions
	} else {

		r.IsCa = res.IsCa

		r.MaxIssuerPathLength = res.MaxIssuerPathLength

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions = &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions{empty: true}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesCaOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds struct {
	empty        bool    `json:"-"`
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds
	} else {

		r.ObjectIdPath = res.ObjectIdPath

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds = &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds{empty: true}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesPolicyIds) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions struct {
	empty    bool                                                                                                                       `json:"-"`
	ObjectId *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId `json:"objectId"`
	Critical *bool                                                                                                                      `json:"critical"`
	Value    *string                                                                                                                    `json:"value"`
}

type jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions
	} else {

		r.ObjectId = res.ObjectId

		r.Critical = res.Critical

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions = &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions{empty: true}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId struct {
	empty        bool    `json:"-"`
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId
	} else {

		r.ObjectIdPath = res.ObjectIdPath

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId = &CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{empty: true}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyAllowedConfigListAllowedConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyOverwriteConfigValues struct {
	empty                bool                                                                            `json:"-"`
	ReusableConfig       *string                                                                         `json:"reusableConfig"`
	ReusableConfigValues *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues `json:"reusableConfigValues"`
}

type jsonCertificateAuthorityCertificatePolicyOverwriteConfigValues CertificateAuthorityCertificatePolicyOverwriteConfigValues

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValues) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyOverwriteConfigValues
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValues
	} else {

		r.ReusableConfig = res.ReusableConfig

		r.ReusableConfigValues = res.ReusableConfigValues

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyOverwriteConfigValues is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValues *CertificateAuthorityCertificatePolicyOverwriteConfigValues = &CertificateAuthorityCertificatePolicyOverwriteConfigValues{empty: true}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValues) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValues) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValues) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues struct {
	empty                bool                                                                                                 `json:"-"`
	KeyUsage             *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage              `json:"keyUsage"`
	CaOptions            *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions             `json:"caOptions"`
	PolicyIds            []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds            `json:"policyIds"`
	AiaOcspServers       []string                                                                                             `json:"aiaOcspServers"`
	AdditionalExtensions []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions `json:"additionalExtensions"`
}

type jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues
	} else {

		r.KeyUsage = res.KeyUsage

		r.CaOptions = res.CaOptions

		r.PolicyIds = res.PolicyIds

		r.AiaOcspServers = res.AiaOcspServers

		r.AdditionalExtensions = res.AdditionalExtensions

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues = &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues{empty: true}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValues) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage struct {
	empty                    bool                                                                                                             `json:"-"`
	BaseKeyUsage             *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage              `json:"baseKeyUsage"`
	ExtendedKeyUsage         *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage          `json:"extendedKeyUsage"`
	UnknownExtendedKeyUsages []CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages `json:"unknownExtendedKeyUsages"`
}

type jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage
	} else {

		r.BaseKeyUsage = res.BaseKeyUsage

		r.ExtendedKeyUsage = res.ExtendedKeyUsage

		r.UnknownExtendedKeyUsages = res.UnknownExtendedKeyUsages

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage = &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage{empty: true}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage struct {
	empty             bool  `json:"-"`
	DigitalSignature  *bool `json:"digitalSignature"`
	ContentCommitment *bool `json:"contentCommitment"`
	KeyEncipherment   *bool `json:"keyEncipherment"`
	DataEncipherment  *bool `json:"dataEncipherment"`
	KeyAgreement      *bool `json:"keyAgreement"`
	CertSign          *bool `json:"certSign"`
	CrlSign           *bool `json:"crlSign"`
	EncipherOnly      *bool `json:"encipherOnly"`
	DecipherOnly      *bool `json:"decipherOnly"`
}

type jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage
	} else {

		r.DigitalSignature = res.DigitalSignature

		r.ContentCommitment = res.ContentCommitment

		r.KeyEncipherment = res.KeyEncipherment

		r.DataEncipherment = res.DataEncipherment

		r.KeyAgreement = res.KeyAgreement

		r.CertSign = res.CertSign

		r.CrlSign = res.CrlSign

		r.EncipherOnly = res.EncipherOnly

		r.DecipherOnly = res.DecipherOnly

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage = &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage{empty: true}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageBaseKeyUsage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage struct {
	empty           bool  `json:"-"`
	ServerAuth      *bool `json:"serverAuth"`
	ClientAuth      *bool `json:"clientAuth"`
	CodeSigning     *bool `json:"codeSigning"`
	EmailProtection *bool `json:"emailProtection"`
	TimeStamping    *bool `json:"timeStamping"`
	OcspSigning     *bool `json:"ocspSigning"`
}

type jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage
	} else {

		r.ServerAuth = res.ServerAuth

		r.ClientAuth = res.ClientAuth

		r.CodeSigning = res.CodeSigning

		r.EmailProtection = res.EmailProtection

		r.TimeStamping = res.TimeStamping

		r.OcspSigning = res.OcspSigning

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage = &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage{empty: true}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageExtendedKeyUsage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages struct {
	empty        bool    `json:"-"`
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages
	} else {

		r.ObjectIdPath = res.ObjectIdPath

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages = &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages{empty: true}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesKeyUsageUnknownExtendedKeyUsages) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions struct {
	empty               bool   `json:"-"`
	IsCa                *bool  `json:"isCa"`
	MaxIssuerPathLength *int64 `json:"maxIssuerPathLength"`
}

type jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions
	} else {

		r.IsCa = res.IsCa

		r.MaxIssuerPathLength = res.MaxIssuerPathLength

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions = &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions{empty: true}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesCaOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds struct {
	empty        bool    `json:"-"`
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds
	} else {

		r.ObjectIdPath = res.ObjectIdPath

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds = &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds{empty: true}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesPolicyIds) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions struct {
	empty    bool                                                                                                        `json:"-"`
	ObjectId *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId `json:"objectId"`
	Critical *bool                                                                                                       `json:"critical"`
	Value    *string                                                                                                     `json:"value"`
}

type jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions
	} else {

		r.ObjectId = res.ObjectId

		r.Critical = res.Critical

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions = &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions{empty: true}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId struct {
	empty        bool    `json:"-"`
	ObjectIdPath []int64 `json:"objectIdPath"`
}

type jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId
	} else {

		r.ObjectIdPath = res.ObjectIdPath

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId = &CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId{empty: true}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyOverwriteConfigValuesReusableConfigValuesAdditionalExtensionsObjectId) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations struct {
	empty              bool    `json:"-"`
	CountryCode        *string `json:"countryCode"`
	Organization       *string `json:"organization"`
	OrganizationalUnit *string `json:"organizationalUnit"`
	Locality           *string `json:"locality"`
	Province           *string `json:"province"`
	StreetAddress      *string `json:"streetAddress"`
	PostalCode         *string `json:"postalCode"`
}

type jsonCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations

func (r *CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations
	} else {

		r.CountryCode = res.CountryCode

		r.Organization = res.Organization

		r.OrganizationalUnit = res.OrganizationalUnit

		r.Locality = res.Locality

		r.Province = res.Province

		r.StreetAddress = res.StreetAddress

		r.PostalCode = res.PostalCode

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations *CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations = &CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations{empty: true}

func (r *CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyAllowedLocationsAndOrganizations) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyAllowedSans struct {
	empty                     bool     `json:"-"`
	AllowedDnsNames           []string `json:"allowedDnsNames"`
	AllowedUris               []string `json:"allowedUris"`
	AllowedEmailAddresses     []string `json:"allowedEmailAddresses"`
	AllowedIps                []string `json:"allowedIps"`
	AllowGlobbingDnsWildcards *bool    `json:"allowGlobbingDnsWildcards"`
	AllowCustomSans           *bool    `json:"allowCustomSans"`
}

type jsonCertificateAuthorityCertificatePolicyAllowedSans CertificateAuthorityCertificatePolicyAllowedSans

func (r *CertificateAuthorityCertificatePolicyAllowedSans) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyAllowedSans
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyAllowedSans
	} else {

		r.AllowedDnsNames = res.AllowedDnsNames

		r.AllowedUris = res.AllowedUris

		r.AllowedEmailAddresses = res.AllowedEmailAddresses

		r.AllowedIps = res.AllowedIps

		r.AllowGlobbingDnsWildcards = res.AllowGlobbingDnsWildcards

		r.AllowCustomSans = res.AllowCustomSans

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyAllowedSans is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyAllowedSans *CertificateAuthorityCertificatePolicyAllowedSans = &CertificateAuthorityCertificatePolicyAllowedSans{empty: true}

func (r *CertificateAuthorityCertificatePolicyAllowedSans) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyAllowedSans) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyAllowedSans) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityCertificatePolicyAllowedIssuanceModes struct {
	empty                    bool  `json:"-"`
	AllowCsrBasedIssuance    *bool `json:"allowCsrBasedIssuance"`
	AllowConfigBasedIssuance *bool `json:"allowConfigBasedIssuance"`
}

type jsonCertificateAuthorityCertificatePolicyAllowedIssuanceModes CertificateAuthorityCertificatePolicyAllowedIssuanceModes

func (r *CertificateAuthorityCertificatePolicyAllowedIssuanceModes) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityCertificatePolicyAllowedIssuanceModes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityCertificatePolicyAllowedIssuanceModes
	} else {

		r.AllowCsrBasedIssuance = res.AllowCsrBasedIssuance

		r.AllowConfigBasedIssuance = res.AllowConfigBasedIssuance

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityCertificatePolicyAllowedIssuanceModes is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityCertificatePolicyAllowedIssuanceModes *CertificateAuthorityCertificatePolicyAllowedIssuanceModes = &CertificateAuthorityCertificatePolicyAllowedIssuanceModes{empty: true}

func (r *CertificateAuthorityCertificatePolicyAllowedIssuanceModes) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityCertificatePolicyAllowedIssuanceModes) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityCertificatePolicyAllowedIssuanceModes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CertificateAuthorityIssuingOptions struct {
	empty               bool  `json:"-"`
	IncludeCaCertUrl    *bool `json:"includeCaCertUrl"`
	IncludeCrlAccessUrl *bool `json:"includeCrlAccessUrl"`
}

type jsonCertificateAuthorityIssuingOptions CertificateAuthorityIssuingOptions

func (r *CertificateAuthorityIssuingOptions) UnmarshalJSON(data []byte) error {
	var res jsonCertificateAuthorityIssuingOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCertificateAuthorityIssuingOptions
	} else {

		r.IncludeCaCertUrl = res.IncludeCaCertUrl

		r.IncludeCrlAccessUrl = res.IncludeCrlAccessUrl

	}
	return nil
}

// This object is used to assert a desired state where this CertificateAuthorityIssuingOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyCertificateAuthorityIssuingOptions *CertificateAuthorityIssuingOptions = &CertificateAuthorityIssuingOptions{empty: true}

func (r *CertificateAuthorityIssuingOptions) Empty() bool {
	return r.empty
}

func (r *CertificateAuthorityIssuingOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *CertificateAuthorityIssuingOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *CertificateAuthority) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "privateca",
		Type:    "CertificateAuthority",
		Version: "beta",
	}
}

const CertificateAuthorityMaxPage = -1

type CertificateAuthorityList struct {
	Items []*CertificateAuthority

	nextToken string

	pageSize int32

	project string

	location string
}

func (l *CertificateAuthorityList) HasNext() bool {
	return l.nextToken != ""
}

func (l *CertificateAuthorityList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listCertificateAuthority(ctx, l.project, l.location, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListCertificateAuthority(ctx context.Context, project, location string) (*CertificateAuthorityList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListCertificateAuthorityWithMaxResults(ctx, project, location, CertificateAuthorityMaxPage)

}

func (c *Client) ListCertificateAuthorityWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*CertificateAuthorityList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listCertificateAuthority(ctx, project, location, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &CertificateAuthorityList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,

		location: location,
	}, nil
}

func (c *Client) GetCertificateAuthority(ctx context.Context, r *CertificateAuthority) (*CertificateAuthority, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getCertificateAuthorityRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalCertificateAuthority(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeCertificateAuthorityNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteCertificateAuthority(ctx context.Context, r *CertificateAuthority) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("CertificateAuthority resource is nil")
	}
	c.Config.Logger.Info("Deleting CertificateAuthority...")
	deleteOp := deleteCertificateAuthorityOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllCertificateAuthority deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllCertificateAuthority(ctx context.Context, project, location string, filter func(*CertificateAuthority) bool) error {
	listObj, err := c.ListCertificateAuthority(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllCertificateAuthority(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllCertificateAuthority(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyCertificateAuthority(ctx context.Context, rawDesired *CertificateAuthority, opts ...dcl.ApplyOption) (*CertificateAuthority, error) {

	var resultNewState *CertificateAuthority
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyCertificateAuthorityHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyCertificateAuthorityHelper(c *Client, ctx context.Context, rawDesired *CertificateAuthority, opts ...dcl.ApplyOption) (*CertificateAuthority, error) {
	c.Config.Logger.Info("Beginning ApplyCertificateAuthority...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.certificateAuthorityDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	opStrings := dcl.DeduplicateOperations(fieldDiffs)
	diffs, err := convertFieldDiffToCertificateAuthorityOp(opStrings, fieldDiffs, opts)
	if err != nil {
		return nil, err
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
	var ops []certificateAuthorityApiOperation
	if create {
		ops = append(ops, &createCertificateAuthorityOperation{})
	} else if recreate {

		ops = append(ops, &deleteCertificateAuthorityOperation{})

		ops = append(ops, &createCertificateAuthorityOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeCertificateAuthorityDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetCertificateAuthority(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createCertificateAuthorityOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapCertificateAuthority(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeCertificateAuthorityNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeCertificateAuthorityNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeCertificateAuthorityDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffCertificateAuthority(c, newDesired, newState)
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
