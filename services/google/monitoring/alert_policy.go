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
package monitoring

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type AlertPolicy struct {
	Name                 *string                      `json:"name"`
	DisplayName          *string                      `json:"displayName"`
	Documentation        *AlertPolicyDocumentation    `json:"documentation"`
	UserLabels           map[string]string            `json:"userLabels"`
	Conditions           []AlertPolicyConditions      `json:"conditions"`
	Combiner             *AlertPolicyCombinerEnum     `json:"combiner"`
	Disabled             *bool                        `json:"disabled"`
	Enabled              *AlertPolicyEnabled          `json:"enabled"`
	Validity             *AlertPolicyValidity         `json:"validity"`
	NotificationChannels []string                     `json:"notificationChannels"`
	CreationRecord       *AlertPolicyCreationRecord   `json:"creationRecord"`
	MutationRecord       *AlertPolicyMutationRecord   `json:"mutationRecord"`
	IncidentStrategy     *AlertPolicyIncidentStrategy `json:"incidentStrategy"`
	Metadata             *AlertPolicyMetadata         `json:"metadata"`
	Project              *string                      `json:"project"`
}

func (r *AlertPolicy) String() string {
	return dcl.SprintResource(r)
}

// The enum AlertPolicyConditionsResourceStateFilterEnum.
type AlertPolicyConditionsResourceStateFilterEnum string

// AlertPolicyConditionsResourceStateFilterEnumRef returns a *AlertPolicyConditionsResourceStateFilterEnum with the value of string s
// If the empty string is provided, nil is returned.
func AlertPolicyConditionsResourceStateFilterEnumRef(s string) *AlertPolicyConditionsResourceStateFilterEnum {
	if s == "" {
		return nil
	}

	v := AlertPolicyConditionsResourceStateFilterEnum(s)
	return &v
}

