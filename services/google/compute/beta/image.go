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

type Image struct {
	ArchiveSizeBytes             *int64                             `json:"archiveSizeBytes"`
	Description                  *string                            `json:"description"`
	DiskSizeGb                   *int64                             `json:"diskSizeGb"`
	Family                       *string                            `json:"family"`
	GuestOsFeature               []ImageGuestOsFeature              `json:"guestOsFeature"`
	ImageEncryptionKey           *ImageImageEncryptionKey           `json:"imageEncryptionKey"`
	Labels                       map[string]string                  `json:"labels"`
	License                      []string                           `json:"license"`
	Name                         *string                            `json:"name"`
	RawDisk                      *ImageRawDisk                      `json:"rawDisk"`
	ShieldedInstanceInitialState *ImageShieldedInstanceInitialState `json:"shieldedInstanceInitialState"`
	SelfLink                     *string                            `json:"selfLink"`
	SourceDisk                   *string                            `json:"sourceDisk"`
	SourceDiskEncryptionKey      *ImageSourceDiskEncryptionKey      `json:"sourceDiskEncryptionKey"`
	SourceDiskId                 *string                            `json:"sourceDiskId"`
	SourceImage                  *string                            `json:"sourceImage"`
	SourceImageEncryptionKey     *ImageSourceImageEncryptionKey     `json:"sourceImageEncryptionKey"`
	SourceImageId                *string                            `json:"sourceImageId"`
	SourceSnapshot               *string                            `json:"sourceSnapshot"`
	SourceSnapshotEncryptionKey  *ImageSourceSnapshotEncryptionKey  `json:"sourceSnapshotEncryptionKey"`
	SourceSnapshotId             *string                            `json:"sourceSnapshotId"`
	SourceType                   *ImageSourceTypeEnum               `json:"sourceType"`
	Status                       *ImageStatusEnum                   `json:"status"`
	StorageLocation              []string                           `json:"storageLocation"`
	Deprecated                   *ImageDeprecated                   `json:"deprecated"`
	Project                      *string                            `json:"project"`
}

func (r *Image) String() string {
	return dcl.SprintResource(r)
}

// The enum ImageGuestOsFeatureTypeEnum.
type ImageGuestOsFeatureTypeEnum string

// ImageGuestOsFeatureTypeEnumRef returns a *ImageGuestOsFeatureTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ImageGuestOsFeatureTypeEnumRef(s string) *ImageGuestOsFeatureTypeEnum {
	if s == "" {
		return nil
	}

	v := ImageGuestOsFeatureTypeEnum(s)
	return &v
}

