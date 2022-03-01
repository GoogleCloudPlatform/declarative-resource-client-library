// Copyright 2022 Google LLC. All Rights Reserved.
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

type Dashboard struct {
	Name         *string                `json:"name"`
	DisplayName  *string                `json:"displayName"`
	GridLayout   *DashboardGridLayout   `json:"gridLayout"`
	MosaicLayout *DashboardMosaicLayout `json:"mosaicLayout"`
	RowLayout    *DashboardRowLayout    `json:"rowLayout"`
	ColumnLayout *DashboardColumnLayout `json:"columnLayout"`
	Project      *string                `json:"project"`
	Etag         *string                `json:"etag"`
}

func (r *Dashboard) String() string {
	return dcl.SprintResource(r)
}

// The enum DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum.
type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum string

// DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef returns a *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnumRef(s string) *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum {
	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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

// The enum DashboardWidgetXyChartDataSetsPlotTypeEnum.
type DashboardWidgetXyChartDataSetsPlotTypeEnum string

// DashboardWidgetXyChartDataSetsPlotTypeEnumRef returns a *DashboardWidgetXyChartDataSetsPlotTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartDataSetsPlotTypeEnumRef(s string) *DashboardWidgetXyChartDataSetsPlotTypeEnum {
	v := DashboardWidgetXyChartDataSetsPlotTypeEnum(s)
	return &v
}

func (v DashboardWidgetXyChartDataSetsPlotTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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

// The enum DashboardWidgetXyChartThresholdsColorEnum.
type DashboardWidgetXyChartThresholdsColorEnum string

// DashboardWidgetXyChartThresholdsColorEnumRef returns a *DashboardWidgetXyChartThresholdsColorEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetXyChartThresholdsColorEnumRef(s string) *DashboardWidgetXyChartThresholdsColorEnum {
	v := DashboardWidgetXyChartThresholdsColorEnum(s)
	return &v
}

func (v DashboardWidgetXyChartThresholdsColorEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartThresholdsDirectionEnum(s)
	return &v
}

func (v DashboardWidgetXyChartThresholdsDirectionEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartXAxisScaleEnum(s)
	return &v
}

func (v DashboardWidgetXyChartXAxisScaleEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartYAxisScaleEnum(s)
	return &v
}

func (v DashboardWidgetXyChartYAxisScaleEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetXyChartChartOptionsModeEnum(s)
	return &v
}

func (v DashboardWidgetXyChartChartOptionsModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterRankingMethodEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilterDirectionEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterRankingMethodEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum(s)
	return &v
}

func (v DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilterDirectionEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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

// The enum DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum.
type DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum string

// DashboardWidgetScorecardSparkChartViewSparkChartTypeEnumRef returns a *DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func DashboardWidgetScorecardSparkChartViewSparkChartTypeEnumRef(s string) *DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum {
	v := DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum(s)
	return &v
}

func (v DashboardWidgetScorecardSparkChartViewSparkChartTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardThresholdsColorEnum(s)
	return &v
}

func (v DashboardWidgetScorecardThresholdsColorEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetScorecardThresholdsDirectionEnum(s)
	return &v
}

func (v DashboardWidgetScorecardThresholdsDirectionEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
	v := DashboardWidgetTextFormatEnum(s)
	return &v
}

func (v DashboardWidgetTextFormatEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardGridLayout *DashboardGridLayout = &DashboardGridLayout{empty: true}

func (r *DashboardGridLayout) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardMosaicLayout *DashboardMosaicLayout = &DashboardMosaicLayout{empty: true}

func (r *DashboardMosaicLayout) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardMosaicLayoutTiles *DashboardMosaicLayoutTiles = &DashboardMosaicLayoutTiles{empty: true}

func (r *DashboardMosaicLayoutTiles) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardRowLayout *DashboardRowLayout = &DashboardRowLayout{empty: true}

func (r *DashboardRowLayout) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardRowLayoutRows *DashboardRowLayoutRows = &DashboardRowLayoutRows{empty: true}

func (r *DashboardRowLayoutRows) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardColumnLayout *DashboardColumnLayout = &DashboardColumnLayout{empty: true}

func (r *DashboardColumnLayout) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardColumnLayoutColumns *DashboardColumnLayoutColumns = &DashboardColumnLayoutColumns{empty: true}

func (r *DashboardColumnLayoutColumns) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidget *DashboardWidget = &DashboardWidget{empty: true}

func (r *DashboardWidget) Empty() bool {
	return r.empty
}

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
	empty             bool                                `json:"-"`
	DataSets          []DashboardWidgetXyChartDataSets    `json:"dataSets"`
	TimeshiftDuration *string                             `json:"timeshiftDuration"`
	Thresholds        []DashboardWidgetXyChartThresholds  `json:"thresholds"`
	XAxis             *DashboardWidgetXyChartXAxis        `json:"xAxis"`
	YAxis             *DashboardWidgetXyChartYAxis        `json:"yAxis"`
	ChartOptions      *DashboardWidgetXyChartChartOptions `json:"chartOptions"`
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

		r.TimeshiftDuration = res.TimeshiftDuration

		r.Thresholds = res.Thresholds

		r.XAxis = res.XAxis

		r.YAxis = res.YAxis

		r.ChartOptions = res.ChartOptions

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChart is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChart *DashboardWidgetXyChart = &DashboardWidgetXyChart{empty: true}

func (r *DashboardWidgetXyChart) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSets *DashboardWidgetXyChartDataSets = &DashboardWidgetXyChartDataSets{empty: true}

func (r *DashboardWidgetXyChartDataSets) Empty() bool {
	return r.empty
}

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

		r.UnitOverride = res.UnitOverride

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQuery is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQuery *DashboardWidgetXyChartDataSetsTimeSeriesQuery = &DashboardWidgetXyChartDataSetsTimeSeriesQuery{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQuery) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilter) Empty() bool {
	return r.empty
}

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
	empty              bool                                                                                            `json:"-"`
	AlignmentPeriod    *string                                                                                         `json:"alignmentPeriod"`
	PerSeriesAligner   *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum   `json:"perSeriesAligner"`
	CrossSeriesReducer *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum `json:"crossSeriesReducer"`
	GroupByFields      []string                                                                                        `json:"groupByFields"`
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

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) Empty() bool {
	return r.empty
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterAggregation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation struct {
	empty              bool                                                                                                     `json:"-"`
	AlignmentPeriod    *string                                                                                                  `json:"alignmentPeriod"`
	PerSeriesAligner   *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum   `json:"perSeriesAligner"`
	CrossSeriesReducer *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum `json:"crossSeriesReducer"`
	GroupByFields      []string                                                                                                 `json:"groupByFields"`
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

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) Empty() bool {
	return r.empty
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) HashCode() string {
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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatio) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumerator) Empty() bool {
	return r.empty
}

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
	empty              bool                                                                                                          `json:"-"`
	AlignmentPeriod    *string                                                                                                       `json:"alignmentPeriod"`
	PerSeriesAligner   *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum   `json:"perSeriesAligner"`
	CrossSeriesReducer *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum `json:"crossSeriesReducer"`
	GroupByFields      []string                                                                                                      `json:"groupByFields"`
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

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) Empty() bool {
	return r.empty
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) HashCode() string {
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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominator) Empty() bool {
	return r.empty
}

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
	empty              bool                                                                                                            `json:"-"`
	AlignmentPeriod    *string                                                                                                         `json:"alignmentPeriod"`
	PerSeriesAligner   *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum   `json:"perSeriesAligner"`
	CrossSeriesReducer *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum `json:"crossSeriesReducer"`
	GroupByFields      []string                                                                                                        `json:"groupByFields"`
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

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) Empty() bool {
	return r.empty
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation struct {
	empty              bool                                                                                                          `json:"-"`
	AlignmentPeriod    *string                                                                                                       `json:"alignmentPeriod"`
	PerSeriesAligner   *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum   `json:"perSeriesAligner"`
	CrossSeriesReducer *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum `json:"crossSeriesReducer"`
	GroupByFields      []string                                                                                                      `json:"groupByFields"`
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

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) Empty() bool {
	return r.empty
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) HashCode() string {
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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter = &DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{empty: true}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) Empty() bool {
	return r.empty
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetXyChartDataSetsTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) HashCode() string {
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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartThresholds *DashboardWidgetXyChartThresholds = &DashboardWidgetXyChartThresholds{empty: true}

func (r *DashboardWidgetXyChartThresholds) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartXAxis *DashboardWidgetXyChartXAxis = &DashboardWidgetXyChartXAxis{empty: true}

func (r *DashboardWidgetXyChartXAxis) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartYAxis *DashboardWidgetXyChartYAxis = &DashboardWidgetXyChartYAxis{empty: true}

func (r *DashboardWidgetXyChartYAxis) Empty() bool {
	return r.empty
}

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
	empty bool                                        `json:"-"`
	Mode  *DashboardWidgetXyChartChartOptionsModeEnum `json:"mode"`
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

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetXyChartChartOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetXyChartChartOptions *DashboardWidgetXyChartChartOptions = &DashboardWidgetXyChartChartOptions{empty: true}

func (r *DashboardWidgetXyChartChartOptions) Empty() bool {
	return r.empty
}

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

		r.GaugeView = res.GaugeView

		r.SparkChartView = res.SparkChartView

		r.Thresholds = res.Thresholds

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecard is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecard *DashboardWidgetScorecard = &DashboardWidgetScorecard{empty: true}

func (r *DashboardWidgetScorecard) Empty() bool {
	return r.empty
}

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

		r.UnitOverride = res.UnitOverride

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQuery is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQuery *DashboardWidgetScorecardTimeSeriesQuery = &DashboardWidgetScorecardTimeSeriesQuery{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQuery) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilter) Empty() bool {
	return r.empty
}

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
	empty              bool                                                                                      `json:"-"`
	AlignmentPeriod    *string                                                                                   `json:"alignmentPeriod"`
	PerSeriesAligner   *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationPerSeriesAlignerEnum   `json:"perSeriesAligner"`
	CrossSeriesReducer *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregationCrossSeriesReducerEnum `json:"crossSeriesReducer"`
	GroupByFields      []string                                                                                  `json:"groupByFields"`
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

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) Empty() bool {
	return r.empty
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterAggregation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation struct {
	empty              bool                                                                                               `json:"-"`
	AlignmentPeriod    *string                                                                                            `json:"alignmentPeriod"`
	PerSeriesAligner   *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationPerSeriesAlignerEnum   `json:"perSeriesAligner"`
	CrossSeriesReducer *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregationCrossSeriesReducerEnum `json:"crossSeriesReducer"`
	GroupByFields      []string                                                                                           `json:"groupByFields"`
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

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) Empty() bool {
	return r.empty
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterSecondaryAggregation) HashCode() string {
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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterPickTimeSeriesFilter) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatio) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumerator) Empty() bool {
	return r.empty
}

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
	empty              bool                                                                                                    `json:"-"`
	AlignmentPeriod    *string                                                                                                 `json:"alignmentPeriod"`
	PerSeriesAligner   *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationPerSeriesAlignerEnum   `json:"perSeriesAligner"`
	CrossSeriesReducer *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregationCrossSeriesReducerEnum `json:"crossSeriesReducer"`
	GroupByFields      []string                                                                                                `json:"groupByFields"`
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

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) Empty() bool {
	return r.empty
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioNumeratorAggregation) HashCode() string {
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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominator) Empty() bool {
	return r.empty
}

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
	empty              bool                                                                                                      `json:"-"`
	AlignmentPeriod    *string                                                                                                   `json:"alignmentPeriod"`
	PerSeriesAligner   *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationPerSeriesAlignerEnum   `json:"perSeriesAligner"`
	CrossSeriesReducer *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregationCrossSeriesReducerEnum `json:"crossSeriesReducer"`
	GroupByFields      []string                                                                                                  `json:"groupByFields"`
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

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) Empty() bool {
	return r.empty
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioDenominatorAggregation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation struct {
	empty              bool                                                                                                    `json:"-"`
	AlignmentPeriod    *string                                                                                                 `json:"alignmentPeriod"`
	PerSeriesAligner   *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationPerSeriesAlignerEnum   `json:"perSeriesAligner"`
	CrossSeriesReducer *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregationCrossSeriesReducerEnum `json:"crossSeriesReducer"`
	GroupByFields      []string                                                                                                `json:"groupByFields"`
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

	}
	return nil
}