func (v AlertPolicyConditionsResourceStateFilterEnum) Validate() error {
	for _, s := range []string{"RESOURCE_STATE_FILTER_UNSPECIFIED", "INCLUDE_INACTIVE", "EXCLUDE_INACTIVE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AlertPolicyConditionsResourceStateFilterEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum.
type AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum string

// AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnumRef returns a *AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnumRef(s string) *AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum {
	if s == "" {
		return nil
	}

	v := AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum(s)
	return &v
}

func (v AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum) Validate() error {
	for _, s := range []string{"ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_MAKE_DISTRIBUTION", "ALIGN_PERCENT_CHANGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum.
type AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum string

// AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnumRef returns a *AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnumRef(s string) *AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum {
	if s == "" {
		return nil
	}

	v := AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum(s)
	return &v
}

func (v AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum.
type AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum string

// AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnumRef returns a *AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnumRef(s string) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum {
	if s == "" {
		return nil
	}

	v := AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum(s)
	return &v
}

func (v AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum) Validate() error {
	for _, s := range []string{"ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_MAKE_DISTRIBUTION", "ALIGN_PERCENT_CHANGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum.
type AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum string

// AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnumRef returns a *AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnumRef(s string) *AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum {
	if s == "" {
		return nil
	}

	v := AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum(s)
	return &v
}

func (v AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AlertPolicyConditionsConditionThresholdComparisonEnum.
type AlertPolicyConditionsConditionThresholdComparisonEnum string

// AlertPolicyConditionsConditionThresholdComparisonEnumRef returns a *AlertPolicyConditionsConditionThresholdComparisonEnum with the value of string s
// If the empty string is provided, nil is returned.
func AlertPolicyConditionsConditionThresholdComparisonEnumRef(s string) *AlertPolicyConditionsConditionThresholdComparisonEnum {
	if s == "" {
		return nil
	}

	v := AlertPolicyConditionsConditionThresholdComparisonEnum(s)
	return &v
}

func (v AlertPolicyConditionsConditionThresholdComparisonEnum) Validate() error {
	for _, s := range []string{"COMPARISON_UNSPECIFIED", "COMPARISON_GT", "COMPARISON_GE", "COMPARISON_LT", "COMPARISON_LE", "COMPARISON_EQ", "COMPARISON_NE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AlertPolicyConditionsConditionThresholdComparisonEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum.
type AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum string

// AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnumRef returns a *AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnumRef(s string) *AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum {
	if s == "" {
		return nil
	}

	v := AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum(s)
	return &v
}

func (v AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum) Validate() error {
	for _, s := range []string{"ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_MAKE_DISTRIBUTION", "ALIGN_PERCENT_CHANGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum.
type AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum string

// AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnumRef returns a *AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnumRef(s string) *AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum {
	if s == "" {
		return nil
	}

	v := AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum(s)
	return &v
}

func (v AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum.
type AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum string

// AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnumRef returns a *AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnumRef(s string) *AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum {
	if s == "" {
		return nil
	}

	v := AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum(s)
	return &v
}

func (v AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum) Validate() error {
	for _, s := range []string{"ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_MAKE_DISTRIBUTION", "ALIGN_PERCENT_CHANGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum.
type AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum string

// AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnumRef returns a *AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnumRef(s string) *AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum {
	if s == "" {
		return nil
	}

	v := AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum(s)
	return &v
}

func (v AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AlertPolicyConditionsConditionRateComparisonEnum.
type AlertPolicyConditionsConditionRateComparisonEnum string

// AlertPolicyConditionsConditionRateComparisonEnumRef returns a *AlertPolicyConditionsConditionRateComparisonEnum with the value of string s
// If the empty string is provided, nil is returned.
func AlertPolicyConditionsConditionRateComparisonEnumRef(s string) *AlertPolicyConditionsConditionRateComparisonEnum {
	if s == "" {
		return nil
	}

	v := AlertPolicyConditionsConditionRateComparisonEnum(s)
	return &v
}

func (v AlertPolicyConditionsConditionRateComparisonEnum) Validate() error {
	for _, s := range []string{"COMPARISON_UNSPECIFIED", "COMPARISON_GT", "COMPARISON_GE", "COMPARISON_LT", "COMPARISON_LE", "COMPARISON_EQ", "COMPARISON_NE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AlertPolicyConditionsConditionRateComparisonEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AlertPolicyConditionsConditionProcessCountComparisonEnum.
type AlertPolicyConditionsConditionProcessCountComparisonEnum string

// AlertPolicyConditionsConditionProcessCountComparisonEnumRef returns a *AlertPolicyConditionsConditionProcessCountComparisonEnum with the value of string s
// If the empty string is provided, nil is returned.
func AlertPolicyConditionsConditionProcessCountComparisonEnumRef(s string) *AlertPolicyConditionsConditionProcessCountComparisonEnum {
	if s == "" {
		return nil
	}

	v := AlertPolicyConditionsConditionProcessCountComparisonEnum(s)
	return &v
}

func (v AlertPolicyConditionsConditionProcessCountComparisonEnum) Validate() error {
	for _, s := range []string{"COMPARISON_UNSPECIFIED", "COMPARISON_GT", "COMPARISON_GE", "COMPARISON_LT", "COMPARISON_LE", "COMPARISON_EQ", "COMPARISON_NE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AlertPolicyConditionsConditionProcessCountComparisonEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AlertPolicyCombinerEnum.
type AlertPolicyCombinerEnum string

// AlertPolicyCombinerEnumRef returns a *AlertPolicyCombinerEnum with the value of string s
// If the empty string is provided, nil is returned.
func AlertPolicyCombinerEnumRef(s string) *AlertPolicyCombinerEnum {
	if s == "" {
		return nil
	}

	v := AlertPolicyCombinerEnum(s)
	return &v
}

func (v AlertPolicyCombinerEnum) Validate() error {
	for _, s := range []string{"COMBINE_UNSPECIFIED", "AND", "OR", "AND_WITH_MATCHING_RESOURCE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AlertPolicyCombinerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum AlertPolicyIncidentStrategyTypeEnum.
type AlertPolicyIncidentStrategyTypeEnum string

// AlertPolicyIncidentStrategyTypeEnumRef returns a *AlertPolicyIncidentStrategyTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func AlertPolicyIncidentStrategyTypeEnumRef(s string) *AlertPolicyIncidentStrategyTypeEnum {
	if s == "" {
		return nil
	}

	v := AlertPolicyIncidentStrategyTypeEnum(s)
	return &v
}

func (v AlertPolicyIncidentStrategyTypeEnum) Validate() error {
	for _, s := range []string{"TYPE_UNSPECIFIED", "CUSTOM", "APP_ENGINE", "CLOUD_ENDPOINTS", "ISTIO", "ISTIO_SERVICE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "AlertPolicyIncidentStrategyTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type AlertPolicyDocumentation struct {
	empty    bool    `json:"-"`
	Content  *string `json:"content"`
	MimeType *string `json:"mimeType"`
}

type jsonAlertPolicyDocumentation AlertPolicyDocumentation

func (r *AlertPolicyDocumentation) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyDocumentation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyDocumentation
	} else {

		r.Content = res.Content

		r.MimeType = res.MimeType

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyDocumentation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyDocumentation *AlertPolicyDocumentation = &AlertPolicyDocumentation{empty: true}

func (r *AlertPolicyDocumentation) Empty() bool {
	return r.empty
}

func (r *AlertPolicyDocumentation) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyDocumentation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditions struct {
	empty                            bool                                                   `json:"-"`
	Name                             *string                                                `json:"name"`
	DisplayName                      *string                                                `json:"displayName"`
	ResourceStateFilter              *AlertPolicyConditionsResourceStateFilterEnum          `json:"resourceStateFilter"`
	ConditionThreshold               *AlertPolicyConditionsConditionThreshold               `json:"conditionThreshold"`
	ConditionAbsent                  *AlertPolicyConditionsConditionAbsent                  `json:"conditionAbsent"`
	ConditionMatchedLog              *AlertPolicyConditionsConditionMatchedLog              `json:"conditionMatchedLog"`
	ConditionClusterOutlier          *AlertPolicyConditionsConditionClusterOutlier          `json:"conditionClusterOutlier"`
	ConditionRate                    *AlertPolicyConditionsConditionRate                    `json:"conditionRate"`
	ConditionUpMon                   *AlertPolicyConditionsConditionUpMon                   `json:"conditionUpMon"`
	ConditionProcessCount            *AlertPolicyConditionsConditionProcessCount            `json:"conditionProcessCount"`
	ConditionTimeSeriesQueryLanguage *AlertPolicyConditionsConditionTimeSeriesQueryLanguage `json:"conditionTimeSeriesQueryLanguage"`
	ConditionMonitoringQueryLanguage *AlertPolicyConditionsConditionMonitoringQueryLanguage `json:"conditionMonitoringQueryLanguage"`
}

type jsonAlertPolicyConditions AlertPolicyConditions

func (r *AlertPolicyConditions) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditions
	} else {

		r.Name = res.Name

		r.DisplayName = res.DisplayName

		r.ResourceStateFilter = res.ResourceStateFilter

		r.ConditionThreshold = res.ConditionThreshold

		r.ConditionAbsent = res.ConditionAbsent

		r.ConditionMatchedLog = res.ConditionMatchedLog

		r.ConditionClusterOutlier = res.ConditionClusterOutlier

		r.ConditionRate = res.ConditionRate

		r.ConditionUpMon = res.ConditionUpMon

		r.ConditionProcessCount = res.ConditionProcessCount

		r.ConditionTimeSeriesQueryLanguage = res.ConditionTimeSeriesQueryLanguage

		r.ConditionMonitoringQueryLanguage = res.ConditionMonitoringQueryLanguage

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditions *AlertPolicyConditions = &AlertPolicyConditions{empty: true}

func (r *AlertPolicyConditions) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditions) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThreshold struct {
	empty                   bool                                                             `json:"-"`
	Filter                  *string                                                          `json:"filter"`
	Aggregations            []AlertPolicyConditionsConditionThresholdAggregations            `json:"aggregations"`
	DenominatorFilter       *string                                                          `json:"denominatorFilter"`
	DenominatorAggregations []AlertPolicyConditionsConditionThresholdDenominatorAggregations `json:"denominatorAggregations"`
	Comparison              *AlertPolicyConditionsConditionThresholdComparisonEnum           `json:"comparison"`
	ThresholdValue          *float64                                                         `json:"thresholdValue"`
	Duration                *string                                                          `json:"duration"`
	Trigger                 *AlertPolicyConditionsConditionThresholdTrigger                  `json:"trigger"`
}

type jsonAlertPolicyConditionsConditionThreshold AlertPolicyConditionsConditionThreshold

func (r *AlertPolicyConditionsConditionThreshold) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThreshold
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThreshold
	} else {

		r.Filter = res.Filter

		r.Aggregations = res.Aggregations

		r.DenominatorFilter = res.DenominatorFilter

		r.DenominatorAggregations = res.DenominatorAggregations

		r.Comparison = res.Comparison

		r.ThresholdValue = res.ThresholdValue

		r.Duration = res.Duration

		r.Trigger = res.Trigger

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThreshold is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThreshold *AlertPolicyConditionsConditionThreshold = &AlertPolicyConditionsConditionThreshold{empty: true}

func (r *AlertPolicyConditionsConditionThreshold) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThreshold) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThreshold) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdAggregations struct {
	empty                        bool                                                                             `json:"-"`
	AlignmentPeriod              *string                                                                          `json:"alignmentPeriod"`
	PerSeriesAligner             *AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum         `json:"perSeriesAligner"`
	CrossSeriesReducer           *AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum       `json:"crossSeriesReducer"`
	GroupByFields                []string                                                                         `json:"groupByFields"`
	ReduceFractionLessThanParams *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams `json:"reduceFractionLessThanParams"`
	ReduceMakeDistributionParams *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams `json:"reduceMakeDistributionParams"`
}

type jsonAlertPolicyConditionsConditionThresholdAggregations AlertPolicyConditionsConditionThresholdAggregations

func (r *AlertPolicyConditionsConditionThresholdAggregations) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdAggregations
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdAggregations
	} else {

		r.AlignmentPeriod = res.AlignmentPeriod

		r.PerSeriesAligner = res.PerSeriesAligner

		r.CrossSeriesReducer = res.CrossSeriesReducer

		r.GroupByFields = res.GroupByFields

		r.ReduceFractionLessThanParams = res.ReduceFractionLessThanParams

		r.ReduceMakeDistributionParams = res.ReduceMakeDistributionParams

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdAggregations is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdAggregations *AlertPolicyConditionsConditionThresholdAggregations = &AlertPolicyConditionsConditionThresholdAggregations{empty: true}

func (r *AlertPolicyConditionsConditionThresholdAggregations) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdAggregations) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdAggregations) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams struct {
	empty     bool     `json:"-"`
	Threshold *float64 `json:"threshold"`
}

type jsonAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams
	} else {

		r.Threshold = res.Threshold

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams = &AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams{empty: true}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams struct {
	empty            bool                                                                                             `json:"-"`
	BucketOptions    *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions    `json:"bucketOptions"`
	ExemplarSampling *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling `json:"exemplarSampling"`
}

type jsonAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams
	} else {

		r.BucketOptions = res.BucketOptions

		r.ExemplarSampling = res.ExemplarSampling

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams = &AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams{empty: true}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions struct {
	empty              bool                                                                                                            `json:"-"`
	LinearBuckets      *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

type jsonAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions
	} else {

		r.LinearBuckets = res.LinearBuckets

		r.ExponentialBuckets = res.ExponentialBuckets

		r.ExplicitBuckets = res.ExplicitBuckets

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions = &AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions{empty: true}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

type jsonAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.Width = res.Width

		r.Offset = res.Offset

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets = &AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{empty: true}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

type jsonAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.GrowthFactor = res.GrowthFactor

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets = &AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{empty: true}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

type jsonAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	} else {

		r.Bounds = res.Bounds

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets = &AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{empty: true}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling struct {
	empty        bool     `json:"-"`
	MinimumValue *float64 `json:"minimumValue"`
}

type jsonAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling
	} else {

		r.MinimumValue = res.MinimumValue

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling = &AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling{empty: true}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdDenominatorAggregations struct {
	empty                        bool                                                                                        `json:"-"`
	AlignmentPeriod              *string                                                                                     `json:"alignmentPeriod"`
	PerSeriesAligner             *AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum         `json:"perSeriesAligner"`
	CrossSeriesReducer           *AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum       `json:"crossSeriesReducer"`
	GroupByFields                []string                                                                                    `json:"groupByFields"`
	ReduceFractionLessThanParams *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams `json:"reduceFractionLessThanParams"`
	ReduceMakeDistributionParams *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams `json:"reduceMakeDistributionParams"`
}

type jsonAlertPolicyConditionsConditionThresholdDenominatorAggregations AlertPolicyConditionsConditionThresholdDenominatorAggregations

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregations) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdDenominatorAggregations
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregations
	} else {

		r.AlignmentPeriod = res.AlignmentPeriod

		r.PerSeriesAligner = res.PerSeriesAligner

		r.CrossSeriesReducer = res.CrossSeriesReducer

		r.GroupByFields = res.GroupByFields

		r.ReduceFractionLessThanParams = res.ReduceFractionLessThanParams

		r.ReduceMakeDistributionParams = res.ReduceMakeDistributionParams

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdDenominatorAggregations is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregations *AlertPolicyConditionsConditionThresholdDenominatorAggregations = &AlertPolicyConditionsConditionThresholdDenominatorAggregations{empty: true}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregations) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregations) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregations) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams struct {
	empty     bool     `json:"-"`
	Threshold *float64 `json:"threshold"`
}

type jsonAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams
	} else {

		r.Threshold = res.Threshold

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams = &AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams{empty: true}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams struct {
	empty            bool                                                                                                        `json:"-"`
	BucketOptions    *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions    `json:"bucketOptions"`
	ExemplarSampling *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling `json:"exemplarSampling"`
}

type jsonAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams
	} else {

		r.BucketOptions = res.BucketOptions

		r.ExemplarSampling = res.ExemplarSampling

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams = &AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams{empty: true}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions struct {
	empty              bool                                                                                                                       `json:"-"`
	LinearBuckets      *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

type jsonAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions
	} else {

		r.LinearBuckets = res.LinearBuckets

		r.ExponentialBuckets = res.ExponentialBuckets

		r.ExplicitBuckets = res.ExplicitBuckets

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions = &AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions{empty: true}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

type jsonAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.Width = res.Width

		r.Offset = res.Offset

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets = &AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{empty: true}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

type jsonAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.GrowthFactor = res.GrowthFactor

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets = &AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{empty: true}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

type jsonAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	} else {

		r.Bounds = res.Bounds

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets = &AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{empty: true}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling struct {
	empty        bool     `json:"-"`
	MinimumValue *float64 `json:"minimumValue"`
}

type jsonAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling
	} else {

		r.MinimumValue = res.MinimumValue

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling = &AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling{empty: true}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionThresholdTrigger struct {
	empty   bool     `json:"-"`
	Count   *int64   `json:"count"`
	Percent *float64 `json:"percent"`
}

type jsonAlertPolicyConditionsConditionThresholdTrigger AlertPolicyConditionsConditionThresholdTrigger

func (r *AlertPolicyConditionsConditionThresholdTrigger) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionThresholdTrigger
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionThresholdTrigger
	} else {

		r.Count = res.Count

		r.Percent = res.Percent

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionThresholdTrigger is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionThresholdTrigger *AlertPolicyConditionsConditionThresholdTrigger = &AlertPolicyConditionsConditionThresholdTrigger{empty: true}

func (r *AlertPolicyConditionsConditionThresholdTrigger) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionThresholdTrigger) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionThresholdTrigger) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionAbsent struct {
	empty        bool                                               `json:"-"`
	Filter       *string                                            `json:"filter"`
	Aggregations []AlertPolicyConditionsConditionAbsentAggregations `json:"aggregations"`
	Duration     *AlertPolicyConditionsConditionAbsentDuration      `json:"duration"`
	Trigger      *AlertPolicyConditionsConditionAbsentTrigger       `json:"trigger"`
}

type jsonAlertPolicyConditionsConditionAbsent AlertPolicyConditionsConditionAbsent

func (r *AlertPolicyConditionsConditionAbsent) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionAbsent
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionAbsent
	} else {

		r.Filter = res.Filter

		r.Aggregations = res.Aggregations

		r.Duration = res.Duration

		r.Trigger = res.Trigger

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionAbsent is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionAbsent *AlertPolicyConditionsConditionAbsent = &AlertPolicyConditionsConditionAbsent{empty: true}

func (r *AlertPolicyConditionsConditionAbsent) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionAbsent) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionAbsent) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionAbsentAggregations struct {
	empty                        bool                                                                          `json:"-"`
	AlignmentPeriod              *string                                                                       `json:"alignmentPeriod"`
	PerSeriesAligner             *AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum         `json:"perSeriesAligner"`
	CrossSeriesReducer           *AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum       `json:"crossSeriesReducer"`
	GroupByFields                []string                                                                      `json:"groupByFields"`
	ReduceFractionLessThanParams *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams `json:"reduceFractionLessThanParams"`
	ReduceMakeDistributionParams *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams `json:"reduceMakeDistributionParams"`
}

type jsonAlertPolicyConditionsConditionAbsentAggregations AlertPolicyConditionsConditionAbsentAggregations

func (r *AlertPolicyConditionsConditionAbsentAggregations) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionAbsentAggregations
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionAbsentAggregations
	} else {

		r.AlignmentPeriod = res.AlignmentPeriod

		r.PerSeriesAligner = res.PerSeriesAligner

		r.CrossSeriesReducer = res.CrossSeriesReducer

		r.GroupByFields = res.GroupByFields

		r.ReduceFractionLessThanParams = res.ReduceFractionLessThanParams

		r.ReduceMakeDistributionParams = res.ReduceMakeDistributionParams

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionAbsentAggregations is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionAbsentAggregations *AlertPolicyConditionsConditionAbsentAggregations = &AlertPolicyConditionsConditionAbsentAggregations{empty: true}

func (r *AlertPolicyConditionsConditionAbsentAggregations) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionAbsentAggregations) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionAbsentAggregations) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams struct {
	empty     bool     `json:"-"`
	Threshold *float64 `json:"threshold"`
}

type jsonAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams
	} else {

		r.Threshold = res.Threshold

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams = &AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams{empty: true}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams struct {
	empty            bool                                                                                          `json:"-"`
	BucketOptions    *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions    `json:"bucketOptions"`
	ExemplarSampling *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling `json:"exemplarSampling"`
}

type jsonAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams
	} else {

		r.BucketOptions = res.BucketOptions

		r.ExemplarSampling = res.ExemplarSampling

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams = &AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams{empty: true}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions struct {
	empty              bool                                                                                                         `json:"-"`
	LinearBuckets      *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

type jsonAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions
	} else {

		r.LinearBuckets = res.LinearBuckets

		r.ExponentialBuckets = res.ExponentialBuckets

		r.ExplicitBuckets = res.ExplicitBuckets

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions = &AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions{empty: true}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

type jsonAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.Width = res.Width

		r.Offset = res.Offset

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets = &AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{empty: true}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

type jsonAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.GrowthFactor = res.GrowthFactor

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets = &AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{empty: true}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

type jsonAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	} else {

		r.Bounds = res.Bounds

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets = &AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{empty: true}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling struct {
	empty        bool     `json:"-"`
	MinimumValue *float64 `json:"minimumValue"`
}

type jsonAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling
	} else {

		r.MinimumValue = res.MinimumValue

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling = &AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling{empty: true}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionAbsentDuration struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

type jsonAlertPolicyConditionsConditionAbsentDuration AlertPolicyConditionsConditionAbsentDuration

func (r *AlertPolicyConditionsConditionAbsentDuration) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionAbsentDuration
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionAbsentDuration
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionAbsentDuration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionAbsentDuration *AlertPolicyConditionsConditionAbsentDuration = &AlertPolicyConditionsConditionAbsentDuration{empty: true}

func (r *AlertPolicyConditionsConditionAbsentDuration) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionAbsentDuration) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionAbsentDuration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionAbsentTrigger struct {
	empty   bool     `json:"-"`
	Count   *int64   `json:"count"`
	Percent *float64 `json:"percent"`
}

type jsonAlertPolicyConditionsConditionAbsentTrigger AlertPolicyConditionsConditionAbsentTrigger

func (r *AlertPolicyConditionsConditionAbsentTrigger) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionAbsentTrigger
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionAbsentTrigger
	} else {

		r.Count = res.Count

		r.Percent = res.Percent

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionAbsentTrigger is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionAbsentTrigger *AlertPolicyConditionsConditionAbsentTrigger = &AlertPolicyConditionsConditionAbsentTrigger{empty: true}

func (r *AlertPolicyConditionsConditionAbsentTrigger) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionAbsentTrigger) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionAbsentTrigger) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionMatchedLog struct {
	empty           bool              `json:"-"`
	Filter          *string           `json:"filter"`
	LabelExtractors map[string]string `json:"labelExtractors"`
}

