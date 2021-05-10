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

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Dashboard struct {
	Name         *string                `json:"name"`
	DisplayName  *string                `json:"displayName"`
	GridLayout   *DashboardGridLayout   `json:"gridLayout"`
	MosaicLayout *DashboardMosaicLayout `json:"mosaicLayout"`
	RowLayout    *DashboardRowLayout    `json:"rowLayout"`
	ColumnLayout *DashboardColumnLayout `json:"columnLayout"`
	Project      *string                `json:"project"`
}

func (r *Dashboard) String() string {
	return dcl.SprintResource(r)
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) Validate() error {
	for _, s := range []string{"ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_MAKE_DISTRIBUTION", "ALIGN_PERCENT_CHANGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) Validate() error {
	for _, s := range []string{"ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_MAKE_DISTRIBUTION", "ALIGN_PERCENT_CHANGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) Validate() error {
	for _, s := range []string{"METHOD_UNSPECIFIED", "METHOD_MEAN", "METHOD_MAX", "METHOD_MIN", "METHOD_SUM", "METHOD_LATEST"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) Validate() error {
	for _, s := range []string{"DIRECTION_UNSPECIFIED", "TOP", "BOTTOM"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) Validate() error {
	for _, s := range []string{"ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_MAKE_DISTRIBUTION", "ALIGN_PERCENT_CHANGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) Validate() error {
	for _, s := range []string{"ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_MAKE_DISTRIBUTION", "ALIGN_PERCENT_CHANGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) Validate() error {
	for _, s := range []string{"ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_MAKE_DISTRIBUTION", "ALIGN_PERCENT_CHANGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) Validate() error {
	for _, s := range []string{"METHOD_UNSPECIFIED", "METHOD_MEAN", "METHOD_MAX", "METHOD_MIN", "METHOD_SUM", "METHOD_LATEST"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) Validate() error {
	for _, s := range []string{"DIRECTION_UNSPECIFIED", "TOP", "BOTTOM"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum) Validate() error {
	for _, s := range []string{"DEFAULT_CLOUD", "AUTO"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartDataSetsPlotTypeEnum.
type DashboardWidgetXyChartDataSetsPlotTypeEnum string

// DashboardWidgetXyChartDataSetsPlotTypeEnumRef returns a *DashboardWidgetXyChartDataSetsPlotTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsPlotTypeEnumRef(s string) *DashboardWidgetXyChartDataSetsPlotTypeEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartDataSetsPlotTypeEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsPlotTypeEnum) Validate() error {
	for _, s := range []string{"PLOT_TYPE_UNSPECIFIED", "LINE", "STACKED_AREA", "STACKED_BAR", "HEATMAP"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartDataSetsPlotTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum.
type DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum string

// DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnumRef returns a *DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnumRef(s string) *DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum(s)
	return &v
}

func (v DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum) Validate() error {
	for _, s := range []string{"LOGICAL_OPERATOR_UNSPECIFIED", "AND", "OR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum.
type DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum string

// DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnumRef returns a *DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnumRef(s string) *DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum(s)
	return &v
}

func (v DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum) Validate() error {
	for _, s := range []string{"COMPARATOR_UNSPECIFIED", "EQUALS", "NOT_EQUALS", "CONTAINS", "NOT_CONTAINS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum.
type DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum string

// DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnumRef returns a *DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnumRef(s string) *DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum(s)
	return &v
}

func (v DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum) Validate() error {
	for _, s := range []string{"LOGICAL_OPERATOR_UNSPECIFIED", "AND", "OR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum.
type DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum string

// DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnumRef returns a *DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnumRef(s string) *DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum(s)
	return &v
}

func (v DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum) Validate() error {
	for _, s := range []string{"COMPARATOR_UNSPECIFIED", "EQUALS", "NOT_EQUALS", "CONTAINS", "NOT_CONTAINS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum.
type DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum string

// DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnumRef returns a *DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnumRef(s string) *DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum(s)
	return &v
}

func (v DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum) Validate() error {
	for _, s := range []string{"LOGICAL_OPERATOR_UNSPECIFIED", "AND", "OR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum.
type DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum string

// DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnumRef returns a *DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnumRef(s string) *DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum(s)
	return &v
}

func (v DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum) Validate() error {
	for _, s := range []string{"COMPARATOR_UNSPECIFIED", "EQUALS", "NOT_EQUALS", "CONTAINS", "NOT_CONTAINS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnum.
type DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnum string

// DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnumRef returns a *DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnumRef(s string) *DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnum(s)
	return &v
}

func (v DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnum) Validate() error {
	for _, s := range []string{"TYPE_UNSPECIFIED", "CUSTOM", "APP_ENGINE", "CLOUD_ENDPOINTS", "ISTIO", "ISTIO_SERVICE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum.
type DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum string

// DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnumRef returns a *DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnumRef(s string) *DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum(s)
	return &v
}

func (v DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum) Validate() error {
	for _, s := range []string{"LOGICAL_OPERATOR_UNSPECIFIED", "AND", "OR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum.
type DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum string

// DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnumRef returns a *DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnumRef(s string) *DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum(s)
	return &v
}

func (v DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum) Validate() error {
	for _, s := range []string{"COMPARATOR_UNSPECIFIED", "EQUALS", "NOT_EQUALS", "CONTAINS", "NOT_CONTAINS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum.
type DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum string

// DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnumRef returns a *DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnumRef(s string) *DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartThresholdsColorEnum.
type DashboardWidgetXyChartThresholdsColorEnum string

// DashboardWidgetXyChartThresholdsColorEnumRef returns a *DashboardWidgetXyChartThresholdsColorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartThresholdsColorEnumRef(s string) *DashboardWidgetXyChartThresholdsColorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartThresholdsColorEnum(s)
	return &v
}

func (v DashboardWidgetXyChartThresholdsColorEnum) Validate() error {
	for _, s := range []string{"COLOR_UNSPECIFIED", "GREY", "BLUE", "GREEN", "YELLOW", "ORANGE", "RED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartThresholdsColorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartThresholdsDirectionEnum.
type DashboardWidgetXyChartThresholdsDirectionEnum string

// DashboardWidgetXyChartThresholdsDirectionEnumRef returns a *DashboardWidgetXyChartThresholdsDirectionEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartThresholdsDirectionEnumRef(s string) *DashboardWidgetXyChartThresholdsDirectionEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartThresholdsDirectionEnum(s)
	return &v
}

func (v DashboardWidgetXyChartThresholdsDirectionEnum) Validate() error {
	for _, s := range []string{"DIRECTION_UNSPECIFIED", "ABOVE", "BELOW"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartThresholdsDirectionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartXAxisScaleEnum.
type DashboardWidgetXyChartXAxisScaleEnum string

// DashboardWidgetXyChartXAxisScaleEnumRef returns a *DashboardWidgetXyChartXAxisScaleEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartXAxisScaleEnumRef(s string) *DashboardWidgetXyChartXAxisScaleEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartXAxisScaleEnum(s)
	return &v
}

func (v DashboardWidgetXyChartXAxisScaleEnum) Validate() error {
	for _, s := range []string{"SCALE_UNSPECIFIED", "LINEAR", "LOG10"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartXAxisScaleEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartYAxisScaleEnum.
type DashboardWidgetXyChartYAxisScaleEnum string

// DashboardWidgetXyChartYAxisScaleEnumRef returns a *DashboardWidgetXyChartYAxisScaleEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartYAxisScaleEnumRef(s string) *DashboardWidgetXyChartYAxisScaleEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartYAxisScaleEnum(s)
	return &v
}

func (v DashboardWidgetXyChartYAxisScaleEnum) Validate() error {
	for _, s := range []string{"SCALE_UNSPECIFIED", "LINEAR", "LOG10"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartYAxisScaleEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetXyChartChartOptionsModeEnum.
type DashboardWidgetXyChartChartOptionsModeEnum string

// DashboardWidgetXyChartChartOptionsModeEnumRef returns a *DashboardWidgetXyChartChartOptionsModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartChartOptionsModeEnumRef(s string) *DashboardWidgetXyChartChartOptionsModeEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetXyChartChartOptionsModeEnum(s)
	return &v
}

func (v DashboardWidgetXyChartChartOptionsModeEnum) Validate() error {
	for _, s := range []string{"MODE_UNSPECIFIED", "COLOR", "X_RAY", "STATS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetXyChartChartOptionsModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.
type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum string

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) Validate() error {
	for _, s := range []string{"ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_MAKE_DISTRIBUTION", "ALIGN_PERCENT_CHANGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum.
type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum string

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum.
type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum string

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) Validate() error {
	for _, s := range []string{"ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_MAKE_DISTRIBUTION", "ALIGN_PERCENT_CHANGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum.
type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum string

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum.
type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum string

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) Validate() error {
	for _, s := range []string{"METHOD_UNSPECIFIED", "METHOD_MEAN", "METHOD_MAX", "METHOD_MIN", "METHOD_SUM", "METHOD_LATEST"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum.
type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum string

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) Validate() error {
	for _, s := range []string{"DIRECTION_UNSPECIFIED", "TOP", "BOTTOM"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum.
type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum string

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) Validate() error {
	for _, s := range []string{"ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_MAKE_DISTRIBUTION", "ALIGN_PERCENT_CHANGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum.
type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum string

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum.
type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum string

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) Validate() error {
	for _, s := range []string{"ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_MAKE_DISTRIBUTION", "ALIGN_PERCENT_CHANGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum.
type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum string

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum.
type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum string

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) Validate() error {
	for _, s := range []string{"ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_MAKE_DISTRIBUTION", "ALIGN_PERCENT_CHANGE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum.
type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum string

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum.
type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum string

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) Validate() error {
	for _, s := range []string{"METHOD_UNSPECIFIED", "METHOD_MEAN", "METHOD_MAX", "METHOD_MIN", "METHOD_SUM", "METHOD_LATEST"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum.
type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum string

// DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) Validate() error {
	for _, s := range []string{"DIRECTION_UNSPECIFIED", "TOP", "BOTTOM"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum.
type DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum string

// DashboardWidgetScorecardTimeSeriesQueryApiSourceEnumRef returns a *DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardTimeSeriesQueryApiSourceEnumRef(s string) *DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum) Validate() error {
	for _, s := range []string{"DEFAULT_CLOUD", "AUTO"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum.
type DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum string

// DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnumRef returns a *DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnumRef(s string) *DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum(s)
	return &v
}

func (v DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum) Validate() error {
	for _, s := range []string{"LOGICAL_OPERATOR_UNSPECIFIED", "AND", "OR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum.
type DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum string

// DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnumRef returns a *DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnumRef(s string) *DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum(s)
	return &v
}

func (v DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum) Validate() error {
	for _, s := range []string{"COMPARATOR_UNSPECIFIED", "EQUALS", "NOT_EQUALS", "CONTAINS", "NOT_CONTAINS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum.
type DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum string

// DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnumRef returns a *DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnumRef(s string) *DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum(s)
	return &v
}

func (v DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum) Validate() error {
	for _, s := range []string{"LOGICAL_OPERATOR_UNSPECIFIED", "AND", "OR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum.
type DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum string

// DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnumRef returns a *DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnumRef(s string) *DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum(s)
	return &v
}

func (v DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum) Validate() error {
	for _, s := range []string{"COMPARATOR_UNSPECIFIED", "EQUALS", "NOT_EQUALS", "CONTAINS", "NOT_CONTAINS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum.
type DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum string

// DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnumRef returns a *DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnumRef(s string) *DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum(s)
	return &v
}

func (v DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum) Validate() error {
	for _, s := range []string{"LOGICAL_OPERATOR_UNSPECIFIED", "AND", "OR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum.
type DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum string

// DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnumRef returns a *DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnumRef(s string) *DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum(s)
	return &v
}

func (v DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum) Validate() error {
	for _, s := range []string{"COMPARATOR_UNSPECIFIED", "EQUALS", "NOT_EQUALS", "CONTAINS", "NOT_CONTAINS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnum.
type DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnum string

// DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnumRef returns a *DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnumRef(s string) *DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnum(s)
	return &v
}

func (v DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnum) Validate() error {
	for _, s := range []string{"TYPE_UNSPECIFIED", "CUSTOM", "APP_ENGINE", "CLOUD_ENDPOINTS", "ISTIO", "ISTIO_SERVICE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum.
type DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum string

// DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnumRef returns a *DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnumRef(s string) *DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum(s)
	return &v
}

func (v DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum) Validate() error {
	for _, s := range []string{"LOGICAL_OPERATOR_UNSPECIFIED", "AND", "OR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum.
type DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum string

// DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnumRef returns a *DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnumRef(s string) *DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum(s)
	return &v
}

func (v DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum) Validate() error {
	for _, s := range []string{"COMPARATOR_UNSPECIFIED", "EQUALS", "NOT_EQUALS", "CONTAINS", "NOT_CONTAINS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum.
type DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum string

// DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnumRef returns a *DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnumRef(s string) *DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum) Validate() error {
	for _, s := range []string{"REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05", "REDUCE_FRACTION_LESS_THAN", "REDUCE_MAKE_DISTRIBUTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum.
type DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum string

// DashboardWidgetScorecardSparkChartViewSparkChartTypeEnumRef returns a *DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardSparkChartViewSparkChartTypeEnumRef(s string) *DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(s)
	return &v
}

func (v DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum) Validate() error {
	for _, s := range []string{"SPARK_CHART_TYPE_UNSPECIFIED", "SPARK_LINE", "SPARK_BAR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardThresholdsColorEnum.
type DashboardWidgetScorecardThresholdsColorEnum string

// DashboardWidgetScorecardThresholdsColorEnumRef returns a *DashboardWidgetScorecardThresholdsColorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardThresholdsColorEnumRef(s string) *DashboardWidgetScorecardThresholdsColorEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardThresholdsColorEnum(s)
	return &v
}

func (v DashboardWidgetScorecardThresholdsColorEnum) Validate() error {
	for _, s := range []string{"COLOR_UNSPECIFIED", "GREY", "BLUE", "GREEN", "YELLOW", "ORANGE", "RED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardThresholdsColorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetScorecardThresholdsDirectionEnum.
type DashboardWidgetScorecardThresholdsDirectionEnum string

// DashboardWidgetScorecardThresholdsDirectionEnumRef returns a *DashboardWidgetScorecardThresholdsDirectionEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardThresholdsDirectionEnumRef(s string) *DashboardWidgetScorecardThresholdsDirectionEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetScorecardThresholdsDirectionEnum(s)
	return &v
}

func (v DashboardWidgetScorecardThresholdsDirectionEnum) Validate() error {
	for _, s := range []string{"DIRECTION_UNSPECIFIED", "ABOVE", "BELOW"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetScorecardThresholdsDirectionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum DashboardWidgetTextFormatEnum.
type DashboardWidgetTextFormatEnum string

// DashboardWidgetTextFormatEnumRef returns a *DashboardWidgetTextFormatEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetTextFormatEnumRef(s string) *DashboardWidgetTextFormatEnum {
	if s == "" {
		return nil
	}

	v := DashboardWidgetTextFormatEnum(s)
	return &v
}

func (v DashboardWidgetTextFormatEnum) Validate() error {
	for _, s := range []string{"FORMAT_UNSPECIFIED", "MARKDOWN", "RAW"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "DashboardWidgetTextFormatEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type DashboardGridLayout struct {
	empty   bool              `json:"-"`
	Columns *int64            `json:"columns"`
	Widgets []DashboardWidget `json:"widgets"`
}

type jsonDashboardGridLayout DashboardGridLayout

func (r *DashboardGridLayout) UnmarshalJSON(data []byte) error {
	var res jsonDashboardGridLayout
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardGridLayout
	} else {

		r.Columns = res.Columns

		r.Widgets = res.Widgets

	}
	return nil
}

// This object is used to assert a desired state where this DashboardGridLayout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardGridLayout *DashboardGridLayout = &DashboardGridLayout{empty: true}

func (r *DashboardGridLayout) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardGridLayout) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardMosaicLayout struct {
	empty   bool                         `json:"-"`
	Columns *int64                       `json:"columns"`
	Tiles   []DashboardMosaicLayoutTiles `json:"tiles"`
}

type jsonDashboardMosaicLayout DashboardMosaicLayout

func (r *DashboardMosaicLayout) UnmarshalJSON(data []byte) error {
	var res jsonDashboardMosaicLayout
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardMosaicLayout
	} else {

		r.Columns = res.Columns

		r.Tiles = res.Tiles

	}
	return nil
}

// This object is used to assert a desired state where this DashboardMosaicLayout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardMosaicLayout *DashboardMosaicLayout = &DashboardMosaicLayout{empty: true}

func (r *DashboardMosaicLayout) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardMosaicLayout) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardMosaicLayoutTiles struct {
	empty  bool             `json:"-"`
	XPos   *int64           `json:"xPos"`
	YPos   *int64           `json:"yPos"`
	Width  *int64           `json:"width"`
	Height *int64           `json:"height"`
	Widget *DashboardWidget `json:"widget"`
}

type jsonDashboardMosaicLayoutTiles DashboardMosaicLayoutTiles

func (r *DashboardMosaicLayoutTiles) UnmarshalJSON(data []byte) error {
	var res jsonDashboardMosaicLayoutTiles
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardMosaicLayoutTiles
	} else {

		r.XPos = res.XPos

		r.YPos = res.YPos

		r.Width = res.Width

		r.Height = res.Height

		r.Widget = res.Widget

	}
	return nil
}

// This object is used to assert a desired state where this DashboardMosaicLayoutTiles is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardMosaicLayoutTiles *DashboardMosaicLayoutTiles = &DashboardMosaicLayoutTiles{empty: true}

func (r *DashboardMosaicLayoutTiles) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardMosaicLayoutTiles) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardRowLayout struct {
	empty bool                     `json:"-"`
	Rows  []DashboardRowLayoutRows `json:"rows"`
}

type jsonDashboardRowLayout DashboardRowLayout

func (r *DashboardRowLayout) UnmarshalJSON(data []byte) error {
	var res jsonDashboardRowLayout
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardRowLayout
	} else {

		r.Rows = res.Rows

	}
	return nil
}

// This object is used to assert a desired state where this DashboardRowLayout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardRowLayout *DashboardRowLayout = &DashboardRowLayout{empty: true}

func (r *DashboardRowLayout) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardRowLayout) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardRowLayoutRows struct {
	empty   bool              `json:"-"`
	Weight  *int64            `json:"weight"`
	Widgets []DashboardWidget `json:"widgets"`
}

type jsonDashboardRowLayoutRows DashboardRowLayoutRows

func (r *DashboardRowLayoutRows) UnmarshalJSON(data []byte) error {
	var res jsonDashboardRowLayoutRows
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardRowLayoutRows
	} else {

		r.Weight = res.Weight

		r.Widgets = res.Widgets

	}
	return nil
}

// This object is used to assert a desired state where this DashboardRowLayoutRows is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardRowLayoutRows *DashboardRowLayoutRows = &DashboardRowLayoutRows{empty: true}

func (r *DashboardRowLayoutRows) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardRowLayoutRows) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardColumnLayout struct {
	empty   bool                           `json:"-"`
	Columns []DashboardColumnLayoutColumns `json:"columns"`
}

type jsonDashboardColumnLayout DashboardColumnLayout

func (r *DashboardColumnLayout) UnmarshalJSON(data []byte) error {
	var res jsonDashboardColumnLayout
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardColumnLayout
	} else {

		r.Columns = res.Columns

	}
	return nil
}

// This object is used to assert a desired state where this DashboardColumnLayout is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardColumnLayout *DashboardColumnLayout = &DashboardColumnLayout{empty: true}

func (r *DashboardColumnLayout) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardColumnLayout) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardColumnLayoutColumns struct {
	empty   bool              `json:"-"`
	Weight  *int64            `json:"weight"`
	Widgets []DashboardWidget `json:"widgets"`
}

type jsonDashboardColumnLayoutColumns DashboardColumnLayoutColumns

func (r *DashboardColumnLayoutColumns) UnmarshalJSON(data []byte) error {
	var res jsonDashboardColumnLayoutColumns
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardColumnLayoutColumns
	} else {

		r.Weight = res.Weight

		r.Widgets = res.Widgets

	}
	return nil
}

// This object is used to assert a desired state where this DashboardColumnLayoutColumns is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardColumnLayoutColumns *DashboardColumnLayoutColumns = &DashboardColumnLayoutColumns{empty: true}

func (r *DashboardColumnLayoutColumns) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardColumnLayoutColumns) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidget struct {
	empty     bool                      `json:"-"`
	Title     *string                   `json:"title"`
	XyChart   *DashboardWidgetXyChart   `json:"xyChart"`
	Scorecard *DashboardWidgetScorecard `json:"scorecard"`
	Text      *DashboardWidgetText      `json:"text"`
	Blank     *DashboardWidgetBlank     `json:"blank"`
}

type jsonDashboardWidget DashboardWidget

func (r *DashboardWidget) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidget
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidget
	} else {

		r.Title = res.Title

		r.XyChart = res.XyChart

		r.Scorecard = res.Scorecard

		r.Text = res.Text

		r.Blank = res.Blank

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidget is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidget *DashboardWidget = &DashboardWidget{empty: true}

func (r *DashboardWidget) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidget) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChart struct {
	empty             bool                                   `json:"-"`
	DataSets          []DashboardWidgetXyChartDataSets       `json:"dataSets"`
	SourceDrilldown   *DashboardWidgetXyChartSourceDrilldown `json:"sourceDrilldown"`
	MetricDrilldown   *DashboardWidgetXyChartMetricDrilldown `json:"metricDrilldown"`
	TimeshiftDuration *string                                `json:"timeshiftDuration"`
	Thresholds        []DashboardWidgetXyChartThresholds     `json:"thresholds"`
	XAxis             *DashboardWidgetXyChartXAxis           `json:"xAxis"`
	YAxis             *DashboardWidgetXyChartYAxis           `json:"yAxis"`
	ChartOptions      *DashboardWidgetXyChartChartOptions    `json:"chartOptions"`
}

type jsonDashboardWidgetXyChart DashboardWidgetXyChart

func (r *DashboardWidgetXyChart) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChart
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChart
	} else {

		r.DataSets = res.DataSets

		r.SourceDrilldown = res.SourceDrilldown

		r.MetricDrilldown = res.MetricDrilldown

		r.TimeshiftDuration = res.TimeshiftDuration

		r.Thresholds = res.Thresholds

		r.XAxis = res.XAxis

		r.YAxis = res.YAxis

		r.ChartOptions = res.ChartOptions

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChart is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChart *DashboardWidgetXyChart = &DashboardWidgetXyChart{empty: true}

func (r *DashboardWidgetXyChart) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChart) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSets struct {
	empty              bool                                           `json:"-"`
	TimeSeriesQuery    *DashboardWidgetXyChartDataSetsTimeSeriesQuery `json:"timeSeriesQuery"`
	PlotType           *DashboardWidgetXyChartDataSetsPlotTypeEnum    `json:"plotType"`
	LegendTemplate     *string                                        `json:"legendTemplate"`
	MinAlignmentPeriod *string                                        `json:"minAlignmentPeriod"`
}

type jsonDashboardWidgetXyChartDataSets DashboardWidgetXyChartDataSets

func (r *DashboardWidgetXyChartDataSets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSets
	} else {

		r.TimeSeriesQuery = res.TimeSeriesQuery

		r.PlotType = res.PlotType

		r.LegendTemplate = res.LegendTemplate

		r.MinAlignmentPeriod = res.MinAlignmentPeriod

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSets *DashboardWidgetXyChartDataSets = &DashboardWidgetXyChartDataSets{empty: true}

func (r *DashboardWidgetXyChartDataSets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQuery struct {
	empty                   bool                                                                `json:"-"`
	TimeSeriesFilter        *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter      `json:"timeSeriesFilter"`
	TimeSeriesFilterRatio   *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio `json:"timeSeriesFilterRatio"`
	TimeSeriesQueryLanguage *string                                                             `json:"timeSeriesQueryLanguage"`
	ApiSource               *DashboardWidgetXyChartDataSetsTimeSeriesQueryApiSourceEnum         `json:"apiSource"`
	UnitOverride            *string                                                             `json:"unitOverride"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQuery DashboardWidgetXyChartDataSetsTimeSeriesQuery

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQuery) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQuery
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQuery
	} else {

		r.TimeSeriesFilter = res.TimeSeriesFilter

		r.TimeSeriesFilterRatio = res.TimeSeriesFilterRatio

		r.TimeSeriesQueryLanguage = res.TimeSeriesQueryLanguage

		r.ApiSource = res.ApiSource

		r.UnitOverride = res.UnitOverride

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQuery is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQuery *DashboardWidgetXyChartDataSetsTimeSeriesQuery = &DashboardWidgetXyChartDataSetsTimeSeriesQuery{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQuery) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQuery) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter struct {
	empty                bool                                                                               `json:"-"`
	Filter               *string                                                                            `json:"filter"`
	Aggregation          *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation          `json:"aggregation"`
	SecondaryAggregation *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation `json:"secondaryAggregation"`
	PickTimeSeriesFilter *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter `json:"pickTimeSeriesFilter"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter
	} else {

		r.Filter = res.Filter

		r.Aggregation = res.Aggregation

		r.SecondaryAggregation = res.SecondaryAggregation

		r.PickTimeSeriesFilter = res.PickTimeSeriesFilter

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation struct {
	empty                        bool                                                                                                  `json:"-"`
	AlignmentPeriod              *string                                                                                               `json:"alignmentPeriod"`
	PerSeriesAligner             *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum         `json:"perSeriesAligner"`
	CrossSeriesReducer           *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum       `json:"crossSeriesReducer"`
	GroupByFields                []string                                                                                              `json:"groupByFields"`
	ReduceFractionLessThanParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams `json:"reduceFractionLessThanParams"`
	ReduceMakeDistributionParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams `json:"reduceMakeDistributionParams"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation
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

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams struct {
	empty     bool     `json:"-"`
	Threshold *float64 `json:"threshold"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams
	} else {

		r.Threshold = res.Threshold

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams struct {
	empty            bool                                                                                                                  `json:"-"`
	BucketOptions    *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions    `json:"bucketOptions"`
	ExemplarSampling *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling `json:"exemplarSampling"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams
	} else {

		r.BucketOptions = res.BucketOptions

		r.ExemplarSampling = res.ExemplarSampling

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions struct {
	empty              bool                                                                                                                                 `json:"-"`
	LinearBuckets      *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions
	} else {

		r.LinearBuckets = res.LinearBuckets

		r.ExponentialBuckets = res.ExponentialBuckets

		r.ExplicitBuckets = res.ExplicitBuckets

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.Width = res.Width

		r.Offset = res.Offset

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.GrowthFactor = res.GrowthFactor

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	} else {

		r.Bounds = res.Bounds

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling struct {
	empty        bool     `json:"-"`
	MinimumValue *float64 `json:"minimumValue"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling
	} else {

		r.MinimumValue = res.MinimumValue

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation struct {
	empty                        bool                                                                                                           `json:"-"`
	AlignmentPeriod              *string                                                                                                        `json:"alignmentPeriod"`
	PerSeriesAligner             *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum         `json:"perSeriesAligner"`
	CrossSeriesReducer           *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum       `json:"crossSeriesReducer"`
	GroupByFields                []string                                                                                                       `json:"groupByFields"`
	ReduceFractionLessThanParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams `json:"reduceFractionLessThanParams"`
	ReduceMakeDistributionParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams `json:"reduceMakeDistributionParams"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
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

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams struct {
	empty     bool     `json:"-"`
	Threshold *float64 `json:"threshold"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams
	} else {

		r.Threshold = res.Threshold

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams struct {
	empty            bool                                                                                                                           `json:"-"`
	BucketOptions    *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions    `json:"bucketOptions"`
	ExemplarSampling *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling `json:"exemplarSampling"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams
	} else {

		r.BucketOptions = res.BucketOptions

		r.ExemplarSampling = res.ExemplarSampling

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions struct {
	empty              bool                                                                                                                                          `json:"-"`
	LinearBuckets      *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions
	} else {

		r.LinearBuckets = res.LinearBuckets

		r.ExponentialBuckets = res.ExponentialBuckets

		r.ExplicitBuckets = res.ExplicitBuckets

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.Width = res.Width

		r.Offset = res.Offset

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.GrowthFactor = res.GrowthFactor

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	} else {

		r.Bounds = res.Bounds

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling struct {
	empty        bool     `json:"-"`
	MinimumValue *float64 `json:"minimumValue"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling
	} else {

		r.MinimumValue = res.MinimumValue

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter struct {
	empty         bool                                                                                                `json:"-"`
	RankingMethod *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum `json:"rankingMethod"`
	NumTimeSeries *int64                                                                                              `json:"numTimeSeries"`
	Direction     *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum     `json:"direction"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
	} else {

		r.RankingMethod = res.RankingMethod

		r.NumTimeSeries = res.NumTimeSeries

		r.Direction = res.Direction

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio struct {
	empty                bool                                                                                    `json:"-"`
	Numerator            *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator            `json:"numerator"`
	Denominator          *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator          `json:"denominator"`
	SecondaryAggregation *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation `json:"secondaryAggregation"`
	PickTimeSeriesFilter *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter `json:"pickTimeSeriesFilter"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio
	} else {

		r.Numerator = res.Numerator

		r.Denominator = res.Denominator

		r.SecondaryAggregation = res.SecondaryAggregation

		r.PickTimeSeriesFilter = res.PickTimeSeriesFilter

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator struct {
	empty       bool                                                                                    `json:"-"`
	Filter      *string                                                                                 `json:"filter"`
	Aggregation *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation `json:"aggregation"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator
	} else {

		r.Filter = res.Filter

		r.Aggregation = res.Aggregation

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation struct {
	empty                        bool                                                                                                                `json:"-"`
	AlignmentPeriod              *string                                                                                                             `json:"alignmentPeriod"`
	PerSeriesAligner             *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum         `json:"perSeriesAligner"`
	CrossSeriesReducer           *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum       `json:"crossSeriesReducer"`
	GroupByFields                []string                                                                                                            `json:"groupByFields"`
	ReduceFractionLessThanParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams `json:"reduceFractionLessThanParams"`
	ReduceMakeDistributionParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams `json:"reduceMakeDistributionParams"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
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

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams struct {
	empty     bool     `json:"-"`
	Threshold *float64 `json:"threshold"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams
	} else {

		r.Threshold = res.Threshold

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams struct {
	empty            bool                                                                                                                                `json:"-"`
	BucketOptions    *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions    `json:"bucketOptions"`
	ExemplarSampling *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling `json:"exemplarSampling"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams
	} else {

		r.BucketOptions = res.BucketOptions

		r.ExemplarSampling = res.ExemplarSampling

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions struct {
	empty              bool                                                                                                                                               `json:"-"`
	LinearBuckets      *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions
	} else {

		r.LinearBuckets = res.LinearBuckets

		r.ExponentialBuckets = res.ExponentialBuckets

		r.ExplicitBuckets = res.ExplicitBuckets

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.Width = res.Width

		r.Offset = res.Offset

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.GrowthFactor = res.GrowthFactor

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	} else {

		r.Bounds = res.Bounds

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling struct {
	empty        bool     `json:"-"`
	MinimumValue *float64 `json:"minimumValue"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling
	} else {

		r.MinimumValue = res.MinimumValue

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator struct {
	empty       bool                                                                                      `json:"-"`
	Filter      *string                                                                                   `json:"filter"`
	Aggregation *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation `json:"aggregation"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator
	} else {

		r.Filter = res.Filter

		r.Aggregation = res.Aggregation

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation struct {
	empty                        bool                                                                                                                  `json:"-"`
	AlignmentPeriod              *string                                                                                                               `json:"alignmentPeriod"`
	PerSeriesAligner             *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum         `json:"perSeriesAligner"`
	CrossSeriesReducer           *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum       `json:"crossSeriesReducer"`
	GroupByFields                []string                                                                                                              `json:"groupByFields"`
	ReduceFractionLessThanParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams `json:"reduceFractionLessThanParams"`
	ReduceMakeDistributionParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams `json:"reduceMakeDistributionParams"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
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

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams struct {
	empty     bool     `json:"-"`
	Threshold *float64 `json:"threshold"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams
	} else {

		r.Threshold = res.Threshold

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams struct {
	empty            bool                                                                                                                                  `json:"-"`
	BucketOptions    *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions    `json:"bucketOptions"`
	ExemplarSampling *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling `json:"exemplarSampling"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams
	} else {

		r.BucketOptions = res.BucketOptions

		r.ExemplarSampling = res.ExemplarSampling

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions struct {
	empty              bool                                                                                                                                                 `json:"-"`
	LinearBuckets      *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions
	} else {

		r.LinearBuckets = res.LinearBuckets

		r.ExponentialBuckets = res.ExponentialBuckets

		r.ExplicitBuckets = res.ExplicitBuckets

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.Width = res.Width

		r.Offset = res.Offset

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.GrowthFactor = res.GrowthFactor

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	} else {

		r.Bounds = res.Bounds

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling struct {
	empty        bool     `json:"-"`
	MinimumValue *float64 `json:"minimumValue"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling
	} else {

		r.MinimumValue = res.MinimumValue

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation struct {
	empty                        bool                                                                                                                `json:"-"`
	AlignmentPeriod              *string                                                                                                             `json:"alignmentPeriod"`
	PerSeriesAligner             *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum         `json:"perSeriesAligner"`
	CrossSeriesReducer           *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum       `json:"crossSeriesReducer"`
	GroupByFields                []string                                                                                                            `json:"groupByFields"`
	ReduceFractionLessThanParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams `json:"reduceFractionLessThanParams"`
	ReduceMakeDistributionParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams `json:"reduceMakeDistributionParams"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
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

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams struct {
	empty     bool     `json:"-"`
	Threshold *float64 `json:"threshold"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams
	} else {

		r.Threshold = res.Threshold

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams struct {
	empty            bool                                                                                                                                `json:"-"`
	BucketOptions    *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions    `json:"bucketOptions"`
	ExemplarSampling *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling `json:"exemplarSampling"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams
	} else {

		r.BucketOptions = res.BucketOptions

		r.ExemplarSampling = res.ExemplarSampling

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions struct {
	empty              bool                                                                                                                                               `json:"-"`
	LinearBuckets      *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions
	} else {

		r.LinearBuckets = res.LinearBuckets

		r.ExponentialBuckets = res.ExponentialBuckets

		r.ExplicitBuckets = res.ExplicitBuckets

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.Width = res.Width

		r.Offset = res.Offset

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.GrowthFactor = res.GrowthFactor

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	} else {

		r.Bounds = res.Bounds

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling struct {
	empty        bool     `json:"-"`
	MinimumValue *float64 `json:"minimumValue"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling
	} else {

		r.MinimumValue = res.MinimumValue

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter struct {
	empty         bool                                                                                                     `json:"-"`
	RankingMethod *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum `json:"rankingMethod"`
	NumTimeSeries *int64                                                                                                   `json:"numTimeSeries"`
	Direction     *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum     `json:"direction"`
}

type jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
	} else {

		r.RankingMethod = res.RankingMethod

		r.NumTimeSeries = res.NumTimeSeries

		r.Direction = res.Direction

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartSourceDrilldown struct {
	empty                         bool                                                                 `json:"-"`
	ResourceTypeDrilldown         *DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown          `json:"resourceTypeDrilldown"`
	ResourceLabelDrilldowns       []DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns       `json:"resourceLabelDrilldowns"`
	MetadataSystemLabelDrilldowns []DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns `json:"metadataSystemLabelDrilldowns"`
	MetadataUserLabelDrilldowns   []DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns   `json:"metadataUserLabelDrilldowns"`
	GroupNameDrilldown            *DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown             `json:"groupNameDrilldown"`
	ServiceNameDrilldown          *DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown           `json:"serviceNameDrilldown"`
	ServiceTypeDrilldown          *DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown           `json:"serviceTypeDrilldown"`
}

type jsonDashboardWidgetXyChartSourceDrilldown DashboardWidgetXyChartSourceDrilldown

func (r *DashboardWidgetXyChartSourceDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartSourceDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartSourceDrilldown
	} else {

		r.ResourceTypeDrilldown = res.ResourceTypeDrilldown

		r.ResourceLabelDrilldowns = res.ResourceLabelDrilldowns

		r.MetadataSystemLabelDrilldowns = res.MetadataSystemLabelDrilldowns

		r.MetadataUserLabelDrilldowns = res.MetadataUserLabelDrilldowns

		r.GroupNameDrilldown = res.GroupNameDrilldown

		r.ServiceNameDrilldown = res.ServiceNameDrilldown

		r.ServiceTypeDrilldown = res.ServiceTypeDrilldown

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartSourceDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartSourceDrilldown *DashboardWidgetXyChartSourceDrilldown = &DashboardWidgetXyChartSourceDrilldown{empty: true}

func (r *DashboardWidgetXyChartSourceDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartSourceDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown struct {
	empty        bool     `json:"-"`
	TargetValues []string `json:"targetValues"`
}

type jsonDashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown

func (r *DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown
	} else {

		r.TargetValues = res.TargetValues

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown *DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown = &DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown{empty: true}

func (r *DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartSourceDrilldownResourceTypeDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns struct {
	empty             bool                                                                             `json:"-"`
	Label             *string                                                                          `json:"label"`
	LogicalOperator   *DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum `json:"logicalOperator"`
	ValueRestrictions []DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions  `json:"valueRestrictions"`
}

type jsonDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns

func (r *DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns
	} else {

		r.Label = res.Label

		r.LogicalOperator = res.LogicalOperator

		r.ValueRestrictions = res.ValueRestrictions

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns *DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns = &DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns{empty: true}

func (r *DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldowns) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions struct {
	empty       bool                                                                                         `json:"-"`
	TargetValue *string                                                                                      `json:"targetValue"`
	Comparator  *DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum `json:"comparator"`
}

type jsonDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions

func (r *DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions
	} else {

		r.TargetValue = res.TargetValue

		r.Comparator = res.Comparator

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions *DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions = &DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions{empty: true}

func (r *DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartSourceDrilldownResourceLabelDrilldownsValueRestrictions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns struct {
	empty             bool                                                                                   `json:"-"`
	Label             *string                                                                                `json:"label"`
	LogicalOperator   *DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum `json:"logicalOperator"`
	ValueRestrictions []DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions  `json:"valueRestrictions"`
}

type jsonDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns

func (r *DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns
	} else {

		r.Label = res.Label

		r.LogicalOperator = res.LogicalOperator

		r.ValueRestrictions = res.ValueRestrictions

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns *DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns = &DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns{empty: true}

func (r *DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldowns) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions struct {
	empty       bool                                                                                               `json:"-"`
	TargetValue *string                                                                                            `json:"targetValue"`
	Comparator  *DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum `json:"comparator"`
}

type jsonDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions

func (r *DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions
	} else {

		r.TargetValue = res.TargetValue

		r.Comparator = res.Comparator

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions *DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions = &DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions{empty: true}

func (r *DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns struct {
	empty             bool                                                                                 `json:"-"`
	Label             *string                                                                              `json:"label"`
	LogicalOperator   *DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum `json:"logicalOperator"`
	ValueRestrictions []DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions  `json:"valueRestrictions"`
}

type jsonDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns

func (r *DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns
	} else {

		r.Label = res.Label

		r.LogicalOperator = res.LogicalOperator

		r.ValueRestrictions = res.ValueRestrictions

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns *DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns = &DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns{empty: true}

func (r *DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldowns) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions struct {
	empty       bool                                                                                             `json:"-"`
	TargetValue *string                                                                                          `json:"targetValue"`
	Comparator  *DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum `json:"comparator"`
}

type jsonDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions

func (r *DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions
	} else {

		r.TargetValue = res.TargetValue

		r.Comparator = res.Comparator

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions *DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions = &DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions{empty: true}

func (r *DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown struct {
	empty        bool     `json:"-"`
	TargetValues []string `json:"targetValues"`
}

type jsonDashboardWidgetXyChartSourceDrilldownGroupNameDrilldown DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown

func (r *DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartSourceDrilldownGroupNameDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartSourceDrilldownGroupNameDrilldown
	} else {

		r.TargetValues = res.TargetValues

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartSourceDrilldownGroupNameDrilldown *DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown = &DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown{empty: true}

func (r *DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartSourceDrilldownGroupNameDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown struct {
	empty        bool     `json:"-"`
	TargetValues []string `json:"targetValues"`
}

type jsonDashboardWidgetXyChartSourceDrilldownServiceNameDrilldown DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown

func (r *DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartSourceDrilldownServiceNameDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartSourceDrilldownServiceNameDrilldown
	} else {

		r.TargetValues = res.TargetValues

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartSourceDrilldownServiceNameDrilldown *DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown = &DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown{empty: true}

func (r *DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartSourceDrilldownServiceNameDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown struct {
	empty bool                                                                 `json:"-"`
	Types []DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldownTypesEnum `json:"types"`
}

type jsonDashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown

func (r *DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown
	} else {

		r.Types = res.Types

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown *DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown = &DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown{empty: true}

func (r *DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartSourceDrilldownServiceTypeDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartMetricDrilldown struct {
	empty                  bool                                                         `json:"-"`
	MetricTypeDrilldown    *DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown    `json:"metricTypeDrilldown"`
	MetricLabelDrilldowns  []DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns `json:"metricLabelDrilldowns"`
	MetricGroupByDrilldown *DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown `json:"metricGroupByDrilldown"`
}

type jsonDashboardWidgetXyChartMetricDrilldown DashboardWidgetXyChartMetricDrilldown

func (r *DashboardWidgetXyChartMetricDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartMetricDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartMetricDrilldown
	} else {

		r.MetricTypeDrilldown = res.MetricTypeDrilldown

		r.MetricLabelDrilldowns = res.MetricLabelDrilldowns

		r.MetricGroupByDrilldown = res.MetricGroupByDrilldown

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartMetricDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartMetricDrilldown *DashboardWidgetXyChartMetricDrilldown = &DashboardWidgetXyChartMetricDrilldown{empty: true}

func (r *DashboardWidgetXyChartMetricDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartMetricDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown struct {
	empty       bool    `json:"-"`
	TargetValue *string `json:"targetValue"`
}

type jsonDashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown

func (r *DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown
	} else {

		r.TargetValue = res.TargetValue

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown *DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown = &DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown{empty: true}

func (r *DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartMetricDrilldownMetricTypeDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns struct {
	empty             bool                                                                           `json:"-"`
	Label             *string                                                                        `json:"label"`
	LogicalOperator   *DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum `json:"logicalOperator"`
	ValueRestrictions []DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions  `json:"valueRestrictions"`
}

type jsonDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns

func (r *DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns
	} else {

		r.Label = res.Label

		r.LogicalOperator = res.LogicalOperator

		r.ValueRestrictions = res.ValueRestrictions

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns *DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns = &DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns{empty: true}

func (r *DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldowns) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions struct {
	empty       bool                                                                                       `json:"-"`
	TargetValue *string                                                                                    `json:"targetValue"`
	Comparator  *DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum `json:"comparator"`
}

type jsonDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions

func (r *DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions
	} else {

		r.TargetValue = res.TargetValue

		r.Comparator = res.Comparator

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions *DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions = &DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions{empty: true}

func (r *DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartMetricDrilldownMetricLabelDrilldownsValueRestrictions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown struct {
	empty                bool                                                                    `json:"-"`
	ResourceLabels       []string                                                                `json:"resourceLabels"`
	MetricLabels         []string                                                                `json:"metricLabels"`
	MetadataSystemLabels []string                                                                `json:"metadataSystemLabels"`
	MetadataUserLabels   []string                                                                `json:"metadataUserLabels"`
	Reducer              *DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldownReducerEnum `json:"reducer"`
}

type jsonDashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown

func (r *DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown
	} else {

		r.ResourceLabels = res.ResourceLabels

		r.MetricLabels = res.MetricLabels

		r.MetadataSystemLabels = res.MetadataSystemLabels

		r.MetadataUserLabels = res.MetadataUserLabels

		r.Reducer = res.Reducer

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown *DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown = &DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown{empty: true}

func (r *DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartMetricDrilldownMetricGroupByDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartThresholds struct {
	empty     bool                                           `json:"-"`
	Label     *string                                        `json:"label"`
	Value     *float64                                       `json:"value"`
	Color     *DashboardWidgetXyChartThresholdsColorEnum     `json:"color"`
	Direction *DashboardWidgetXyChartThresholdsDirectionEnum `json:"direction"`
}

type jsonDashboardWidgetXyChartThresholds DashboardWidgetXyChartThresholds

func (r *DashboardWidgetXyChartThresholds) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartThresholds
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartThresholds
	} else {

		r.Label = res.Label

		r.Value = res.Value

		r.Color = res.Color

		r.Direction = res.Direction

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartThresholds is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartThresholds *DashboardWidgetXyChartThresholds = &DashboardWidgetXyChartThresholds{empty: true}

func (r *DashboardWidgetXyChartThresholds) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartThresholds) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartXAxis struct {
	empty bool                                  `json:"-"`
	Label *string                               `json:"label"`
	Scale *DashboardWidgetXyChartXAxisScaleEnum `json:"scale"`
}

type jsonDashboardWidgetXyChartXAxis DashboardWidgetXyChartXAxis

func (r *DashboardWidgetXyChartXAxis) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartXAxis
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartXAxis
	} else {

		r.Label = res.Label

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartXAxis is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartXAxis *DashboardWidgetXyChartXAxis = &DashboardWidgetXyChartXAxis{empty: true}

func (r *DashboardWidgetXyChartXAxis) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartXAxis) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartYAxis struct {
	empty bool                                  `json:"-"`
	Label *string                               `json:"label"`
	Scale *DashboardWidgetXyChartYAxisScaleEnum `json:"scale"`
}

type jsonDashboardWidgetXyChartYAxis DashboardWidgetXyChartYAxis

func (r *DashboardWidgetXyChartYAxis) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartYAxis
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartYAxis
	} else {

		r.Label = res.Label

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartYAxis is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartYAxis *DashboardWidgetXyChartYAxis = &DashboardWidgetXyChartYAxis{empty: true}

func (r *DashboardWidgetXyChartYAxis) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartYAxis) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartChartOptions struct {
	empty      bool                                        `json:"-"`
	Mode       *DashboardWidgetXyChartChartOptionsModeEnum `json:"mode"`
	ShowLegend *bool                                       `json:"showLegend"`
}

type jsonDashboardWidgetXyChartChartOptions DashboardWidgetXyChartChartOptions

func (r *DashboardWidgetXyChartChartOptions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetXyChartChartOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetXyChartChartOptions
	} else {

		r.Mode = res.Mode

		r.ShowLegend = res.ShowLegend

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartChartOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartChartOptions *DashboardWidgetXyChartChartOptions = &DashboardWidgetXyChartChartOptions{empty: true}

func (r *DashboardWidgetXyChartChartOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartChartOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecard struct {
	empty           bool                                     `json:"-"`
	TimeSeriesQuery *DashboardWidgetScorecardTimeSeriesQuery `json:"timeSeriesQuery"`
	SourceDrilldown *DashboardWidgetScorecardSourceDrilldown `json:"sourceDrilldown"`
	MetricDrilldown *DashboardWidgetScorecardMetricDrilldown `json:"metricDrilldown"`
	GaugeView       *DashboardWidgetScorecardGaugeView       `json:"gaugeView"`
	SparkChartView  *DashboardWidgetScorecardSparkChartView  `json:"sparkChartView"`
	Thresholds      []DashboardWidgetScorecardThresholds     `json:"thresholds"`
}

type jsonDashboardWidgetScorecard DashboardWidgetScorecard

func (r *DashboardWidgetScorecard) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecard
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecard
	} else {

		r.TimeSeriesQuery = res.TimeSeriesQuery

		r.SourceDrilldown = res.SourceDrilldown

		r.MetricDrilldown = res.MetricDrilldown

		r.GaugeView = res.GaugeView

		r.SparkChartView = res.SparkChartView

		r.Thresholds = res.Thresholds

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecard is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecard *DashboardWidgetScorecard = &DashboardWidgetScorecard{empty: true}

func (r *DashboardWidgetScorecard) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecard) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQuery struct {
	empty                   bool                                                          `json:"-"`
	TimeSeriesFilter        *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter      `json:"timeSeriesFilter"`
	TimeSeriesFilterRatio   *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio `json:"timeSeriesFilterRatio"`
	TimeSeriesQueryLanguage *string                                                       `json:"timeSeriesQueryLanguage"`
	ApiSource               *DashboardWidgetScorecardTimeSeriesQueryApiSourceEnum         `json:"apiSource"`
	UnitOverride            *string                                                       `json:"unitOverride"`
}

type jsonDashboardWidgetScorecardTimeSeriesQuery DashboardWidgetScorecardTimeSeriesQuery

func (r *DashboardWidgetScorecardTimeSeriesQuery) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQuery
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQuery
	} else {

		r.TimeSeriesFilter = res.TimeSeriesFilter

		r.TimeSeriesFilterRatio = res.TimeSeriesFilterRatio

		r.TimeSeriesQueryLanguage = res.TimeSeriesQueryLanguage

		r.ApiSource = res.ApiSource

		r.UnitOverride = res.UnitOverride

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQuery is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQuery *DashboardWidgetScorecardTimeSeriesQuery = &DashboardWidgetScorecardTimeSeriesQuery{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQuery) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQuery) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter struct {
	empty                bool                                                                         `json:"-"`
	Filter               *string                                                                      `json:"filter"`
	Aggregation          *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation          `json:"aggregation"`
	SecondaryAggregation *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation `json:"secondaryAggregation"`
	PickTimeSeriesFilter *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter `json:"pickTimeSeriesFilter"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter
	} else {

		r.Filter = res.Filter

		r.Aggregation = res.Aggregation

		r.SecondaryAggregation = res.SecondaryAggregation

		r.PickTimeSeriesFilter = res.PickTimeSeriesFilter

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation struct {
	empty                        bool                                                                                            `json:"-"`
	AlignmentPeriod              *string                                                                                         `json:"alignmentPeriod"`
	PerSeriesAligner             *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum         `json:"perSeriesAligner"`
	CrossSeriesReducer           *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum       `json:"crossSeriesReducer"`
	GroupByFields                []string                                                                                        `json:"groupByFields"`
	ReduceFractionLessThanParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams `json:"reduceFractionLessThanParams"`
	ReduceMakeDistributionParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams `json:"reduceMakeDistributionParams"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation
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

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams struct {
	empty     bool     `json:"-"`
	Threshold *float64 `json:"threshold"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams
	} else {

		r.Threshold = res.Threshold

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceFractionLessThanParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams struct {
	empty            bool                                                                                                            `json:"-"`
	BucketOptions    *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions    `json:"bucketOptions"`
	ExemplarSampling *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling `json:"exemplarSampling"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams
	} else {

		r.BucketOptions = res.BucketOptions

		r.ExemplarSampling = res.ExemplarSampling

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions struct {
	empty              bool                                                                                                                           `json:"-"`
	LinearBuckets      *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions
	} else {

		r.LinearBuckets = res.LinearBuckets

		r.ExponentialBuckets = res.ExponentialBuckets

		r.ExplicitBuckets = res.ExplicitBuckets

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.Width = res.Width

		r.Offset = res.Offset

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.GrowthFactor = res.GrowthFactor

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	} else {

		r.Bounds = res.Bounds

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling struct {
	empty        bool     `json:"-"`
	MinimumValue *float64 `json:"minimumValue"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling
	} else {

		r.MinimumValue = res.MinimumValue

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationReduceMakeDistributionParamsExemplarSampling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation struct {
	empty                        bool                                                                                                     `json:"-"`
	AlignmentPeriod              *string                                                                                                  `json:"alignmentPeriod"`
	PerSeriesAligner             *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum         `json:"perSeriesAligner"`
	CrossSeriesReducer           *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum       `json:"crossSeriesReducer"`
	GroupByFields                []string                                                                                                 `json:"groupByFields"`
	ReduceFractionLessThanParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams `json:"reduceFractionLessThanParams"`
	ReduceMakeDistributionParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams `json:"reduceMakeDistributionParams"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation
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

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams struct {
	empty     bool     `json:"-"`
	Threshold *float64 `json:"threshold"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams
	} else {

		r.Threshold = res.Threshold

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceFractionLessThanParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams struct {
	empty            bool                                                                                                                     `json:"-"`
	BucketOptions    *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions    `json:"bucketOptions"`
	ExemplarSampling *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling `json:"exemplarSampling"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams
	} else {

		r.BucketOptions = res.BucketOptions

		r.ExemplarSampling = res.ExemplarSampling

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions struct {
	empty              bool                                                                                                                                    `json:"-"`
	LinearBuckets      *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions
	} else {

		r.LinearBuckets = res.LinearBuckets

		r.ExponentialBuckets = res.ExponentialBuckets

		r.ExplicitBuckets = res.ExplicitBuckets

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.Width = res.Width

		r.Offset = res.Offset

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.GrowthFactor = res.GrowthFactor

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	} else {

		r.Bounds = res.Bounds

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling struct {
	empty        bool     `json:"-"`
	MinimumValue *float64 `json:"minimumValue"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling
	} else {

		r.MinimumValue = res.MinimumValue

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter struct {
	empty         bool                                                                                          `json:"-"`
	RankingMethod *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum `json:"rankingMethod"`
	NumTimeSeries *int64                                                                                        `json:"numTimeSeries"`
	Direction     *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum     `json:"direction"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter
	} else {

		r.RankingMethod = res.RankingMethod

		r.NumTimeSeries = res.NumTimeSeries

		r.Direction = res.Direction

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio struct {
	empty                bool                                                                              `json:"-"`
	Numerator            *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator            `json:"numerator"`
	Denominator          *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator          `json:"denominator"`
	SecondaryAggregation *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation `json:"secondaryAggregation"`
	PickTimeSeriesFilter *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter `json:"pickTimeSeriesFilter"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio
	} else {

		r.Numerator = res.Numerator

		r.Denominator = res.Denominator

		r.SecondaryAggregation = res.SecondaryAggregation

		r.PickTimeSeriesFilter = res.PickTimeSeriesFilter

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator struct {
	empty       bool                                                                              `json:"-"`
	Filter      *string                                                                           `json:"filter"`
	Aggregation *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation `json:"aggregation"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator
	} else {

		r.Filter = res.Filter

		r.Aggregation = res.Aggregation

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation struct {
	empty                        bool                                                                                                          `json:"-"`
	AlignmentPeriod              *string                                                                                                       `json:"alignmentPeriod"`
	PerSeriesAligner             *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum         `json:"perSeriesAligner"`
	CrossSeriesReducer           *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum       `json:"crossSeriesReducer"`
	GroupByFields                []string                                                                                                      `json:"groupByFields"`
	ReduceFractionLessThanParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams `json:"reduceFractionLessThanParams"`
	ReduceMakeDistributionParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams `json:"reduceMakeDistributionParams"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation
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

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams struct {
	empty     bool     `json:"-"`
	Threshold *float64 `json:"threshold"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams
	} else {

		r.Threshold = res.Threshold

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceFractionLessThanParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams struct {
	empty            bool                                                                                                                          `json:"-"`
	BucketOptions    *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions    `json:"bucketOptions"`
	ExemplarSampling *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling `json:"exemplarSampling"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams
	} else {

		r.BucketOptions = res.BucketOptions

		r.ExemplarSampling = res.ExemplarSampling

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions struct {
	empty              bool                                                                                                                                         `json:"-"`
	LinearBuckets      *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions
	} else {

		r.LinearBuckets = res.LinearBuckets

		r.ExponentialBuckets = res.ExponentialBuckets

		r.ExplicitBuckets = res.ExplicitBuckets

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.Width = res.Width

		r.Offset = res.Offset

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.GrowthFactor = res.GrowthFactor

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	} else {

		r.Bounds = res.Bounds

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling struct {
	empty        bool     `json:"-"`
	MinimumValue *float64 `json:"minimumValue"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling
	} else {

		r.MinimumValue = res.MinimumValue

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationReduceMakeDistributionParamsExemplarSampling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator struct {
	empty       bool                                                                                `json:"-"`
	Filter      *string                                                                             `json:"filter"`
	Aggregation *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation `json:"aggregation"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator
	} else {

		r.Filter = res.Filter

		r.Aggregation = res.Aggregation

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation struct {
	empty                        bool                                                                                                            `json:"-"`
	AlignmentPeriod              *string                                                                                                         `json:"alignmentPeriod"`
	PerSeriesAligner             *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum         `json:"perSeriesAligner"`
	CrossSeriesReducer           *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum       `json:"crossSeriesReducer"`
	GroupByFields                []string                                                                                                        `json:"groupByFields"`
	ReduceFractionLessThanParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams `json:"reduceFractionLessThanParams"`
	ReduceMakeDistributionParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams `json:"reduceMakeDistributionParams"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation
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

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams struct {
	empty     bool     `json:"-"`
	Threshold *float64 `json:"threshold"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams
	} else {

		r.Threshold = res.Threshold

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceFractionLessThanParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams struct {
	empty            bool                                                                                                                            `json:"-"`
	BucketOptions    *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions    `json:"bucketOptions"`
	ExemplarSampling *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling `json:"exemplarSampling"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams
	} else {

		r.BucketOptions = res.BucketOptions

		r.ExemplarSampling = res.ExemplarSampling

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions struct {
	empty              bool                                                                                                                                           `json:"-"`
	LinearBuckets      *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions
	} else {

		r.LinearBuckets = res.LinearBuckets

		r.ExponentialBuckets = res.ExponentialBuckets

		r.ExplicitBuckets = res.ExplicitBuckets

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.Width = res.Width

		r.Offset = res.Offset

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.GrowthFactor = res.GrowthFactor

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	} else {

		r.Bounds = res.Bounds

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling struct {
	empty        bool     `json:"-"`
	MinimumValue *float64 `json:"minimumValue"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling
	} else {

		r.MinimumValue = res.MinimumValue

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationReduceMakeDistributionParamsExemplarSampling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation struct {
	empty                        bool                                                                                                          `json:"-"`
	AlignmentPeriod              *string                                                                                                       `json:"alignmentPeriod"`
	PerSeriesAligner             *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum         `json:"perSeriesAligner"`
	CrossSeriesReducer           *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum       `json:"crossSeriesReducer"`
	GroupByFields                []string                                                                                                      `json:"groupByFields"`
	ReduceFractionLessThanParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams `json:"reduceFractionLessThanParams"`
	ReduceMakeDistributionParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams `json:"reduceMakeDistributionParams"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation
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

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams struct {
	empty     bool     `json:"-"`
	Threshold *float64 `json:"threshold"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams
	} else {

		r.Threshold = res.Threshold

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceFractionLessThanParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams struct {
	empty            bool                                                                                                                          `json:"-"`
	BucketOptions    *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions    `json:"bucketOptions"`
	ExemplarSampling *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling `json:"exemplarSampling"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams
	} else {

		r.BucketOptions = res.BucketOptions

		r.ExemplarSampling = res.ExemplarSampling

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParams) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions struct {
	empty              bool                                                                                                                                         `json:"-"`
	LinearBuckets      *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets      `json:"linearBuckets"`
	ExponentialBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets `json:"exponentialBuckets"`
	ExplicitBuckets    *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets    `json:"explicitBuckets"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions
	} else {

		r.LinearBuckets = res.LinearBuckets

		r.ExponentialBuckets = res.ExponentialBuckets

		r.ExplicitBuckets = res.ExplicitBuckets

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	Width            *float64 `json:"width"`
	Offset           *float64 `json:"offset"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.Width = res.Width

		r.Offset = res.Offset

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsLinearBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets struct {
	empty            bool     `json:"-"`
	NumFiniteBuckets *int64   `json:"numFiniteBuckets"`
	GrowthFactor     *float64 `json:"growthFactor"`
	Scale            *float64 `json:"scale"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets
	} else {

		r.NumFiniteBuckets = res.NumFiniteBuckets

		r.GrowthFactor = res.GrowthFactor

		r.Scale = res.Scale

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExponentialBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets struct {
	empty  bool      `json:"-"`
	Bounds []float64 `json:"bounds"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets
	} else {

		r.Bounds = res.Bounds

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsBucketOptionsExplicitBuckets) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling struct {
	empty        bool     `json:"-"`
	MinimumValue *float64 `json:"minimumValue"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling
	} else {

		r.MinimumValue = res.MinimumValue

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationReduceMakeDistributionParamsExemplarSampling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter struct {
	empty         bool                                                                                               `json:"-"`
	RankingMethod *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum `json:"rankingMethod"`
	NumTimeSeries *int64                                                                                             `json:"numTimeSeries"`
	Direction     *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum     `json:"direction"`
}

type jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter
	} else {

		r.RankingMethod = res.RankingMethod

		r.NumTimeSeries = res.NumTimeSeries

		r.Direction = res.Direction

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardSourceDrilldown struct {
	empty                         bool                                                                   `json:"-"`
	ResourceTypeDrilldown         *DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown          `json:"resourceTypeDrilldown"`
	ResourceLabelDrilldowns       []DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns       `json:"resourceLabelDrilldowns"`
	MetadataSystemLabelDrilldowns []DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns `json:"metadataSystemLabelDrilldowns"`
	MetadataUserLabelDrilldowns   []DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns   `json:"metadataUserLabelDrilldowns"`
	GroupNameDrilldown            *DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown             `json:"groupNameDrilldown"`
	ServiceNameDrilldown          *DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown           `json:"serviceNameDrilldown"`
	ServiceTypeDrilldown          *DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown           `json:"serviceTypeDrilldown"`
}

type jsonDashboardWidgetScorecardSourceDrilldown DashboardWidgetScorecardSourceDrilldown

func (r *DashboardWidgetScorecardSourceDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardSourceDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardSourceDrilldown
	} else {

		r.ResourceTypeDrilldown = res.ResourceTypeDrilldown

		r.ResourceLabelDrilldowns = res.ResourceLabelDrilldowns

		r.MetadataSystemLabelDrilldowns = res.MetadataSystemLabelDrilldowns

		r.MetadataUserLabelDrilldowns = res.MetadataUserLabelDrilldowns

		r.GroupNameDrilldown = res.GroupNameDrilldown

		r.ServiceNameDrilldown = res.ServiceNameDrilldown

		r.ServiceTypeDrilldown = res.ServiceTypeDrilldown

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardSourceDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardSourceDrilldown *DashboardWidgetScorecardSourceDrilldown = &DashboardWidgetScorecardSourceDrilldown{empty: true}

func (r *DashboardWidgetScorecardSourceDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardSourceDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown struct {
	empty        bool     `json:"-"`
	TargetValues []string `json:"targetValues"`
}

type jsonDashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown

func (r *DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown
	} else {

		r.TargetValues = res.TargetValues

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown *DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown = &DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown{empty: true}

func (r *DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardSourceDrilldownResourceTypeDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns struct {
	empty             bool                                                                               `json:"-"`
	Label             *string                                                                            `json:"label"`
	LogicalOperator   *DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsLogicalOperatorEnum `json:"logicalOperator"`
	ValueRestrictions []DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions  `json:"valueRestrictions"`
}

type jsonDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns

func (r *DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns
	} else {

		r.Label = res.Label

		r.LogicalOperator = res.LogicalOperator

		r.ValueRestrictions = res.ValueRestrictions

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns *DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns = &DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns{empty: true}

func (r *DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldowns) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions struct {
	empty       bool                                                                                           `json:"-"`
	TargetValue *string                                                                                        `json:"targetValue"`
	Comparator  *DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictionsComparatorEnum `json:"comparator"`
}

type jsonDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions

func (r *DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions
	} else {

		r.TargetValue = res.TargetValue

		r.Comparator = res.Comparator

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions *DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions = &DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions{empty: true}

func (r *DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardSourceDrilldownResourceLabelDrilldownsValueRestrictions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns struct {
	empty             bool                                                                                     `json:"-"`
	Label             *string                                                                                  `json:"label"`
	LogicalOperator   *DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsLogicalOperatorEnum `json:"logicalOperator"`
	ValueRestrictions []DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions  `json:"valueRestrictions"`
}

type jsonDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns

func (r *DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns
	} else {

		r.Label = res.Label

		r.LogicalOperator = res.LogicalOperator

		r.ValueRestrictions = res.ValueRestrictions

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns *DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns = &DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns{empty: true}

func (r *DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldowns) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions struct {
	empty       bool                                                                                                 `json:"-"`
	TargetValue *string                                                                                              `json:"targetValue"`
	Comparator  *DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictionsComparatorEnum `json:"comparator"`
}

type jsonDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions

func (r *DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions
	} else {

		r.TargetValue = res.TargetValue

		r.Comparator = res.Comparator

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions *DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions = &DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions{empty: true}

func (r *DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardSourceDrilldownMetadataSystemLabelDrilldownsValueRestrictions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns struct {
	empty             bool                                                                                   `json:"-"`
	Label             *string                                                                                `json:"label"`
	LogicalOperator   *DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsLogicalOperatorEnum `json:"logicalOperator"`
	ValueRestrictions []DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions  `json:"valueRestrictions"`
}

type jsonDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns

func (r *DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns
	} else {

		r.Label = res.Label

		r.LogicalOperator = res.LogicalOperator

		r.ValueRestrictions = res.ValueRestrictions

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns *DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns = &DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns{empty: true}

func (r *DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldowns) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions struct {
	empty       bool                                                                                               `json:"-"`
	TargetValue *string                                                                                            `json:"targetValue"`
	Comparator  *DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictionsComparatorEnum `json:"comparator"`
}

type jsonDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions

func (r *DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions
	} else {

		r.TargetValue = res.TargetValue

		r.Comparator = res.Comparator

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions *DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions = &DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions{empty: true}

func (r *DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardSourceDrilldownMetadataUserLabelDrilldownsValueRestrictions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown struct {
	empty        bool     `json:"-"`
	TargetValues []string `json:"targetValues"`
}

type jsonDashboardWidgetScorecardSourceDrilldownGroupNameDrilldown DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown

func (r *DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardSourceDrilldownGroupNameDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardSourceDrilldownGroupNameDrilldown
	} else {

		r.TargetValues = res.TargetValues

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardSourceDrilldownGroupNameDrilldown *DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown = &DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown{empty: true}

func (r *DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardSourceDrilldownGroupNameDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown struct {
	empty        bool     `json:"-"`
	TargetValues []string `json:"targetValues"`
}

type jsonDashboardWidgetScorecardSourceDrilldownServiceNameDrilldown DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown

func (r *DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardSourceDrilldownServiceNameDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardSourceDrilldownServiceNameDrilldown
	} else {

		r.TargetValues = res.TargetValues

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardSourceDrilldownServiceNameDrilldown *DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown = &DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown{empty: true}

func (r *DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardSourceDrilldownServiceNameDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown struct {
	empty bool                                                                   `json:"-"`
	Types []DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldownTypesEnum `json:"types"`
}

type jsonDashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown

func (r *DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown
	} else {

		r.Types = res.Types

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown *DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown = &DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown{empty: true}

func (r *DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardSourceDrilldownServiceTypeDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardMetricDrilldown struct {
	empty                  bool                                                           `json:"-"`
	MetricTypeDrilldown    *DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown    `json:"metricTypeDrilldown"`
	MetricLabelDrilldowns  []DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns `json:"metricLabelDrilldowns"`
	MetricGroupByDrilldown *DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown `json:"metricGroupByDrilldown"`
}

type jsonDashboardWidgetScorecardMetricDrilldown DashboardWidgetScorecardMetricDrilldown

func (r *DashboardWidgetScorecardMetricDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardMetricDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardMetricDrilldown
	} else {

		r.MetricTypeDrilldown = res.MetricTypeDrilldown

		r.MetricLabelDrilldowns = res.MetricLabelDrilldowns

		r.MetricGroupByDrilldown = res.MetricGroupByDrilldown

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardMetricDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardMetricDrilldown *DashboardWidgetScorecardMetricDrilldown = &DashboardWidgetScorecardMetricDrilldown{empty: true}

func (r *DashboardWidgetScorecardMetricDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardMetricDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown struct {
	empty       bool    `json:"-"`
	TargetValue *string `json:"targetValue"`
}

type jsonDashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown

func (r *DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown
	} else {

		r.TargetValue = res.TargetValue

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown *DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown = &DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown{empty: true}

func (r *DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardMetricDrilldownMetricTypeDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns struct {
	empty             bool                                                                             `json:"-"`
	Label             *string                                                                          `json:"label"`
	LogicalOperator   *DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsLogicalOperatorEnum `json:"logicalOperator"`
	ValueRestrictions []DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions  `json:"valueRestrictions"`
}

type jsonDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns

func (r *DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns
	} else {

		r.Label = res.Label

		r.LogicalOperator = res.LogicalOperator

		r.ValueRestrictions = res.ValueRestrictions

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns *DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns = &DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns{empty: true}

func (r *DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldowns) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions struct {
	empty       bool                                                                                         `json:"-"`
	TargetValue *string                                                                                      `json:"targetValue"`
	Comparator  *DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictionsComparatorEnum `json:"comparator"`
}

type jsonDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions

func (r *DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions
	} else {

		r.TargetValue = res.TargetValue

		r.Comparator = res.Comparator

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions *DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions = &DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions{empty: true}

func (r *DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardMetricDrilldownMetricLabelDrilldownsValueRestrictions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown struct {
	empty                bool                                                                      `json:"-"`
	ResourceLabels       []string                                                                  `json:"resourceLabels"`
	MetricLabels         []string                                                                  `json:"metricLabels"`
	MetadataSystemLabels []string                                                                  `json:"metadataSystemLabels"`
	MetadataUserLabels   []string                                                                  `json:"metadataUserLabels"`
	Reducer              *DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldownReducerEnum `json:"reducer"`
}

type jsonDashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown

func (r *DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown
	} else {

		r.ResourceLabels = res.ResourceLabels

		r.MetricLabels = res.MetricLabels

		r.MetadataSystemLabels = res.MetadataSystemLabels

		r.MetadataUserLabels = res.MetadataUserLabels

		r.Reducer = res.Reducer

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown *DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown = &DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown{empty: true}

func (r *DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardMetricDrilldownMetricGroupByDrilldown) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardGaugeView struct {
	empty      bool     `json:"-"`
	LowerBound *float64 `json:"lowerBound"`
	UpperBound *float64 `json:"upperBound"`
}

type jsonDashboardWidgetScorecardGaugeView DashboardWidgetScorecardGaugeView

func (r *DashboardWidgetScorecardGaugeView) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardGaugeView
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardGaugeView
	} else {

		r.LowerBound = res.LowerBound

		r.UpperBound = res.UpperBound

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardGaugeView is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardGaugeView *DashboardWidgetScorecardGaugeView = &DashboardWidgetScorecardGaugeView{empty: true}

func (r *DashboardWidgetScorecardGaugeView) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardGaugeView) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardSparkChartView struct {
	empty              bool                                                      `json:"-"`
	SparkChartType     *DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum `json:"sparkChartType"`
	MinAlignmentPeriod *string                                                   `json:"minAlignmentPeriod"`
}

type jsonDashboardWidgetScorecardSparkChartView DashboardWidgetScorecardSparkChartView

func (r *DashboardWidgetScorecardSparkChartView) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardSparkChartView
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardSparkChartView
	} else {

		r.SparkChartType = res.SparkChartType

		r.MinAlignmentPeriod = res.MinAlignmentPeriod

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardSparkChartView is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardSparkChartView *DashboardWidgetScorecardSparkChartView = &DashboardWidgetScorecardSparkChartView{empty: true}

func (r *DashboardWidgetScorecardSparkChartView) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardSparkChartView) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardThresholds struct {
	empty     bool                                             `json:"-"`
	Label     *string                                          `json:"label"`
	Value     *float64                                         `json:"value"`
	Color     *DashboardWidgetScorecardThresholdsColorEnum     `json:"color"`
	Direction *DashboardWidgetScorecardThresholdsDirectionEnum `json:"direction"`
}

type jsonDashboardWidgetScorecardThresholds DashboardWidgetScorecardThresholds

func (r *DashboardWidgetScorecardThresholds) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetScorecardThresholds
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetScorecardThresholds
	} else {

		r.Label = res.Label

		r.Value = res.Value

		r.Color = res.Color

		r.Direction = res.Direction

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardThresholds is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardThresholds *DashboardWidgetScorecardThresholds = &DashboardWidgetScorecardThresholds{empty: true}

func (r *DashboardWidgetScorecardThresholds) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardThresholds) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetText struct {
	empty   bool                           `json:"-"`
	Content *string                        `json:"content"`
	Format  *DashboardWidgetTextFormatEnum `json:"format"`
}

type jsonDashboardWidgetText DashboardWidgetText

func (r *DashboardWidgetText) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetText
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetText
	} else {

		r.Content = res.Content

		r.Format = res.Format

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetText is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetText *DashboardWidgetText = &DashboardWidgetText{empty: true}

func (r *DashboardWidgetText) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetText) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetBlank struct {
	empty bool `json:"-"`
}

type jsonDashboardWidgetBlank DashboardWidgetBlank

func (r *DashboardWidgetBlank) UnmarshalJSON(data []byte) error {
	var res jsonDashboardWidgetBlank
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDashboardWidgetBlank
	} else {

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetBlank is
// empty.  Go lacks global const objects, but this object should be treated
// as one.  Modifying this object will have undesirable results.
var EmptyDashboardWidgetBlank *DashboardWidgetBlank = &DashboardWidgetBlank{empty: true}

func (r *DashboardWidgetBlank) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetBlank) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Dashboard) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "monitoring",
		Type:    "Dashboard",
		Version: "monitoring",
	}
}

const DashboardMaxPage = -1

type DashboardList struct {
	Items []*Dashboard

	nextToken string

	pageSize int32

	project string
}

func (l *DashboardList) HasNext() bool {
	return l.nextToken != ""
}

func (l *DashboardList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listDashboard(ctx, l.project, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListDashboard(ctx context.Context, project string) (*DashboardList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	return c.ListDashboardWithMaxResults(ctx, project, DashboardMaxPage)

}

func (c *Client) ListDashboardWithMaxResults(ctx context.Context, project string, pageSize int32) (*DashboardList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	items, token, err := c.listDashboard(ctx, project, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &DashboardList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,

		project: project,
	}, nil
}

func (c *Client) GetDashboard(ctx context.Context, r *Dashboard) (*Dashboard, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	b, err := c.getDashboardRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalDashboard(b, c)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeDashboardNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteDashboard(ctx context.Context, r *Dashboard) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	if r == nil {
		return fmt.Errorf("Dashboard resource is nil")
	}
	c.Config.Logger.Info("Deleting Dashboard...")
	deleteOp := deleteDashboardOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllDashboard deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllDashboard(ctx context.Context, project string, filter func(*Dashboard) bool) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	listObj, err := c.ListDashboard(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllDashboard(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllDashboard(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyDashboard(ctx context.Context, rawDesired *Dashboard, opts ...dcl.ApplyOption) (*Dashboard, error) {
	c.Config.Logger.Info("Beginning ApplyDashboard...")
	c.Config.Logger.Infof("User specified desired state: %v", rawDesired)

	ctx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	initial, desired, diffs, err := c.dashboardDiffsForRawDesired(ctx, rawDesired, opts...)
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
	var ops []dashboardApiOperation
	if create {
		ops = append(ops, &createDashboardOperation{})
	} else if recreate {

		ops = append(ops, &deleteDashboardOperation{})

		ops = append(ops, &createDashboardOperation{})
		// We should re-canonicalize based on a nil existing resource.
		desired, err = canonicalizeDashboardDesiredState(rawDesired, nil)
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
	rawNew, err := c.GetDashboard(ctx, desired.urlNormalized())
	if err != nil {
		return nil, err
	}

	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createDashboardOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.Info("Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapDashboard(r, c)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeDashboardNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.Infof("Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeDashboardNewState(c, rawNew, rawDesired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.Infof("Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeDashboardDesiredState(rawDesired, newState)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffDashboard(c, newDesired, newState)
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