// This object is used to assert a desired state where this DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) Empty() bool {
	return r.empty
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioSecondaryAggregation) HashCode() string {
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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter = &DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter{empty: true}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) Empty() bool {
	return r.empty
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) String() string {
	return dcl.SprintResource(r)
}

func (r *DashboardWidgetScorecardTimeSeriesQueryTimeSeriesFilterRatioPickTimeSeriesFilter) HashCode() string {
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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardGaugeView *DashboardWidgetScorecardGaugeView = &DashboardWidgetScorecardGaugeView{empty: true}

func (r *DashboardWidgetScorecardGaugeView) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardSparkChartView *DashboardWidgetScorecardSparkChartView = &DashboardWidgetScorecardSparkChartView{empty: true}

func (r *DashboardWidgetScorecardSparkChartView) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetScorecardThresholds *DashboardWidgetScorecardThresholds = &DashboardWidgetScorecardThresholds{empty: true}

func (r *DashboardWidgetScorecardThresholds) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetText *DashboardWidgetText = &DashboardWidgetText{empty: true}

func (r *DashboardWidgetText) Empty() bool {
	return r.empty
}

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
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDashboardWidgetBlank *DashboardWidgetBlank = &DashboardWidgetBlank{empty: true}

func (r *DashboardWidgetBlank) Empty() bool {
	return r.empty
}

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

func (r *Dashboard) ID() (string, error) {
	if err := extractDashboardFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":         dcl.ValueOrEmptyString(nr.Name),
		"displayName":  dcl.ValueOrEmptyString(nr.DisplayName),
		"gridLayout":   dcl.ValueOrEmptyString(nr.GridLayout),
		"mosaicLayout": dcl.ValueOrEmptyString(nr.MosaicLayout),
		"rowLayout":    dcl.ValueOrEmptyString(nr.RowLayout),
		"columnLayout": dcl.ValueOrEmptyString(nr.ColumnLayout),
		"project":      dcl.ValueOrEmptyString(nr.Project),
		"etag":         dcl.ValueOrEmptyString(nr.Etag),
	}
	return dcl.Nprintf("projects/{{project}}/dashboards/{{name}}", params), nil
}