type jsonAlertPolicyConditionsConditionMatchedLog AlertPolicyConditionsConditionMatchedLog

func (r *AlertPolicyConditionsConditionMatchedLog) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionMatchedLog
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionMatchedLog
	} else {

		r.Filter = res.Filter

		r.LabelExtractors = res.LabelExtractors

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionMatchedLog is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionMatchedLog *AlertPolicyConditionsConditionMatchedLog = &AlertPolicyConditionsConditionMatchedLog{empty: true}

func (r *AlertPolicyConditionsConditionMatchedLog) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionMatchedLog) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionMatchedLog) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionClusterOutlier struct {
	empty  bool    `json:"-"`
	Filter *string `json:"filter"`
}

type jsonAlertPolicyConditionsConditionClusterOutlier AlertPolicyConditionsConditionClusterOutlier

func (r *AlertPolicyConditionsConditionClusterOutlier) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionClusterOutlier
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionClusterOutlier
	} else {

		r.Filter = res.Filter

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionClusterOutlier is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionClusterOutlier *AlertPolicyConditionsConditionClusterOutlier = &AlertPolicyConditionsConditionClusterOutlier{empty: true}

func (r *AlertPolicyConditionsConditionClusterOutlier) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionClusterOutlier) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionClusterOutlier) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionRate struct {
	empty          bool                                              `json:"-"`
	Filter         *string                                           `json:"filter"`
	Aggregations   []AlertPolicyConditionsConditionRateAggregations  `json:"aggregations"`
	Comparison     *AlertPolicyConditionsConditionRateComparisonEnum `json:"comparison"`
	ThresholdValue *float64                                          `json:"thresholdValue"`
	TimeWindow     *AlertPolicyConditionsConditionRateTimeWindow     `json:"timeWindow"`
	Trigger        *AlertPolicyConditionsConditionRateTrigger        `json:"trigger"`
}