func (v ImageGuestOsFeatureTypeEnum) Validate() error {
	for _, s := range []string{"FEATURE_TYPE_UNSPECIFIED", "VIRTIO_SCSI_MULTIQUEUE", "WINDOWS", "MULTI_IP_SUBNET", "UEFI_COMPATIBLE", "SECURE_BOOT", "SEV_CAPABLE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ImageGuestOsFeatureTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ImageRawDiskContainerTypeEnum.
type ImageRawDiskContainerTypeEnum string

// ImageRawDiskContainerTypeEnumRef returns a *ImageRawDiskContainerTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ImageRawDiskContainerTypeEnumRef(s string) *ImageRawDiskContainerTypeEnum {
	if s == "" {
		return nil
	}

	v := ImageRawDiskContainerTypeEnum(s)
	return &v
}

func (v ImageRawDiskContainerTypeEnum) Validate() error {
	for _, s := range []string{"TAR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ImageRawDiskContainerTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ImageShieldedInstanceInitialStatePkFileTypeEnum.
type ImageShieldedInstanceInitialStatePkFileTypeEnum string

// ImageShieldedInstanceInitialStatePkFileTypeEnumRef returns a *ImageShieldedInstanceInitialStatePkFileTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ImageShieldedInstanceInitialStatePkFileTypeEnumRef(s string) *ImageShieldedInstanceInitialStatePkFileTypeEnum {
	if s == "" {
		return nil
	}

	v := ImageShieldedInstanceInitialStatePkFileTypeEnum(s)
	return &v
}

func (v ImageShieldedInstanceInitialStatePkFileTypeEnum) Validate() error {
	for _, s := range []string{"UNDEFINED", "BIN", "X509"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ImageShieldedInstanceInitialStatePkFileTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ImageShieldedInstanceInitialStateKekFileTypeEnum.
type ImageShieldedInstanceInitialStateKekFileTypeEnum string

// ImageShieldedInstanceInitialStateKekFileTypeEnumRef returns a *ImageShieldedInstanceInitialStateKekFileTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ImageShieldedInstanceInitialStateKekFileTypeEnumRef(s string) *ImageShieldedInstanceInitialStateKekFileTypeEnum {
	if s == "" {
		return nil
	}

	v := ImageShieldedInstanceInitialStateKekFileTypeEnum(s)
	return &v
}

func (v ImageShieldedInstanceInitialStateKekFileTypeEnum) Validate() error {
	for _, s := range []string{"UNDEFINED", "BIN", "X509"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ImageShieldedInstanceInitialStateKekFileTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ImageShieldedInstanceInitialStateDbFileTypeEnum.
type ImageShieldedInstanceInitialStateDbFileTypeEnum string

// ImageShieldedInstanceInitialStateDbFileTypeEnumRef returns a *ImageShieldedInstanceInitialStateDbFileTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ImageShieldedInstanceInitialStateDbFileTypeEnumRef(s string) *ImageShieldedInstanceInitialStateDbFileTypeEnum {
	if s == "" {
		return nil
	}

	v := ImageShieldedInstanceInitialStateDbFileTypeEnum(s)
	return &v
}

func (v ImageShieldedInstanceInitialStateDbFileTypeEnum) Validate() error {
	for _, s := range []string{"UNDEFINED", "BIN", "X509"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ImageShieldedInstanceInitialStateDbFileTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ImageShieldedInstanceInitialStateDbxFileTypeEnum.
type ImageShieldedInstanceInitialStateDbxFileTypeEnum string

// ImageShieldedInstanceInitialStateDbxFileTypeEnumRef returns a *ImageShieldedInstanceInitialStateDbxFileTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ImageShieldedInstanceInitialStateDbxFileTypeEnumRef(s string) *ImageShieldedInstanceInitialStateDbxFileTypeEnum {
	if s == "" {
		return nil
	}

	v := ImageShieldedInstanceInitialStateDbxFileTypeEnum(s)
	return &v
}

func (v ImageShieldedInstanceInitialStateDbxFileTypeEnum) Validate() error {
	for _, s := range []string{"UNDEFINED", "BIN", "X509"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ImageShieldedInstanceInitialStateDbxFileTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ImageSourceTypeEnum.
type ImageSourceTypeEnum string

// ImageSourceTypeEnumRef returns a *ImageSourceTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ImageSourceTypeEnumRef(s string) *ImageSourceTypeEnum {
	if s == "" {
		return nil
	}

	v := ImageSourceTypeEnum(s)
	return &v
}

func (v ImageSourceTypeEnum) Validate() error {
	for _, s := range []string{"RAW"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ImageSourceTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ImageStatusEnum.
type ImageStatusEnum string

// ImageStatusEnumRef returns a *ImageStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func ImageStatusEnumRef(s string) *ImageStatusEnum {
	if s == "" {
		return nil
	}

	v := ImageStatusEnum(s)
	return &v
}

func (v ImageStatusEnum) Validate() error {
	for _, s := range []string{"PENDING", "READY", "FAILED", "DELETING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ImageStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ImageDeprecatedStateEnum.
type ImageDeprecatedStateEnum string

// ImageDeprecatedStateEnumRef returns a *ImageDeprecatedStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func ImageDeprecatedStateEnumRef(s string) *ImageDeprecatedStateEnum {
	if s == "" {
		return nil
	}

	v := ImageDeprecatedStateEnum(s)
	return &v
}

func (v ImageDeprecatedStateEnum) Validate() error {
	for _, s := range []string{"DEPRECATED", "OBSOLETE", "DELETED", "ACTIVE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ImageDeprecatedStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ImageGuestOsFeature struct {
	empty bool                         `json:"-"`
	Type  *ImageGuestOsFeatureTypeEnum `json:"type"`
}

type jsonImageGuestOsFeature ImageGuestOsFeature

func (r *ImageGuestOsFeature) UnmarshalJSON(data []byte) error {
	var res jsonImageGuestOsFeature
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyImageGuestOsFeature
	} else {

		r.Type = res.Type

	}
	return nil
}

// This object is used to assert a desired state where this ImageGuestOsFeature is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyImageGuestOsFeature *ImageGuestOsFeature = &ImageGuestOsFeature{empty: true}

func (r *ImageGuestOsFeature) Empty() bool {
	return r.empty
}

func (r *ImageGuestOsFeature) String() string {
	return dcl.SprintResource(r)
}

func (r *ImageGuestOsFeature) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ImageImageEncryptionKey struct {
	empty                bool    `json:"-"`
	RawKey               *string `json:"rawKey"`
	KmsKeyName           *string `json:"kmsKeyName"`
	Sha256               *string `json:"sha256"`
	KmsKeyServiceAccount *string `json:"kmsKeyServiceAccount"`
}

type jsonImageImageEncryptionKey ImageImageEncryptionKey

func (r *ImageImageEncryptionKey) UnmarshalJSON(data []byte) error {
	var res jsonImageImageEncryptionKey
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyImageImageEncryptionKey
	} else {

		r.RawKey = res.RawKey

		r.KmsKeyName = res.KmsKeyName

		r.Sha256 = res.Sha256

		r.KmsKeyServiceAccount = res.KmsKeyServiceAccount

	}
	return nil
}

// This object is used to assert a desired state where this ImageImageEncryptionKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyImageImageEncryptionKey *ImageImageEncryptionKey = &ImageImageEncryptionKey{empty: true}

func (r *ImageImageEncryptionKey) Empty() bool {
	return r.empty
}

func (r *ImageImageEncryptionKey) String() string {
	return dcl.SprintResource(r)
}

func (r *ImageImageEncryptionKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ImageRawDisk struct {
	empty         bool                           `json:"-"`
	Source        *string                        `json:"source"`
	Sha1Checksum  *string                        `json:"sha1Checksum"`
	ContainerType *ImageRawDiskContainerTypeEnum `json:"containerType"`
}

type jsonImageRawDisk ImageRawDisk

func (r *ImageRawDisk) UnmarshalJSON(data []byte) error {
	var res jsonImageRawDisk
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyImageRawDisk
	} else {

		r.Source = res.Source

		r.Sha1Checksum = res.Sha1Checksum

		r.ContainerType = res.ContainerType

	}
	return nil
}

// This object is used to assert a desired state where this ImageRawDisk is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyImageRawDisk *ImageRawDisk = &ImageRawDisk{empty: true}

func (r *ImageRawDisk) Empty() bool {
	return r.empty
}

func (r *ImageRawDisk) String() string {
	return dcl.SprintResource(r)
}

func (r *ImageRawDisk) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ImageShieldedInstanceInitialState struct {
	empty bool                                   `json:"-"`
	Pk    *ImageShieldedInstanceInitialStatePk   `json:"pk"`
	Kek   []ImageShieldedInstanceInitialStateKek `json:"kek"`
	Db    []ImageShieldedInstanceInitialStateDb  `json:"db"`
	Dbx   []ImageShieldedInstanceInitialStateDbx `json:"dbx"`
}

type jsonImageShieldedInstanceInitialState ImageShieldedInstanceInitialState

func (r *ImageShieldedInstanceInitialState) UnmarshalJSON(data []byte) error {
	var res jsonImageShieldedInstanceInitialState
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyImageShieldedInstanceInitialState
	} else {

		r.Pk = res.Pk

		r.Kek = res.Kek

		r.Db = res.Db

		r.Dbx = res.Dbx

	}
	return nil
}

// This object is used to assert a desired state where this ImageShieldedInstanceInitialState is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyImageShieldedInstanceInitialState *ImageShieldedInstanceInitialState = &ImageShieldedInstanceInitialState{empty: true}

func (r *ImageShieldedInstanceInitialState) Empty() bool {
	return r.empty
}

func (r *ImageShieldedInstanceInitialState) String() string {
	return dcl.SprintResource(r)
}

func (r *ImageShieldedInstanceInitialState) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ImageShieldedInstanceInitialStatePk struct {
	empty    bool                                             `json:"-"`
	Content  *string                                          `json:"content"`
	FileType *ImageShieldedInstanceInitialStatePkFileTypeEnum `json:"fileType"`
}

type jsonImageShieldedInstanceInitialStatePk ImageShieldedInstanceInitialStatePk

func (r *ImageShieldedInstanceInitialStatePk) UnmarshalJSON(data []byte) error {
	var res jsonImageShieldedInstanceInitialStatePk
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyImageShieldedInstanceInitialStatePk
	} else {

		r.Content = res.Content

		r.FileType = res.FileType

	}
	return nil
}

// This object is used to assert a desired state where this ImageShieldedInstanceInitialStatePk is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyImageShieldedInstanceInitialStatePk *ImageShieldedInstanceInitialStatePk = &ImageShieldedInstanceInitialStatePk{empty: true}

func (r *ImageShieldedInstanceInitialStatePk) Empty() bool {
	return r.empty
}

func (r *ImageShieldedInstanceInitialStatePk) String() string {
	return dcl.SprintResource(r)
}

func (r *ImageShieldedInstanceInitialStatePk) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ImageShieldedInstanceInitialStateKek struct {
	empty    bool                                              `json:"-"`
	Content  *string                                           `json:"content"`
	FileType *ImageShieldedInstanceInitialStateKekFileTypeEnum `json:"fileType"`
}

type jsonImageShieldedInstanceInitialStateKek ImageShieldedInstanceInitialStateKek

func (r *ImageShieldedInstanceInitialStateKek) UnmarshalJSON(data []byte) error {
	var res jsonImageShieldedInstanceInitialStateKek
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyImageShieldedInstanceInitialStateKek
	} else {

		r.Content = res.Content

		r.FileType = res.FileType

	}
	return nil
}

// This object is used to assert a desired state where this ImageShieldedInstanceInitialStateKek is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyImageShieldedInstanceInitialStateKek *ImageShieldedInstanceInitialStateKek = &ImageShieldedInstanceInitialStateKek{empty: true}

func (r *ImageShieldedInstanceInitialStateKek) Empty() bool {
	return r.empty
}

func (r *ImageShieldedInstanceInitialStateKek) String() string {
	return dcl.SprintResource(r)
}

func (r *ImageShieldedInstanceInitialStateKek) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ImageShieldedInstanceInitialStateDb struct {
	empty    bool                                             `json:"-"`
	Content  *string                                          `json:"content"`
	FileType *ImageShieldedInstanceInitialStateDbFileTypeEnum `json:"fileType"`
}

type jsonImageShieldedInstanceInitialStateDb ImageShieldedInstanceInitialStateDb

func (r *ImageShieldedInstanceInitialStateDb) UnmarshalJSON(data []byte) error {
	var res jsonImageShieldedInstanceInitialStateDb
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyImageShieldedInstanceInitialStateDb
	} else {

		r.Content = res.Content

		r.FileType = res.FileType

	}
	return nil
}

// This object is used to assert a desired state where this ImageShieldedInstanceInitialStateDb is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyImageShieldedInstanceInitialStateDb *ImageShieldedInstanceInitialStateDb = &ImageShieldedInstanceInitialStateDb{empty: true}

func (r *ImageShieldedInstanceInitialStateDb) Empty() bool {
	return r.empty
}

func (r *ImageShieldedInstanceInitialStateDb) String() string {
	return dcl.SprintResource(r)
}

func (r *ImageShieldedInstanceInitialStateDb) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ImageShieldedInstanceInitialStateDbx struct {
	empty    bool                                              `json:"-"`
	Content  *string                                           `json:"content"`
	FileType *ImageShieldedInstanceInitialStateDbxFileTypeEnum `json:"fileType"`
}

type jsonImageShieldedInstanceInitialStateDbx ImageShieldedInstanceInitialStateDbx

func (r *ImageShieldedInstanceInitialStateDbx) UnmarshalJSON(data []byte) error {
	var res jsonImageShieldedInstanceInitialStateDbx
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyImageShieldedInstanceInitialStateDbx
	} else {

		r.Content = res.Content

		r.FileType = res.FileType

	}
	return nil
}

// This object is used to assert a desired state where this ImageShieldedInstanceInitialStateDbx is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyImageShieldedInstanceInitialStateDbx *ImageShieldedInstanceInitialStateDbx = &ImageShieldedInstanceInitialStateDbx{empty: true}

func (r *ImageShieldedInstanceInitialStateDbx) Empty() bool {
	return r.empty
}

func (r *ImageShieldedInstanceInitialStateDbx) String() string {
	return dcl.SprintResource(r)
}

func (r *ImageShieldedInstanceInitialStateDbx) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ImageSourceDiskEncryptionKey struct {
	empty                bool    `json:"-"`
	RawKey               *string `json:"rawKey"`
	KmsKeyName           *string `json:"kmsKeyName"`
	Sha256               *string `json:"sha256"`
	KmsKeyServiceAccount *string `json:"kmsKeyServiceAccount"`
}

type jsonImageSourceDiskEncryptionKey ImageSourceDiskEncryptionKey

func (r *ImageSourceDiskEncryptionKey) UnmarshalJSON(data []byte) error {
	var res jsonImageSourceDiskEncryptionKey
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyImageSourceDiskEncryptionKey
	} else {

		r.RawKey = res.RawKey

		r.KmsKeyName = res.KmsKeyName

		r.Sha256 = res.Sha256

		r.KmsKeyServiceAccount = res.KmsKeyServiceAccount

	}
	return nil
}

// This object is used to assert a desired state where this ImageSourceDiskEncryptionKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyImageSourceDiskEncryptionKey *ImageSourceDiskEncryptionKey = &ImageSourceDiskEncryptionKey{empty: true}

func (r *ImageSourceDiskEncryptionKey) Empty() bool {
	return r.empty
}

func (r *ImageSourceDiskEncryptionKey) String() string {
	return dcl.SprintResource(r)
}

func (r *ImageSourceDiskEncryptionKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ImageSourceImageEncryptionKey struct {
	empty                bool    `json:"-"`
	RawKey               *string `json:"rawKey"`
	KmsKeyName           *string `json:"kmsKeyName"`
	Sha256               *string `json:"sha256"`
	KmsKeyServiceAccount *string `json:"kmsKeyServiceAccount"`
}

type jsonImageSourceImageEncryptionKey ImageSourceImageEncryptionKey

func (r *ImageSourceImageEncryptionKey) UnmarshalJSON(data []byte) error {
	var res jsonImageSourceImageEncryptionKey
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyImageSourceImageEncryptionKey
	} else {

		r.RawKey = res.RawKey

		r.KmsKeyName = res.KmsKeyName

		r.Sha256 = res.Sha256

		r.KmsKeyServiceAccount = res.KmsKeyServiceAccount

	}
	return nil
}

// This object is used to assert a desired state where this ImageSourceImageEncryptionKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyImageSourceImageEncryptionKey *ImageSourceImageEncryptionKey = &ImageSourceImageEncryptionKey{empty: true}

func (r *ImageSourceImageEncryptionKey) Empty() bool {
	return r.empty
}

func (r *ImageSourceImageEncryptionKey) String() string {
	return dcl.SprintResource(r)
}

func (r *ImageSourceImageEncryptionKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ImageSourceSnapshotEncryptionKey struct {
	empty                bool    `json:"-"`
	RawKey               *string `json:"rawKey"`
	KmsKeyName           *string `json:"kmsKeyName"`
	Sha256               *string `json:"sha256"`
	KmsKeyServiceAccount *string `json:"kmsKeyServiceAccount"`
}

type jsonImageSourceSnapshotEncryptionKey ImageSourceSnapshotEncryptionKey

func (r *ImageSourceSnapshotEncryptionKey) UnmarshalJSON(data []byte) error {
	var res jsonImageSourceSnapshotEncryptionKey
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyImageSourceSnapshotEncryptionKey
	} else {

		r.RawKey = res.RawKey

		r.KmsKeyName = res.KmsKeyName

		r.Sha256 = res.Sha256

		r.KmsKeyServiceAccount = res.KmsKeyServiceAccount

	}
	return nil
}

// This object is used to assert a desired state where this ImageSourceSnapshotEncryptionKey is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyImageSourceSnapshotEncryptionKey *ImageSourceSnapshotEncryptionKey = &ImageSourceSnapshotEncryptionKey{empty: true}

func (r *ImageSourceSnapshotEncryptionKey) Empty() bool {
	return r.empty
}

func (r *ImageSourceSnapshotEncryptionKey) String() string {
	return dcl.SprintResource(r)
}

func (r *ImageSourceSnapshotEncryptionKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ImageDeprecated struct {
	empty       bool                      `json:"-"`
	State       *ImageDeprecatedStateEnum `json:"state"`
	Replacement *string                   `json:"replacement"`
	Deprecated  *string                   `json:"deprecated"`
	Obsolete    *string                   `json:"obsolete"`
	Deleted     *string                   `json:"deleted"`
}

type jsonImageDeprecated ImageDeprecated

func (r *ImageDeprecated) UnmarshalJSON(data []byte) error {
	var res jsonImageDeprecated
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyImageDeprecated
	} else {

		r.State = res.State

		r.Replacement = res.Replacement

		r.Deprecated = res.Deprecated

		r.Obsolete = res.Obsolete

		r.Deleted = res.Deleted

	}
	return nil
}

// This object is used to assert a desired state where this ImageDeprecated is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyImageDeprecated *ImageDeprecated = &ImageDeprecated{empty: true}

func (r *ImageDeprecated) Empty() bool {
	return r.empty
}

func (r *ImageDeprecated) String() string {
	return dcl.SprintResource(r)
}

func (r *ImageDeprecated) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Image) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "compute",
		Type:    "Image",
		Version: "beta",
	}
}