const DashboardMaxPage = -1

type DashboardList struct {
	Items []*Dashboard

	nextToken string

	pageSize int32

	resource *Dashboard
}

func (l *DashboardList) HasNext() bool {
	return l.nextToken != ""
}

func (l *DashboardList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listDashboard(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListDashboard(ctx context.Context, project string) (*DashboardList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListDashboardWithMaxResults(ctx, project, DashboardMaxPage)

}

func (c *Client) ListDashboardWithMaxResults(ctx context.Context, project string, pageSize int32) (*DashboardList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Dashboard{
		Project: &project,
	}
	items, token, err := c.listDashboard(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &DashboardList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetDashboard(ctx context.Context, r *Dashboard) (*Dashboard, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractDashboardFields(r)

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
	result, err := unmarshalDashboard(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeDashboardNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractDashboardFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteDashboard(ctx context.Context, r *Dashboard) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Dashboard resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Dashboard...")
	deleteOp := deleteDashboardOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllDashboard deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllDashboard(ctx context.Context, project string, filter func(*Dashboard) bool) error {
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
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Dashboard
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyDashboardHelper(c, ctx, rawDesired, opts...)
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

func applyDashboardHelper(c *Client, ctx context.Context, rawDesired *Dashboard, opts ...dcl.ApplyOption) (*Dashboard, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyDashboard...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractDashboardFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.dashboardDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToDashboardDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
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
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
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
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyDashboardDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyDashboardDiff(c *Client, ctx context.Context, desired *Dashboard, rawDesired *Dashboard, ops []dashboardApiOperation, opts ...dcl.ApplyOption) (*Dashboard, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
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

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapDashboard(r, c, rawDesired)
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

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeDashboardNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeDashboardDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractDashboardFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractDashboardFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffDashboard(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