type jsonAlertPolicyConditionsConditionRate AlertPolicyConditionsConditionRate

func (r *AlertPolicyConditionsConditionRate) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionRate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionRate
	} else {

		r.Filter = res.Filter

		r.Aggregations = res.Aggregations

		r.Comparison = res.Comparison

		r.ThresholdValue = res.ThresholdValue

		r.TimeWindow = res.TimeWindow

		r.Trigger = res.Trigger

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionRate is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionRate *AlertPolicyConditionsConditionRate = &AlertPolicyConditionsConditionRate{empty: true}

func (r *AlertPolicyConditionsConditionRate) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionRate) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionRate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionRateAggregations struct {
	empty                        bool                                                                        `json:"-"`
	AlignmentPeriod              *string                                                                     `json:"alignmentPeriod"`
	PerSeriesAligner             *AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum         `json:"perSeriesAligner"`
	CrossSeriesReducer           *AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum       `json:"crossSeriesReducer"`
	GroupByFields                []string                                                                    `json:"groupByFields"`
	ReduceFractionLessThanParams *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams `json:"reduceFractionLessThanParams"`
	ReduceMakeDistributionParams *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams `json:"reduceMakeDistributionParams"`
}

type jsonAlertPolicyConditionsConditionRateAggregations AlertPolicyConditionsConditionRateAggregations