const ImageMaxPage = -1

type ImageList struct {
	Items []*Image

	nextToken string

	pageSize int32

	project string
}

func (l *ImageList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ImageList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listImage(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListImage(ctx context.Context, project string) (*ImageList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListImageWithMaxResults(ctx, project, ImageMaxPage)

}

func (c *Client) ListImageWithMaxResults(ctx context.Context, project string, pageSize int32) (*ImageList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listImage(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ImageList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

// URLNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Image) URLNormalized() *Image {
	normalized := dcl.Copy(*r).(Image)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Family = dcl.SelfLinkToName(r.Family)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.SourceDisk = dcl.SelfLinkToName(r.SourceDisk)
	normalized.SourceDiskId = dcl.SelfLinkToName(r.SourceDiskId)
	normalized.SourceImage = dcl.SelfLinkToName(r.SourceImage)
	normalized.SourceImageId = dcl.SelfLinkToName(r.SourceImageId)
	normalized.SourceSnapshot = dcl.SelfLinkToName(r.SourceSnapshot)
	normalized.SourceSnapshotId = dcl.SelfLinkToName(r.SourceSnapshotId)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (c *Client) GetImage(ctx context.Context, r *Image) (*Image, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getImageRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalImage(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeImageNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteImage(ctx context.Context, r *Image) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Image resource is nil")
	}
	c.Config.Logger.Info("Deleting Image...")
	deleteOp := deleteImageOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllImage deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllImage(ctx context.Context, project string, filter func(*Image) bool) error {
	listObj, err := c.ListImage(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllImage(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllImage(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyImage(ctx context.Context, rawDesired *Image, opts ...dcl.ApplyOption) (*Image, error) {
	var resultNewState *Image
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyImageHelper(c, ctx, rawDesired, opts...)
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

func applyImageHelper(c *Client, ctx context.Context, rawDesired *Image, opts ...dcl.ApplyOption) (*Image, error) {
	c.Config.Logger.Info("Beginning ApplyImage...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.imageDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToImageDiffs(c.Config, fieldDiffs, opts)
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
				recreate = true
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []imageApiOperation
	if create {
		ops = append(ops, &createImageOperation{})
	} else if recreate {
		ops = append(ops, &deleteImageOperation{})
		ops = append(ops, &createImageOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeImageDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetImage(ctx, desired.URLNormalized())
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createImageOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapImage(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeImageNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeImageNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeImageDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffImage(c, newDesired, newState)
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