func (r *AlertPolicyConditionsConditionRateAggregations) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionRateAggregations
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionRateAggregations
	} else {

		r.AlignmentPeriod = res.AlignmentPeriod

		r.PerSeriesAligner = res.PerSeriesAligner

		r.CrossSeriesReducer = res.CrossSeriesReducer

		r.GroupByFields = res.GroupByFields

		r.ReduceFractionLessThanParams = res.ReduceFractionLessThanParams

		r.ReduceMakeDistributionParams = res.ReduceMakeDistributionParams

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionRateAggregations is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionRateAggregations *AlertPolicyConditionsConditionRateAggregations = &AlertPolicyConditionsConditionRateAggregations{empty: true}

func (r *AlertPolicyConditionsConditionRateAggregations) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionRateAggregations) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionRateAggregations) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams struct {
	empty     bool     `json:"-"`
	Threshold *float64 `json:"threshold"`
}

type jsonAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams

func (r *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams
	} else {

		r.Threshold = res.Threshold

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams = &AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams{empty: true}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams struct {
	empty            bool                                                                                        `json:"-"`
	BucketOptions    *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions    `json:"bucketOptions"`
	ExemplarSampling *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling `json:"exemplarSampling"`
}

type jsonAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams
	} else {

		r.BucketOptions = res.BucketOptions

		r.ExemplarSampling = res.ExemplarSampling

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams = &AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams{empty: true}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions struct {
	empty              bool                                                                                                       `json:"-"`
	LinearBuckets      *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

type jsonAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions
	} else {

		r.LinearBuckets = res.LinearBuckets

		r.ExponentialBuckets = res.ExponentialBuckets

		r.ExplicitBuckets = res.ExplicitBuckets

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions = &AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions{empty: true}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

type jsonAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.Width = res.Width

		r.Offset = res.Offset

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets = &AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets{empty: true}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

type jsonAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.GrowthFactor = res.GrowthFactor

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets = &AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets{empty: true}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

type jsonAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	} else {

		r.Bounds = res.Bounds

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets = &AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets{empty: true}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling struct {
	empty        bool     `json:"-"`
	MinimumValue *float64 `json:"minimumValue"`
}

type jsonAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling
	} else {

		r.MinimumValue = res.MinimumValue

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling = &AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling{empty: true}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionRateTimeWindow struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

type jsonAlertPolicyConditionsConditionRateTimeWindow AlertPolicyConditionsConditionRateTimeWindow

func (r *AlertPolicyConditionsConditionRateTimeWindow) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionRateTimeWindow
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionRateTimeWindow
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionRateTimeWindow is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionRateTimeWindow *AlertPolicyConditionsConditionRateTimeWindow = &AlertPolicyConditionsConditionRateTimeWindow{empty: true}

func (r *AlertPolicyConditionsConditionRateTimeWindow) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionRateTimeWindow) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionRateTimeWindow) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionRateTrigger struct {
	empty   bool     `json:"-"`
	Count   *int64   `json:"count"`
	Percent *float64 `json:"percent"`
}

type jsonAlertPolicyConditionsConditionRateTrigger AlertPolicyConditionsConditionRateTrigger

func (r *AlertPolicyConditionsConditionRateTrigger) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionRateTrigger
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionRateTrigger
	} else {

		r.Count = res.Count

		r.Percent = res.Percent

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionRateTrigger is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionRateTrigger *AlertPolicyConditionsConditionRateTrigger = &AlertPolicyConditionsConditionRateTrigger{empty: true}

func (r *AlertPolicyConditionsConditionRateTrigger) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionRateTrigger) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionRateTrigger) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionUpMon struct {
	empty      bool                                         `json:"-"`
	Filter     *string                                      `json:"filter"`
	EndpointId *string                                      `json:"endpointId"`
	CheckId    *string                                      `json:"checkId"`
	Duration   *AlertPolicyConditionsConditionUpMonDuration `json:"duration"`
	Trigger    *AlertPolicyConditionsConditionUpMonTrigger  `json:"trigger"`
}

type jsonAlertPolicyConditionsConditionUpMon AlertPolicyConditionsConditionUpMon

func (r *AlertPolicyConditionsConditionUpMon) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionUpMon
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionUpMon
	} else {

		r.Filter = res.Filter

		r.EndpointId = res.EndpointId

		r.CheckId = res.CheckId

		r.Duration = res.Duration

		r.Trigger = res.Trigger

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionUpMon is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionUpMon *AlertPolicyConditionsConditionUpMon = &AlertPolicyConditionsConditionUpMon{empty: true}

func (r *AlertPolicyConditionsConditionUpMon) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionUpMon) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionUpMon) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionUpMonDuration struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

type jsonAlertPolicyConditionsConditionUpMonDuration AlertPolicyConditionsConditionUpMonDuration

func (r *AlertPolicyConditionsConditionUpMonDuration) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionUpMonDuration
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionUpMonDuration
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionUpMonDuration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionUpMonDuration *AlertPolicyConditionsConditionUpMonDuration = &AlertPolicyConditionsConditionUpMonDuration{empty: true}

func (r *AlertPolicyConditionsConditionUpMonDuration) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionUpMonDuration) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionUpMonDuration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionUpMonTrigger struct {
	empty   bool     `json:"-"`
	Count   *int64   `json:"count"`
	Percent *float64 `json:"percent"`
}

type jsonAlertPolicyConditionsConditionUpMonTrigger AlertPolicyConditionsConditionUpMonTrigger

func (r *AlertPolicyConditionsConditionUpMonTrigger) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionUpMonTrigger
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionUpMonTrigger
	} else {

		r.Count = res.Count

		r.Percent = res.Percent

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionUpMonTrigger is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionUpMonTrigger *AlertPolicyConditionsConditionUpMonTrigger = &AlertPolicyConditionsConditionUpMonTrigger{empty: true}

func (r *AlertPolicyConditionsConditionUpMonTrigger) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionUpMonTrigger) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionUpMonTrigger) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionProcessCount struct {
	empty                 bool                                                      `json:"-"`
	Process               *string                                                   `json:"process"`
	User                  *string                                                   `json:"user"`
	Filter                *string                                                   `json:"filter"`
	Comparison            *AlertPolicyConditionsConditionProcessCountComparisonEnum `json:"comparison"`
	ProcessCountThreshold *int64                                                    `json:"processCountThreshold"`
	Trigger               *AlertPolicyConditionsConditionProcessCountTrigger        `json:"trigger"`
	Duration              *AlertPolicyConditionsConditionProcessCountDuration       `json:"duration"`
}

type jsonAlertPolicyConditionsConditionProcessCount AlertPolicyConditionsConditionProcessCount

func (r *AlertPolicyConditionsConditionProcessCount) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionProcessCount
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionProcessCount
	} else {

		r.Process = res.Process

		r.User = res.User

		r.Filter = res.Filter

		r.Comparison = res.Comparison

		r.ProcessCountThreshold = res.ProcessCountThreshold

		r.Trigger = res.Trigger

		r.Duration = res.Duration

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionProcessCount is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionProcessCount *AlertPolicyConditionsConditionProcessCount = &AlertPolicyConditionsConditionProcessCount{empty: true}

func (r *AlertPolicyConditionsConditionProcessCount) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionProcessCount) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionProcessCount) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionProcessCountTrigger struct {
	empty   bool     `json:"-"`
	Count   *int64   `json:"count"`
	Percent *float64 `json:"percent"`
}

type jsonAlertPolicyConditionsConditionProcessCountTrigger AlertPolicyConditionsConditionProcessCountTrigger

func (r *AlertPolicyConditionsConditionProcessCountTrigger) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionProcessCountTrigger
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionProcessCountTrigger
	} else {

		r.Count = res.Count

		r.Percent = res.Percent

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionProcessCountTrigger is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionProcessCountTrigger *AlertPolicyConditionsConditionProcessCountTrigger = &AlertPolicyConditionsConditionProcessCountTrigger{empty: true}

func (r *AlertPolicyConditionsConditionProcessCountTrigger) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionProcessCountTrigger) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionProcessCountTrigger) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionProcessCountDuration struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

type jsonAlertPolicyConditionsConditionProcessCountDuration AlertPolicyConditionsConditionProcessCountDuration

func (r *AlertPolicyConditionsConditionProcessCountDuration) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionProcessCountDuration
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionProcessCountDuration
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionProcessCountDuration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionProcessCountDuration *AlertPolicyConditionsConditionProcessCountDuration = &AlertPolicyConditionsConditionProcessCountDuration{empty: true}

func (r *AlertPolicyConditionsConditionProcessCountDuration) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionProcessCountDuration) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionProcessCountDuration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionTimeSeriesQueryLanguage struct {
	empty   bool    `json:"-"`
	Query   *string `json:"query"`
	Summary *string `json:"summary"`
}

type jsonAlertPolicyConditionsConditionTimeSeriesQueryLanguage AlertPolicyConditionsConditionTimeSeriesQueryLanguage

func (r *AlertPolicyConditionsConditionTimeSeriesQueryLanguage) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionTimeSeriesQueryLanguage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionTimeSeriesQueryLanguage
	} else {

		r.Query = res.Query

		r.Summary = res.Summary

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionTimeSeriesQueryLanguage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionTimeSeriesQueryLanguage *AlertPolicyConditionsConditionTimeSeriesQueryLanguage = &AlertPolicyConditionsConditionTimeSeriesQueryLanguage{empty: true}

func (r *AlertPolicyConditionsConditionTimeSeriesQueryLanguage) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionTimeSeriesQueryLanguage) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionTimeSeriesQueryLanguage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionMonitoringQueryLanguage struct {
	empty    bool                                                           `json:"-"`
	Query    *string                                                        `json:"query"`
	Duration *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration `json:"duration"`
	Trigger  *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger  `json:"trigger"`
}

type jsonAlertPolicyConditionsConditionMonitoringQueryLanguage AlertPolicyConditionsConditionMonitoringQueryLanguage

func (r *AlertPolicyConditionsConditionMonitoringQueryLanguage) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionMonitoringQueryLanguage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionMonitoringQueryLanguage
	} else {

		r.Query = res.Query

		r.Duration = res.Duration

		r.Trigger = res.Trigger

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionMonitoringQueryLanguage is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionMonitoringQueryLanguage *AlertPolicyConditionsConditionMonitoringQueryLanguage = &AlertPolicyConditionsConditionMonitoringQueryLanguage{empty: true}

func (r *AlertPolicyConditionsConditionMonitoringQueryLanguage) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionMonitoringQueryLanguage) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionMonitoringQueryLanguage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionMonitoringQueryLanguageDuration struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

type jsonAlertPolicyConditionsConditionMonitoringQueryLanguageDuration AlertPolicyConditionsConditionMonitoringQueryLanguageDuration

func (r *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionMonitoringQueryLanguageDuration
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionMonitoringQueryLanguageDuration
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionMonitoringQueryLanguageDuration is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionMonitoringQueryLanguageDuration *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration = &AlertPolicyConditionsConditionMonitoringQueryLanguageDuration{empty: true}

func (r *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionMonitoringQueryLanguageDuration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger struct {
	empty   bool     `json:"-"`
	Count   *int64   `json:"count"`
	Percent *float64 `json:"percent"`
}

type jsonAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger

func (r *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger
	} else {

		r.Count = res.Count

		r.Percent = res.Percent

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger = &AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger{empty: true}

func (r *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) Empty() bool {
	return r.empty
}

func (r *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyEnabled struct {
	empty bool  `json:"-"`
	Value *bool `json:"value"`
}

type jsonAlertPolicyEnabled AlertPolicyEnabled

func (r *AlertPolicyEnabled) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyEnabled
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyEnabled
	} else {

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyEnabled is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyEnabled *AlertPolicyEnabled = &AlertPolicyEnabled{empty: true}

func (r *AlertPolicyEnabled) Empty() bool {
	return r.empty
}

func (r *AlertPolicyEnabled) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyEnabled) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyValidity struct {
	empty   bool                         `json:"-"`
	Code    *int64                       `json:"code"`
	Message *string                      `json:"message"`
	Details []AlertPolicyValidityDetails `json:"details"`
}

type jsonAlertPolicyValidity AlertPolicyValidity

func (r *AlertPolicyValidity) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyValidity
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyValidity
	} else {

		r.Code = res.Code

		r.Message = res.Message

		r.Details = res.Details

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyValidity is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyValidity *AlertPolicyValidity = &AlertPolicyValidity{empty: true}

func (r *AlertPolicyValidity) Empty() bool {
	return r.empty
}

func (r *AlertPolicyValidity) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyValidity) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyValidityDetails struct {
	empty   bool    `json:"-"`
	TypeUrl *string `json:"typeUrl"`
	Value   *string `json:"value"`
}

type jsonAlertPolicyValidityDetails AlertPolicyValidityDetails

func (r *AlertPolicyValidityDetails) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyValidityDetails
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyValidityDetails
	} else {

		r.TypeUrl = res.TypeUrl

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyValidityDetails is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyValidityDetails *AlertPolicyValidityDetails = &AlertPolicyValidityDetails{empty: true}

func (r *AlertPolicyValidityDetails) Empty() bool {
	return r.empty
}

func (r *AlertPolicyValidityDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyValidityDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyCreationRecord struct {
	empty      bool                                 `json:"-"`
	MutateTime *AlertPolicyCreationRecordMutateTime `json:"mutateTime"`
	MutatedBy  *string                              `json:"mutatedBy"`
}

type jsonAlertPolicyCreationRecord AlertPolicyCreationRecord

func (r *AlertPolicyCreationRecord) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyCreationRecord
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyCreationRecord
	} else {

		r.MutateTime = res.MutateTime

		r.MutatedBy = res.MutatedBy

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyCreationRecord is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyCreationRecord *AlertPolicyCreationRecord = &AlertPolicyCreationRecord{empty: true}

func (r *AlertPolicyCreationRecord) Empty() bool {
	return r.empty
}

func (r *AlertPolicyCreationRecord) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyCreationRecord) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyCreationRecordMutateTime struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

type jsonAlertPolicyCreationRecordMutateTime AlertPolicyCreationRecordMutateTime

func (r *AlertPolicyCreationRecordMutateTime) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyCreationRecordMutateTime
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyCreationRecordMutateTime
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyCreationRecordMutateTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyCreationRecordMutateTime *AlertPolicyCreationRecordMutateTime = &AlertPolicyCreationRecordMutateTime{empty: true}

func (r *AlertPolicyCreationRecordMutateTime) Empty() bool {
	return r.empty
}

func (r *AlertPolicyCreationRecordMutateTime) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyCreationRecordMutateTime) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyMutationRecord struct {
	empty      bool                                 `json:"-"`
	MutateTime *AlertPolicyMutationRecordMutateTime `json:"mutateTime"`
	MutatedBy  *string                              `json:"mutatedBy"`
}

type jsonAlertPolicyMutationRecord AlertPolicyMutationRecord

func (r *AlertPolicyMutationRecord) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyMutationRecord
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyMutationRecord
	} else {

		r.MutateTime = res.MutateTime

		r.MutatedBy = res.MutatedBy

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyMutationRecord is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyMutationRecord *AlertPolicyMutationRecord = &AlertPolicyMutationRecord{empty: true}

func (r *AlertPolicyMutationRecord) Empty() bool {
	return r.empty
}

func (r *AlertPolicyMutationRecord) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyMutationRecord) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyMutationRecordMutateTime struct {
	empty   bool   `json:"-"`
	Seconds *int64 `json:"seconds"`
	Nanos   *int64 `json:"nanos"`
}

type jsonAlertPolicyMutationRecordMutateTime AlertPolicyMutationRecordMutateTime

func (r *AlertPolicyMutationRecordMutateTime) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyMutationRecordMutateTime
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyMutationRecordMutateTime
	} else {

		r.Seconds = res.Seconds

		r.Nanos = res.Nanos

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyMutationRecordMutateTime is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyMutationRecordMutateTime *AlertPolicyMutationRecordMutateTime = &AlertPolicyMutationRecordMutateTime{empty: true}

func (r *AlertPolicyMutationRecordMutateTime) Empty() bool {
	return r.empty
}

func (r *AlertPolicyMutationRecordMutateTime) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyMutationRecordMutateTime) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyIncidentStrategy struct {
	empty bool                                 `json:"-"`
	Type  *AlertPolicyIncidentStrategyTypeEnum `json:"type"`
}

type jsonAlertPolicyIncidentStrategy AlertPolicyIncidentStrategy

func (r *AlertPolicyIncidentStrategy) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyIncidentStrategy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyIncidentStrategy
	} else {

		r.Type = res.Type

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyIncidentStrategy is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyIncidentStrategy *AlertPolicyIncidentStrategy = &AlertPolicyIncidentStrategy{empty: true}

func (r *AlertPolicyIncidentStrategy) Empty() bool {
	return r.empty
}

func (r *AlertPolicyIncidentStrategy) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyIncidentStrategy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type AlertPolicyMetadata struct {
	empty    bool     `json:"-"`
	SloNames []string `json:"sloNames"`
}

type jsonAlertPolicyMetadata AlertPolicyMetadata

func (r *AlertPolicyMetadata) UnmarshalJSON(data []byte) error {
	var res jsonAlertPolicyMetadata
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyAlertPolicyMetadata
	} else {

		r.SloNames = res.SloNames

	}
	return nil
}

// This object is used to assert a desired state where this AlertPolicyMetadata is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyAlertPolicyMetadata *AlertPolicyMetadata = &AlertPolicyMetadata{empty: true}

func (r *AlertPolicyMetadata) Empty() bool {
	return r.empty
}

func (r *AlertPolicyMetadata) String() string {
	return dcl.SprintResource(r)
}

func (r *AlertPolicyMetadata) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *AlertPolicy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "monitoring",
		Type:    "AlertPolicy",
		Version: "monitoring",
	}
}

const AlertPolicyMaxPage = -1

type AlertPolicyList struct {
	Items []*AlertPolicy

	nextToken string

	pageSize int32

	project string
}

func (l *AlertPolicyList) HasNext() bool {
	return l.nextToken != ""
}

func (l *AlertPolicyList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listAlertPolicy(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListAlertPolicy(ctx context.Context, project string) (*AlertPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListAlertPolicyWithMaxResults(ctx, project, AlertPolicyMaxPage)

}

func (c *Client) ListAlertPolicyWithMaxResults(ctx context.Context, project string, pageSize int32) (*AlertPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	items, token, err := c.listAlertPolicy(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &AlertPolicyList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetAlertPolicy(ctx context.Context, r *AlertPolicy) (*AlertPolicy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getAlertPolicyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalAlertPolicy(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeAlertPolicyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteAlertPolicy(ctx context.Context, r *AlertPolicy) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("AlertPolicy resource is nil")
	}
	c.Config.Logger.Info("Deleting AlertPolicy...")
	deleteOp := deleteAlertPolicyOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllAlertPolicy deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllAlertPolicy(ctx context.Context, project string, filter func(*AlertPolicy) bool) error {
	listObj, err := c.ListAlertPolicy(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllAlertPolicy(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllAlertPolicy(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyAlertPolicy(ctx context.Context, rawDesired *AlertPolicy, opts ...dcl.ApplyOption) (*AlertPolicy, error) {

	var resultNewState *AlertPolicy
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyAlertPolicyHelper(c, ctx, rawDesired, opts...)
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

func applyAlertPolicyHelper(c *Client, ctx context.Context, rawDesired *AlertPolicy, opts ...dcl.ApplyOption) (*AlertPolicy, error) {
	c.Config.Logger.Info("Beginning ApplyAlertPolicy...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.alertPolicyDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []alertPolicyApiOperation
	if create {
		ops = append(ops, &createAlertPolicyOperation{})
	} else if recreate {

		ops = append(ops, &deleteAlertPolicyOperation{})

		ops = append(ops, &createAlertPolicyOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeAlertPolicyDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetAlertPolicy(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createAlertPolicyOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapAlertPolicy(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeAlertPolicyNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeAlertPolicyNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeAlertPolicyDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffAlertPolicy(c, newDesired, newState)
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
